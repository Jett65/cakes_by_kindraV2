// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: cakes.sql

package datebase

import (
	"context"

	"github.com/google/uuid"
)

const createCake = `-- name: CreateCake :one
INSERT INTO cakes (id, type, layer_number, tiere_number, size, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, type, layer_number, tiere_number, size, price
`

type CreateCakeParams struct {
	ID          uuid.UUID
	Type        string
	LayerNumber int16
	TiereNumber int16
	Size        string
	Price       string
}

func (q *Queries) CreateCake(ctx context.Context, arg CreateCakeParams) (Cake, error) {
	row := q.db.QueryRowContext(ctx, createCake,
		arg.ID,
		arg.Type,
		arg.LayerNumber,
		arg.TiereNumber,
		arg.Size,
		arg.Price,
	)
	var i Cake
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.LayerNumber,
		&i.TiereNumber,
		&i.Size,
		&i.Price,
	)
	return i, err
}

const deleteCake = `-- name: DeleteCake :exec
DELETE FROM cakes WHERE id=$1
`

func (q *Queries) DeleteCake(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCake, id)
	return err
}

const getAllCakes = `-- name: GetAllCakes :many
SELECT id, type, layer_number, tiere_number, size, price FROM cakes
`

func (q *Queries) GetAllCakes(ctx context.Context) ([]Cake, error) {
	rows, err := q.db.QueryContext(ctx, getAllCakes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cake
	for rows.Next() {
		var i Cake
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.LayerNumber,
			&i.TiereNumber,
			&i.Size,
			&i.Price,
		); err != nil {
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

const getCakeById = `-- name: GetCakeById :one
SELECT id, type, layer_number, tiere_number, size, price FROM cakes WHERE id=$1
`

func (q *Queries) GetCakeById(ctx context.Context, id uuid.UUID) (Cake, error) {
	row := q.db.QueryRowContext(ctx, getCakeById, id)
	var i Cake
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.LayerNumber,
		&i.TiereNumber,
		&i.Size,
		&i.Price,
	)
	return i, err
}

const updateCake = `-- name: UpdateCake :one
UPDATE cakes SET type=$2, layer_number=$3, tiere_number=$4, size=$5, price=$6
WHERE id = $1
RETURNING id, type, layer_number, tiere_number, size, price
`

type UpdateCakeParams struct {
	ID          uuid.UUID
	Type        string
	LayerNumber int16
	TiereNumber int16
	Size        string
	Price       string
}

func (q *Queries) UpdateCake(ctx context.Context, arg UpdateCakeParams) (Cake, error) {
	row := q.db.QueryRowContext(ctx, updateCake,
		arg.ID,
		arg.Type,
		arg.LayerNumber,
		arg.TiereNumber,
		arg.Size,
		arg.Price,
	)
	var i Cake
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.LayerNumber,
		&i.TiereNumber,
		&i.Size,
		&i.Price,
	)
	return i, err
}
