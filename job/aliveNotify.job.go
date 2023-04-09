package job

import (
	"fmt"
	"klineService/database"
	discordstructs "klineService/services/discordStructs"
	"time"
)

func (j *Job) AliveNotify(db *database.Database, aliveLoopChannel *chan bool) {
	conn := db.CreateConnection()

	lastBtc := j.KlineRepo.FindLastPrice(nil, conn, 3, 1, 2)
	lastEth := j.KlineRepo.FindLastPrice(nil, conn, 4, 1, 2)

	var notify discordstructs.Notification
	notify.Channel = "Health"
	notify.Content = fmt.Sprintf("(%v) KlineService is Healthy! (BTCUSDT: %v, ETHUSDT: %v)", time.Now().Format("2006-01-02 15:04:05"), lastBtc, lastEth)
	j.DiscordInterface.SendNotification(&notify)

	conn.Close()

	time.Sleep(time.Minute * 5)
	*aliveLoopChannel <- true
}
