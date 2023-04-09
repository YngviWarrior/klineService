package services

import (
	"context"
	"encoding/json"
	"fmt"
	database "klineService/database"
	asset "klineService/entities/asset"
	kline "klineService/entities/kline"
	"klineService/utils"
	"log"

	"github.com/adshao/go-binance/v2"
)

func (s *Binance) LiveKlines(ctx context.Context, db *database.Database, quitChannel *chan bool) {
	assets := s.Redis.GetCache("Assets", "string")

	var cachedSlice []*asset.Asset
	err := json.Unmarshal([]byte(assets.(string)), &cachedSlice)

	if err != nil {
		log.Panic("Binance liveKlines unMarshal 01: ", err)
	}

	klines := make(map[string]string)
	for _, v := range cachedSlice {
		if v.Symbol == "BRL" {
			klines[fmt.Sprintf("USDT%s", v.Symbol)] = "1m"
		} else if v.Symbol != "USDT" {
			klines[fmt.Sprintf("%sUSDT", v.Symbol)] = "1m"
		}
	}

	wsKlineHandler := func(info *binance.WsKlineEvent) {
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

		k.Exchange = 1
		k.Mts = uint64(info.Kline.StartTime)
		k.Open = utils.ParseFloat(info.Kline.Open)
		k.Close = utils.ParseFloat(info.Kline.Close)
		k.High = utils.ParseFloat(info.Kline.High)
		k.Low = utils.ParseFloat(info.Kline.Low)
		k.Volume = utils.ParseFloat(info.Kline.Volume)
		k.TestNet = s.TestAPI

		go s.RabbitMQ.SendCotation(&k)

		if !s.KlineRepoInterface.CreateDirect(db.Pool, &k) {
			return
		}
	}

	errHandler := func(err error) {
		log.Panic("Socket Err 01: ", err)
	}

	_, stop, err := binance.WsCombinedKlineServe(klines, wsKlineHandler, errHandler)

	if err != nil {
		log.Panic("Socket Err 02: ", err)
		return
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				stop <- struct{}{}
			}
		}
	}()
}
