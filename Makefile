postgres:
	docker run --name mypostgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 -d postgres:16.0-alpine3.18

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it mypostgres dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down 1


# in powershell::
sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

acessPSQL:
	docker exec -it mypostgres /bin/sh
and:
	psql -U root -d simplebank


.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server



# to connect to a db:
# \c mydatabase
