-- name: CreateFlavor :one
INSERT INTO flavors (id, name, add_price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllFlavors :many
SELECT * FROM flavors;

-- name: GetFlavorById :one
SELECT * FROM flavors WHERE id=$1;

-- name: UpdateFlavor :one
UPDATE flavors SET name=$2, add_price=$3
WHERE id=$1
RETURNING *;

-- name: DeleteFlavor :exec
DELETE FROM flavors WHERE id=$1;

