-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE teams;
-- +goose StatementEnd

-- goose -dir migrations postgres "postgres://postgres:admin@localhost:5432/mpl?sslmode=disable" up