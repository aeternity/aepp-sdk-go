package aeternity

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v8/account"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
	"github.com/aeternity/aepp-sdk-go/v8/utils"
)

type mockNodeForTxReceiptWatch struct {
	i uint64
}

func (m *mockNodeForTxReceiptWatch) GetHeight() (uint64, error) {
	m.i++
	return m.i, nil
}

// GetTransactionByHash pretends that the transaction was not mined until block 9, and this is only visible when the mockClient is at height 10.
func (m *mockNodeForTxReceiptWatch) GetTransactionByHash(hash string) (tx *models.GenericSignedTx, err error) {
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
				node:       &mockNodeForTxReceiptWatch{},
			},
			wantErr: false,
		},
		{
			name: "Tx did not get mined",
			args: args{
				mined:      make(chan bool),
				waitBlocks: 10,
				node:       &mockNodeForTxReceiptWatch{i: 20},
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

	bobAddress := "ak_wJ3iKZcqvgdnQ6YVz8pY2xPjtVTNNEL61qF4AYQdksZfXZLks"

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

type mockNodeForSignBroadcast struct {
	shouldAcceptTx bool
}

func (m *mockNodeForSignBroadcast) PostTransaction(string, string) error {
	if m.shouldAcceptTx {
		return nil
	}
	return fmt.Errorf("Dummy PostTransaction error")
}
func TestSignBroadcast(t *testing.T) {
	dummyAcc, err := account.New()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		tx             transactions.Transaction
		signingAccount *account.Account
		n              naet.PostTransactioner
		networkID      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Everything okay, transaction was accepted",
			args: args{
				tx: &transactions.SpendTx{
					SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
					RecipientID: "ak_wJ3iKZcqvgdnQ6YVz8pY2xPjtVTNNEL61qF4AYQdksZfXZLks",
					Amount:      &big.Int{},
					Fee:         &big.Int{},
					Payload:     nil,
					TTL:         0,
					Nonce:       0,
				},
				signingAccount: dummyAcc,
				n:              &mockNodeForSignBroadcast{shouldAcceptTx: true},
				networkID:      "dummy_network_id",
			},
			wantErr: false,
		},
		{
			name: "Error, transaction not accepted",
			args: args{
				tx: &transactions.SpendTx{
					SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
					RecipientID: "ak_wJ3iKZcqvgdnQ6YVz8pY2xPjtVTNNEL61qF4AYQdksZfXZLks",
					Amount:      &big.Int{},
					Fee:         &big.Int{},
					Payload:     nil,
					TTL:         0,
					Nonce:       0,
				},
				signingAccount: dummyAcc,
				n:              &mockNodeForSignBroadcast{shouldAcceptTx: false},
				networkID:      "dummy_network_id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SignBroadcast(tt.args.tx, tt.args.signingAccount, tt.args.n, tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignBroadcast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_findVMABIVersion(t *testing.T) {
	type args struct {
		nodeVersion     string
		compilerBackend string
	}
	tests := []struct {
		name           string
		args           args
		wantVMVersion  uint16
		wantABIVersion uint16
		wantErr        bool
	}{
		{
			name: "node version 5, FATE backend",
			args: args{
				nodeVersion:     "5",
				compilerBackend: "fate",
			},
			wantVMVersion:  5,
			wantABIVersion: 3,
		},
		{
			name: "node version 5, AEVM backend",
			args: args{
				nodeVersion:     "5",
				compilerBackend: "aevm",
			},
			wantVMVersion:  6,
			wantABIVersion: 1,
		},
		{
			name: "node version 4, AEVM backend",
			args: args{
				nodeVersion:     "4",
				compilerBackend: "aevm",
			},
			wantVMVersion:  4,
			wantABIVersion: 1,
		},
		{
			name: "node version 4, does not actually support FATE, so it should return answer for AEVM anyway",
			args: args{
				nodeVersion:     "4",
				compilerBackend: "fate",
			},
			wantVMVersion:  4,
			wantABIVersion: 1,
		},
		{
			name: "Other versions of the node are not supported",
			args: args{
				nodeVersion:     "3",
				compilerBackend: "aevm",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVMVersion, gotABIVersion, err := findVMABIVersion(tt.args.nodeVersion, tt.args.compilerBackend)
			if (err != nil) != tt.wantErr {
				t.Errorf("findVMABIVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVMVersion != tt.wantVMVersion {
				t.Errorf("findVMABIVersion() gotVMVersion = %v, want %v", gotVMVersion, tt.wantVMVersion)
			}
			if gotABIVersion != tt.wantABIVersion {
				t.Errorf("findVMABIVersion() gotABIVersion = %v, want %v", gotABIVersion, tt.wantABIVersion)
			}
		})
	}
}
