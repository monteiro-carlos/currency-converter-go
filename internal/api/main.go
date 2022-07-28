package main

import (
	"log"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/container"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/routes"
)

// @title Currency Converter API
// @version 1.0
// @description This API is used to consult and convert from a wide range of currencies to BRL
// @termsOfService http://swagger.io/terms/
// @contact.name Carlos Fernandes
// @query.collection.format multi
// @in header
// @schemes http https.
func main() {
	dep, err := container.New()
	if err != nil {
		log.Fatal(err)
	}
	routes.Handler(dep)
}
