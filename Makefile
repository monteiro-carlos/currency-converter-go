.PHONY: app-run
app-run:
	@go run ./internal/api/main.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix

.PHONY: docker-db
docker-db:
	@docker-compose -f ./docker-compose-db.yml up

.PHONY: docker-db-stop
docker-db-stop:
	@docker-compose -f ./docker-compose-db.yml stop

.PHONY: generate
generate: swagger
	@go generate ./...

.PHONY: swagger
swagger:
	@go run github.com/swaggo/swag/cmd/swag@v1.7.4 init -g internal/api/main.go -o internal/swagger/docs

.PHONY: docker-db-test
docker-db-test:
	@docker-compose -f ./integration_test/docker-compose.yml up

.PHONY: docker-db-test-stop
docker-db-test-stop:
	@docker-compose -f ./integration_test/docker-compose.yml stop

.PHONY: test
test:
	@go test -v ./...