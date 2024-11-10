-- +goose Up
-- +goose StatementBegin
INSERT INTO teams (id, name) VALUES 
(1, 'AE'), 
(2, 'BTR'), 
(3, 'DEWA'), 
(4, 'EVOS'),
(5, 'FNOC'), 
(6, 'GEEK'), 
(7, 'RBL'), 
(8, 'RRQ'), 
(9, 'TLID');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM teams;
-- +goose StatementEnd
