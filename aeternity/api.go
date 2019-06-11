package aeternity

import (
	"fmt"

	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetStatus post transaction
func (c *Client) GetStatus() (status *models.Status, err error) {
	r, err := c.External.GetStatus(nil)
	if err != nil {
		return
	}
	status = r.Payload
	return
}

// PostTransaction post transaction
func (c *Client) PostTransaction(signedEncodedTx, signedEncodedTxHash string) (err error) {
	p := external.NewPostTransactionParams().WithBody(&models.Tx{
		Tx: &signedEncodedTx,
	})
	r, err := c.External.PostTransaction(p)
	if err != nil {
		return
	}
	if *r.Payload.TxHash != signedEncodedTxHash {
		err = fmt.Errorf("Transaction hash mismatch, expected %s got %s", signedEncodedTxHash, *r.Payload.TxHash)
	}
	return
}

// GetTopBlock get the top block of the chain
func (c *Client) GetTopBlock() (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	r, err := c.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// GetHeight get the height of the chain
func (c *Client) GetHeight() (height uint64, err error) {
	tb, err := c.GetTopBlock()
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

// GetCurrentKeyBlock get current key block
func (c *Client) GetCurrentKeyBlock() (kb *models.KeyBlock, err error) {
	r, err := c.External.GetCurrentKeyBlock(nil)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	kb = r.Payload
	return
}

// GetAccount retrieve an account by its address (public key)
// it is particularly useful to obtain the nonce for spending transactions
func (c *Client) GetAccount(accountID string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(accountID)
	r, err := c.External.GetAccountByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	account = r.Payload
	return
}

// GetNameEntryByName return the name entry
func (c *Client) GetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	p := external.NewGetNameEntryByNameParams().WithName(name)
	r, err := c.External.GetNameEntryByName(p)

	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}

	nameEntry = r.Payload
	return
}

// GetGenerationByHeight gets the keyblock and all its microblocks
func (c *Client) GetGenerationByHeight(height uint64) (g *models.Generation, err error) {
	p := external.NewGetGenerationByHeightParams().WithHeight(height)
	r, err := c.External.GetGenerationByHeight(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	g = r.Payload
	return
}

// GetMicroBlockTransactionsByHash get the transactions of a microblock
func (c *Client) GetMicroBlockTransactionsByHash(microBlockID string) (txs *models.GenericTxs, err error) {
	p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(microBlockID)
	r, err := c.External.GetMicroBlockTransactionsByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetMicroBlockHeaderByHash get the header of a micro block
func (c *Client) GetMicroBlockHeaderByHash(microBlockID string) (txs *models.MicroBlockHeader, err error) {
	p := external.NewGetMicroBlockHeaderByHashParams().WithHash(microBlockID)
	r, err := c.External.GetMicroBlockHeaderByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetKeyBlockByHash get a key block by its hash
func (c *Client) GetKeyBlockByHash(keyBlockID string) (txs *models.KeyBlock, err error) {
	p := external.NewGetKeyBlockByHashParams().WithHash(keyBlockID)
	r, err := c.External.GetKeyBlockByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetTransactionByHash get a transaction by it's hash
func (c *Client) GetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := c.External.GetTransactionByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	tx = r.Payload
	return
}

// GetOracleByPubkey get an oracle by it's public key
func (c *Client) GetOracleByPubkey(pubkey string) (oracle *models.RegisteredOracle, err error) {
	p := external.NewGetOracleByPubkeyParams().WithPubkey(pubkey)
	r, err := c.External.GetOracleByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracle = r.Payload
	return
}

// GetOracleQueriesByPubkey get a list of queries made to a particular oracle
func (c *Client) GetOracleQueriesByPubkey(pubkey string) (oracleQueries *models.OracleQueries, err error) {
	p := external.NewGetOracleQueriesByPubkeyParams().WithPubkey(pubkey)
	r, err := c.External.GetOracleQueriesByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	oracleQueries = r.Payload
	return
}

// GetContractByID gets a contract by ct_ ID
func (c *Client) GetContractByID(ctID string) (contract *models.ContractObject, err error) {
	p := external.NewGetContractParams().WithPubkey(ctID)
	r, err := c.External.GetContract(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	contract = r.Payload
	return
}
