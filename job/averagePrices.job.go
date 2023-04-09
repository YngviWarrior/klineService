package job

import (
	"klineService/database"
	bybitstructs "klineService/services/byBitStructs"
	rabbitmqstructs "klineService/services/rabbitMQStructs"
	"klineService/utils"
	"time"
)

func (j *Job) AveragePrices(db *database.Database, loopChannel *chan bool) {
	conn := db.CreateConnection()

	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from := time.Date(to.Year(), to.Month(), to.Day()-1, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)

	dayList := j.KlineRepo.FindAvg(nil, conn, from.UnixMicro(), to.UnixMicro(), j.Test)

	for _, v := range dayList {
		var a rabbitmqstructs.InputAvgMessageDto

		a.Asset = v.Asset
		a.AssetSymbol = v.AssetSymbol
		a.AssetQuote = v.AssetQuote
		a.AssetQuoteSymbol = v.AssetQuoteSymbol
		a.Roc = v.Roc
		a.Avg = v.Close
		a.Period = "Day"

		j.RabbitMQ.SendAveragePrice(&a)
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-7, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)

	weekList := j.KlineRepo.FindAvg(nil, conn, from.UnixMicro(), to.UnixMicro(), j.Test)

	for _, v := range weekList {
		var a rabbitmqstructs.InputAvgMessageDto

		a.Asset = v.Asset
		a.AssetSymbol = v.AssetSymbol
		a.AssetQuote = v.AssetQuote
		a.AssetQuoteSymbol = v.AssetQuoteSymbol
		a.Roc = v.Roc
		a.Avg = v.Close
		a.Period = "Week"

		j.RabbitMQ.SendAveragePrice(&a)
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-30, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)

	monthList := j.KlineRepo.FindAvg(nil, conn, from.UnixMicro(), to.UnixMicro(), j.Test)

	for _, v := range monthList {
		var a rabbitmqstructs.InputAvgMessageDto

		a.Asset = v.Asset
		a.AssetSymbol = v.AssetSymbol
		a.AssetQuote = v.AssetQuote
		a.AssetQuoteSymbol = v.AssetQuoteSymbol
		a.Roc = v.Roc
		a.Avg = v.Close
		a.Period = "Month"

		j.RabbitMQ.SendAveragePrice(&a)
	}

	var p bybitstructs.GetKlinesParams

	for _, assetInfo := range dayList {
		var sum float64 = 0
		var first float64 = 0
		var last float64 = 0

		switch assetInfo.Asset {
		case 3:
			p.Symbol = "BTCUSDT"
		case 4:
			p.Symbol = "ETHUSDT"
		}

		p.Interval = "1"
		p.Limit = 200

		switch assetInfo.Exchange {
		case 1:
		case 2:
			resp, _ := j.ByBitInterface.GetKlines("spot", p.Symbol, p.Interval, 0, 0, p.Limit)

			if len(resp) == 0 {
				continue
			}

			for i, v := range resp {
				if i == 0 {
					first = utils.ParseFloat(v.Close)
				}

				if i == len(resp)-1 {
					last = utils.ParseFloat(v.Close)
				}

				sum += utils.ParseFloat(v.Close)
			}
		}

		smaAvg := sum / float64(p.Limit)

		if first == 0 {
			continue
		}

		var a rabbitmqstructs.InputAvgMessageDto

		a.Asset = assetInfo.Asset
		a.AssetSymbol = assetInfo.AssetSymbol
		a.AssetQuote = assetInfo.AssetQuote
		a.AssetQuoteSymbol = assetInfo.AssetQuoteSymbol
		a.Roc = ((last / first) - 1) * 100
		a.Avg = smaAvg
		a.Period = "Sma200"

		j.RabbitMQ.SendAveragePrice(&a)
	}

	conn.Close()

	time.Sleep(time.Second * 5)
	*loopChannel <- true
}
