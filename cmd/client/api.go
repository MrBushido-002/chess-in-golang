package main

import(
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
)
type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(username string, password string) (string, error) {
	creds := LoginCredentials{Username: username, Password: password}
	data, err := json.Marshal(creds)
	if err != nil {
		return "", err
	}

	res, err := http.Post("http://localhost:2209/players/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	var token string
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func RegisterUser(username string, password string) error {
	creds := LoginCredentials{Username: username, Password: password}
	data, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	res, err := http.Post("http://localhost:2209/players/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("registration failed: %s", res.Status)
	}
	return nil
}

func CreateGame(token string) (string, error) {
	var game struct {
		GameID string `json:"GameID"`
	}
	
	req, err := http.NewRequest("POST", "http://localhost:2209/games/", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("game creation failed: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&game)
	if err != nil {
		return "", err
	}
	return game.GameID, nil
}

func JoinGame(token string, gameID string) (string, error) {
	var game struct {
		GameID string `json:"GameID"`
	}
	
	url := fmt.Sprintf("http://localhost:2209/games/%s/join", gameID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to join game: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&game)
	if err != nil {
		return "", err
	}
	return game.GameID, nil
}


type GameState struct {
		GameID      string `json:"GameID"`
		BoardState  string `json:"BoardState"`
		Turn        string `json:"Turn"`
		Status      string `json:"Status"`
		WhitePlayerID string `json:"WhitePlayerID"`
		BlackPlayerID string `json:"BlackPlayerID"`
	}

func GetGame(token string, gameID string) (GameState, error) {
	var gamestate GameState
	url := fmt.Sprintf("http://localhost:2209/games/%s", gameID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GameState{}, err
	}

	req.Header.Set("authorization", "Bearer " + token)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return GameState{}, err
	}

	if res.StatusCode != http.StatusOK {
		return GameState{}, fmt.Errorf("failed to fetch game data: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&gamestate)
	if err != nil {
		return GameState{}, err
	}
	return gamestate, nil

}


type moveRequest struct {
    Start struct {
        Rank int `json:"Rank"`
        File int `json:"File"`
    } `json:"start"`
    End struct {
        Rank int `json:"Rank"`
        File int `json:"File"`
    } `json:"end"`
}

func MakeMove(token string, gameID, moveStr string) error {
	startFile := int(moveStr[0] - 'a')
	startRank := 8 - int(moveStr[1]-'0') 
	endFile := int(moveStr[2] - 'a')
	endRank := 8 - int(moveStr[3]-'0')
	
	move := moveRequest{}
	move.Start.Rank = startRank
	move.Start.File = startFile
	move.End.Rank = endRank
	move.End.File = endFile

	data, err := json.Marshal(move)
	if err != nil {
		return err
	}
	
	
	url := fmt.Sprintf("http://localhost:2209/games/%s/moves", gameID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not make move: %s", res.Status)
	}

	return nil

}

type MoveRecord struct {
	MoveID int `json:"move_id"`
	GameID string `json:"game_id"`
	Move string `json:"move"`
	Color string `json:"color"`
}

func GetMoves(token string, gameID string) ([]MoveRecord, error) {
	var moves []MoveRecord
	
	url := fmt.Sprintf("http://localhost:2209/games/%s/moves", gameID)



	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} 

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not fetch game: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&moves)
	if err != nil {
		return nil, err
	}

	return moves, nil
}

func GetReplay(token string, gameID string) ([]string, error) {
    url := fmt.Sprintf("http://localhost:2209/games/%s/replay", gameID)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + token)
    
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    
    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("could not fetch replay: %s", res.Status)
    }
    
    var fenStrings []string
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&fenStrings)
    if err != nil {
        return nil, err
    }
    return fenStrings, nil
}