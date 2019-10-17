package aeternity

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/binary"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/naet"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
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
func GenerateGetTTL(n naet.GetHeighter) GetTTLFunc {
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
func GenerateGetNextNonce(n naet.GetAccounter) GetNextNonceFunc {
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
func GenerateGetAnythingByName(n naet.GetNameEntryByNamer) GetAnythingByNameFunc {
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
func NewContextFromURL(url string, address string, debug bool) (ctx *Context, node *naet.Node) {
	node = naet.NewNode(url, debug)
	return NewContextFromNode(node, address), node
}

// NewContextFromNode is a convenience function that creates a Context and its
// TTL/Nonce closures from a Node instance
func NewContextFromNode(node *naet.Node, address string) (ctx *Context) {
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
func (c *Context) SpendTx(senderID string, recipientID string, amount, fee *big.Int, payload []byte) (tx *transactions.SpendTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	return transactions.NewSpendTx(senderID, recipientID, amount, fee, payload, txTTL, accountNonce), err
}

// NamePreclaimTx creates a name preclaim transaction, filling in the account
// nonce and transaction TTL automatically. It also generates a commitment ID
// and salt, later used to claim the name.
func (c *Context) NamePreclaimTx(name string, fee *big.Int) (tx *transactions.NamePreclaimTx, nameSalt *big.Int, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
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
	tx = transactions.NewNamePreclaimTx(c.Address, cm, fee, txTTL, accountNonce)

	return
}

// NameClaimTx creates a claim transaction, filling in the account nonce and
// transaction TTL automatically.
func (c *Context) NameClaimTx(name string, nameSalt, nameFee, fee *big.Int) (tx *transactions.NameClaimTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	tx = transactions.NewNameClaimTx(c.Address, name, nameSalt, nameFee, fee, txTTL, accountNonce)

	return tx, err
}

// NameUpdateTx creates a name update transaction, filling in the account nonce
// and transaction TTL automatically.
func (c *Context) NameUpdateTx(name string, targetAddress string) (tx *transactions.NameUpdateTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	nm, err := transactions.NameID(name)
	if err != nil {
		return
	}

	absNameTTL, err := c.GetTTL(config.Client.Names.NameTTL)
	if err != nil {
		return
	}
	// create the transaction
	tx = transactions.NewNameUpdateTx(c.Address, nm, []string{targetAddress}, absNameTTL, config.Client.Names.ClientTTL, config.Client.Fee, txTTL, accountNonce)

	return
}

// NameTransferTx creates a name transfer transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) NameTransferTx(name string, recipientAddress string) (tx *transactions.NameTransferTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	nm, err := transactions.NameID(name)
	if err != nil {
		return
	}

	tx = transactions.NewNameTransferTx(c.Address, nm, recipientAddress, config.Client.Fee, txTTL, accountNonce)
	return
}

// NameRevokeTx creates a name revoke transaction, filling in the account nonce
// and transaction TTL automatically.
func (c *Context) NameRevokeTx(name string) (tx *transactions.NameRevokeTx, err error) {
	txTTL, accountNonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	nm, err := transactions.NameID(name)
	if err != nil {
		return
	}

	tx = transactions.NewNameRevokeTx(c.Address, nm, config.Client.Fee, txTTL, accountNonce)
	return
}

// OracleRegisterTx creates an oracle register transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) OracleRegisterTx(querySpec, responseSpec string, queryFee *big.Int, oracleTTLType, oracleTTLValue uint64, VMVersion uint16) (tx *transactions.OracleRegisterTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewOracleRegisterTx(c.Address, nonce, querySpec, responseSpec, queryFee, oracleTTLType, oracleTTLValue, VMVersion, config.Client.Fee, ttl)
	return tx, nil
}

// OracleExtendTx creates an oracle extend transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) OracleExtendTx(oracleID string, ttlType, ttlValue uint64) (tx *transactions.OracleExtendTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewOracleExtendTx(oracleID, nonce, ttlType, ttlValue, config.Client.Fee, ttl)
	return tx, nil
}

// OracleQueryTx creates an oracle query transaction, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) OracleQueryTx(OracleID, Query string, QueryFee *big.Int, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue uint64) (tx *transactions.OracleQueryTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewOracleQueryTx(c.Address, nonce, OracleID, Query, QueryFee, QueryTTLType, QueryTTLValue, ResponseTTLType, ResponseTTLValue, config.Client.Fee, ttl)
	return tx, nil
}

// OracleRespondTx creates an oracle response transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) OracleRespondTx(OracleID string, QueryID string, Response string, TTLType uint64, TTLValue uint64) (tx *transactions.OracleRespondTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewOracleRespondTx(OracleID, nonce, QueryID, Response, TTLType, TTLValue, config.Client.Fee, ttl)
	return tx, nil
}

// ContractCreateTx creates a contract create transaction, filling in the
// account nonce and transaction TTL automatically.
func (c *Context) ContractCreateTx(Code string, CallData string, VMVersion, AbiVersion uint16, Deposit, Amount, GasLimit, Fee *big.Int) (tx *transactions.ContractCreateTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewContractCreateTx(c.Address, nonce, Code, VMVersion, AbiVersion, Deposit, Amount, GasLimit, config.Client.GasPrice, Fee, ttl, CallData)
	return tx, nil
}

// ContractCallTx creates a contract call transaction,, filling in the account
// nonce and transaction TTL automatically.
func (c *Context) ContractCallTx(ContractID, CallData string, AbiVersion uint16, Amount, GasLimit, GasPrice, Fee *big.Int) (tx *transactions.ContractCallTx, err error) {
	ttl, nonce, err := c.GetTTLNonce(c.Address, config.Client.TTL)
	if err != nil {
		return
	}

	tx = transactions.NewContractCallTx(c.Address, nonce, ContractID, Amount, GasLimit, GasPrice, AbiVersion, CallData, Fee, ttl)
	return tx, nil
}

// getTransactionByHashHeighter is used by WaitForTransactionForXBlocks to
// specify that the node/mock node passed in should support
// GetTransactionByHash() and GetHeight()
type getTransactionByHashHeighter interface {
	naet.GetTransactionByHasher
	naet.GetHeighter
}

// ErrWaitTransaction is returned by WaitForTransactionForXBlocks() to let
// callers distinguish between network errors and transaction acceptance errors.
type ErrWaitTransaction struct {
	NetworkErr     bool
	TransactionErr bool
	Err            error
}

func (b ErrWaitTransaction) Error() string {
	var errType string
	if b.TransactionErr {
		errType = "TransactionErr"
	} else {
		errType = "NetworkErr"
	}

	return fmt.Sprintf("%s: %s", errType, b.Err.Error())
}

// WaitForTransactionForXBlocks blocks until a transaction has been mined or X
// blocks have gone by, after which it returns an error. The node polling
// interval can be config.Configured with config.Tuning.ChainPollInterval.
func WaitForTransactionForXBlocks(c getTransactionByHashHeighter, txHash string, x uint64) (blockHeight uint64, blockHash string, wtError error) {
	nodeHeight, err := c.GetHeight()
	if err != nil {
		wtError = ErrWaitTransaction{
			NetworkErr:     true,
			TransactionErr: false,
			Err:            err,
		}
		return
	}
	endHeight := nodeHeight + x
	for nodeHeight <= endHeight {
		nodeHeight, err = c.GetHeight()
		if err != nil {
			wtError = ErrWaitTransaction{
				NetworkErr:     true,
				TransactionErr: false,
				Err:            err,
			}
			return
		}
		tx, err := c.GetTransactionByHash(txHash)
		if err != nil {
			wtError = ErrWaitTransaction{
				NetworkErr:     false,
				TransactionErr: true,
				Err:            err,
			}
			return
		}

		if tx.BlockHeight.LargerThanZero() {
			bh := big.Int(tx.BlockHeight)
			return bh.Uint64(), *tx.BlockHash, nil
		}
		time.Sleep(time.Millisecond * time.Duration(config.Tuning.ChainPollInterval))
	}
	wtError = ErrWaitTransaction{
		NetworkErr:     false,
		TransactionErr: true,
		Err:            fmt.Errorf("%v blocks have gone by and %v still isn't in a block", x, txHash),
	}
	return
}

// SignBroadcastTransaction signs a transaction and broadcasts it to a node.
func SignBroadcastTransaction(tx rlp.Encoder, signingAccount *account.Account, n naet.PostTransactioner, networkID string) (signedTxStr, hash, signature string, err error) {
	signedTx, hash, signature, err := transactions.SignHashTx(signingAccount, tx, networkID)
	if err != nil {
		return
	}

	signedTxStr, err = transactions.SerializeTx(signedTx)
	if err != nil {
		return
	}

	err = n.PostTransaction(signedTxStr, hash)
	if err != nil {
		return
	}
	return
}

type broadcastWaitTransactionNodeCapabilities interface {
	naet.PostTransactioner
	getTransactionByHashHeighter
}

// SignBroadcastWaitTransaction is a convenience function that combines
// SignBroadcastTransaction and WaitForTransactionForXBlocks.
func SignBroadcastWaitTransaction(tx rlp.Encoder, signingAccount *account.Account, n broadcastWaitTransactionNodeCapabilities, networkID string, x uint64) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	signedTxStr, hash, signature, err = SignBroadcastTransaction(tx, signingAccount, n, networkID)
	if err != nil {
		return
	}
	blockHeight, blockHash, err = WaitForTransactionForXBlocks(n, hash, x)
	return
}

// generateCommitmentID gives a commitment ID 'cm_...' given a particular AENS
// name. It is split into the deterministic part computeCommitmentID(), which
// can be tested, and the part incorporating random salt generateCommitmentID()
//
// since the salt is a uint256, which Erlang handles well, but Go has nothing
// similar to it, it is imperative that the salt be kept as a bytearray unless
// you really have to convert it into an integer. Which you usually don't,
// because it's a salt.
func generateCommitmentID(name string) (ch string, salt *big.Int, err error) {
	// Generate 32 random bytes for a salt
	saltBytes := make([]byte, 32)
	_, err = rand.Read(saltBytes)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return
	}

	ch, err = computeCommitmentID(name, saltBytes)

	salt = new(big.Int)
	salt.SetBytes(saltBytes)

	return ch, salt, err
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !strconv.IsPrint(r) {
			return false
		}
	}
	return true
}

func computeCommitmentID(name string, salt []byte) (ch string, err error) {
	var nh = []byte{}
	if strings.HasSuffix(name, ".test") {
		nh = append(Namehash(name), salt...)

	} else {
		// Since UTF-8 ~ ASCII, just use the string directly. QuoteToASCII
		// includes an extra byte at the start and end of the string, messing up
		// the hashing process.
		if !isPrintable(name) {
			return "", fmt.Errorf("The name %s must contain only printable characters", name)
		}

		nh = append([]byte(name), salt...)
	}
	nh, err = binary.Blake2bHash(nh)
	if err != nil {
		return
	}
	ch = binary.Encode(binary.PrefixCommitment, nh)
	return
}

// Namehash calculate the Namehash of a string. Names within aeternity are
// generally referred to only by their namehashes.
//
// The implementation is the same as ENS EIP-137
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-137.md#namehash-algorithm
// but using Blake2b.
func Namehash(name string) []byte {
	buf := make([]byte, 32)
	for _, s := range strings.Split(name, ".") {
		sh, _ := binary.Blake2bHash([]byte(s))
		buf, _ = binary.Blake2bHash(append(buf, sh...))
	}
	return buf
}
