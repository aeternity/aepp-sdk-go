package aeternity

// HashPrefix a prefix for an aeternity object hash
type HashPrefix string

const (
	// PrefixAccount prefix of an account address
	PrefixAccount = HashPrefix("ak$")
	// PrefixTx prefix of a transaction
	PrefixTx = HashPrefix("tx$")
	// PrefixTxHash prefix of a transaction hash
	PrefixTxHash = HashPrefix("th$")
	// PrefixBlockHash prefix of a block hash
	PrefixBlockHash = HashPrefix("bh$")
	// PrefixContract prefix of a contract address
	PrefixContract = HashPrefix("ct$")
	// PrefixNameHash prefix of an a name hash
	PrefixNameHash = HashPrefix("nm$")
	// PrefixSignature prefix of an a signature
	PrefixSignature = HashPrefix("sg$")
	// PrefixBlockTxHash prefix of a block transaction hash
	PrefixBlockTxHash = HashPrefix("bx$")
	// PrefixBlockStateHash prefix of a block state hash
	PrefixBlockStateHash = HashPrefix("bs$")
	// PrefixChannel prefix of a channel
	PrefixChannel = HashPrefix("ch$")
	// PrefixNameCommitment prefix of a name commmitemtn hash
	PrefixNameCommitment = HashPrefix("cm$")
)
