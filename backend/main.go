package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/norun9/postmantest/util"
)

var (
	fixtures *testfixtures.Loader
)

func main() {
	util.Prepare()
}
