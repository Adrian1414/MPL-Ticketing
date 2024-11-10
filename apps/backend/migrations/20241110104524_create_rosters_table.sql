-- +goose Up
-- +goose StatementBegin
-- migrations/20231103_create_rosters_table.sql
CREATE TABLE IF NOT EXISTS rosters (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    team_id INT REFERENCES teams(id) ON DELETE CASCADE,
    photo TEXT
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rosters;
-- +goose StatementEnd
