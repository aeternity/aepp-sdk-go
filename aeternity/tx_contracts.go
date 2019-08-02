package aeternity

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func encodeVMABI(VMVersion, ABIVersion uint16) []byte {
	vmBytes := big.NewInt(int64(VMVersion)).Bytes()
	abiBytes := big.NewInt(int64(ABIVersion)).Bytes()
	vmAbiBytes := []byte{}
	vmAbiBytes = append(vmAbiBytes, vmBytes...)
	vmAbiBytes = append(vmAbiBytes, leftPadByteSlice(2, abiBytes)...)
	return vmAbiBytes
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

// EncodeRLP implements rlp.Encoder
func (tx *ContractCreateTx) EncodeRLP(w io.Writer) (err error) {
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
	callDataBinary, err := Decode(tx.CallData)
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
