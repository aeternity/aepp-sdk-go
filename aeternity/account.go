package aeternity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

// Account holds the signing/private key and the aeternity account address
type Account struct {
	SigningKey ed25519.PrivateKey
	Address    string
}

func loadAccountFromPrivateKeyRaw(privb []byte) (account *Account, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Invalid private key")
		}
	}()
	account = loadAccountFromPrivateKey(ed25519.PrivateKey(privb))
	return
}

func loadAccountFromPrivateKey(priv ed25519.PrivateKey) (account *Account) {
	account = &Account{
		SigningKey: priv,
		Address:    Encode(PrefixAccountPubkey, []byte(fmt.Sprintf("%s", priv.Public()))),
	}
	return
}

// NewAccount generates a new Account
func NewAccount() (account *Account, err error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	account = loadAccountFromPrivateKey(priv)
	return
}

// AccountFromHexString creates an Account from a hexstring
func AccountFromHexString(hexPrivateKey string) (account *Account, err error) {
	raw, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		return
	}
	return loadAccountFromPrivateKeyRaw(raw)
}

// SigningKeyToHexString returns the SigningKey as an hex string
func (account *Account) SigningKeyToHexString() (signingKeyHex string) {
	signingKeyHex = hex.EncodeToString([]byte(account.SigningKey))
	return
}

// Sign a message with the signing/private key
func (account *Account) Sign(message []byte) (signature []byte) {
	signature = ed25519.Sign(account.SigningKey, message)
	return
}

// Verify a message with the signing/private key
func Verify(address string, message, signature []byte) (valid bool, err error) {
	pub, err := Decode(address)
	if err != nil {
		return
	}
	valid = ed25519.Verify(ed25519.PublicKey(pub), message, signature)
	return
}
