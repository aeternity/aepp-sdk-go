package aeternity

import (
	"fmt"

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
	senderID    string
	recipientID string
	amount      utils.BigInt
	fee         utils.BigInt
	payload     string
	ttl         uint64
	nonce       uint64
}

// RLP returns a byte serialized representation
func (t SpendTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	sID, err := buildIDTag(IDTagAccount, t.senderID)
	if err != nil {
		return
	}
	// build id for the recipient
	rID, err := buildIDTag(IDTagAccount, t.recipientID)
	if err != nil {
		return
	}
	amountBytes := t.amount.Bytes()
	fmt.Print(amountBytes)
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSpendTransaction,
		rlpMessageVersion,
		sID,
		rID,
		t.amount,
		t.fee,
		t.ttl,
		t.nonce,
		[]byte(t.payload))
	return
}

// NewSpendTx is a constructor for a SpendTx struct
func NewSpendTx(senderID, recipientID string, amount, fee utils.BigInt, payload string, ttl, nonce uint64) SpendTx {
	return SpendTx{senderID, recipientID, amount, fee, payload, ttl, nonce}
}

// NamePreclaimTx represents a transaction where one reserves a name on AENS without revealing it yet
type NamePreclaimTx struct {
	accountID    string
	commitmentID string
	fee          uint64
	ttl          uint64
	nonce        uint64
}

// RLP returns a byte serialized representation
func (t NamePreclaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.accountID)
	if err != nil {
		return
	}
	// build id for the committment
	cID, err := buildIDTag(IDTagCommitment, t.commitmentID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServicePreclaimTransaction,
		rlpMessageVersion,
		aID,
		t.nonce,
		cID,
		uint64(t.fee),
		t.ttl)
	return
}

// NewNamePreclaimTx is a constructor for a NamePreclaimTx struct
func NewNamePreclaimTx(accountID, commitmentID string, fee uint64, ttl, nonce uint64) NamePreclaimTx {
	return NamePreclaimTx{accountID, commitmentID, fee, ttl, nonce}
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
type NameClaimTx struct {
	accountID string
	name      string
	nameSalt  uint64
	fee       uint64
	ttl       uint64
	nonce     uint64
}

// RLP returns a byte serialized representation
func (t NameClaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.accountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, t.name)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		t.nonce,
		nID,
		uint64(t.nameSalt),
		uint64(t.fee),
		t.ttl)
	return
}

// NewNameClaimTx is a constructor for a NameClaimTx struct
func NewNameClaimTx(accountID, name string, nameSalt, fee uint64, ttl, nonce uint64) NameClaimTx {
	return NameClaimTx{accountID, name, nameSalt, fee, ttl, nonce}
}

// NameUpdateTx represents a transaction where one extends the lifetime of a reserved name on AENS
type NameUpdateTx struct {
	accountID string
	nameID    string
	pointers  []string
	nameTTL   uint64
	clientTTL uint64
	fee       uint64
	ttl       uint64
	nonce     uint64
}

// RLP returns a byte serialized representation
func (t NameUpdateTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, t.accountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, t.nameID)
	if err != nil {
		return
	}
	// build id for pointers
	ptrs, err := buildPointers(t.pointers)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		t.nonce,
		nID,
		uint64(t.nameTTL),
		ptrs,
		uint64(t.clientTTL),
		uint64(t.fee),
		t.ttl)
	return
}

// NewNameUpdateTx is a constructor for a NameUpdateTx struct
func NewNameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee uint64, ttl, nonce uint64) NameUpdateTx {
	return NameUpdateTx{accountID, nameID, pointers, nameTTL, clientTTL, fee, ttl, nonce}
}

// OracleRegisterTx represents a transaction that registers an oracle on the blockchain's state
type OracleRegisterTx struct {
	accountID      string
	accountNonce   uint64
	querySpec      string
	responseSpec   string
	queryFee       utils.BigInt
	oracleTTLType  uint64
	oracleTTLValue uint64
	abiVersion     uint64
	txFee          utils.BigInt
	txTTL          uint64
}

// RLP returns a byte serialized representation
func (t OracleRegisterTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the account
	aID, err := buildIDTag(IDTagAccount, t.accountID)
	if err != nil {
		return
	}
	qBytes := t.queryFee.Bytes()
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleRegisterTransaction,
		rlpMessageVersion,
		aID,
		t.accountNonce,
		[]byte(t.querySpec),
		[]byte(t.responseSpec),
		qBytes,
		t.oracleTTLType,
		t.oracleTTLValue,
		t.txFee.Bytes(),
		t.txTTL,
		t.abiVersion)
	return
}

// NewOracleRegisterTx is a constructor for a OracleRegisterTx struct
func NewOracleRegisterTx(accountID string, accountNonce uint64, querySpec, responseSpec string, queryFee utils.BigInt, oracleTTLType, oracleTTLValue, abiVersion uint64, txFee utils.BigInt, txTTL uint64) OracleRegisterTx {
	return OracleRegisterTx{accountID, accountNonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, abiVersion, txFee, txTTL}
}

// OracleExtendTx represents a transaction that extends the lifetime of an oracle
type OracleExtendTx struct {
	oracleID     string
	accountNonce uint64
	ttlType      uint64
	ttlValue     uint64
	fee          utils.BigInt
	ttl          uint64
}

// RLP returns a byte serialized representation
func (t OracleExtendTx) RLP() (rlpRawMsg []byte, err error) {
	aID, err := buildIDTag(IDTagOracle, t.oracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		aID,
		t.accountNonce,
		t.ttlType,
		t.ttlValue,
		t.fee.Bytes(),
		t.ttl)
	return
}

// NewOracleExtendTx is a constructor for a OracleExtendTx struct
func NewOracleExtendTx(oracleID string, accountNonce, ttlType, ttlValue uint64, fee utils.BigInt, ttl uint64) OracleExtendTx {
	return OracleExtendTx{oracleID, accountNonce, ttlType, ttlValue, fee, ttl}
}
