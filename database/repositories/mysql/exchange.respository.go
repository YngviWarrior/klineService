package mysql

import (
	"database/sql"
	entity "klineService/entities/exchange"
	"log"
)

type ExchangeRepository struct{}

type ExchangeRepositoryInterface interface {
	Create(tx *sql.Tx, name string) bool
	List(tx *sql.Tx) (list []*entity.Exchange)
}

func (*ExchangeRepository) List(tx *sql.Tx) (list []*entity.Exchange) {
	stmt, err := tx.Prepare(`
		SELECT exchange, name, active
		FROM exchange
	`)

	if err != nil {
		log.Panicln("UEKR 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("UEKR 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u entity.Exchange

		err := res.Scan(&u.Exchange, &u.Name, &u.Active)

		if err != nil {
			log.Panicln("UEKR 03: ", err)
			return
		}

		list = append(list, &u)
	}

	return
}

func (*ExchangeRepository) Create(tx *sql.Tx, name string) bool {
	stmt, err := tx.Prepare(`INSERT INTO exchange (name) VALUES (?)`)

	if err != nil {
		log.Panicln("ERC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(name)

	if err != nil {
		log.Panicln("ERC 02: ", err)
		return false
	}

	return true
}
