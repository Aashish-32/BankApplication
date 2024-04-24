postgres:
	docker run --name mypostgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 -d postgres:16.0-alpine3.18

createdb: 
	docker exec -it mypostgres createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it mypostgres dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc




