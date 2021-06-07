package integrationtest

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v8/account"
	"github.com/aeternity/aepp-sdk-go/v8/aeternity"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var alicePrivateKey = "e6a91d633c77cf5771329d3354b3bcef1bc5e032c43d70b6d35af923ce1eb74dcea7ade470c9f99d9d4e400880a86f1d49bb444b62f11a9ebb64bbcfeb73fef3"
var bobPrivateKey = "7065616e38c1da983bc619188efe19bbddc8c149ddfcd3ed1c294294957a18477b47ed425587f4abd5064fe61d5a0121949a4125e8b700a2d14f0bbbafb8b2c6"
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

func fundAccount(t *testing.T, node *naet.Node, source, destination *account.Account, amount *big.Int) {
	ttlnoncer := transactions.NewTTLNoncer(node)
	fmt.Println("Funding account", destination.Address)
	tx, err := transactions.NewSpendTx(source.Address, destination.Address, amount, []byte{}, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	receipt, err := aeternity.SignBroadcast(tx, source, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	err = aeternity.WaitSynchronous(receipt, config.Client.WaitBlocks, node)
	if err != nil {
		t.Fatal(err)
	}
}
