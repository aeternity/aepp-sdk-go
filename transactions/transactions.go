package transactions

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v6/binary"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/naet"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/utils"
	rlp "github.com/randomshinichi/rlpae"
)

// TransactionTypes is a map between the ObjectTags defined above and the
// corresponding Tx struct. It is wrapped by a function to guarantee you cannot
// modify this map, because Golang does not have const maps.
func TransactionTypes() map[uint]Transaction {
	return map[uint]Transaction{
		ObjectTagSignedTransaction:                   &SignedTx{},
		ObjectTagSpendTransaction:                    &SpendTx{},
		ObjectTagNameServiceClaimTransaction:         &NameClaimTx{},
		ObjectTagNameServicePreclaimTransaction:      &NamePreclaimTx{},
		ObjectTagNameServiceUpdateTransaction:        &NameUpdateTx{},
		ObjectTagNameServiceRevokeTransaction:        &NameRevokeTx{},
		ObjectTagNameServiceTransferTransaction:      &NameTransferTx{},
		ObjectTagOracleRegisterTransaction:           &OracleRegisterTx{},
		ObjectTagOracleQueryTransaction:              &OracleQueryTx{},
		ObjectTagOracleResponseTransaction:           &OracleRespondTx{},
		ObjectTagOracleExtendTransaction:             &OracleExtendTx{},
		ObjectTagContractCreateTransaction:           &ContractCreateTx{},
		ObjectTagContractCallTransaction:             &ContractCallTx{},
		ObjectTagGeneralizedAccountAttachTransaction: &GAAttachTx{},
		ObjectTagGeneralizedAccountMetaTransaction:   &GAMetaTx{},
	}
}

// RLP message version used in RLP serialization
const (
	rlpMessageVersion uint = 1
)

// Address-like bytearrays are converted in to an ID (uint8 bytearray) for RLP
// serialization. ID Tags differentiate between them.
// https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
const (
	IDTagAccount    uint8 = 1
	IDTagName       uint8 = 2
	IDTagCommitment uint8 = 3
	IDTagOracle     uint8 = 4
	IDTagContract   uint8 = 5
	IDTagChannel    uint8 = 6
)

// Object tags are used to differentiate between different types of bytearrays
// in RLP serialization. see
// https://github.com/aeternity/protocol/blob/master/serializations.md#binary-serialization
const (
	ObjectTagAccount                             uint = 10
	ObjectTagSignedTransaction                   uint = 11
	ObjectTagSpendTransaction                    uint = 12
	ObjectTagOracle                              uint = 20
	ObjectTagOracleQuery                         uint = 21
	ObjectTagOracleRegisterTransaction           uint = 22
	ObjectTagOracleQueryTransaction              uint = 23
	ObjectTagOracleResponseTransaction           uint = 24
	ObjectTagOracleExtendTransaction             uint = 25
	ObjectTagNameServiceName                     uint = 30
	ObjectTagNameServiceCommitment               uint = 31
	ObjectTagNameServiceClaimTransaction         uint = 32
	ObjectTagNameServicePreclaimTransaction      uint = 33
	ObjectTagNameServiceUpdateTransaction        uint = 34
	ObjectTagNameServiceRevokeTransaction        uint = 35
	ObjectTagNameServiceTransferTransaction      uint = 36
	ObjectTagContract                            uint = 40
	ObjectTagContractCall                        uint = 41
	ObjectTagContractCreateTransaction           uint = 42
	ObjectTagContractCallTransaction             uint = 43
	ObjectTagChannelCreateTransaction            uint = 50
	ObjectTagChannelDepositTransaction           uint = 51
	ObjectTagChannelWithdrawTransaction          uint = 52
	ObjectTagChannelForceProgressTransaction     uint = 521
	ObjectTagChannelCloseMutualTransaction       uint = 53
	ObjectTagChannelCloseSoloTransaction         uint = 54
	ObjectTagChannelSlashTransaction             uint = 55
	ObjectTagChannelSettleTransaction            uint = 56
	ObjectTagChannelOffChainTransaction          uint = 57
	ObjectTagChannelOffChainUpdateTransfer       uint = 570
	ObjectTagChannelOffChainUpdateDeposit        uint = 571
	ObjectTagChannelOffChainUpdateWithdrawal     uint = 572
	ObjectTagChannelOffChainUpdateCreateContract uint = 573
	ObjectTagChannelOffChainUpdateCallContract   uint = 574
	ObjectTagChannel                             uint = 58
	ObjectTagChannelSnapshotTransaction          uint = 59
	ObjectTagPoi                                 uint = 60
	ObjectTagGeneralizedAccountAttachTransaction uint = 80
	ObjectTagGeneralizedAccountMetaTransaction   uint = 81
	ObjectTagMicroBody                           uint = 101
	ObjectTagLightMicroBlock                     uint = 102
)

func leftPadByteSlice(length int, data []byte) []byte {
	dataLen := len(data)
	t := make([]byte, length-dataLen)
	paddedSlice := append(t, data...)
	return paddedSlice
}

func buildOracleQueryID(sender string, senderNonce uint64, recipient string) (id string, err error) {
	queryIDBin := []byte{}
	senderBin, err := binary.Decode(sender)
	if err != nil {
		return
	}
	queryIDBin = append(queryIDBin, senderBin...)

	senderNonceBytes := utils.NewIntFromUint64(senderNonce).Bytes()
	senderNonceBytesPadded := leftPadByteSlice(32, senderNonceBytes)
	queryIDBin = append(queryIDBin, senderNonceBytesPadded...)

	recipientBin, err := binary.Decode(recipient)
	if err != nil {
		return
	}
	queryIDBin = append(queryIDBin, recipientBin...)

	hashedQueryID, err := binary.Blake2bHash(queryIDBin)
	if err != nil {
		return
	}
	id = binary.Encode(binary.PrefixOracleQueryID, hashedQueryID)
	return
}

func buildContractID(sender string, senderNonce uint64) (ctID string, err error) {
	senderBin, err := binary.Decode(sender)
	if err != nil {
		return ctID, err
	}

	l := big.Int{}
	l.SetUint64(senderNonce)

	ctIDUnhashed := append(senderBin, l.Bytes()...)
	ctIDHashed, err := binary.Blake2bHash(ctIDUnhashed)
	if err != nil {
		return ctID, err
	}

	ctID = binary.Encode(binary.PrefixContractPubkey, ctIDHashed)
	return ctID, err
}

// Transaction is used to indicate a transaction of any type.
type Transaction interface {
	rlp.Encoder
}

// TransactionFeeCalculable is a set of methods that simplify calculating the tx
// fee. In particular, SetFee and GetFee let the code increase the fee further
// in case the newer, calculated fee ends up increasing the size of the
// transaction (and thus necessitates an even larger fee)
type TransactionFeeCalculable interface {
	Transaction
	SetFee(*big.Int)
	GetFee() *big.Int
	GetGasLimit() *big.Int
}

// calculateSignature calculates the signature of the SignedTx.Tx. Although it does not use
// the SignedTx itself, it takes a SignedTx as an argument because if it took a
// rlp.Encoder as an interface, one might expect the signature to be of the
// SignedTx itself, which won't work.
func calculateSignature(kp *account.Account, tx *SignedTx, networkID string) (signature []byte, err error) {
	txRaw, err := rlp.EncodeToBytes(tx.Tx)
	if err != nil {
		return []byte{}, err
	}
	// add the network_id to the transaction
	msg := append([]byte(networkID), txRaw...)
	// sign the transaction
	signature = kp.Sign(msg)
	return
}

// Hash calculates the hash of a SignedTx. It is intended to be used after
// SignedTx.Signatures has been filled out.
func Hash(tx *SignedTx) (txhash string, err error) {
	rlpTxRaw, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}
	rlpTxHashRaw, err := binary.Blake2bHash(rlpTxRaw)
	if err != nil {
		return "", err
	}

	txhash = binary.Encode(binary.PrefixTransactionHash, rlpTxHashRaw)
	return txhash, nil
}

// SignHashTx wraps a *Tx struct in a SignedTx, then returns its signature and
// hash.
func SignHashTx(kp *account.Account, tx Transaction, networkID string) (signedTx *SignedTx, txhash, signature string, err error) {
	signedTx = NewSignedTx([][]byte{}, tx)
	var signatureBytes []byte

	if _, ok := tx.(*GAMetaTx); !ok {
		signatureBytes, err = calculateSignature(kp, signedTx, networkID)
		if err != nil {
			return
		}
		signedTx.Signatures = append(signedTx.Signatures, signatureBytes)
		signature = binary.Encode(binary.PrefixSignature, signatureBytes)

	}

	txhash, err = Hash(signedTx)
	if err != nil {
		return
	}
	return signedTx, txhash, signature, nil
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
		switch binary.GetHashPrefix(p) {
		case binary.PrefixAccountPubkey:
			// pID, err := buildIDTag(IDTagAccount, p)
			key := "account_pubkey"
			ptrs[i] = NewNamePointer(key, p)
			if err != nil {
				break
			}
		case binary.PrefixOraclePubkey:
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

func calcRlpLen(o interface{}) (size int64, err error) {
	rlpEncoded, err := rlp.EncodeToBytes(o)
	if err != nil {
		return
	}
	size = int64(len(rlpEncoded))
	return
}

func baseGasForTxType(tx Transaction) (baseGas *big.Int) {
	baseGas = big.NewInt(0)
	switch tx.(type) {
	case *ContractCreateTx, *GAAttachTx, *GAMetaTx:
		return baseGas.Mul(big.NewInt(5), config.Client.BaseGas)
	case *ContractCallTx:
		return baseGas.Mul(big.NewInt(30), config.Client.BaseGas)
	default:
		return config.Client.BaseGas
	}
}

func calcFee(tx TransactionFeeCalculable) (fee *big.Int, err error) {
	gas := big.NewInt(0)
	// baseGas + len(tx)*GasPerByte + contractExecutionGas
	baseGas := baseGasForTxType(tx)
	gas.Add(gas, baseGas)

	s, err := calcRlpLen(tx)
	if err != nil {
		return
	}
	txLenGasPerByte := big.NewInt(s)
	txLenGasPerByte.Mul(txLenGasPerByte, config.Client.GasPerByte)
	gas.Add(gas, txLenGasPerByte)

	gas.Add(gas, tx.GetGasLimit())

	fee = new(big.Int)
	fee = gas.Mul(gas, config.Client.GasPrice)

	tx.SetFee(fee)
	return
}

// CalculateFee calculates the required transaction fee, and increases the fee
// further in case the newer fee ends up increasing the transaction size.
func CalculateFee(tx TransactionFeeCalculable) (err error) {
	var fee, newFee *big.Int
	for {
		fee = tx.GetFee()
		newFee, err = calcFee(tx)
		if err != nil {
			break
		}

		if fee.Cmp(newFee) == 0 {
			break
		}
	}
	return
}

// SerializeTx takes a Tx, runs its RLP() method, and base encodes the result.
func SerializeTx(tx rlp.Encoder) (string, error) {
	w := &bytes.Buffer{}
	err := rlp.Encode(w, tx)
	if err != nil {
		return "", err
	}
	txStr := binary.Encode(binary.PrefixTransaction, w.Bytes())
	return txStr, nil
}

// DeserializeTxStr takes a tx_ string and returns the corresponding Tx struct
func DeserializeTxStr(txRLP string) (Transaction, error) {
	rawRLP, err := binary.Decode(txRLP)
	if err != nil {
		return nil, err
	}
	return DeserializeTx(rawRLP)
}

// DeserializeTx takes a RLP serialized transaction as a bytearray and returns
// the corresponding Tx struct
func DeserializeTx(rawRLP []byte) (Transaction, error) {
	tx, err := GetTransactionType(rawRLP)
	if err != nil {
		return nil, err
	}
	err = rlp.DecodeBytes(rawRLP, tx)
	return tx, err
}

// GetTransactionType reads the RLP input and returns a blank Tx struct of the correct type
func GetTransactionType(rawRLP []byte) (tx Transaction, err error) {
	f := binary.DecodeRLPMessage(rawRLP)[0] // [33] interface, needs to be cast to []uint8
	objTag := uint(f.([]uint8)[0])          // [33] cast to []uint8, get rid of the slice, cast to uint
	return TransactionTypes()[objTag], nil
}

// SignedTx wraps around other Tx structs to hold the signature.
type SignedTx struct {
	Signatures [][]byte
	Tx         rlp.Encoder
}

// EncodeRLP implements rlp.Encoder
func (tx *SignedTx) EncodeRLP(w io.Writer) (err error) {
	/*
		DO NOT WANT
		[
			[11]
			[1]
			[
				[236 231 90 243 220 196 194 60 197 146 118 25 164 100 106 136 121 102 44 60 54 186 255 231 125 101 99 245 135 206 127 202 47 114 210 160 204 85 98 246 178 145 76 58 59 165 110 97 131 144 141 124 223 118 254 14 37 79 8 99 73 97 190 10]
			]
			[
				[12]
				[1]
				[1 206 167 173 228 112 201 249 157 157 78 64 8 128 168 111 29 73 187 68 75 98 241 26 158 187 100 187 207 235 115 254 243]
				[1 31 19 163 176 139 240 1 64 6 98 166 139 105 216 117 247 128 60 236 76 8 100 127 110 213 216 76 120 151 189 80 163]
				[255 255 255 255 255 255 255 255]
				[15 141 103 108 248 0]
				[2 172]
				[1]
				[72 101 108 108 111 32 87 111 114 108 100]
			]
		]

		WANT
		[
		[11]
		[1]
		[
			[173 20 154 64 81 213 186 62 125 201 233 189 58 130 84 238 72 139 204 93 244 135 85 176 84 140 19 30 41 84 113 189 36 16 190 47 230 28 129 84 152 173 60 131 60 55 60 8 127 98 209 161 161 125 188 163 226 193 93 208 202 255 99 1]
		]
		[248 102 12 1 161 1 206 167 173 228 112 201 249 157 157 78 64 8 128 168 111 29 73 187 68 75 98 241 26 158 187 100 187 20...
		]
	*/
	// RLP serialize the wrapped Tx into a plain bytearray.
	wrappedTxRLPBytes, err := rlp.EncodeToBytes(tx.Tx)
	if err != nil {
		return
	}
	// RLP Serialize the SignedTx
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagSignedTransaction,
		rlpMessageVersion,
		tx.Signatures,
		wrappedTxRLPBytes,
	)
	if err != nil {
		return err
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return err
	}
	return nil
}

type signedTxRLP struct {
	ObjectTag         uint
	RlpMessageVersion uint
	Signatures        [][]byte
	WrappedTx         []byte
}

func (stx *signedTxRLP) ReadRLP(s *rlp.Stream) (err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, stx); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *SignedTx) DecodeRLP(s *rlp.Stream) (err error) {
	stx := &signedTxRLP{}
	if err = stx.ReadRLP(s); err != nil {
		return
	}
	wtx, err := DeserializeTx(stx.WrappedTx)
	if err != nil {
		return
	}

	tx.Signatures = stx.Signatures
	tx.Tx = wtx
	return
}

// NewSignedTx ensures that all fields of SignedTx are filled out.
func NewSignedTx(Signatures [][]byte, tx rlp.Encoder) (s *SignedTx) {
	return &SignedTx{
		Signatures: Signatures,
		Tx:         tx,
	}
}

func buildRLPMessage(tag uint, version uint, fields ...interface{}) (rlpRawMsg []byte, err error) {
	// create a message of the transaction and signature
	data := []interface{}{tag, version}
	data = append(data, fields...)
	// fmt.Printf("TX %+v\n\n", data)
	// encode the message using rlp
	rlpRawMsg, err = rlp.EncodeToBytes(data)
	// fmt.Printf("ENCODED %+v\n\n", data)
	return
}

// buildIDTag assemble an id() object see
// https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
func buildIDTag(IDTag uint8, encodedHash string) (v []uint8, err error) {
	raw, err := binary.Decode(encodedHash)
	v = []uint8{IDTag}
	for _, x := range raw {
		v = append(v, uint8(x))
	}
	return
}

// readIDTag disassemble an id() object see
// https://github.com/aeternity/protocol/blob/master/serializations.md#the-id-type
func readIDTag(v []uint8) (IDTag uint8, encodedHash string, err error) {
	IDTag = v[0]
	hash := []byte{}
	for _, x := range v[1:] {
		hash = append(hash, byte(x))
	}

	var prefix binary.HashPrefix
	switch IDTag {
	case IDTagAccount:
		prefix = binary.PrefixAccountPubkey
	case IDTagName:
		prefix = binary.PrefixName
	case IDTagCommitment:
		prefix = binary.PrefixCommitment
	case IDTagOracle:
		prefix = binary.PrefixOraclePubkey
	case IDTagContract:
		prefix = binary.PrefixContractPubkey
	case IDTagChannel:
		prefix = binary.PrefixChannel
	default:
		return 0, "", fmt.Errorf("readIDTag() does not recognize this IDTag (first byte in input array): %v", IDTag)
	}

	encodedHash = binary.Encode(prefix, hash)
	return
}

// VerifySignedTx verifies the signature of a signed transaction, in its RLP
// serialized, base64 encoded tx_ form.
//
// The network ID is also used when calculating the signature, so the network ID
// that the transaction was intended for should be provided too.
func VerifySignedTx(accountID string, txSigned string, networkID string) (valid bool, err error) {
	txRawSigned, _ := binary.Decode(txSigned)
	txRLP := binary.DecodeRLPMessage(txRawSigned)

	// RLP format of signed signature: [[Tag], [Version], [Signatures...],
	// [Transaction]]
	tx := txRLP[3].([]byte)
	txSignature := txRLP[2].([]interface{})[0].([]byte)

	msg := append([]byte(networkID), tx...)

	valid, err = account.Verify(accountID, msg, txSignature)
	if err != nil {
		return
	}
	return
}

// TTLer defines a function that will return an appropriate TTL for a
// transaction.
type TTLer func(offset uint64) (ttl uint64, err error)

// Noncer defines a function that will return an unused account nonce
// for making a transaction.
type Noncer func(accountID string) (nonce uint64, err error)

// TTLNoncer describes a function that combines the roles of TTLer
// and Noncer
type TTLNoncer func(address string, offset uint64) (ttl, nonce uint64, err error)

// GenerateTTLer returns the chain height + offset
func GenerateTTLer(n naet.GetHeighter) TTLer {
	return func(offset uint64) (ttl uint64, err error) {
		height, err := n.GetHeight()
		if err != nil {
			return
		}
		ttl = height + offset
		return
	}
}

// GenerateNoncer retrieves the current accountNonce and adds 1 to it for
// use in transaction building
func GenerateNoncer(n naet.GetAccounter) Noncer {
	return func(accountID string) (nextNonce uint64, err error) {
		a, err := n.GetAccount(accountID)
		if err != nil {
			return
		}
		nextNonce = *a.Nonce + 1
		return
	}
}

// GenerateTTLNoncer combines the commonly used together functions of TTLer
// and Noncer
func GenerateTTLNoncer(ttlFunc TTLer, nonceFunc Noncer) TTLNoncer {
	return func(accountID string, offset uint64) (ttl, nonce uint64, err error) {
		ttl, err = ttlFunc(offset)
		if err != nil {
			return
		}
		nonce, err = nonceFunc(accountID)
		if err != nil {
			return
		}
		return
	}
}
