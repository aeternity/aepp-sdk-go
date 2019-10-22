package integrationtest

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/aeternity"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
)

func TestOracleWorkflow(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	// Setup temporary test account and fund it
	testAccount, err := account.New()
	if err != nil {
		t.Fatal(err)
	}
	fundAccount(t, node, alice, testAccount, big.NewInt(1000000000000000000))

	// Register
	register, err := transactions.NewOracleRegisterTx(testAccount.Address, "hello", "helloback", config.Client.Oracles.QueryFee, config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, config.Client.Oracles.ABIVersion, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Register %+v\n", register)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(register, testAccount, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

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
	extend, err := transactions.NewOracleExtendTx(testAccount.Address, oraclePubKey, config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Extend %+v\n", extend)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(extend, testAccount, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm that the oracle's TTL changed
	oracle, err = node.GetOracleByPubkey(oraclePubKey)
	if err != nil {
		t.Fatalf("APIGetOracleByPubkey: %s", err)
	}
	if oracle.TTL == oracleTTL {
		t.Fatalf("The Oracle's TTL did not change after OracleExtendTx. Got %v but expected %v", oracle.TTL, oracleTTL)
	}

	// Query
	query, err := transactions.NewOracleQueryTx(testAccount.Address, oraclePubKey, "How was your day?", config.Client.Oracles.QueryFee, 0, 100, 0, 100, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Query %+v\n", query)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(query, testAccount, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

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
	fmt.Println(oraclePubKey, *oqID)

	// Respond
	respond, err := transactions.NewOracleRespondTx(testAccount.Address, oraclePubKey, *oqID, "My day was fine thank you", config.OracleTTLTypeDelta, config.Client.Oracles.ResponseTTLValue, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Respond %+v\n", respond)
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(respond, testAccount, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
}
