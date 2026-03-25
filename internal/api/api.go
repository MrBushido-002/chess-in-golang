package api

import "github.com/jackc/pgx/v5"

type APIConfig struct {
	DB *pgx.Conn
	JWTSecret string
}