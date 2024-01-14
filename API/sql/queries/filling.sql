-- name: CreateFilling :one
INSERT INTO filling (id, name, price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllFillings :many
SELECT * FROM filling;

-- name: GetFillingById :one
SELECT * FROM filling WHERE id=$1;

-- name: UpdateFilling :one
UPDATE filling SET name=$2, price=$3
WHERE id=$1
RETURNING *;

-- name: DeleteFilling :exec
DELETE FROM filling WHERE id=$1;
