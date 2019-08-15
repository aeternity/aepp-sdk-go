package aeternity

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

type mockClient struct {
	i uint64
}

func (m *mockClient) GetHeight() (uint64, error) {
	m.i++
	return m.i, nil
}

// GetTransactionByHash pretends that the transaction was not mined until block 9, and this is only visible when the mockClient is at height 10.
func (m *mockClient) GetTransactionByHash(hash string) (tx *models.GenericSignedTx, err error) {
	unminedHeight, _ := utils.NewIntFromString("-1")
	minedHeight, _ := utils.NewIntFromString("9")

	bh := "bh_someblockhash"
	tx = &models.GenericSignedTx{
		BlockHash:   &bh,
		BlockHeight: utils.BigInt{},
		Hash:        &hash,
		Signatures:  nil,
	}

	if m.i == 10 {
		tx.BlockHeight.Set(minedHeight)
	} else {
		tx.BlockHeight.Set(unminedHeight)
	}
	return tx, nil
}
func TestWaitForTransactionForXBlocks(t *testing.T) {
	m := new(mockClient)
	blockHeight, blockHash, err := WaitForTransactionForXBlocks(m, "th_transactionhash", 10)
	if err != nil {
		t.Fatal(err)
	}
	if blockHeight != 9 {
		t.Fatalf("Expected mock blockHeight 9, got %v", blockHeight)
	}
	if blockHash != "bh_someblockhash" {
		t.Fatalf("Expected mock blockHash bh_someblockhash, got %s", blockHash)
	}
}
