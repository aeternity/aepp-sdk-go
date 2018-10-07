//
// from https://github.com/InWeCrypto/cryptox
//

package aeternity

import (
  "bytes"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "time"

  "github.com/pborman/uuid"
  "golang.org/x/crypto/pbkdf2"
  "golang.org/x/crypto/scrypt"
  "golang.org/x/crypto/sha3"
)

var (
  standardScryptN = 1 << 18
  standardScryptP = 1
  lightScryptN    = 1 << 12
  lightScryptP    = 6
  scryptR         = 8
  scryptDklen     = 32
  scryptKDFName   = "scrypt"
  pbkdf2Name      = "pbkdf2"
)

// Errors
var (
  ErrDecrypt = errors.New("could not decrypt key with given passphrase")
)

// Provider keystore serializer provider
type Provider interface {
  Read(data []byte, password string) (*Key, error)
  Write(key *Key, password string, attrs map[string]interface{}) ([]byte, error)
  KdfTypeName() []string
}

// Key keystore handled key object
type Key struct {
  ID         []byte
  Address    string
  PrivateKey []byte
}

func newKey(privateKey []byte, address string) (k *Key) {
  return &Key{
    ID:         uuid.NewRandom(),
    Address:    address,
    PrivateKey: privateKey,
  }
  return
}

// KdfParams .
type KdfParams struct {
  DkLen int    `json:"dklen"` // DK length
  Salt  string `json:"salt"`  // salt string
}

type encryptedKeyJSONV3 struct {
  Address string     `json:"address"`
  Crypto  cryptoJSON `json:"crypto"`
  ID      string     `json:"id"`
  Version int        `json:"version"`
}

type cryptoJSON struct {
  Cipher       string                 `json:"cipher"`
  CipherText   string                 `json:"ciphertext"`
  CipherParams cipherparamsJSON       `json:"cipherparams"`
  KDF          string                 `json:"kdf"`
  KDFParams    map[string]interface{} `json:"kdfparams"`
  MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
  IV string `json:"iv"`
}

var providers = []Provider{
  &Web3KeyStore{},
}

// Decrypt read key from keystore
func Decrypt(data []byte, password string) (*Key, error) {
  provider := &Web3KeyStore{}

  return provider.Read(data, password)
}

// Encrypt encrypt key as keystore data
func Encrypt(key *Key, password string, attrs map[string]interface{}) ([]byte, error) {
  provider := &Web3KeyStore{}

  return provider.Write(key, password, attrs)
}

func selectProvider(keystoreType string) (Provider, bool) {
  for _, provider := range providers {
    for _, support := range provider.KdfTypeName() {
      if support == keystoreType {
        return provider, true
      }
    }
  }

  return nil, false
}

// Web3KeyStore scrypt keystore keystore
type Web3KeyStore struct {
}

// Read .
func (keystore *Web3KeyStore) Read(data []byte, password string) (*Key, error) {

  // Parse the json into a simple map to fetch the key version
  kv := make(map[string]interface{})
  if err := json.Unmarshal(data, &kv); err != nil {
    return nil, err
  }

  if version, ok := kv["version"].(string); ok && version != "3" {
    return nil, fmt.Errorf("cryptox library only support keystore version 3")
  }

  k := new(encryptedKeyJSONV3)

  if err := json.Unmarshal(data, k); err != nil {
    return nil, err
  }

  keyBytes, keyID, err := keystore.decryptKeyV3(k, password)

  if err != nil {
    return nil, err
  }

  return &Key{
    ID:         uuid.UUID(keyID),
    Address:    k.Address,
    PrivateKey: keyBytes,
  }, nil

}

func (keystore *Web3KeyStore) decryptKeyV3(
  keyProtected *encryptedKeyJSONV3,
  password string) (keyBytes []byte, keyID []byte, err error) {

  if keyProtected.Crypto.Cipher != "aes-128-ctr" {
    return nil, nil, fmt.Errorf("Cipher not supported: %v", keyProtected.Crypto.Cipher)
  }

  keyID = uuid.Parse(keyProtected.ID)
  mac, err := hex.DecodeString(keyProtected.Crypto.MAC)

  if err != nil {
    return nil, nil, err
  }

  iv, err := hex.DecodeString(keyProtected.Crypto.CipherParams.IV)
  if err != nil {
    return nil, nil, err
  }

  cipherText, err := hex.DecodeString(keyProtected.Crypto.CipherText)
  if err != nil {
    return nil, nil, err
  }

  derivedKey, err := getKDFKey(keyProtected.Crypto, password)
  if err != nil {
    return nil, nil, err
  }

  hasher := sha3.NewLegacyKeccak256()

  hasher.Write(derivedKey[16:32])
  hasher.Write(cipherText)

  calculatedMAC := hasher.Sum(nil)

  if !bytes.Equal(calculatedMAC, mac) {
    return nil, nil, fmt.Errorf("%s\n%s\n%s",
      ErrDecrypt,
      hex.EncodeToString(calculatedMAC),
      hex.EncodeToString(mac))
  }

  plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)

  if err != nil {
    return nil, nil, err
  }

  return plainText, keyID, err
}

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
  // AES-128 is selected due to size of encryptKey.
  aesBlock, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }
  stream := cipher.NewCTR(aesBlock, iv)
  outText := make([]byte, len(inText))
  stream.XORKeyStream(outText, inText)
  return outText, err
}

func ensureInt(x interface{}) int {
  res, ok := x.(int)
  if !ok {
    res = int(x.(float64))
  }
  return res
}

func getKDFKey(cryptoJSON cryptoJSON, auth string) ([]byte, error) {
  authArray := []byte(auth)
  salt, err := hex.DecodeString(cryptoJSON.KDFParams["salt"].(string))
  if err != nil {
    return nil, err
  }
  dkLen := ensureInt(cryptoJSON.KDFParams["dklen"])

  if cryptoJSON.KDF == scryptKDFName {
    n := ensureInt(cryptoJSON.KDFParams["n"])
    r := ensureInt(cryptoJSON.KDFParams["r"])
    p := ensureInt(cryptoJSON.KDFParams["p"])
    return scrypt.Key(authArray, salt, n, r, p, dkLen)

  } else if cryptoJSON.KDF == "pbkdf2" {
    c := ensureInt(cryptoJSON.KDFParams["c"])
    prf := cryptoJSON.KDFParams["prf"].(string)
    if prf != "hmac-sha256" {
      return nil, fmt.Errorf("Unsupported PBKDF2 PRF: %s", prf)
    }
    key := pbkdf2.Key(authArray, salt, c, dkLen, sha256.New)
    return key, nil
  }

  return nil, fmt.Errorf("Unsupported KDF: %s", cryptoJSON.KDF)
}

// Write .
func (keystore *Web3KeyStore) Write(key *Key, password string, attrs map[string]interface{}) ([]byte, error) {
  authArray := []byte(password)
  salt := GetEntropyCSPRNG(32)

  scryptN := lightScryptN
  scryptP := lightScryptP

  if attrs != nil {
    if scryptN, ok := attrs["ScryptN"]; ok {
      scryptN = scryptN.(int)
    }

    if scryptP, ok := attrs["ScryptP"]; ok {
      scryptP = scryptP.(int)
    }
  }

  derivedKey, err := scrypt.Key(authArray, salt, scryptN, scryptR, scryptP, scryptDklen)

  if err != nil {
    return nil, err
  }

  encryptKey := derivedKey[:16]

  keyBytes := key.PrivateKey

  if len(key.PrivateKey) < 32 {
    keyBytes := make([]byte, 32)

    copy(keyBytes, key.PrivateKey)
  }

  iv := GetEntropyCSPRNG(aes.BlockSize) // 16

  cipherText, err := aesCTRXOR(encryptKey, keyBytes, iv)
  if err != nil {
    return nil, err
  }

  hasher := sha3.NewLegacyKeccak256()

  hasher.Write(derivedKey[16:32])
  hasher.Write(cipherText)

  mac := hasher.Sum(nil)

  scryptParamsJSON := make(map[string]interface{}, 5)
  scryptParamsJSON["n"] = scryptN
  scryptParamsJSON["r"] = scryptR
  scryptParamsJSON["p"] = scryptP
  scryptParamsJSON["dklen"] = scryptDklen
  scryptParamsJSON["salt"] = hex.EncodeToString(salt)

  cipherParamsJSON := cipherparamsJSON{
    IV: hex.EncodeToString(iv),
  }

  cryptoStruct := cryptoJSON{
    Cipher:       "aes-128-ctr",
    CipherText:   hex.EncodeToString(cipherText),
    CipherParams: cipherParamsJSON,
    KDF:          scryptKDFName,
    KDFParams:    scryptParamsJSON,
    MAC:          hex.EncodeToString(mac),
  }
  encryptedKeyJSONV3 := encryptedKeyJSONV3{
    key.Address,
    cryptoStruct,
    uuid.UUID(key.ID).String(),
    3,
  }
  return json.Marshal(encryptedKeyJSONV3)
}

// KdfTypeName get the keystore keystore's kdf alogirthm type
func (keystore *Web3KeyStore) KdfTypeName() []string {
  return []string{
    scryptKDFName,
    pbkdf2Name,
  }
}

// GetEntropyCSPRNG .
func GetEntropyCSPRNG(n int) []byte {
  mainBuff := make([]byte, n)
  _, err := io.ReadFull(rand.Reader, mainBuff)
  if err != nil {
    panic("reading from crypto/rand failed: " + err.Error())
  }
  return mainBuff
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
