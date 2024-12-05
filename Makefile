postgres:
	docker run --name postgres-12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=bye -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:bye@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	echo "y" | migrate -path db/migration -database "postgres://root:bye@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown