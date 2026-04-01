package api

import (
	"net/http"
	"github.com/MrBushido-002/chess-in-golang/internal/auth"
	"github.com/MrBushido-002/chess-in-golang/internal/db"
	"github.com/MrBushido-002/chess-in-golang/internal/game"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"strings"
	"strconv"
	"fmt"

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

func(cfg *APIConfig) HandleGetReplay(w http.ResponseWriter, r *http.Request) {
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

	moves, err := queries.GetMoves(context.Background(), game_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not retrieve game data"))
		return
	}
	FENstrings := []string{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"}
	board := game.FENParser(FENstrings[0])

	for _, moveRecord := range moves {
		var startRank, startFile, endRank, endFile int
		_, err := fmt.Sscanf(moveRecord.Move, "{%d %d} {%d %d}", &startRank, &startFile, &endRank, &endFile)
		if err != nil {
			// try new format
			parts := strings.Split(moveRecord.Move, ",")
			if len(parts) == 4 {
				startRank, _ = strconv.Atoi(parts[0])
				startFile, _ = strconv.Atoi(parts[1])
				endRank, _ = strconv.Atoi(parts[2])
				endFile, _ = strconv.Atoi(parts[3])
			}
		}
	    m := game.Move{
			Start: game.Square{Rank: startRank, File: startFile},
			End:   game.Square{Rank: endRank, File: endFile},
		}
		board = game.HypotheticalMove(board, m)
		FENstrings = append(FENstrings, game.BoardToFEN(board))
	}
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(FENstrings)
}