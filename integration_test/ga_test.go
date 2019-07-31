package integrationtest

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
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

	// Read the auth contract from a file, compile and prepare its init() calldata
	authSource := readFile(t, "authorize.aes")
	authBytecode, err := compiler.CompileContract(authSource)
	if err != nil {
		t.Fatal(err)
	}
	authCalldata, err := compiler.EncodeCalldata(authSource, "init", []string{alice.Address})
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
	gaTx := aeternity.NewGAAttachTx(testAccount.Address, 1, authBytecode, auth.TypeInfo[0].FuncHash, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.BaseGas, aeternity.Config.Client.GasPrice, aeternity.Config.Client.Fee, ttl, authCalldata)
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

	// GAMetaTx with SpendTx
	ttl, err = h.GetTTL(aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatal(err)
	}

	// SpendTx - wrap it in SignedTx with 0 signatures before including in GAMetaTx
	spendTx := aeternity.NewSpendTx(testAccount.Address, bob.Address, *big.NewInt(5000), aeternity.Config.Client.Fee, []byte{}, ttl, 0)

	spendTxRLPBytes, err := EncodeRLPToBytes(&spendTx)
	if err != nil {
		t.Fatal(err)
	}

	gaMetaTx := aeternity.NewGAMetaTx(testAccount.Address, "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACArgmMvLPdJq0/eccQZx/kn0CmjZNS2PRRsu167v6sUZQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABdclTA==", aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.BaseGas, aeternity.Config.Client.GasPrice, aeternity.Config.Client.Fee, ttl, &spendTx)

	w := &bytes.Buffer{}
	err = rlp.Encode(w, &gaMetaTx)
	if err != nil {
		t.Fatal(err)
	}

	gaMetaTxRaw, err := aeternity.CreateSignedTransaction(w.Bytes(), [][]byte{})
	if err != nil {
		t.Fatal(err)
	}

	gaMetaTxStr := aeternity.Encode(aeternity.PrefixTransaction, gaMetaTxRaw)
	fmt.Println("GAMetaTx", gaMetaTxStr)
	err = aeNode.PostTransaction(gaMetaTxStr, "th_fdsa")
	if err != nil {
		t.Error(err)
	}

}
