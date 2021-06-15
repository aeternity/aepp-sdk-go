package transactions

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v8/binary"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	rlp "github.com/aeternity/rlp-go"
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
	GasLimit     *big.Int
	GasPrice     *big.Int
	Fee          *big.Int
	TTL          uint64
	CallData     string
}

// EncodeRLP implements rlp.Encoder
func (tx *GAAttachTx) EncodeRLP(w io.Writer) (err error) {
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
		ObjectTagGeneralizedAccountAttachTransaction,
		rlpMessageVersion,
		aID,
		tx.AccountNonce,
		codeBinary,
		tx.AuthFunc,
		encodeVMABI(tx.VMVersion, tx.AbiVersion),
		tx.Fee,
		tx.TTL,
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

type gaAttachRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AccountNonce      uint64
	CodeBinary        []byte
	AuthFunc          []byte
	VMABI             []byte
	Fee               *big.Int
	TTL               uint64
	GasLimit          *big.Int
	GasPrice          *big.Int
	CallDataBinary    []byte
}

func (g *gaAttachRLP) ReadRLP(s *rlp.Stream) (aID, code, calldata string, vmversion, abiversion uint16, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, g); err != nil {
		return
	}
	if _, aID, err = readIDTag(g.AccountID); err != nil {
		return
	}
	code = binary.Encode(binary.PrefixContractByteArray, g.CodeBinary)
	calldata = binary.Encode(binary.PrefixContractByteArray, g.CallDataBinary)
	vmversion, abiversion = decodeVMABI(g.VMABI)
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *GAAttachTx) DecodeRLP(s *rlp.Stream) (err error) {
	gtx := &gaAttachRLP{}
	aID, code, calldata, vmversion, abiversion, err := gtx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.OwnerID = aID
	tx.AccountNonce = gtx.AccountNonce
	tx.Code = code
	tx.AuthFunc = gtx.AuthFunc
	tx.VMVersion = vmversion
	tx.AbiVersion = abiversion
	tx.GasLimit = gtx.GasLimit
	tx.GasPrice = gtx.GasPrice
	tx.Fee = gtx.Fee
	tx.TTL = gtx.TTL
	tx.CallData = calldata
	return
}

// SetFee implements Transaction
func (tx *GAAttachTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *GAAttachTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *GAAttachTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Mul(config.Client.BaseGas, big.NewInt(5))
	gasComponent, err := normalGasComponent(tx, tx.GasLimit)
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewGAAttachTx creates a GAAttachTx
func NewGAAttachTx(ownerID string, code string, authFunc []byte, vmVersion uint16, abiVersion uint16, gasLimit *big.Int, gasPrice *big.Int, callData string, ttlnoncer TTLNoncer) (tx *GAAttachTx, err error) {
	ttl, _, accountNonce, err := ttlnoncer(ownerID, config.Client.TTL)
	if err != nil {
		return
	}

	tx = &GAAttachTx{
		OwnerID:      ownerID,
		AccountNonce: accountNonce,
		Code:         code,
		AuthFunc:     authFunc,
		VMVersion:    vmVersion,
		AbiVersion:   abiVersion,
		GasLimit:     gasLimit,
		GasPrice:     gasPrice,
		Fee:          config.Client.Fee,
		TTL:          ttl,
		CallData:     callData,
	}
	CalculateFee(tx)
	return
}

// GAMetaTx wraps a normal Tx (that is not a GAAttachTx) that will only be
// executed if the contract located at GaID returns true
type GAMetaTx struct {
	AccountID  string
	AuthData   string
	AbiVersion uint16
	GasLimit   *big.Int
	GasPrice   *big.Int
	Fee        *big.Int
	Tx         *SignedTx
}

// EncodeRLP implements rlp.Encoder
func (tx *GAMetaTx) EncodeRLP(w io.Writer) (err error) {
	aID, err := buildIDTag(IDTagAccount, tx.AccountID)
	if err != nil {
		return
	}
	authDataBinary, err := binary.Decode(tx.AuthData)
	if err != nil {
		return
	}
	txRLP, err := rlp.EncodeToBytes(tx.Tx)
	if err != nil {
		return
	}

	rlpRawMsg, err := buildRLPMessage(
		ObjectTagGeneralizedAccountMetaTransaction,
		rlpMessageVersion2,
		aID,
		authDataBinary,
		tx.AbiVersion,
		tx.Fee,
		tx.GasLimit,
		tx.GasPrice,
		txRLP,
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

type gaMetaRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	AccountID         []uint8
	AuthDataBinary    []byte
	AbiVersion        uint16
	Fee               *big.Int
	GasLimit          *big.Int
	GasPrice          *big.Int
	WrappedTx         []byte
}

func (g *gaMetaRLP) ReadRLP(s *rlp.Stream) (aID, authdata string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, g); err != nil {
		return
	}
	if _, aID, err = readIDTag(g.AccountID); err != nil {
		return
	}

	authdata = binary.Encode(binary.PrefixContractByteArray, g.AuthDataBinary)
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *GAMetaTx) DecodeRLP(s *rlp.Stream) (err error) {
	gtx := &gaMetaRLP{}
	aID, authdata, err := gtx.ReadRLP(s)
	if err != nil {
		return
	}
	wtx, err := DeserializeTx(gtx.WrappedTx)
	if err != nil {
		return
	}

	tx.AccountID = aID
	tx.AuthData = authdata
	tx.AbiVersion = gtx.AbiVersion
	tx.GasLimit = gtx.GasLimit
	tx.GasPrice = gtx.GasPrice
	tx.Fee = gtx.Fee
	tx.Tx = wtx.(*SignedTx)
	return
}

// SetFee implements Transaction
func (tx *GAMetaTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *GAMetaTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
// TODO Recursive tx gas calculation not implemented
func (tx *GAMetaTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Mul(config.Client.BaseGas, big.NewInt(5))
	gasComponent, err := normalGasComponent(tx, tx.GasLimit)
	if err != nil {
		return
	}
	g = baseGas.Add(baseGas, gasComponent)
	return
}

// NewGAMetaTx creates a GAMetaTx
func NewGAMetaTx(accountID string, authData string, abiVersion uint16, gasLimit *big.Int, gasPrice *big.Int, wrappedTx Transaction, ttlnoncer TTLNoncer) (tx *GAMetaTx, err error) {
	tx = &GAMetaTx{
		AccountID:  accountID,
		AuthData:   authData,
		AbiVersion: abiVersion,
		GasLimit:   gasLimit,
		GasPrice:   gasPrice,
		Fee:        config.Client.Fee,
		Tx: &SignedTx{
			Signatures: [][]byte{},
			Tx:         wrappedTx,
		},
	}
	CalculateFee(tx)
	return
}
