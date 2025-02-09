# Build and run the app
build-run:
	@echo "Building the application..."
	go build -o bin/api ./cmd/app
	@echo "Running the application..."
	./bin/api

# Run the migrations (without seeding)
migrate:
	@echo "Running migrations..."
	go run ./cmd/migration_runner/main.go

# Run the migrations and seed the database
migrate-seed:
	@echo "Running migrations and seeding the database..."
	go run ./cmd/migration_runner/main.go --seed

# Start the app using go run
run:
	@echo "Starting the app locally using go run..."
	go run ./cmd/app/main.go

# run tests
test:
	@echo "Starting the app locally using go run..."
	go test ./...
