package aeternity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

// Account holds the signing key and the aeternity account address
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
		Address:    encodeP(PrefixAccount, []byte(fmt.Sprintf("%s", priv.Public()))),
	}
	return
}

// NewAccount genereate a new keypair
func NewAccount() (account *Account, err error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	account = loadAccountFromPrivateKey(priv)
	return
}

// AccountFromHexString load an account from hex string
func AccountFromHexString(hexPrivateKey string) (account *Account, err error) {
	raw, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		return
	}
	return loadAccountFromPrivateKeyRaw(raw)
}

// SigningKeyToHexString return the SigningKey as an hex string
func (account *Account) SigningKeyToHexString() (signingKeyHex string) {
	signingKeyHex = hex.EncodeToString([]byte(account.SigningKey))
	return
}

// Sign a message with a private key
func (account *Account) Sign(message []byte) (signature []byte) {
	signature = ed25519.Sign(account.SigningKey, message)
	return
}

// Verify a message with a private key
func Verify(address string, message, signature []byte) (valid bool, err error) {
	pub, err := decode(address)
	if err != nil {
		return
	}
	valid = ed25519.Verify(ed25519.PublicKey(pub), message, signature)
	return
}
