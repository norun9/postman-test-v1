package util

import (
	"database/sql"
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	fixtures *testfixtures.Loader
)

func Prepare() {
	dbconf := "test:test@tcp(127.0.0.1:5306)/testdb?charset=utf8mb4"
	var err error
	dbConn, err := sql.Open("mysql", dbconf)
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
