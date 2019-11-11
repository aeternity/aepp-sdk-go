package transactions

import (
	"io"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v7/binary"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/utils"
	rlp "github.com/randomshinichi/rlpae"
)

// SpendTx represents a simple transaction where one party sends another AE
type SpendTx struct {
	SenderID    string
	RecipientID string
	Amount      *big.Int
	Fee         *big.Int
	Payload     []byte
	TTL         uint64
	Nonce       uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *SpendTx) EncodeRLP(w io.Writer) (err error) {
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
	rlpRawMsg, err := buildRLPMessage(
		ObjectTagSpendTransaction,
		rlpMessageVersion,
		sID,
		rID,
		tx.Amount,
		tx.Fee,
		tx.TTL,
		tx.Nonce,
		[]byte(tx.Payload))

	if err != nil {
		return
	}
	_, err = w.Write(rlpRawMsg)
	if err != nil {
		return
	}
	return nil
}

type spendRLP struct {
	ObjectTagSpendTransaction uint
	RlpMessageVersion         uint
	SenderID                  []uint8
	ReceiverID                []uint8
	Amount                    *big.Int
	Fee                       *big.Int
	TTL                       uint64
	Nonce                     uint64
	Payload                   []byte
}

func (stx *spendRLP) ReadRLP(s *rlp.Stream) (sID, rID string, err error) {
	var blob []byte
	if blob, err = s.Raw(); err != nil {
		return
	}
	if err = rlp.DecodeBytes(blob, stx); err != nil {
		return
	}
	if _, sID, err = readIDTag(stx.SenderID); err != nil {
		return
	}
	if _, rID, err = readIDTag(stx.ReceiverID); err != nil {
		return
	}
	return
}

// DecodeRLP implements rlp.Decoder
func (tx *SpendTx) DecodeRLP(s *rlp.Stream) (err error) {
	stx := &spendRLP{}
	sID, rID, err := stx.ReadRLP(s)
	if err != nil {
		return
	}

	tx.SenderID = sID
	tx.RecipientID = rID
	tx.Amount = stx.Amount
	tx.Fee = stx.Fee
	tx.TTL = stx.TTL
	tx.Nonce = stx.Nonce
	tx.Payload = stx.Payload
	return
}

// JSON representation of a Tx is useful for querying the node's debug endpoint
func (tx *SpendTx) JSON() (string, error) {
	baseEncodedPayload := binary.Encode(binary.PrefixByteArray, tx.Payload)
	swaggerT := models.SpendTx{
		Amount:      utils.BigInt(*tx.Amount),
		Fee:         utils.BigInt(*tx.Fee),
		Nonce:       tx.Nonce,
		Payload:     &baseEncodedPayload,
		RecipientID: &tx.RecipientID,
		SenderID:    &tx.SenderID,
		TTL:         tx.TTL,
	}
	output, err := swaggerT.MarshalBinary()
	return string(output), err
}

// SetFee implements Transaction
func (tx *SpendTx) SetFee(f *big.Int) {
	tx.Fee = f
}

// GetFee implements Transaction
func (tx *SpendTx) GetFee() *big.Int {
	return tx.Fee
}

// CalcGas implements Transaction
func (tx *SpendTx) CalcGas() (g *big.Int, err error) {
	baseGas := new(big.Int)
	baseGas.Add(baseGas, config.Client.BaseGas)
	gasComponent, err := normalGasComponent(tx, big.NewInt(0))
	if err != nil {
		return
	}
	g = new(big.Int)
	g = g.Add(baseGas, gasComponent)
	return
}

// NewSpendTx is a constructor for a SpendTx struct
func NewSpendTx(senderID, recipientID string, amount *big.Int, payload []byte, ttlnoncer TTLNoncer) (tx *SpendTx, err error) {
	ttl, nonce, err := ttlnoncer(senderID, config.Client.TTL)
	if err != nil {
		return
	}

	tx = &SpendTx{senderID, recipientID, amount, config.Client.Fee, payload, ttl, nonce}
	CalculateFee(tx)
	return
}
