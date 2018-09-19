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
	PrefixKeyBlockHash = HashPrefix("kh_")
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

// verion used in the rlp message
const (
	rlpMessageVersion uint = 1
)

// Tag constant for ids (type uint8)
// see https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
// <<Tag:1/unsigned-integer-unit:8, Hash:32/binary-unit:8>>
const (
	IDTagAccount    uint8 = 1
	IDTagName       uint8 = 2
	IDTagCommitment uint8 = 3
	IDTagOracle     uint8 = 4
	IDTagContract   uint8 = 5
	IDTagChannel    uint8 = 6
)

// Object tags
// see https://github.com/aeternity/protocol/blob/master/serializations.md#binary-serialization
const (
	ObjectTagAccount                             uint = 10
	ObjectTagSignedTransaction                   uint = 11
	ObjectTagSpendTransaction                    uint = 12
	ObjectTagOracle                              uint = 20
	ObjectTagOracleQuery                         uint = 21
	ObjectTagOracleRegisterTransaction           uint = 22
	ObjectTagOracleQueryTransaction              uint = 23
	ObjectTagOracleResponseTransaction           uint = 24
	ObjectTagOracleExtendTransaction             uint = 25
	ObjectTagNameServiceName                     uint = 30
	ObjectTagNameServiceCommitment               uint = 31
	ObjectTagNameServiceClaimTransaction         uint = 32
	ObjectTagNameServicePreclaimTransaction      uint = 33
	ObjectTagNameServiceUpdateTransaction        uint = 34
	ObjectTagNameServiceRevokeTransaction        uint = 35
	ObjectTagNameServiceTransferTransaction      uint = 36
	ObjectTagContract                            uint = 40
	ObjectTagContractCall                        uint = 41
	ObjectTagContractCreateTransaction           uint = 42
	ObjectTagContractCallTransaction             uint = 43
	ObjectTagChannelCreateTransaction            uint = 50
	ObjectTagChannelDepositTransaction           uint = 51
	ObjectTagChannelWithdrawTransaction          uint = 52
	ObjectTagChannelForceProgressTransaction     uint = 521
	ObjectTagChannelCloseMutualTransaction       uint = 53
	ObjectTagChannelCloseSoloTransaction         uint = 54
	ObjectTagChannelSlashTransaction             uint = 55
	ObjectTagChannelSettleTransaction            uint = 56
	ObjectTagChannelOffChainTransaction          uint = 57
	ObjectTagChannelOffChainUpdateTransfer       uint = 570
	ObjectTagChannelOffChainUpdateDeposit        uint = 571
	ObjectTagChannelOffChainUpdateWithdrawal     uint = 572
	ObjectTagChannelOffChainUpdateCreateContract uint = 573
	ObjectTagChannelOffChainUpdateCallContract   uint = 574
	ObjectTagChannel                             uint = 58
	ObjectTagChannelSnapshotTransaction          uint = 59
	ObjectTagPoi                                 uint = 60
	ObjectTagMicroBody                           uint = 101
	ObjectTagLightMicroBlock                     uint = 102
)
