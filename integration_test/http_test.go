package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
	"github.com/stretchr/testify/assert"
)

func TestAddCurrencyRateManually(t *testing.T) {
	tests := []struct {
		name     string
		input    models.CurrencyPayload
		wantCode int
		resp     models.CurrencyPayload
	}{
		{
			name:     "adding currency rate manually",
			wantCode: http.StatusOK,
			input:    currencyPayload(),
			resp:     currencyPayload(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dep, err := setupEnviroment()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			payloadRequestJson, _ := json.Marshal(currencyPayload())
			dep.router.POST("/currency", dep.currencyHandler.CreateCurrencyRateManually)
			req, _ := http.NewRequest("POST", "/currency", bytes.NewBuffer(payloadRequestJson))

			res := httptest.NewRecorder()
			dep.router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := new(models.CurrencyPayload)
			if err := json.Unmarshal(res.Body.Bytes(), actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, want.Currency.Name, actual.Currency.Name)
			assert.Equal(t, want.Currency.Code, actual.Currency.Code)
			assert.Equal(t, want.Rate, actual.Rate)
		})
	}
}

func TestUpdateCurrencyRatesOnline(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		resp     []models.CurrencyPayload
	}{
		{
			name:     "updating all currency rates online",
			wantCode: http.StatusOK,
			resp:     make([]models.CurrencyPayload, getCurrencyCodesQuantity()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dep, err := setupEnviroment()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			dep.router.GET("/currency/update", dep.currencyHandler.UpdateCurrencyRatesOnline)
			req, _ := http.NewRequest("GET", "/currency/update", nil)

			res := httptest.NewRecorder()
			dep.router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := make([]models.CurrencyPayload, 0)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, len(want), len(actual))
			assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(want))
		})
	}
}

func TestGetAllCurrencyRates(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		resp     []models.CurrencyPayload
	}{
		{
			name:     "getting all currency rates",
			wantCode: http.StatusOK,
			resp:     make([]models.CurrencyPayload, getCurrencyCodesQuantity()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dep, err := setupEnviroment()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			dep.router.GET("/currency", dep.currencyHandler.GetAllCurrencyRates)
			req, _ := http.NewRequest("GET", "/currency", nil)

			res := httptest.NewRecorder()
			dep.router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := make([]models.CurrencyPayload, 0)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, len(want), len(actual))
			assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(want))
		})
	}
}

func TestConvertToAllCurrencyRates(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		input    models.ConversionRequest
		resp     []models.ConversionResponse
	}{
		{
			name:     "converting to all currency rates",
			wantCode: http.StatusOK,
			input:    conversionRequest(),
			resp:     make([]models.ConversionResponse, getCurrencyCodesQuantity()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dep, err := setupEnviroment()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			conversionRequestJson, _ := json.Marshal(conversionRequest())
			dep.router.POST("/currency/convert", dep.currencyHandler.ConvertToAllCurrencies)
			req, _ := http.NewRequest("POST", "/currency/convert", bytes.NewBuffer(conversionRequestJson))

			res := httptest.NewRecorder()
			dep.router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := make([]models.CurrencyPayload, 0)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, len(want), len(actual))
			assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(want))
		})
	}
}

func TestGetCurrencyRateByCode(t *testing.T) {
	tests := []struct {
		name       string
		wantCode   int
		inputParam string
		resp       models.CurrencyPayload
	}{
		{
			name:       "getting USD currency rate",
			wantCode:   http.StatusOK,
			inputParam: "USD",
			resp:       currencyPayloadRespWithCode("USD"),
		},
		{
			name:       "getting EUR currency rate",
			wantCode:   http.StatusOK,
			inputParam: "EUR",
			resp:       currencyPayloadRespWithCode("EUR"),
		},
		{
			name:       "getting INR currency rate",
			wantCode:   http.StatusOK,
			inputParam: "INR",
			resp:       currencyPayloadRespWithCode("INR"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dep, err := setupEnviroment()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			dep.router.GET("/currency/:code", dep.currencyHandler.GetCurrencyByCode)
			req, _ := http.NewRequest("GET", "/currency/"+test.inputParam, nil)

			res := httptest.NewRecorder()
			dep.router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := new(models.CurrencyPayload)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(want))
			assert.Equal(t, want.Currency.Code, actual.Currency.Code)
			assert.Equal(t, want.Currency.Name, actual.Currency.Name)
		})
	}
}
