package job

import (
	"fmt"
	"klineService/database"
	"klineService/database/repositories/mysql"
	"klineService/entities/kline"
	"klineService/services"
	"strconv"
	"time"
)

var bybitService services.ByBitInterface = &services.ByBit{}
var klineRepo mysql.KlineRepositoryInterface = &mysql.KlineRepository{}
var assetRepo mysql.AssetRepositoryInterface = &mysql.AssetRepository{}

func SyncKlineTable(db *database.Database) {

	conn := db.CreateConnection()
	asset := assetRepo.List(nil, conn)
	conn.Close()

	fmt.Println("Syncronizing Klines")

	for _, a := range asset {
		conn := db.CreateConnection()
		fMts := klineRepo.FindFirstMts(nil, conn, int64(a.Asset), 1, 2)
		conn.Close()

		if (fMts == kline.Kline{}) {
			continue
		}

		start := time.UnixMilli(int64(fMts.Mts))
		end := start.Add(time.Hour)

		for end.Before(time.Now()) {
			request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, start.Unix(), end.Unix(), 0)
			time.Sleep(time.Second)
			end = end.Add(time.Hour)
		}

		end = end.Add(time.Hour)
		request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, 0, end.Unix(), 0)
	}

	fmt.Println("Klines Syncronization is Finished")
}

func request(db *database.Database, symbol string, asset uint64, start, end, limit int64) {
	resp, _ := bybitService.GetKlines(symbol, "1m", start, end, limit)

	if len(resp) == 0 {
		return
	}

	for _, v := range resp {
		var c kline.Kline
		c.Asset = asset
		c.AssetQuote = 1
		c.Exchange = 2

		mts, _ := strconv.Atoi(v.Open_time)

		c.Mts = uint64(mts)
		c.Open, _ = strconv.ParseFloat(v.Open, 64)
		c.Close, _ = strconv.ParseFloat(v.Close, 64)
		c.High, _ = strconv.ParseFloat(v.High, 64)
		c.Low, _ = strconv.ParseFloat(v.Low, 64)
		c.Volume, _ = strconv.ParseFloat(v.Volume, 64)

		klineRepo.CreateDirect(db.Pool, &c)
	}
}
