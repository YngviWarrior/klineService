package job

import (
	"context"
	"klineService/database"
	"klineService/database/repositories/mysql"
	"klineService/services"
)

type Job struct {
	Test bool

	BinanceInterface services.BinanceInterface
	ByBitInterface   services.ByBitInterface
	RabbitMQ         services.RabbitMQInterface
	Redis            services.RedisInterface
	DiscordInterface services.DiscordInterface

	KlineRepo mysql.KlineRepositoryInterface
	AssetRepo mysql.AssetRepositoryInterface
}

type JobInterface interface {
	AliveNotify(db *database.Database, loopChannel *chan bool)
	AveragePrices(db *database.Database, loopChannel *chan bool)
	AssetManager(cancel context.CancelFunc, db *database.Database, assetManagerLoopChannel *chan bool, quitLoopChannel *chan bool)
	LiveStream(ctx context.Context, db *database.Database, quitChannel *chan bool)
	SyncKlineTable(db *database.Database)
}
