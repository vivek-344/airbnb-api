postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=vivek -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root room_db

dropdb:
	docker exec -it postgres dropdb room_db

migrateup:
	migrate -path db/migration -database "postgresql://root:vivek@localhost:5432/room_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:vivek@localhost:5432/room_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server