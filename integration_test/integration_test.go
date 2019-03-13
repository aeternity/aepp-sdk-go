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

	aeternity.Config.Node.URL = "http://localhost:3013"
	aeternity.Config.Node.NetworkID = "ae_docker"
	aeCli := aeternity.NewCli(aeternity.Config.Node.URL, false)

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
	ttl, nonce, err := aeternity.GetTTLNonce(aeCli.Node, sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	base64TxMsg, err := aeternity.SpendTxStr(sender, recipient, *amount, *fee, message, ttl, nonce)
	if err != nil {
		t.Fatalf("SpendTx errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeternity.BroadcastTransaction(aeCli.Node, signedBase64TxMsg)
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

func TestSpendTxLargeWithNode(t *testing.T) {
	// This is a separate test because the account may not have enough funds for this test when the node has just started.
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	aeternity.Config.Node.URL = "http://localhost:3013"
	aeternity.Config.Node.NetworkID = "ae_docker"
	aeCli := aeternity.NewCli(aeternity.Config.Node.URL, false)

	amount := utils.RequireBigIntFromString("18446744073709551615") // max uint64
	fee := utils.NewBigIntFromUint64(uint64(2e13))
	var expectedAmount = utils.NewBigInt()

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.Set(&amount.Int)
	} else {
		expectedAmount.Add(&recipientAccount.Balance.Int, &amount.Int)
	}

	ttl, nonce, err := aeternity.GetTTLNonce(aeCli.Node, sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	base64TxMsg, err := aeternity.SpendTxStr(sender, recipient, *amount, *fee, message, ttl, nonce)
	if err != nil {
		t.Fatalf("SpendTx errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeternity.BroadcastTransaction(aeCli.Node, signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(&expectedAmount.Int) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}
