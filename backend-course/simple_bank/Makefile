pull_postgre:
	docker pull postgres:14.4

init_postgre:
	docker run --name docker_postgres -p 5433:5432 -e POSTGRES_USER=joshua -e POSTGRES_PASSWORD=capitalX123 -d postgres:14.4

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://joshua:capitalX123@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://joshua:capitalX123@localhost:5433/simple_bank?sslmode=disable" -verbose down

createdb:
	docker exec -it docker_postgres createdb --username=joshua simple_bank

deletedb:
	docker exec -it docker_postgres dropdb --username=joshua simple_bank

sqlc:
	sqlc generate

cleanup:
	docker stop docker_postgres
	docker rm docker_postgres

.PHONY: pull_postgre init_postgre migrate migrateup migratedown createdb deletedb cleanup