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

