package aeternity

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/aeternity/aepp-sdk-go/v7/utils"
)

type mockClient struct {
	i uint64
}

func (m *mockClient) GetHeight() (uint64, error) {
	m.i++
	return m.i, nil
}

// GetTransactionByHash pretends that the transaction was not mined until block 9, and this is only visible when the mockClient is at height 10.
func (m *mockClient) GetTransactionByHash(hash string) (tx *models.GenericSignedTx, err error) {
	unminedHeight, _ := utils.NewIntFromString("-1")
	minedHeight, _ := utils.NewIntFromString("9")

	bh := "bh_someblockhash"
	tx = &models.GenericSignedTx{
		BlockHash:   &bh,
		BlockHeight: utils.BigInt{},
		Hash:        &hash,
		Signatures:  nil,
	}

	if m.i == 10 {
		tx.BlockHeight.Set(minedHeight)
	} else {
		tx.BlockHeight.Set(unminedHeight)
	}
	return tx, nil
}
func TestWaitForTransactionForXBlocks(t *testing.T) {
	m := new(mockClient)
	blockHeight, blockHash, err := WaitForTransactionForXBlocks(m, "th_transactionhash", 10)
	if err != nil {
		t.Fatal(err)
	}
	if blockHeight != 9 {
		t.Fatalf("Expected mock blockHeight 9, got %v", blockHeight)
	}
	if blockHash != "bh_someblockhash" {
		t.Fatalf("Expected mock blockHash bh_someblockhash, got %s", blockHash)
	}
}

func Example() {
	// Set the Network ID. For this example, setting the config.Node.NetworkID
	// is actually not needed - but if you have other code that also needs to
	// access NetworkID somehow, do it this way.
	config.Node.NetworkID = config.NetworkIDTestnet

	alice, err := account.FromHexString("deadbeef")
	if err != nil {
		fmt.Println("Could not create alice's Account:", err)
	}

	bobAddress := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

	// create a connection to a node, represented by *Node
	node := naet.NewNode("http://localhost:3013", false)

	// create the closures that autofill the correct account nonce and transaction TTL
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	// create the SpendTransaction
	msg := "Reason For Payment"
	tx, err := transactions.NewSpendTx(alice.Address, bobAddress, big.NewInt(1e9), []byte(msg), ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	signedTxStr, hash, signature, foundAtBlockHeight, foundAtBlockHash, err := SignBroadcastWaitTransaction(tx, alice, node, config.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
	}
	fmt.Println(signedTxStr, hash, signature, foundAtBlockHeight, foundAtBlockHash)

	// check the recipient's balance
	time.Sleep(2 * time.Second)
	bobState, err := node.GetAccount(bobAddress)
	if err != nil {
		fmt.Println("Couldn't get Bob's account data:", err)
	}

	fmt.Println(bobState.Balance)
}
