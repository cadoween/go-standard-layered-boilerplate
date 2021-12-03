package config

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreSQL() (*pgxpool.Pool, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT")),
		Path:   os.Getenv("DATABASE_NAME"),
	}

	q := dsn.Query()
	q.Add("sslmode", os.Getenv("DATABASE_SSLMODE"))

	dsn.RawQuery = q.Encode()

	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
