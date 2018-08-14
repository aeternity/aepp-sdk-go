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

func encode(in []byte) string {
	h := func(data []byte) []byte {
		d := sha256.New()
		d.Write(data)
		return d.Sum(nil)
	}
	c := h(h(in))
	return base58.Encode(append(in, c[0:4]...))
}

func encodeP(prefix string, data []byte) string {
	return fmt.Sprint(prefix, encode(data))
}

func decode(in string) (out []byte, err error) {
	strings.HasPrefix("aa", "")
	if string(in[2]) == "$" {
		in = in[3:]
	}
	out = base58.Decode(in)
	c := encode(out[len(out)-4:])
	if in != c {
		err = fmt.Errorf("Invalid checksum")
	}
	return
}
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
