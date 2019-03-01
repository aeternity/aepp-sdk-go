package aeternity_test

import (
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestSpendTxStr(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	amount := utils.BigInt{Int: &big.Int{}}
	amount.SetInt64(10)
	fee := utils.BigInt{Int: &big.Int{}}
	fee.SetInt64(10)
	message := "Hello World"
	ttl := uint64(10)
	nonce := uint64(1)
	eBase64TxMsg := "tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCgoKAYtIZWxsbyBXb3JsZPSZjdM="

	base64TxMsg, err := aeternity.SpendTxStr(sender, recipient, amount, fee, ttl, nonce, message)
	if err != nil {
		t.Fatalf("SpendTx could not create a SpendTransaction: %v", err)
	}
	if base64TxMsg != eBase64TxMsg {
		t.Fatalf("SpendTx returned a wrong tx_ blob, got %s, want %s", base64TxMsg, eBase64TxMsg)
	}
}
