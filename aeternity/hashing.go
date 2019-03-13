package aeternity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aeternity/aepp-sdk-go/rlp"
	"github.com/btcsuite/btcutil/base58"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/blake2b"
)

func hashSha256(data []byte) []byte {
	d := sha256.New()
	d.Write(data)
	return d.Sum(nil)
}

// Encode a byte array into base58/base64 with chacksum and a prefix
func Encode(prefix HashPrefix, data []byte) string {
	checksum := hashSha256(hashSha256(data))
	in := append(data, checksum[0:4]...)
	switch objectEncoding[prefix] {
	case Base58c:
		return fmt.Sprint(prefix, base58.Encode(in))
	case Base64c:
		return fmt.Sprint(prefix, base64.StdEncoding.EncodeToString(in))
	default:
		panic(fmt.Sprint("Encoding not supported"))
	}

}

// Decode a string encoded with base58/base64 + checksum to a byte array
func Decode(in string) (out []byte, err error) {
	// prefix and hash
	var p HashPrefix
	var h string
	var raw []byte

	// Validation
	// 3 (**_) + 5 (Single byte, prefixed with Base58 4 character hash)
	// then split it into p(refix) and h(ash)
	if len(in) <= 8 || string(in[2]) != PrefixSeparator {
		err = fmt.Errorf("Invalid object encoding")
		return
	}
	p = HashPrefix(in[0:3])
	h = in[3:]

	switch objectEncoding[p] {
	case Base58c:
		raw = base58.Decode(h)
	case Base64c:
		raw, _ = base64.StdEncoding.DecodeString(h)
	}
	if len(raw) < 5 {
		err = fmt.Errorf("Invalid input, %s cannot be decoded", in)
		return nil, err
	}
	out = raw[:len(raw)-4]
	if chk := Encode(p, out); in != chk {
		err = fmt.Errorf("Invalid checksum, expected %s got %s", chk, in)
		return nil, err
	}
	return out, nil
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

// naming
func computeCommitmentID(name string) (ch string, salt []byte, err error) {
	salt, err = randomBytes(32)
	if err != nil {
		return
	}
	// TODO: this is done using the api (concatenating )
	nh := append(namehash(name), salt...)
	nh, _ = hash(nh)
	// nh := namehash(name)
	ch = Encode(PrefixCommitment, nh)
	return
}

func buildRLPMessage(tag uint, version uint, fields ...interface{}) (rlpRawMsg []byte, err error) {
	// create a message of the transaction and signature
	data := []interface{}{tag, version}
	data = append(data, fields...)
	// fmt.Printf("TX %#v\n\n", data)
	// encode the message using rlp
	rlpRawMsg, err = rlp.EncodeToBytes(data)
	// fmt.Printf("ENCODED %#v\n\n", data)
	return
}

// buildIDTag assemble an id() object
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md#the-id-type
func buildIDTag(IDTag uint8, encodedHash string) (v []uint8, err error) {
	raw, err := Decode(encodedHash)
	v = []uint8{IDTag}
	for _, x := range raw {
		v = append(v, uint8(x))
	}
	return
}

func decodeRLPMessage(rawBytes []byte) []interface{} {
	res := []interface{}{}
	rlp.DecodeBytes(rawBytes, &res)
	return res
}
