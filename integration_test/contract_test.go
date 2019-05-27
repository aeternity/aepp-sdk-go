package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestContracts(t *testing.T) {
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeternity.Config.Node.NetworkID = networkID
	aeClient := aeternity.NewClient(nodeURL, false)
	contractsAlice := aeternity.Contract{Client: aeClient, Account: acc}

	var ctID string
	var callData string
	var txHash string

	code := "cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib"
	callData = "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA=="
	fmt.Println("Contract Create")
	tx, err := contractsAlice.ContractCreateTx(code, callData, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Deposit, aeternity.Config.Client.Contracts.Amount, *utils.NewBigIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewBigIntFromUint64(564480000000000))
	if err != nil {
		t.Fatal(err)
	}

	ctID, _ = tx.ContractID()
	txStr, _ := aeternity.BaseEncodeTx(&tx)
	fmt.Printf("%#v\n", tx)
	fmt.Println(ctID)
	txHash, err = signBroadcast(txStr, acc, aeClient)
	fmt.Println(txHash)
	err = waitForTransaction(aeClient, txHash)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Confirm that contract was created")
	time.Sleep(1000 * time.Millisecond)
	_, err = aeClient.APIGetContractByID(ctID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Contract Call")
	callData = "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBo8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7dbVl"
	callTx, err := contractsAlice.ContractCallTx(ctID, callData, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Amount, *utils.NewBigIntFromUint64(1e5), aeternity.Config.Client.Contracts.GasPrice, *utils.NewBigIntFromUint64(665480000000000))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", callTx)
	fmt.Println(callTx.FeeEstimate())
	callTxStr, _ := aeternity.BaseEncodeTx(&callTx)
	fmt.Println(callTxStr)
	txHash, err = signBroadcast(callTxStr, acc, aeClient)
	fmt.Println(txHash)

	err = waitForTransaction(aeClient, txHash)
	if err != nil {
		t.Fatal(err)
	}

}
