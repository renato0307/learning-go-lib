package finance

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFinanceFunctions(t *testing.T) {
	// act
	ff := NewFinanceFunctions("DummyApiUrl", "DummyApiKey")

	// assert
	assert.Equal(t, "DummyApiUrl", ff.ApiUrl)
	assert.Equal(t, "DummyApiKey", ff.ApiKey)
}

func TestNewFinanceFunctionsWithEmptyApiUrl(t *testing.T) {
	// act
	ff := NewFinanceFunctions("", "DummyApiKey")

	// assert
	assert.NotEmpty(t, ff.ApiUrl)
	assert.Contains(t, ff.ApiUrl, "fcsapi.com")
}

func TestConvertCurrency(t *testing.T) {
	// arrange
	expected := `
		{
			"code": 200,
			"info": {
				"_t": "2021-12-27 21:49:18 UTC",
				"credit_count": 1,
				"server_time": "2021-12-27 21:49:18 UTC"
			},
			"msg": "Successfully",
			"response": [
				{
					"c": "1.13268",
					"ch": "-0.00013",
					"cp": "-0.01%",
					"h": "1.13281",
					"id": "1",
					"l": "1.13246",
					"o": "1.13281",
					"s": "EUR/USD",
					"t": "1640638800",
					"tm": "2021-12-27 21:00:00",
					"up": "2021-12-27 21:49:10"
				}
			],
			"status": true
		}
	`

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	// act
	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 11.3268, math.Round(result*10000)/10000)
}

func TestConvertCurrencyWithApiError(t *testing.T) {
	// arrange
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))
	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	// act
	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestConvertCurrencyWithInvalidJsonBody(t *testing.T) {
	// arrange
	expected := "invalid json"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	// act
	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestConvertCurrencyWithInvalidPriceClose(t *testing.T) {
	// arrange
	expected := // the "c" value is not a float
		`
		{
			"code": 200,
			"info": {
				"_t": "2021-12-27 21:49:18 UTC",
				"credit_count": 1,
				"server_time": "2021-12-27 21:49:18 UTC"
			},
			"msg": "Successfully",
			"response": [
				{
					"c": "xxxx",
					"ch": "-0.00013",
					"cp": "-0.01%",
					"h": "1.13281",
					"id": "1",
					"l": "1.13246",
					"o": "1.13281",
					"s": "EUR/USD",
					"t": "1640638800",
					"tm": "2021-12-27 21:00:00",
					"up": "2021-12-27 21:49:10"
				}
			],
			"status": true
		}
	`

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	// act
	_, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error parsing the conversion data")
}

func TestConvertCurrencyWithInvalidProtocolUrl(t *testing.T) {
	// arrange
	invalidFcsapiUrl := "xxxx://someurl"
	ff := NewFinanceFunctions(invalidFcsapiUrl, "DummyApiKey")

	// act
	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
	assert.Equal(t, 0.0, result)
}

func TestConvertCurrencyWithUrlVarReplacement(t *testing.T) {
	// arrange
	invalidFcsapiUrl := "https://fcsapi.com/api-v3/forex/candle?symbol=%s/%s&period=1h&access_key=%s"
	ff := NewFinanceFunctions(invalidFcsapiUrl, "DummyApiKey")

	// act
	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	// assert
	assert.NotNil(t, err)
	assert.NotContains(t, err.Error(), "%s")
	assert.Equal(t, 0.0, result)
}

func TestReadBody(t *testing.T) {
	// arrange
	bodyReader := errReader(0)

	// act
	body, err := readBody(bodyReader)

	// assert
	assert.Error(t, err)
	assert.Empty(t, body)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func (errReader) Close() error {
	return nil
}
