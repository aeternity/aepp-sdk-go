package binary

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	rlp "github.com/aeternity/rlp-go"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
)

func hashSha256(data []byte) []byte {
	d := sha256.New()
	d.Write(data)
	return d.Sum(nil)
}

// Encode a byte array into base58/base64 with checksum and a prefix.
//
// in aeternity, bytearrays are always base-encoded with a prefix that indicates
// what the bytearray is. For example, accounts "ak_...." are a plain bytearray
// that is base58 encoded and prefixed with "ak_" to indicate that it is an
// account.
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
//
// in aeternity, bytearrays are always base-encoded with a prefix that indicates
// what the bytearray is. For example, accounts "ak_...." are a plain bytearray
// that is base58 encoded and prefixed with "ak_" to indicate that it is an
// account.
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

// Blake2bHash calculates the blake2b 32bit hash of the input byte array
func Blake2bHash(in []byte) (out []byte, err error) {
	h, err := blake2b.New(32, nil)
	if err != nil {
		return
	}
	h.Write(in)
	out = h.Sum(nil)
	return
}

// DecodeRLPMessage takes an RLP serialized bytearray and parses the RLP to
// return the deserialized, structured data as bytearrays ([]interfaces that
// should be later coerced into specific types). Only meant for debugging
// purposes - to parse serialized RLP into a useful type, see DeseralizeTx.
func DecodeRLPMessage(rawBytes []byte) []interface{} {
	res := []interface{}{}
	rlp.DecodeBytes(rawBytes, &res)
	return res
}
