package aeternity

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

// OracleRegisterTx represents a transaction that registers an oracle on the blockchain's state
type OracleRegisterTx struct {
	AccountID      string
	AccountNonce   uint64
	QuerySpec      string
	ResponseSpec   string
	QueryFee       big.Int
	OracleTTLType  uint64
	OracleTTLValue uint64
	AbiVersion     uint16
	Fee            big.Int
	TTL            uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *OracleRegisterTx) EncodeRLP(w io.Writer) (err error) {
	// build id for the account
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagOracleRegisterTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		[]byte(tx.QuerySpec),
		[]byte(tx.ResponseSpec),
		tx.QueryFee,
		tx.OracleTTLType,
		tx.OracleTTLValue,
		tx.Fee,
		tx.TTL,
		tx.AbiVersion)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
// BUG: Account Nonce won'tx be represented in JSON output if nonce is 0, thanks to swagger.json
func (tx *OracleRegisterTx) JSON() (string, error) {
	// # Oracles
	// ORACLE_TTL_TYPE_DELTA = 'delta'
	// ORACLE_TTL_TYPE_BLOCK = 'block'
	// From reading the code, 0 is "delta", 1 is "block"
	// # VM Identifiers
	// # vm version specification
	// # https://github.com/aeternity/protocol/blob/master/contracts/contract_vms.md#virtual-machines-on-the-%C3%A6ternity-blockchain
	// NO_VM = 0
	// VM_SOPHIA = 1
	// VM_SOLIDITY = 2
	// VM_SOPHIA_IMPROVEMENTS = 3
	// # abi
	// NO_ABI = 0
	// ABI_SOPHIA = 1
	// ABI_SOLIDITY = 2

	ttlTypeStr := ttlTypeIntToStr(tx.OracleTTLType)

	swaggerT := models.OracleRegisterTx{
		AbiVersion: tx.AbiVersion,
		AccountID:  &tx.AccountID,
		Fee:        utils.BigInt(tx.Fee),
		Nonce:      tx.AccountNonce,
		OracleTTL: &models.TTL{
			Type:  &ttlTypeStr,
			Value: &tx.OracleTTLValue,
		},
		QueryFee:       utils.BigInt(tx.QueryFee),
		QueryFormat:    &tx.QuerySpec,
		ResponseFormat: &tx.ResponseSpec,
		TTL:            tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleRegisterTx is a constructor for a OracleRegisterTx struct
func NewOracleRegisterTx(accountID string, accountNonce uint64, querySpec, responseSpec string, queryFee big.Int, oracleTTLType, oracleTTLValue uint64, abiVersion uint16, txFee big.Int, txTTL uint64) OracleRegisterTx {
	return OracleRegisterTx{accountID, accountNonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, abiVersion, txFee, txTTL}
}

// OracleExtendTx represents a transaction that extends the lifetime of an oracle
type OracleExtendTx struct {
	OracleID       string
	AccountNonce   uint64
	OracleTTLType  uint64
	OracleTTLValue uint64
	Fee            big.Int
	TTL            uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *OracleExtendTx) EncodeRLP(w io.Writer) (err error) {
	oID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		oID,
		tx.AccountNonce,
		tx.OracleTTLType,
		tx.OracleTTLValue,
		tx.Fee,
		tx.TTL)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *OracleExtendTx) JSON() (string, error) {
	oracleTTLTypeStr := ttlTypeIntToStr(tx.OracleTTLType)

	swaggerT := models.OracleExtendTx{
		Fee:      utils.BigInt(tx.Fee),
		Nonce:    tx.AccountNonce,
		OracleID: &tx.OracleID,
		OracleTTL: &models.RelativeTTL{
			Type:  &oracleTTLTypeStr,
			Value: &tx.OracleTTLValue,
		},
		TTL: tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleExtendTx is a constructor for a OracleExtendTx struct
func NewOracleExtendTx(oracleID string, accountNonce, oracleTTLType, oracleTTLValue uint64, Fee big.Int, TTL uint64) OracleExtendTx {
	return OracleExtendTx{oracleID, accountNonce, oracleTTLType, oracleTTLValue, Fee, TTL}
}

// OracleQueryTx represents a transaction that a program sends to query an oracle
type OracleQueryTx struct {
	SenderID         string
	AccountNonce     uint64
	OracleID         string
	Query            string
	QueryFee         big.Int
	QueryTTLType     uint64
	QueryTTLValue    uint64
	ResponseTTLType  uint64
	ResponseTTLValue uint64
	Fee              big.Int
	TTL              uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *OracleQueryTx) EncodeRLP(w io.Writer) (err error) {
	accountID, err := buildIDTag(IDTagAccount, tx.SenderID)
	if err != nil {
		return
	}

	oracleID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagOracleQueryTransaction,
		rlpMessageVersion,
		accountID,
		tx.AccountNonce,
		oracleID,
		[]byte(tx.Query),
		tx.QueryFee,
		tx.QueryTTLType,
		tx.QueryTTLValue,
		tx.ResponseTTLType,
		tx.ResponseTTLValue,
		tx.Fee,
		tx.TTL)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *OracleQueryTx) JSON() (string, error) {
	responseTTLTypeStr := ttlTypeIntToStr(tx.ResponseTTLType)
	queryTTLTypeStr := ttlTypeIntToStr(tx.QueryTTLType)

	swaggerT := models.OracleQueryTx{
		Fee:      utils.BigInt(tx.Fee),
		Nonce:    tx.AccountNonce,
		OracleID: &tx.OracleID,
		Query:    &tx.Query,
		QueryFee: utils.BigInt(tx.QueryFee),
		QueryTTL: &models.TTL{
			Type:  &queryTTLTypeStr,
			Value: &tx.QueryTTLValue,
		},
		ResponseTTL: &models.RelativeTTL{
			Type:  &responseTTLTypeStr,
			Value: &tx.ResponseTTLValue,
		},
		SenderID: &tx.SenderID,
		TTL:      tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleQueryTx is a constructor for a OracleQueryTx struct
func NewOracleQueryTx(SenderID string, AccountNonce uint64, OracleID, Query string, QueryFee big.Int, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64, Fee big.Int, TTL uint64) OracleQueryTx {
	return OracleQueryTx{SenderID, AccountNonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, Fee, TTL}
}

// OracleRespondTx represents a transaction that an oracle sends to respond to an incoming query
type OracleRespondTx struct {
	OracleID         string
	AccountNonce     uint64
	QueryID          string
	Response         string
	ResponseTTLType  uint64
	ResponseTTLValue uint64
	Fee              big.Int
	TTL              uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *OracleRespondTx) EncodeRLP(w io.Writer) (err error) {
	oID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}
	queryIDBytes, err := Decode(tx.QueryID)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagOracleResponseTransaction,
		rlpMessageVersion,
		oID,
		tx.AccountNonce,
		queryIDBytes,
		tx.Response,
		tx.ResponseTTLType,
		tx.ResponseTTLValue,
		tx.Fee,
		tx.TTL)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *OracleRespondTx) JSON() (string, error) {
	responseTTLTypeStr := ttlTypeIntToStr(tx.ResponseTTLType)

	swaggerT := models.OracleRespondTx{
		Fee:      utils.BigInt(tx.Fee),
		Nonce:    tx.AccountNonce,
		OracleID: &tx.OracleID,
		QueryID:  &tx.QueryID,
		Response: &tx.Response,
		ResponseTTL: &models.RelativeTTL{
			Type:  &responseTTLTypeStr,
			Value: &tx.ResponseTTLValue,
		},
		TTL: tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err

}

// NewOracleRespondTx is a constructor for a OracleRespondTx struct
func NewOracleRespondTx(OracleID string, AccountNonce uint64, QueryID string, Response string, TTLType uint64, TTLValue uint64, Fee big.Int, TTL uint64) OracleRespondTx {
	return OracleRespondTx{OracleID, AccountNonce, QueryID, Response, TTLType, TTLValue, Fee, TTL}
}
