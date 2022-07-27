.PHONY: app-run
app-run:
	@go run main.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix

.PHONY: docker-db
docker-db:
	@docker-compose -f ./docker-compose-db.yml up

.PHONY: swagger
swagger:
	@go run github.com/swaggo/swag/cmd/swag init -g internal/api/main.go -o internal/swagger/docs --parseDependency --parseInternal --parseDepth 1