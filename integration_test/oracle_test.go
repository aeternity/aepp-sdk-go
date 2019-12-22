package integrationtest

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

func myFunction(query string) (response string, err error) {
	return "I don't care what you say, I am an oracle~!", nil
}
func TestOracleHLL(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	ctx := aeternity.NewContext(alice, node)
	oracle := aeternity.NewOracle(myFunction, node, ctx, "", "", config.Tuning.ChainPollInterval)
	oracle.Listen()
}

func TestOracleWorkflow(t *testing.T) {
	alice, _ := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	ctx := aeternity.NewContext(alice, node)

	// Setup temporary test account and fund it
	testAccount, err := account.New()
	if err != nil {
		t.Fatal(err)
	}
	fundAccount(t, node, alice, testAccount, big.NewInt(1000000000000000000))

	// Register
	register, err := transactions.NewOracleRegisterTx(testAccount.Address, "hello", "helloback", config.Client.Oracles.QueryFee, config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, config.Client.Oracles.ABIVersion, ctx.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Register %+v\n", register)
	_, err = ctx.SignBroadcastWait(register, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm that the oracle exists
	oraclePubKey := register.ID()
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
	extend, err := transactions.NewOracleExtendTx(testAccount.Address, oraclePubKey, config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, ctx.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Extend %+v\n", extend)
	_, err = ctx.SignBroadcastWait(extend, config.Client.WaitBlocks)
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
	query, err := transactions.NewOracleQueryTx(testAccount.Address, oraclePubKey, "How was your day?", config.Client.Oracles.QueryFee, 0, 100, 0, 100, ctx.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Query %+v\n", query)
	_, err = ctx.SignBroadcastWait(query, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

	// Find the Oracle Query ID to reply to
	oqID, err := query.ID()
	if err != nil {
		t.Fatal(err)
	}

	// Respond
	respond, err := transactions.NewOracleRespondTx(testAccount.Address, oraclePubKey, oqID, "My day was fine thank you", config.OracleTTLTypeDelta, config.Client.Oracles.ResponseTTLValue, ctx.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Respond %+v\n", respond)
	_, err = ctx.SignBroadcastWait(respond, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
}
