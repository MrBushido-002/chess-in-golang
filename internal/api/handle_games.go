package api

import (
	"net/http"
	"github.com/MrBushido-002/chess-in-golang/internal/auth"
	"github.com/MrBushido-002/chess-in-golang/internal/db"
	"context"
	"encoding/json"
	"github.com/google/uuid"

)


func (cfg *APIConfig) HandleCreateGame(w http.ResponseWriter, r *http.Request) {
	queries := db.New(cfg.DB)

	

	player_id, err := auth.AuthenticateRequest(r, cfg.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	game, err := queries.CreateGame(context.Background(), uuid.NullUUID{UUID: player_id, Valid: true})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not create game"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func (cfg *APIConfig) HandleJoinGame(w http.ResponseWriter, r *http.Request) {
	queries := db.New(cfg.DB)
	game_id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game ID"))
		return
	}

	player_id, err := auth.AuthenticateRequest(r, cfg.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	game, err := queries.JoinGame(context.Background(), db.JoinGameParams{
		GameID: game_id,
		BlackPlayerID: uuid.NullUUID{UUID: player_id, Valid: true},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not join game"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

func (cfg *APIConfig) HandleGetGames(w http.ResponseWriter, r *http.Request) {
	queries := db.New(cfg.DB)
	
	_, err := auth.AuthenticateRequest(r, cfg.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	
	game_id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game ID"))
		return
	}
	
	gameData, err := queries.GetGame(context.Background(), game_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not retrieve game data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gameData)
}