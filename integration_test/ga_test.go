package integrationtest

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/golden"
	"github.com/aeternity/aepp-sdk-go/utils"
	rlp "github.com/randomshinichi/rlpae"
)

func EncodeRLPToBytes(tx rlp.Encoder) (b []byte, err error) {
	w := new(bytes.Buffer)
	err = rlp.Encode(w, tx)
	if err != nil {
		return
	}
	return w.Bytes(), nil
}

func TestGeneralizedAccounts(t *testing.T) {
	alice, bob := setupAccounts(t)
	aeNode := setupNetwork(t, privatenetURL, false)
	compiler := aeternity.NewCompiler(aeternity.Config.Client.Contracts.CompilerURL, false)
	h := aeternity.Helpers{Node: aeNode}

	// Take note of Bob's balance, and after this test, we expect it to have this much more AE
	amount := utils.NewIntFromUint64(5000)
	expected := new(big.Int)
	bobState, err := aeNode.GetAccount(bob.Address)
	if err != nil {
		expected.Set(amount)
	} else {
		bS := big.Int(bobState.Balance)
		expected.Add(&bS, amount)
	}

	// Read the auth contract from a file, compile and prepare its init() calldata
	authBytecode, err := compiler.CompileContract(golden.AuthorizeSource)
	if err != nil {
		t.Fatal(err)
	}
	authInitCalldata, err := compiler.EncodeCalldata(golden.AuthorizeSource, "init", []string{alice.Address})
	if err != nil {
		t.Fatal(err)
	}

	// Create throwaway test account, fund it and ensure it is a POA
	testAccount, err := aeternity.NewAccount()
	if err != nil {
		t.Fatal(err)
	}
	fundAccount(t, aeNode, alice, testAccount, big.NewInt(1000000000000000000))
	testAccountState, err := aeNode.GetAccount(testAccount.Address)
	if err != nil {
		t.Fatal(err)
	}
	if testAccountState.Kind != "basic" {
		t.Fatalf("%s is supposed to be a basic account but wasn't", testAccount.Address)
	}

	// GAAttachTx
	// Create a Contract{} struct from the compiled bytecode to get its authfunc hash
	auth, err := aeternity.NewContractFromString(authBytecode)
	if err != nil {
		t.Fatal(err)
	}
	ttl, err := h.GetTTL(aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatal(err)
	}
	gaTx := aeternity.NewGAAttachTx(testAccount.Address, 1, authBytecode, auth.TypeInfo[0].FuncHash, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.BaseGas, aeternity.Config.Client.GasPrice, aeternity.Config.Client.Fee, ttl, authInitCalldata)
	txHash := signBroadcast(t, &gaTx, testAccount, aeNode)
	_, _, err = waitForTransaction(aeNode, txHash)
	if err != nil {
		t.Error(err)
	}

	// The test account should now be a generalized account
	checkGeneralizedAccount := func() {
		testAccountState, err = aeNode.GetAccount(testAccount.Address)
		if err != nil {
			t.Fatal(err)
		}
	}
	delay(checkGeneralizedAccount)
	if testAccountState.Kind != "generalized" {
		t.Fatalf("%s was supposed to be a generalized account but isn't", testAccount.Address)
	}

	// GAMetaTx
	// Get the TTL (not really needed, could be 0 too)
	ttl, err = h.GetTTL(aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatal(err)
	}

	// spendTx will be wrapped in a SignedTx with 0 signatures before being
	// included in GAMetaTx. The constructor NewGAMetaTx() does this for you.
	// authData is authorize(3)
	authData := "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBGtXufEG2HuMYcRcNwsGAeqymslunKf692bHnvwI5K6wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAU3aKBNm"
	gas := utils.NewIntFromUint64(10000) // the node will fail the authentication if there isn't enough gas
	spendTx := aeternity.NewSpendTx(testAccount.Address, bob.Address, *big.NewInt(5000), aeternity.Config.Client.Fee, []byte{}, ttl, 0)
	gaMetaTx := aeternity.NewGAMetaTx(testAccount.Address, authData, aeternity.Config.Client.Contracts.ABIVersion, *gas, aeternity.Config.Client.GasPrice, aeternity.Config.Client.Fee, ttl, &spendTx)

	gaMetaTxFinal, hash, _, err := aeternity.SignHashTx(testAccount, &gaMetaTx, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Fatal(err)
	}

	gaMetaTxStr, err := aeternity.SerializeTx(&gaMetaTxFinal)
	if err != nil {
		t.Fatal(err)
	}
	err = aeNode.PostTransaction(gaMetaTxStr, hash)
	if err != nil {
		t.Fatal(err)
	}

	// check bob.Address, make sure it got the SpendTx
	getBobsAccount := func() {
		bobState, err = aeNode.GetAccount(bob.Address)
		if err != nil {
			t.Fatalf("Couldn't get Bob's account data: %v", err)
		}
	}
	delay(getBobsAccount)
	b := big.Int(bobState.Balance)

	if expected.Cmp(&b) != 0 {
		t.Fatalf("Bob should have %v, but has %v instead", expected.String(), bobState.Balance.String())
	}
}
