package cmd

import "github.com/aeternity/aepp-sdk-go/v9/binary"

// IsAddress does some minor checks to ensure that the string is an ak_ address
func IsAddress(a string) bool {
	if len(a) > 0 && a[:3] == string(binary.PrefixAccountPubkey) {
		return true
	}
	return false
}

// IsBytecode does some minor checks to ensure that the string is a cb_ bytecode
func IsBytecode(a string) bool {
	if len(a) > 0 && a[:3] == string(binary.PrefixContractByteArray) {
		return true
	}
	return false
}

// IsTransaction does some minor checks to ensure that the string is a tx_ transaction
func IsTransaction(a string) bool {
	if len(a) > 0 && a[:3] == string(binary.PrefixTransaction) {
		return true
	}
	return false
}
