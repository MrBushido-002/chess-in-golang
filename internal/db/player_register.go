package db

import(
	
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5"
	"context"
)

func RegisterPlayer(conn *pgx.Conn, username string, password string) (CreatePlayerRow, error) {
	queries := New(conn)
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return CreatePlayerRow{}, err
	}
	Player, err := queries.CreatePlayer(context.Background(), CreatePlayerParams{
		Username: username,
		HashedPassword: string(hashedPassword),
	})
	if err != nil {
		return CreatePlayerRow{}, err
	}

	return Player, nil

}