package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnectionDB(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
}
