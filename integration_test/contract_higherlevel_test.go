package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
)

func TestCreateContract(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	c := naet.NewCompiler("http://localhost:3080", false)
	alice, _ := setupAccounts(t)

	simplestorage := `
contract SimpleStorage =
  record state = { data : int }
  entrypoint init(value : int) : state = { data = value }
  function get() : int = state.data
  stateful function set(value : int) = put(state{data = value})`

	_, err := aeternity.CreateContract(n, c, alice, simplestorage, "init", []string{"42"}, config.CompilerBackendFATE)
	if err != nil {
		t.Error(err)
	}
}

func TestCallContract(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	c := naet.NewCompiler("http://localhost:3080", false)
	alice, _ := setupAccounts(t)

	simplestorage := `
contract SimpleStorage =
  record state = { data : int }
  entrypoint init(value : int) : state = { data = value }
  function get() : int = state.data
  stateful function set(value : int) = put(state{data = value})`

	ctID, _, err := aeternity.CreateContract(n, c, alice, simplestorage, "init", []string{"42"}, config.CompilerBackendFATE)
	if err != nil {
		t.Fatal(err)
	}

	err = aeternity.CallContract(n, c, alice, ctID, simplestorage, "get", []string{}, config.CompilerBackendFATE)
	if err != nil {
		t.Error(err)
	}
}
