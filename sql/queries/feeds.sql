-- name: FeedCreate :one
INSERT INTO "feeds" ("created_at", "updated_at", "name", "url", "user_id")
VALUES (coalesce($1, now()), coalesce($2, now()), $3, $4, $5)
RETURNING *;

-- name: FeedGetAll :many
SELECT * FROM "feeds"
ORDER BY "name";

-- name: FeedDeleteAll :exec
DELETE FROM "feeds";