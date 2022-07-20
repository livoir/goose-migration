-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users(id, name, username, password) VALUES (?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET name = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE from users WHERE id = ?;

# -- name: TestSelectGender :many
# SELECT gender FROM users;