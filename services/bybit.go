package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	database "klineService/database"
	asset "klineService/entities/asset"
	kline "klineService/entities/kline"
	bybitstructs "klineService/services/byBitStructs"
	"klineService/services/bybit-api/ws"
	"klineService/utils"
	"log"
	"net/http"
)

const byBitBaseUrlTest = "https://api-testnet.bybit.com"
const byBitBaseUrl = "https://api.bybit.com"

const byBitWSSUrl = "wss://stream.bybit.com/spot/public/v3"
const byBitWSSUrlTest = "wss://stream-testnet.bybit.com/spot/public/v3"

func (s *ByBit) ServerTimestamp() (response bybitstructs.GetServerTimestamp) {
	client := &http.Client{}
	var req *http.Request
	var err error

	switch s.TestAPI {
	case true:
		req, err = http.NewRequest("GET", byBitBaseUrlTest+"/spot/v1/time", nil)
	default:
		req, err = http.NewRequest("GET", byBitBaseUrl+"/spot/v1/time", nil)
	}

	if err != nil {
		log.Println("BBST 01: ", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Req klines exec: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Println("BBST 02: ", err)
	}

	defer resp.Body.Close()

	return
}

func (s *ByBit) GetKlines(category, symbol string, resolution string, start, end, limit int64) (list []*bybitstructs.Kline, err error) {
	client := &http.Client{}
	var req *http.Request

	switch s.TestAPI {
	case true:
		req, err = http.NewRequest("GET", byBitBaseUrlTest+"/v5/market/kline", nil)
	default:
		req, err = http.NewRequest("GET", byBitBaseUrl+"/v5/market/kline", nil)
	}

	if err != nil {
		log.Panic("BBGK 01: ", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()

	if limit > 0 {
		q.Add("limit", fmt.Sprintf("%v", limit))
	}

	if end > 0 {
		q.Add("endTime", fmt.Sprintf("%v", end))
	}

	if start > 0 {
		q.Add("startTime", fmt.Sprintf("%v", start))
	}

	q.Add("category", category)
	q.Add("interval", resolution)
	q.Add("symbol", symbol)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Panic("BBGK 02: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var res bybitstructs.KlineResponse
	err = json.Unmarshal(bodyBytes, &res)

	if err != nil || res.RetMsg != "OK" {
		log.Panicf("BBGK 03: %v, %s", err, res.RetMsg)
	} else {
		for _, v := range res.Result.List {
			// fmt.Println(v)
			var k bybitstructs.Kline
			for i, w := range v {
				if i == 0 {
					k.Open_time = fmt.Sprintf("%v", w)
				}
				if i == 1 {
					k.Open = fmt.Sprintf("%v", w)
				}
				if i == 2 {
					k.High = fmt.Sprintf("%v", w)
				}
				if i == 3 {
					k.Low = fmt.Sprintf("%v", w)
				}
				if i == 4 {
					k.Close = fmt.Sprintf("%v", w)
				}
				if i == 5 {
					k.Volume = fmt.Sprintf("%v", w)
				}
			}
			// fmt.Println(k)
			list = append(list, &k)
		}
	}

	defer resp.Body.Close()

	return
}

func (s *ByBit) LiveKlines(ctx context.Context, db *database.Database, quitChannel *chan bool) {
	var url = byBitWSSUrl
	switch s.TestAPI {
	case true:
		url = byBitWSSUrlTest
	}

	cfg := &ws.Configuration{
		Addr:          url,
		AutoReconnect: true,
		DebugMode:     false,
	}

	b := ws.New(cfg)
	err := b.Start()

	if err != nil {
		fmt.Println(err)
	}

	assets := s.Redis.GetCache("Assets", "string")

	var cachedSlice []*asset.Asset
	err = json.Unmarshal([]byte(assets.(string)), &cachedSlice)

	if err != nil {
		log.Panic("Bybit liveKlines unMarshal: ", err)
	}

	for _, v := range cachedSlice {
		if v.Symbol == "BRL" {
			b.Subscribe(fmt.Sprintf("kline.1m.USDT%s", v.Symbol))
		} else if v.Symbol != "USDT" {
			b.Subscribe(fmt.Sprintf("kline.1m.%sUSDT", v.Symbol))
		}
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				b.Close()
			}
		}
	}()

	b.On("kline", func(symbol string, info ws.KLine) {
		var k kline.Kline

		k.AssetQuote = 1
		k.AssetQuoteSymbol = "USDT"

		for _, asset := range cachedSlice {
			if info.Symbol[:3] == asset.Symbol {
				k.Asset = asset.Asset
				k.AssetSymbol = asset.Symbol
			}

			if info.Symbol[:4] == asset.Symbol {
				k.Asset = asset.Asset
				k.AssetSymbol = asset.Symbol
				k.AssetQuote = 2
				k.AssetQuoteSymbol = "BRL"
			}
		}

		k.Mts = uint64(info.OpenTime)
		k.Exchange = 2
		k.Open = utils.ParseFloat(info.Open)
		k.Close = utils.ParseFloat(info.Close)
		k.High = utils.ParseFloat(info.High)
		k.Low = utils.ParseFloat(info.Low)
		k.Volume = utils.ParseFloat(info.Volume)
		k.TestNet = s.TestAPI

		go s.RabbitMQ.SendCotation(&k)

		if !s.KlineRepoInterface.CreateDirect(db.Pool, &k) {
			return
		}
	})
}
