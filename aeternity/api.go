package aeternity

import (
	"fmt"
	"reflect"

	"github.com/aeternity/aepp-sdk-go/v5/swagguard/node/client/external"
	"github.com/aeternity/aepp-sdk-go/v5/swagguard/node/models"
)

// getErrorReason parses the external.*BadRequest structs returned by the
// Swagger API endpoints and gets to the Reason field of the wrapped
// models.Error struct, returning the reason why the call failed.
func getErrorReason(v interface{}) (msg string) {
	var p func(v reflect.Value) (msg string)
	p = func(v reflect.Value) (msg string) {
		switch v.Kind() {
		// If it is a pointer we need to unwrap and call once again
		case reflect.Ptr:
			if v.IsValid() {
				msg = p(v.Elem())
			}
		case reflect.Struct:
			if v.Type() == reflect.TypeOf(models.Error{}) {
				msg = fmt.Sprint(reflect.Indirect(v.FieldByName("Reason")))
				break
			}
			for i := 0; i < v.NumField(); i++ {
				msg = p(v.Field(i))
			}
		}
		return
	}
	msg = p(reflect.ValueOf(v))
	if len(msg) == 0 {
		msg = fmt.Sprint(v)
	}
	return
}

// NodeInterface collects together all the interfaces defined in this file to
// describe the capabilities of a generic connection to an aeternity (or mock)
// node
type NodeInterface interface {
	GetAccounter
	GetTopBlocker
	GetStatuser
	GetHeighter
	GetKeyBlockByHasher
	GetOracleByPubkeyer
	GetGenerationByHeighter
	GetTransactionByHasher
	GetNameEntryByNamer
	GetMicroBlockHeaderByHasher
	GetMicroBlockTransactionsByHasher
	PostTransactioner
}

// GetStatuser guarantees that one can run a GetStatus() method on the
// mocked/real network connection
type GetStatuser interface {
	GetStatus() (*models.Status, error)
}

// GetStatus returns a node's Status.
func (c *Node) GetStatus() (status *models.Status, err error) {
	r, err := c.External.GetStatus(nil)
	if err != nil {
		return
	}
	status = r.Payload
	return
}

// PostTransactioner guarantees that one can run a PostTransaction() method on
// the mocked/real network connection
type PostTransactioner interface {
	PostTransaction(string, string) error
}

// PostTransaction post transaction
func (c *Node) PostTransaction(signedEncodedTx, signedEncodedTxHash string) (err error) {
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

// GetTopBlocker guarantees that one can run a GetTopBlock() method on the
// mocked/real network connection
type GetTopBlocker interface {
	GetTopBlock() (*models.KeyBlockOrMicroBlockHeader, error)
}

// GetTopBlock get the top block of the chain
func (c *Node) GetTopBlock() (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	r, err := c.External.GetTopBlock(nil)
	if err != nil {
		return
	}
	kb = r.Payload
	return
}

// GetHeighter guarantees that one can run a GetHeight() method on the
// mocked/real network connection
type GetHeighter interface {
	GetHeight() (uint64, error)
}

// GetHeight get the height of the chain
func (c *Node) GetHeight() (height uint64, err error) {
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
func (c *Node) GetCurrentKeyBlock() (kb *models.KeyBlock, err error) {
	r, err := c.External.GetCurrentKeyBlock(nil)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	kb = r.Payload
	return
}

// GetAccounter guarantees that one can run a GetAccount() method on the
// mocked/real network connection
type GetAccounter interface {
	GetAccount(accountID string) (*models.Account, error)
}

// GetAccount retrieve an account by its address (public key) it is particularly
// useful to obtain the nonce for spending transactions
func (c *Node) GetAccount(accountID string) (account *models.Account, err error) {
	p := external.NewGetAccountByPubkeyParams().WithPubkey(accountID)
	r, err := c.External.GetAccountByPubkey(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	account = r.Payload
	return
}

// GetNameEntryByNamer guarantees that one can run a GetNameEntryByName() method
// on the mocked/real network connection
type GetNameEntryByNamer interface {
	GetNameEntryByName(string) (*models.NameEntry, error)
}

// GetNameEntryByName return the name entry
func (c *Node) GetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	p := external.NewGetNameEntryByNameParams().WithName(name)
	r, err := c.External.GetNameEntryByName(p)

	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}

	nameEntry = r.Payload
	return
}

// GetGenerationByHeighter guarantees that one can run a GetGenerationByHeight()
// method on the mocked/real network connection
type GetGenerationByHeighter interface {
	GetGenerationByHeight(height uint64) (*models.Generation, error)
}

// GetGenerationByHeight gets the keyblock and all its microblocks
func (c *Node) GetGenerationByHeight(height uint64) (g *models.Generation, err error) {
	p := external.NewGetGenerationByHeightParams().WithHeight(height)
	r, err := c.External.GetGenerationByHeight(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	g = r.Payload
	return
}

// GetMicroBlockTransactionsByHasher guarantees that one can run a
// GetMicroBlockTransactionsByHash() method on the mocked/real network
// connection
type GetMicroBlockTransactionsByHasher interface {
	GetMicroBlockTransactionsByHash(string) (*models.GenericTxs, error)
}

// GetMicroBlockTransactionsByHash get the transactions of a microblock
func (c *Node) GetMicroBlockTransactionsByHash(microBlockID string) (txs *models.GenericTxs, err error) {
	p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(microBlockID)
	r, err := c.External.GetMicroBlockTransactionsByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetMicroBlockHeaderByHasher guarantees that one can run a
// GetMicroBlockHeaderByHash() method on the mocked/real network connection
type GetMicroBlockHeaderByHasher interface {
	GetMicroBlockHeaderByHash(string) (*models.MicroBlockHeader, error)
}

// GetMicroBlockHeaderByHash get the header of a micro block
func (c *Node) GetMicroBlockHeaderByHash(microBlockID string) (txs *models.MicroBlockHeader, err error) {
	p := external.NewGetMicroBlockHeaderByHashParams().WithHash(microBlockID)
	r, err := c.External.GetMicroBlockHeaderByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetKeyBlockByHasher guarantees that one can run a GetKeyBlockByHash() method
// on the mocked/real network connection
type GetKeyBlockByHasher interface {
	GetKeyBlockByHash(string) (*models.KeyBlock, error)
}

// GetKeyBlockByHash get a key block by its hash
func (c *Node) GetKeyBlockByHash(keyBlockID string) (txs *models.KeyBlock, err error) {
	p := external.NewGetKeyBlockByHashParams().WithHash(keyBlockID)
	r, err := c.External.GetKeyBlockByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	txs = r.Payload
	return
}

// GetTransactionByHasher guarantees that one can run a GetTransactionByHash()
// method on the mocked/real network connection
type GetTransactionByHasher interface {
	GetTransactionByHash(string) (*models.GenericSignedTx, error)
}

// GetTransactionByHash get a transaction by it's hash
func (c *Node) GetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	p := external.NewGetTransactionByHashParams().WithHash(txHash)
	r, err := c.External.GetTransactionByHash(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	tx = r.Payload
	return
}

// GetOracleByPubkeyer guarantees that one can run a GetOracleByPubkey() method
// on the mocked/real network connection
type GetOracleByPubkeyer interface {
	GetOracleByPubkey(string) (*models.RegisteredOracle, error)
}

// GetOracleByPubkey get an oracle by it's public key
func (c *Node) GetOracleByPubkey(pubkey string) (oracle *models.RegisteredOracle, err error) {
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
func (c *Node) GetOracleQueriesByPubkey(pubkey string) (oracleQueries *models.OracleQueries, err error) {
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
func (c *Node) GetContractByID(ctID string) (contract *models.ContractObject, err error) {
	p := external.NewGetContractParams().WithPubkey(ctID)
	r, err := c.External.GetContract(p)
	if err != nil {
		err = fmt.Errorf("Error: %v", getErrorReason(err))
		return
	}
	contract = r.Payload
	return
}
