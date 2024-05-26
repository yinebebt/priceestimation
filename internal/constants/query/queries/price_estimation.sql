-- name: AddPriceEstimation :one
INSERT INTO price_estimation (
                               product_name,
                               price,
                               user_id,
                               location_id
) VALUES ($1,$2,$3, $4) RETURNING *;

-- name: GetPriceEstimation :one
SELECT * FROM price_estimation
WHERE id = $1 LIMIT 1;

-- name: GetPriceEstimations :many
SELECT * FROM price_estimation
ORDER BY created_at;

-- name: UpdatePriceEstimation :one
UPDATE price_estimation
SET price = $1 WHERE id = $1 RETURNING *;

-- name: DeletePriceEstimation :exec
DELETE FROM price_estimation
WHERE id = $1;
