package transactions

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/aeternity/aepp-sdk-go/v9/binary"
	"github.com/aeternity/aepp-sdk-go/v9/config"
	"github.com/aeternity/aepp-sdk-go/v9/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v9/utils"
	rlp "github.com/aeternity/rlp-go"
)

// NameID computes the nm_ string of a given AENS name.
func NameID(name string) (nm string, err error) {
	s, err := binary.Blake2bHash([]byte(name))
	if err != nil {
		return
	}
	return binary.Encode(binary.PrefixName, s), nil
}

// generateCommitmentID gives a commitment ID 'cm_...' given a particular AENS
// name. It is split into the deterministic part computeCommitmentID(), which
// can be tested, and the part incorporating random salt generateCommitmentID()
//
// since the salt is a uint256, which Erlang handles well, but Go has nothing
// similar to it, it is imperative that the salt be kept as a bytearray unless
// you really have to convert it into an integer. Which you usually don't,
// because it's a salt.
func generateCommitmentID(name string) (ch string, salt *big.Int, err error) {
	// Generate 32 random bytes for a salt
	saltBytes := make([]byte, 32)
	_, err = rand.Read(saltBytes)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return
	}

	ch, err = computeCommitmentID(name, saltBytes)

	salt = new(big.Int)
	salt.SetBytes(saltBytes)

	return ch, salt, err
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !strconv.IsPrint(r) {
			return false
		}
	}
	return true
}

func computeCommitmentID(name string, salt []byte) (ch string, err error) {
	var nh = []byte{}
	if strings.HasSuffix(name, ".test") {
		nh = append(Namehash(name), salt...)

	} else {
		// Since UTF-8 ~ ASCII, just use the string directly. QuoteToASCII
		// includes an extra byte at the start and end of the string, messing up
		// the hashing process.
		if !isPrintable(name) {
			return "", fmt.Errorf("the name %s must contain only printable characters", name)
		}

		nh = append([]byte(name), salt...)
	}
	nh, err = binary.Blake2bHash(nh)
	if err != nil {
		return
	}
	ch = binary.Encode(binary.PrefixCommitment, nh)
	return
}

// Namehash calculate the Namehash of a string. Names within aeternity are
// generally referred to only by their namehashes.
//
// The implementation is the same as ENS EIP-137
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-137.md#namehash-algorithm
// but using Blake2b.
func Namehash(name string) []byte {
	buf := make([]byte, 32)
	for _, s := range strings.Split(name, ".") {
		sh, _ := binary.Blake2bHash([]byte(s))
		buf, _ = binary.Blake2bHash(append(buf, sh...))
	}
	return buf
}

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
	fee := utils.BigInt(*tx.Fee)
	swaggerT := models.NamePreclaimTx{
		AccountID:    &tx.AccountID,
		CommitmentID: &tx.CommitmentID,
		Fee:          &fee,
		Nonce:        tx.AccountNonce,
		TTL:          tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *NamePreclaimTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *NamePreclaimTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *NamePreclaimTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewNamePreclaimTx is a constructor for a NamePreclaimTx struct
func NewNamePreclaimTx(accountID, name string, ttlnoncer TTLNoncer) (tx *NamePreclaimTx, nameSalt *big.Int, err error) {
	ttl, _, accountNonce, err := ttlnoncer(accountID, config.Client.TTL)
	if err != nil {
		return
	}
	// calculate the commitment and get the preclaim salt since the salt is 32
	// bytes long, you must use a big.Int to convert it into an integer
	cm, nameSalt, err := generateCommitmentID(name)
	if err != nil {
		return
	}

	tx = &NamePreclaimTx{accountID, cm, config.Client.Fee, ttl, accountNonce}
	CalculateFee(tx)
	return
}

// NameClaimTx represents a transaction where one claims a previously reserved name on AENS
// The revealed name is simply sent in plaintext in RLP, while in JSON representation
// it is base58 encoded.
type NameClaimTx struct {
	AccountID    string
	Name         string
	NameSalt     *big.Int
	NameFee      *big.Int
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
		2,
		aID,
		tx.AccountNonce,
		tx.Name,
		tx.NameSalt,
		tx.NameFee,
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
	NameFee           *big.Int
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
	tx.NameFee = ntx.NameFee
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
	fee := utils.BigInt(*tx.Fee)
	nameSalt := utils.BigInt(*tx.NameSalt)
	swaggerT := models.NameClaimTx{
		AccountID: &tx.AccountID,
		Fee:       &fee,
		Name:      &nameAPIEncoded,
		NameSalt:  &nameSalt,
		NameFee:   utils.BigInt(*tx.NameFee),
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *NameClaimTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *NameClaimTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *NameClaimTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewNameClaimTx is a constructor for a NameClaimTx struct
func NewNameClaimTx(accountID, name string, nameSalt, nameFee *big.Int, ttlnoncer TTLNoncer) (tx *NameClaimTx, err error) {
	ttl, _, accountNonce, err := ttlnoncer(accountID, config.Client.TTL)
	if err != nil {
		return
	}

	tx = &NameClaimTx{accountID, name, nameSalt, nameFee, config.Client.Fee, ttl, accountNonce}
	CalculateFee(tx)
	return
}

// CalculateMinNameFee returns the starting bid price for a name on AENS. The
// name argument should include its TLD, e.g. "fdsa.test".
func CalculateMinNameFee(name string) (minNameFee *big.Int) {
	n := strings.Split(name, ".") // n = ['fdsa', '.test']
	minNameFee = new(big.Int)
	l := len(n[0])
	nf := config.NameAuctionFee(l)
	minNameFee.SetUint64(nf)
	return
}

// NamePointer is a go-native representation of swagger generated
// models.NamePointer.
type NamePointer struct {
	Key     string
	Pointer string
}

// Swagger generates the go-swagger representation of this model
func (np *NamePointer) Swagger() *models.NamePointer {
	return &models.NamePointer{
		Key: &np.Key,
		ID:  &np.Pointer,
	}
}

// EncodeRLP implements rlp.Encoder interface.
func (np *NamePointer) EncodeRLP(w io.Writer) (err error) {
	var idTag uint8
	// figure out ID tag for RLP serialization from NamePointer.Key
	switch np.Key {
	case "account_pubkey":
		idTag = IDTagAccount
	case "oracle_pubkey":
		idTag = IDTagOracle
	case "contract_pubkey":
		idTag = IDTagContract
	case "channel":
		idTag = IDTagChannel
	// if NamePointer.Key is custom, encode it as a channel for now until AENS
	// supports generic byte blobs
	default:
		idTag = IDTagChannel
	}

	ID, err := buildIDTag(idTag, np.Pointer)
	if err != nil {
		return
	}

	err = rlp.Encode(w, []interface{}{np.Key, ID})
	if err != nil {
		return
	}
	return err
}

type namePointerRLP struct {
	Key     string
	Pointer []uint8
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
	_, ID, err := readIDTag(npRLP.Pointer)
	if err != nil {
		return
	}

	np.Key = npRLP.Key
	np.Pointer = ID
	return err
}

// Validate checks that the pointer's type matches its declared type in the Key,
// and that its length is <32 bytes.
func (np *NamePointer) Validate() (err error) {
	var typeValidation error
	switch np.Key {
	case "account_pubkey":
		typeValidation = np.validateAccountPubkey()
	case "oracle_pubkey":
		typeValidation = np.validateOraclePubkey()
	case "contract_pubkey":
		typeValidation = np.validateContractPubkey()
	case "channel":
		typeValidation = np.validateChannel()
	default:
		typeValidation = nil
	}
	return typeValidation
}

func (np *NamePointer) validateAccountPubkey() error {
	if !strings.HasPrefix(np.Pointer, string(binary.PrefixAccountPubkey)) {
		return fmt.Errorf("if the Key is \"account_pubkey\", the Pointer must start with %s", binary.PrefixAccountPubkey)
	}
	return nil
}
func (np *NamePointer) validateOraclePubkey() error {
	if !strings.HasPrefix(np.Pointer, string(binary.PrefixOraclePubkey)) {
		return fmt.Errorf("if the Key is \"oracle_pubkey\", the Pointer must start with %s", binary.PrefixOraclePubkey)
	}
	return nil
}

func (np *NamePointer) validateContractPubkey() error {
	if !strings.HasPrefix(np.Pointer, string(binary.PrefixContractPubkey)) {
		return fmt.Errorf("if the Key is \"contract_pubkey\", the Pointer must start with %s", binary.PrefixContractPubkey)
	}
	return nil
}

func (np *NamePointer) validateChannel() error {
	if !strings.HasPrefix(np.Pointer, string(binary.PrefixChannel)) {
		return fmt.Errorf("if the Key is \"channel\", the Pointer must start with %s", binary.PrefixChannel)
	}
	return nil
}

// NewNamePointer is a constructor for NamePointer struct.
func NewNamePointer(key string, pointer string) (n *NamePointer, err error) {
	n = &NamePointer{
		Key:     key,
		Pointer: pointer,
	}
	err = n.Validate()
	if err != nil {
		return nil, err
	}
	return
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

	fee := utils.BigInt(*tx.Fee)
	swaggerT := models.NameUpdateTx{
		AccountID: &tx.AccountID,
		ClientTTL: &tx.ClientTTL,
		Fee:       &fee,
		NameID:    &tx.NameID,
		NameTTL:   &tx.NameTTL,
		Nonce:     tx.AccountNonce,
		Pointers:  swaggerNamePointers,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *NameUpdateTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *NameUpdateTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *NameUpdateTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewNameUpdateTx is a constructor for a NameUpdateTx struct
func NewNameUpdateTx(accountID, name string, pointers []*NamePointer, clientTTL uint64, ttlnoncer TTLNoncer) (tx *NameUpdateTx, err error) {
	ttl, height, accountNonce, err := ttlnoncer(accountID, config.Client.TTL)
	if err != nil {
		return
	}
	nameTTL := height + config.Client.Names.NameTTL

	nm, err := NameID(name)
	if err != nil {
		return
	}

	tx = &NameUpdateTx{accountID, nm, pointers, nameTTL, clientTTL, config.Client.Fee, ttl, accountNonce}
	CalculateFee(tx)
	return
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
	fee := utils.BigInt(*tx.Fee)
	swaggerT := models.NameRevokeTx{
		AccountID: &tx.AccountID,
		Fee:       &fee,
		NameID:    &tx.NameID,
		Nonce:     tx.AccountNonce,
		TTL:       tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *NameRevokeTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *NameRevokeTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *NameRevokeTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewNameRevokeTx is a constructor for a NameRevokeTx struct
func NewNameRevokeTx(accountID, name string, ttlnoncer TTLNoncer) (tx *NameRevokeTx, err error) {
	ttl, _, accountNonce, err := ttlnoncer(accountID, config.Client.TTL)
	if err != nil {
		return
	}

	nm, err := NameID(name)
	if err != nil {
		return
	}

	tx = &NameRevokeTx{accountID, nm, config.Client.Fee, ttl, accountNonce}
	CalculateFee(tx)
	return
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
	fee := utils.BigInt(*tx.Fee)
	swaggerT := models.NameTransferTx{
		AccountID:   &tx.AccountID,
		Fee:         &fee,
		NameID:      &tx.NameID,
		Nonce:       tx.AccountNonce,
		RecipientID: &tx.RecipientID,
		TTL:         tx.TTL,
	}

	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *NameTransferTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *NameTransferTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *NameTransferTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewNameTransferTx is a constructor for a NameTransferTx struct
func NewNameTransferTx(accountID, name, recipientID string, ttlnoncer TTLNoncer) (tx *NameTransferTx, err error) {
	ttl, _, accountNonce, err := ttlnoncer(accountID, config.Client.TTL)
	if err != nil {
		return
	}

	nm, err := NameID(name)
	if err != nil {
		return
	}

	tx = &NameTransferTx{accountID, nm, recipientID, config.Client.Fee, ttl, accountNonce}
	CalculateFee(tx)
	return
}
