package account

import (
	"bytes"
	"fmt"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/nacl/sign"
)

// ParseMnemonic uses BIP39 to parse a mnemonic. BIP39 uses a password along
// with the mnemonic to arrive at the final seed bytearray value, but aeternity
// only ever uses a blank string for the password, so the point of ParseMnemonic
// is to reduce ambiguity/confusion.
func ParseMnemonic(mnemonic string) (masterSeed []byte, err error) {
	masterSeed = bip39.NewSeed(mnemonic, "")
	return
}

// BIP32KeyToAeKey translates a BIP32 Key into an aeternity Account.
func BIP32KeyToAeKey(key *Key) (acc *Account, err error) {
	keyReader := bytes.NewReader(key.Key)
	_, privKey, err := sign.GenerateKey(keyReader)
	if err != nil {
		return
	}
	return loadFromPrivateKeyRaw(privKey[:])
}

// DerivePathFromSeed derives a BIP32 Key given a seed (usually derived from a
// mnemonic) and an account index and address index, as standardized by AEX-10.
func DerivePathFromSeed(masterSeed []byte, accountIndex, addressIndex uint32) (key *Key, err error) {
	path := buildPath(accountIndex, addressIndex)
	mK, err := NewMasterKey(masterSeed)
	if err != nil {
		return
	}

	return derivePath(mK, path)
}

// buildPath ensures that only aeternity BIP32 addresses can be calculated. Due
// to ed25519, only hardened path nodes are supported. Hardened nodes are
// denoted with apostrophes ', e.g. "m/44'/457'/0'/0'/0'".
func buildPath(accountIndex, addressIndex uint32) string {
	return fmt.Sprintf("m/44'/457'/%d'/0'/%d'", accountIndex, addressIndex)
}

func derivePath(key *Key, path string) (childKey *Key, err error) {
	parsedPath, err := ParsePath(path)
	if err != nil {
		return
	}

	childKey = key
	for _, p := range parsedPath.Elements {
		if p.Master {
			continue
		}
		childKey, err = childKey.NewChildKey(p.ChildNumber)
		if err != nil {
			return
		}
	}
	return
}
