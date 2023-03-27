package job

import (
	"fmt"
	"klineService/database"
	"klineService/database/repositories/mysql"
	"klineService/services"
	discordstructs "klineService/services/discordStructs"
	"time"
)

func AliveNotify(db *database.Database, aliveLoopChannel *chan bool) {
	var klineRepo mysql.KlineRepositoryInterface = &mysql.KlineRepository{}

	conn := db.CreateConnection()

	lastBtc := klineRepo.FindLastPrice(nil, conn, 3, 1, 2)
	lastEth := klineRepo.FindLastPrice(nil, conn, 4, 1, 2)
	var DiscordInterface services.DiscordInterface = &services.Discord{}

	var notify discordstructs.Notification
	notify.Channel = "Health"
	notify.Content = fmt.Sprintf("(%v) KlineService is Healthy! (BTCUSDT: %v, ETHUSDT: %v)", time.Now().Format("2006-01-02 15:04:05"), lastBtc, lastEth)
	DiscordInterface.SendNotification(&notify)

	conn.Close()

	time.Sleep(time.Minute * 5)
	*aliveLoopChannel <- true
}
