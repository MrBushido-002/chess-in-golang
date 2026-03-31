-- name: CreateGame :one
INSERT INTO games(white_player_id)
    VALUES(
        $1
    )
    RETURNING *;

-- name: GetGame :one 
SELECT * FROM games WHERE game_id = $1;

-- name: UpdateBoardState :one 
UPDATE games
SET board_state = $2, turn = $3 WHERE game_id = $1
RETURNING *;

-- name: UpdateGameStatus :one
UPDATE games
set status = $2 WHERE game_id = $1
RETURNING *;

-- name: JoinGame :one
UPDATE games 
SET black_player_id = $2, status = 'active' WHERE game_id = $1
RETURNING *;

