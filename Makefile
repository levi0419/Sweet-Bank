postgres:
	docker run --name personal -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password1234 -d postgres:14-alpine

createdb:
	docker exec -it personal createdb --username=root --owner=root personal_test

dropdb:
	docker exec -it personal dropdb personal_test

migrateup:
	migrate -path db/migration -database "postgresql://root:password1234@localhost:5433/personal_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password1234@localhost:5433/personal_test?sslmode=disable" -verbose down

sqlcpull:
	docker pull kjconroy/sqlc

sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlcpull test
