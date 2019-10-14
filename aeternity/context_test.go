package aeternity

import (
	"fmt"
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
)

func ExampleContext() {
	// Set the Network ID. For this example, setting the config.Node.NetworkID
	// is actually not needed - but if you have other code that also needs to
	// access NetworkID somehow, do it this way.
	config.Node.NetworkID = config.NetworkIDTestnet

	alice, err := account.FromHexString("deadbeef")
	if err != nil {
		fmt.Println("Could not create alice's Account:", err)
	}

	bobAddress := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

	// create a Context for the node you will use and address you're going to
	// sign the transaction with
	ctx, node := NewContextFromURL("http://localhost:3013", alice.Address, false)

	// create the SpendTransaction
	amount := big.NewInt(1e9)
	fee := big.NewInt(1e6)
	msg := "Reason For Payment"
	tx, err := ctx.SpendTx(alice.Address, bobAddress, amount, fee, []byte(msg))
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	// Optional: minimize the fee to save money!
	err = transactions.CalculateFee(tx)
	if err != nil {
		fmt.Println("Could not calculate the transaction fee", err)
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
