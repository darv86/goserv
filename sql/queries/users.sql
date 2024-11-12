-- every SqlC statement starts with comment,
-- the name of a func, which will be generated by go from this file
-- and one record wll be returned
-- name: CreateUser :one
INSERT INTO "users" (id, created_at, updated_at, name)
-- CreateUser will accept as a parameters these 4 arguments
VALUES ($1, $2, $3, $4)
RETURNING *;