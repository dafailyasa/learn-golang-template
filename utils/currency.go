package util

import "fmt"

const (
	EUR = "EUR"
	IDR = "IDR"
	USD = "USD"
)

func IsSupportCurrency(currency string) error {
	switch currency {
	case EUR, IDR, USD:
		return nil
	}
	return fmt.Errorf("currently currency just support (%s,%s,%s)", USD, EUR, IDR)
}
