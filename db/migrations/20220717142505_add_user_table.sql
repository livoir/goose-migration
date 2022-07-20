-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id integer primary key auto_increment,
    name varchar(255),
    username varchar(255),
    password varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
