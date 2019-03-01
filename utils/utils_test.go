package utils

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestBigInt(t *testing.T) {
	realBig := big.Int{}
	realBig.SetUint64(math.MaxUint64)
	result := big.Int{}
	fmt.Println(realBig, result)
	fmt.Println(result.Add(&realBig, big.NewInt(10)))

	var customBig = BigInt{&big.Int{}}
	customBig.SetUint64(math.MaxUint64)
	fmt.Println(customBig)

	var resultBig = BigInt{&big.Int{}}
	resultBig.Add(customBig.Int, big.NewInt(1000))
	fmt.Println(resultBig)
}

func TestBigIntNewStr(t *testing.T) {
	a, err := NewBigIntStr("20000000000000000000") // 2e19
	if err != nil {
		t.Fatal(err)
	}

	ex := BigInt{Int: &big.Int{}}
	ex.SetString("20000000000000000000", 10)
	if ex.Cmp(a.Int) != 0 {
		t.Fatalf("Expected 20000000000000000000 but got %v", a.String())
	}

}

func TestBigIntValidate(t *testing.T) {
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

func TestBigIntUnmarshalJSON(t *testing.T) {
	maxuint64AsString := []byte("18446744073709551615")
	maxuint64AsBigInt := BigInt{&big.Int{}}
	maxuint64AsBigInt.SetString("18446744073709551615", 10)

	overuint64AsString := []byte("28446744073709551615")
	overuint64AsBigInt := BigInt{&big.Int{}}
	overuint64AsBigInt.SetString("28446744073709551615", 10)

	negBigIntAsString := []byte("-1")
	negBigIntAsBigInt := BigInt{&big.Int{}}
	negBigIntAsBigInt.SetString("-1", 10)

	var amount = BigInt{&big.Int{}}
	var tests = []struct {
		input    []byte
		expected BigInt
	}{
		{[]byte("0"), BigInt{&big.Int{}}},
		{maxuint64AsString, maxuint64AsBigInt},
		{overuint64AsString, overuint64AsBigInt},
		{negBigIntAsString, negBigIntAsBigInt},
	}

	for _, test := range tests {
		amount.UnmarshalJSON(test.input)
		if amount.Cmp(test.expected.Int) != 0 {
			t.Errorf("Test Failed: %v inputted, %v expected, %#v received", test.input, test.expected, amount)
		}
	}
}
