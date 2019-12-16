package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

func TestRegisterName(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	alice, _ := setupAccounts(t)

	ctx, err := aeternity.NewContext(alice, n)
	if err != nil {
		t.Fatal(err)
	}
	name := "somelongnamefdsafdffsa.chain"
	nameFee := transactions.CalculateMinNameFee(name)

	aens := aeternity.NewAENS(ctx)
	_, err = aens.RegisterName(name, nameFee)
	if err != nil {
		t.Error(err)
	}
}
