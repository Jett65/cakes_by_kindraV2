// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: flavors.sql

package datebase

import (
	"context"

	"github.com/google/uuid"
)

const createFlavor = `-- name: CreateFlavor :one
INSERT INTO flavors (id, name, add_price)
VALUES ($1, $2, $3)
RETURNING id, name, add_price
`

type CreateFlavorParams struct {
	ID       uuid.UUID
	Name     string
	AddPrice string
}

func (q *Queries) CreateFlavor(ctx context.Context, arg CreateFlavorParams) (Flavor, error) {
	row := q.db.QueryRowContext(ctx, createFlavor, arg.ID, arg.Name, arg.AddPrice)
	var i Flavor
	err := row.Scan(&i.ID, &i.Name, &i.AddPrice)
	return i, err
}

const deleteFlavor = `-- name: DeleteFlavor :exec
DELETE FROM flavors WHERE id=$1
`

func (q *Queries) DeleteFlavor(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFlavor, id)
	return err
}

const getAllFlavors = `-- name: GetAllFlavors :many
SELECT id, name, add_price FROM flavors
`

func (q *Queries) GetAllFlavors(ctx context.Context) ([]Flavor, error) {
	rows, err := q.db.QueryContext(ctx, getAllFlavors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Flavor
	for rows.Next() {
		var i Flavor
		if err := rows.Scan(&i.ID, &i.Name, &i.AddPrice); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFlavorById = `-- name: GetFlavorById :one
SELECT id, name, add_price FROM flavors WHERE id=$1
`

func (q *Queries) GetFlavorById(ctx context.Context, id uuid.UUID) (Flavor, error) {
	row := q.db.QueryRowContext(ctx, getFlavorById, id)
	var i Flavor
	err := row.Scan(&i.ID, &i.Name, &i.AddPrice)
	return i, err
}

const updateFlavor = `-- name: UpdateFlavor :one
UPDATE flavors SET name=$2, add_price=$3
WHERE id=$1
RETURNING id, name, add_price
`

type UpdateFlavorParams struct {
	ID       uuid.UUID
	Name     string
	AddPrice string
}

func (q *Queries) UpdateFlavor(ctx context.Context, arg UpdateFlavorParams) (Flavor, error) {
	row := q.db.QueryRowContext(ctx, updateFlavor, arg.ID, arg.Name, arg.AddPrice)
	var i Flavor
	err := row.Scan(&i.ID, &i.Name, &i.AddPrice)
	return i, err
}
