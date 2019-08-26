package integrationtest

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
var recipientPrivateKey = os.Getenv("INTEGRATION_TEST_RECEIVER_PRIVATE_KEY")
var privatenetURL = "http://localhost:3013"
var testnetURL = "http://sdk-testnet.aepps.com"
var networkID = "ae_docker"

func setupNetwork(t *testing.T, nodeURL string, debug bool) *aeternity.Node {
	aeternity.Config.Node.NetworkID = networkID
	client := aeternity.NewNode(nodeURL, debug)
	t.Logf("nodeURL: %s, networkID: %s", nodeURL, aeternity.Config.Node.NetworkID)
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

func getHeight(aeNode *aeternity.Node) (h uint64) {
	h, err := aeNode.GetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	// fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(aeNode *aeternity.Node, hash string) (height uint64, microblockHash string, err error) {
	height, microblockHash, err = aeternity.WaitForTransactionForXBlocks(aeNode, hash, 10)
	if err != nil {
		// Sometimes, the tests want the tx to fail. Return the err to let them know.
		return 0, "", err
	}
	fmt.Println("Transaction was found at", height, "microblockHash", microblockHash, "err", err)
	return height, microblockHash, err
}

func fundAccount(t *testing.T, n *aeternity.Node, source, destination *aeternity.Account, amount *big.Int) {
	h := aeternity.Helpers{Node: n}
	ctx := aeternity.Context{Address: source.Address, Helpers: h}

	fmt.Println("Funding account", destination.Address)
	tx, err := ctx.SpendTx(source.Address, destination.Address, *amount, aeternity.Config.Client.Fee, []byte{})
	if err != nil {
		t.Fatal(err)
	}
	_, hash, _, err := aeternity.SignBroadcastTransaction(&tx, source, n, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = waitForTransaction(n, hash)
	if err != nil {
		t.Fatal(err)
	}
}
