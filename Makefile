postgres:
	docker run --name mypostgres --network bank-network -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 -d postgres:16.0-alpine3.18

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

bankContainer:
	docker run --name bank --network bank-network -p 8000:8000 -e dbsource=postgresql://root:password@mypostgres:5432/simplebank?sslmode=disable -e GIN_MODE=release simplebank

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Aashish-32/bank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/Aashish-32/bank/worker TaskDistributor

acessPSQL:
	docker exec -it mypostgres /bin/sh
and:
	psql -U root -d simplebank


.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock



# to connect to a db:
# \c mydatabase

# make a file executable:
# icacls wait-for-it.sh /grant Users:RX
