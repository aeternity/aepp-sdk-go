package aeternity

import (
	"fmt"

	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetStatus post transaction
func (ae *Client) GetStatus() (status *models.Status, err error) {
	return getStatus(ae)
}

func getStatus(node *Client) (status *models.Status, err error) {
	r, err := node.External.GetStatus(nil)
	if err != nil {
		return
	}
	status = r.Payload
	return
}

// PostTransaction post transaction
func (ae *Client) PostTransaction(signedEncodedTx, signedEncodedTxHash string) (err error) {
	return postTransaction(ae, signedEncodedTx, signedEncodedTxHash)
}

// postTransaction post a transaction to the chain
func postTransaction(node *Client, signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := external.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: models.EncodedByteArray(signedEncodedTx),
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

// GetTopBlock get the top block of the chain
func (ae *Client) GetTopBlock() (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	return getTopBlock(ae)
}

// GetHeight get the height of the chain
func (ae *Client) GetHeight() (height uint64, err error) {
	tb, err := getTopBlock(ae)
	if err != nil {
		return
	}
	if tb.KeyBlock == nil {
		height = uint64(tb.MicroBlock.Height)
		return
	}
	height = uint64(tb.KeyBlock.Height)
	return
}

// getTopBlock get the top block of the chain
// wraps the generated code to avoid too much changes
// in case of the swagger api call changes
func getTopBlock(node *Client) (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	r, err := node.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// GetCurrentKeyBlock get current key block
func (ae *Client) GetCurrentKeyBlock() (kb *models.KeyBlock, err error) {
	return getCurrentKeyBlock(ae)
}

func getCurrentKeyBlock(node *Client) (kb *models.KeyBlock, err error) {
	r, err := node.External.GetCurrentKeyBlock(nil)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	kb = r.Payload
	return
}

// GetAccount return the account
func (ae *Client) GetAccount(accountID string) (account *models.Account, err error) {
	return getAccount(ae, accountID)
}

// getAccount retrieve an account by its address (public key)
// it is particularly useful to obtain the nonce for spending transactions
func getAccount(node *Client, accountID string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(accountID)
	r, err := node.External.GetAccountByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	account = r.Payload
	return
}

// GetNameEntryByName return the name entry
func (ae *Client) GetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	return getNameEntryByName(ae, name)
}

func getNameEntryByName(node *Client, name string) (nameEntry *models.NameEntry, err error) {
	p := external.NewGetNameEntryByNameParams().WithName(name)
	r, err := node.External.GetNameEntryByName(p)

	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}

	nameEntry = r.Payload
	return
}

// GetMicroBlockTransactionsByHash get the transactions of a microblock
func (ae *Client) GetMicroBlockTransactionsByHash(microBlockID string) (txs *models.GenericTxs, err error) {
	return getMicroBlockTransactionsByHash(ae, microBlockID)
}

func getMicroBlockTransactionsByHash(node *Client, microBlockID string) (txs *models.GenericTxs, err error) {
	p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(microBlockID)
	r, err := node.External.GetMicroBlockTransactionsByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetMicroBlockHeaderByHash get the header of a micro block
func (ae *Client) GetMicroBlockHeaderByHash(microBlockID string) (txs *models.MicroBlockHeader, err error) {
	return getMicroBlockHeaderByHash(ae, microBlockID)
}

func getMicroBlockHeaderByHash(node *Client, microBlockID string) (txs *models.MicroBlockHeader, err error) {
	p := external.NewGetMicroBlockHeaderByHashParams().WithHash(microBlockID)
	r, err := node.External.GetMicroBlockHeaderByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetKeyBlockByHash get a key block by its hash
func (ae *Client) GetKeyBlockByHash(keyBlockID string) (txs *models.KeyBlock, err error) {
	return getKeyBlockByHash(ae, keyBlockID)
}

func getKeyBlockByHash(node *Client, keyBlockID string) (txs *models.KeyBlock, err error) {
	p := external.NewGetKeyBlockByHashParams().WithHash(keyBlockID)
	r, err := node.External.GetKeyBlockByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetTransactionByHash get a transaction by it's hash
func (ae *Client) GetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	return getTransactionByHash(ae, txHash)
}

// getTransactionByHash retrieve a transaction by it's hash
func getTransactionByHash(node *Client, txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := node.External.GetTransactionByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	tx = r.Payload
	return
}

// GetOracleByPubkey get an oracle by it's public key
func (ae *Client) GetOracleByPubkey(pubkey string) (oracle *models.RegisteredOracle, err error) {
	return getOracleByPubkey(ae, pubkey)
}

// getOracleByPubkey get an oracle by it's public key
func getOracleByPubkey(node *Client, pubkey string) (oracle *models.RegisteredOracle, err error) {
	p := external.NewGetOracleByPubkeyParams().WithPubkey(pubkey)
	r, err := node.External.GetOracleByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracle = r.Payload
	return
}

// GetOracleQueriesByPubkey get a list of queries made to a particular oracle
func (ae *Client) GetOracleQueriesByPubkey(pubkey string) (oracleQueries *models.OracleQueries, err error) {
	return getOracleQueriesByPubkey(ae, pubkey)
}

// getOracleQueriesByPubkey get a list of queries made to a particular oracle
func getOracleQueriesByPubkey(node *Client, pubkey string) (oracleQueries *models.OracleQueries, err error) {
	p := external.NewGetOracleQueriesByPubkeyParams().WithPubkey(pubkey)
	r, err := node.External.GetOracleQueriesByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracleQueries = r.Payload
	return
}

// GetContractByID gets a contract by ct_ ID
func (ae *Client) GetContractByID(ctID string) (contract *models.ContractObject, err error) {
	return getContractByID(ae, ctID)
}

// getContractByID get a contract by ct_ ID
func getContractByID(node *Client, ctID string) (contract *models.ContractObject, err error) {
	p := external.NewGetContractParams().WithPubkey(ctID)
	r, err := node.External.GetContract(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	contract = r.Payload
	return
}
