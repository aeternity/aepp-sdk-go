package account

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aeternity/aepp-sdk-go/v5/binary"
	"golang.org/x/crypto/ed25519"
)

// Account holds the signing/private key and the aeternity account address
type Account struct {
	SigningKey ed25519.PrivateKey
	Address    string
}

func loadFromPrivateKeyRaw(privb []byte) (account *Account, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Invalid private key")
		}
	}()
	account = loadFromPrivateKey(ed25519.PrivateKey(privb))
	return
}

func loadFromPrivateKey(priv ed25519.PrivateKey) (account *Account) {
	account = &Account{
		SigningKey: priv,
		Address:    binary.Encode(binary.PrefixAccountPubkey, []byte(fmt.Sprintf("%s", priv.Public()))),
	}
	return
}

// New generates a new Account
func New() (account *Account, err error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	account = loadFromPrivateKey(priv)
	return
}

// FromHexString creates an Account from a hexstring
func FromHexString(hexPrivateKey string) (account *Account, err error) {
	raw, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		return
	}
	return loadFromPrivateKeyRaw(raw)
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
	pub, err := binary.Decode(address)
	if err != nil {
		return
	}
	valid = ed25519.Verify(ed25519.PublicKey(pub), message, signature)
	return
}

// StoreToKeyStoreFile saves an encrypted Account to a JSON file
func StoreToKeyStoreFile(account *Account, password, walletName string) (filePath string, err error) {
	// keystore will be saved in current directory
	basePath, _ := os.Getwd()

	// generate the keystore file
	jks, err := KeystoreSeal(account, password)
	if err != nil {
		return
	}
	// build the wallet path
	filePath = filepath.Join(basePath, keyFileName(account.Address))
	if len(walletName) > 0 {
		filePath = filepath.Join(basePath, walletName)
	}
	// write the file to disk
	err = ioutil.WriteFile(filePath, jks, 0600)
	return
}

// LoadFromKeyStoreFile loads an encrypted Account from a JSON file
func LoadFromKeyStoreFile(keyFile, password string) (account *Account, err error) {
	// find out the real path of the wallet
	filePath, err := GetWalletPath(keyFile)
	if err != nil {
		return
	}
	// load the json file
	jks, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	// decrypt keystore
	account, err = KeystoreOpen(jks, password)
	return
}

// GetWalletPath checks if a file exists at the specified path.
func GetWalletPath(path string) (walletPath string, err error) {
	// if file exists then load the file
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		walletPath = path
		return
	}
	return
}
