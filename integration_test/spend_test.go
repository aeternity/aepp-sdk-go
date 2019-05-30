package integrationtest

import (
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

// Tests for 2 things: sending an amount that is max uint64, and that the node accepts the minimum fee
// that is calculated via tx.EstimateFee().
func TestSpendTx(t *testing.T) {
	node := setupNetwork(t)
	alice, bob := setupAccounts(t)

	amount := utils.RequireBigIntFromString("18446744073709551615") // max uint64
	fee := utils.NewBigIntFromUint64(uint64(2e13))
	msg := "Hello World"

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	expected := utils.NewBigInt()
	bobState, err := node.APIGetAccount(bob.Address)
	if err != nil {
		expected.Set(amount.Int)
	} else {
		expected.Add(bobState.Balance.Int, amount.Int)
	}

	ttl, nonce, err := node.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, bob.Address, *amount, *fee, msg, ttl, nonce)
	// minimize the fee to save money!
	est, _ := tx.FeeEstimate()
	t.Log("Estimated vs Actual Fee:", est, tx.Fee)
	tx.Fee = *est

	txB64, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	t.Log(txB64)

	// sign the transaction, output params for debugging
	_, err = signBroadcast(txB64, alice, node)
	if err != nil {
		t.Fatal(err)
	}
	// check the recipient's balance
	time.Sleep(200 * time.Millisecond)
	bobState, err = node.APIGetAccount(bob.Address)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if bobState.Balance.Cmp(expected.Int) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expected.String(), bobState.Balance.String())
	}
}
