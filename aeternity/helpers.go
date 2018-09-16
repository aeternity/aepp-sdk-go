package aeternity

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
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

// getTopBlock get the top block of the chain
// wraps the generated code to avoid too much changes
// in case of the swagger api call changes
func getTopBlock(epochCli *apiclient.Epoch) (kb *models.KeyBlock, err error) {
	r, err := epochCli.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// return the current key block
func getCurrentKeyBlock(epochCli *apiclient.Epoch) (kb *models.KeyBlock, err error) {
	r, err := epochCli.External.GetCurrentKeyBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// getAbsoluteHeight return the chain height adding the offset
func getAbsoluteHeight(epochCli *apiclient.Epoch, offset int64) (height int64, err error) {
	kb, err := getTopBlock(epochCli)
	if err != nil {
		return
	}
	height = kb.Height + offset
	return
}

// getAccount retrieve an account by its address (public key)
// it is particularly useful to obtain the nonce for spending transactions
func getAccount(epochCli *apiclient.Epoch, address string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(address)
	r, err := epochCli.External.GetAccountByPubkey(p)
	if err != nil {
		return
	}
	account = r.Payload
	return
}

// getNextNonce retrieve the next nonce for an account
// it has to query the chain to do so
func getNextNonce(epochCli *apiclient.Epoch, acccount *KeyPair) (nextNonce uint64, err error) {
	a, err := getAccount(epochCli, acccount.Address)
	if err != nil {
		return
	}
	nextNonce = *a.Nonce + 1
	return
}

// getTransaction retrieve a transaction by it's hash
func getTransaction(epochCli *apiclient.Epoch, txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := epochCli.External.GetTransactionByHash(p)
	if err != nil {
		return
	}
	tx = r.Payload
	return
}

// waitForTransaction to appear on the chain
func waitForTransaction(epochCli *apiclient.Epoch, txHash string) (blockHeight int64, blockHash string, err error) {
	// caclulate the date for the timeout
	ctm := Config.P.Tuning.ChainTimeout
	tm := time.Now().Add(time.Millisecond * time.Duration(ctm))
	// start querying the transaction
	for {
		if time.Now().After(tm) {
			// TODO: should use the chain height instead of a timeout
			err = fmt.Errorf("Timeout waiting for the transaction to appear")
			break // timeout execed
		}
		tx, err := getTransaction(epochCli, txHash)
		if err != nil {
			break
		}
		if len(tx.BlockHash) > 0 {
			blockHeight = tx.BlockHeight
			blockHash = fmt.Sprint(tx.BlockHash)
			break
		}
		time.Sleep(time.Millisecond * time.Duration(Config.P.Tuning.ChainPollInteval))
	}
	return
}

// waitForTransactionUntillHeight waits for a transaction until heightLimit (inclusive) is reached
func waitForTransactionUntillHeight(epochCli *apiclient.Epoch, height int64, txHash string) (blockHeight int64, blockHash, microBlockHash string, tx *models.GenericSignedTx, err error) {

	kb, err := getCurrentKeyBlock(epochCli)
	if err != nil {
		return
	}
	// current height
	targetHeight := kb.Height
	nextHeight := kb.Height
	// hold the generation
	var g *models.Generation

Main:
	for {
		// check the top height
		if targetHeight > height {
			err = fmt.Errorf("Transaction %s not found, height %d", txHash, height)
			break
		}
		// get the generation of the targetHeight
		fmt.Println("Try with block ", targetHeight)
		p := external.NewGetGenerationByHeightParams().WithHeight(targetHeight)
		r, err := epochCli.External.GetGenerationByHeight(p)
		if err != nil {
			break
		}
		g = r.Payload
		// search for transaction in the microblocks
		for _, mbh := range g.MicroBlocks {
			// get the microblok
			mbhs := fmt.Sprint(mbh)
			p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(mbhs)
			r, mErr := epochCli.External.GetMicroBlockTransactionsByHash(p)
			if mErr != nil {
				err = mErr
				break Main
			}
			// go through all the hashes
			for _, btx := range r.Payload.Transactions {
				if fmt.Sprint(btx.Hash) == txHash {
					// transaction found !!
					blockHash = fmt.Sprint(g.KeyBlock.Hash)
					blockHeight = g.KeyBlock.Height
					microBlockHash = mbhs
					tx = btx
					break Main
				}
			}
		}
		// here we want to query one more time the current generation
		// before switching to the next one in case microblocks have been added meanwhile
		// update targetHeight
		if nextHeight > targetHeight {
			targetHeight = nextHeight
		}
		// update nextHeight
		kb, err = getCurrentKeyBlock(epochCli)
		if err != nil {
			break
		}
		nextHeight = kb.Height
	}

	return
}

// postTransaction post a transaction to the chain
func postTransaction(epochCli *apiclient.Epoch, signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := external.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: signedEncodedTx,
	})
	r, err := epochCli.External.PostTransaction(p)
	if err != nil {
		return
	}
	if r.Payload.TxHash != models.EncodedHash(signedEncodedTxHash) {
		err = fmt.Errorf("Transaction hash mismatch, expected %s got %s", signedEncodedTxHash, r.Payload.TxHash)
	}
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
