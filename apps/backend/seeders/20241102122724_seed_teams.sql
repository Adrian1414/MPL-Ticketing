-- +goose Up
-- +goose StatementBegin
INSERT INTO teams (name) VALUES ('EVOS'), ('FNOC'), ('RRQ'), ('TLID'), ('GEEK'), ('RBL'), ('DEWA'), ('AE'), ('BTR');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM teams WHERE name IN ('EVOS'), ('FNOC'), ('RRQ'), ('TLID'), ('GEEK'), ('RBL'), ('DEWA'), ('AE'), ('BTR');
-- +goose StatementEnd

-- goose -dir seeders postgres "postgres://postgres:admin@localhost:5432/mpl?sslmode=disable" up
