package aeternity

import (
  "crypto/rand"
  "crypto/sha256"
  "encoding/base64"
  "fmt"
  "github.com/btcsuite/btcutil/base58"
  "github.com/satori/go.uuid"
  "golang.org/x/crypto/blake2b"
  "strings"
)

func hashSha256(data []byte) []byte {
  d := sha256.New()
  d.Write(data)
  return d.Sum(nil)
}

// encode encode a byte array into base58 with chacksum and a prefix
func encode(prefix HashPrefix, data []byte) string {
  checksum := hashSha256(hashSha256(data))
  in := append(data, checksum[0:4]...)
  switch objectEncoding[prefix] {
  case Base58c:
    return base58.Encode(in)
  case Base64c:
    return base64.StdEncoding.EncodeToString(in)
  default:
    panic(fmt.Sprint("Encoding not supported"))
  }

}

// decode decode a string encoded with base58 + checksum to a byte array
func decode(in string) (out []byte, err error) {
  // prefix and hash
  var p HashPrefix
  var h string
  // validate input
  if len(in) <= 3 || string(in[2]) == PrefixSeparator {
    err = fmt.Errorf("Invalid object encoding")
    return
  }
  // TODO: check for a valid encoding

  if len(in) >= 3 && string(in[2]) == PrefixSeparator {
    p = HashPrefix(in[0:3])
    h = in[3:]
  }
  switch objectEncoding[in[0:3]] {
  }
  raw := base58.Decode(in)
  if len(raw) < 5 {
    err = fmt.Errorf("Invalid input, %s cannot be decoded", in)
    return
  }
  out = raw[:len(raw)-4]
  if chk := encode(out); in != chk {
    err = fmt.Errorf("Invalid checksum, expected %s got %s", chk, in)
  }
  return
}

// hash calculate the blacke2b 32bit hash of the input byte array
func hash(in []byte) (out []byte, err error) {
  h, err := blake2b.New(32, nil)
  if err != nil {
    return
  }
  h.Write(in)
  out = h.Sum(nil)
  return
}

// namehash calculate the namehash of a string
// TODO: link to the
func namehash(name string) []byte {
  buf := make([]byte, 32)
  for _, s := range strings.Split(name, ".") {
    sh, _ := hash([]byte(s))
    buf, _ = hash(append(buf, sh...))
  }
  return buf
}

// randomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func randomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  // Note that err == nil only if we read len(b) bytes.
  if err != nil {
    return nil, err
  }
  return b, nil
}

// generate an uuid v4 string
func uuidV4() (u string) {
  return fmt.Sprint(uuid.NewV4())
}
