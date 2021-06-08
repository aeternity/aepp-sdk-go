package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v8/account"
	"github.com/aeternity/aepp-sdk-go/v8/aeternity"
	"github.com/aeternity/aepp-sdk-go/v8/config"
)

func TestNoncer(t *testing.T) {
	n := setupNetwork(t, privatenetURL, false)
	emptyAccount, err := account.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := aeternity.NewContext(emptyAccount, n)
	_, _, accountNonce, err := ctx.TTLNoncer()(emptyAccount.Address, config.Client.TTL)
	if err != nil {
		t.Fatal(err)
	}
	if accountNonce != 0 {
		t.Fatal("Invalid nonce of new account", accountNonce)
	}
}
