package database

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbFinal *sqlx.DB
	once    sync.Once
	errDb   error
)

func ConnectDB() (*sqlx.DB, error) {

	once.Do(func() {
		dsn := "root:root@tcp(localhost:14045)/readarticle"

		db, err := sqlx.Connect("mysql", dsn)
		if err != nil {
			errDb = err
			return
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		err = db.Ping()

		if err != nil {
			errDb = err
			return
		}
		dbFinal = db
		defer db.Close()
	})

	return dbFinal, errDb

}
