-- name: AddLocation :one
INSERT INTO location (
                      country,
                      region,
                      zone,
                      city) values ($1,$2,$3,$4) returning *;

-- name: GetLocation :one
SELECT * FROM location
WHERE id = $1 LIMIT 1;

-- name: GetLocations :many
SELECT * FROM location
ORDER BY created_at;

-- name: DeleteLocation :exec
DELETE FROM location
WHERE id = $1;