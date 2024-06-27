package storage

import (
	"EMSv0.1/app/internal/config"
	"context"
	"database/sql"
	"fmt"
	"time"
)

// NewStorage init storage
func NewStorage(cfg config.Postgres) (*sql.DB, error) {
	const op = "storage.NewStorage"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// ping storage db
	pingErrChan := make(chan error, 1)
	go func() {
		pingErrChan <- db.Ping()
	}()

	select {
	case <-ctx.Done():
		db.Close()
		return nil, fmt.Errorf("%s: timeout", op)
	case err = <-pingErrChan:
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return db, nil
}
