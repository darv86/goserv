// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
)

const feedCreate = `-- name: FeedCreate :one
INSERT INTO "feeds" ("created_at", "updated_at", "name", "url", "user_id")
VALUES (coalesce($1, now()), coalesce($2, now()), $3, $4, $5)
RETURNING id, created_at, updated_at, name, url, user_id
`

type FeedCreateParams struct {
	Column1 interface{} `json:"column_1"`
	Column2 interface{} `json:"column_2"`
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	UserID  int64       `json:"user_id"`
}

func (q *Queries) FeedCreate(ctx context.Context, arg FeedCreateParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, feedCreate,
		arg.Column1,
		arg.Column2,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

const feedDeleteAll = `-- name: FeedDeleteAll :exec
DELETE FROM "feeds"
`

func (q *Queries) FeedDeleteAll(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, feedDeleteAll)
	return err
}

const feedGetAll = `-- name: FeedGetAll :many
SELECT id, created_at, updated_at, name, url, user_id FROM "feeds"
ORDER BY "name"
`

func (q *Queries) FeedGetAll(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, feedGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
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
