package aeternity

import (
	"fmt"
	"io"

	"github.com/aeternity/aepp-sdk-go/generated/models"
	"github.com/aeternity/aepp-sdk-go/rlp"
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

func ttlTypeIntToStr(i uint64) string {
	var oracleTTLTypeStr string
	if i == 0 {
		oracleTTLTypeStr = "delta"
	} else {
		oracleTTLTypeStr = "block"
	}
	return oracleTTLTypeStr
}

func buildPointers(pointers []string) (ptrs []*NamePointer, err error) {
	// TODO: handle errors
	ptrs = make([]*NamePointer, len(pointers))
	for i, p := range pointers {
		switch GetHashPrefix(p) {
		case PrefixAccountPubkey:
			// pID, err := buildIDTag(IDTagAccount, p)
			key := "account_pubkey"
			ptrs[i] = NewNamePointer(key, p)
			if err != nil {
				break
			}
		case PrefixOraclePubkey:
			// pID, err := buildIDTag(IDTagOracle, p)
			key := "oracle_pubkey"
			ptrs[i] = NewNamePointer(key, p)
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

// JSON representation of a Tx is useful for querying the node's debug endpoint
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
	Fee          utils.BigInt
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
		t.Fee.Int,
		t.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (t *NamePreclaimTx) JSON() (string, error) {
	swaggerT := models.NamePreclaimTx{
		AccountID:    models.EncodedHash(t.AccountID),
		CommitmentID: models.EncodedHash(t.CommitmentID),
		Fee:          t.Fee,
		Nonce:        t.Nonce,
		TTL:          t.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewNamePreclaimTx is a constructor for a NamePreclaimTx struct
func NewNamePreclaimTx(accountID, commitmentID string, fee utils.BigInt, ttl, nonce uint64) NamePreclaimTx {
	return NamePreclaimTx{accountID, commitmentID, fee, ttl, nonce}
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
// The revealed name is simply sent in plaintext in RLP, while in JSON representation
// it is base58 encoded.
type NameClaimTx struct {
	AccountID string
	Name      string
	NameSalt  utils.BigInt
	Fee       utils.BigInt
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

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		t.Nonce,
		t.Name,
		t.NameSalt.Int,
		t.Fee.Int,
		t.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoints
func (t *NameClaimTx) JSON() (string, error) {
	// When talking JSON to the node, the name should be 'API encoded'
	// (base58), not namehash-ed.
	nameAPIEncoded := Encode(PrefixName, []byte(t.Name))
	swaggerT := models.NameClaimTx{
		AccountID: models.EncodedHash(t.AccountID),
		Fee:       t.Fee,
		Name:      &nameAPIEncoded,
		NameSalt:  t.NameSalt,
		Nonce:     t.Nonce,
		TTL:       t.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewNameClaimTx is a constructor for a NameClaimTx struct
func NewNameClaimTx(accountID, name string, nameSalt utils.BigInt, fee utils.BigInt, ttl, nonce uint64) NameClaimTx {
	return NameClaimTx{accountID, name, nameSalt, fee, ttl, nonce}
}

// NamePointer extends the swagger gener ated models.NamePointer to provide RLP serialization
type NamePointer struct {
	*models.NamePointer
}

// EncodeRLP implements rlp.Encoder interface.
func (t *NamePointer) EncodeRLP(w io.Writer) (err error) {
	accountID, err := buildIDTag(IDTagAccount, string(t.NamePointer.ID))
	if err != nil {
		return
	}

	err = rlp.Encode(w, []interface{}{t.Key, accountID})
	if err != nil {
		return
	}
	return err
}

// NewNamePointer is a constructor for a swagger generated NamePointer struct.
// It returns a pointer because
func NewNamePointer(key string, id string) *NamePointer {
	np := models.NamePointer{ID: models.EncodedHash(id), Key: &key}
	return &NamePointer{
		NamePointer: &np,
	}
}

// NameUpdateTx represents a transaction where one extends the lifetime of a reserved name on AENS
type NameUpdateTx struct {
	AccountID string
	NameID    string
	Pointers  []*NamePointer
	NameTTL   uint64
	ClientTTL uint64
	Fee       utils.BigInt
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
	// build id for the name
	nID, err := buildIDTag(IDTagName, t.NameID)
	if err != nil {
		return
	}

	// reverse the NamePointer order as compared to the JSON serialization, because the node seems to want it that way
	i := 0
	j := len(t.Pointers) - 1
	reversedPointers := make([]*NamePointer, len(t.Pointers))
	for i <= len(t.Pointers)-1 {
		reversedPointers[i] = t.Pointers[j]
		i++
		j--
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceUpdateTransaction,
		rlpMessageVersion,
		aID,
		t.Nonce,
		nID,
		t.NameTTL,
		reversedPointers,
		t.ClientTTL,
		t.Fee.Int,
		t.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (t *NameUpdateTx) JSON() (string, error) {
	swaggerNamePointers := []*models.NamePointer{}
	for _, np := range t.Pointers {
		swaggerNamePointers = append(swaggerNamePointers, np.NamePointer)
	}

	swaggerT := models.NameUpdateTx{
		AccountID: models.EncodedHash(t.AccountID),
		ClientTTL: &t.ClientTTL,
		Fee:       t.Fee,
		NameID:    models.EncodedHash(t.NameID),
		NameTTL:   &t.NameTTL,
		Nonce:     t.Nonce,
		Pointers:  swaggerNamePointers,
		TTL:       t.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewNameUpdateTx is a constructor for a NameUpdateTx struct
func NewNameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee utils.BigInt, ttl, nonce uint64) NameUpdateTx {
	parsedPointers, err := buildPointers(pointers)
	if err != nil {
		panic(err)
	}
	return NameUpdateTx{accountID, nameID, parsedPointers, nameTTL, clientTTL, fee, ttl, nonce}
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

// JSON representation of a Tx is useful for querying the node's debug endpoint
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

	ttlTypeStr := ttlTypeIntToStr(t.OracleTTLType)

	swaggerT := models.OracleRegisterTx{
		AbiVersion: &abiVersionCasted,
		AccountID:  models.EncodedHash(t.AccountID),
		Fee:        t.TxFee,
		Nonce:      t.AccountNonce,
		OracleTTL: &models.TTL{
			Type:  &ttlTypeStr,
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
	oID, err := buildIDTag(IDTagOracle, t.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		oID,
		t.AccountNonce,
		t.TTLType,
		t.TTLValue,
		t.Fee.Int,
		t.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (t *OracleExtendTx) JSON() (string, error) {
	oracleTTLTypeStr := ttlTypeIntToStr(t.TTLType)

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

// OracleQueryTx represents a transaction that a program sends to query an oracle
type OracleQueryTx struct {
	SenderID         string
	AccountNonce     uint64
	OracleID         string
	Query            string
	QueryFee         utils.BigInt
	QueryTTLType     uint64
	QueryTTLValue    uint64
	ResponseTTLType  uint64
	ResponseTTLValue uint64
	TxFee            utils.BigInt
	TxTTL            uint64
}

// RLP returns a byte serialized representation
func (t *OracleQueryTx) RLP() (rlpRawMsg []byte, err error) {
	accountID, err := buildIDTag(IDTagAccount, t.SenderID)
	if err != nil {
		return
	}

	oracleID, err := buildIDTag(IDTagOracle, t.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleQueryTransaction,
		rlpMessageVersion,
		accountID,
		t.AccountNonce,
		oracleID,
		[]byte(t.Query),
		t.QueryFee.Int,
		t.QueryTTLType,
		t.QueryTTLValue,
		t.ResponseTTLType,
		t.ResponseTTLValue,
		t.TxFee.Int,
		t.TxTTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (t *OracleQueryTx) JSON() (string, error) {
	responseTTLTypeStr := ttlTypeIntToStr(t.ResponseTTLType)
	queryTTLTypeStr := ttlTypeIntToStr(t.QueryTTLType)

	swaggerT := models.OracleQueryTx{
		Fee:      t.TxFee,
		Nonce:    t.AccountNonce,
		OracleID: models.EncodedHash(t.OracleID),
		Query:    &t.Query,
		QueryFee: t.QueryFee,
		QueryTTL: &models.TTL{
			Type:  &queryTTLTypeStr,
			Value: &t.QueryTTLValue,
		},
		ResponseTTL: &models.RelativeTTL{
			Type:  &responseTTLTypeStr,
			Value: &t.ResponseTTLValue,
		},
		SenderID: models.EncodedHash(t.SenderID),
		TTL:      t.TxTTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// NewOracleQueryTx is a constructor for a OracleQueryTx struct
func NewOracleQueryTx(SenderID string, AccountNonce uint64, OracleID, Query string, QueryFee utils.BigInt, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64, TxFee utils.BigInt, TxTTL uint64) OracleQueryTx {
	return OracleQueryTx{SenderID, AccountNonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, TxFee, TxTTL}
}

// OracleRespondTx represents a transaction that an oracle sends to respond to an incoming query
type OracleRespondTx struct {
	OracleID         string
	AccountNonce     uint64
	QueryID          string
	Response         string
	ResponseTTLType  uint64
	ResponseTTLValue uint64
	TxFee            utils.BigInt
	TxTTL            uint64
}

// RLP returns a byte serialized representation
func (t *OracleRespondTx) RLP() (rlpRawMsg []byte, err error) {
	oID, err := buildIDTag(IDTagOracle, t.OracleID)
	if err != nil {
		return
	}
	queryIDBytes, err := Decode(t.QueryID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleResponseTransaction,
		rlpMessageVersion,
		oID,
		t.AccountNonce,
		queryIDBytes,
		t.Response,
		t.ResponseTTLType,
		t.ResponseTTLValue,
		t.TxFee.Int,
		t.TxTTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (t *OracleRespondTx) JSON() (string, error) {
	responseTTLTypeStr := ttlTypeIntToStr(t.ResponseTTLType)

	swaggerT := models.OracleRespondTx{
		Fee:      t.TxFee,
		Nonce:    t.AccountNonce,
		OracleID: models.EncodedHash(t.OracleID),
		QueryID:  models.EncodedHash(t.QueryID),
		Response: &t.Response,
		ResponseTTL: &models.RelativeTTL{
			Type:  &responseTTLTypeStr,
			Value: &t.ResponseTTLValue,
		},
		TTL: t.TxTTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err

}

// NewOracleRespondTx is a constructor for a OracleRespondTx struct
func NewOracleRespondTx(OracleID string, AccountNonce uint64, QueryID string, Response string, TTLType uint64, TTLValue uint64, TxFee utils.BigInt, TxTTL uint64) OracleRespondTx {
	return OracleRespondTx{OracleID, AccountNonce, QueryID, Response, TTLType, TTLValue, TxFee, TxTTL}
}
