-- name: CreateUser :exec
INSERT INTO users (username, password, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1;
