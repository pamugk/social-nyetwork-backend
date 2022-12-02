package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

func init() {
	var err error

	dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	log.Println("Initialized DB pool")
}

func CloseDbPool() {
	if dbpool != nil {
		dbpool.Close()
	}
}

func unwrap(text pgtype.Text)*string {
	if (text.Valid) {
		return &text.String
	}
	return nil
}
