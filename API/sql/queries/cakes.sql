-- name: CreateCake :one
INSERT INTO cakes (id, type, layer_number, tiere_number, size, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllCakes :many
SELECT * FROM cakes;

-- name: GetCakeById :one
SELECT * FROM cakes WHERE id=$1;

-- name: DeleteCake :exec
DELETE FROM cakes WHERE id=$1;

-- name: UpdateCake :one
UPDATE cakes SET type=$2, layer_number=$3, tiere_number=$4, size=$5, price=$6
WHERE id = $1
RETURNING *;
