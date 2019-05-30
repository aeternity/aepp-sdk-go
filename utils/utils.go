package utils

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/go-openapi/strfmt"
	gonanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// DefaultAlphabet default alphabet for string generation
	DefaultAlphabet = "asdfghjklqwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
)

// IsEqStr tells if two strings a and b are equals after trimming spaces and lowercasing
func IsEqStr(a, b string) bool {
	return strings.ToLower(strings.TrimSpace(a)) == strings.ToLower(strings.TrimSpace(b))
}

// IsEmptyStr tells if a string is empty or not
func IsEmptyStr(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// DefaultIfEmptyStr set a default for a string if it is nulled
func DefaultIfEmptyStr(s *string, defaultS string) {
	if IsEmptyStr(*s) {
		*s = defaultS
	}
}

// DefaultIfEmptyInt set the value of an int to a default if it is nulled (0)
func DefaultIfEmptyInt(v *int, defaultV int) {
	if *v <= 0 {
		*v = defaultV
	}
}

// DefaultIfEmptyInt64 set the value of an int to a default if it is nulled (0)
func DefaultIfEmptyInt64(v *int64, defaultV int64) {
	if *v <= 0 {
		*v = defaultV
	}
}

// DefaultIfEmptyUint8 set the value of an int to a default if it is nulled (0)
func DefaultIfEmptyUint8(v *uint8, defaultV uint8) {
	if *v <= 0 {
		*v = defaultV
	}
}

// DefaultIfEmptyUint32 set the value of an int to a default if it is nulled (0)
func DefaultIfEmptyUint32(v *uint32, defaultV uint32) {
	if *v <= 0 {
		*v = defaultV
	}
}

// DefaultIfEmptyUint64 set the value of an int to a default if it is nulled (0)
func DefaultIfEmptyUint64(v *uint64, defaultV uint64) {
	if *v <= 0 {
		*v = defaultV
	}
}

// RandomString generate a random string of required lenght an with requested alphabet
func RandomString(alphabet string, length int) (s string, err error) {
	if IsEmptyStr(alphabet) {
		err = fmt.Errorf("alphabet must not be empty")
		return
	}
	if length <= 0 {
		err = fmt.Errorf("string length must be longer than 0")
		return
	}
	return gonanoid.Generate(alphabet, length)
}

// RandomStringL generate a string that can be used as secrete api key
func RandomStringL(l int) string {
	secret, _ := RandomString(DefaultAlphabet, l)
	return secret
}

// IsInt64 check if a string is a int64
func IsInt64(str string) (isInt bool, val int64) {
	if v, err := strconv.Atoi(str); err == nil {
		val = int64(v)
		isInt = true
	}
	return
}

// IsPositiveInt64 check if a string is a positive integer
func IsPositiveInt64(str string) (isInt bool, val int64) {
	isInt, val = IsInt64(str)
	isInt = val > 0
	return
}

// AskYes prompt a yes/no question to the prompt
func AskYes(question string, defaultYes bool) (isYes bool) {
	fmt.Print(question)
	if defaultYes {
		fmt.Print(" [yes]: ")
	} else {
		fmt.Print(" [no]: ")
	}
	reader := bufio.NewReader(os.Stdin)
	reply, _ := reader.ReadString('\n')
	DefaultIfEmptyStr(&reply, "yes")
	if IsEqStr(reply, "yes") {
		return true
	}
	return
}

// AskPassword ask a password
func AskPassword(question string) (password string, err error) {
	fmt.Println(question)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	password = string(bytePassword)
	return
}

// BigInt is composed of a big.Int, but includes a Validate() method for swagger and other convenience functions.
// Once created, it can be used just like a big.Int.
type BigInt struct {
	*big.Int
}

// Validate ensures that the BigInt's value is >= 0.
// The actual check does not need 'formats' from swagger, which is why Validate() wraps that function.
func (b *BigInt) Validate(formats strfmt.Registry) error {
	v := b.LargerOrEqualToZero()
	if !v {
		return fmt.Errorf("%v was not >=0", b.Int.String())
	}
	return nil
}

// LargerThanZero checks that the number is >=0
func (b *BigInt) LargerThanZero() bool {
	zero := NewBigInt()

	if b.Cmp(zero.Int) != 1 {
		return false
	}
	return true
}

// LargerOrEqualToZero checks that the number is >=0
func (b *BigInt) LargerOrEqualToZero() bool {
	zero := NewBigInt()

	if b.Cmp(zero.Int) == -1 {
		return false
	}
	return true
}

// UnmarshalJSON ensures that BigInt.Int is always initialized, even if the JSON value is nil.
func (b *BigInt) UnmarshalJSON(text []byte) error {
	if b.Int == nil {
		b.Int = &big.Int{}
	}

	return b.Int.UnmarshalJSON(text)
}

// NewBigInt returns a new BigInt with its Int struct field initialized
func NewBigInt() (i *BigInt) {
	return &BigInt{new(big.Int)}
}

// NewBigIntFromString returns a new BigInt from a string representation
func NewBigIntFromString(number string) (i *BigInt, err error) {
	i = &BigInt{new(big.Int)}
	_, success := i.SetString(number, 10)
	if success == false {
		return nil, errors.New("Could not parse string as a number")
	}
	return i, nil
}

// RequireBigIntFromString returns a new BigInt from a string representation or panics if NewBigIntFromString would have returned an error.
func RequireBigIntFromString(number string) *BigInt {
	i, err := NewBigIntFromString(number)
	if err != nil {
		panic(err)
	}
	return i
}

// NewBigIntFromUint64 returns a new BigInt from a uint64 representation
func NewBigIntFromUint64(number uint64) (i *BigInt) {
	i = &BigInt{new(big.Int)}
	i.SetUint64(number)
	return i
}
