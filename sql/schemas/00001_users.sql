-- comment below is used by goose to migrate up
-- (the last migration version for users schema)
-- +goose Up
CREATE TABLE IF NOT EXISTS "users" (
	"id" bigint generated always as identity,
	"created_at" timestamp not null,
	"updated_at" timestamp not null,
	"name" text not null
);

ALTER TABLE "users" ADD CONSTRAINT "pkUsers" PRIMARY KEY ("id");
CREATE UNIQUE INDEX "akUsers" ON "users" ("name");

-- +goose Down
DROP TABLE IF EXISTS "users";
