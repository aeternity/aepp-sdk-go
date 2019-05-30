package integrationtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
var recipientPrivateKey = os.Getenv("INTEGRATION_TEST_RECEIVER_PRIVATE_KEY")
var nodeURL = "http://localhost:3013"
var networkID = "ae_docker"

func setupNetwork(t *testing.T) *aeternity.Client {
	t.Log("setup integration test")
	aeternity.Config.Node.NetworkID = networkID
	client := aeternity.NewClient(nodeURL, false)
	return client
}

func setupAccounts(t *testing.T) (*aeternity.Account, *aeternity.Account) {
	alice, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	bob, err := aeternity.AccountFromHexString(recipientPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alice: %s, Bob: %s", alice.Address, bob.Address)
	return alice, bob
}

func signBroadcast(tx string, acc *aeternity.Account, aeClient *aeternity.Client) (hash string, err error) {
	signedTxStr, hash, _, err := aeternity.SignEncodeTxStr(acc, tx, aeternity.Config.Node.NetworkID)
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

func getHeight(aeClient *aeternity.Client) (h uint64) {
	h, err := aeClient.APIGetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(aeClient *aeternity.Client, hash string) (err error) {
	height := getHeight(aeClient)
	fmt.Println("Waiting for Transaction...")
	height, blockHash, microBlockHash, _, err := aeClient.WaitForTransactionUntilHeight(height+10, hash)
	if err != nil {
		// Sometimes, the tests want the tx to fail. Return the err to let them know.
		return err
	}
	fmt.Println("Transaction was found at", height, "blockhash", blockHash, "microBlockHash", microBlockHash, "err", err)
	return nil
}
