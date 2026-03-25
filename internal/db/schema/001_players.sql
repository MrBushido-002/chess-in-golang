-- +goose Up

CREATE TABLE players(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(20) NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL
);

-- +goose Down

DROP TABLE players;