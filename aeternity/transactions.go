package aeternity

import (
	"fmt"
	"io"
	"math/big"

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

func calcFeeStd(tx Tx, txLen int) *big.Int {
	// (Config.Client.BaseGas + len(txRLP) * Config.Client.GasPerByte) * Config.Client.GasPrice
	//                                   txLenGasPerByte
	fee := new(big.Int)
	txLenGasPerByte := new(big.Int)

	txLenGasPerByte.Mul(utils.NewIntFromUint64(uint64(txLen)), &Config.Client.GasPerByte)
	fee.Add(&Config.Client.BaseGas, txLenGasPerByte)
	fee.Mul(fee, &Config.Client.GasPrice)
	return fee
}

func calcFeeContract(gas *big.Int, baseGasMultiplier int64, length int) *big.Int {
	// (Config.Client.BaseGas * 5) + gaslimit + (len(txRLP) * Config.Client.GasPerByte) * Config.Client.GasPrice
	//           baseGas5                                txLenGasPerByte
	baseGas5 := new(big.Int)
	txLenBig := new(big.Int)
	answer := new(big.Int)

	baseGas5.Mul(&Config.Client.BaseGas, big.NewInt(baseGasMultiplier))
	txLenBig.SetUint64(uint64(length))
	txLenGasPerByte := new(big.Int)
	txLenGasPerByte.Mul(txLenBig, &Config.Client.GasPerByte)

	answer.Add(baseGas5, gas)
	answer.Add(answer, txLenGasPerByte)
	answer.Mul(answer, &Config.Client.GasPrice)
	return answer
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func calcSizeEstimate(tx Tx, fee *big.Int) (int, error) {
	feeRlp, err := rlp.EncodeToBytes(fee)
	if err != nil {
		return 0, err
	}
	feeRlpLen := len(feeRlp)

	rlpRawMsg, err := tx.RLP()
	if err != nil {
		return 0, err
	}

	return len(rlpRawMsg) - feeRlpLen + 8, nil
}

// Tx interface guarantees that code using Tx can rely on these functions being present
type Tx interface {
	RLP() ([]byte, error)
}

// BaseEncodeTx takes a Tx, runs its RLP() method, and base encodes the result.
func BaseEncodeTx(tx Tx) (string, error) {
	txRaw, err := tx.RLP()
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
	Amount      big.Int
	Fee         big.Int
	Payload     string
	TTL         uint64
	Nonce       uint64
}

// RLP returns a byte serialized representation
func (tx *SpendTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	sID, err := buildIDTag(IDTagAccount, tx.SenderID)
	if err != nil {
		return
	}
	// build id for the recipient
	rID, err := buildIDTag(IDTagAccount, tx.RecipientID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSpendTransaction,
		rlpMessageVersion,
		sID,
		rID,
		tx.Amount,
		tx.Fee,
		tx.TTL,
		tx.Nonce,
		[]byte(tx.Payload))
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *SpendTx) JSON() (string, error) {
	swaggerT := models.SpendTx{
		Amount:      utils.BigInt(tx.Amount),
		Fee:         utils.BigInt(tx.Fee),
		Nonce:       tx.Nonce,
		Payload:     &tx.Payload,
		RecipientID: &tx.RecipientID,
		SenderID:    &tx.SenderID,
		TTL:         tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *SpendTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *SpendTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewSpendTx is a constructor for a SpendTx struct
func NewSpendTx(senderID, recipientID string, amount, fee big.Int, payload string, ttl, nonce uint64) SpendTx {
	return SpendTx{senderID, recipientID, amount, fee, payload, ttl, nonce}
}

// NamePreclaimTx represents a transaction where one reserves a name on AENS without revealing it yet
type NamePreclaimTx struct {
	AccountID    string
	CommitmentID string
	Fee          big.Int
	TTL          uint64
	AccountNonce uint64
}

// RLP returns a byte serialized representation
func (tx *NamePreclaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// build id for the committment
	cID, err := buildIDTag(IDTagCommitment, tx.CommitmentID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServicePreclaimTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		cID,
		tx.Fee,
		tx.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NamePreclaimTx) JSON() (string, error) {
	swaggerT := models.NamePreclaimTx{
		AccountID:    &tx.AccountID,
		CommitmentID: &tx.CommitmentID,
		Fee:          utils.BigInt(tx.Fee),
		Nonce:        tx.AccountNonce,
		TTL:          tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NamePreclaimTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *NamePreclaimTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewNamePreclaimTx is a constructor for a NamePreclaimTx struct
func NewNamePreclaimTx(accountID, commitmentID string, fee big.Int, ttl, accountNonce uint64) NamePreclaimTx {
	return NamePreclaimTx{accountID, commitmentID, fee, ttl, accountNonce}
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
// The revealed name is simply sent in plaintext in RLP, while in JSON representation
// it is base58 encoded.
type NameClaimTx struct {
	AccountID    string
	Name         string
	NameSalt     big.Int
	Fee          big.Int
	TTL          uint64
	AccountNonce uint64
}

// RLP returns a byte serialized representation
func (tx *NameClaimTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		tx.Name,
		tx.NameSalt,
		tx.Fee,
		tx.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoints
func (tx *NameClaimTx) JSON() (string, error) {
	// When talking JSON to the node, the name should be 'API encoded'
	// (base58), not namehash-ed.
	nameAPIEncoded := Encode(PrefixName, []byte(tx.Name))
	swaggerT := models.NameClaimTx{
		AccountID: &tx.AccountID,
		Fee:       utils.BigInt(tx.Fee),
		Name:      &nameAPIEncoded,
		NameSalt:  utils.BigInt(tx.NameSalt),
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameClaimTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *NameClaimTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewNameClaimTx is a constructor for a NameClaimTx struct
func NewNameClaimTx(accountID, name string, nameSalt big.Int, fee big.Int, ttl, accountNonce uint64) NameClaimTx {
	return NameClaimTx{accountID, name, nameSalt, fee, ttl, accountNonce}
}

// NamePointer extends the swagger gener ated models.NamePointer to provide RLP serialization
type NamePointer struct {
	*models.NamePointer
}

// EncodeRLP implements rlp.Encoder interface.
func (np *NamePointer) EncodeRLP(w io.Writer) (err error) {
	accountID, err := buildIDTag(IDTagAccount, *np.NamePointer.ID)
	if err != nil {
		return
	}

	err = rlp.Encode(w, []interface{}{np.Key, accountID})
	if err != nil {
		return
	}
	return err
}

// NewNamePointer is a constructor for a swagger generated NamePointer struct.
// It returns a pointer because
func NewNamePointer(key string, id string) *NamePointer {
	np := models.NamePointer{ID: &id, Key: &key}
	return &NamePointer{
		NamePointer: &np,
	}
}

// NameUpdateTx represents a transaction where one extends the lifetime of a reserved name on AENS
type NameUpdateTx struct {
	AccountID    string
	NameID       string
	Pointers     []*NamePointer
	NameTTL      uint64
	ClientTTL    uint64
	Fee          big.Int
	TTL          uint64
	AccountNonce uint64
}

// RLP returns a byte serialized representation
func (tx *NameUpdateTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// build id for the name
	nID, err := buildIDTag(IDTagName, tx.NameID)
	if err != nil {
		return
	}

	// reverse the NamePointer order as compared to the JSON serialization, because the node seems to want it that way
	i := 0
	j := len(tx.Pointers) - 1
	reversedPointers := make([]*NamePointer, len(tx.Pointers))
	for i <= len(tx.Pointers)-1 {
		reversedPointers[i] = tx.Pointers[j]
		i++
		j--
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceUpdateTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		nID,
		tx.NameTTL,
		reversedPointers,
		tx.ClientTTL,
		tx.Fee,
		tx.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameUpdateTx) JSON() (string, error) {
	swaggerNamePointers := []*models.NamePointer{}
	for _, np := range tx.Pointers {
		swaggerNamePointers = append(swaggerNamePointers, np.NamePointer)
	}

	swaggerT := models.NameUpdateTx{
		AccountID: &tx.AccountID,
		ClientTTL: &tx.ClientTTL,
		Fee:       utils.BigInt(tx.Fee),
		NameID:    &tx.NameID,
		NameTTL:   &tx.NameTTL,
		Nonce:     tx.AccountNonce,
		Pointers:  swaggerNamePointers,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameUpdateTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *NameUpdateTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewNameUpdateTx is a constructor for a NameUpdateTx struct
func NewNameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee big.Int, ttl, accountNonce uint64) NameUpdateTx {
	parsedPointers, err := buildPointers(pointers)
	if err != nil {
		panic(err)
	}
	return NameUpdateTx{accountID, nameID, parsedPointers, nameTTL, clientTTL, fee, ttl, accountNonce}
}

// NameRevokeTx represents a transaction that revokes the name, i.e. has the same effect as waiting for the Name's TTL to expire.
type NameRevokeTx struct {
	AccountID    string
	NameID       string
	Fee          big.Int
	TTL          uint64
	AccountNonce uint64
}

// RLP returns a byte serialized representation
func (tx *NameRevokeTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// build id for the name
	nID, err := buildIDTag(IDTagName, tx.NameID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceRevokeTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		nID,
		tx.Fee,
		tx.TTL)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameRevokeTx) JSON() (string, error) {
	swaggerT := models.NameRevokeTx{
		AccountID: &tx.AccountID,
		Fee:       utils.BigInt(tx.Fee),
		NameID:    &tx.NameID,
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameRevokeTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *NameRevokeTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewNameRevokeTx is a constructor for a NameRevokeTx struct
func NewNameRevokeTx(accountID, name string, fee big.Int, ttl, accountNonce uint64) NameRevokeTx {
	return NameRevokeTx{accountID, name, fee, ttl, accountNonce}
}

// NameTransferTx represents a transaction that transfers ownership of one name to another account.
type NameTransferTx struct {
	AccountID    string
	NameID       string
	RecipientID  string
	Fee          big.Int
	TTL          uint64
	AccountNonce uint64
}

// RLP returns a byte serialized representation
func (tx *NameTransferTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}

	// build id for the recipient
	rID, err := buildIDTag(IDTagAccount, tx.RecipientID)
	if err != nil {
		return
	}

	// build id for the name
	nID, err := buildIDTag(IDTagName, tx.NameID)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceTransferTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		nID,
		rID,
		tx.Fee,
		tx.TTL,
	)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameTransferTx) JSON() (string, error) {
	swaggerT := models.NameTransferTx{
		AccountID:   &tx.AccountID,
		Fee:         utils.BigInt(tx.Fee),
		NameID:      &tx.NameID,
		Nonce:       tx.AccountNonce,
		RecipientID: &tx.RecipientID,
		TTL:         tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameTransferTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *NameTransferTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeStd(tx, txLenEstimated)
	return estimatedFee, nil
}

// NewNameTransferTx is a constructor for a NameTransferTx struct
func NewNameTransferTx(AccountID, NameID, RecipientID string, Fee big.Int, TTL, AccountNonce uint64) NameTransferTx {
	return NameTransferTx{AccountID, NameID, RecipientID, Fee, TTL, AccountNonce}
}

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

// RLP returns a byte serialized representation
func (tx *OracleRegisterTx) RLP() (rlpRawMsg []byte, err error) {
	// build id for the account
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
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

// RLP returns a byte serialized representation
func (tx *OracleExtendTx) RLP() (rlpRawMsg []byte, err error) {
	oID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		oID,
		tx.AccountNonce,
		tx.OracleTTLType,
		tx.OracleTTLValue,
		tx.Fee,
		tx.TTL)
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

// RLP returns a byte serialized representation
func (tx *OracleQueryTx) RLP() (rlpRawMsg []byte, err error) {
	accountID, err := buildIDTag(IDTagAccount, tx.SenderID)
	if err != nil {
		return
	}

	oracleID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
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

// RLP returns a byte serialized representation
func (tx *OracleRespondTx) RLP() (rlpRawMsg []byte, err error) {
	oID, err := buildIDTag(IDTagOracle, tx.OracleID)
	if err != nil {
		return
	}
	queryIDBytes, err := Decode(tx.QueryID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
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

// ContractCreateTx represents a transaction that creates a smart contract
type ContractCreateTx struct {
	OwnerID      string
	AccountNonce uint64
	Code         string
	VMVersion    uint16
	AbiVersion   uint16
	Deposit      big.Int
	Amount       big.Int
	Gas          big.Int
	GasPrice     big.Int
	Fee          big.Int
	TTL          uint64
	CallData     string
}

func encodeVMABI(VMVersion, ABIVersion uint16) []byte {
	vmBytes := big.NewInt(int64(VMVersion)).Bytes()
	abiBytes := big.NewInt(int64(ABIVersion)).Bytes()
	vmAbiBytes := []byte{}
	vmAbiBytes = append(vmAbiBytes, vmBytes...)
	vmAbiBytes = append(vmAbiBytes, leftPadByteSlice(2, abiBytes)...)
	return vmAbiBytes
}

// RLP returns a byte serialized representation
func (tx *ContractCreateTx) RLP() (rlpRawMsg []byte, err error) {
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

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagContractCreateTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		codeBinary,
		encodeVMABI(tx.VMVersion, tx.AbiVersion),
		tx.Fee,
		tx.TTL,
		tx.Deposit,
		tx.Amount,
		tx.Gas,
		tx.GasPrice,
		callDataBinary,
	)
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *ContractCreateTx) JSON() (string, error) {
	g := tx.Gas.Uint64()
	swaggerT := models.ContractCreateTx{
		OwnerID:    &tx.OwnerID,
		Nonce:      tx.AccountNonce,
		Code:       &tx.Code,
		VMVersion:  &tx.VMVersion,
		AbiVersion: &tx.AbiVersion,
		Deposit:    utils.BigInt(tx.Deposit),
		Amount:     utils.BigInt(tx.Amount),
		Gas:        &g,
		GasPrice:   utils.BigInt(tx.GasPrice),
		Fee:        utils.BigInt(tx.Fee),
		TTL:        tx.TTL,
		CallData:   &tx.CallData,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *ContractCreateTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *ContractCreateTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeContract(&tx.Gas, 5, txLenEstimated)
	return estimatedFee, nil
}

// ContractID returns the ct_ ID that this transaction would produce, which depends on the OwnerID and AccountNonce.
func (tx *ContractCreateTx) ContractID() (string, error) {
	return buildContractID(tx.OwnerID, tx.AccountNonce)
}

// NewContractCreateTx is a constructor for a ContractCreateTx struct
func NewContractCreateTx(OwnerID string, AccountNonce uint64, Code string, VMVersion, AbiVersion uint16, Deposit, Amount, Gas, GasPrice, Fee big.Int, TTL uint64, CallData string) ContractCreateTx {
	return ContractCreateTx{
		OwnerID:      OwnerID,
		AccountNonce: AccountNonce,
		Code:         Code,
		VMVersion:    VMVersion,
		AbiVersion:   AbiVersion,
		Deposit:      Deposit,
		Amount:       Amount,
		Gas:          Gas,
		GasPrice:     GasPrice,
		Fee:          Fee,
		TTL:          TTL,
		CallData:     CallData,
	}
}

// ContractCallTx represents calling an existing smart contract
// VMVersion is not included in RLP serialized representation (implied by contract type already)
type ContractCallTx struct {
	CallerID     string
	AccountNonce uint64
	ContractID   string
	Amount       big.Int
	Gas          big.Int
	GasPrice     big.Int
	AbiVersion   uint16
	CallData     string
	Fee          big.Int
	TTL          uint64
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *ContractCallTx) JSON() (string, error) {
	gas := tx.Gas.Uint64()
	swaggerT := models.ContractCallTx{
		CallerID:   &tx.CallerID,
		Nonce:      tx.AccountNonce,
		ContractID: &tx.ContractID,
		Amount:     utils.BigInt(tx.Amount),
		Gas:        &gas,
		GasPrice:   utils.BigInt(tx.GasPrice),
		AbiVersion: &tx.AbiVersion,
		CallData:   &tx.CallData,
		Fee:        utils.BigInt(tx.Fee),
		TTL:        tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// RLP returns a byte serialized representation
func (tx *ContractCallTx) RLP() (rlpRawMsg []byte, err error) {
	cID, err := buildIDTag(IDTagAccount, tx.CallerID)
	if err != nil {
		return
	}
	ctID, err := buildIDTag(IDTagContract, tx.ContractID)
	if err != nil {
		return
	}
	callDataBinary, err := Decode(tx.CallData)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagContractCallTransaction,
		rlpMessageVersion,
		cID,
		tx.AccountNonce,
		ctID,
		tx.AbiVersion,
		tx.Fee,
		tx.TTL,
		tx.Amount,
		tx.Gas,
		tx.GasPrice,
		callDataBinary,
	)
	return
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *ContractCallTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, &tx.Fee)
}

// FeeEstimate estimates the fee needed for the node to accept this transaction, assuming the fee is 8 bytes long when RLP serialized.
func (tx *ContractCallTx) FeeEstimate() (*big.Int, error) {
	txLenEstimated, err := tx.sizeEstimate()
	if err != nil {
		return new(big.Int), err
	}
	estimatedFee := calcFeeContract(&tx.Gas, 30, txLenEstimated)
	return estimatedFee, nil
}

// NewContractCallTx is a constructor for a ContractCallTx struct
func NewContractCallTx(CallerID string, AccountNonce uint64, ContractID string, Amount, Gas, GasPrice big.Int, AbiVersion uint16, CallData string, Fee big.Int, TTL uint64) ContractCallTx {
	return ContractCallTx{
		CallerID:     CallerID,
		AccountNonce: AccountNonce,
		ContractID:   ContractID,
		Amount:       Amount,
		Gas:          Gas,
		GasPrice:     GasPrice,
		AbiVersion:   AbiVersion,
		CallData:     CallData,
		Fee:          Fee,
		TTL:          TTL,
	}
}
