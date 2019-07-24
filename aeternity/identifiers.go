package aeternity

import (
	"fmt"
)

// HashPrefix a prefix for an aeternity object hash
type HashPrefix string

// ObjectEncoding the encoding of an object
type ObjectEncoding string

// Encoding strategies
const (
	Base58c = ObjectEncoding("b58c")
	Base64c = ObjectEncoding("b64c")
)

// Prefixes
const (
	// Prefix separator
	PrefixSeparator = "_"

	// Base58 prefixes
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

	// Base 64 encoded transactions
	PrefixByteArray         = HashPrefix("ba_")
	PrefixContractByteArray = HashPrefix("cb_")
	PrefixOracleResponse    = HashPrefix("or_")
	PrefixOracleQuery       = HashPrefix("ov_")
	PrefixProofOfInclusion  = HashPrefix("pi_")
	PrefixStateTrees        = HashPrefix("ss_")
	PrefixState             = HashPrefix("st_")
	PrefixTransaction       = HashPrefix("tx_")
)

// store the encoding
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
