gen:
	mockgen -source=internal/service/service.go \
	-destination=internal/repository/mocks/mock_repository.go -package mock_repository

unit-test:
	go test ./internal/... -v

lint:
	golangci-lint run

run:
	go run cmd/main.go