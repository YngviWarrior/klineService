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
	"sync"
	"time"
)

func (j *Job) SyncKlineTable(db *database.Database) {
	assets := j.Redis.GetCache("Assets", "string")
	if assets == nil {
		log.Panic("No assets found")
	}

	var cachedSlice []*asset.Asset
	err := json.Unmarshal([]byte(assets.(string)), &cachedSlice)
	if err != nil {
		log.Panic("BBSKT Unmarshal: ", err)
	}

	log.Println("Syncronizing Klines")
	var wg sync.WaitGroup
	wg.Add(len(cachedSlice))

	for _, a := range cachedSlice {
		go func(a *asset.Asset, wg *sync.WaitGroup) {
			conn := db.CreateConnection()

			fMts := j.KlineRepo.FindFirstMts(nil, conn, int64(a.Asset), 1, 2, j.Test)
			conn.Close()

			if (fMts == kline.Kline{}) {
				wg.Done()
				return
			}
			// fmt.Println(a)
			start := time.UnixMilli(int64(fMts.Mts) / 1000)
			end := start.Add(time.Hour * 5)

			var count int64
			for start.Before(time.Now()) {
				j.request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, start.UnixMilli(), end.UnixMilli(), 0)
				count++
				// fmt.Println(count)
				time.Sleep(time.Second)
				start = start.Add(time.Hour * 5)
				end = end.Add(time.Hour * 5)
				// fmt.Println(start, end)
				if !start.Before(time.Now()) {
					j.request(db, fmt.Sprintf("%sUSDT", a.Symbol), a.Asset, 0, end.UnixMilli(), 0)
					count++
				}
			}

			log.Printf("%sUSDT Klines Syncronization Looped %d times \n", strings.ToUpper(a.Symbol), count)
			wg.Done()
		}(a, &wg)
	}

	wg.Wait()
	log.Println("Klines Syncronization is Finished")
}

func (j *Job) request(db *database.Database, symbol string, asset uint64, start, end, limit int64) {
	resp, err := j.ByBitInterface.GetKlines("spot", symbol, "1", start, end, limit)

	if len(resp) == 0 {
		log.Println("SKTR 01: ", err)
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
