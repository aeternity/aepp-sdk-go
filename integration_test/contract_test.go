package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"gotest.tools/golden"
)

func TestContracts(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	ttlnoncer := transactions.NewTTLNoncer(node)

	var ctID string

	identityBytecode := string(golden.Get(t, "identity_bytecode.txt"))
	identityInitCalldata := string(golden.Get(t, "identity_initcalldata.txt"))
	create, err := transactions.NewContractCreateTx(alice.Address, identityBytecode, config.Client.Contracts.VMVersion, config.Client.Contracts.ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, identityInitCalldata, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	ctID, _ = create.ContractID()
	_, txHash, _, _, _, err := aeternity.SignBroadcastWaitTransaction(create, alice, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Create %s, %+v %s\n", ctID, create, txHash)

	// Confirm that contract was created
	getContract := func() {
		_, err = node.GetContractByID(ctID)
		if err != nil {
			t.Fatal(err)
		}
	}
	delay(getContract)

	identityMain42Calldata := string(golden.Get(t, "identity_main42.txt"))
	callTx, err := transactions.NewContractCallTx(alice.Address, ctID, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, identityMain42Calldata, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, txHash, _, _, _, err = aeternity.SignBroadcastWaitTransaction(callTx, alice, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Call %+v %s\n", callTx, txHash)
}
