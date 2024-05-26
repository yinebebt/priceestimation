-- name: AddUser :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    password) values ($1,$2,$3,$4) returning *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;