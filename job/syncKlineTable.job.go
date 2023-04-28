package job

import (
	"encoding/json"
	"fmt"
	"klineService/database"
	"klineService/entities/asset"
	"klineService/entities/kline"
	"log"
	"strconv"
	"strings"
	"time"
)

func (j *Job) SyncKlineTable(db *database.Database) {
	assets := j.Redis.GetCache("Assets", "string")

	var cachedSlice []*asset.Asset
	err := json.Unmarshal([]byte(assets.(string)), &cachedSlice)

	if err != nil {
		log.Panic("Bybit liveKlines unMarshal: ", err)
	}

	log.Println("Syncronizing Klines")

	for _, a := range cachedSlice {
		conn := db.CreateConnection()
		fMts := j.KlineRepo.FindFirstMts(nil, conn, int64(a.Asset), 1, 2, j.Test)
		conn.Close()

		if (fMts == kline.Kline{}) {
			continue
		}

		start := time.UnixMilli(int64(fMts.Mts) / 1000)
		end := start.Add(time.Hour * 5)

		var count int64
		for end.Before(time.Now()) {
			j.request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, start.UnixMilli(), end.UnixMilli(), 0)
			count++

			time.Sleep(time.Second)
			end = end.Add(time.Hour * 5)

			if !end.Before(time.Now()) {
				j.request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, 0, end.UnixMilli(), 0)
				count++
			}
		}

		log.Printf("%sUSDT Klines Syncronization Looped %d times", strings.ToUpper(a.Symbol), count)
	}

	log.Println("Klines Syncronization is Finished")
}

func (j *Job) request(db *database.Database, symbol string, asset uint64, start, end, limit int64) {
	resp, _ := j.ByBitInterface.GetKlines("spot", symbol, "1", start, end, limit)

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

		j.KlineRepo.CreateDirect(db.Pool, &c)
	}
}
