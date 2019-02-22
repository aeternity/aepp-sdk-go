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

// BigInt is used by swagger as a big.Int type with Validate() function
type BigInt struct {
	*big.Int
}

// Validate is an exported function that swagger uses.
// However, the implementation does not need 'formats', so it is broken
// out into validate().
func (b *BigInt) Validate(formats strfmt.Registry) error {
	return b.validate()
}

// validate checks that the number is >=0
func (b *BigInt) validate() error {
	var zero big.Int
	var convertedCustomBigInt = b

	if convertedCustomBigInt.Cmp(&zero) != 1 {
		return errors.New("swagger deserialization: Balance Validation failed")
	}
	return nil
}
