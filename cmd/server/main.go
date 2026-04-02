package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
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
    defer conn.Close()

    cfg := api.APIConfig{
        DB:        conn,
        JWTSecret: os.Getenv("SECRET_KEY"),
    }

    fmt.Println("Database connected successfully!")

    mux := http.NewServeMux()
    mux.HandleFunc("POST /players/register", cfg.HandleRegisterPlayer)
    mux.HandleFunc("POST /players/login", cfg.HandlePlayerLogin)
    mux.HandleFunc("POST /games", cfg.HandleCreateGame)
    mux.HandleFunc("POST /games/{id}/join", cfg.HandleJoinGame)
    mux.HandleFunc("POST /games/{id}/moves", cfg.HandelMakeMove)
    mux.HandleFunc("GET /games/{id}", cfg.HandleGetGames)
    mux.HandleFunc("GET /games/{id}/moves", cfg.HandelGetMoves)
    mux.HandleFunc("GET /games/{id}/replay", cfg.HandleGetReplay)

    port := os.Getenv("PORT")
    if port == "" {
        port = "2209"
    }
    fmt.Println("Server starting on port", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}
