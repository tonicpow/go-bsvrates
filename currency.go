package bsvrates

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/shopspring/decimal"
)

// SatoshisPerBitcoin is the fixed amount of Satoshis per Bitcoin denomination
const SatoshisPerBitcoin = 1e8

// Regex for formatting commas
var commaRegEx = regexp.MustCompile("(\\d+)(\\d{3})")

// FormatCommas formats the integer with strings
func FormatCommas(num int) string {
	numString := strconv.Itoa(num)
	for {
		formatted := commaRegEx.ReplaceAllString(numString, "$1,$2")
		if formatted == numString {
			return formatted
		}
		numString = formatted
	}
}

// ConvertSatsToBSV converts sats to bsv
func ConvertSatsToBSV(sats int) float64 {
	return float64(sats) * 0.00000001
}

// ConvertPriceToSatoshis will get the satoshis (amount) from the current rate.
// IE: 1 BSV = $150 and you want to know what $1 is in satoshis
func ConvertPriceToSatoshis(currentRate float64, amount float64) (int64, error) {

	// Cannot use 0 (division by zero?!)
	if amount == 0 {
		return 0, fmt.Errorf("an amount must be set")
	} else if currentRate <= 0 {
		return 0, fmt.Errorf("current rate must be a positive value")
	}

	// Do conversion to satoshis (percentage) using decimal package to avoid float issues
	// => 1e8 * amount / currentRate
	// (use 1e8 since rate is in Bitcoin not Satoshis)
	satoshisDecimal := decimal.NewFromInt(SatoshisPerBitcoin).Mul(decimal.NewFromFloat(amount)).Div(decimal.NewFromFloat(currentRate))

	// Drop decimals after since can only have whole Satoshis
	return satoshisDecimal.Ceil().IntPart(), nil
}

// FormatCentsToDollars formats the integer for currency in USD (cents to dollars)
func FormatCentsToDollars(cents int) string {
	return strconv.FormatFloat(float64(cents)/100.0, 'f', 2, 64)
}

// ConvertFloatToIntUSD converts a float to int
func ConvertFloatToIntUSD(floatValue float64) int64 {
	return int64(floatValue*100 + 0.5)
}

// TransformCurrencyToInt takes the decimal format of the currency and returns the integer value
// Currently only supports USD and BSV
func TransformCurrencyToInt(decimalValue float64, currency Currency) (int64, error) {
	if currency == CurrencyDollars {
		return ConvertFloatToIntUSD(decimalValue), nil
	} else if currency == CurrencyBitcoin {
		return ConvertFloatToIntBSV(decimalValue), nil
	}
	return 0, fmt.Errorf("currency %s cannot be transformed", currency.Name())
}

// TransformIntToCurrency will take the int and return a float value.
// Currently only supports USD and BSV
func TransformIntToCurrency(intValue int, currency Currency) (string, error) {
	if currency == CurrencyDollars {
		return FormatCentsToDollars(intValue), nil
	} else if currency == CurrencyBitcoin {
		return fmt.Sprintf("%8.8f", ConvertSatsToBSV(intValue)), nil
	}
	return "", fmt.Errorf("currency %s cannot be transformed", currency.Name())
}

// ConvertFloatToIntBSV converts the BSV float value to the sats value
func ConvertFloatToIntBSV(floatValue float64) int64 {

	// Do conversion to satoshis (percentage) using decimal package to avoid float issues
	// => 1e8 * amount / currentRate
	// (use 1e8 since rate is in Bitcoin not Satoshis)
	satoshisDecimal := decimal.NewFromInt(SatoshisPerBitcoin).Mul(decimal.NewFromFloat(floatValue))

	// Drop decimals after since can only have whole Satoshis
	return satoshisDecimal.Ceil().IntPart()
}
