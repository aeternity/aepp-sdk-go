package aeternity

import (
	"fmt"
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/naet"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
	rlp "github.com/randomshinichi/rlpae"
)

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
