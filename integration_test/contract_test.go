package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/golden"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestContracts(t *testing.T) {
	alice, _ := setupAccounts(t)
	aeNode := setupNetwork(t, privatenetURL, false)
	helpers := aeternity.Helpers{Node: aeNode}
	contractsAlice := aeternity.Context{Helpers: helpers, Address: alice.Address}

	var ctID string
	var txHash string

	create, err := contractsAlice.ContractCreateTx(golden.IdentityBytecode, golden.IdentityInitCalldata, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Deposit, aeternity.Config.Client.Contracts.Amount, *utils.NewIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewIntFromUint64(564480000000000))
	if err != nil {
		t.Fatal(err)
	}
	ctID, _ = create.ContractID()
	fmt.Printf("Create %s, %+v\n", ctID, create)
	_, txHash, _, err = aeternity.SignBroadcastTransaction(&create, alice, aeNode, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = waitForTransaction(aeNode, txHash)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm that contract was created
	getContract := func() {
		_, err = aeNode.GetContractByID(ctID)
		if err != nil {
			t.Fatal(err)
		}
	}
	delay(getContract)

	callTx, err := contractsAlice.ContractCallTx(ctID, golden.IdentityCalldata, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Amount, *utils.NewIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewIntFromUint64(665480000000000))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Call %+v\n", callTx)
	_, txHash, _, err = aeternity.SignBroadcastTransaction(&callTx, alice, aeNode, networkID)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = waitForTransaction(aeNode, txHash)
	if err != nil {
		t.Fatal(err)
	}
}
