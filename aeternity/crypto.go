package aeternity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
)

// KeyPair holds the signing key and the aeternity account address
type KeyPair struct {
	SigningKey ed25519.PrivateKey
	Address    string
}

// Load parse a string into a private key
func Load(priv string) (kp *KeyPair, err error) {
	privb, err := hex.DecodeString(priv)
	if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Invalid private key")
		}
	}()
	kp = &KeyPair{SigningKey: ed25519.PrivateKey(privb)}
	pubb := []byte(fmt.Sprintf("%s", kp.SigningKey.Public()))
	kp.Address = encodeP(PrefixAccount, pubb)
	return
}

func h(data []byte) []byte {
	d := sha256.New()
	d.Write(data)
	return d.Sum(nil)
}

// encode encode a byte array into base58 with chacksum
func encode(in []byte) string {
	c := h(h(in))
	return base58.Encode(append(in, c[0:4]...))
}

// encodeP encode a byte array into base58 with chacksum and a prefix
func encodeP(prefix HashPrefix, data []byte) string {
	return fmt.Sprint(prefix, encode(data))
}

// decode decode a string encoded with base58 + checksum to a byte array
func decode(in string) (out []byte, err error) {
	if len(in) >= 3 && string(in[2]) == "$" {
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

// Sign a message with a private key
func (k *KeyPair) Sign(message []byte) (signature []byte) {
	signature = ed25519.Sign(k.SigningKey, message)
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
