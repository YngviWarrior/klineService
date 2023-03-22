package mysql

import (
	"database/sql"
	"klineService/database/repositories"
	entity "klineService/entities/asset"
	"log"
)

type AssetRepository struct{}

type AssetRepositoryInterface interface {
	List(tx *sql.Tx, conn *sql.Conn) (list []*entity.Asset)
}

func (*AssetRepository) List(tx *sql.Tx, conn *sql.Conn) (list []*entity.Asset) {
	query := `SELECT asset, symbol, active FROM asset`
	stmt, err := repositories.Prepare(tx, conn, query)

	if err != nil {
		log.Panicln("CRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("CRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var c entity.Asset

		err := res.Scan(&c.Asset, &c.Symbol, &c.Active)

		if err != nil {
			log.Panicln("CRL 03: ", err)
			return
		}

		list = append(list, &c)
	}

	return
}
