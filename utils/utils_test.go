package utils

import (
	"errors"
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	var amount = BigInt{&big.Int{}}
	var tests = []struct {
		input    int64
		expected error
	}{
		{math.MinInt64, errors.New("the error message is not important")},
		{-1, errors.New("the error message is not important")},
		{0, errors.New("the error message is not important")},
		{1, nil},
		{math.MaxInt64, nil},
	}

	for _, test := range tests {
		amount.SetInt64(test.input)
		err := amount.validate()
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %#v received", test.input, test.expected, err)
		}
	}
}
