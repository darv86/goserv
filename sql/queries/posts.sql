-- name: PostCreate :one
INSERT INTO "posts" (
	"title",
	"url",
	"feed_id"
)
VALUES ($1, $2, $3)
RETURNING *;
