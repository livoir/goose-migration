-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD column gender varchar(10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
drop column gender
-- +goose StatementEnd
