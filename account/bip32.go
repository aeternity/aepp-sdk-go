package account

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// ErrHardenedChildPublicKey is returned when trying to create a hardened child
// from a public key.
var ErrHardenedChildPublicKey = errors.New("Can't create hardened child from public key")

// ErrHardenedOnly is returned when a node in the path is not hardened.
var ErrHardenedOnly = errors.New("ed25519 only works with hardened children")

// Key represents a bip32 extended key, used
// for deriving subaccounts for HD wallets.
type Key struct {
	Key         []byte // 33 bytes
	ChildNumber uint32 // 4 bytes
	ChainCode   []byte // 32 bytes NEEDED FOR CHILDKEY DERIVATION
	Depth       uint32 // 1 bytes
	IsPrivate   bool   // unserialized
}

// NewMasterKey creates a new master extended key from a seed
func NewMasterKey(seed []byte) (*Key, error) {
	// Generate key and chaincode
	hmac := hmac.New(sha512.New, []byte("ed25519 seed"))
	_, err := hmac.Write(seed)
	if err != nil {
		return nil, err
	}
	intermediary := hmac.Sum(nil)

	// Split it into our key and chain code
	keyBytes := intermediary[:32]
	chainCode := intermediary[32:]

	// Create the key struct
	key := &Key{
		ChainCode:   chainCode,
		Key:         keyBytes,
		Depth:       0,
		ChildNumber: 0,
		IsPrivate:   true,
	}

	return key, nil
}

// NewChildKey derives a child key from a given parent as outlined by bip32
func (key *Key) NewChildKey(childIdx uint32) (*Key, error) {
	if childIdx < FirstHardenedChild {
		return nil, ErrHardenedOnly
	}

	intermediary, err := key.getIntermediary(childIdx)
	if err != nil {
		return nil, err
	}
	// Create child Key with data common to all both scenarios
	childKey := &Key{
		ChildNumber: childIdx,
		Key:         intermediary[:32],
		ChainCode:   intermediary[32:],
		Depth:       key.Depth + 1,
		IsPrivate:   key.IsPrivate,
	}

	return childKey, nil
}

func hexify(i []byte) string {
	return hex.EncodeToString(i)
}

func (key *Key) getIntermediary(childIdx uint32) ([]byte, error) {
	// Get intermediary to create key and chaincode from
	// Hardened children are based on the private key
	// NonHardened children are based on the public key
	childIndexBytes := uint32Bytes(childIdx)

	var data []byte
	if childIdx < FirstHardenedChild {
		return nil, ErrHardenedOnly
	}
	data = append([]byte{0x0}, key.Key...)
	data = append(data, childIndexBytes...)

	hmac := hmac.New(sha512.New, key.ChainCode)
	_, err := hmac.Write(data)
	if err != nil {
		return nil, err
	}
	return hmac.Sum(nil), nil
}

func uint32Bytes(i uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, i)
	return bytes
}
