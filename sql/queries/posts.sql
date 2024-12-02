-- name: PostCreate :one
INSERT INTO "posts" (
	"title",
	"url",
	"feed_id"
)
VALUES ($1, $2, $3)
RETURNING *;

-- name: PostByUser :many
SELECT * FROM "posts"
WHERE "feed_id" IN (
	SELECT "id" FROM "feeds"
	WHERE "user_id" = $1
);