package api

import "database/sql"

type APIConfig struct {
	DB *sql.DB
	JWTSecret string
}