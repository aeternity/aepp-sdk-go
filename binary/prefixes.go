package binary

import "fmt"

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
