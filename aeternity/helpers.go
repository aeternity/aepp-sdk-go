package aeternity

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

// GetHeightAccountNamer is used by Helper{} methods to describe the
// capabilities of whatever should be passed in as conn
type GetHeightAccountNamer interface {
	GetHeighter
	GetAccounter
	GetNameEntryByNamer
}

// getTransactionByHashHeighter is used by WaitForTransactionUntilHeight
type getTransactionByHashHeighter interface {
	GetTransactionByHasher
	GetHeighter
}

// HelpersInterface describes an interface for the helper functions GetTTLNonce
// and friends so they are mockable, without having to mock out the Node/network
// connection.
type HelpersInterface interface {
	GetTTL(offset uint64) (ttl uint64, err error)
	GetNextNonce(accountID string) (nextNonce uint64, err error)
	GetTTLNonce(accountID string, offset uint64) (height uint64, nonce uint64, err error)
	GetAccountsByName(name string) (addresses []string, err error)
	GetOraclesByName(name string) (oracleIDs []string, err error)
	GetContractsByName(name string) (contracts []string, err error)
	GetChannelsByName(name string) (channels []string, err error)
}

// Helpers is a struct to contain the GetTTLNonce helper functions and feed them
// with a node connection
type Helpers struct {
	Node GetHeightAccountNamer
}

// GetTTL returns the chain height + offset
func (h Helpers) GetTTL(offset uint64) (ttl uint64, err error) {
	height, err := h.Node.GetHeight()
	if err != nil {
		return
	}

	ttl = height + offset

	return
}

// GetNextNonce retrieves the current accountNonce and adds 1 to it for use in transaction building
func (h Helpers) GetNextNonce(accountID string) (nextNonce uint64, err error) {
	a, err := h.Node.GetAccount(accountID)
	if err != nil {
		return
	}
	nextNonce = *a.Nonce + 1
	return
}

// GetTTLNonce combines the commonly used together functions of GetTTL and GetNextNonce
func (h Helpers) GetTTLNonce(accountID string, offset uint64) (height uint64, nonce uint64, err error) {
	height, err = h.GetTTL(offset)
	if err != nil {
		return
	}

	nonce, err = h.GetNextNonce(accountID)
	if err != nil {
		return
	}
	return
}

// getAnythingByName is the underlying implementation of Get*ByName
func (h Helpers) getAnythingByName(name string, key string) (results []string, err error) {
	n, err := h.Node.GetNameEntryByName(name)
	if err != nil {
		return []string{}, err
	}
	for _, p := range n.Pointers {
		if *p.Key == key {
			results = append(results, *p.ID)
		}
	}
	return results, nil
}

// GetAccountsByName returns any account_pubkey entries that it finds in a name's Pointers.
func (h Helpers) GetAccountsByName(name string) (addresses []string, err error) {
	return h.getAnythingByName(name, "account_pubkey")
}

// GetOraclesByName returns any oracle_pubkey entries that it finds in a name's Pointers.
func (h Helpers) GetOraclesByName(name string) (oracleIDs []string, err error) {
	return h.getAnythingByName(name, "oracle_pubkey")
}

// GetContractsByName returns any contract_pubkey entries that it finds in a name's Pointers.
func (h Helpers) GetContractsByName(name string) (contracts []string, err error) {
	return h.getAnythingByName(name, "contract_pubkey")
}

// GetChannelsByName returns any channel entries that it finds in a name's Pointers.
func (h Helpers) GetChannelsByName(name string) (channels []string, err error) {
	return h.getAnythingByName(name, "channel")
}

// Context stores relevant context (node connection, account address) that one might not want to spell out each time one creates a transaction
type Context struct {
	Address string
	Helpers HelpersInterface
}

// NewContextFromURL is a convenience function that associates a Node with a
// Helper struct for you.
func NewContextFromURL(url string, address string) Context {
	h := Helpers{Node: NewNode(url, false)}
	return Context{Helpers: h, Address: address}
}

// SpendTx creates a spend transaction
func (c *Context) SpendTx(senderID string, recipientID string, amount, fee big.Int, payload []byte) (tx SpendTx, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	return NewSpendTx(senderID, recipientID, amount, fee, payload, txTTL, accountNonce), err
}

// NamePreclaimTx creates a name preclaim transaction and salt (required for claiming)
// It should return the Tx struct, not the base64 encoded RLP, to ease subsequent inspection.
func (c *Context) NamePreclaimTx(name string, fee big.Int) (tx NamePreclaimTx, nameSalt *big.Int, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// calculate the commitment and get the preclaim salt
	// since the salt is 32 bytes long, you must use a big.Int to convert it into an integer
	cm, nameSalt, err := generateCommitmentID(name)
	if err != nil {
		return
	}

	// build the transaction
	tx = NewNamePreclaimTx(c.Address, cm, fee, txTTL, accountNonce)
	if err != nil {
		return
	}

	return
}

// NameClaimTx creates a claim transaction
func (c *Context) NameClaimTx(name string, nameSalt big.Int, fee big.Int) (tx NameClaimTx, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	tx = NewNameClaimTx(c.Address, name, nameSalt, fee, txTTL, accountNonce)

	return tx, err
}

// NameUpdateTx perform a name update
func (c *Context) NameUpdateTx(name string, targetAddress string) (tx NameUpdateTx, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))
	absNameTTL, err := c.Helpers.GetTTL(Config.Client.Names.NameTTL)
	if err != nil {
		return NameUpdateTx{}, err
	}
	// create the transaction
	tx = NewNameUpdateTx(c.Address, encodedNameHash, []string{targetAddress}, absNameTTL, Config.Client.Names.ClientTTL, Config.Client.Names.UpdateFee, txTTL, accountNonce)

	return
}

// NameTransferTx transfer a name to another owner
func (c *Context) NameTransferTx(name string, recipientAddress string) (tx NameTransferTx, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameTransferTx(c.Address, encodedNameHash, recipientAddress, Config.Client.Fee, txTTL, accountNonce)
	return
}

// NameRevokeTx revoke a name
func (c *Context) NameRevokeTx(name string) (tx NameRevokeTx, err error) {
	txTTL, accountNonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameRevokeTx(c.Address, encodedNameHash, Config.Client.Fee, txTTL, accountNonce)
	return
}

// OracleRegisterTx create a new oracle
func (c *Context) OracleRegisterTx(querySpec, responseSpec string, queryFee big.Int, oracleTTLType, oracleTTLValue uint64, abiVersion uint16) (tx OracleRegisterTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return OracleRegisterTx{}, err
	}

	tx = NewOracleRegisterTx(c.Address, nonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, abiVersion, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleExtendTx extend the lifetime of an existing oracle
func (c *Context) OracleExtendTx(oracleID string, ttlType, ttlValue uint64) (tx OracleExtendTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return OracleExtendTx{}, err
	}

	tx = NewOracleExtendTx(oracleID, nonce, ttlType, ttlValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleQueryTx ask something of an oracle
func (c *Context) OracleQueryTx(OracleID, Query string, QueryFee big.Int, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64) (tx OracleQueryTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return OracleQueryTx{}, err
	}

	tx = NewOracleQueryTx(c.Address, nonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleRespondTx the oracle responds by sending this transaction
func (c *Context) OracleRespondTx(OracleID string, QueryID string, Response string, TTLType uint64, TTLValue uint64) (tx OracleRespondTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return OracleRespondTx{}, err
	}

	tx = NewOracleRespondTx(OracleID, nonce, QueryID, Response, TTLType, TTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// ContractCreateTx returns a transaction for creating a contract on the chain
func (c *Context) ContractCreateTx(Code string, CallData string, VMVersion, AbiVersion uint16, Deposit, Amount, Gas, GasPrice, Fee big.Int) (tx ContractCreateTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return ContractCreateTx{}, err
	}

	tx = NewContractCreateTx(c.Address, nonce, Code, VMVersion, AbiVersion, Deposit, Amount, Gas, GasPrice, Fee, ttl, CallData)
	return tx, nil
}

// ContractCallTx returns a transaction for calling a contract on the chain
func (c *Context) ContractCallTx(ContractID, CallData string, AbiVersion uint16, Amount, Gas, GasPrice, Fee big.Int) (tx ContractCallTx, err error) {
	ttl, nonce, err := c.Helpers.GetTTLNonce(c.Address, Config.Client.TTL)
	if err != nil {
		return ContractCallTx{}, err
	}

	tx = NewContractCallTx(c.Address, nonce, ContractID, Amount, Gas, GasPrice, AbiVersion, CallData, Fee, ttl)
	return tx, nil
}

// StoreAccountToKeyStoreFile store an account to a json file
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

// LoadAccountFromKeyStoreFile load file from the keystore
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

// GetWalletPath try to locate a wallet
func GetWalletPath(path string) (walletPath string, err error) {
	// if file exists then load the file
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		walletPath = path
		return
	}
	return
}

// SignEncodeTxStr sign and encode a transaction format as string (ex. tx_xyz)
func SignEncodeTxStr(kp *Account, tx string, networkID string) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
	txRaw, err := Decode(tx)
	if err != nil {
		fmt.Println("Error decoding tx from base64")
		os.Exit(1)
	}

	signedEncodedTx, signedEncodedTxHash, signature, err = SignEncodeTx(kp, txRaw, networkID)
	return
}

// VerifySignedTx verifies a tx_ with signature
func VerifySignedTx(accountID string, txSigned string, networkID string) (valid bool, err error) {
	txRawSigned, _ := Decode(txSigned)
	txRLP := DecodeRLPMessage(txRawSigned)

	// RLP format of signed signature: [[Tag], [Version], [Signatures...], [Transaction]]
	tx := txRLP[3].([]byte)
	txSignature := txRLP[2].([]interface{})[0].([]byte)

	msg := append([]byte(networkID), tx...)

	valid, err = Verify(accountID, msg, txSignature)
	if err != nil {
		return
	}
	return
}

// WaitForTransactionUntilHeight waits for a transaction until heightLimit (inclusive) is reached
func WaitForTransactionUntilHeight(c getTransactionByHashHeighter, txHash string, untilHeight uint64) (blockHeight uint64, blockHash string, err error) {
	var nodeHeight uint64
	for nodeHeight <= untilHeight {
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
		time.Sleep(time.Millisecond * time.Duration(Config.Tuning.ChainPollInteval))
	}
	return 0, "", fmt.Errorf("It is already height %v and %v still isn't in a block", nodeHeight, txHash)
}

// BroadcastTransaction differs from Client.PostTransaction() in that the latter just handles
// the HTTP request via swagger, the former recalculates the txhash and compares it to the node's
//  response after POSTing the transaction.
func BroadcastTransaction(c PostTransactioner, txSignedBase64 string) (err error) {
	// Get back to RLP to calculate txhash
	txRLP, _ := Decode(txSignedBase64)

	// calculate the hash of the decoded txRLP
	rlpTxHashRaw, _ := hash(txRLP)
	// base58/64 encode the hash with the th_ prefix
	signedEncodedTxHash := Encode(PrefixTransactionHash, rlpTxHashRaw)

	// send it to the network
	err = c.PostTransaction(txSignedBase64, signedEncodedTxHash)
	return
}
