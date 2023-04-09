package job

import (
	"context"
	"klineService/database"
)

func (j *Job) LiveStream(ctx context.Context, db *database.Database, quitChannel *chan bool) {
	go j.ByBitInterface.LiveKlines(ctx, db, quitChannel)
	go j.BinanceInterface.LiveKlines(ctx, db, quitChannel)
}
