package services

import (
	"context"
	"klineService/database"
	"klineService/database/repositories/mysql"
	repository "klineService/database/repositories/mysql"
	"klineService/entities/kline"
	bybitstructs "klineService/services/byBitStructs"
	dR "klineService/services/discordStructs"
	rabbitmqstructs "klineService/services/rabbitMQStructs"

	"github.com/go-redis/redis/v8"
)

type ByBit struct {
	TestAPI bool

	KlineRepoInterface repository.KlineRepositoryInterface
	RabbitMQ           RabbitMQInterface
	Redis              RedisInterface
}

type Binance struct {
	TestAPI bool

	KlineRepoInterface repository.KlineRepositoryInterface
	RabbitMQ           RabbitMQInterface
	Redis              RedisInterface
}

type Discord struct{}

type RabbitMQ struct{}

type Redis struct {
	Client *redis.Client

	Database  *database.Database
	AssetRepo mysql.AssetRepositoryInterface
}

type RedisInterface interface {
	InitCache()
	GetInstance() *redis.Client
	GetCache(key, primitiveType string) (val any)
}

type RabbitMQInterface interface {
	SendCotation(kline *kline.Kline)
	SendAveragePrice(avgMessage *rabbitmqstructs.InputAvgMessageDto)
}

type DiscordInterface interface {
	SendNotification(params *dR.Notification)
}

type ByBitInterface interface {
	ServerTimestamp() (response bybitstructs.GetServerTimestamp)
	LiveKlines(ctx context.Context, db *database.Database, quitChannel *chan bool)
	GetKlines(category, symbol string, resolution string, start, end, limit int64) (list []*bybitstructs.Kline, err error)
}

type BinanceInterface interface {
	LiveKlines(ctx context.Context, db *database.Database, quitChannel *chan bool)
}

// Global usage.
var CacheInstance Redis
