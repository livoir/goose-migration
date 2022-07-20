-- +goose Up
-- +goose StatementBegin
ALTER TABLE users_roles
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE users_roles
    ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users_roles
    DROP FOREIGN KEY fk_role_id;

ALTER TABLE users_roles
    DROP FOREIGN KEY fk_user_id;
-- +goose StatementEnd
