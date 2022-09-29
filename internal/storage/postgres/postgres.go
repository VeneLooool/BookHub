package postgres

import (
	"github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}
