-- +goose Up
-- +goose StatementBegin
INSERT INTO roles VALUES (1, 'admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM roles WHERE id = 1;
-- +goose StatementEnd
