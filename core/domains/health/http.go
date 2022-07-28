package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health/models"
	_ "github.com/swaggo/swag/example/celler/httputil" // necessary for swagger
)

type Handler struct {
	Service ServiceI
}

// ReadinessHandler godoc
// @Summary Shows if all the API's components are up
// @Description Shows if all the API's components are up
// @Produce json
// @Success 200 {object} models.HealthCheckReadiness
// @Failure 500 {object} models.HealthCheckReadiness
// @Tags HealthCheck
// @Router /health/readiness [get].
func (h *Handler) ReadinessHandler(c *gin.Context) {
	databaseStatus := h.Service.CheckDatabase()

	checks := []models.Check{
		databaseStatus,
	}

	readinessHealth := true
	statusCode := http.StatusOK

	for _, check := range checks {
		if !check.Healthy {
			readinessHealth = false
			statusCode = http.StatusInternalServerError
		}
	}

	readiness := models.HealthCheckReadiness{
		Checks:    checks,
		Healthy:   readinessHealth,
		CheckedAt: time.Now(),
		Status:    statusCode,
	}

	c.JSON(statusCode, &readiness)
}

// LivenessHandler godoc
// @Summary Shows if this API is up
// @Description Shows if this API is up
// @Produce json
// @Success 202 {object} models.HealthCheckLiveness
// @Tags HealthCheck
// @Router /health/liveness [get].
func (h *Handler) LivenessHandler(c *gin.Context) {
	liveness := models.HealthCheckLiveness{
		Healthy:   true,
		CheckedAt: time.Now(),
		Status:    http.StatusOK,
	}

	c.JSON(http.StatusOK, &liveness)
}
