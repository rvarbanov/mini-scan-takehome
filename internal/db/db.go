package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/rvarbanov/mini-scan-takehome/internal/model"

	_ "github.com/lib/pq"
)

type DBInterface interface {
	StoreScan(ctx context.Context, scan model.Scan) error
}

type DB struct {
	db *sql.DB
}

func NewDB(host, port, user, pass, name string) DBInterface {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		name,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db: db}
}

func (d *DB) StoreScan(ctx context.Context, scan model.Scan) error {
	query := `
		INSERT INTO scan (
			ip,
			port,
			service,
			data,
			timestamp,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		ON CONFLICT (ip, port, service) 
		DO UPDATE SET 
			data = EXCLUDED.data,
			timestamp = EXCLUDED.timestamp,
			updated_at = EXCLUDED.updated_at`

	result, err := d.db.ExecContext(ctx, query,
		scan.IP,
		scan.Port,
		scan.Service,
		scan.Data,
		scan.Timestamp,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to store scan: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return err
}
