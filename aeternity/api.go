package aeternity

import (
	"fmt"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
)

// Ae.API*() methods are the stable interface to Go code that uses this SDK.
// Logic implementation is handled by the unexported functions.

// APIGetStatus post transaction
func (ae *Ae) APIGetStatus() (status *models.Status, err error) {
	return getStatus(ae.Node)
}

func getStatus(node *apiclient.Node) (status *models.Status, err error) {
	r, err := node.External.GetStatus(nil)
	if err != nil {
		return
	}
	status = r.Payload
	return
}

// APIPostTransaction post transaction
func (ae *Ae) APIPostTransaction(signedEncodedTx, signedEncodedTxHash string) (err error) {
	return postTransaction(ae.Node, signedEncodedTx, signedEncodedTxHash)
}

// postTransaction post a transaction to the chain
func postTransaction(node *apiclient.Node, signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := external.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: &signedEncodedTx,
	})
	r, err := node.External.PostTransaction(p)
	if err != nil {
		return
	}
	if r.Payload.TxHash != models.EncodedHash(signedEncodedTxHash) {
		err = fmt.Errorf("Transaction hash mismatch, expected %s got %s", signedEncodedTxHash, r.Payload.TxHash)
	}
	return
}

// APIGetTopBlock get the top block of the chain
func (ae *Ae) APIGetTopBlock() (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	return getTopBlock(ae.Node)
}

// APIGetHeight get the height of the chain
func (ae *Ae) APIGetHeight() (height uint64, err error) {
	tb, err := getTopBlock(ae.Node)
	if err != nil {
		return
	}
	if tb.KeyBlock == nil {
		height = *tb.MicroBlock.Height
		return
	}
	height = *tb.KeyBlock.Height
	return
}

// getTopBlock get the top block of the chain
// wraps the generated code to avoid too much changes
// in case of the swagger api call changes
func getTopBlock(node *apiclient.Node) (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	r, err := node.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// APIGetCurrentKeyBlock get current key block
func (ae *Ae) APIGetCurrentKeyBlock() (kb *models.KeyBlock, err error) {
	return getCurrentKeyBlock(ae.Node)
}

func getCurrentKeyBlock(node *apiclient.Node) (kb *models.KeyBlock, err error) {
	r, err := node.External.GetCurrentKeyBlock(nil)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	kb = r.Payload
	return
}

// APIGetAccount return the account
func (ae *Ae) APIGetAccount(accountID string) (account *models.Account, err error) {
	return getAccount(ae.Node, accountID)
}

// getAccount retrieve an account by its address (public key)
// it is particularly useful to obtain the nonce for spending transactions
func getAccount(node *apiclient.Node, accountID string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(accountID)
	r, err := node.External.GetAccountByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	account = r.Payload
	return
}

// APIGetNameEntryByName return the name entry
func (ae *Ae) APIGetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	return getNameEntryByName(ae.Node, name)
}

func getNameEntryByName(node *apiclient.Node, name string) (nameEntry *models.NameEntry, err error) {
	p := external.NewGetNameEntryByNameParams().WithName(name)
	r, err := node.External.GetNameEntryByName(p)

	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}

	nameEntry = r.Payload
	return
}

// APIGetMicroBlockTransactionsByHash get the transactions of a microblock
func (ae *Ae) APIGetMicroBlockTransactionsByHash(microBlockID string) (txs *models.GenericTxs, err error) {
	return getMicroBlockTransactionsByHash(ae.Node, microBlockID)
}

func getMicroBlockTransactionsByHash(node *apiclient.Node, microBlockID string) (txs *models.GenericTxs, err error) {
	p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(microBlockID)
	r, err := node.External.GetMicroBlockTransactionsByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetMicroBlockHeaderByHash get the header of a micro block
func (ae *Ae) APIGetMicroBlockHeaderByHash(microBlockID string) (txs *models.MicroBlockHeader, err error) {
	return getMicroBlockHeaderByHash(ae.Node, microBlockID)
}

func getMicroBlockHeaderByHash(node *apiclient.Node, microBlockID string) (txs *models.MicroBlockHeader, err error) {
	p := external.NewGetMicroBlockHeaderByHashParams().WithHash(microBlockID)
	r, err := node.External.GetMicroBlockHeaderByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetKeyBlockByHash get a key block by its hash
func (ae *Ae) APIGetKeyBlockByHash(keyBlockID string) (txs *models.KeyBlock, err error) {
	return getKeyBlockByHash(ae.Node, keyBlockID)
}

func getKeyBlockByHash(node *apiclient.Node, keyBlockID string) (txs *models.KeyBlock, err error) {
	p := external.NewGetKeyBlockByHashParams().WithHash(keyBlockID)
	r, err := node.External.GetKeyBlockByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetTransactionByHash get a transaction by it's hash
func (ae *Ae) APIGetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	return getTransactionByHash(ae.Node, txHash)
}

// getTransactionByHash retrieve a transaction by it's hash
func getTransactionByHash(node *apiclient.Node, txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := node.External.GetTransactionByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	tx = r.Payload
	return
}

// APIGetOracleByPubkey get an oracle by it's public key
func (ae *Ae) APIGetOracleByPubkey(pubkey string) (oracle *models.RegisteredOracle, err error) {
	return getOracleByPubkey(ae.Node, pubkey)
}

// getOracleByPubkey get an oracle by it's public key
func getOracleByPubkey(node *apiclient.Node, pubkey string) (oracle *models.RegisteredOracle, err error) {
	p := external.NewGetOracleByPubkeyParams().WithPubkey(pubkey)
	r, err := node.External.GetOracleByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracle = r.Payload
	return
}
