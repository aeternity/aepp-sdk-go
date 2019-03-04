package integration_test

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestSpendTxWithNode(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	aeternity.Config.Epoch.URL = "http://localhost:3013"
	aeternity.Config.Epoch.NetworkID = "ae_docker"
	aeCli := aeternity.NewCli(aeternity.Config.Epoch.URL, false)

	// In case this test has been run before, get recipient's account info. If it exists, expectedAmount = amount + 10
	var expectedAmount big.Int
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.SetInt64(10)
	} else {
		expectedAmount.Add(&recipientAccount.Balance.Int, big.NewInt(10))
		fmt.Printf("Recipient already exists with balance %v, expectedAmount after test is %s\n", recipientAccount.Balance.String(), expectedAmount.String())
	}

	amount := utils.NewBigInt()
	amount.SetInt64(10)
	fee := utils.NewBigInt()
	fee.SetUint64(uint64(2e13))
	ttl, nonce, err := aeternity.GetTTLNonce(aeCli.Epoch, sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	base64TxMsg, err := aeternity.SpendTxStr(sender, recipient, *amount, *fee, ttl, nonce, message)
	if err != nil {
		t.Fatalf("SpendTx errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, txHash, signature, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)
	fmt.Println(txHash)
	fmt.Println(signature)

	// send the signed transaction to the node
	err = aeternity.BroadcastTransaction(signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(&expectedAmount) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}
