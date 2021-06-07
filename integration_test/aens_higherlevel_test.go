package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v8/aeternity"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
)

func TestRegisterName(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	alice, _ := setupAccounts(t)

	ctx := aeternity.NewContext(alice, n)
	name := randomName(22)
	nameFee := transactions.CalculateMinNameFee(name)

	aens := aeternity.NewAENS(ctx)
	_, err := aens.RegisterName(name, nameFee)
	if err != nil {
		t.Error(err)
	}
}
