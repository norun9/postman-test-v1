package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/norun9/postmantest/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	// DefaultBlackList :
	DefaultBlackList = boil.Blacklist("created_at")
)

// Client :
type Client struct {
	dbClient *sql.DB
}

// NewClient :
func NewClient(db *sql.DB) Client {
	return Client{db}
}

func (c Client) Get(ctx context.Context) SQLHandler {
	if tx := util.GetDBTx(ctx); tx != nil {
		return tx
	}
	return c.dbClient
}

// GetGlobalDB :
func (c Client) GetGlobalDB() SQLHandler {
	return c.dbClient
}

// SQLHandler :
type SQLHandler interface {
	boil.ContextExecutor
}

// NewDB initialize databases
func NewDB(sqlDB *sql.DB) Client {
	util.InitLocal()
	boil.SetLocation(time.Local)
	// デバックを有効化
	boil.DebugMode = true
	return NewClient(sqlDB)
}
