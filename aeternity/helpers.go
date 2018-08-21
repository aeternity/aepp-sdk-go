package aeternity

import (
	"encoding/binary"
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
func NewCli(epochURL string, debug bool) *Ae {
	// create the transport
	host, schemas := urlComponents(epochURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClinet := apiclient.New(transport, strfmt.Default)
	aecli := &Ae{
		Epoch: openAPIClinet,
		Wallet: &Wallet{
			epochCli: openAPIClinet,
		},
		Aens: &Aens{
			epochCli: openAPIClinet,
		},
		Contract: &Contract{
			epochCli: openAPIClinet,
		},
	}
	return aecli
}

// NewCliW obtain a new epochCli instance
func NewCliW(epochURL string, kp *KeyPair, debug bool) *Ae {
	// create the transport
	host, schemas := urlComponents(epochURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClinet := apiclient.New(transport, strfmt.Default)
	aecli := &Ae{
		Epoch: openAPIClinet,
		Wallet: &Wallet{
			epochCli: openAPIClinet,
			owner:    kp,
		},
		Aens: &Aens{
			epochCli: openAPIClinet,
			owner:    kp,
		},
		Contract: &Contract{
			epochCli: openAPIClinet,
			owner:    kp,
		},
	}
	return aecli
}

func urlComponents(url string) (host string, schemas []string) {
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

// GetAbsoluteHeight return the chain height adding the offset
func GetAbsoluteHeight(epochCli *apiclient.Epoch, offset int64) (height int64) {
	if r, err := epochCli.Operations.GetTop(nil); err == nil {
		height = r.Payload.Height + offset
	}
	return
}

// PostTransaction post a transaction to the chain
func PostTransaction(epochCli *apiclient.Epoch, signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := operations.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: signedEncodedTx,
	})
	r, err := epochCli.Operations.PostTransaction(p)
	if err != nil {
		return
	}
	if r.Payload.TxHash != models.EncodedHash(signedEncodedTxHash) {
		err = fmt.Errorf("Transaction hash mismatch, expected %s got %s", signedEncodedTxHash, r.Payload.TxHash)
	}
	return
}

// SignEncodeTx sign and encode a transaction
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
	data := struct {
		Tag        uint
		Vsn        uint
		Signatures [][]byte
		TxRaw      []byte
	}{
		11,
		1,
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

// Ae the aeternity client
type Ae struct {
	*apiclient.Epoch
	*Wallet
	*Aens
	*Contract
}

// Wallet high level abstraction for operation on a wallet
type Wallet struct {
	epochCli *apiclient.Epoch
	owner    *KeyPair
}

// Aens abstractions for aens operations
type Aens struct {
	epochCli     *apiclient.Epoch
	owner        *KeyPair
	name         string
	preClaimSalt []byte
}

// Contract abstractions for contracts
type Contract struct {
	epochCli *apiclient.Epoch
	owner    *KeyPair
	source   string
}

// WithKeyPair associate a keypair with the client
func (ae *Ae) WithKeyPair(kp *KeyPair) *Ae {
	ae.Wallet.owner = kp
	ae.Aens.owner = kp
	ae.Contract.owner = kp
	return ae
}

// Spend transfer tokens from an account to another
func (w *Wallet) Spend(recipientAddress string, amount int64) (txHash, signature string, err error) {
	// calculate the absolute ttl for the transaction
	absoluteTTL := GetAbsoluteHeight(w.epochCli, Config.P.Client.TxTTL)
	// create spend transaction
	ps, err := w.epochCli.Operations.PostSpend(operations.NewPostSpendParams().WithBody(&models.SpendTx{
		RecipientPubkey: models.EncodedHash(recipientAddress),
		Sender:          models.EncodedHash(w.owner.Address),
		Amount:          &amount,
		TTL:             absoluteTTL,
		Fee:             &Config.P.Client.Fee,
		Payload:         &Config.P.Tuning.TxPayload,
	}))
	if err != nil {
		return
	}
	// this is the transaction to sign
	postSpendTransaction := ps.Payload.Tx
	// sign the above transaction with the private key
	tx, txHash, signature, err := SignEncodeTx(w.owner, postSpendTransaction)
	if err != nil {
		return
	}
	// post the signed transaction to the chain
	pt, err := w.epochCli.Operations.PostTx(operations.NewPostTxParams().WithBody(
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

// naming

func commitmentHash(name string) (ch string, salt []byte, err error) {
	salt, err = randomBytes(32)
	if err != nil {
		return
	}
	nh := append(namehash(name), salt...)
	ch = encodeP(PrefixNameCommitment, nh)
	return
}

// Claim perform a name claiming
func (n *Aens) Claim(name string, targetAddress string) (err error) {
	// get the ttl offset
	ttl := GetAbsoluteHeight(n.epochCli, Config.P.Client.TxTTL)
	// calculate the commitment and get the preclaim salt
	cm, salt, err := commitmentHash(name)
	if err != nil {
		return
	}
	// preclaim transaction
	pc := operations.NewPostNamePreclaimParams().WithBody(
		&models.NamePreclaimTx{
			Account:    models.EncodedHash(n.owner.Address),
			Commitment: &cm,
			Fee:        &Config.P.Client.Names.PreClaimFee,
			TTL:        ttl,
		},
	)
	pcr, err := n.epochCli.Operations.PostNamePreclaim(pc)
	if err != nil {
		return
	}
	// this is the transaction to sign
	postPreclaimTransaction := pcr.Payload.Tx
	// sign the above transaction with the private key
	tx, txHash, _, err := SignEncodeTx(n.owner, postPreclaimTransaction)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = PostTransaction(n.epochCli, tx, txHash)
	if err != nil {
		return
	}
	// claim transaction
	encodedName := encodeP("nm", []byte(name))
	int64Salt := int64(binary.BigEndian.Uint64(salt))
	c := operations.NewPostNameClaimParams().WithBody(
		&models.NameClaimTx{
			Account:  models.EncodedHash(n.owner.Address),
			Name:     &encodedName,
			NameSalt: &int64Salt,
			Fee:      &Config.P.Client.Names.ClaimFee,
			TTL:      ttl,
		},
	)
	cr, err := n.epochCli.Operations.PostNameClaim(c)
	if err != nil {
		return
	}
	// this is the transaction to sign
	postClaimTransaction := cr.Payload.Tx
	// sign the above transaction with the private key
	tx, txHash, _, err = SignEncodeTx(n.owner, postClaimTransaction)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = PostTransaction(n.epochCli, tx, txHash)
	if err != nil {
		return
	}
	// update name
	encodedNameHash := encodeP(PrefixNameHash, namehash(name))
	absClientTTL := GetAbsoluteHeight(n.epochCli, Config.P.Client.Names.ClientTTL)
	absNameTTL := GetAbsoluteHeight(n.epochCli, Config.P.Client.Names.TTL)
	pointers := fmt.Sprintf(`{ "account_pubkey": "%s" }`, targetAddress)
	u := operations.NewPostNameUpdateParams().WithBody(&models.NameUpdateTx{
		Account:   models.EncodedHash(n.owner.Address),
		NameHash:  &encodedNameHash,
		ClientTTL: &absClientTTL,
		NameTTL:   &absNameTTL,
		TTL:       ttl,
		Fee:       &Config.P.Client.Names.UpdateFee,
		Pointers:  &pointers,
	})
	ur, err := n.epochCli.Operations.PostNameUpdate(u)
	if err != nil {
		return
	}
	// this is the transaction to sign
	postUpdateTransaction := ur.Payload.Tx
	// sign the above transaction with the private key
	tx, txHash, _, err = SignEncodeTx(n.owner, postUpdateTransaction)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = PostTransaction(n.epochCli, tx, txHash)
	return
}
