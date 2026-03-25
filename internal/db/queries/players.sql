-- name: CreatePlayer :one
INSERT INTO players(id, username, hashed_password)
    VALUES(
        gen_random_uuid(),
        $1,
        $2
    )
    RETURNING id, username;

-- name: GetUserInfo :one
SELECT id, hashed_password FROM players WHERE username = $1;