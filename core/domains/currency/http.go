package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/log"
	"go.uber.org/zap"
)

type Handler struct {
	Service ServiceI
	log     *log.Logger
}

func NewHandler(service ServiceI, logger *log.Logger) *Handler {
	return &Handler{
		Service: service,
		log:     logger,
	}
}

// GetAllCurrencyRates godoc
// @Summary Gets all currency rates
// @Description Gets all currency rates for the previously specified currencies
// @Tags Currency
// @Produce json
// @Success 200 {object} []models.CurrencyPayload
// @Failure 404 {object} models.ErrorResponse
// @Router /currency [get].
func (h *Handler) GetAllCurrencyRates(c *gin.Context) {
	currencyRates, err := h.Service.GetAllCurrencyRates()
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, currencyRates)
}

// CreateCurrencyRateManually godoc
// @Summary Add a currency rate manually
// @Description Add a currency rate by inserting data manually
// @Tags Currency
// @Accept json
// @Produce json
// @Param currencyRate body models.CurrencyPayload true "CurrencyPayload Model"
// @Success 200 {object} string
// @Failure 404 {object} models.ErrorResponse
// @Router /currency [post].
func (h *Handler) CreateCurrencyRateManually(c *gin.Context) {
	var currencyPayload models.CurrencyPayload
	if err := c.ShouldBindJSON(&currencyPayload); err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	err := h.Service.AddNewCurrencyManually(&currencyPayload)
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Currency rate manually added",
	})
}

// UpdateCurrencyRatesOnline godoc
// @Summary Updates all currencies online simultaneously
// @Description Updates all currencies online simultaneously getting data from an external source
// @Tags Currency
// @Produce json
// @Success 200 {object} []models.CurrencyPayload
// @Failure 404 {object} models.ErrorResponse
// @Router /currency/update [get].
func (h *Handler) UpdateCurrencyRatesOnline(c *gin.Context) {
	rates, err := h.Service.UpdateCurrenciesDatabase()
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, rates)
}

// ConvertToAllCurrencies godoc
// @Summary Convert a value to all currencies
// @Description Convert a given value to all previously specified currencies
// @Tags Currency
// @Accept json
// @Produce json
// @Param value body models.ConversionRequest true "ConversionRequest Model"
// @Success 200 {object} []models.ConversionResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /currency/convert [post].
func (h *Handler) ConvertToAllCurrencies(c *gin.Context) {
	var conversionReq models.ConversionRequest
	if err := c.ShouldBindJSON(&conversionReq); err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	conversions, err := h.Service.ConvertValueToAllCurrencies(&conversionReq)
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, conversions)
}

// GetCurrencyByCode godoc
// @Summary Gets a currency rate by currency code
// @Description Gets a currency rate by the currency code given through params
// @Tags Currency
// @Produce json
// @Param code path string true "Currency Code"
// @Success 200 {object} []models.CurrencyPayload
// @Failure 404 {object} models.ErrorResponse
// @Router /currency [get].
func (h *Handler) GetCurrencyByCode(c *gin.Context) {
	code := c.Params.ByName("code")
	currencyPayload, err := h.Service.GetCurrencyRatesByCode(code)
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, currencyPayload)
}
