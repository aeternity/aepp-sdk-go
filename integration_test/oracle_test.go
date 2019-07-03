package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestOracleWorkflow(t *testing.T) {
	alice, _ := setupAccounts(t)
	client := setupNetwork(t, privatenetURL)

	oracleAlice := aeternity.Context{Client: client, Address: alice.Address}

	// Register
	queryFee := utils.NewIntFromUint64(1000)
	register, err := oracleAlice.OracleRegisterTx("hello", "helloback", *queryFee, uint64(0), uint64(100), 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Register %+v\n", register)
	registerHash := signBroadcast(t, &register, alice, client)
	_, _, _ = waitForTransaction(client, registerHash)

	// Confirm that the oracle exists
	oraclePubKey := strings.Replace(alice.Address, "ak_", "ok_", 1)
	var oracle *models.RegisteredOracle
	getOracle := func() {
		oracle, err = client.GetOracleByPubkey(oraclePubKey)
		if err != nil {
			t.Fatalf("APIGetOracleByPubkey: %s", err)
		}
	}
	delay(getOracle)

	// Extend
	// save the oracle's initial TTL so we can compare it with after OracleExtendTx
	oracleTTL := oracle.TTL
	extend, err := oracleAlice.OracleExtendTx(oraclePubKey, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Extend %+v\n", extend)
	extendHash := signBroadcast(t, &extend, alice, client)
	_, _, _ = waitForTransaction(client, extendHash)

	// Confirm that the oracle's TTL changed
	oracle, err = client.GetOracleByPubkey(oraclePubKey)
	if err != nil {
		t.Fatalf("APIGetOracleByPubkey: %s", err)
	}
	if oracle.TTL == oracleTTL {
		t.Fatalf("The Oracle's TTL did not change after OracleExtendTx. Got %v but expected %v", oracle.TTL, oracleTTL)
	}

	// Query
	query, err := oracleAlice.OracleQueryTx(oraclePubKey, "How was your day?", *queryFee, 0, 100, 0, 100)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Query %+v\n", query)
	queryHash := signBroadcast(t, &query, alice, client)
	_, _, _ = waitForTransaction(client, queryHash)

	// Find the Oracle Query ID to reply to
	fmt.Println("Sleeping a bit before querying node for OracleID")

	var oracleQueries *models.OracleQueries
	getOracleQueries := func() {
		oracleQueries, err = client.GetOracleQueriesByPubkey(oraclePubKey)
		if err != nil {
			t.Fatalf("APIGetOracleQueriesByPubkey: %s", err)
		}
	}
	delay(getOracleQueries)
	oqID := oracleQueries.OracleQueries[0].ID

	// Respond
	respond, err := oracleAlice.OracleRespondTx(oraclePubKey, *oqID, "My day was fine thank you", 0, 100)
	fmt.Printf("Respond %+v\n", respond)
	respondHash := signBroadcast(t, &respond, alice, client)
	_, _, _ = waitForTransaction(client, respondHash)
}
