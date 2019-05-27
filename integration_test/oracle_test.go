package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestOracleWorkflow(t *testing.T) {
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeternity.Config.Node.NetworkID = networkID
	aeClient := aeternity.NewClient(nodeURL, false)

	oracleAlice := aeternity.Oracle{Client: aeClient, Account: acc}

	fmt.Println("OracleRegisterTx")
	queryFee := utils.NewBigIntFromUint64(1000)
	oracleRegisterTx, err := oracleAlice.OracleRegisterTx("hello", "helloback", *queryFee, 0, 100, 0, 0)
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
	oracleExtendTx, err := oracleAlice.OracleExtendTx(oraclePubKey, 0, 1000)
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
	oracleQueryTx, err := oracleAlice.OracleQueryTx(oraclePubKey, "How was your day?", *queryFee, 0, 100, 0, 100)
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
	oracleRespondTx, err := oracleAlice.OracleRespondTx(oraclePubKey, oqID, "My day was fine thank you", 0, 100)
	oracleRespondTxStr, _ := aeternity.BaseEncodeTx(&oracleRespondTx)
	oracleRespondTxHash, err := signBroadcast(oracleRespondTxStr, acc, aeClient)
	if err != nil {
		t.Error(err)
	}
	_ = waitForTransaction(aeClient, oracleRespondTxHash)

}
