package integrationtest

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v5/account"
	"github.com/aeternity/aepp-sdk-go/v5/config"
	"github.com/aeternity/aepp-sdk-go/v5/naet"
	"github.com/aeternity/aepp-sdk-go/v5/aeternity"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
var recipientPrivateKey = os.Getenv("INTEGRATION_TEST_RECEIVER_PRIVATE_KEY")
var privatenetURL = "http://localhost:3013"
var testnetURL = "http://sdk-testnet.aepps.com"
var networkID = "ae_docker"

func setupNetwork(t *testing.T, nodeURL string, debug bool) *naet.Node {
	config.Node.NetworkID = networkID
	client := naet.NewNode(nodeURL, debug)
	t.Logf("nodeURL: %s, networkID: %s", nodeURL, config.Node.NetworkID)
	return client
}

func setupAccounts(t *testing.T) (*account.Account, *account.Account) {
	alice, err := account.FromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	bob, err := account.FromHexString(recipientPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alice: %s, Bob: %s", alice.Address, bob.Address)
	return alice, bob
}

func readFile(t *testing.T, filename string) (r string) {
	rb, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(rb)
}

type delayableCode func()

func delay(f delayableCode) {
	time.Sleep(2000 * time.Millisecond)
	f()
}

func getHeight(node *naet.Node) (h uint64) {
	h, err := node.GetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	// fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(node *naet.Node, hash string) (height uint64, microblockHash string, err error) {
	height, microblockHash, err = aeternity.WaitForTransactionForXBlocks(node, hash, 10)
	if err != nil {
		// Sometimes, the tests want the tx to fail. Return the err to let them know.
		return 0, "", err
	}
	fmt.Println("Transaction was found at", height, "microblockHash", microblockHash, "err", err)
	return height, microblockHash, err
}

func fundAccount(t *testing.T, node *naet.Node, source, destination *account.Account, amount *big.Int) {
	ctx := aeternity.NewContextFromNode(node, source.Address)

	fmt.Println("Funding account", destination.Address)
	tx, err := ctx.SpendTx(source.Address, destination.Address, amount, config.Client.Fee, []byte{})
	if err != nil {
		t.Fatal(err)
	}
	_, hash, _, err := aeternity.SignBroadcastTransaction(tx, source, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = waitForTransaction(node, hash)
	if err != nil {
		t.Fatal(err)
	}
}
