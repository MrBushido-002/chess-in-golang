package db

import(
	
	"golang.org/x/crypto/bcrypt"
	"context"
	"database/sql"
)

func RegisterPlayer(conn *sql.DB, username string, password string) (CreatePlayerRow, error) {
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