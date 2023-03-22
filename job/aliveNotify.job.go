package job

import (
	"fmt"
	"klineService/services"
	discordstructs "klineService/services/discordStructs"
	"time"
)

func AliveNotify(aliveLoopChannel *chan bool) {
	var DiscordInterface services.DiscordInterface = &services.Discord{}

	var send discordstructs.Notification
	send.Channel = "Health"

	send.Content = fmt.Sprintf("%v <-----> Kline Sockets Online", time.Now().Format("2006-01-02 15:04:05"))

	DiscordInterface.SendNotification(&send)
	time.Sleep(time.Minute * 5)
	*aliveLoopChannel <- true

}
