package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitConnection() {
	dbURL := "postgres://postgres:Awaisniaz@localhost:5432/user-management-system?sslmode=disable"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Could not ping DB: %v", err)
	}
	log.Println("✅ PostgreSQL connected")
	DB = pool

}
