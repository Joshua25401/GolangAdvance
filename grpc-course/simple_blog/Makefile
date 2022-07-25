cleanup:
	# Stop running containers (mongodb and mongoexpress)
	docker stop docker_mdb docker_me
	# Remove the containers
	docker rm docker_mdb docker_me

compose:
	# Running docker compose ...
	docker-compose up

compose_down:
	# Stopping all docker services...
	docker-compose down

tidy:
	# Tidying Go Modules
	go mod tidy

mongo_driver:
	# Getting mongo driver
	go get go.mongodb.org/mongo-driver/mongo

protoc:
	# Generate proto code golang
	protoc --go_out=. --go-grpc_out=. -Iproto/ proto/*.proto

build_all:
# Build the server
	go build -o ./bin/server/ ./server/
	# Build the client
	go build -o ./bin/client/ ./client/

run_server:
	# Running the server...
	./bin/server/server

run_client:
	# Running the client...
	./bin/client/client

PHONY.: cleanup compose tidy mongo_driver protoc run_client run_server