package aeternity

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	cryptoSecretType   = "ed25519"
	cryptoSymmetricAlg = "xsalsa20-poly1305"
	kdf                = "argon2id"
	kdfKeySize         = 32
	formatVersion      = 1
)

// KeystoreJSON keystore format
type keystoreJSON struct {
	PublicKey string `json:"public_key"`
	Crypto    crypto `json:"crypto"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Version   int    `json:"version"`
}

type crypto struct {
	SecretType   string       `json:"secret_type"`
	SymmetricAlg string       `json:"symmetric_alg"`
	Ciphertext   string       `json:"ciphertext"`
	CipherParams cipherParams `json:"cipher_params"`
	Kdf          string       `json:"kdf"`
	KdfParams    kdfParams    `json:"kdf_params"`
}

type cipherParams struct {
	Nonce string `json:"nonce"`
}

type kdfParams struct {
	Memlimit    uint32 `json:"memlimit_kib"`
	Opslimit    uint32 `json:"opslimit"` // time
	Salt        string `json:"salt"`
	Parallelism uint8  `json:"parallelism"`
}

// KeystoreOpen open and decrypt a keystore
func KeystoreOpen(data []byte, password string) (account *Account, err error) {
	k := keystoreJSON{}
	err = json.Unmarshal(data, &k)
	if err != nil {
		return
	}
	// build the key from the password
	salt, err := hex.DecodeString(k.Crypto.KdfParams.Salt)
	argonKey := argon2.IDKey([]byte(password), salt,
		k.Crypto.KdfParams.Opslimit,
		k.Crypto.KdfParams.Memlimit,
		k.Crypto.KdfParams.Parallelism,
		kdfKeySize)
	var key [kdfKeySize]byte
	copy(key[:], argonKey)
	// retrieve the nonce
	v, err := hex.DecodeString(k.Crypto.CipherParams.Nonce)
	var decryptNonce [24]byte
	copy(decryptNonce[:], v)
	// now retrieve the cypertext
	v, err = hex.DecodeString(k.Crypto.Ciphertext)
	//
	decrypted, ok := secretbox.Open(nil, v, &decryptNonce, &key)
	if !ok {
		err = fmt.Errorf("Cannot decrypt secret")
		return
	}
	// now load the account
	account, err = loadAccountFromPrivateKeyRaw(decrypted)
	return
}

// KeystoreSeal create an encrypted json keystore with the private key of the account
func KeystoreSeal(account *Account, password string) (j []byte, e error) {
	// normalize pwd
	salt, err := randomBytes(16)
	if err != nil {
		return
	}
	argonKey := argon2.IDKey([]byte(password), salt,
		Config.Tuning.CryptoKdfOpslimit,
		Config.Tuning.CryptoKdfMemlimit,
		Config.Tuning.CryptoKdfThreads,
		kdfKeySize)

	var key [kdfKeySize]byte
	copy(key[:], argonKey)
	// generate nonce
	nonce, err := randomBytes(24)
	if err != nil {
		return
	}
	var n24 [24]byte
	copy(n24[:], nonce)
	//
	privateKeyRaw := []byte(account.SigningKey)
	encrypted := secretbox.Seal(nil, privateKeyRaw, &n24, &key)
	// serialize
	k := keystoreJSON{
		ID:        uuidV4(),
		Name:      keyFileName(account.Address),
		Version:   formatVersion,
		PublicKey: account.Address,
		Crypto: crypto{
			SecretType:   cryptoSecretType,
			SymmetricAlg: cryptoSymmetricAlg,
			CipherParams: cipherParams{Nonce: hex.EncodeToString(nonce)},
			Ciphertext:   hex.EncodeToString(encrypted),
			Kdf:          kdf,
			KdfParams: kdfParams{
				Memlimit:    Config.Tuning.CryptoKdfMemlimit,
				Opslimit:    Config.Tuning.CryptoKdfOpslimit,
				Salt:        hex.EncodeToString(salt),
				Parallelism: Config.Tuning.CryptoKdfThreads,
			},
		},
	}
	// print json
	j, err = json.Marshal(k)
	return
}

// keyFileName implements the naming convention for keyfiles:
// UTC--<created_at UTC ISO8601>-<address hex>
func keyFileName(keyAddr string) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), keyAddr)
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
