-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "feeds" (
	"id" bigint generated always as identity,
	"created_at" timestamp not null default now(),
	"updated_at" timestamp not null default now(),
	"name" text not null,
	"url" text not null UNIQUE,
	"user_id" bigint not null references "users"("id") on delete cascade
);

ALTER TABLE "feeds" ADD CONSTRAINT "pkFeeds" PRIMARY KEY ("id");
CREATE UNIQUE INDEX "akFeeds" ON "feeds" ("name");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "feeds";
-- +goose StatementEnd
