package job

import (
	"klineService/database"
	"klineService/database/repositories/mysql"
	"klineService/services"
)

func LiveStream(db *database.Database) {
	var BinanceInterface services.BinanceInterface = &services.Binance{}
	var ByBitInterface services.ByBitInterface = &services.ByBit{}
	var assetRepo mysql.AssetRepositoryInterface = &mysql.AssetRepository{}

	conn := db.CreateConnection()
	assetList := assetRepo.List(nil, conn)
	conn.Close()

	go ByBitInterface.LiveKlines(db, assetList)
	go BinanceInterface.LiveKlines(db, assetList)
}
