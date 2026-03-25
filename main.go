package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"context"

	"github.com/MrBushido-002/chess-in-golang/internal/api"
	"github.com/MrBushido-002/chess-in-golang/internal/db"
	"github.com/joho/godotenv"

)

func main() {
	godotenv.Load()
	
	dbURL := os.Getenv("DB_URL")
    conn, err := db.NewConnection(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	cfg := api.APIConfig{
		DB: conn,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	fmt.Println("Database connected successfully!")
	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /games", api.HandleGetGames)
	mux.HandleFunc("POST /players/register", cfg.HandleRegisterPlayer)
	mux.HandleFunc("POST /players/login", cfg.HandlePlayerLogin)

	server := &http.Server{
		Addr: ":2209",
		Handler: mux,
	}

	fmt.Println("Server starting on port 2209")
	log.Fatal(server.ListenAndServe())
}
