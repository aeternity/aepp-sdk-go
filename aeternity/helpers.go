package aeternity

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/aeternity/aepp-sdk-go/generated/models"
)

type getHeighter interface {
	GetHeight() (uint64, error)
}

type getAccounter interface {
	GetAccount(string) (*models.Account, error)
}

type getHeightAccounter interface {
	getHeighter
	getAccounter
}

// GetTTL returns the chain height + offset
func GetTTL(c getHeighter, offset uint64) (ttl uint64, err error) {
	height, err := c.GetHeight()
	if err != nil {
		return
	}

	ttl = height + offset

	return
}

// GetNextNonce retrieves the current accountNonce and adds 1 to it for use in transaction building
func GetNextNonce(c getAccounter, accountID string) (nextNonce uint64, err error) {
	a, err := c.GetAccount(accountID)
	if err != nil {
		return
	}
	nextNonce = *a.Nonce + 1
	return
}

// GetTTLNonce combines the commonly used together functions of GetTTL and GetNextNonce
func GetTTLNonce(c getHeightAccounter, accountID string, offset uint64) (height uint64, nonce uint64, err error) {
	height, err = GetTTL(c, offset)
	if err != nil {
		return
	}

	nonce, err = GetNextNonce(c, accountID)
	if err != nil {
		return
	}
	return
}

type getTransactionByHasher interface {
	GetTransactionByHash(string) (*models.GenericSignedTx, error)
}
type getTransactionByHashHeighter interface {
	getTransactionByHasher
	getHeighter
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

type transactionPoster interface {
	PostTransaction(string, string) (err error)
}

// BroadcastTransaction differs from Client.PostTransaction() in that the latter just handles
// the HTTP request via swagger, the former recalculates the txhash and compares it to the node's
//  response after POSTing the transaction.
func BroadcastTransaction(c transactionPoster, txSignedBase64 string) (err error) {
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

// PrintGenerationByHeight utility function to print a generation by it's height
// TODO this belongs in cmd and needs to be tested with error cases
func (c *Client) PrintGenerationByHeight(height uint64) {
	r, err := c.GetGenerationByHeight(height)
	if err == nil {
		PrintObject("generation", r)
		// search for transaction in the microblocks
		for _, mbh := range r.MicroBlocks {
			// get the microblok
			mbhs := fmt.Sprint(mbh)
			r, err := c.GetMicroBlockTransactionsByHash(mbhs)
			if err != nil {
				Pp("Error:", err)
			}
			// go through all the hashes
			for _, btx := range r.Transactions {
				p, err := c.GetTransactionByHash(fmt.Sprint(btx.Hash))
				if err == nil {
					PrintObject("transaction", p)
				}
			}
		}
	} else {
		fmt.Println("Something went wrong in PrintGenerationByHeight")
	}
}

// NamePreclaimTx creates a name preclaim transaction and salt (required for claiming)
// It should return the Tx struct, not the base64 encoded RLP, to ease subsequent inspection.
func (n *Aens) NamePreclaimTx(name string, fee big.Int) (tx NamePreclaimTx, nameSalt *big.Int, err error) {
	txTTL, accountNonce, err := GetTTLNonce(n.Client, n.Account.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// calculate the commitment and get the preclaim salt
	// since the salt is 32 bytes long, you must use a big.Int to convert it into an integer
	cm, nameSalt, err := generateCommitmentID(name)
	if err != nil {
		return NamePreclaimTx{}, new(big.Int), err
	}

	// build the transaction
	tx = NewNamePreclaimTx(n.Account.Address, cm, fee, txTTL, accountNonce)
	if err != nil {
		return NamePreclaimTx{}, new(big.Int), err
	}

	return
}

// NameClaimTx creates a claim transaction
func (n *Aens) NameClaimTx(name string, nameSalt big.Int, fee big.Int) (tx NameClaimTx, err error) {
	txTTL, accountNonce, err := GetTTLNonce(n.Client, n.Account.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	tx = NewNameClaimTx(n.Account.Address, name, nameSalt, fee, txTTL, accountNonce)

	return tx, err
}

// NameUpdateTx perform a name update
func (n *Aens) NameUpdateTx(name string, targetAddress string) (tx NameUpdateTx, err error) {
	txTTL, accountNonce, err := GetTTLNonce(n.Client, n.Account.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))
	absNameTTL, err := GetTTL(n.Client, Config.Client.Names.NameTTL)
	if err != nil {
		return NameUpdateTx{}, err
	}
	// create the transaction
	tx = NewNameUpdateTx(n.Account.Address, encodedNameHash, []string{targetAddress}, absNameTTL, Config.Client.Names.ClientTTL, Config.Client.Names.UpdateFee, txTTL, accountNonce)

	return
}

// NameTransferTx transfer a name to another owner
func (n *Aens) NameTransferTx(name string, recipientAddress string) (tx NameTransferTx, err error) {
	txTTL, accountNonce, err := GetTTLNonce(n.Client, n.Account.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameTransferTx(n.Account.Address, encodedNameHash, recipientAddress, Config.Client.Fee, txTTL, accountNonce)
	return
}

// NameRevokeTx revoke a name
func (n *Aens) NameRevokeTx(name string, recipientAddress string) (tx NameRevokeTx, err error) {
	txTTL, accountNonce, err := GetTTLNonce(n.Client, n.Account.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))

	tx = NewNameRevokeTx(n.Account.Address, encodedNameHash, Config.Client.Fee, txTTL, accountNonce)
	return
}

// OracleRegisterTx create a new oracle
func (o *Oracle) OracleRegisterTx(querySpec, responseSpec string, queryFee big.Int, oracleTTLType, oracleTTLValue uint64, abiVersion uint16) (tx OracleRegisterTx, err error) {
	ttl, nonce, err := GetTTLNonce(o.Client, o.Account.Address, Config.Client.TTL)
	if err != nil {
		return OracleRegisterTx{}, err
	}

	tx = NewOracleRegisterTx(o.Account.Address, nonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, abiVersion, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleExtendTx extend the lifetime of an existing oracle
func (o *Oracle) OracleExtendTx(oracleID string, ttlType, ttlValue uint64) (tx OracleExtendTx, err error) {
	ttl, nonce, err := GetTTLNonce(o.Client, o.Account.Address, Config.Client.TTL)
	if err != nil {
		return OracleExtendTx{}, err
	}

	tx = NewOracleExtendTx(oracleID, nonce, ttlType, ttlValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleQueryTx ask something of an oracle
func (o *Oracle) OracleQueryTx(OracleID, Query string, QueryFee big.Int, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64) (tx OracleQueryTx, err error) {
	ttl, nonce, err := GetTTLNonce(o.Client, o.Account.Address, Config.Client.TTL)
	if err != nil {
		return OracleQueryTx{}, err
	}

	tx = NewOracleQueryTx(o.Account.Address, nonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// OracleRespondTx the oracle responds by sending this transaction
func (o *Oracle) OracleRespondTx(OracleID string, QueryID string, Response string, TTLType uint64, TTLValue uint64) (tx OracleRespondTx, err error) {
	ttl, nonce, err := GetTTLNonce(o.Client, o.Account.Address, Config.Client.TTL)
	if err != nil {
		return OracleRespondTx{}, err
	}

	tx = NewOracleRespondTx(OracleID, nonce, QueryID, Response, TTLType, TTLValue, Config.Client.Fee, ttl)
	return tx, nil
}

// ContractCreateTx returns a transaction for creating a contract on the chain
func (c *Contract) ContractCreateTx(Code string, CallData string, VMVersion, AbiVersion uint16, Deposit, Amount, Gas, GasPrice, Fee big.Int) (tx ContractCreateTx, err error) {
	ttl, nonce, err := GetTTLNonce(c.Client, c.Account.Address, Config.Client.TTL)
	if err != nil {
		return ContractCreateTx{}, err
	}

	tx = NewContractCreateTx(c.Account.Address, nonce, Code, VMVersion, AbiVersion, Deposit, Amount, Gas, GasPrice, Fee, ttl, CallData)
	return tx, nil
}

// ContractCallTx returns a transaction for calling a contract on the chain
func (c *Contract) ContractCallTx(ContractID, CallData string, AbiVersion uint16, Amount, Gas, GasPrice, Fee big.Int) (tx ContractCallTx, err error) {
	ttl, nonce, err := GetTTLNonce(c.Client, c.Account.Address, Config.Client.TTL)
	if err != nil {
		return ContractCallTx{}, err
	}

	tx = NewContractCallTx(c.Account.Address, nonce, ContractID, Amount, Gas, GasPrice, AbiVersion, CallData, Fee, ttl)
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
