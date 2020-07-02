package bsvrates

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Various units used when describing a BitCoin monetary amount
const (
	UnitBSV            Unit = 0
	UnitSatoshi        Unit = -8
	SatoshisPerBitcoin      = 1e8
)

// Regex for formatting commas
var commaRegEx = regexp.MustCompile("(\\d+)(\\d{3})")

// Unit is a unit different from the base BitCoin unit
type Unit int

// String returns the unit as a string
func (u Unit) String() string {
	switch u {
	case UnitBSV:
		return CurrencyToName(CurrencyBitcoin)
	case UnitSatoshi:
		return "satoshi"
	default:
		return strconv.FormatInt(int64(u), 10) + " sats"
	}
}

// Amount represents the base BitCoin monetary unit which is a Satoshi
type Amount int64

// round converts a floating point number
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating point value representing some value in BitCoin
func NewAmount(f float64) (Amount, error) {
	switch {
	case math.IsNaN(f):
		fallthrough
	case math.IsInf(f, 1):
		fallthrough
	case math.IsInf(f, -1):
		return 0, fmt.Errorf("invalid BitCoin amount")
	}

	return round(f * SatoshisPerBitcoin), nil
}

// ToUnit converts a monetary amount counted in BitCoin base units to a floating point value representing an amount of BitCoin
func (a Amount) ToUnit(u Unit) float64 {
	return float64(a) / math.Pow10(int(u+8))
}

// Format formats a monetary amount counted in BitCoin base units as a string for a given unit
func (a Amount) Format(u Unit, showUnits bool) string {
	if showUnits {
		return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + " " + u.String()
	}
	return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64)
}

// String is the equivalent of calling Format with UnitBSV
func (a Amount) String() string {
	return a.Format(UnitBSV, true)
}

// ToSatoshi will return the satoshi value
func (a Amount) ToSatoshi() string {
	return a.Format(UnitSatoshi, false)
}

// TransformCurrencyToInt takes the decimal format of the currency and returns the integer value
// Currently only supports USD and BSV
func TransformCurrencyToInt(decimalValue float64, currency Currency) (intVal int, err error) {
	if currency == CurrencyDollars {
		intVal = ConvertFloatToIntUSD(decimalValue)
		return
	} else if currency == CurrencyBitcoin {
		intVal, err = ConvertFloatToIntBSV(decimalValue)
		return
	}
	err = fmt.Errorf("currency %s cannot be transformed", currency.Name())
	return
}

// TransformIntToCurrency will take the int and return a float value
// Currently only supports USD and BSV
func TransformIntToCurrency(intVal int, currency Currency) (decimalValue string, err error) {
	if currency == CurrencyDollars {
		decimalValue = FormatCentsToDollars(intVal)
		return
	} else if currency == CurrencyBitcoin {
		decimalValue = fmt.Sprintf("%8.8f", ConvertSatsToBSV(intVal))
		return
	}
	err = fmt.Errorf("currency %s cannot be transformed", currency.Name())
	return
}

// FormatCentsToDollars formats the integer for currency in USD  (cents to dollars)
func FormatCentsToDollars(m int) string {
	return strconv.FormatFloat(float64(m)/100.0, 'f', 2, 64)
}

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

// ConvertIntPriceToFloat converts int to float
func ConvertIntPriceToFloat(price uint64) float64 {
	// Convert integer price to decimal price without float math
	if price == 0 {
		return 0.0
	}
	priceString := strconv.FormatUint(price, 10)
	if price < 100 {
		priceString = "00" + priceString
		priceString = priceString[len(priceString)-3:]
	}
	priceChars := []rune(priceString)
	l := len(priceChars) - 1
	priceChars = append(priceChars, ' ')
	priceChars[l+1], priceChars[l] = priceChars[l], priceChars[l-1]
	priceChars[l-1] = '.'
	floatPrice, _ := strconv.ParseFloat(string(priceChars), 64)
	return floatPrice
}

// ConvertFloatToIntUSD converts a float to int
func ConvertFloatToIntUSD(floatValue float64) int {
	return int(floatValue*100 + 0.5)
}

// ConvertFloatToIntBSV converts the float of BSV to the sats value
func ConvertFloatToIntBSV(floatValue float64) (sats int, err error) {
	var amount Amount
	if amount, err = NewAmount(floatValue); err != nil {
		return
	}

	return strconv.Atoi(amount.ToSatoshi())
}

// ConvertSatsToBSV converts sats to bsv
func ConvertSatsToBSV(sats int) float64 {
	return float64(sats) * 0.00000001
}

// ConvertPriceToSatoshis will get the satoshis (amount) from the current rate
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

	// Create the float of bitcoin (percentage)
	bitcoin := 1 / (currentRate / amount)

	// Create a new amount using that bitcoin
	var bitcoinAmount Amount
	if bitcoinAmount, err = NewAmount(bitcoin); err != nil {
		return
	}

	// Convert bitcoin to satoshis
	satoshis, err = strconv.ParseInt(bitcoinAmount.ToSatoshi(), 10, 64)
	return
}
