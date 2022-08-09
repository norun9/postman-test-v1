package util

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	fixtures *testfixtures.Loader
)

func Prepare() {
	dbconf := "test:test@tcp(127.0.0.1:5306)/testdb?charset=utf8mb4"
	var err error
	var dbConn *sql.DB
	if dbConn, err = sql.Open("mysql", dbconf); err != nil {
		panic(err)
	}
	if err = dbConn.Ping(); err != nil {
		panic(err)
	}
	if fixtures, err = testfixtures.New(
		testfixtures.Database(dbConn),
		testfixtures.Dialect("mysql"),
		testfixtures.Paths("testdata/fixture/common"),
	); err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	if err = fixtures.Load(); err != nil {
		log.Println(err)
	}
}

// InitLocal :
func InitLocal() {
	location := "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}
