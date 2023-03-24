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

	var db database.Database
	db.CreatePool(50)

	job.SyncKlineTable(&db)
	go job.LiveStream(&db)

	avgLoopChannel := make(chan bool)
	aliveLoopChannel := make(chan bool)

	go job.AveragePrices(&db, &avgLoopChannel)
	go job.AliveNotify(&db, &aliveLoopChannel)

	for {
		select {
		case <-avgLoopChannel:
			time.Sleep(time.Second * 15)
			go job.AveragePrices(&db, &avgLoopChannel)
		case <-aliveLoopChannel:
			time.Sleep(time.Minute * 5)
			go job.AliveNotify(&db, &aliveLoopChannel)
		}
	}
}
