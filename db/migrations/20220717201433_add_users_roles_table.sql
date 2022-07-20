-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_roles(
    id integer primary key auto_increment,
    user_id integer not null,
    role_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_roles;
-- +goose StatementEnd
