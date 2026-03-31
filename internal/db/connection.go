package db

import (
    "fmt"
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
)

func NewConnection(connString string) (*sql.DB, error) {
    conn, err := sql.Open("pgx", connString)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %w", err)
    }
    return conn, nil
}