package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/simplebank/util"
	"log"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDb)
	m.Run()
}
