package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/aeternity"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/utils"
	"gotest.tools/golden"
)

func TestContracts(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	contractsAlice := aeternity.NewContextFromNode(node, alice.Address)

	var ctID string

	identityBytecode := string(golden.Get(t, "identity_bytecode.txt"))
	identityInitCalldata := string(golden.Get(t, "identity_initcalldata.txt"))
	create, err := contractsAlice.ContractCreateTx(identityBytecode, identityInitCalldata, config.Client.Contracts.VMVersion, config.Client.Contracts.ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, utils.NewIntFromUint64(1e5), utils.NewIntFromUint64(564480000000000))
	if err != nil {
		t.Fatal(err)
	}
	ctID, _ = create.ContractID()
	fmt.Printf("Create %s, %+v\n", ctID, create)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(create, alice, node, networkID, config.Client.WaitBlocks)
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
	callTx, err := contractsAlice.ContractCallTx(ctID, identityMain42Calldata, config.Client.Contracts.ABIVersion, config.Client.Contracts.Amount, utils.NewIntFromUint64(1e5), config.Client.GasPrice, utils.NewIntFromUint64(665480000000000))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Call %+v\n", callTx)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(callTx, alice, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
}
