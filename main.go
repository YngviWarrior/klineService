package main

import (
	"klineService/database"
	"klineService/job"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}
	// var DiscordInterface services.DiscordInterface = &services.Discord{}

	var db database.Database
	db.CreatePool(50)

	go job.LiveStream(&db)

	loopChannel := make(chan bool)
	aliveLoopChannel := make(chan bool)

	go job.AveragePrices(&db, &loopChannel)
	go job.AliveNotify(&loopChannel)

	for {
		select {
		case <-loopChannel:
			time.Sleep(time.Second * 60)
			go job.AveragePrices(&db, &loopChannel)
		case <-aliveLoopChannel:
			time.Sleep(time.Minute * 5)
			go job.AliveNotify(&aliveLoopChannel)
		}
	}
}
