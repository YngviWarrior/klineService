package main

import (
	"klineService/database"
	"klineService/job"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	avgLoopChannel := make(chan bool)
	aliveLoopChannel := make(chan bool)

	var db database.Database
	db.CreatePool(50)

	job.SyncKlineTable(&db)
	go job.LiveStream(&db)

	go job.AveragePrices(&db, &avgLoopChannel)
	go job.AliveNotify(&db, &aliveLoopChannel)
	log.Println("Kline Service is Running.")

	for {
		select {
		case <-avgLoopChannel:
			go job.AveragePrices(&db, &avgLoopChannel)
		case <-aliveLoopChannel:
			go job.AliveNotify(&db, &aliveLoopChannel)
		}
	}
}
