-- name: CreateMove :one
INSERT INTO moves(game_id, move, color)
    VALUES(
        $1,
        $2,
        $3
    )
    RETURNING *;

-- name: GetMoves :many
SELECT * FROM moves WHERE game_id = $1
ORDER BY move_id;