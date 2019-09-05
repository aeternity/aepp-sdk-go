package aeternity

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	rlp "github.com/randomshinichi/rlpae"
)

// GetTTLFunc defines a function that will return an appropriate TTL for a
// transaction.
type GetTTLFunc func(offset uint64) (ttl uint64, err error)

// GetNextNonceFunc defines a function that will return an unused account nonce
// for making a transaction.
type GetNextNonceFunc func(accountID string) (nonce uint64, err error)

// GetTTLNonceFunc describes a function that combines the roles of GetTTLFunc
// and GetNextNonceFunc
type GetTTLNonceFunc func(address string, offset uint64) (ttl, nonce uint64, err error)

// GenerateGetTTL returns the chain height + offset
func GenerateGetTTL(n GetHeighter) GetTTLFunc {
	return func(offset uint64) (ttl uint64, err error) {
		height, err := n.GetHeight()
		if err != nil {
			return
		}
		ttl = height + offset
		return
	}
}

// GenerateGetNextNonce retrieves the current accountNonce and adds 1 to it for
// use in transaction building
func GenerateGetNextNonce(n GetAccounter) GetNextNonceFunc {
	return func(accountID string) (nextNonce uint64, err error) {
		a, err := n.GetAccount(accountID)
		if err != nil {
			return
		}
		nextNonce = *a.Nonce + 1
		return
	}
}

// GenerateGetTTLNonce combines the commonly used together functions of GetTTL
// and GetNextNonce
func GenerateGetTTLNonce(ttlFunc GetTTLFunc, nonceFunc GetNextNonceFunc) GetTTLNonceFunc {
	return func(accountID string, offset uint64) (ttl, nonce uint64, err error) {
		ttl, err = ttlFunc(offset)
		if err != nil {
			return
		}
		nonce, err = nonceFunc(accountID)
		if err != nil {
			return
		}
		return
	}
}

// GetAnythingByNameFunc describes a function that returns lookup results for a
// AENS name
type GetAnythingByNameFunc func(name, key string) (results []string, err error)

// GenerateGetAnythingByName is the underlying implementation of Get*ByName
func GenerateGetAnythingByName(n GetNameEntryByNamer) GetAnythingByNameFunc {
	return func(name string, key string) (results []string, err error) {
		nameEntry, err := n.GetNameEntryByName(name)
		if err != nil {
			return []string{}, err
		}
		for _, p := range nameEntry.Pointers {
			if *p.Key == key {
				results = append(results, *p.ID)
			}
		}
		return results, nil
	}
}

// GetAccountsByName returns any account_pubkey entries that it finds in a
// name's Pointers.
func GetAccountsByName(n GetAnythingByNameFunc, name string) (addresses []string, err error) {
	return n(name, "account_pubkey")
}

// GetOraclesByName returns any oracle_pubkey entries that it finds in a name's
// Pointers.
func GetOraclesByName(n GetAnythingByNameFunc, name string) (oracleIDs []string, err error) {
	return n(name, "oracle_pubkey")
}

// GetContractsByName returns any contract_pubkey entries that it finds in a
// name's Pointers.
func GetContractsByName(n GetAnythingByNameFunc, name string) (contracts []string, err error) {
	return n(name, "contract_pubkey")
}

// GetChannelsByName returns any channel entries that it finds in a name's
// Pointers.
func GetChannelsByName(n GetAnythingByNameFunc, name string) (channels []string, err error) {
	return n(name, "channel")
}

// Context is a struct that eases transaction creation. Specifically, the role
// of Context is to automate/hide the tedious details of transaction creation,
// such as filling in an unused account nonce and an appropriate TTL, so that
// the transaction creator only has to worry about the details relevant to the
// task at hand.
type Context struct {
	GetTTL      GetTTLFunc
	GetNonce    GetNextNonceFunc
	GetTTLNonce GetTTLNonceFunc
	Address     string
}

// NewContextFromURL is a convenience function that creates a Context and its
// TTL/Nonce closures from a URL
func NewContextFromURL(url string, address string, debug bool) (ctx *Context, node *Node) {
	node = NewNode(url, debug)
	return NewContextFromNode(node, address), node
}

// NewContextFromNode is a convenience function that creates a Context and its
// TTL/Nonce closures from a Node instance
func NewContextFromNode(node *Node, address string) (ctx *Context) {
	ttlFunc := GenerateGetTTL(node)
	nonceFunc := GenerateGetNextNonce(node)
	ttlNonceFunc := GenerateGetTTLNonce(ttlFunc, nonceFunc)
	ctx = &Context{
		GetTTL:      ttlFunc,
		GetNonce:    nonceFunc,
		GetTTLNonce: ttlNonceFunc,
		Address:     address,
	}
	return
}

// SpendTx creates a spend transaction, filling in the account nonce and
// transaction TTL automatically.
func (c *Context) SpendTx(senderID string, recipientID string, amount, fee big.Int, payload []byte) (tx *SpendTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	return NewSpendTx(senderID, recipientID, amount, fee, payload, txTTL, accountNonce), err
}

// NamePreclaimTx creates a name preclaim transaction, filling in the account
// nonce and transaction TTL automatically. It also generates a commitment ID
// and salt, later used to claim the name.
func (c *Context) NamePreclaimTx(name string, fee big.Int) (tx *NamePreclaimTx, nameSalt *big.Int, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// calculate the commitment and get the preclaim salt since the salt is 32
	// bytes long, you must use a big.Int to convert it into an integer
	cm, nameSalt, err := generateCommitmentID(name)
	if err != nil {
		return
	}

	// build the transaction
	tx = NewNamePreclaimTx(c.Address, cm, fee, txTTL, accountNonce)

	return
}

// NameClaimTx creates a claim transaction, filling in the account nonce and
// transaction TTL automatically.
func (c *Context) NameClaimTx(name string, nameSalt big.Int, fee big.Int) (tx *NameClaimTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	tx = NewNameClaimTx(c.Address, name, nameSalt, fee, txTTL, accountNonce)

	return tx, err
}

// NameUpdateTx creates a name update transaction, filling in the account nonce
// and transaction TTL automatically.
func (c *Context) NameUpdateTx(name string, targetAddress string) (tx *NameUpdateTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))
	absNameTTL, err := c.GetTTL(Config.Client.Names.NameTTL)
	if err != nil {
		return
	}
	// create the transaction
	tx = NewNameUpdateTx(c.Address, encodedNameHash, []string{targetAddress}, absNameTTL, Config.Client.Names.ClientTTL, Config.Client.Names.UpdateFee, txTTL, accountNonce)

	return
}

// NameTransferTx creates a name transfer transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) NameTransferTx(name string, recipientAddress string) (tx *NameTransferTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameTransferTx(c.Address, encodedNameHash, recipientAddress, Config.Client.Fee, txTTL, accountNonce)
	return
}

// NameRevokeTx creates a name revoke transaction, filling in the account nonce
// and transaction TTL automatically.
func (c *Context) NameRevokeTx(name string) (tx *NameRevokeTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameRevokeTx(c.Address, encodedNameHash, Config.Client.Fee, txTTL, accountNonce)
	return
}

// OracleRegisterTx creates an oracle register transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) OracleRegisterTx(querySpec, responseSpec string, queryFee big.Int, oracleTTLType, oracleTTLValue uint64, VMVersion uint16) (tx *OracleRegisterTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewOracleRegisterTx(c.Address, nonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, VMVersion, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleExtendTx creates an oracle extend transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) OracleExtendTx(oracleID string, ttlType, ttlValue uint64) (tx *OracleExtendTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewOracleExtendTx(oracleID, nonce, ttlType, ttlValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleQueryTx creates an oracle query transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) OracleQueryTx(OracleID, Query string, QueryFee big.Int, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64) (tx *OracleQueryTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewOracleQueryTx(c.Address, nonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleRespondTx creates an oracle response transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) OracleRespondTx(OracleID string, QueryID string, Response string, TTLType uint64, TTLValue uint64) (tx *OracleRespondTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewOracleRespondTx(OracleID, nonce, QueryID, Response, TTLType, TTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// ContractCreateTx creates a contract create transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) ContractCreateTx(Code string, CallData string, VMVersion, AbiVersion uint16, Deposit, Amount, Gas, GasPrice, Fee big.Int) (tx *ContractCreateTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewContractCreateTx(c.Address, nonce, Code, VMVersion, AbiVersion, Deposit, Amount, Gas, GasPrice, Fee, ttl, CallData)
	return tx, nil
}

// ContractCallTx creates a contract call transaction,, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) ContractCallTx(ContractID, CallData string, AbiVersion uint16, Amount, Gas, GasPrice, Fee big.Int) (tx *ContractCallTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	tx = NewContractCallTx(c.Address, nonce, ContractID, Amount, Gas, GasPrice, AbiVersion, CallData, Fee, ttl)
	return tx, nil
}

// StoreAccountToKeyStoreFile saves an encrypted Account to a JSON file
func StoreAccountToKeyStoreFile(account *Account, password, walletName string) (filePath string, err error) {
	// keystore will be saved in current directory
	basePath, _ := os.Getwd()

	// generate the keystore file
	jks, err := KeystoreSeal(account, password)
	if err != nil {
		return
	}
	// build the wallet path
	filePath = filepath.Join(basePath, keyFileName(account.Address))
	if len(walletName) > 0 {
		filePath = filepath.Join(basePath, walletName)
	}
	// write the file to disk
	err = ioutil.WriteFile(filePath, jks, 0600)
	return
}

// LoadAccountFromKeyStoreFile loads an encrypted Account from a JSON file
func LoadAccountFromKeyStoreFile(keyFile, password string) (account *Account, err error) {
	// find out the real path of the wallet
	filePath, err := GetWalletPath(keyFile)
	if err != nil {
		return
	}
	// load the json file
	jks, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	// decrypt keystore
	account, err = KeystoreOpen(jks, password)
	return
}

// GetWalletPath checks if a file exists at the specified path.
func GetWalletPath(path string) (walletPath string, err error) {
	// if file exists then load the file
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		walletPath = path
		return
	}
	return
}

// VerifySignedTx verifies the signature of a signed transaction, in its RLP
// serialized, base64 encoded tx_ form.
//
// The network ID is also used when calculating the signature, so the network ID
// that the transaction was intended for should be provided too.
func VerifySignedTx(accountID string, txSigned string, networkID string) (valid bool, err error) {
	txRawSigned, _ := Decode(txSigned)
	txRLP := DecodeRLPMessage(txRawSigned)

	// RLP format of signed signature: [[Tag], [Version], [Signatures...],
	// [Transaction]]
	tx := txRLP[3].([]byte)
	txSignature := txRLP[2].([]interface{})[0].([]byte)

	msg := append([]byte(networkID), tx...)

	valid, err = Verify(accountID, msg, txSignature)
	if err != nil {
		return
	}
	return
}

// getTransactionByHashHeighter is used by WaitForTransactionForXBlocks to
// specify that the node/mock node passed in should support
// GetTransactionByHash() and GetHeight()
type getTransactionByHashHeighter interface {
	GetTransactionByHasher
	GetHeighter
}

// WaitForTransactionForXBlocks blocks until a transaction has been mined or X
// blocks have gone by, after which it returns an error. The node polling
// interval can be configured with Config.Tuning.ChainPollInterval.
func WaitForTransactionForXBlocks(c getTransactionByHashHeighter, txHash string, x uint64) (blockHeight uint64, blockHash string, err error) {
	nodeHeight, err := c.GetHeight()
	if err != nil {
		return
	}
	endHeight := nodeHeight + x
	for nodeHeight <= endHeight {
		nodeHeight, err = c.GetHeight()
		if err != nil {
			return 0, "", err
		}
		tx, err := c.GetTransactionByHash(txHash)
		if err != nil {
			return 0, "", err
		}

		if tx.BlockHeight.LargerThanZero() {
			bh := big.Int(tx.BlockHeight)
			return bh.Uint64(), *tx.BlockHash, nil
		}
		time.Sleep(time.Millisecond * time.Duration(Config.Tuning.ChainPollInterval))
	}
	return 0, "", fmt.Errorf("%v blocks have gone by and %v still isn't in a block", x, txHash)
}

// SignBroadcastTransaction signs a transaction and broadcasts it to a node.
func SignBroadcastTransaction(tx rlp.Encoder, signingAccount *Account, n *Node, networkID string) (signedTxStr, hash, signature string, err error) {
	signedTx, hash, signature, err := SignHashTx(signingAccount, tx, networkID)
	if err != nil {
		return
	}

	signedTxStr, err = SerializeTx(signedTx)
	if err != nil {
		return
	}

	err = n.PostTransaction(signedTxStr, hash)
	if err != nil {
		return
	}
	return
}

// SignBroadcastWaitTransaction is a convenience function that combines
// SignBroadcastTransaction and WaitForTransactionForXBlocks.
func SignBroadcastWaitTransaction(tx rlp.Encoder, signingAccount *Account, n *Node, networkID string, x uint64) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	signedTxStr, hash, signature, err = SignBroadcastTransaction(tx, signingAccount, n, networkID)
	if err != nil {
		return
	}
	blockHeight, blockHash, err = WaitForTransactionForXBlocks(n, hash, x)
	return
}
