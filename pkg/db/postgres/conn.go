package postgres

import (
	"fmt"
	"github.com/VeneLooool/BookHub/internal/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPsqlDB(c *config.Config) (db *sqlx.DB, err error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)
	db, err = sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
