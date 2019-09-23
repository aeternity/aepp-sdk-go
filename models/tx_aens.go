package models

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/binary"
	"github.com/aeternity/aepp-sdk-go/v5/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v5/utils"
	rlp "github.com/randomshinichi/rlpae"
)

// NamePreclaimTx represents a transaction where one reserves a name on AENS without revealing it yet
type NamePreclaimTx struct {
	AccountID    string
	CommitmentID string
	Fee          *big.Int
	TTL          uint64
	AccountNonce uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *NamePreclaimTx) EncodeRLP(w io.Writer) (err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	// build id for the commitment
	cID, err := buildIDTag(IDTagCommitment, tx.CommitmentID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagNameServicePreclaimTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		cID,
		tx.Fee,
		tx.TTL)

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	return
}

type namePreclaimRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	CommitmentID      []uint8
	Fee               *big.Int
	TTL               uint64
}

func (n *namePreclaimRLP) ReadRLP(s *rlp.Stream) (aID, cID string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, n); err != nil {
		return
	}
	if _, aID, err = readIDTag(n.AccountID); err != nil {
		return
	}
	_, cID, err = readIDTag(n.CommitmentID)
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *NamePreclaimTx) DecodeRLP(s *rlp.Stream) (err error) {
	ntx := &namePreclaimRLP{}
	aID, cID, err := ntx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.AccountID = aID
	tx.CommitmentID = cID
	tx.Fee = ntx.Fee
	tx.TTL = ntx.TTL
	tx.AccountNonce = ntx.AccountNonce
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NamePreclaimTx) JSON() (string, error) {
	swaggerT := models.NamePreclaimTx{
		AccountID:    &tx.AccountID,
		CommitmentID: &tx.CommitmentID,
		Fee:          utils.BigInt(*tx.Fee),
		Nonce:        tx.AccountNonce,
		TTL:          tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NamePreclaimTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, tx.Fee)
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
func NewNamePreclaimTx(accountID, commitmentID string, fee *big.Int, ttl, accountNonce uint64) *NamePreclaimTx {
	return &NamePreclaimTx{accountID, commitmentID, fee, ttl, accountNonce}
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
// The revealed name is simply sent in plaintext in RLP, while in JSON representation
// it is base58 encoded.
type NameClaimTx struct {
	AccountID    string
	Name         string
	NameSalt     *big.Int
	Fee          *big.Int
	TTL          uint64
	AccountNonce uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *NameClaimTx) EncodeRLP(w io.Writer) (err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		tx.Name,
		tx.NameSalt,
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

type nameClaimRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	Name              string
	NameSalt          *big.Int
	Fee               *big.Int
	TTL               uint64
}

func (n *nameClaimRLP) ReadRLP(s *rlp.Stream) (aID string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, n); err != nil {
		return
	}
	if _, aID, err = readIDTag(n.AccountID); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *NameClaimTx) DecodeRLP(s *rlp.Stream) (err error) {
	ntx := &nameClaimRLP{}
	aID, err := ntx.ReadRLP(s)
	if err != nil {
		return
	}
	tx.AccountID = aID
	tx.Name = ntx.Name
	tx.NameSalt = ntx.NameSalt
	tx.Fee = ntx.Fee
	tx.TTL = ntx.TTL
	tx.AccountNonce = ntx.AccountNonce
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoints
func (tx *NameClaimTx) JSON() (string, error) {
	// When talking JSON to the node, the name should be 'API encoded'
	// (base58), not namehash-ed.
	nameAPIEncoded := binary.Encode(binary.PrefixName, []byte(tx.Name))
	swaggerT := models.NameClaimTx{
		AccountID: &tx.AccountID,
		Fee:       utils.BigInt(*tx.Fee),
		Name:      &nameAPIEncoded,
		NameSalt:  utils.BigInt(*tx.NameSalt),
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameClaimTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, tx.Fee)
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
func NewNameClaimTx(accountID, name string, nameSalt, fee *big.Int, ttl, accountNonce uint64) *NameClaimTx {
	return &NameClaimTx{accountID, name, nameSalt, fee, ttl, accountNonce}
}

// NamePointer is a go-native representation of swagger generated
// models.NamePointer.
type NamePointer struct {
	Key string
	ID  string
}

// Swagger generates the go-swagger representation of this model
func (np *NamePointer) Swagger() *models.NamePointer {
	return &models.NamePointer{
		Key: &np.Key,
		ID:  &np.ID,
	}
}

// EncodeRLP implements rlp.Encoder interface.
func (np *NamePointer) EncodeRLP(w io.Writer) (err error) {
	accountID, err := buildIDTag(IDTagAccount, np.ID)
	if err != nil {
		return
	}

	err = rlp.Encode(w, []interface{}{np.Key, accountID})
	if err != nil {
		return
	}
	return err
}

type namePointerRLP struct {
	Key       string
	AccountID []uint8
}

// DecodeRLP implements rlp.Decoder interface.
func (np *NamePointer) DecodeRLP(s *rlp.Stream) (err error) {
	var blob []byte
	npRLP := &namePointerRLP{}

	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, npRLP); err != nil {
		return
	}
	_, aID, err := readIDTag(npRLP.AccountID)
	if err != nil {
		return
	}

	np.Key = npRLP.Key
	np.ID = aID
	return err
}

// NewNamePointer is a constructor for NamePointer struct.
func NewNamePointer(key string, id string) *NamePointer {
	return &NamePointer{
		Key: key,
		ID:  id,
	}
}

// NameUpdateTx represents a transaction where one extends the lifetime of a reserved name on AENS
type NameUpdateTx struct {
	AccountID    string
	NameID       string
	Pointers     []*NamePointer
	NameTTL      uint64
	ClientTTL    uint64
	Fee          *big.Int
	TTL          uint64
	AccountNonce uint64
}

func reverse(input []*NamePointer) []*NamePointer {
	i := 0
	j := len(input) - 1
	reversedPointers := make([]*NamePointer, len(input))
	for i <= len(input)-1 {
		reversedPointers[i] = input[j]
		i++
		j--
	}
	return reversedPointers
}

// EncodeRLP implements rlp.Encoder
func (tx *NameUpdateTx) EncodeRLP(w io.Writer) (err error) {
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
	reversedPointers := reverse(tx.Pointers)

	// create the transaction
	rlpRawMsg, err := buildRLPMessage(
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

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return
}

type nameUpdateRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	Name              []uint8
	NameTTL           uint64
	Pointers          []*NamePointer
	ClientTTL         uint64
	Fee               *big.Int
	TTL               uint64
}

func (n *nameUpdateRLP) ReadRLP(s *rlp.Stream) (aID, name string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, n); err != nil {
		return
	}
	if _, aID, err = readIDTag(n.AccountID); err != nil {
		return
	}
	if _, name, err = readIDTag(n.Name); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *NameUpdateTx) DecodeRLP(s *rlp.Stream) (err error) {
	ntx := &nameUpdateRLP{}
	aID, name, err := ntx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.AccountID = aID
	tx.NameID = name
	tx.Pointers = reverse(ntx.Pointers)
	tx.NameTTL = ntx.NameTTL
	tx.ClientTTL = ntx.ClientTTL
	tx.Fee = ntx.Fee
	tx.TTL = ntx.TTL
	tx.AccountNonce = ntx.AccountNonce
	return nil
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameUpdateTx) JSON() (string, error) {
	swaggerNamePointers := []*models.NamePointer{}
	for _, np := range tx.Pointers {
		swaggerNamePointers = append(swaggerNamePointers, np.Swagger())
	}

	swaggerT := models.NameUpdateTx{
		AccountID: &tx.AccountID,
		ClientTTL: &tx.ClientTTL,
		Fee:       utils.BigInt(*tx.Fee),
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
	return calcSizeEstimate(tx, tx.Fee)
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
func NewNameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee *big.Int, ttl, accountNonce uint64) *NameUpdateTx {
	parsedPointers, err := buildPointers(pointers)
	if err != nil {
		panic(err)
	}
	return &NameUpdateTx{accountID, nameID, parsedPointers, nameTTL, clientTTL, fee, ttl, accountNonce}
}

// NameRevokeTx represents a transaction that revokes the name, i.e. has the same effect as waiting for the Name's TTL to expire.
type NameRevokeTx struct {
	AccountID    string
	NameID       string
	Fee          *big.Int
	TTL          uint64
	AccountNonce uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *NameRevokeTx) EncodeRLP(w io.Writer) (err error) {
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

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagNameServiceRevokeTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		nID,
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

type nameRevokeRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	Name              []uint8
	Fee               *big.Int
	TTL               uint64
}

func (n *nameRevokeRLP) ReadRLP(s *rlp.Stream) (aID, name string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, n); err != nil {
		return
	}
	if _, aID, err = readIDTag(n.AccountID); err != nil {
		return
	}
	if _, name, err = readIDTag(n.Name); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder interface.
func (tx *NameRevokeTx) DecodeRLP(s *rlp.Stream) (err error) {
	ntx := &nameRevokeRLP{}
	aID, name, err := ntx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.AccountID = aID
	tx.NameID = name
	tx.Fee = ntx.Fee
	tx.TTL = ntx.TTL
	tx.AccountNonce = ntx.AccountNonce
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameRevokeTx) JSON() (string, error) {
	swaggerT := models.NameRevokeTx{
		AccountID: &tx.AccountID,
		Fee:       utils.BigInt(*tx.Fee),
		NameID:    &tx.NameID,
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// sizeEstimate returns the size of the transaction when RLP serialized, assuming the Fee has a length of 8 bytes.
func (tx *NameRevokeTx) sizeEstimate() (int, error) {
	return calcSizeEstimate(tx, tx.Fee)
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
func NewNameRevokeTx(accountID, name string, fee *big.Int, ttl, accountNonce uint64) *NameRevokeTx {
	return &NameRevokeTx{accountID, name, fee, ttl, accountNonce}
}

// NameTransferTx represents a transaction that transfers ownership of one name to another account.
type NameTransferTx struct {
	AccountID    string
	NameID       string
	RecipientID  string
	Fee          *big.Int
	TTL          uint64
	AccountNonce uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *NameTransferTx) EncodeRLP(w io.Writer) (err error) {
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
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagNameServiceTransferTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		nID,
		rID,
		tx.Fee,
		tx.TTL,
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

type nameTransferRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	Name              []uint8
	RecipientID       []uint8
	Fee               *big.Int
	TTL               uint64
}

func (n *nameTransferRLP) ReadRLP(s *rlp.Stream) (aID, name, rID string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, n); err != nil {
		return
	}
	if _, aID, err = readIDTag(n.AccountID); err != nil {
		return
	}
	if _, name, err = readIDTag(n.Name); err != nil {
		return
	}
	if _, rID, err = readIDTag(n.RecipientID); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder interface.
func (tx *NameTransferTx) DecodeRLP(s *rlp.Stream) (err error) {
	ntx := &nameTransferRLP{}
	aID, name, rID, err := ntx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.AccountID = aID
	tx.NameID = name
	tx.RecipientID = rID
	tx.Fee = ntx.Fee
	tx.TTL = ntx.TTL
	tx.AccountNonce = ntx.AccountNonce
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *NameTransferTx) JSON() (string, error) {
	swaggerT := models.NameTransferTx{
		AccountID:   &tx.AccountID,
		Fee:         utils.BigInt(*tx.Fee),
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
	return calcSizeEstimate(tx, tx.Fee)
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
func NewNameTransferTx(AccountID, NameID, RecipientID string, Fee *big.Int, TTL, AccountNonce uint64) *NameTransferTx {
	return &NameTransferTx{AccountID, NameID, RecipientID, Fee, TTL, AccountNonce}
}
