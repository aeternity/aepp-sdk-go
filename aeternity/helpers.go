package aeternity

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
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

// SignBroadcast signs a transaction and broadcasts it to a node.
func SignBroadcast(tx transactions.Transaction, signingAccount *account.Account, n naet.PostTransactioner, networkID string) (txReceipt *TxReceipt, err error) {
	signedTx, hash, signature, err := transactions.SignHashTx(signingAccount, tx, networkID)
	if err != nil {
		return
	}

	signedTxStr, err := transactions.SerializeTx(signedTx)
	if err != nil {
		return
	}

	err = n.PostTransaction(signedTxStr, hash)
	if err != nil {
		return
	}

	txReceipt = NewTxReceipt(&tx, signedTxStr, hash, signature)
	return
}

// WaitSynchronous blocks until TxReceipt.Watch() reports that a transaction was
// mined/not mined. It is intended as a convenience function since it makes an
// asynchronous operation synchronous.
func WaitSynchronous(txReceipt *TxReceipt, waitBlocks uint64, n getTransactionByHashHeighter) (err error) {
	minedChan := make(chan bool)
	go txReceipt.Watch(minedChan, waitBlocks, n)
	mined := <-minedChan
	if !mined {
		return txReceipt.Error
	}
	return nil
}

// TxReceipt represents the status of a sent transaction
type TxReceipt struct {
	Tx          *transactions.Transaction
	SignedTx    string
	Hash        string
	Signature   string
	BlockHeight uint64
	BlockHash   string
	Mined       bool
	Error       error
}

func (t *TxReceipt) String() string {
	return fmt.Sprintf("Mined: %v\nTx: %+v\nSigned: %s\nHash: %s\nSignature: %s\nBlockHeight: %d\nBlockHash: %s", t.Mined, *t.Tx, t.SignedTx, t.Hash, t.Signature, t.BlockHeight, t.BlockHash)
}

func NewTxReceipt(tx *transactions.Transaction, signedTx, hash, signature string) (txReceipt *TxReceipt) {
	txReceipt = &TxReceipt{
		Tx:        tx,
		SignedTx:  signedTx,
		Hash:      hash,
		Signature: signature,
	}

	return
}

// Watch polls until a transaction has been mined or X blocks have gone by,
// after which it errors out via TxReceiptWatchResult. The node polling interval
// can be configured with config.Tuning.ChainPollInterval, which accepts a
// time.Duration.
func (t *TxReceipt) Watch(mined chan bool, waitBlocks uint64, node getTransactionByHashHeighter) {
	nodeHeight, err := node.GetHeight()
	if err != nil {
		t.Error = err
		mined <- false
	}
	endHeight := nodeHeight + waitBlocks
	for nodeHeight <= endHeight {
		nodeHeight, err = node.GetHeight()
		if err != nil {
			t.Error = err
			mined <- false
		}
		tx, err := node.GetTransactionByHash(t.Hash)
		if err != nil {
			t.Error = err
			mined <- false
		}

		if tx.BlockHeight.LargerThanZero() {
			bh := big.Int(tx.BlockHeight)
			t.BlockHeight = bh.Uint64()
			t.BlockHash = *tx.BlockHash
			t.Mined = true
			mined <- true
		}
		time.Sleep(config.Tuning.ChainPollInterval)
	}

	t.Error = fmt.Errorf("%v blocks have gone by and %v still isn't in a block", waitBlocks, t.Hash)
	mined <- false
}

func findVMABIVersion(nodeVersion, compilerBackend string) (VMVersion, ABIVersion uint16, err error) {
	if nodeVersion[0] == '5' && compilerBackend == "fate" {
		return 5, 3, nil
	} else if nodeVersion[0] == '5' && compilerBackend == "aevm" {
		return 6, 1, nil
	} else if nodeVersion[0] == '4' {
		return 4, 1, nil
	} else {
		return 0, 0, errors.New("Other node versions unsupported")
	}
}
