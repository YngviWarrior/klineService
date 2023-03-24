package mysql

import (
	"database/sql"
	"fmt"
	"klineService/database/repositories"
	entity "klineService/entities/kline"
	"log"
)

type KlineRepository struct{}

type KlineRepositoryInterface interface {
	FindLastPrice(tx *sql.Tx, conn *sql.Conn, asset, asset_quote, exchange uint64) (close float64)
	FindLimit(tx *sql.Tx, conn *sql.Conn, asset, asset_quote, exchange, limit int64) (list []*entity.Kline)
	FindFirstMts(tx *sql.Tx, conn *sql.Conn, asset, asset_quote, exchange int64) (c entity.Kline)
	FindFirst(tx *sql.Tx, conn *sql.Conn, asset, asset_quote, exchange, from int64) (c entity.Kline)
	FindAvg(tx *sql.Tx, conn *sql.Conn, from, to int64) (list []*entity.Kline)
	Create(tx *sql.Tx, conn *sql.Conn, kline *entity.Kline) bool
	CreateDirect(tx *sql.DB, kline *entity.Kline) bool
}

func (*KlineRepository) FindLastPrice(tx *sql.Tx, conn *sql.Conn, asset, asset_quote, exchange uint64) (close float64) {
	query := `
		SELECT close
		FROM kline 
		WHERE asset = ? AND asset_quote = ? AND exchange = ?
		ORDER BY mts DESC 
		LIMIT 1`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRFLP 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(asset, asset_quote, exchange).Scan(&close)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFLP 02: ", err)
		return
	}

	return
}

func (*KlineRepository) Create(tx *sql.Tx, conn *sql.Conn, kline *entity.Kline) bool {
	constraint := `
		ON DUPLICATE KEY UPDATE 
		asset = ` + fmt.Sprintf("%v", kline.Asset) + `,
		asset_quote = ` + fmt.Sprintf("%v", kline.AssetQuote) + `,
		exchange = ` + fmt.Sprintf("%v", kline.Exchange) + `,
		open = ` + fmt.Sprintf("%v", kline.Open) + `,
		close = ` + fmt.Sprintf("%v", kline.Close) + `,
		high = ` + fmt.Sprintf("%v", kline.High) + `,
		low = ` + fmt.Sprintf("%v", kline.Low) + `,
		volume = ` + fmt.Sprintf("%v", kline.Volume) + `
	`
	query := `INSERT INTO kline(asset, asset_quote, exchange, mts, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)` + constraint

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRC 01: ", err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(kline.Asset, kline.AssetQuote, kline.Exchange, (kline.Mts * 1000), kline.Open, kline.Close, kline.High, kline.Low, kline.Volume)

	if err != nil {
		log.Panicln("CRC 02: ", err)
		return false
	}

	return true
}

func (*KlineRepository) CreateDirect(db *sql.DB, kline *entity.Kline) bool {
	constraint := `
		ON DUPLICATE KEY UPDATE 
		asset = ` + fmt.Sprintf("%v", kline.Asset) + `,
		asset_quote = ` + fmt.Sprintf("%v", kline.AssetQuote) + `,
		exchange = ` + fmt.Sprintf("%v", kline.Exchange) + `,
		open = ` + fmt.Sprintf("%v", kline.Open) + `,
		close = ` + fmt.Sprintf("%v", kline.Close) + `,
		high = ` + fmt.Sprintf("%v", kline.High) + `,
		low = ` + fmt.Sprintf("%v", kline.Low) + `,
		volume = ` + fmt.Sprintf("%v", kline.Volume) + `
	`
	query := `INSERT INTO kline(asset, asset_quote, exchange, mts, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)` + constraint

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Panicln("CRCD 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(kline.Asset, kline.AssetQuote, kline.Exchange, (kline.Mts * 1000), kline.Open, kline.Close, kline.High, kline.Low, kline.Volume)

	if err != nil {
		log.Panicln("CRCD 02: ", err)
		return false
	}

	return true
}

func (*KlineRepository) FindLimit(tx *sql.Tx, conn *sql.Conn, asset, assetQuote, exchange, limit int64) (list []*entity.Kline) {
	query := `
		SELECT asset, asset_quote, exchange, close, mts
		FROM kline 
		WHERE asset = ? AND asset_quote = ? AND exchange = ?
		ORDER BY mts DESC
		LIMIT ?`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRFL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(asset, assetQuote, exchange, limit)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFL 02: ", err)
		return
	}

	for res.Next() {
		var c entity.Kline

		err := res.Scan(&c.Asset, &c.AssetQuote, &c.Exchange, &c.Close, &c.Mts)

		if err != nil {
			log.Panic("CRFL 03: ", err)
		}

		list = append(list, &c)
	}

	return
}

func (*KlineRepository) FindFirstMts(tx *sql.Tx, conn *sql.Conn, asset, assetQuote, exchange int64) (c entity.Kline) {
	query := `
		SELECT asset, asset_quote, exchange, close, mts
		FROM kline 
		WHERE asset = ? AND asset_quote = ? AND exchange = ?
		ORDER BY mts DESC
		LIMIT 1`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("KRFFM 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(asset, assetQuote, exchange).Scan(&c.Asset, &c.AssetQuote, &c.Exchange, &c.Close, &c.Mts)
	fmt.Println(c)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("KRFFM 02: ", err)
		return
	}

	return
}

func (*KlineRepository) FindFirst(tx *sql.Tx, conn *sql.Conn, asset, assetQuote, exchange, from int64) (c entity.Kline) {
	query := `
		SELECT asset, asset_quote, exchange, close, mts
		FROM kline 
		WHERE mts > ? AND asset = ? AND asset_quote = ? AND exchange = ?
		LIMIT 1`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRF 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(from, asset, assetQuote, exchange).Scan(&c.Asset, &c.AssetQuote, &c.Exchange, &c.Close, &c.Mts)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRF 02: ", err)
		return
	}

	return
}

func (*KlineRepository) FindAvg(tx *sql.Tx, conn *sql.Conn, from, to int64) (list []*entity.Kline) {
	query := `
		SELECT asset, asset_quote, exchange, AVG(close), (((MIN(close) - MAX(close)) / MAX(close)) * 100) as roc
		FROM kline 
		WHERE mts BETWEEN ? AND ?
		GROUP BY asset, asset_quote, exchange`

	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRFA 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(from, to)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFA 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var c entity.Kline

		err := res.Scan(&c.Asset, &c.AssetQuote, &c.Exchange, &c.Close, &c.Roc)

		if err != nil {
			log.Panic("CRFA 03: ", err)
		}

		list = append(list, &c)
	}

	return
}
