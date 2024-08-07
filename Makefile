# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/ticketopia-user-service/main.go

# Run the application
run:
	@go run cmd/ticketopia-user-service/main.go



# Test the application
test:
	@echo "Testing..."
	@go test ./... -v



# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


migratedown:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/ticketopia-user-service?sslmode=disable" -verbose down

migratedown1:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/ticketopia-user-service?sslmode=disable" -verbose down 1

proto:
	rm -f pkg/pb/*.go
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
	proto/user/v1/*.proto

.PHONY: all build run test clean migratedown migratedown1 proto
