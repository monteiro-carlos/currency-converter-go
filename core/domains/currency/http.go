package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
)

type Handler struct {
	Service ServiceI
}

func (h *Handler) GetAllCurrencyRates(c *gin.Context) {
	currencyRates, err := h.Service.GetAllCurrencyRates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, &currencyRates)
}

func (h *Handler) CreateCurrencyRateManually(c *gin.Context) {
	var currencyPayload models.CurrencyPayload
	if err := c.ShouldBindJSON(&currencyPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err := h.Service.AddNewCurrencyManually(&currencyPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Currency rate manually added",
	})
}