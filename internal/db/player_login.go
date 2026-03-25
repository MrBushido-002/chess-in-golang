package db

import(
	
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5"
	"context"
	"os"
	"github.com/MrBushido-002/chess-in-golang/internal/auth"
)

func PlayerLogin(conn *pgx.Conn, username string, password string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	queries := New(conn)
	expiresIn := 12 * time.Hour
	
	playerInfo, err := queries.GetUserInfo(context.Background(), username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(playerInfo.HashedPassword), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	tokenString, err := auth.MakeJWT(playerInfo.ID, secretKey, expiresIn)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}