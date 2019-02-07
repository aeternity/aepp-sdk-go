package aeternity_test

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestSpendTransaction(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"
	eBase64TxMsg := "tx_+FgMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIAAHi0hlbGxvIFdvcmxk0pjHOg=="

	base64TxMsg, _, _, err := aeternity.SpendTransaction(sender, recipient, 10, 10, message)
	if err != nil {
		t.Errorf("SpendTransaction errored out: %v", err)
	}
	if base64TxMsg != eBase64TxMsg {
		t.Errorf("SpendTransaction returned a wrong tx_ blob, got %s, want %s", base64TxMsg, eBase64TxMsg)
	}

	eBase64Tx := "tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIAAHgDJ1jfU="
	base64Tx, _, _, err := aeternity.SpendTransaction(sender, recipient, 10, 10, "")
	if err != nil {
		t.Errorf("SpendTransaction errored out: %v", err)
	}
	if base64Tx != eBase64Tx {
		t.Errorf("SpendTransaction without a message returned a wrong tx_ blob, got %s, want %s", base64TxMsg, eBase64TxMsg)
	}

}
