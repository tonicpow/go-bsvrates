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
func ConvertPriceToSatoshis(currentRate float64, amount float64) (satoshis int64, err error) {

	// Cannot use 0 (division by zero?!)
	if amount == 0 {
		err = fmt.Errorf("an amount must be set")
		return
	} else if currentRate <= 0 {
		err = fmt.Errorf("current rate must be a positive value")
		return
	}

	// Do conversion to satoshis (percentage) using decimal package to avoid float issues
	// => 1e8 * amount / currentRate
	// (use 1e8 since rate is in Bitcoin not Satoshis)
	satoshisDecimal := decimal.NewFromInt(1e8).Mul(decimal.NewFromFloat(amount)).Div(decimal.NewFromFloat(currentRate))

	// Drop decimals after since can only have whole Satoshis
	satoshis = satoshisDecimal.Ceil().IntPart()
	return
}
