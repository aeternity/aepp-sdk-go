package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v9/aeternity"
	"github.com/aeternity/aepp-sdk-go/v9/config"
	"github.com/aeternity/aepp-sdk-go/v9/naet"
	"github.com/aeternity/aepp-sdk-go/v9/swagguard/node/models"
)

func TestCreateContract(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	c := naet.NewCompiler("http://localhost:3080", false)
	alice, _ := setupAccounts(t)
	ctx := aeternity.NewContext(alice, n)
	ctx.SetCompiler(c)
	contract := aeternity.NewContract(ctx)

	simplestorage := `
contract SimpleStorage =
  record state = { data : int }
  entrypoint init(value : int) : state = { data = value }
  function get() : int = state.data
  stateful function set(value : int) = put(state{data = value})`

	ctID, _, err := contract.Deploy(simplestorage, "init", []string{"42"})
	if err != nil {
		t.Error(err)
	}

	_, err = n.GetContractByID(ctID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallContract(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	c := naet.NewCompiler("http://localhost:3080", false)
	alice, _ := setupAccounts(t)
	ctx := aeternity.NewContext(alice, n)
	ctx.SetCompiler(c)
	contract := aeternity.NewContract(ctx)

	simplestorage := `
contract SimpleStorage =
  record state = { data : int }
  entrypoint init(value : int) : state = { data = value }
  function get() : int = state.data
  stateful function set(value : int) = put(state{data = value})`

	ctID, _, err := contract.Deploy(simplestorage, "init", []string{"42"})
	if err != nil {
		t.Fatal(err)
	}

	callReceipt, err := contract.Call(ctID, simplestorage, "get", []string{})
	if err != nil {
		t.Error(err)
	}

	callInfoChan := make(chan *models.ContractCallObject)
	errChan := make(chan error)
	go aeternity.DefaultCallResultListener(n, callReceipt.Hash, callInfoChan, errChan, config.Tuning.ChainPollInterval)
	<-callInfoChan
	err = <-errChan
	if err != nil {
		t.Fatal(err)
	}
}
