-- name: CreateFrosting :one
INSERT INTO frosting (id, name, add_price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllFrostings :many
SELECT * FROM frosting;

-- name: GetFrostingById :one
SELECT * FROM frosting WHERE id=$1;

-- name: UpdateFrosting :one
UPDATE frosting SET name=$2, add_price=$3
WHERE id=$1
RETURNING *;

-- name: DeleteFrosting :exec
DELETE FROM frosting WHERE id=$1;
