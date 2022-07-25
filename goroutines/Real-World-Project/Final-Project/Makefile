BINARY_NAME=myapp
DSN="host=localhost port=5433 user=joshua password=capitalX123 dbname=concurrency sslmode=disable"
REDIS="127.0.0.1:6379"

# Build: Build binary
build:
	@echo "Bulding binaries..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/web
	@echo "Done Built!"

# Run: builds and runs the application
run : build
	@echo "Starting...."
	@env DSN=${DSN} REDIS=${REDIS} ./${BINARY_NAME} &
	@echo "Started!"

# Clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

# Start: an alias to run
start: run

# Stop: stops running the application
stop:
	@echo "Stopping...."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"

# Restart: stops and starts the application
restart: stop start

# Test: runs all tests
test:
	go test -v ./...

dependency:
	go get github.com/jackc/pgconn
	go get github.com/jackc/pgx/v4
	go get github.com/alexedwards/scs/v2
	go get github.com/alexedwards/scs/redisstore
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/chi/v5/middleware

docker:
	docker-compose up -d

rmdocker:
	docker-compose down


PHONY. : dependency