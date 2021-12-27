package finance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const fcsapiUrl string = "https://fcsapi.com/api-v3/forex/candle?symbol=%s/%s&period=1h&access_key=%s"

// fasApiResponse represents the return from the Last Candle API from FSC.
// More details in the docs: https://fcsapi.com/document/forex-api#lastcandle
type fsaApiResponse struct {
	Code     int                     `json:"code"`
	Info     fsaApiResponseInfo      `json:"info"`
	Message  string                  `json:"msg"`
	Response []fsaApiResponseDetails `json:"response"`
}

type fsaApiResponseInfo struct {
	T           string `json:"_t"`
	CreditCount int    `json:"credit_count"`
	ServerTime  string `json:"server_time"`
}

type fsaApiResponseDetails struct {
	PriceClose           string `json:"c"`
	ChangeInOneDayCandle string `json:"ch"`
	ChangeInPercentage   string `json:"cp"`
	High                 string `json:"h"`
	ID                   string `json:"id"`
	Low                  string `json:"l"`
	Open                 string `json:"o"`
	Symbol               string `json:"s"`
	WhenUnix             string `json:"t"`
	WhenUtc              string `json:"tm"`
	WhenLastUpdateUtc    string `json:"up"`
}

// NewFinanceFunctions create a new FinanceFunctions instance. If the apiUrl is
// empty a default value will be set.
func NewFinanceFunctions(apiUrl, apiKey string) FinanceFunctions {
	ff := FinanceFunctions{
		ApiUrl: apiUrl,
		ApiKey: apiKey,
	}

	if ff.ApiUrl == "" {
		ff.ApiUrl = fcsapiUrl // sets the default value
	}

	return ff
}

// ConvertCurrency converts an amount from one currency into another using
// the https://fcsapi.com/ last candle API.
func (ff *FinanceFunctions) ConvertCurrency(
	from string,
	to string,
	amount float64) (float64, error) {

	response, err := ff.callLastCandleApi(from, to)
	if err != nil {
		return 0, err
	}

	priceClose, err := strconv.ParseFloat(response.Response[0].PriceClose, 64)
	if err != nil {
		err = fmt.Errorf("error parsing the conversion data: %s", err.Error())
		return 0, err
	}

	convertedAmount := priceClose * amount

	return convertedAmount, nil
}

// callLastCandleApi calls the Last Candle API, parsing the response into a
// struct.
func (ff *FinanceFunctions) callLastCandleApi(from, to string) (fsaApiResponse, error) {
	response := fsaApiResponse{}

	// calls the API and checks the result for errors
	url := ff.ApiUrl
	if strings.Count(ff.ApiUrl, "%s") > 0 { // when using httptest the URL contains no %s
		url = fmt.Sprintf(ff.ApiUrl, from, to, ff.ApiKey)
	}

	httpResponse, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("error getting the conversion data: %s", err.Error())
		return response, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		err := fmt.Errorf("error getting the conversion data: %d", httpResponse.StatusCode)
		return response, err
	}

	defer httpResponse.Body.Close()

	// reads and response and converts JSON string to a struct
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		err = fmt.Errorf("error reading the conversion data: %s", err.Error())
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("error parsing the conversion data: %s", err.Error())
		return response, err
	}

	return response, nil
}
