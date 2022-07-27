package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/container"
	docs "github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/swagger/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	healthHandler := &health.Handler{
		Service: dep.Services.Health,
	}
	hc := router.Group("/health")
	{
		hc.GET("/liveness", healthHandler.LivenessHandler)
		hc.GET("/readiness", healthHandler.ReadinessHandler)
	}
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":5000")
}

//docs "github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/swagger/docs"
//swaggerfiles "github.com/swaggo/files"
//ginSwagger "github.com/swaggo/gin-swagger"
