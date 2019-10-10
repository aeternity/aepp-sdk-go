package integrationtest

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/aeternity"
	"github.com/aeternity/aepp-sdk-go/v6/swagguard/node/models"
)

func TestOracleWorkflow(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)

	// Setup temporary test account and fund it
	testAccount, err := account.New()
	if err != nil {
		t.Fatal(err)
	}
	fundAccount(t, node, alice, testAccount, big.NewInt(1000000000000000000))
	oracleAccount := aeternity.NewContextFromNode(node, testAccount.Address)

	// Register
	register, err := oracleAccount.OracleRegisterTx("hello", "helloback", config.Client.Oracles.QueryFee, config.Client.Oracles.QueryTTLType, config.Client.Oracles.QueryTTLValue, config.Client.Oracles.VMVersion)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Register %+v\n", register)
	_, registerHash, _, err := aeternity.SignBroadcastTransaction(register, testAccount, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = waitForTransaction(node, registerHash)

	// Confirm that the oracle exists
	oraclePubKey := strings.Replace(testAccount.Address, "ak_", "ok_", 1)
	var oracle *models.RegisteredOracle
	getOracle := func() {
		oracle, err = node.GetOracleByPubkey(oraclePubKey)
		if err != nil {
			t.Fatalf("APIGetOracleByPubkey: %s", err)
		}
	}
	delay(getOracle)

	// Extend
	// save the oracle's initial TTL so we can compare it with after OracleExtendTx
	oracleTTL := oracle.TTL
	extend, err := oracleAccount.OracleExtendTx(oraclePubKey, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Extend %+v\n", extend)
	_, extendHash, _, err := aeternity.SignBroadcastTransaction(extend, testAccount, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = waitForTransaction(node, extendHash)

	// Confirm that the oracle's TTL changed
	oracle, err = node.GetOracleByPubkey(oraclePubKey)
	if err != nil {
		t.Fatalf("APIGetOracleByPubkey: %s", err)
	}
	if oracle.TTL == oracleTTL {
		t.Fatalf("The Oracle's TTL did not change after OracleExtendTx. Got %v but expected %v", oracle.TTL, oracleTTL)
	}

	// Query
	query, err := oracleAccount.OracleQueryTx(oraclePubKey, "How was your day?", config.Client.Oracles.QueryFee, 0, 100, 0, 100)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Query %+v\n", query)
	_, queryHash, _, err := aeternity.SignBroadcastTransaction(query, testAccount, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = waitForTransaction(node, queryHash)

	// Find the Oracle Query ID to reply to
	fmt.Println("Sleeping a bit before querying node for OracleID")

	var oracleQueries *models.OracleQueries
	getOracleQueries := func() {
		oracleQueries, err = node.GetOracleQueriesByPubkey(oraclePubKey)
		if err != nil {
			t.Fatalf("APIGetOracleQueriesByPubkey: %s", err)
		}
	}
	delay(getOracleQueries)
	oqID := oracleQueries.OracleQueries[0].ID

	// Respond
	respond, err := oracleAccount.OracleRespondTx(oraclePubKey, *oqID, "My day was fine thank you", 0, 100)
	fmt.Printf("Respond %+v\n", respond)
	_, respondHash, _, err := aeternity.SignBroadcastTransaction(respond, testAccount, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = waitForTransaction(node, respondHash)
}
