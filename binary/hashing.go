package binary

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	rlp "github.com/randomshinichi/rlpae"
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

// HashPrefix describes a prefix that is attached to every base-encoded
// bytearray used in aeternity to describe its function.
//
// For example, the "ak_" HashPrefix describes an account address and "ct_"
// HashPrefix describes a contract address.
type HashPrefix string

// GetHashPrefix returns a HashPrefix of a string. It panics if the hash
// contains only the prefix (length 3).
func GetHashPrefix(hash string) (p HashPrefix) {
	if len(hash) <= 3 {
		panic(fmt.Sprintln("Invalid hash", hash))
	}
	p = HashPrefix(hash[0:3])
	return
}

// ObjectEncoding is an enum string that describes whether a bytearray is base58
// or base64 encoded
type ObjectEncoding string

// Base58/Base64 encoding definitions
const (
	Base58c = ObjectEncoding("b58c")
	Base64c = ObjectEncoding("b64c")
)

// Prefixes
const (
	// Prefix separator
	PrefixSeparator = "_"

	// Base58 encoded bytearrays
	PrefixAccountPubkey         = HashPrefix("ak_")
	PrefixBlockProofOfFraudHash = HashPrefix("bf_")
	PrefixBlockStateHash        = HashPrefix("bs_")
	PrefixBlockTransactionHash  = HashPrefix("bx_")
	PrefixChannel               = HashPrefix("ch_")
	PrefixCommitment            = HashPrefix("cm_")
	PrefixContractPubkey        = HashPrefix("ct_")
	PrefixKeyBlockHash          = HashPrefix("kh_")
	PrefixMicroBlockHash        = HashPrefix("mh_")
	PrefixName                  = HashPrefix("nm_")
	PrefixOraclePubkey          = HashPrefix("ok_")
	PrefixOracleQueryID         = HashPrefix("oq_")
	PrefixPeerPubkey            = HashPrefix("pp_")
	PrefixSignature             = HashPrefix("sg_")
	PrefixTransactionHash       = HashPrefix("th_")

	// Base64 encoded bytearrays
	PrefixByteArray         = HashPrefix("ba_")
	PrefixContractByteArray = HashPrefix("cb_")
	PrefixOracleResponse    = HashPrefix("or_")
	PrefixOracleQuery       = HashPrefix("ov_")
	PrefixProofOfInclusion  = HashPrefix("pi_")
	PrefixStateTrees        = HashPrefix("ss_")
	PrefixState             = HashPrefix("st_")
	PrefixTransaction       = HashPrefix("tx_")
)

// objectEncoding maps a HashPrefix like "ak_" to its base encoding scheme
// (base58).
var objectEncoding = map[HashPrefix]ObjectEncoding{
	PrefixByteArray:             Base64c,
	PrefixContractByteArray:     Base64c,
	PrefixOracleResponse:        Base64c,
	PrefixOracleQuery:           Base64c,
	PrefixProofOfInclusion:      Base64c,
	PrefixStateTrees:            Base64c,
	PrefixState:                 Base64c,
	PrefixTransaction:           Base64c,
	PrefixAccountPubkey:         Base58c,
	PrefixBlockProofOfFraudHash: Base58c,
	PrefixBlockStateHash:        Base58c,
	PrefixBlockTransactionHash:  Base58c,
	PrefixChannel:               Base58c,
	PrefixCommitment:            Base58c,
	PrefixContractPubkey:        Base58c,
	PrefixKeyBlockHash:          Base58c,
	PrefixMicroBlockHash:        Base58c,
	PrefixName:                  Base58c,
	PrefixOraclePubkey:          Base58c,
	PrefixOracleQueryID:         Base58c,
	PrefixPeerPubkey:            Base58c,
	PrefixSignature:             Base58c,
	PrefixTransactionHash:       Base58c,
}
