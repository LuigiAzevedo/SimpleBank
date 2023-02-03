startdocker:
	service docker start
startpostgres:
	docker start postgres
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=luigi -e POSTGRES_PASSWORD=for7drokidr4 -d postgres
createdb:
	docker exec -it postgres createdb --username=luigi --owner=luigi simple_bank
dropdb:
	docker exec -it postgres dropdb --username=luigi simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://luigi:for7drokidr4@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://luigi:for7drokidr4@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://luigi:for7drokidr4@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://luigi:for7drokidr4@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc: 
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/LuigiAzevedo/simplebank/db/sqlc Store

.PHONY: startdocker startpostgres postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock