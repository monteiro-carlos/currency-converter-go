package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/container"
)

func Handler(dep *container.Dependency) {
	router := gin.Default()
	currencyHandler := &currency.Handler{
		Service: dep.Services.Currency,
	}
	g := router.Group("/currency")
	{
		g.GET("/", currencyHandler.GetAllCurrencyRates)
		g.POST("/", currencyHandler.CreateCurrencyRateManually)
		g.GET("/update", currencyHandler.UpdateCurrencyRatesOnline)
		g.POST("/convert", currencyHandler.ConvertToAllCurrencies)
		g.GET("/code/:code", currencyHandler.GetCurrencyByCode)
	}
	router.Run(":5000")
}
