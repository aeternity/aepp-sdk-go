package aeternity

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/aeternity/aepp-sdk-go/v7/utils"
)

type mockNode struct {
	i uint64
}

func (m *mockNode) GetHeight() (uint64, error) {
	m.i++
	return m.i, nil
}

// GetTransactionByHash pretends that the transaction was not mined until block 9, and this is only visible when the mockClient is at height 10.
func (m *mockNode) GetTransactionByHash(hash string) (tx *models.GenericSignedTx, err error) {
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

func TestTxReceipt_Watch(t *testing.T) {
	config.Tuning.ChainPollInterval = 1 * time.Microsecond
	type args struct {
		mined      chan bool
		waitBlocks uint64
		node       transactionWaiter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Tx was mined successfully",
			args: args{
				mined:      make(chan bool),
				waitBlocks: 10,
				node:       &mockNode{},
			},
			wantErr: false,
		},
		{
			name: "Tx did not get mined",
			args: args{
				mined:      make(chan bool),
				waitBlocks: 10,
				node:       &mockNode{i: 20},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &transactions.SpendTx{}
			txReceipt := NewTxReceipt(tx, "tx_signed", "th_somehash", "sg_somesignature")

			go txReceipt.Watch(tt.args.mined, tt.args.waitBlocks, tt.args.node)
			result := <-tt.args.mined
			if result && tt.wantErr {
				t.Fatal(txReceipt.Error)
			}
		})
	}
}

func Example() {
	// Set the Network ID. For this example, setting the config.Node.NetworkID
	// is actually not needed - but if you have other code that also needs to
	// access NetworkID somehow, do it this way.
	config.Node.NetworkID = config.NetworkIDTestnet

	alice, err := account.FromHexString("deadbeef")
	if err != nil {
		fmt.Println("Could not create alice's Account:", err)
	}

	bobAddress := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

	// create a connection to a node, represented by *Node
	node := naet.NewNode("http://localhost:3013", false)

	// create the closures that autofill the correct account nonce and transaction TTL
	ttlnoncer := transactions.NewTTLNoncer(node)

	// create the SpendTransaction
	msg := "Reason For Payment"
	tx, err := transactions.NewSpendTx(alice.Address, bobAddress, big.NewInt(1e9), []byte(msg), ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	spendTxReceipt, err := SignBroadcast(tx, alice, node, config.Node.NetworkID)
	if err != nil {
		fmt.Println("could not send transaction:", err)
	}
	err = WaitSynchronous(spendTxReceipt, 10, node)
	if err != nil {
		fmt.Println("transaction was not accepted by the blockchain:", err)
	}
	fmt.Printf("%#v\n", spendTxReceipt)

	// check the recipient's balance
	time.Sleep(2 * time.Second)
	bobState, err := node.GetAccount(bobAddress)
	if err != nil {
		fmt.Println("Couldn't get Bob's account data:", err)
	}

	fmt.Println(bobState.Balance)
}
