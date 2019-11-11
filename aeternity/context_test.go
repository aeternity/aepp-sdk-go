package aeternity

import (
	"fmt"
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

func Example() {
	// Set the Network ID. For this example, setting the config.Node.NetworkID
	// is actually not needed - but if you have other code that also needs to
	// access NetworkID somehow, do it this way.
	config.Node.NetworkID = config.NetworkIDTestnet

	alice, err := account.FromHexString("deadbeef")
	if err != nil {
		fmt.Println("Could not create alice's Account:", err)
	}

	bobAddress := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

	// create a connection to a node, represented by *Node
	node := naet.NewNode("http://localhost:3013", false)

	// create the closures that autofill the correct account nonce and transaction TTL
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	// create the SpendTransaction
	msg := "Reason For Payment"
	tx, err := transactions.NewSpendTx(alice.Address, bobAddress, big.NewInt(1e9), []byte(msg), ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	_, _, _, _, _, err = SignBroadcastWaitTransaction(tx, alice, node, config.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
	}

	// check the recipient's balance
	time.Sleep(2 * time.Second)
	bobState, err := node.GetAccount(bobAddress)
	if err != nil {
		fmt.Println("Couldn't get Bob's account data:", err)
	}

	fmt.Println(bobState.Balance)
}
