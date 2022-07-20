-- +goose Up
-- +goose StatementBegin
# ALTER TABLE users
#     ADD gender varchar(10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
