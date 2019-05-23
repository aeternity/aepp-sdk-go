package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

// Tests for 2 things: sending an amount that is max uint64, and that the node accepts the minimum fee
// that is calculated via tx.EstimateFee().
func TestSpendTx(t *testing.T) {
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	aeternity.Config.Node.URL = nodeURL
	aeternity.Config.Node.NetworkID = networkID
	aeCli := aeternity.NewClient(aeternity.Config.Node.URL, false)

	amount := utils.RequireBigIntFromString("18446744073709551615") // max uint64
	fee := utils.NewBigIntFromUint64(uint64(2e13))
	var expectedAmount = utils.NewBigInt()

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.Set(amount.Int)
	} else {
		expectedAmount.Add(recipientAccount.Balance.Int, amount.Int)
	}

	ttl, nonce, err := aeCli.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, recipient, *amount, *fee, message, ttl, nonce)
	// minimize the fee to save money!
	estimatedFee, _ := tx.FeeEstimate()
	fmt.Println("Estimated vs Actual Fee:", estimatedFee, tx.Fee)
	tx.Fee = *estimatedFee

	base64TxMsg, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeCli.BroadcastTransaction(signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(expectedAmount.Int) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}
