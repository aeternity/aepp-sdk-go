package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

// IsEqStr tells if two strings a and b are equals after trimming spaces and lowercasing
func IsEqStr(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
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

// AskYes prompt a yes/no question to the prompt
func AskYes(question string, defaultYes bool) (isYes bool) {
	defaultStrVal := "yes"
	if !defaultYes {
		defaultStrVal = "no"
	}
	fmt.Print(question, " [", defaultStrVal, "]: ")
	reader := bufio.NewReader(os.Stdin)
	reply, _ := reader.ReadString('\n')
	DefaultIfEmptyStr(&reply, defaultStrVal)
	if IsEqStr(reply, "yes") {
		return true
	}
	return
}

// AskPassword ask a password
func AskPassword(question string) (password string, err error) {
	fmt.Println(question)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	password = string(bytePassword)
	return
}
