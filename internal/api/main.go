package main

import (
	"log"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/container"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/routes"
)

func main() {
	dep, err := container.New()
	if err != nil {
		log.Fatal(err)
	}
	routes.Handler(dep)
}
