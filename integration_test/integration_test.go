package integration_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
var recipientPrivateKey = os.Getenv("INTEGRATION_TEST_RECEIVER_PRIVATE_KEY")
var nodeURL = "http://localhost:3013"
var networkID = "ae_docker"

// Tests for 2 things: sending an amount that is max uint64, and that the node accepts the minimum fee
// that is calculated via tx.EstimateFee().
func TestSpendTxWithNode(t *testing.T) {
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	aeternity.Config.Node.URL = nodeURL
	aeternity.Config.Node.NetworkID = networkID
	aeCli := aeternity.NewCli(aeternity.Config.Node.URL, false)

	amount := utils.RequireBigIntFromString("18446744073709551615") // max uint64
	fee := utils.NewBigIntFromUint64(uint64(2e13))
	var expectedAmount = utils.NewBigInt()

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.Set(amount.Int)
	} else {
		expectedAmount.Add(recipientAccount.Balance.Int, amount.Int)
	}

	ttl, nonce, err := aeCli.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, recipient, *amount, *fee, message, ttl, nonce)
	// minimize the fee to save money!
	estimatedFee, _ := tx.FeeEstimate()
	fmt.Println("Estimated vs Actual Fee:", estimatedFee, tx.Fee)
	tx.Fee = *estimatedFee

	base64TxMsg, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeCli.BroadcastTransaction(signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(expectedAmount.Int) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}

func signBroadcast(tx string, acc *aeternity.Account, aeClient *aeternity.Ae) (hash string, err error) {
	signedTxStr, hash, _, err := aeternity.SignEncodeTxStr(acc, tx, aeternity.Config.Node.NetworkID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = aeClient.BroadcastTransaction(signedTxStr)
	if err != nil {
		panic(err)
	}

	return hash, nil

}

func getHeight(aeClient *aeternity.Ae) (h uint64) {
	h, err := aeClient.APIGetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(aeClient *aeternity.Ae, hash string) (err error) {
	height := getHeight(aeClient)
	fmt.Println("Waiting for Transaction...")
	height, blockHash, microBlockHash, _, err := aeClient.WaitForTransactionUntilHeight(height+10, hash)
	if err != nil {
		// Sometimes, the tests want the tx to fail. Return the err to let them know.
		return err
	}
	fmt.Println("Transaction was found at", height, "blockhash", blockHash, "microBlockHash", microBlockHash, "err", err)
	return nil
}

func getNameEntry(aeClient *aeternity.Ae, name string) (responseJSON string, err error) {
	response, err := aeClient.APIGetNameEntryByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := response.MarshalBinary()
	responseJSON = string(r)
	return responseJSON, nil
}

func TestAENSWorkflow(t *testing.T) {
	name := "fdsa.test"
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeClient := aeternity.NewCli(nodeURL, false).WithAccount(acc)
	aeternity.Config.Node.NetworkID = networkID
	aeternity.Config.Client.Fee = *utils.RequireBigIntFromString("100000000000000")

	// Preclaim the name
	fmt.Println("PreclaimTx")
	preclaimTx, salt, err := aeClient.Aens.NamePreclaimTx(name, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	preclaimTxStr, _ := aeternity.BaseEncodeTx(&preclaimTx)
	fmt.Println("PreclaimTx and Salt:", preclaimTxStr, salt)

	hash, err := signBroadcast(preclaimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NamePreclaimTx", hash)

	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Claim the name
	fmt.Println("NameClaimTx")
	claimTx, err := aeClient.Aens.NameClaimTx(name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	claimTxStr, _ := aeternity.BaseEncodeTx(&claimTx)
	fmt.Println("ClaimTx:", claimTxStr)

	hash, err = signBroadcast(claimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameClaimTx")

	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Verify that the name exists
	entryAfterNameClaim, err := getNameEntry(aeClient, name)
	fmt.Println(entryAfterNameClaim)

	// Update the name, make it point to something
	fmt.Println("NameUpdateTx")
	updateTx, err := aeClient.Aens.NameUpdateTx(name, acc.Address)
	updateTxStr, _ := aeternity.BaseEncodeTx(&updateTx)
	fmt.Println("UpdateTx:", updateTxStr)

	_, err = signBroadcast(updateTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameUpdateTx")

	// Verify that the name was updated
	// Sleep a little, it takes time for the entry update to show up
	fmt.Printf("Sleeping a bit before querying /names/%s...\n", name)
	time.Sleep(1000 * time.Millisecond)
	entryAfterNameUpdate, _ := getNameEntry(aeClient, name)
	fmt.Println(entryAfterNameUpdate)

	if !strings.Contains(entryAfterNameUpdate, acc.Address) {
		t.Errorf("The AENS entry should now point to %s but doesn't: %s", acc.Address, entryAfterNameUpdate)
	}

	// Transfer the name to a recipient
	fmt.Println("NameTransferTx")
	acc2, err := aeternity.AccountFromHexString(recipientPrivateKey)
	if err != nil {
		t.Error(err)
	}
	transferTx, err := aeClient.Aens.NameTransferTx(name, acc2.Address)
	transferTxStr, _ := aeternity.BaseEncodeTx(&transferTx)
	fmt.Println("TransferTx:", transferTxStr)
	hash, err = signBroadcast(transferTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameTransferTx")
	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Receiver updates the name, makes it point to himself
	aeClient2 := aeternity.NewCli(nodeURL, false).WithAccount(acc2)
	fmt.Println("NameUpdateTx Signed By Recipient")
	updateTx2, err := aeClient2.Aens.NameUpdateTx(name, acc2.Address)
	updateTx2Str, _ := aeternity.BaseEncodeTx(&updateTx2)
	fmt.Println("UpdateTx:", updateTx2Str)

	_, err = signBroadcast(updateTx2Str, acc2, aeClient2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameUpdateTx Signed By Recipient")

	// Revoke the name - shouldn't work because it is signed by the sender, who no longer owns the address
	fmt.Println("NameRevokeTx")
	revokeTx, err := aeClient.Aens.NameRevokeTx(name, acc.Address)
	revokeTxStr, _ := aeternity.BaseEncodeTx(&revokeTx)
	fmt.Println("RevokeTx:", revokeTxStr)
	hash, err = signBroadcast(revokeTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameRevokeTx")
	// Wait for a bit
	revokeTxShouldHaveFailed := waitForTransaction(aeClient, hash)
	if revokeTxShouldHaveFailed == nil {
		t.Error("After transferring the name to Recipient, the Sender should not have been able to revoke the name")
	} else {
		fmt.Println(revokeTxShouldHaveFailed)
	}

	// Revoke the name - signed by the recipient
	fmt.Println("NameRevokeTx Signed By Recipient")
	revokeTx2, err := aeClient2.Aens.NameRevokeTx(name, acc2.Address)
	revokeTx2Str, _ := aeternity.BaseEncodeTx(&revokeTx2)
	fmt.Println("RevokeTx Signed By Recipient:", revokeTx2Str)
	hash, err = signBroadcast(revokeTx2Str, acc2, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameRevokeTx Signed By Recipient")
	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

}

func TestOracleWorkflow(t *testing.T) {
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeternity.Config.Node.NetworkID = networkID
	aeClient := aeternity.NewCli(nodeURL, false).WithAccount(acc)

	fmt.Println("OracleRegisterTx")
	queryFee := utils.NewBigIntFromUint64(1000)
	oracleRegisterTx, err := aeClient.Oracle.OracleRegisterTx("hello", "helloback", *queryFee, 0, 100, 0, 0)
	if err != nil {
		t.Error(err)
	}
	oracleRegisterTxStr, _ := aeternity.BaseEncodeTx(&oracleRegisterTx)
	oracleRegisterTxHash, err := signBroadcast(oracleRegisterTxStr, acc, aeClient)
	if err != nil {
		t.Error(err)
	}

	_ = waitForTransaction(aeClient, oracleRegisterTxHash)

	// Confirm that the oracle exists
	oraclePubKey := strings.Replace(acc.Address, "ak_", "ok_", 1)
	oracle, err := aeClient.APIGetOracleByPubkey(oraclePubKey)
	if err != nil {
		t.Errorf("APIGetOracleByPubkey: %s", err)
	}

	fmt.Println("OracleExtendTx")
	// save the oracle's initial TTL so we can compare it with after OracleExtendTx
	oracleTTL := *oracle.TTL
	oracleExtendTx, err := aeClient.Oracle.OracleExtendTx(oraclePubKey, 0, 1000)
	if err != nil {
		t.Error(err)
	}
	oracleExtendTxStr, _ := aeternity.BaseEncodeTx(&oracleExtendTx)
	oracleExtendTxHash, err := signBroadcast(oracleExtendTxStr, acc, aeClient)
	if err != nil {
		t.Error(err)
	}
	_ = waitForTransaction(aeClient, oracleExtendTxHash)

	oracle, err = aeClient.APIGetOracleByPubkey(oraclePubKey)
	if err != nil {
		t.Errorf("APIGetOracleByPubkey: %s", err)
	}
	if *oracle.TTL == oracleTTL {
		t.Errorf("The Oracle's TTL did not change after OracleExtendTx. Got %v but expected %v", *oracle.TTL, oracleTTL)
	}

	fmt.Println("OracleQueryTx")
	oracleQueryTx, err := aeClient.Oracle.OracleQueryTx(oraclePubKey, "How was your day?", *queryFee, 0, 100, 0, 100)
	if err != nil {
		t.Error(err)
	}
	oracleQueryTxStr, _ := aeternity.BaseEncodeTx(&oracleQueryTx)
	oracleQueryTxHash, err := signBroadcast(oracleQueryTxStr, acc, aeClient)
	if err != nil {
		t.Error(err)
	}
	_ = waitForTransaction(aeClient, oracleQueryTxHash)

	fmt.Println("OracleRespondTx")
	// Find the Oracle Query ID to reply to
	oracleQueries, err := aeClient.APIGetOracleQueriesByPubkey(oraclePubKey)
	if err != nil {
		t.Errorf("APIGetOracleQueriesByPubkey: %s", err)
	}
	oqID := string(oracleQueries.OracleQueries[0].ID)
	oracleRespondTx, err := aeClient.Oracle.OracleRespondTx(oraclePubKey, oqID, "My day was fine thank you", 0, 100)
	oracleRespondTxStr, _ := aeternity.BaseEncodeTx(&oracleRespondTx)
	oracleRespondTxHash, err := signBroadcast(oracleRespondTxStr, acc, aeClient)
	if err != nil {
		t.Error(err)
	}
	_ = waitForTransaction(aeClient, oracleRespondTxHash)

}

func TestContracts(t *testing.T) {
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeternity.Config.Node.NetworkID = networkID
	aeClient := aeternity.NewCli(nodeURL, false).WithAccount(acc)

	ttl, nonce, _ := aeClient.GetTTLNonce(acc.Address, aeternity.Config.Client.TTL)
	code := "cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib"
	callData := "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA=="
	fmt.Println("Contract Create")
	tx := aeternity.NewContractCreateTx(acc.Address, nonce, code, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.Contracts.Deposit, aeternity.Config.Client.Contracts.Amount, *utils.NewBigIntFromUint64(1e6), aeternity.Config.Client.Contracts.GasPrice, *utils.NewBigIntFromUint64(1.2e17), ttl, callData)

	ctID, _ := tx.ContractID()
	txStr, _ := aeternity.BaseEncodeTx(&tx)
	fmt.Printf("%#v\n", tx)
	fmt.Println(ctID)
	txHash, err := signBroadcast(txStr, acc, aeClient)
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
}
