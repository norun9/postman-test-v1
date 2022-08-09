package db

import (
	"database/sql"
	"fmt"
	"log"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/norun9/postmantest/pkg/config"
)

const timezone string = "Asia%2FTokyo"

var (
	fixtures *testfixtures.Loader
)

// NewMySQL connect db init
func NewMySQL(c config.MySQL) *sql.DB {
	if c.Pseudo {
		return nil
	}
	// "test:test@tcp(127.0.0.1:5306)/testdb?charset=utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s", c.DBUserName, c.DBPassword, c.DBHost, c.DBPort, c.DBName, timezone)
	openedDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	openedDB.SetMaxIdleConns(c.MaxIdleConns)
	openedDB.SetMaxOpenConns(c.MaxOpenConns)
	openedDB.SetConnMaxLifetime(c.ConnMaxLifetime)

	if err = openedDB.Ping(); err != nil {
		panic(err)
	}

	if fixtures, err = testfixtures.New(
		testfixtures.Database(openedDB),
		testfixtures.Dialect("mysql"),
		testfixtures.Paths("../../testdata/fixture/common"),
	); err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	if err = fixtures.Load(); err != nil {
		log.Println(err)
	}

	return openedDB
}
