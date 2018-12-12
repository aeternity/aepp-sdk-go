package aeternity

import (
	"fmt"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
)

// APIGetStatus post transaction
func (ae *Ae) APIGetStatus() (status *models.Status, err error) {
	return getStatus(ae.Epoch)
}

func getStatus(epoch *apiclient.Epoch) (status *models.Status, err error) {
	r, err := epoch.External.GetStatus(nil)
	if err != nil {
		return
	}
	status = r.Payload
	return
}

// APIPostTransaction post transaction
func (ae *Ae) APIPostTransaction(signedEncodedTx, signedEncodedTxHash string) (err error) {
	return postTransaction(ae.Epoch, signedEncodedTx, signedEncodedTxHash)
}

// postTransaction post a transaction to the chain
func postTransaction(epoch *apiclient.Epoch, signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := external.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: &signedEncodedTx,
	})
	r, err := epoch.External.PostTransaction(p)
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
	return getTopBlock(ae.Epoch)
}

// APIGetHeight get the height of the chain
func (ae *Ae) APIGetHeight() (height uint64, err error) {
	tb, err := getTopBlock(ae.Epoch)
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
func getTopBlock(epoch *apiclient.Epoch) (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	r, err := epoch.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// APIGetCurrentKeyBlock get current key block
func (ae *Ae) APIGetCurrentKeyBlock() (kb *models.KeyBlock, err error) {
	return getCurrentKeyBlock(ae.Epoch)
}

func getCurrentKeyBlock(epoch *apiclient.Epoch) (kb *models.KeyBlock, err error) {
	r, err := epoch.External.GetCurrentKeyBlock(nil)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	kb = r.Payload
	return
}

// APIGetAccount return the account
func (ae *Ae) APIGetAccount(accountID string) (account *models.Account, err error) {
	return getAccount(ae.Epoch, accountID)
}

// getAccount retrieve an account by its address (public key)
// it is particularly useful to obtain the nonce for spending transactions
func getAccount(epoch *apiclient.Epoch, accountID string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(accountID)
	r, err := epoch.External.GetAccountByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	account = r.Payload
	return
}

// APIGetNameEntryByName return the name entry
func (ae *Ae) APIGetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	return getNameEntryByName(ae.Epoch, name)
}

func getNameEntryByName(epoch *apiclient.Epoch, name string) (nameEntry *models.NameEntry, err error) {
	p := external.NewGetNameEntryByNameParams().WithName(name)
	r, err := epoch.External.GetNameEntryByName(p)

	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}

	nameEntry = r.Payload
	return
}

// APIGetMicroBlockTransactionsByHash get the transactions of a microblock
func (ae *Ae) APIGetMicroBlockTransactionsByHash(microBlockID string) (txs *models.GenericTxs, err error) {
	return getMicroBlockTransactionsByHash(ae.Epoch, microBlockID)
}

func getMicroBlockTransactionsByHash(epoch *apiclient.Epoch, microBlockID string) (txs *models.GenericTxs, err error) {
	p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(microBlockID)
	r, err := epoch.External.GetMicroBlockTransactionsByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetMicroBlockHeaderByHash get the header of a micro block
func (ae *Ae) APIGetMicroBlockHeaderByHash(microBlockID string) (txs *models.MicroBlockHeader, err error) {
	return getMicroBlockHeaderByHash(ae.Epoch, microBlockID)
}

func getMicroBlockHeaderByHash(epoch *apiclient.Epoch, microBlockID string) (txs *models.MicroBlockHeader, err error) {
	p := external.NewGetMicroBlockHeaderByHashParams().WithHash(microBlockID)
	r, err := epoch.External.GetMicroBlockHeaderByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetKeyBlockByHash get a key block by its hash
func (ae *Ae) APIGetKeyBlockByHash(keyBlockID string) (txs *models.KeyBlock, err error) {
	return getKeyBlockByHash(ae.Epoch, keyBlockID)
}

func getKeyBlockByHash(epoch *apiclient.Epoch, keyBlockID string) (txs *models.KeyBlock, err error) {
	p := external.NewGetKeyBlockByHashParams().WithHash(keyBlockID)
	r, err := epoch.External.GetKeyBlockByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// APIGetTransactionByHash get a transaction by it's hash
func (ae *Ae) APIGetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	return getTransactionByHash(ae.Epoch, txHash)
}

// getTransactionByHash retrieve a transaction by it's hash
func getTransactionByHash(epoch *apiclient.Epoch, txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := epoch.External.GetTransactionByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	tx = r.Payload
	return
}

// APIGetOracleByPubkey get an oracle by it's public key
func (ae *Ae) APIGetOracleByPubkey(pubkey string) (oracle *models.RegisteredOracle, err error) {
	return getOracleByPubkey(ae.Epoch, pubkey)
}

// getOracleByPubkey get an oracle by it's public key
func getOracleByPubkey(epoch *apiclient.Epoch, pubkey string) (oracle *models.RegisteredOracle, err error) {
	p := external.NewGetOracleByPubkeyParams().WithPubkey(pubkey)
	r, err := epoch.External.GetOracleByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracle = r.Payload
	return
}
