# Build Go binary.
build:
	go build -o bin/main ./cmd/app/main.go

# Start dev server.
start:
	air




# Run tests.
test:
	/opt/homebrew/bin/go test ./...

test_frontend:
	/opt/homebrew/bin/go test ./pkg/frontend/...

test_frontend_repository:
	/opt/homebrew/bin/go test ./pkg/frontend/adapter/repository/... -v

test_frontend_usecase:
	/opt/homebrew/bin/go test ./pkg/frontend/usecase/... -v

test_frontend_controller:
	/opt/homebrew/bin/go test ./pkg/frontend/adapter/controller/... -v

.PHONY: build install setup_db start migrate_schema seed test test_frontend test_frontend_repository test_frontend_usecase test_frontend_controller