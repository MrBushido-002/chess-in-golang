package api

import (
	"net/http"
	"github.com/MrBushido-002/chess-in-golang/internal/auth"
	"github.com/MrBushido-002/chess-in-golang/internal/db"
	"github.com/MrBushido-002/chess-in-golang/internal/game"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"fmt"

)

type makeMoveRequest struct {
    Start game.Square `json:"start"`
    End   game.Square `json:"end"`
}

func (cfg *APIConfig) HandelMakeMove(w http.ResponseWriter, r *http.Request) {
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

	gameData, err := queries.GetGame(context.Background(), game_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not create gameData"))
		return
	}

	var player_color game.Color
	var opponent_color game.Color
	if player_id == gameData.WhitePlayerID.UUID {
		player_color = game.White
		opponent_color = game.Black
	} else if player_id == gameData.BlackPlayerID.UUID {
		player_color = game.Black
		opponent_color = game.White
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not in this gameData"))
		return
	}

	if string(player_color) != string(gameData.Turn) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not your turn"))
		return
	}

	var req makeMoveRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get move"))
		return
	}

	move := game.Move{Start: req.Start, End: req.End}
	move_string := fmt.Sprintf("%v %v", move.Start, move.End)


	if game.IsValidMove(game.FENParser(gameData.BoardState), move, player_color) != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid move"))
		return
	}

	new_board := game.HypotheticalMove(game.FENParser(gameData.BoardState), move)

	new_game_data := game.BoardToFEN(new_board)

	var nextTurn db.PlayerColor
	if player_color == game.White {
		nextTurn = db.PlayerColorBlack
	} else {
		nextTurn = db.PlayerColorWhite
	}


	_, err = queries.UpdateBoardState(context.Background(), db.UpdateBoardStateParams{
		GameID: game_id,
		BoardState: new_game_data,
		Turn: nextTurn,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid move"))
		return
	}

	_, err = queries.CreateMove(context.Background(), db.CreateMoveParams{
		GameID: game_id,
		Move: move_string,
		Color: db.PlayerColor(player_color),
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid move"))
		return
	}
	if game.IsCheckMate(new_board, opponent_color) == true {
		queries.UpdateGameStatus(context.Background(), db.UpdateGameStatusParams{
			Status: db.GameStatusComplete,
			GameID:     game_id,
		})
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Checkmate!"))
    return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(new_game_data)
}

func (cfg *APIConfig) HandelGetMoves(w http.ResponseWriter, r *http.Request) {
	queries := db.New(cfg.DB)

	game_id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game ID"))
		return
	}

	moves, err := queries.GetMoves(context.Background(), game_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not fetch moves"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(moves)
}