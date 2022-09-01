package repository

import "database/sql"

type MarketDb struct {
	db *sql.DB
}

func NewMarketDb(db *sql.DB) *MarketDb {
	return &MarketDb{db: db}
}
