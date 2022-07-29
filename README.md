
# Currency Converter API

An API to convert currencies from BRL (Brazilian Reais) to another currencies.

## Stack

* Golang
* PostgreSQL
* Gin
* Gorm
* Docker
* Swagger
* golangci-lint
* zap/logger

## API Docs

- After running project, the swagger documentation is available at: [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)
- Also, there's a Postman Collection inside `internal/docs` folder.



## Running

Clone project

```bash
  git clone https://github.com/monteiro-carlos/eng-gruposbf-backend-golang.git
```

Enter projects directory

```bash
  cd eng-gruposbf-backend-golang
```

Generate swagger files

```bash
  make swagger
```

Install dependencies

```bash
  go mod tidy
```

Start database

```bash
  make docker-db
```

Setup envs from `./.env.sample`


Inicialize project from `internal/api/main.go`


## Runnning tests

To run all tests, just run the following

```bash
  make test
```


## Reference

- [Golang Standard Project Layout](https://github.com/golang-standards/project-layout)
- Onion Architecture
- Microservices Architecture
## Authors

- [@monteiro-carlos](https://github.com/monteiro-carlos/)
