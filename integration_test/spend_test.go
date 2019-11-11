package integrationtest

import (
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/aeternity/aepp-sdk-go/v7/utils"
)

// Tests for 2 things: sending an amount that is max uint64, and that the node
// accepts the minimum fee that is calculated via tx.EstimateFee().
func TestSpendTx(t *testing.T) {
	node := setupNetwork(t, privatenetURL, false)
	alice, bob := setupAccounts(t)

	amount := utils.RequireIntFromString("18446744073709551615") // max uint64
	msg := "Hello World"

	// In case the recipient account already has funds, get recipient's account
	// info. If it exists, expectedAmount = existing balance + amount + fee
	expected := new(big.Int)
	bobState, err := node.GetAccount(bob.Address)
	if err != nil {
		expected.Set(amount)
	} else {
		bS := big.Int(bobState.Balance)
		expected.Add(&bS, amount)
	}

	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	tx, err := transactions.NewSpendTx(alice.Address, bob.Address, amount, []byte(msg), ttlnoncer)
	if err != nil {
		t.Error(err)
	}
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(tx, alice, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	// check the recipient's balance

	getBobsAccount := func() {
		bobState, err = node.GetAccount(bob.Address)
		if err != nil {
			t.Fatalf("Couldn't get Bob's account data: %v", err)
		}
	}
	delay(getBobsAccount)
	b := big.Int(bobState.Balance)

	if expected.Cmp(&b) != 0 {
		t.Fatalf("Bob should have %v, but has %v instead", expected.String(), bobState.Balance.String())
	}
}
