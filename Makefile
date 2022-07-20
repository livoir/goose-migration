database := "goose_migration"
connection := "root:root@tcp(localhost:3306)/$(database)?parseTime=true&multiStatements=true"
dir := ./db/migrations
goose := goose -dir $(dir) mysql $(connection)

status:
	$(goose) status
create:
	$(goose) create $(name) sql
up:
	$(goose) up
down:
	$(goose) down
up-to:
	$(goose) up-to $(v)
up-1:
	$(goose) up-by-one
down-to:
	$(goose) down-to $(v)
fix:
	$(goose) fix
reset:
	$(goose) reset
version:
	$(goose) version
redo:
	$(goose) redo
test:
	go test -v -cover ./...