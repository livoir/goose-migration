-- +goose Up
-- +goose StatementBegin
ALTER TABLE users drop column gender;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
