-- name: CreateUser :one
INSERT INTO users(id, name, created_at, updated_at) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUsers :exec
DELETE FROM users;