package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestGeneralizedAccounts(t *testing.T) {
	alice, _ := setupAccounts(t)
	aeNode := setupNetwork(t, privatenetURL, false)
	compiler := aeternity.NewCompiler(aeternity.Config.Client.Contracts.CompilerURL, false)

	authSource := readFile(t, "authorize.aes")
	authBytecode, err := compiler.CompileContract(authSource)
	if err != nil {
		t.Fatal(err)
	}
	authCalldata, err := compiler.EncodeCalldata(authSource, "init", []string{alice.Address})
	if err != nil {
		t.Fatal(err)
	}

	auth, err := aeternity.NewContractFromString(authBytecode)
	if err != nil {
		t.Fatal(err)
	}
	gaAlice := aeternity.NewGAAttachTx(alice.Address, 1, authBytecode, auth.TypeInfo[0].FuncHash, aeternity.Config.Client.Contracts.VMVersion, aeternity.Config.Client.Contracts.ABIVersion, aeternity.Config.Client.BaseGas, aeternity.Config.Client.GasPrice, aeternity.Config.Client.Fee, aeternity.Config.Client.TTL, authCalldata)
	txHash := signBroadcast(t, &gaAlice, alice, aeNode)
	_, _, err = waitForTransaction(aeNode, txHash)
	if err != nil {
		t.Fatal(err)
	}
}
