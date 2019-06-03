package utils

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/go-openapi/strfmt"
)

// BigInt is an alias for math/big.Int, for use with Swagger generated code.
// Even though it has some corresponding methods, convert it as soon as possible into big.Int.
type BigInt big.Int

// String casts BigInt into big.Int and uses its String method.
func (b *BigInt) String() string {
	bc := big.Int(*b)
	return bc.String()
}

// Validate ensures that the BigInt's value is >= 0.
// The actual check does not need 'formats' from swagger, which is why Validate() wraps that function.
func (b *BigInt) Validate(formats strfmt.Registry) error {
	v := b.LargerOrEqualToZero()
	if !v {
		return fmt.Errorf("%v was not >=0", b.String())
	}
	return nil
}

// LargerThanZero returns true if it is >0
func (b *BigInt) LargerThanZero() bool {
	zero := new(BigInt)

	if b.Cmp(zero) != 1 {
		return false
	}
	return true
}

// LargerOrEqualToZero checks that the number is >=0
func (b *BigInt) LargerOrEqualToZero() bool {
	zero := new(BigInt)

	if b.Cmp(zero) == -1 {
		return false
	}
	return true
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *BigInt) UnmarshalJSON(text []byte) error {
	bc := new(big.Int)
	err := bc.UnmarshalJSON(text)
	if err != nil {
		return err
	}
	b.Set(bc)
	return nil
}

// Set makes a BigInt equal to a given big.Int.
func (b *BigInt) Set(i *big.Int) *BigInt {
	iB := BigInt(*i)
	*b = iB
	return b
}

// Cmp compares two BigInts just like big.Int
func (b *BigInt) Cmp(i *BigInt) int {
	b2 := big.Int(*b)
	i2 := big.Int(*i)
	return b2.Cmp(&i2)
}

// NewBigIntFromString returns a new math/big.Int from a string representation
func NewBigIntFromString(number string) (i *big.Int, err error) {
	i = new(big.Int)
	_, success := i.SetString(number, 10)
	if success == false {
		return nil, errors.New("Could not parse string as a number")
	}
	return i, nil
}

// RequireBigIntFromString returns a new  big.Int from a string representation or panics if NewBigIntFromString would have returned an error.
func RequireBigIntFromString(number string) *big.Int {
	i, err := NewBigIntFromString(number)
	if err != nil {
		panic(err)
	}
	return i
}

// NewBigIntFromUint64 returns a new big.Int from a uint64 representation
func NewBigIntFromUint64(number uint64) (i *big.Int) {
	i = new(big.Int)
	i.SetUint64(number)
	return i
}
