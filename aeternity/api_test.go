package aeternity_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestSubmitSpendTransactionToTestnet(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	rawTx, err := aeternity.SpendTx(sender, recipient, message, 10, 10, 1, 1)
	if err != nil {
		fmt.Printf("Could not create raw SpendTransaction because %v", err)
	}

	ae := aeternity.NewCli("http://localhost:3013", true)
	genesis, err := aeternity.AccountFromHexString(senderPrivateKey)

	signedEncodedTx, signedEncodedTxHash, _, _ := aeternity.SignEncodeTx(genesis, rawTx, "ae_docker")

	err = ae.APIPostTransaction(signedEncodedTx, signedEncodedTxHash)
	if err != nil {
		fmt.Printf("Could not POST signed SpendTransaction because %v", err)
	}

}
