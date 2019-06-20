package utils

import (
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestBigIntNewStr(t *testing.T) {
	// Use the convenience function NewBigIntFromString to set the number
	a, err := NewIntFromString("20000000000000000000") // 2e19
	if err != nil {
		t.Fatal(err)
	}

	// Make a BigInt the old fashioned, raw way.
	ex := new(big.Int)
	ex.SetString("20000000000000000000", 10)

	// They should equal each other.
	if ex.Cmp(a) != 0 {
		t.Fatalf("Expected 20000000000000000000 but got %v", a.String())
	}

}

func TestBigIntLargerOrEqualToZero(t *testing.T) {
	var amount = new(big.Int)
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
		amountc := BigInt(*amount)
		err := amountc.LargerOrEqualToZero()
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %+v received", test.input, test.expected, err)
		}
	}
}

func TestBigIntLargerThanZero(t *testing.T) {
	var amount = new(big.Int)
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
		amountc := BigInt(*amount)
		err := amountc.LargerThanZero()
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %+v received", test.input, test.expected, err)
		}
	}
}

func TestBigInt_UnmarshalJSON(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		b       BigInt
		args    args
		wantErr bool
	}{{
		name:    "Deserialize this",
		args:    args{[]byte("1600000000000000000129127208515966861305")},
		b:       BigInt(*RequireIntFromString("1600000000000000000129127208515966861305")),
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BigInt{}
			if err := b.UnmarshalJSON(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("BigInt.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if b.Cmp(&tt.b) != 0 {
				t.Errorf("The values are not the same: b: %s, want: %s", b.String(), tt.b.String())
			}
		})
	}
}
