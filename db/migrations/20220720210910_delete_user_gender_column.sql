-- +goose Up
-- +goose StatementBegin
# ALTER TABLE users DROP column gender;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
