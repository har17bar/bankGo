postgres:
	 docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createDb:
	docker exec -it postgres12 createdb simple_bank
dropDb:
	docker exec -it postgres12 dropdb simple_bank
migrateCreate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateUp:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up
migrateDown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down
migration_fix:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createDb dropDb migrateUp migrateDown sqlc test
