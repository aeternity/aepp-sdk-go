package integrationtest

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

// Tests for 2 things: sending an amount that is max uint64, and that the node accepts the minimum fee
// that is calculated via tx.EstimateFee().
func TestSpendTx(t *testing.T) {
	node := setupNetwork(t, privatenetURL, false)
	alice, bob := setupAccounts(t)

	amount := utils.RequireIntFromString("18446744073709551615") // max uint64
	fee := utils.NewIntFromUint64(uint64(2e13))
	msg := "Hello World"

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	expected := new(big.Int)
	bobState, err := node.GetAccount(bob.Address)
	if err != nil {
		expected.Set(amount)
	} else {
		bS := big.Int(bobState.Balance)
		expected.Add(&bS, amount)
	}

	// create a Context for the address you're going to sign the transaction
	// with, and an aeternity node to talk to/query the address's nonce.
	ctx := aeternity.NewContext(node, alice.Address)

	// create the SpendTransaction
	tx, err := ctx.SpendTx(alice.Address, bob.Address, *amount, *fee, []byte(msg))
	if err != nil {
		t.Error(err)
	}

	// minimize the fee to save money!
	est, _ := tx.FeeEstimate()
	fmt.Println("Estimated vs Actual Fee:", est, tx.Fee)
	tx.Fee = *est

	// sign the transaction, output params for debugging
	hash := signBroadcast(t, &tx, alice, node)
	// check the recipient's balance

	// Wait for a bit
	_, _, _ = waitForTransaction(node, hash)

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
