package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health/models"
)

type Handler struct {
	Service ServiceI
}

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

	c.JSON(http.StatusOK, &readiness)
}

func (h *Handler) LivenessHandler(c *gin.Context) {
	liveness := models.HealthCheckLiveness{
		Healthy:   true,
		CheckedAt: time.Now(),
		Status:    http.StatusOK,
	}

	c.JSON(http.StatusOK, &liveness)
}
