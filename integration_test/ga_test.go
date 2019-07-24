package integrationtest

import (
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestPOAtoGA(t *testing.T) {
	node := setupNetwork(t, privatenetURL)
	alice, _ := setupAccounts(t)

	aliceState, err := node.GetAccount(alice.Address)
	if err != nil {
		t.Fatal(err)
	}
	if aliceState.Kind != "basic" {
		t.Fatalf("Alice's account should be a basic account at first, but got %s", aliceState.Kind)
	}

	blindAuth := `contract BlindAuth =
	record state = { nonce : int, owner : address }

	function init(owner' : address) = { nonce = 1, owner = owner' }

	stateful function authorize(s : signature) : bool =
	  switch(Auth.tx_hash)
		None          => abort("Not in Auth context")
		Some(tx_hash) => 
			put(state{ nonce = state.nonce + 1 })
			true

	function to_sign(h : hash, n : int) : hash =
	  Crypto.blake2b((h, n))

	private function require(b : bool, err : string) =
	  if(!b) abort(err)
	`
	h := aeternity.Helpers{Node: node}
	c := aeternity.NewCompiler(aeternity.Config.Client.Contracts.CompilerURL, false)
	ttl, nonce, err := h.GetTTLNonce(alice.Address, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatal(err)
	}

	bytecode, err := c.CompileContract(blindAuth)
	if err != nil {
		t.Fatal(err)
	}

	tx := aeternity.NewGAAttachTx(
		alice.Address,
		nonce,
		bytecode,
		authfunc,
		aeternity.Config.Client.Contracts.VMVersion,
		aeternity.Config.Client.Contracts.ABIVersion,
		*big.NewInt(200000000000000), // deposit
		*big.NewInt(500),             // gas
		*big.NewInt(1000000000),      // gasprice
		aeternity.Config.Client.Fee,
		ttl,
		calldata)
}
