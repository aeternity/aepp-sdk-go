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
	node := setupNetwork(t, privatenetURL)
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

	ttl, nonce, err := aeternity.GetTTLNonce(node, sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, bob.Address, *amount, *fee, msg, ttl, nonce)
	// minimize the fee to save money!
	est, _ := tx.FeeEstimate()
	fmt.Println("Estimated vs Actual Fee:", est, tx.Fee)
	tx.Fee = *est

	// sign the transaction, output params for debugging
	_ = signBroadcast(t, &tx, alice, node)
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
