// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const create = `-- name: Create :one
INSERT INTO "users" (created_at, updated_at, name)
VALUES (coalesce($2, now()), coalesce($3, now()), $1)
RETURNING id, created_at, updated_at, name
`

type CreateParams struct {
	Name    string
	Column2 interface{}
	Column3 interface{}
}

// every SqlC statement starts with comment,
// the name of a func, which will be generated by go from this file
// and one record wll be returned
// func Create will only one parameter with string type,
// instead of CreateParams type
// INSERT INTO "users" (name)
// Create will accept as a parameters these 4 arguments
// COALESCE is the standard sql function,
// which returns 1st not null value
// VALUES ($1, $2, $3, $4)
// syntax * returns all parameters
// RETURNING *;
func (q *Queries) Create(ctx context.Context, arg CreateParams) (User, error) {
	row := q.db.QueryRowContext(ctx, create, arg.Name, arg.Column2, arg.Column3)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getAll = `-- name: GetAll :many
SELECT id, created_at, updated_at, name FROM "users"
ORDER BY "name"
`

func (q *Queries) GetAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
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

const getById = `-- name: GetById :one
SELECT id, created_at, updated_at, name FROM "users"
WHERE "id" = $1
`

func (q *Queries) GetById(ctx context.Context, id sql.NullInt64) (User, error) {
	row := q.db.QueryRowContext(ctx, getById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
