-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name, email, password)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserFromUserName :one
SELECT * FROM users
WHERE user_name = $1;