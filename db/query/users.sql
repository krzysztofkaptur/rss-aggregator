-- name: FetchUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (email) VALUES ($1) RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key=$1;