package money

import (
	"regexp"
	"strconv"
)

type CurrencyPair struct {
	Base    Currency
	Counter Currency
	Ratio   float64
}

// Creates a new currency pair based on the given ratio
func NewCurrencyPair(base, counter Currency, ratio float64) *CurrencyPair {
	return &CurrencyPair{
		base,
		counter,
		ratio,
	}
}

// Creates the currency pair from ISO string
// https://en.wikipedia.org/wiki/Currency_pair
// https://en.wikipedia.org/wiki/ISO_4217
func NewCurrencyPairFromIso(iso string) (*CurrencyPair, error) {
	regex := regexp.MustCompile("([A-Z]{2,3})/([A-Z]{2,3})\\s([0-9]*\\.?[0-9]+)$")

	if !regex.MatchString(iso) {
		return nil, ErrInvalidIsoPair
	}

	matches := regex.FindAllStringSubmatch(iso, 3)[0][1:]

	ratio, err := strconv.ParseFloat(matches[2], 64)

	if err != nil {
		return nil, err
	}

	return NewCurrencyPair(
		Currency(matches[0]),
		Currency(matches[1]),
		float64(ratio),
	), nil
}

// Converts from base to counter currency
func (p CurrencyPair) Convert(money *Money) (*Money, error) {
	if !money.Currency.Equals(p.Base) {
		return nil, ErrNotSameCurrency
	}

	return money.Convert(p.Counter, p.Ratio), nil
}
