package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
	"gotest.tools/golden"
)

func TestContracts(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	contractsAlice := aeternity.NewContextFromNode(node, alice.Address)

	var ctID string
	var txHash string

	identityBytecode := string(golden.Get(t, "identity_bytecode.txt"))
	identityInitCalldata := string(golden.Get(t, "identity_initcalldata.txt"))
	create, err := contractsAlice.ContractCreateTx(identityBytecode, identityInitCalldata, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Deposit, aeternity.Config.Client.Contracts.Amount, *utils.NewIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewIntFromUint64(564480000000000))
	if err != nil {
		t.Fatal(err)
	}
	ctID, _ = create.ContractID()
	fmt.Printf("Create %s, %+v\n", ctID, create)
	_, txHash, _, err = aeternity.SignBroadcastTransaction(create, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = waitForTransaction(node, txHash)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm that contract was created
	getContract := func() {
		_, err = node.GetContractByID(ctID)
		if err != nil {
			t.Fatal(err)
		}
	}
	delay(getContract)

	identityMain42Calldata := string(golden.Get(t, "identity_main42.txt"))
	callTx, err := contractsAlice.ContractCallTx(ctID, identityMain42Calldata, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Amount, *utils.NewIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewIntFromUint64(665480000000000))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Call %+v\n", callTx)
	_, txHash, _, err = aeternity.SignBroadcastTransaction(callTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = waitForTransaction(node, txHash)
	if err != nil {
		t.Fatal(err)
	}
}
