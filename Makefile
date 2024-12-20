removedb:
	docker rm -f postgres

postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=bye -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:bye@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	echo "y" | migrate -path db/migration -database "postgres://root:bye@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

sleep:
	sleep 5

test:
	go test -v -cover ./...

generatedb: removedb postgres sleep createdb migrateup

.PHONY: generatedb postgres createdb dropdb migrateup migratedown sqlc test
