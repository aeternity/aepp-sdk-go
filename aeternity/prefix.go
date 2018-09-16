package aeternity

import (
	"fmt"
)

// HashPrefix a prefix for an aeternity object hash
type HashPrefix string

const (
	// PrefixSeparator the separator for the prefixes
	PrefixSeparator = "_"
	// PrefixAccount prefix of an account address
	PrefixAccount = HashPrefix("ak_")
	// PrefixTx prefix of a transaction
	PrefixTx = HashPrefix("tx_")
	// PrefixTxHash prefix of a transaction hash
	PrefixTxHash = HashPrefix("th_")
	// PrefixKeyBlockHash prefix of a block hash
	PrefixKeyBlockHash = HashPrefix("bh_")
	// PrefixMicroBlockHash prefix of a block hash TODO: what is the real prefix
	PrefixMicroBlockHash = HashPrefix("mh_")
	// PrefixContract prefix of a contract address
	PrefixContract = HashPrefix("ct_")
	// PrefixNameHash prefix of an a name hash
	PrefixNameHash = HashPrefix("nm_")
	// PrefixSignature prefix of an a signature
	PrefixSignature = HashPrefix("sg_")
	// PrefixBlockTxHash prefix of a block transaction hash
	PrefixBlockTxHash = HashPrefix("bx_")
	// PrefixBlockStateHash prefix of a block state hash
	PrefixBlockStateHash = HashPrefix("bs_")
	// PrefixChannel prefix of a channel
	PrefixChannel = HashPrefix("ch_")
	// PrefixNameCommitment prefix of a name commmitment hash
	PrefixNameCommitment = HashPrefix("cm_")
	// PrefixOracle prefix of an oracle
	PrefixOracle = HashPrefix("ok_")
)

// GetHashPrefix get the prefix of an hash, panics if the hash is too short
func GetHashPrefix(hash string) (p HashPrefix) {
	if len(hash) <= 3 {
		panic(fmt.Sprintln("Invalid hash", hash))
	}
	p = HashPrefix(hash[0:3])
	return
}
)
