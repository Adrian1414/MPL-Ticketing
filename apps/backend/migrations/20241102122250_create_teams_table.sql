-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS teams(
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teams;
-- +goose StatementEnd

-- goose -dir migrations postgres "postgres://postgres:admin@localhost:5432/mpl?sslmode=disable" up