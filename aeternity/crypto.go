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

// encode encode a byte array into base58 with chacksum
func encode(in []byte) string {
  c := hashSha256(hashSha256(in))
  return base58.Encode(append(in, c[0:4]...))
}

func encode64(in []byte) string {
	c := hashSha256(hashSha256(in))
	return base64.StdEncoding.EncodeToString(append(in, c[0:4]...))
}

// encodeP encode a byte array into base58 with chacksum and a prefix
func encodeP(prefix HashPrefix, data []byte) string {
  return fmt.Sprint(prefix, encode(data))
}

// encodeP64 encode a byte array into base64 with chacksum and a prefix
func encodeP64(prefix HashPrefix, data []byte) string {
  return fmt.Sprint(prefix, encode64(data))
}

// decode decode a string encoded with base58 + checksum to a byte array
func decode(in string) (out []byte, err error) {
  if len(in) >= 3 && string(in[2]) == PrefixSeparator {
    in = in[3:]
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
