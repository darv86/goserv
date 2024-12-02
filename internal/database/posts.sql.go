// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
)

const postByUser = `-- name: PostByUser :many
SELECT id, title, url, feed_id FROM "posts"
WHERE "feed_id" IN (
	SELECT "id" FROM "feeds"
	WHERE "user_id" = $1
)
`

func (q *Queries) PostByUser(ctx context.Context, userID int64) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, postByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.FeedID,
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

const postCreate = `-- name: PostCreate :one
INSERT INTO "posts" (
	"title",
	"url",
	"feed_id"
)
VALUES ($1, $2, $3)
RETURNING id, title, url, feed_id
`

type PostCreateParams struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	FeedID int64  `json:"feed_id"`
}

func (q *Queries) PostCreate(ctx context.Context, arg PostCreateParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, postCreate, arg.Title, arg.Url, arg.FeedID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Url,
		&i.FeedID,
	)
	return i, err
}
