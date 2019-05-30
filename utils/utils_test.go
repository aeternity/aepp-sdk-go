package utils

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestBigInt(t *testing.T) {
	var customBig = BigInt{new(big.Int)}
	customBig.SetUint64(math.MaxUint64)
	fmt.Println(customBig)

	var resultBig = BigInt{new(big.Int)}
	resultBig.Add(customBig.Int, big.NewInt(1000))
	fmt.Println(resultBig)
}

func TestBigIntNewStr(t *testing.T) {
	// Use the convenience function NewBigIntFromString to set the number
	a, err := NewBigIntFromString("20000000000000000000") // 2e19
	if err != nil {
		t.Fatal(err)
	}

	// Make a BigInt the old fashioned, raw way.
	ex := BigInt{new(big.Int)}
	ex.SetString("20000000000000000000", 10)

	// They should equal each other.
	if ex.Cmp(a.Int) != 0 {
		t.Fatalf("Expected 20000000000000000000 but got %v", a.String())
	}

}

func TestBigIntLargerOrEqualToZero(t *testing.T) {
	var amount = NewBigInt()
	var tests = []struct {
		input    int64
		expected bool
	}{
		{math.MinInt64, false},
		{-1, false},
		{0, true},
		{1, true},
		{math.MaxInt64, true},
	}

	for _, test := range tests {
		amount.SetInt64(test.input)
		err := amount.LargerOrEqualToZero()
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %#v received", test.input, test.expected, err)
		}
	}
}

func TestBigIntLargerThanZero(t *testing.T) {
	var amount = NewBigInt()
	var tests = []struct {
		input    int64
		expected bool
	}{
		{math.MinInt64, false},
		{-1, false},
		{0, false},
		{1, true},
		{math.MaxInt64, true},
	}

	for _, test := range tests {
		amount.SetInt64(test.input)
		err := amount.LargerThanZero()
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %#v received", test.input, test.expected, err)
		}
	}
}

// Code using BigInt.UnmarshalJSON should not need to know that it has to initialize
// the big.Int inside BigInt.
func TestBigIntUnmarshalJSON(t *testing.T) {
	maxuint64AsString := []byte("18446744073709551615")
	maxuint64AsBigInt := RequireBigIntFromString("18446744073709551615")

	overuint64AsString := []byte("28446744073709551615")
	overuint64AsBigInt := RequireBigIntFromString("28446744073709551615")

	negBigIntAsString := []byte("-1")
	negBigIntAsBigInt := RequireBigIntFromString("-1")

	realLifeBalanceString := []byte("966370000100000000000001")
	realLifeBalanceBigInt := RequireBigIntFromString("966370000100000000000001")
	var amount = BigInt{} // leave Int uninitialized, to ensure that UnmarshalJSON will initialize it.
	var tests = []struct {
		input    []byte
		expected *BigInt
	}{
		{[]byte("0"), NewBigInt()},
		{maxuint64AsString, maxuint64AsBigInt},
		{overuint64AsString, overuint64AsBigInt},
		{negBigIntAsString, negBigIntAsBigInt},
		{realLifeBalanceString, realLifeBalanceBigInt},
	}

	for _, test := range tests {
		amount.UnmarshalJSON(test.input)
		if amount.Cmp(test.expected.Int) != 0 {
			t.Errorf("Test Failed: %v inputted, %v expected, %#v received", test.input, test.expected, amount)
		}
	}
}
