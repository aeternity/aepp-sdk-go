package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/aeternity"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
)

func TestRegisterName(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	alice, _ := setupAccounts(t)

	name := "somelongnamefdsafdffsa.chain"
	nameFee := transactions.CalculateMinNameFee(name)
	_, _, _, _, _, err := aeternity.RegisterName(n, alice, name, nameFee)
	if err != nil {
		t.Error(err)
	}
}
