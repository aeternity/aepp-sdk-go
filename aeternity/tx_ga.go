package aeternity

import (
	"io"
	"math/big"

	rlp "github.com/randomshinichi/rlpae"
)

// GAAttachTx is a transaction that converts a plain old account (POA) into a
// Generalized Account. The function in the contract that should be used for
// authenticating transactions from that account is denoted by the parameter
// AuthFunc.
type GAAttachTx struct {
	OwnerID      string
	AccountNonce uint64
	Code         string
	AuthFunc     []byte
	VMVersion    uint16
	AbiVersion   uint16
	Gas          big.Int
	GasPrice     big.Int
	Fee          big.Int
	TTL          uint64
	CallData     string
}

// EncodeRLP implements rlp.Encoder
func (tx *GAAttachTx) EncodeRLP(w io.Writer) (err error) {
	aID, err := buildIDTag(IDTagAccount, tx.OwnerID)
	if err != nil {
		return
	}
	codeBinary, err := Decode(tx.Code)
	if err != nil {
		return
	}
	callDataBinary, err := Decode(tx.CallData)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagGeneralizedAccountAttachTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		codeBinary,
		tx.AuthFunc,
		encodeVMABI(tx.VMVersion, tx.AbiVersion),
		tx.Fee,
		tx.TTL,
		tx.Gas,
		tx.GasPrice,
		callDataBinary,
	)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// NewGAAttachTx creates a GAAttachTx
func NewGAAttachTx(OwnerID string, AccountNonce uint64, Code string, AuthFunc []byte, VMVersion uint16, AbiVersion uint16, Gas big.Int, GasPrice big.Int, Fee big.Int, TTL uint64, CallData string) GAAttachTx {
	return GAAttachTx{
		OwnerID:      OwnerID,
		AccountNonce: AccountNonce,
		Code:         Code,
		AuthFunc:     AuthFunc,
		VMVersion:    VMVersion,
		AbiVersion:   AbiVersion,
		Gas:          Gas,
		GasPrice:     GasPrice,
		Fee:          Fee,
		TTL:          TTL,
		CallData:     CallData,
	}
}

// GAMetaTx wraps a normal Tx (that is not a GAAttachTx) that will only be
// executed if the contract located at GaID returns true
type GAMetaTx struct {
	GaID       string
	AuthData   string
	AbiVersion uint16
	Gas        big.Int
	GasPrice   big.Int
	Fee        big.Int
	TTL        uint64
	Tx         rlp.Encoder
}

// EncodeRLP implements rlp.Encoder
func (tx *GAMetaTx) EncodeRLP(w io.Writer) (err error) {
	gaID, err := buildIDTag(IDTagContract, tx.GaID)
	if err != nil {
		return
	}
	authDataBinary, err := Decode(tx.AuthData)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagGeneralizedAccountMetaTransaction,
		rlpMessageVersion,
		gaID,
		authDataBinary,
		tx.AbiVersion,
		tx.Fee,
		tx.Gas,
		tx.GasPrice,
		tx.TTL,
		tx.Tx,
	)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

// NewGAMetaTx creates a GAMetaTx
func NewGAMetaTx(GaID string, AuthData string, AbiVersion uint16, Gas big.Int, GasPrice big.Int, Fee big.Int, TTL uint64, Tx rlp.Encoder) GAMetaTx {
	return GAMetaTx{
		GaID:       GaID,
		AuthData:   AuthData,
		AbiVersion: AbiVersion,
		Gas:        Gas,
		GasPrice:   GasPrice,
		Fee:        Fee,
		TTL:        TTL,
		Tx:         Tx,
	}
}
