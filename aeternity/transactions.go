package aeternity

import (
	"fmt"

	"github.com/aeternity/aepp-sdk-go/generated/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

// SignEncodeTx sign and encode a transaction
func SignEncodeTx(kp *Account, txRaw []byte, networkID string) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
	// add the network_id to the transaction
	msg := append([]byte(networkID), txRaw...)
	// sign the transaction
	sigRaw := kp.Sign(msg)
	if err != nil {
		return
	}
	// encode the message using rlp
	rlpTxRaw, err := createSignedTransaction(txRaw, [][]byte{sigRaw})
	// encode the rlp message with the prefix
	signedEncodedTx = Encode(PrefixTransaction, rlpTxRaw)
	// compute the hash
	rlpTxHashRaw, err := hash(rlpTxRaw)
	signedEncodedTxHash = Encode(PrefixTransactionHash, rlpTxHashRaw)
	// encode the signature
	signature = Encode(PrefixSignature, sigRaw)
	return
}

func createSignedTransaction(txRaw []byte, signatures [][]byte) (rlpRawMsg []byte, err error) {
	// encode the message using rlp
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSignedTransaction,
		rlpMessageVersion,
		signatures,
		txRaw,
	)
	return
}

func buildPointers(pointers []string) (ptrs [][]uint8, err error) {
	// TODO: handle errors
	ptrs = make([][]uint8, len(pointers))
	for i, p := range pointers {
		switch GetHashPrefix(p) {
		case PrefixAccountPubkey:
			pID, err := buildIDTag(IDTagName, p)
			ptrs[i] = pID
			if err != nil {
				break
			}
		case PrefixOraclePubkey:
			pID, err := buildIDTag(IDTagOracle, p)
			ptrs[i] = pID
			if err != nil {
				break
			}
		default:
			err = fmt.Errorf("Invalid ID %v for pointers", p)
		}
	}
	return
}

// Tx interface guarantees that code using Tx can rely on these functions being present
// Since the methods to Tx-like structs do not modify the Tx themselves - they simply generate values
// from the Tx's values - Tx methods should not have a pointer receiver.
// See https://tour.golang.org/methods/4 or https://dave.cheney.net/2016/03/19/should-methods-be-declared-on-t-or-t
type Tx interface {
	RLP() ([]byte, error)
}

// BaseEncodeTx takes a Tx, runs its RLP() method, and base encodes the result.
func BaseEncodeTx(t Tx) (string, error) {
	txRaw, err := t.RLP()
	if err != nil {
		return "", err
	}
	txStr := Encode(PrefixTransaction, txRaw)
	return txStr, nil
}

// SpendTx represents a simple transaction where one party sends another AE
type SpendTx struct {
	SenderID    string
	RecipientID string
	Amount      utils.BigInt
	Fee         utils.BigInt
	Payload     string
	TTL         uint64
	Nonce       uint64
}

// RLP returns a byte serialized representation
func (t *SpendTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	sID, err := buildIDTag(IDTagAccount, t.SenderID)
	if err != nil {
		return
	}
	// build id for the recipient
	rID, err := buildIDTag(IDTagAccount, t.RecipientID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSpendTransaction,
		rlpMessageVersion,
		sID,
		rID,
		t.Amount.Int,
		t.Fee.Int,
		t.TTL,
		t.Nonce,
		[]byte(t.Payload))
	return
}

func (t *SpendTx) JSON() (string, error) {
	swaggerT := models.SpendTx{
		Amount:      t.Amount,
		Fee:         t.Fee,
		Nonce:       t.Nonce,
		Payload:     &t.Payload,
		RecipientID: models.EncodedHash(t.RecipientID),
		SenderID:    models.EncodedHash(t.SenderID),
		TTL:         t.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewSpendTx is a constructor for a SpendTx struct
func NewSpendTx(senderID, recipientID string, amount, fee utils.BigInt, payload string, ttl, nonce uint64) SpendTx {
	return SpendTx{senderID, recipientID, amount, fee, payload, ttl, nonce}
}

// NamePreclaimTx represents a transaction where one reserves a name on AENS without revealing it yet
type NamePreclaimTx struct {
	AccountID    string
	CommitmentID string
	Fee          uint64
	TTL          uint64
	Nonce        uint64
}

// RLP returns a byte serialized representation
func (t *NamePreclaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.AccountID)
	if err != nil {
		return
	}
	// build id for the committment
	cID, err := buildIDTag(IDTagCommitment, t.CommitmentID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServicePreclaimTransaction,
		rlpMessageVersion,
		aID,
		t.Nonce,
		cID,
		uint64(t.Fee),
		t.TTL)
	return
}

// NewNamePreclaimTx is a constructor for a NamePreclaimTx struct
func NewNamePreclaimTx(accountID, commitmentID string, fee uint64, ttl, nonce uint64) NamePreclaimTx {
	return NamePreclaimTx{accountID, commitmentID, fee, ttl, nonce}
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
type NameClaimTx struct {
	AccountID string
	Name      string
	NameSalt  uint64
	Fee       uint64
	TTL       uint64
	Nonce     uint64
}

// RLP returns a byte serialized representation
func (t *NameClaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.AccountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, t.Name)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		t.Nonce,
		nID,
		uint64(t.NameSalt),
		uint64(t.Fee),
		t.TTL)
	return
}

// NewNameClaimTx is a constructor for a NameClaimTx struct
func NewNameClaimTx(accountID, name string, nameSalt, fee uint64, ttl, nonce uint64) NameClaimTx {
	return NameClaimTx{accountID, name, nameSalt, fee, ttl, nonce}
}

// NameUpdateTx represents a transaction where one extends the lifetime of a reserved name on AENS
type NameUpdateTx struct {
	AccountID string
	NameID    string
	Pointers  []string
	NameTTL   uint64
	ClientTTL uint64
	Fee       uint64
	TTL       uint64
	Nonce     uint64
}

// RLP returns a byte serialized representation
func (t *NameUpdateTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.AccountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, t.NameID)
	if err != nil {
		return
	}
	// build id for pointers
	ptrs, err := buildPointers(t.Pointers)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		t.Nonce,
		nID,
		uint64(t.NameTTL),
		ptrs,
		uint64(t.ClientTTL),
		uint64(t.Fee),
		t.TTL)
	return
}

// NewNameUpdateTx is a constructor for a NameUpdateTx struct
func NewNameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee uint64, ttl, nonce uint64) NameUpdateTx {
	return NameUpdateTx{accountID, nameID, pointers, nameTTL, clientTTL, fee, ttl, nonce}
}

// OracleRegisterTx represents a transaction that registers an oracle on the blockchain's state
type OracleRegisterTx struct {
	AccountID      string
	AccountNonce   uint64
	QuerySpec      string
	ResponseSpec   string
	QueryFee       utils.BigInt
	OracleTTLType  uint64
	OracleTTLValue uint64
	AbiVersion     uint64
	VMVersion      uint64
	TxFee          utils.BigInt
	TxTTL          uint64
}

// RLP returns a byte serialized representation
func (t *OracleRegisterTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the account
	aID, err := buildIDTag(IDTagAccount, t.AccountID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleRegisterTransaction,
		rlpMessageVersion,
		aID,
		t.AccountNonce,
		[]byte(t.QuerySpec),
		[]byte(t.ResponseSpec),
		t.QueryFee.Int,
		t.OracleTTLType,
		t.OracleTTLValue,
		t.TxFee.Int,
		t.TxTTL,
		t.AbiVersion)
	return
}

// BUG: Account Nonce won't be represented in JSON output if nonce is 0, thanks to swagger.json
func (t *OracleRegisterTx) JSON() (string, error) {
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
	abiVersionCasted := int64(t.AbiVersion)
	vmVersionCasted := int64(t.VMVersion)

	var oracleTTLTypeStr string
	if t.OracleTTLType == 0 {
		oracleTTLTypeStr = "delta"
	} else {
		oracleTTLTypeStr = "block"
	}

	swaggerT := models.OracleRegisterTx{
		AbiVersion: &abiVersionCasted,
		AccountID:  models.EncodedHash(t.AccountID),
		Fee:        t.TxFee,
		Nonce:      t.AccountNonce,
		OracleTTL: &models.TTL{
			Type:  &oracleTTLTypeStr,
			Value: &t.OracleTTLValue,
		},
		QueryFee:       t.QueryFee,
		QueryFormat:    &t.QuerySpec,
		ResponseFormat: &t.ResponseSpec,
		TTL:            t.TxTTL,
		VMVersion:      &vmVersionCasted,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleRegisterTx is a constructor for a OracleRegisterTx struct
func NewOracleRegisterTx(accountID string, accountNonce uint64, querySpec, responseSpec string, queryFee utils.BigInt, oracleTTLType, oracleTTLValue, abiVersion uint64, vmVersion uint64, txFee utils.BigInt, txTTL uint64) OracleRegisterTx {
	return OracleRegisterTx{accountID, accountNonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, abiVersion, vmVersion, txFee, txTTL}
}

// OracleExtendTx represents a transaction that extends the lifetime of an oracle
type OracleExtendTx struct {
	OracleID     string
	AccountNonce uint64
	TTLType      uint64
	TTLValue     uint64
	Fee          utils.BigInt
	TTL          uint64
}

// RLP returns a byte serialized representation
func (t *OracleExtendTx) RLP() (rlpRawMsg []byte, err error) {
	aID, err := buildIDTag(IDTagOracle, t.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		aID,
		t.AccountNonce,
		t.TTLType,
		t.TTLValue,
		t.Fee.Bytes(),
		t.TTL)
	return
}

func (t *OracleExtendTx) JSON() (string, error) {
	var oracleTTLTypeStr string
	if t.TTLType == 0 {
		oracleTTLTypeStr = "delta"
	} else {
		oracleTTLTypeStr = "block"
	}

	swaggerT := models.OracleExtendTx{
		Fee:      t.Fee,
		Nonce:    t.AccountNonce,
		OracleID: models.EncodedHash(t.OracleID),
		OracleTTL: &models.RelativeTTL{
			Type:  &oracleTTLTypeStr,
			Value: &t.TTLValue,
		},
		TTL: t.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleExtendTx is a constructor for a OracleExtendTx struct
func NewOracleExtendTx(oracleID string, accountNonce, ttlType, ttlValue uint64, fee utils.BigInt, ttl uint64) OracleExtendTx {
	return OracleExtendTx{oracleID, accountNonce, ttlType, ttlValue, fee, ttl}
}
