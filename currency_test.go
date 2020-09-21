package bsvrates

import (
	"fmt"
	"math"
	"testing"
)

// TestFormatCommas will test the method FormatCommas()
func TestFormatCommas(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		integer  int
		expected string
	}{
		{0, "0"},
		{123, "123"},
		{1234, "1,234"},
		{127127, "127,127"},
		{1271271271271, "1,271,271,271,271"},
	}

	// Test all
	for _, test := range tests {
		if output := FormatCommas(test.integer); output != test.expected {
			t.Errorf("%s Failed: [%d] inputted and [%s] expected, received: [%s]", t.Name(), test.integer, test.expected, output)
		}
	}
}

// BenchmarkFormatCommas benchmarks the method FormatCommas()
func BenchmarkFormatCommas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatCommas(10000)
	}
}

// ExampleFormatCommas example using FormatCommas()
func ExampleFormatCommas() {
	val := FormatCommas(1000000)
	fmt.Printf("%s", val)
	// Output:1,000,000
}

// TestConvertSatsToBSV will test the method ConvertSatsToBSV()
func TestConvertSatsToBSV(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		integer  int
		expected float64
	}{
		{1, 0.00000001},
		{100, 0.00000100},
		{1000, 0.0000100},
		{10000, 0.000100},
		{100000, 0.00100},
		{1000000, 0.0100},
		{10000000, 0.100},
		{100000000, 1.0},
		{12345678912, 123.45678912},
	}

	// Test all
	for _, test := range tests {
		if output := ConvertSatsToBSV(test.integer); output != test.expected {
			t.Errorf("%s Failed: [%d] inputted and [%f] expected, received: [%f]", t.Name(), test.integer, test.expected, output)
		}
	}
}

// BenchmarkConvertSatsToBSV benchmarks the method ConvertSatsToBSV()
func BenchmarkConvertSatsToBSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConvertSatsToBSV(1000)
	}
}

// ExampleConvertSatsToBSV example using ConvertSatsToBSV()
func ExampleConvertSatsToBSV() {
	val := ConvertSatsToBSV(1001)
	fmt.Printf("%f", val)
	// Output:0.000010
}

// TestConvertPriceToSatoshis will test the method ConvertPriceToSatoshis()
func TestConvertPriceToSatoshis(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		currentRate   float64
		amount        float64
		expectedSats  int64
		expectedError bool
	}{
		{157.93895102, 0.01, 6332, false},
		{158.18989656, 10, 6321517, false},
		{158.18989656, 1, 632152, false},
		{158.38610459, 1, 631369, false},
		{158.38610459, 0.01, 6314, false},
		{100, 1, 1000000, false},
		{100, 0.10, 100000, false},
		{100, 0.01, 10000, false},
		{100, 150, 150000000, false},
		{100, 100, 100000000, false},
		{1000, 1, 100000, false},
		{10000, 1, 10000, false},
		{100000, 1, 1000, false},
		{1000000, 1, 100, false},
		{0, 1, 0, true},
		{1, 0, 0, true},
		{math.Inf(1), 0, 0, true},
	}

	// Test all
	for _, test := range tests {
		if output, err := ConvertPriceToSatoshis(test.currentRate, test.amount); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%f] [%f] inputted", t.Name(), test.currentRate, test.amount)
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%f] [%f] inputted, received: [%d] error [%s]", t.Name(), test.currentRate, test.amount, output, err.Error())
		} else if output != test.expectedSats && !test.expectedError {
			t.Errorf("%s Failed: [%f] [%f] inputted and [%d] expected, received: [%d]", t.Name(), test.currentRate, test.amount, test.expectedSats, output)
		}
	}
}

// BenchmarkConvertPriceToSatoshis benchmarks the method ConvertPriceToSatoshis()
func BenchmarkConvertPriceToSatoshis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ConvertPriceToSatoshis(150, 10)
	}
}

// ExampleConvertPriceToSatoshis example using ConvertPriceToSatoshis()
func ExampleConvertPriceToSatoshis() {
	val, _ := ConvertPriceToSatoshis(150, 1)
	fmt.Printf("%d", val)
	// Output:666667
}

// TestFormatCentsToDollars will test the method FormatCentsToDollars()
func TestFormatCentsToDollars(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		integer  int
		expected string
	}{
		{0, "0.00"},
		{-1, "-0.01"},
		{127, "1.27"},
		{199276, "1992.76"},
	}

	// Test all
	for _, test := range tests {
		if output := FormatCentsToDollars(test.integer); output != test.expected {
			t.Errorf("%s Failed: [%d] inputted and [%s] expected, received: [%s]", t.Name(), test.integer, test.expected, output)
		}
	}
}

// BenchmarkFormatCentsToDollars benchmarks the method FormatCentsToDollars()
func BenchmarkFormatCentsToDollars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatCentsToDollars(1000)
	}
}

// ExampleFormatCentsToDollars example using FormatCentsToDollars()
func ExampleFormatCentsToDollars() {
	val := FormatCentsToDollars(1000)
	fmt.Printf("%s", val)
	// Output:10.00
}

// TestTransformCurrencyToInt will test the method TransformCurrencyToInt()
func TestTransformCurrencyToInt(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		decimal       float64
		currency      Currency
		expected      int64
		expectedError bool
	}{
		{0, CurrencyDollars, 0, false},
		{1.27, CurrencyDollars, 127, false},
		{01.27, CurrencyDollars, 127, false},
		{199.272, CurrencyDollars, 19927, false},
		{199.276, CurrencyDollars, 19928, false},
		{0.00000010, CurrencyBitcoin, 10, false},
		{0.000010, CurrencyBitcoin, 1000, false},
		{0.0010, CurrencyBitcoin, 100000, false},
		{0.10, CurrencyBitcoin, 10000000, false},
		{1, CurrencyBitcoin, 100000000, false},
		{0.00000010, 123, 0, true},
	}

	// todo: issue with negative floats (-1.27 = -126)

	// Test all
	for _, test := range tests {
		if output, err := TransformCurrencyToInt(test.decimal, test.currency); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%f] inputted [%s] currency", t.Name(), test.decimal, test.currency.Name())
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%f] inputted, received: [%v] error [%s]", t.Name(), test.decimal, output, err.Error())
		} else if output != test.expected && !test.expectedError {
			t.Errorf("%s Failed: [%f] inputted and [%d] expected, received: [%d]", t.Name(), test.decimal, test.expected, output)
		}
	}
}

// BenchmarkTransformCurrencyToInt benchmarks the method TransformCurrencyToInt()
func BenchmarkTransformCurrencyToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = TransformCurrencyToInt(10.00, CurrencyDollars)
	}
}

// ExampleTransformCurrencyToInt example using TransformCurrencyToInt()
func ExampleTransformCurrencyToInt() {
	val, _ := TransformCurrencyToInt(10.00, CurrencyDollars)
	fmt.Printf("%d", val)
	// Output:1000
}

// TestConvertFloatToIntBSV will test the method ConvertFloatToIntBSV()
func TestConvertFloatToIntBSV(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		float         float64
		expected      int64
		expectedError bool
	}{
		{0, 0, false},
		{1.123456789, 112345679, false},
		{0.00000001, 1, false},
		{0.00000111, 111, false},
		{-0.00000111, -111, false},
		// {math.Inf(1), -111, true}, // This will produce a panic in decimal package
	}

	// Test all
	for _, test := range tests {
		if output := ConvertFloatToIntBSV(test.float); output != test.expected && !test.expectedError {
			t.Errorf("%s Failed: [%f] inputted and [%d] expected, received: [%d]", t.Name(), test.float, test.expected, output)
		}
	}
}

// BenchmarkConvertFloatToIntBSV benchmarks the method ConvertFloatToIntBSV()
func BenchmarkConvertFloatToIntBSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConvertFloatToIntBSV(10.00)
	}
}

// ExampleConvertFloatToIntBSV example using ConvertFloatToIntBSV()
func ExampleConvertFloatToIntBSV() {
	val := ConvertFloatToIntBSV(10.01)
	fmt.Printf("%d", val)
	// Output:1001000000
}
