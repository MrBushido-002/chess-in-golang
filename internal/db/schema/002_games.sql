-- +goose Up

CREATE TYPE game_status AS ENUM ('waiting', 'active', 'complete');
CREATE TYPE player_color AS ENUM ('white', 'black');

CREATE TABLE games (
    game_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    white_player_id UUID REFERENCES players(id),
    black_player_id UUID REFERENCES players(id),
    status game_status NOT NULL DEFAULT 'waiting',
    turn player_color NOT NULL DEFAULT 'white',
    board_state TEXT NOT NULL DEFAULT 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE games;
DROP TYPE game_status;
DROP TYPE player_color;