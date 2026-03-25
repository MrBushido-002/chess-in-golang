package db

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5"
)

func NewConnection(connString string) (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), connString)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %w", err)
    }
    return conn, nil
}