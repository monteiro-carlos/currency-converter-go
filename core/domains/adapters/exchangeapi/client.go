package exchangeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/adapters/exchangeapi/models"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/log"
	"go.uber.org/zap"
)

type Client struct {
	log *log.Logger
}

func NewClient(
	logger *log.Logger,
) (*Client, error) {
	return &Client{
		log: logger,
	}, nil
}

func (c *Client) GetCurrencyName(s string) string {
	names := strings.Split(s, "/")
	return names[0]
}

func (c *Client) GetCurrencyRates() ([]models.ExcRatesResp, error) {
	rates := make([]models.ExcRatesResp, 0)
	var rawResults map[string]interface{}
	reqURL := fmt.Sprintf("%v/%v",
		os.Getenv("FREE_EXC_RATES_API"),
		os.Getenv("CURRENCY_CODES"))

	res, err := http.Get(reqURL) //nolint:gosec
	if err != nil {
		c.log.Zap.Error("can't execute exchange rates api get request",
			zap.Error(err))
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.log.Zap.Error("can't execute exchange rates api get rates request",
			zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(resBody, &rawResults)
	if err != nil {
		c.log.Zap.Error("can't parse exchange rates api get rates body to json",
			zap.Error(err))
		return nil, err
	}

	for _, value := range rawResults {
		var rate models.ExcRatesResp
		jsonStr, err := json.Marshal(value)
		if err != nil {
			c.log.Zap.Error("error", zap.Error(err))
			return nil, err
		}
		if err := json.Unmarshal(jsonStr, &rate); err != nil {
			c.log.Zap.Error("error", zap.Error(err))
			return nil, err
		}
		rates = append(rates, rate)
	}

	return rates, nil
}
