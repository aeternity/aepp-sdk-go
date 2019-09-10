package aeternity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/aeternity/aepp-sdk-go/utils"
	"github.com/btcsuite/btcutil/base58"
	rlp "github.com/randomshinichi/rlpae"
	uuid "github.com/satori/go.uuid"
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

func leftPadByteSlice(length int, data []byte) []byte {
	dataLen := len(data)
	t := make([]byte, length-dataLen)
	paddedSlice := append(t, data...)
	return paddedSlice
}

func buildOracleQueryID(sender string, senderNonce uint64, recipient string) (id string, err error) {
	queryIDBin := []byte{}
	senderBin, err := Decode(sender)
	if err != nil {
		return
	}
	queryIDBin = append(queryIDBin, senderBin...)

	senderNonceBytes := utils.NewIntFromUint64(senderNonce).Bytes()
	senderNonceBytesPadded := leftPadByteSlice(32, senderNonceBytes)
	queryIDBin = append(queryIDBin, senderNonceBytesPadded...)

	recipientBin, err := Decode(recipient)
	if err != nil {
		return
	}
	queryIDBin = append(queryIDBin, recipientBin...)

	hashedQueryID, err := Blake2bHash(queryIDBin)
	if err != nil {
		return
	}
	id = Encode(PrefixOracleQueryID, hashedQueryID)
	return
}

func buildContractID(sender string, senderNonce uint64) (ctID string, err error) {
	senderBin, err := Decode(sender)
	if err != nil {
		return ctID, err
	}

	l := big.Int{}
	l.SetUint64(senderNonce)

	ctIDUnhashed := append(senderBin, l.Bytes()...)
	ctIDHashed, err := Blake2bHash(ctIDUnhashed)
	if err != nil {
		return ctID, err
	}

	ctID = Encode(PrefixContractPubkey, ctIDHashed)
	return ctID, err
}

// Namehash calculate the Namehash of a string. Names within aeternity are
// generally referred to only by their namehashes.
//
// The implementation is the same as ENS EIP-137
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-137.md#namehash-algorithm
// but using Blake2b.
func Namehash(name string) []byte {
	buf := make([]byte, 32)
	for _, s := range strings.Split(name, ".") {
		sh, _ := Blake2bHash([]byte(s))
		buf, _ = Blake2bHash(append(buf, sh...))
	}
	return buf
}

// randomBytes returns securely generated random bytes. It will return an error
// if the system's secure random number generator fails to function correctly,
// in which case the caller should not continue.
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

// generateCommitmentID gives a commitment ID 'cm_...' given a particular AENS
// name. It is split into the deterministic part computeCommitmentID(), which
// can be tested, and the part incorporating random salt generateCommitmentID()
//
// since the salt is a uint256, which Erlang handles well, but Go has nothing
// similar to it, it is imperative that the salt be kept as a bytearray unless
// you really have to convert it into an integer. Which you usually don't,
// because it's a salt.
func generateCommitmentID(name string) (ch string, salt *big.Int, err error) {
	saltBytes, err := randomBytes(32)
	if err != nil {
		return
	}

	ch, err = computeCommitmentID(name, saltBytes)

	salt = new(big.Int)
	salt.SetBytes(saltBytes)

	return ch, salt, err
}

func computeCommitmentID(name string, salt []byte) (ch string, err error) {
	nh := append(Namehash(name), salt...)
	nh, _ = Blake2bHash(nh)
	ch = Encode(PrefixCommitment, nh)
	return
}

func buildRLPMessage(tag uint, version uint, fields ...interface{}) (rlpRawMsg []byte, err error) {
	// create a message of the transaction and signature
	data := []interface{}{tag, version}
	data = append(data, fields...)
	// fmt.Printf("TX %+v\n\n", data)
	// encode the message using rlp
	rlpRawMsg, err = rlp.EncodeToBytes(data)
	// fmt.Printf("ENCODED %+v\n\n", data)
	return
}

// buildIDTag assemble an id() object see
// https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
func buildIDTag(IDTag uint8, encodedHash string) (v []uint8, err error) {
	raw, err := Decode(encodedHash)
	v = []uint8{IDTag}
	for _, x := range raw {
		v = append(v, uint8(x))
	}
	return
}

// readIDTag disassemble an id() object see
// https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
func readIDTag(v []uint8) (IDTag uint8, encodedHash string, err error) {
	IDTag = v[0]
	hash := []byte{}
	for _, x := range v[1:] {
		hash = append(hash, byte(x))
	}

	var prefix HashPrefix
	switch IDTag {
	case IDTagAccount:
		prefix = PrefixAccountPubkey
	case IDTagName:
		prefix = PrefixName
	case IDTagCommitment:
		prefix = PrefixCommitment
	case IDTagOracle:
		prefix = PrefixOraclePubkey
	case IDTagContract:
		prefix = PrefixContractPubkey
	case IDTagChannel:
		prefix = PrefixChannel
	default:
		return 0, "", fmt.Errorf("readIDTag() does not recognize this IDTag (first byte in input array): %v", IDTag)
	}

	encodedHash = Encode(prefix, hash)
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
