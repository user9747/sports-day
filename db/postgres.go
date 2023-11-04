package db

import (
	"log"
	"sports-day/conf"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB = nil

// var readOnlyDatabase *sqlx.DB = nil
var dbOnce sync.Once

// var readOnlyDBOnce sync.Once
// TODO: execute this once and store db detail in global var
func GetDB() *sqlx.DB {
	var err error
	dbOnce.Do(func() {
		database, err = sqlx.Open(conf.PostgresConf.Driver, conf.PostgresConf.ConnectionString)
		if err != nil {
			log.Fatal("get DB connection error", err)
		}

	})
	return database
}
