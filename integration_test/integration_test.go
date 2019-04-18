package integration_test

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")

func TestSpendTxWithNode(t *testing.T) {
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
		expectedAmount.Add(recipientAccount.Balance.Int, big.NewInt(10))
		fmt.Printf("Recipient already exists with balance %v, expectedAmount after test is %s\n", recipientAccount.Balance.String(), expectedAmount.String())
	}

	amount := utils.NewBigInt()
	amount.SetInt64(10)
	fee := utils.NewBigInt()
	fee.SetUint64(uint64(2e13))
	ttl, nonce, err := aeCli.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, recipient, *amount, *fee, message, ttl, nonce)
	base64TxMsg, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, hash, signature, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg, hash, signature)

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

	if recipientAccount.Balance.Cmp(&expectedAmount) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}

func TestSpendTxLargeWithNode(t *testing.T) {
	// This is a separate test because the account may not have enough funds for this test when the node has just started.
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

func signBroadcast(tx string, acc *aeternity.Account, aeClient *aeternity.Ae) (hash string, err error) {
	signedTxStr, hash, _, err := aeternity.SignEncodeTxStr(acc, tx, "ae_docker")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = aeClient.BroadcastTransaction(signedTxStr)
	if err != nil {
		panic(err)
	}

	return hash, nil

}

func getHeight(aeClient *aeternity.Ae) (h uint64) {
	h, err := aeClient.APIGetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(aeClient *aeternity.Ae, height uint64, hash string) {
	height, blockHash, microBlockHash, _, err := aeClient.WaitForTransactionUntilHeight(height+10, hash)
	fmt.Println("Transaction was found at", height, "blockhash", blockHash, "microBlockHash", microBlockHash, "err", err)
}

func TestAENSWorkflow(t *testing.T) {
	name := "fdsa.test"
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeClient := aeternity.NewCli("http://localhost:3013", false).WithAccount(acc)
	aeternity.Config.Client.Fee = *utils.RequireBigIntFromString("100000000000000")

	// Preclaim the name
	fmt.Println("PreclaimTx")
	preclaimTx, salt, err := aeClient.Aens.NamePreclaimTx(name, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	preclaimTxStr, _ := aeternity.BaseEncodeTx(&preclaimTx)
	fmt.Println("PreclaimTx and Salt:", preclaimTxStr, salt)

	hash, err := signBroadcast(preclaimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NamePreclaimTx", hash)
	height := getHeight(aeClient)

	// Wait for a bit
	waitForTransaction(aeClient, height, hash)

	// Claim the name
	fmt.Println("NameClaimTx")
	claimTx, err := aeClient.Aens.NameClaimTx(name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	claimTxStr, _ := aeternity.BaseEncodeTx(&claimTx)
	fmt.Println("ClaimTx:", claimTxStr)

	hash, err = signBroadcast(claimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameClaimTx")
	height = getHeight(aeClient)

	// Wait for a bit
	waitForTransaction(aeClient, height, hash)

	// Verify that the name exists
	nameEntry, err := aeClient.APIGetNameEntryByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	nameEntryJSON, _ := nameEntry.MarshalBinary()
	fmt.Println(string(nameEntryJSON))

	// Update the name, make it point to something
	fmt.Println("NameUpdateTx")
	updateTx, err := aeClient.Aens.NameUpdateTx(name, acc.Address)
	updateTxStr, _ := aeternity.BaseEncodeTx(&updateTx)
	fmt.Println("UpdateTx:", updateTxStr)

	_, err = signBroadcast(updateTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameUpdateTx")

}
