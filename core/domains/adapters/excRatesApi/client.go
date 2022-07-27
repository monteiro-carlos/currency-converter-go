package excRatesApi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/adapters/excRatesApi/models"
)

func GetCurrencyName(s string) string {
	names := strings.Split(s, "/")
	return names[0]
}

func GetCurrencyRates() ([]models.ExcRatesResp, error) {
	var rates []models.ExcRatesResp
	var rawResults map[string]interface{}
	//var results []map[string]interface{}
	reqURL := fmt.Sprintf("%v/%v",
		os.Getenv("FREE_EXC_RATES_API"),
		os.Getenv("CURRENCY_CODES"))

	res, err := http.Get(reqURL)
	if err != nil {
		log.Panic(err, "can't execute exchange rates api get request")
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err, "can't execute fix.io get rates request")
		return nil, err
	}
	err = json.Unmarshal(resBody, &rawResults)
	if err != nil {
		log.Panic(err, "can't parse fix.io get rates body to json")
		return nil, err
	}

	//codes := GetIndividualCurrencyCodes()
	//for _, code := range codes {
	//	result := make(map[string]interface{})
	//	resultCode := code + "BRL"
	//	rawRate := result[resultCode].(map[string]interface{})
	//	results = append(results, rawRate)
	//}
	for _, value := range rawResults {
		var rate models.ExcRatesResp
		jsonStr, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonStr, &rate); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}

	return rates, nil
}
