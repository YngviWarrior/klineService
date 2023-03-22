package services

import (
	"fmt"
	database "klineService/database"
	repository "klineService/database/repositories/mysql"
	asset "klineService/entities/asset"
	kline "klineService/entities/kline"
	"klineService/utils"
	"log"

	"github.com/adshao/go-binance/v2"
)

var binanceBaseUrl string = "https://api.binance.com"

func (s *Binance) LiveKlines(db *database.Database, parities []*asset.Asset) {
	var klineRepoInterface repository.KlineRepositoryInterface = &repository.KlineRepository{}
	var rabbitmq RabbitMQInterface = &RabbitMQ{}

	klines := make(map[string]string)
	for _, v := range parities {
		if v.Symbol == "BRL" {
			klines[fmt.Sprintf("USDT%s", v.Symbol)] = "1m"
		} else if v.Symbol != "USDT" {
			klines[fmt.Sprintf("%sUSDT", v.Symbol)] = "1m"
		}
	}

	wsKlineHandler := func(info *binance.WsKlineEvent) {
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

		k.Exchange = 1
		k.Mts = uint64(info.Kline.StartTime)
		k.Open = utils.ParseFloat(info.Kline.Open)
		k.Close = utils.ParseFloat(info.Kline.Close)
		k.High = utils.ParseFloat(info.Kline.High)
		k.Low = utils.ParseFloat(info.Kline.Low)
		k.Volume = utils.ParseFloat(info.Kline.Volume)

		go rabbitmq.SendCotation(&k)

		if !klineRepoInterface.CreateDirect(db.Pool, &k) {
			return
		}
	}

	errHandler := func(err error) {
		log.Panic("Socket Err 01: ", err)
	}

	_, _, err := binance.WsCombinedKlineServe(klines, wsKlineHandler, errHandler)

	if err != nil {
		log.Panic("Socket Err 02: ", err)
		return
	}
}
