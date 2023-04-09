package main

import (
	"context"
	"klineService/database"
	"klineService/database/repositories/mysql"
	"klineService/job"
	"klineService/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	var j job.Job

	switch os.Getenv("ENVIROMENT") {
	case "dev":
		j.Test = true
	}

	var db database.Database
	db.CreatePool(50)

	//repositories
	var assetRepo mysql.AssetRepository
	var klineRepo mysql.KlineRepository

	//services
	var re services.Redis
	re.Database = &db
	re.AssetRepo = &assetRepo
	re.InitCache()

	var r services.RabbitMQ
	var d services.Discord

	var b services.Binance
	b.TestAPI = j.Test
	b.KlineRepoInterface = &klineRepo
	b.RabbitMQ = &r
	b.Redis = &re

	var bb services.ByBit
	bb.TestAPI = j.Test
	bb.KlineRepoInterface = &klineRepo
	bb.RabbitMQ = &r
	bb.Redis = &re

	//job attribuition
	j.AssetRepo = &assetRepo
	j.KlineRepo = &klineRepo

	j.ByBitInterface = &bb
	j.BinanceInterface = &b
	j.RabbitMQ = &r
	j.Redis = &re
	j.DiscordInterface = &d

	var jobInterface job.JobInterface = &j

	quitLoopChannel := make(chan bool)
	assetManageLoopChannel := make(chan bool)
	avgLoopChannel := make(chan bool)
	aliveLoopChannel := make(chan bool)

	// Start Functionalities
	// time.Sleep(time.Second)
	jobInterface.SyncKlineTable(&db)

	ctx, cancel := context.WithCancel(context.Background())
	go jobInterface.AssetManager(cancel, &db, &assetManageLoopChannel, &quitLoopChannel)
	go jobInterface.LiveStream(ctx, &db, &quitLoopChannel)

	go jobInterface.AveragePrices(&db, &avgLoopChannel)
	go jobInterface.AliveNotify(&db, &aliveLoopChannel)
	log.Println("Kline Service is Running.")

	for {
		select {
		case <-avgLoopChannel:
			go jobInterface.AveragePrices(&db, &avgLoopChannel)
		case <-aliveLoopChannel:
			go jobInterface.AliveNotify(&db, &aliveLoopChannel)
		case <-assetManageLoopChannel:
			go jobInterface.AssetManager(cancel, &db, &assetManageLoopChannel, &quitLoopChannel)
		case <-quitLoopChannel:
			ctx, cancel = context.WithCancel(context.Background())
			go jobInterface.LiveStream(ctx, &db, &quitLoopChannel)
		}
	}
}
