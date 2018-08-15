package aeternity

import (
	"fmt"
	"strings"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/operations"
	"github.com/aeternity/aepp-sdk-go/generated/models"
	"github.com/aeternity/aepp-sdk-go/rlp"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCli obtain a new epochCli instance
func NewCli(epochURL string, debug bool) *apiclient.Epoch {
	// create the transport
	host, schemas := urlComponents(epochURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	return apiclient.New(transport, strfmt.Default)
}

func urlComponents(url string) (host string, schemas []string) {
	// TODO: buuld the right host
	p := strings.Split(url, "://")
	if len(p) == 1 {
		host = p[0]
		schemas = []string{"http"}
		return
	}
	host = p[1]
	schemas = []string{p[0]}
	return
}

func SignEncodeTx(kp *KeyPair, tx models.EncodedHash) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
	txStr := fmt.Sprint(tx)
	// decode the transaction string
	txRaw, err := decode(txStr)
	if err != nil {
		return
	}
	// sign the transaction
	sigRaw := kp.Sign(txRaw)
	if err != nil {
		return
	}
	// create a message of the transaction and signature
	data := []interface{}{
		[]uint{11},
		[]uint{1},
		[][]byte{sigRaw},
		txRaw,
	}
	// encode the message using rlp
	rlpTxRaw, err := rlp.EncodeToBytes(data)
	// encode the rlp message with the prefix
	signedEncodedTx = encodeP(PrefixTx, rlpTxRaw)
	// compute the hash
	rlpTxHashRaw, err := hash(rlpTxRaw)
	signedEncodedTxHash = encodeP(PrefixTxHash, rlpTxHashRaw)
	// encode the signature
	signature = encodeP(PrefixSignature, sigRaw)
	return
}

func Spend(epochCli *apiclient.Epoch, sender *KeyPair, recipientAddress string, amount int64) (txHash, signature string, err error) {
	// compute absolute ttl
	t, err := epochCli.Operations.GetTop(nil)
	if err != nil {
		return
	}
	// calculate the absolute ttl for the transaction
	absoluteTTL := t.Payload.Height + Config.Client.TxTTL
	// create spend transaction
	ps, err := epochCli.Operations.PostSpend(operations.NewPostSpendParams().WithBody(&models.SpendTx{
		RecipientPubkey: models.EncodedHash(recipientAddress),
		Sender:          models.EncodedHash(sender.Address),
		Amount:          &amount,
		TTL:             absoluteTTL,
		Fee:             &Config.Client.Fee,
		Payload:         &Config.Tuning.TxPayload,
	}))
	if err != nil {
		return
	}
	// this is the transaction to sign
	postSpendTransaction := ps.Payload.Tx
	// sign the above transaction with the private key
	tx, txHash, signature, err := SignEncodeTx(sender, postSpendTransaction)
	if err != nil {
		return
	}
	// post the signed transaction to the chain
	pt, err := epochCli.Operations.PostTx(operations.NewPostTxParams().WithBody(
		&models.Tx{Tx: tx},
	))
	if err != nil {
		return
	}
	// verify the transaction hash
	if models.EncodedHash(txHash) != pt.Payload.TxHash {
		err = fmt.Errorf("Transaction hash mismatch, expected %s got %s", txHash, pt.Payload.TxHash)
	}
	return
}
