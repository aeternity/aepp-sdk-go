package integrationtest

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var alicePrivateKey = os.Getenv("AETERNITY_ALICE_PRIVATE_KEY")
var bobPrivateKey = os.Getenv("AETERNITY_BOB_PRIVATE_KEY")
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
	alice, err := account.FromHexString(alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	bob, err := account.FromHexString(bobPrivateKey)
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
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	fmt.Println("Funding account", destination.Address)
	tx, err := transactions.NewSpendTx(source.Address, destination.Address, amount, []byte{}, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(tx, source, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
}
