package aeternity_test

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestOracleRegisterTx(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	fee := utils.NewBigIntFromUint64(100)
	txRaw, err := aeternity.OracleRegisterTx(sender, 0, "likethis", "likethat", 123456789, "delta", 1, *fee, 1, 1)
	if err != nil {
		t.Errorf("Could not create OracleRegisterTx: %s", err)
	}
	txStr := aeternity.Encode(aeternity.PrefixTransaction, txRaw)
	fmt.Println(txStr)
}
