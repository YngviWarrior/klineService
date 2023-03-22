package services

import (
	"klineService/database"
	"klineService/entities/asset"
	"klineService/entities/kline"
	bBR "klineService/services/byBitStructs"
	dR "klineService/services/discordStructs"
	rabbitmqstructs "klineService/services/rabbitMQStructs"
)

type ByBit struct {
	TestAPI bool
}

type Binance struct {
	TestAPI bool
}

type Discord struct{}

type RabbitMQ struct{}

type RabbitMQInterface interface {
	SendCotation(kline *kline.Kline)
	SendAveragePrice(avgMessage *rabbitmqstructs.InputAvgMessageDto)
}

type DiscordInterface interface {
	SendNotification(params *dR.Notification)
}

type ByBitInterface interface {
	ServerTimestamp() (response bBR.GetServerTimestamp)
	LiveKlines(db *database.Database, parities []*asset.Asset)
	GetKlines(params *bBR.GetKlinesParams) (response bBR.GetKlinesResponse)
}

type BinanceInterface interface {
	LiveKlines(db *database.Database, parities []*asset.Asset)
}
