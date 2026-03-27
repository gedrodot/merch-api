package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	User   string
	Passwd string
	Addr   string
	Port   string
	DB     string
}

func New(cfg Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Passwd, cfg.Addr, cfg.Port, cfg.DB)

	fmt.Printf("DEBUG DSN: |%s|\n", dataSource)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to connect: %w", err)
	}

	conn.SetMaxOpenConns(50)
	conn.SetMaxIdleConns(50)
	conn.SetConnMaxLifetime(15 * time.Minute)
	conn.SetConnMaxIdleTime(5 * time.Minute)

	return conn, nil
}
