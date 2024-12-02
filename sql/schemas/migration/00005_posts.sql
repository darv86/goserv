-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "posts" (
	"id" bigint generated always as identity,
	"title" text not null,
	"url" text not null UNIQUE,
	"feed_id" bigint not null references "feeds"("id") on delete cascade
);

ALTER TABLE "posts" ADD CONSTRAINT "pkPosts" PRIMARY KEY ("id");
CREATE UNIQUE INDEX "akPosts" ON "posts" ("url");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "posts";
-- +goose StatementEnd
