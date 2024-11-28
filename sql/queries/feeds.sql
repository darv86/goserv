-- name: FeedCreate :one
INSERT INTO "feeds" ("created_at", "updated_at", "name", "url", "user_id")
VALUES (coalesce($1, now()), coalesce($2, now()), $3, $4, $5)
RETURNING *;

-- name: FeedGetAll :many
SELECT * FROM "feeds"
ORDER BY "name";

-- name: FeedDeleteAll :exec
DELETE FROM "feeds";

-- name: FeedMineDeleteById :one
DELETE FROM "feeds"
WHERE "id" = $1 AND "user_id" = $2
RETURNING *;

-- name: FeedMineGetAll :many
SELECT * FROM "feeds"
WHERE "user_id" = $1
ORDER BY "name";