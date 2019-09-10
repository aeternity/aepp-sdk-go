package aeternity

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/utils"
	rlp "github.com/randomshinichi/rlpae"
)

// Transaction is used to indicate a transaction of any type.
type Transaction interface {
	rlp.Encoder
}

// Sign calculates the signature of the SignedTx.Tx. Although it does not use
// the SignedTx itself, it takes a SignedTx as an argument because if it took a
// rlp.Encoder as an interface, one might expect the signature to be of the
// SignedTx itself, which won't work.
func Sign(kp *Account, tx *SignedTx, networkID string) (signature []byte, err error) {
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
	rlpTxHashRaw, err := Blake2bHash(rlpTxRaw)
	if err != nil {
		return "", err
	}

	txhash = Encode(PrefixTransactionHash, rlpTxHashRaw)
	return txhash, nil
}

// SignHashTx wraps a *Tx struct in a SignedTx, then returns its signature and
// hash.
func SignHashTx(kp *Account, tx Transaction, networkID string) (signedTx *SignedTx, txhash, signature string, err error) {
	signedTx = NewSignedTx([][]byte{}, tx)
	var signatureBytes []byte

	if _, ok := tx.(*GAMetaTx); !ok {
		signatureBytes, err = Sign(kp, signedTx, networkID)
		if err != nil {
			return
		}
		signedTx.Signatures = append(signedTx.Signatures, signatureBytes)
		signature = Encode(PrefixSignature, signatureBytes)

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

func calcFeeStd(tx rlp.Encoder, txLen int) *big.Int {
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
func calcSizeEstimate(tx rlp.Encoder, fee *big.Int) (int, error) {
	feeRlp, err := rlp.EncodeToBytes(fee)
	if err != nil {
		return 0, err
	}
	feeRlpLen := len(feeRlp)

	w := &bytes.Buffer{}
	err = rlp.Encode(w, tx)
	if err != nil {
		return 0, err
	}

	rlpRawMsg := w.Bytes()
	return len(rlpRawMsg) - feeRlpLen + 8, nil
}

// SerializeTx takes a Tx, runs its RLP() method, and base encodes the result.
func SerializeTx(tx rlp.Encoder) (string, error) {
	w := &bytes.Buffer{}
	err := rlp.Encode(w, tx)
	if err != nil {
		return "", err
	}
	txStr := Encode(PrefixTransaction, w.Bytes())
	return txStr, nil
}

// DeserializeTxStr takes a tx_ string and returns the corresponding Tx struct
func DeserializeTxStr(txRLP string) (Transaction, error) {
	rawRLP, err := Decode(txRLP)
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
	f := DecodeRLPMessage(rawRLP)[0] // [33] interface, needs to be cast to []uint8
	objTag := uint(f.([]uint8)[0])   // [33] cast to []uint8, get rid of the slice, cast to uint
	return TransactionTypes[objTag], nil
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
