package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	gonanoid "github.com/matoous/go-nanoid"
)

const (
	// DefaultAlphabet default alphabet for string generation
	DefaultAlphabet = "asdfghjklqwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890#!-"
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

func IsInt64(str string) (isInt bool, val int64) {
	if v, err := strconv.Atoi(str); err == nil {
		val = int64(v)
		isInt = true
	}
	return
}

func IsPositiveInt64(str string) (isInt bool, val int64) {
	isInt, val = IsInt64(str)
	isInt = val > 0
	return
}

// GenerateSecret generate a string that can be used as secrete api key
func GenerateSecret() string {
	l := 50
	secret, _ := RandomString(DefaultAlphabet, l)
	return secret
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
