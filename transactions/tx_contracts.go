package transactions

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v6/binary"
	"github.com/aeternity/aepp-sdk-go/v6/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v6/utils"
	rlp "github.com/randomshinichi/rlpae"
)

func encodeVMABI(VMVersion, ABIVersion uint16) []byte {
	vmBytes := big.NewInt(int64(VMVersion)).Bytes()
	abiBytes := big.NewInt(int64(ABIVersion)).Bytes()
	vmAbiBytes := []byte{}
	vmAbiBytes = append(vmAbiBytes, vmBytes...)
	vmAbiBytes = append(vmAbiBytes, leftPadByteSlice(2, abiBytes)...)
	return vmAbiBytes
}

func decodeVMABI(vmabi []byte) (VMVersion, ABIVersion uint16) {
	v := new(big.Int)
	a := new(big.Int)
	var vmPortion, abiPortion []byte
	l := len(vmabi)
	if (l % 2) == 0 {
		vmPortion = vmabi[0:2]
		abiPortion = vmabi[2:]
	} else {
		vmPortion = []byte{vmabi[0]}
		abiPortion = []byte{vmabi[2]}
	}
	v.SetBytes(vmPortion)
	a.SetBytes(abiPortion)
	return uint16(v.Uint64()), uint16(a.Uint64())
}

// ContractCreateTx represents a transaction that creates a smart contract
type ContractCreateTx struct {
	OwnerID      string
	AccountNonce uint64
	Code         string
	VMVersion    uint16
	AbiVersion   uint16
	Deposit      *big.Int
	Amount       *big.Int
	GasLimit     *big.Int
	GasPrice     *big.Int
	Fee          *big.Int
	TTL          uint64
	CallData     string
}

// EncodeRLP implements rlp.Encoder
func (tx *ContractCreateTx) EncodeRLP(w io.Writer) (err error) {
	aID, err := buildIDTag(IDTagAccount, tx.OwnerID)
	if err != nil {
		return
	}
	codeBinary, err := binary.Decode(tx.Code)
	if err != nil {
		return
	}
	callDataBinary, err := binary.Decode(tx.CallData)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
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
		tx.GasLimit,
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

type contractCreateRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	CodeBinary        []byte
	VMABI             []byte
	Fee               *big.Int
	TTL               uint64
	Deposit           *big.Int
	Amount            *big.Int
	GasLimit          *big.Int
	GasPrice          *big.Int
	CallDataBinary    []byte
}

func (c *contractCreateRLP) ReadRLP(s *rlp.Stream) (aID, code, calldata string, vmversion, abiversion uint16, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, c); err != nil {
		return
	}
	if _, aID, err = readIDTag(c.AccountID); err != nil {
		return
	}

	code = binary.Encode(binary.PrefixContractByteArray, c.CodeBinary)
	calldata = binary.Encode(binary.PrefixContractByteArray, c.CallDataBinary)
	vmversion, abiversion = decodeVMABI(c.VMABI)
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *ContractCreateTx) DecodeRLP(s *rlp.Stream) (err error) {
	ctx := &contractCreateRLP{}
	aID, code, calldata, vmversion, abiversion, err := ctx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.OwnerID = aID
	tx.AccountNonce = ctx.AccountNonce
	tx.Code = code
	tx.VMVersion = vmversion
	tx.AbiVersion = abiversion
	tx.Deposit = ctx.Deposit
	tx.Amount = ctx.Amount
	tx.GasLimit = ctx.GasLimit
	tx.GasPrice = ctx.GasPrice
	tx.Fee = ctx.Fee
	tx.TTL = ctx.TTL
	tx.CallData = calldata
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *ContractCreateTx) JSON() (string, error) {
	g := tx.GasLimit.Uint64()
	swaggerT := models.ContractCreateTx{
		OwnerID:    &tx.OwnerID,
		Nonce:      tx.AccountNonce,
		Code:       &tx.Code,
		VMVersion:  &tx.VMVersion,
		AbiVersion: &tx.AbiVersion,
		Deposit:    utils.BigInt(*tx.Deposit),
		Amount:     utils.BigInt(*tx.Amount),
		Gas:        &g,
		GasPrice:   utils.BigInt(*tx.GasPrice),
		Fee:        utils.BigInt(*tx.Fee),
		TTL:        tx.TTL,
		CallData:   &tx.CallData,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// ContractID returns the ct_ ID that this transaction would produce, which depends on the OwnerID and AccountNonce.
func (tx *ContractCreateTx) ContractID() (string, error) {
	return buildContractID(tx.OwnerID, tx.AccountNonce)
}

// SetFee implements TransactionFeeCalculable
func (tx *ContractCreateTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements TransactionFeeCalculable
func (tx *ContractCreateTx) GetFee() *big.Int {
	return tx.Fee
}

// GetGasLimit implements TransactionFeeCalculable
func (tx *ContractCreateTx) GetGasLimit() *big.Int {
	return tx.GasLimit
}

// NewContractCreateTx is a constructor for a ContractCreateTx struct
func NewContractCreateTx(OwnerID string, AccountNonce uint64, Code string, VMVersion, AbiVersion uint16, Deposit, Amount, GasLimit, GasPrice, Fee *big.Int, TTL uint64, CallData string) *ContractCreateTx {
	return &ContractCreateTx{
		OwnerID:      OwnerID,
		AccountNonce: AccountNonce,
		Code:         Code,
		VMVersion:    VMVersion,
		AbiVersion:   AbiVersion,
		Deposit:      Deposit,
		Amount:       Amount,
		GasLimit:     GasLimit,
		GasPrice:     GasPrice,
		Fee:          Fee,
		TTL:          TTL,
		CallData:     CallData,
	}
}

// ContractCallTx represents calling an existing smart contract
type ContractCallTx struct {
	CallerID     string
	AccountNonce uint64
	ContractID   string
	Amount       *big.Int
	GasLimit     *big.Int
	GasPrice     *big.Int
	AbiVersion   uint16
	CallData     string
	Fee          *big.Int
	TTL          uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *ContractCallTx) EncodeRLP(w io.Writer) (err error) {
	cID, err := buildIDTag(IDTagAccount, tx.CallerID)
	if err != nil {
		return
	}
	ctID, err := buildIDTag(IDTagContract, tx.ContractID)
	if err != nil {
		return
	}
	callDataBinary, err := binary.Decode(tx.CallData)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagContractCallTransaction,
		rlpMessageVersion,
		cID,
		tx.AccountNonce,
		ctID,
		tx.AbiVersion,
		tx.Fee,
		tx.TTL,
		tx.Amount,
		tx.GasLimit,
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

type contractCallRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	CallerID          []uint8
	AccountNonce      uint64
	ContractID        []uint8
	AbiVersion        uint16
	Fee               *big.Int
	TTL               uint64
	Amount            *big.Int
	GasLimit          *big.Int
	GasPrice          *big.Int
	CallDataBinary    []byte
}

func (c *contractCallRLP) ReadRLP(s *rlp.Stream) (cID, ctID, calldata string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, c); err != nil {
		return
	}
	if _, cID, err = readIDTag(c.CallerID); err != nil {
		return
	}
	if _, ctID, err = readIDTag(c.ContractID); err != nil {
		return
	}

	calldata = binary.Encode(binary.PrefixContractByteArray, c.CallDataBinary)
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *ContractCallTx) DecodeRLP(s *rlp.Stream) (err error) {
	ctx := &contractCallRLP{}
	cID, ctID, calldata, err := ctx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.CallerID = cID
	tx.AccountNonce = ctx.AccountNonce
	tx.ContractID = ctID
	tx.Amount = ctx.Amount
	tx.GasLimit = ctx.GasLimit
	tx.GasPrice = ctx.GasPrice
	tx.AbiVersion = ctx.AbiVersion
	tx.CallData = calldata
	tx.Fee = ctx.Fee
	tx.TTL = ctx.TTL
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *ContractCallTx) JSON() (string, error) {
	gas := tx.GasLimit.Uint64()
	swaggerT := models.ContractCallTx{
		CallerID:   &tx.CallerID,
		Nonce:      tx.AccountNonce,
		ContractID: &tx.ContractID,
		Amount:     utils.BigInt(*tx.Amount),
		Gas:        &gas,
		GasPrice:   utils.BigInt(*tx.GasPrice),
		AbiVersion: &tx.AbiVersion,
		CallData:   &tx.CallData,
		Fee:        utils.BigInt(*tx.Fee),
		TTL:        tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements TransactionFeeCalculable
func (tx *ContractCallTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements TransactionFeeCalculable
func (tx *ContractCallTx) GetFee() *big.Int {
	return tx.Fee
}

// GetGasLimit implements TransactionFeeCalculable
func (tx *ContractCallTx) GetGasLimit() *big.Int {
	return tx.GasLimit
}

// NewContractCallTx is a constructor for a ContractCallTx struct
func NewContractCallTx(CallerID string, AccountNonce uint64, ContractID string, Amount, GasLimit, GasPrice *big.Int, AbiVersion uint16, CallData string, Fee *big.Int, TTL uint64) *ContractCallTx {
	return &ContractCallTx{
		CallerID:     CallerID,
		AccountNonce: AccountNonce,
		ContractID:   ContractID,
		Amount:       Amount,
		GasLimit:     GasLimit,
		GasPrice:     GasPrice,
		AbiVersion:   AbiVersion,
		CallData:     CallData,
		Fee:          Fee,
		TTL:          TTL,
	}
}
