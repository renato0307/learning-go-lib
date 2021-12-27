package finance

type Interface interface {
	ConvertCurrency(
		from string,
		to string,
		amount float64) (float64, error)
}

type FinanceFunctions struct {
	ApiUrl string
	ApiKey string
}
