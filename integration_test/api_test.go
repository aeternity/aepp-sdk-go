package integrationtest

import (
	"flag"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
)

/*
For best coverage, this needs to be run against a local node, which needs to
have the transactions submitted to it first. Which makes it an integration test.

Mainnet/testnet doesn't have all transactions, and even if they did,
oracles/names expire, so you can't test all the endpoints anyway.

However some tests need known data to test against to avoid depending on other
API functions, e.g. GetKeyBlockByHash(). These use testnet instead of depending
on e.g. GetTopBlock()
The flag -testnet enables testnet-dependent tests, otherwise everything is
tested against a local node.
*/

type txInfo struct {
	height         uint64
	txHash, mbHash string
}
type txTypes struct {
	name             string
	oracleID         string
	contractID       string
	SpendTx          txInfo
	NamePreclaimTx   txInfo
	NameClaimTx      txInfo
	NameUpdateTx     txInfo
	NameTransferTx   txInfo
	NameRevokeTx     txInfo
	OracleRegisterTx txInfo
	OracleQueryTx    txInfo
	OracleRespondTx  txInfo
	OracleExtendTx   txInfo
	ContractCreateTx txInfo
	ContractCallTx   txInfo
}

var sentTxs txTypes
var useTestNet bool

func signBroadcastWaitForTransaction(t *testing.T, tx aeternity.Tx, acc *aeternity.Account, node *aeternity.Node) (height uint64, txHash string, mbHash string) {
	txHash = signBroadcast(t, tx, acc, node)
	height, mbHash, err := waitForTransaction(node, txHash)
	if err != nil {
		t.Fatal(err)
	}
	info := txInfo{
		height: height,
		txHash: txHash,
		mbHash: mbHash,
	}
	rts := strings.Split(reflect.TypeOf(tx).String(), ".") // ['*aeternity', 'SpendTx']
	switch txType := rts[1]; txType {
	case "SpendTx":
		sentTxs.SpendTx = info
	case "NamePreclaimTx":
		sentTxs.NamePreclaimTx = info
	case "NameClaimTx":
		sentTxs.NameClaimTx = info
	case "NameUpdateTx":
		sentTxs.NameUpdateTx = info
	case "NameTransferTx":
		sentTxs.NameTransferTx = info
	case "NameRevokeTx":
		sentTxs.NameRevokeTx = info
	case "OracleRegisterTx":
		sentTxs.OracleRegisterTx = info
	case "OracleQueryTx":
		sentTxs.OracleQueryTx = info
	case "OracleRespondTx":
		sentTxs.OracleRespondTx = info
	case "OracleExtendTx":
		sentTxs.OracleExtendTx = info
	case "ContractCreateTx":
		sentTxs.ContractCreateTx = info
	case "ContractCallTx":
		sentTxs.ContractCallTx = info
	default:
		t.Fatalf("Where should I put this TxTpye: %s", txType)
	}

	return height, txHash, mbHash
}

func init() {
	flag.BoolVar(&useTestNet, "testnet", false, "Run tests that need an internet connection to testnet")
	flag.Parse()
}

func TestAPI(t *testing.T) {
	privateNet := setupNetwork(t, privatenetURL, false)
	testNet := setupNetwork(t, testnetURL, false)

	alice, bob := setupAccounts(t)

	name := randomName(6)
	helpers := aeternity.Helpers{Node: privateNet}
	ctxAlice := aeternity.NewContext(alice.Address, helpers)
	ctxBob := aeternity.NewContext(bob.Address, helpers)
	// SpendTx
	fmt.Println("SpendTx")
	spendTx, err := ctxAlice.SpendTx(sender, bob.Address, *big.NewInt(1000), aeternity.Config.Client.Fee, []byte(""))
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &spendTx, alice, privateNet)

	// NamePreClaimTx
	fmt.Println("NamePreClaimTx")
	preclaimTx, salt, err := ctxAlice.NamePreclaimTx(name, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &preclaimTx, alice, privateNet)

	// NameClaimTx
	fmt.Println("NameClaimTx")
	claimTx, err := ctxAlice.NameClaimTx(name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &claimTx, alice, privateNet)

	// NameUpdateTx
	fmt.Println("NameUpdateTx")
	updateTx, err := ctxAlice.NameUpdateTx(name, alice.Address)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &updateTx, alice, privateNet)

	// NameTransferTx
	fmt.Println("NameTransferTx")
	transferTx, err := ctxAlice.NameTransferTx(name, bob.Address)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &transferTx, alice, privateNet)

	// NameRevokeTx
	fmt.Println("NameRevokeTx")
	revokeTx, err := ctxBob.NameRevokeTx(name)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &revokeTx, bob, privateNet)

	sentTxs.name = randomName(8)
	// NamePreClaimTx
	fmt.Println("NamePreClaimTx 2nd name for other tests")
	preclaimTx, salt, err = ctxAlice.NamePreclaimTx(sentTxs.name, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &preclaimTx, alice, privateNet)

	// NameClaimTx
	fmt.Println("NameClaimTx 2nd name for other tests")
	claimTx, err = ctxAlice.NameClaimTx(sentTxs.name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &claimTx, alice, privateNet)

	// OracleRegisterTx
	fmt.Println("OracleRegisterTx")
	register, err := ctxAlice.OracleRegisterTx("hello", "helloback", *big.NewInt(1000), uint64(0), uint64(100), 0)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &register, alice, privateNet)

	// OracleExtendTx
	fmt.Println("OracleExtendTx")
	sentTxs.oracleID = strings.Replace(alice.Address, "ak_", "ok_", 1)
	extend, err := ctxAlice.OracleExtendTx(sentTxs.oracleID, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &extend, alice, privateNet)

	// OracleQueryTx
	fmt.Println("OracleQueryTx")
	query, err := ctxAlice.OracleQueryTx(sentTxs.oracleID, "How was your day?", *big.NewInt(1000), 0, 100, 0, 100)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitForTransaction(t, &query, alice, privateNet)

	// OracleRespondTx
	fmt.Println("OracleRespondTx")
	var oracleQueries *models.OracleQueries
	getOracleQueries := func() {
		oracleQueries, err = privateNet.GetOracleQueriesByPubkey(sentTxs.oracleID)
		if err != nil {
			t.Fatalf("APIGetOracleQueriesByPubkey: %s", err)
		}
	}
	delay(getOracleQueries)
	oqID := oracleQueries.OracleQueries[0].ID
	respond, err := ctxAlice.OracleRespondTx(sentTxs.oracleID, *oqID, "My day was fine thank you", 0, 100)
	_, _, _ = signBroadcastWaitForTransaction(t, &respond, alice, privateNet)

	t.Logf("%+v\n", sentTxs)
	t.Run("GetStatus", func(t *testing.T) {
		gotStatus, err := privateNet.GetStatus()
		// t.Logf("%+v\n", gotStatus)
		if *gotStatus.NetworkID != "ae_docker" {
			t.Errorf("Client.GetStatus(): Client testsuite should be run on private testnet (ae_docker), not %s", *gotStatus.NetworkID)
		}
		if err != nil {
			t.Errorf("Client.GetStatus() error = %v", err)
			return
		}
	})
	t.Run("GetTopBlock", func(t *testing.T) {
		_, err := privateNet.GetTopBlock()
		// t.Logf("%+v\n", gotTopBlock)
		if err != nil {
			t.Errorf("Client.GetTopBlock() error = %v", err)
			return
		}
	})
	t.Run("GetHeight", func(t *testing.T) {
		gotHeight, err := privateNet.GetHeight()
		// t.Logf("gotHeight: %d", gotHeight)
		if err != nil {
			t.Errorf("Client.GetHeight() error = %d", err)
			return
		}
		if gotHeight < 1 {
			t.Errorf("Client.GetHeight() returned an invalid height: %v", gotHeight)
		}
	})
	t.Run("GetCurrentKeyBlock", func(t *testing.T) {
		_, err := privateNet.GetCurrentKeyBlock()
		// t.Logf("%+v\n", gotCurrentKeyBlock)
		if err != nil {
			t.Errorf("Client.GetCurrentKeyBlock() error = %v", err)
			return
		}
	})

	t.Run("GetAccount", func(t *testing.T) {
		var account = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
		_, err := privateNet.GetAccount(account)
		// t.Logf("%+v\n", gotAccount)
		if err != nil {
			t.Errorf("Client.GetAccount() error = %v", err)
			return
		}
	})

	t.Run("GetGenerationByHeight", func(t *testing.T) {
		height, err := privateNet.GetHeight()
		if err != nil {
			t.Error(err)
		}
		_, err = privateNet.GetGenerationByHeight(height)
		// t.Logf("%+v\n", gotGenerationByHeight)
		if err != nil {
			t.Errorf("Client.GetGenerationByHeight() error = %v", err)
			return
		}
	})

	t.Run("GetMicroBlockHeaderByHash", func(t *testing.T) {
		_, err := privateNet.GetMicroBlockHeaderByHash(sentTxs.SpendTx.mbHash)
		// t.Logf("%+v\n", gotMicroBlockHeaderByHash)
		if err != nil {
			t.Errorf("Client.GetMicroBlockHeaderByHash() error = %v", err)
			return
		}
	})

	t.Run("GetKeyBlockByHash", func(t *testing.T) {
		if !useTestNet {
			t.Skip("-testnet not specified: skipping test")
		}
		_, err := testNet.GetKeyBlockByHash("kh_2ZPK9GGvXKJ8vfwapBLztd2F8DSr9QdphZRHSdJH8MR298Guao")
		// t.Logf("%+v\n", gotKeyBlockByHash)
		if err != nil {
			t.Errorf("Client.GetKeyBlockByHash() error = %v", err)
			return
		}
	})
	t.Run("GetMicroBlockTransactionsByHash", func(t *testing.T) {
		_, err := privateNet.GetMicroBlockTransactionsByHash(sentTxs.SpendTx.mbHash)
		// t.Logf("%+v\n", gotMicroBlockTransactionsByHash)
		if err != nil {
			t.Errorf("Client.GetMicroBlockTransactionsByHash() error = %v", err)
			return
		}
	})
	t.Run("GetTransactionByHash", func(t *testing.T) {
		type args struct {
			txHash string
		}
		tests := []struct {
			name    string
			args    args
			wantTx  *models.GenericSignedTx
			wantErr bool
		}{
			{
				name: "SpendTx",
				args: args{
					txHash: sentTxs.SpendTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "NamePreclaimTx",
				args: args{
					txHash: sentTxs.NamePreclaimTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "NameClaimTx",
				args: args{
					txHash: sentTxs.NameClaimTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "NameUpdateTx",
				args: args{
					txHash: sentTxs.NameUpdateTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "NameTransferTx",
				args: args{
					txHash: sentTxs.NameTransferTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "NameRevokeTx",
				args: args{
					txHash: sentTxs.NameRevokeTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "OracleRegisterTx",
				args: args{
					txHash: sentTxs.OracleRegisterTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "OracleExtendTx",
				args: args{
					txHash: sentTxs.OracleExtendTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "OracleQueryTx",
				args: args{
					txHash: sentTxs.OracleQueryTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "OracleResponseTx",
				args: args{
					txHash: sentTxs.OracleRespondTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			// {
			// 	name: "ContractCreateTx",
			// 	args: args{
			// 		txHash: sentTxs.???.txHash,
			// 	},
			// 	wantTx:  &models.GenericSignedTx{},
			// 	wantErr: false,
			// },
			// {
			// 	name: "ContractCallTx",
			// 	args: args{
			// 		txHash: sentTxs.???.txHash,
			// 	},
			// 	wantTx:  &models.GenericSignedTx{},
			// 	wantErr: false,
			// },
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				genericTx, err := privateNet.GetTransactionByHash(tt.args.txHash)
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.GetTransactionByHash() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				concreteTx := genericTx.Tx()
				if concreteTx.Type() != tt.name {
					t.Errorf("Expected Tx type %s, but received a %s instead", tt.name, concreteTx.Type())
					return
				}
			})
		}
	})
	t.Run("GetName", func(t *testing.T) {
		_, err := privateNet.GetNameEntryByName(sentTxs.name)
		// t.Logf("%+v\n", gotName)
		if err != nil {
			t.Errorf("Client.GetName() error = %v", err)
			return
		}
	})
	t.Run("GetOracleByPubkey", func(t *testing.T) {
		_, err := privateNet.GetOracleByPubkey(sentTxs.oracleID)
		// t.Logf("%+v\n", gotOracleByPubkey)
		if err != nil {
			t.Errorf("Client.GetOracleByPubkey() error = %v", err)
			return
		}
	})
	t.Run("GetOracleQueriesByPubkey", func(t *testing.T) {
		_, err := privateNet.GetOracleQueriesByPubkey(sentTxs.oracleID)
		// t.Logf("%+v\n", gotOracleQueriesByPubkey)
		if err != nil {
			t.Errorf("Client.GetOracleQueriesByPubkey() error = %v", err)
			return
		}
	})
	// t.Run("GetContractByID", func(t *testing.T) {
	// 	_, err := privateNet.GetContractByID(sentTxs.contractID)
	// 	// t.Logf("%+v\n", gotContractByID)
	// 	if err != nil {
	// 		t.Errorf("Client.GetContractByID() error = %v", err)
	// 		return
	// 	}
	// })
}
