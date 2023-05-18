package job

import (
	"context"
	"encoding/json"
	"klineService/database"
	"klineService/entities/asset"
	"log"
	"time"
)

func (j *Job) AssetManager(cancel context.CancelFunc, db *database.Database, assetManagerLoopChannel *chan bool, quitLoopChannel *chan bool) {
	conn := db.CreateConnection()
	assets := j.AssetRepo.List(nil, conn)
	conn.Close()

	if len(assets) == 0 {
		log.Println("JOB AssetManager: No assets found")
		return
	}

	cachedAssets := j.Redis.GetCache("Assets", "string")

	var cachedSlice []*asset.Asset
	err := json.Unmarshal([]byte(cachedAssets.(string)), &cachedSlice)

	if err != nil {
		log.Panic("JOB AssetManager 01: ", err)
	}

	if len(cachedSlice) != 0 && len(cachedSlice) != len(assets) {
		b, err := json.Marshal(assets)

		if err != nil {
			log.Panic("JOB AssetManager 02: ", err)
		}

		j.Redis.GetInstance().Set(context.TODO(), "Assets", b, 0)
		//reboot liveklines
		cancel()
		*quitLoopChannel <- true
	}

	time.Sleep(time.Second * 10)
	*assetManagerLoopChannel <- true
}
