package database

import (
	"context"
	"database/sql"
	"time"
)

func Open(ctx context.Context, databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)

	if err != nil {
		return nil, err
	}

	// configuring connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	// verifying whether postgres is reachable
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
