package api

import(
	"net/http"
	"encoding/json"
	"github.com/MrBushido-002/chess-in-golang/internal/db"
)

type registerPlayerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func(cfg *APIConfig) HandleRegisterPlayer(w http.ResponseWriter, r *http.Request) {
	var req registerPlayerRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	player, err := db.RegisterPlayer(cfg.DB, req.Username, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not register user"))
		return
	}
	json.NewEncoder(w).Encode(player)
}

type loginPlayerRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func (cfg *APIConfig) HandlePlayerLogin(w http.ResponseWriter, r *http.Request) {
	var req loginPlayerRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	JWT, err := db.PlayerLogin(cfg.DB, req.Username, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid credentials"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JWT)
}