.PHONY: it, schema

it:
	go build && ./goserv

schema:
	pg_dump -U darv -h localhost -d goserv --schema-only -f sql/schemas/schema.sql