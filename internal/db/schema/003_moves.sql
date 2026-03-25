-- +goose Up

CREATE TABLE moves(
    move_id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL REFERENCES games(game_id),
    move VARCHAR(20) NOT NULL,
    color player_color NOT NULL
);

-- +goose Down

DROP TABLE moves;