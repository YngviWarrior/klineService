package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	database "klineService/database"
	repository "klineService/database/repositories/mysql"
	asset "klineService/entities/asset"
	kline "klineService/entities/kline"
	bBR "klineService/services/byBitStructs"
	bybitstructs "klineService/services/byBitStructs"
	"klineService/services/bybit-api/ws"
	"klineService/utils"
	"log"
	"net/http"
)

const byBitBaseUrlTest = "https://api-testnet.bybit.com"
const byBitBaseUrl = "https://api.bybit.com"
const byBitBaseUrl2 = "https://api.bytick.com"

func (s *ByBit) ServerTimestamp() (response bBR.GetServerTimestamp) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.bybit.com/spot/v1/time", nil)

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

func (s *ByBit) GetKlines(symbol string, resolution string, start, end, limit int64) (list []*bybitstructs.Kline, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.bybit.com/spot/quote/v1/kline", nil)

	if err != nil {
		log.Println("Req klines prepare: ", err)
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

	q.Add("interval", resolution)
	q.Add("symbol", symbol)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Req klines exec: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var res bybitstructs.KlineResponse
	err = json.Unmarshal(bodyBytes, &res)

	if err != nil {
		log.Println("json conv: ", err)
	} else {
		for _, v := range res.Result {
			var k bybitstructs.Kline
			for i, w := range v {
				if i == 0 {
					k.Open_time = fmt.Sprintf("%10.0f", w)
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

			list = append(list, &k)
		}
	}

	defer resp.Body.Close()

	return
}

func (s *ByBit) LiveKlines(db *database.Database, parities []*asset.Asset) {
	var KlineRepoInterface repository.KlineRepositoryInterface = &repository.KlineRepository{}
	var rabbitmq RabbitMQInterface = &RabbitMQ{}

	url := "wss://stream.bybit.com/spot/public/v3"

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

	for _, v := range parities {
		if v.Symbol == "BRL" {
			b.Subscribe(fmt.Sprintf("kline.1m.USDT%s", v.Symbol))
		} else if v.Symbol != "USDT" {
			b.Subscribe(fmt.Sprintf("kline.1m.%sUSDT", v.Symbol))
		}
	}

	b.On("kline", func(symbol string, info ws.KLine) {
		var k kline.Kline

		switch info.Symbol {
		case "BTCUSDT":
			k.Asset = 3
			k.AssetSymbol = "BTC"

			k.AssetQuote = 1
			k.AssetQuoteSymbol = "USDT"
		case "ETHUSDT":
			k.Asset = 4
			k.AssetSymbol = "ETH"

			k.AssetQuote = 1
			k.AssetQuoteSymbol = "USDT"
		case "USDTBRL":
			k.Asset = 1
			k.AssetSymbol = "USDT"

			k.AssetQuote = 2
			k.AssetQuoteSymbol = "BRL"
		}

		k.Mts = uint64(info.OpenTime)
		k.Exchange = 2
		k.Open = utils.ParseFloat(info.Open)
		k.Close = utils.ParseFloat(info.Close)
		k.High = utils.ParseFloat(info.High)
		k.Low = utils.ParseFloat(info.Low)
		k.Volume = utils.ParseFloat(info.Volume)

		go rabbitmq.SendCotation(&k)

		if !KlineRepoInterface.CreateDirect(db.Pool, &k) {
			return
		}
	})
}
