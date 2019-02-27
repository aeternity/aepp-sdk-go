package aeternity_test

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestSpendTransaction(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"
	eBase64TxMsg := "tx_+FgMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIAAHi0hlbGxvIFdvcmxk0pjHOg=="

	base64TxMsg, _, _, err := aeternity.SpendTransaction(sender, recipient, 10, 10, message)
	if err != nil {
		t.Errorf("SpendTransaction errored out: %v", err)
	}
	if base64TxMsg != eBase64TxMsg {
		t.Errorf("SpendTransaction returned a wrong tx_ blob, got %s, want %s", base64TxMsg, eBase64TxMsg)
	}

	eBase64Tx := "tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIAAHgDJ1jfU="
	base64Tx, _, _, err := aeternity.SpendTransaction(sender, recipient, 10, 10, "")
	if err != nil {
		t.Errorf("SpendTransaction errored out: %v", err)
	}
	if base64Tx != eBase64Tx {
		t.Errorf("SpendTransaction without a message returned a wrong tx_ blob, got %s, want %s", base64TxMsg, eBase64TxMsg)
	}

}

func TestSpendTransactionWithNode(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	senderAccount, _ := aeternity.AccountFromHexString(senderPrivateKey)
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

	aeternity.Config.Epoch.URL = "http://localhost:3013"
	aeternity.Config.Epoch.NetworkID = "ae_docker"

	// create the SpendTransaction
	base64TxMsg, _, _, err := aeternity.SpendTransaction(sender, recipient, 10, 10, "")
	if err != nil {
		t.Errorf("SpendTransaction errored out: %v", err)
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
		t.Errorf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	aeCli := aeternity.NewCli(aeternity.Config.Epoch.URL, false)
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Errorf("Couldn't get refcipient's account data: %v", err)
	}

	ten := big.NewInt(10)
	if recipientAccount.Balance.Cmp(ten) != 0 {
		t.Errorf("Recipient should have 10AE, but has %v instead", recipientAccount.Balance.Int)
	}
}
