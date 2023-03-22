package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Pool *sql.DB
}

type DatabaseInterface interface {
	CreatePool(size int64)
	CreateConnection() (conn *sql.Conn)
	CreateTransaction(ctx *context.Context, conn *sql.Conn) (tx *sql.Tx)
}

func (d *Database) CreatePool(size int64) {
	d.Pool = CreateMysqlPool(size)
}

func (d *Database) CreateConnection() (conn *sql.Conn) {
	ctx := context.TODO()

	conn, err := d.Pool.Conn(ctx)
	if err != nil {
		log.Panicln("Conn Create: ", err)
	}

	return
}

func (d *Database) CreateTransaction(ctx *context.Context, conn *sql.Conn) (tx *sql.Tx) {
	tx, err := conn.BeginTx(*ctx, &sql.TxOptions{})
	if err != nil {
		log.Panicln("TX Create: ", err)
	}

	return
}

func CreateMysqlPool(size int64) *sql.DB {
	var db *sql.DB
	var err error

	switch os.Getenv("ENVIROMENT") {
	case "dev":
		db, err = sql.Open("mysql", os.Getenv("DB_DEV"))
	default:
		db, err = sql.Open("mysql", os.Getenv("DB_PROD"))
	}

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxOpenConns(int(size))

	return db
}

func CreatePostgresPool() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_DEV"))

	if err != nil {
		log.Fatal("DC 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
