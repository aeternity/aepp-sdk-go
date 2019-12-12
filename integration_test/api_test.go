package integrationtest

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"

	"gotest.tools/golden"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/aeternity"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
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

func signBroadcastWaitKeepTrackOfTx(t *testing.T, tx transactions.Transaction, acc *account.Account, node *naet.Node) (height uint64, txHash string, mbHash string) {
	_, txHash, _, height, mbHash, err := aeternity.SignBroadcastWaitTransaction(tx, acc, node, networkID, config.Client.WaitBlocks)
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
		t.Fatalf("Where should I put this TxType: %s", txType)
	}

	return height, txHash, mbHash
}

func TestMain(m *testing.M) {
	flag.BoolVar(&useTestNet, "testnet", false, "Run tests that need an internet connection to testnet")
	flag.Parse()
	os.Exit(m.Run())
}

func TestAPI(t *testing.T) {
	privateNet := setupNetwork(t, privatenetURL, false)
	testNet := setupNetwork(t, testnetURL, false)
	ttlnoncer := transactions.NewTTLNoncer(privateNet)

	alice, bob := setupAccounts(t)

	name := randomName(int(config.Client.Names.NameAuctionMaxLength + 1))

	// SpendTx
	fmt.Println("SpendTx")
	spendTx, err := transactions.NewSpendTx(alice.Address, bob.Address, big.NewInt(1000), []byte(""), ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, spendTx, alice, privateNet)

	// NamePreClaimTx
	fmt.Println("NamePreClaimTx")
	preclaimTx, salt, err := transactions.NewNamePreclaimTx(alice.Address, name, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, preclaimTx, alice, privateNet)

	// NameClaimTx
	fmt.Println("NameClaimTx")
	nameFee := transactions.CalculateMinNameFee(name)
	claimTx, err := transactions.NewNameClaimTx(alice.Address, name, salt, nameFee, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, claimTx, alice, privateNet)

	// NameUpdateTx
	fmt.Println("NameUpdateTx")
	updateTx, err := transactions.NewNameUpdateTx(alice.Address, name, []string{alice.Address}, config.Client.Names.ClientTTL, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, updateTx, alice, privateNet)

	// NameTransferTx
	fmt.Println("NameTransferTx")
	transferTx, err := transactions.NewNameTransferTx(alice.Address, name, bob.Address, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, transferTx, alice, privateNet)

	// NameRevokeTx
	fmt.Println("NameRevokeTx")
	revokeTx, err := transactions.NewNameRevokeTx(bob.Address, name, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, revokeTx, bob, privateNet)

	sentTxs.name = randomName(int(config.Client.Names.NameAuctionMaxLength + 2))
	// NamePreClaimTx
	fmt.Println("NamePreClaimTx 2nd name for other tests")
	preclaimTx, salt, err = transactions.NewNamePreclaimTx(alice.Address, sentTxs.name, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, preclaimTx, alice, privateNet)

	// NameClaimTx
	fmt.Println("NameClaimTx 2nd name for other tests")
	nameFee2 := transactions.CalculateMinNameFee(sentTxs.name)
	claimTx, err = transactions.NewNameClaimTx(alice.Address, sentTxs.name, salt, nameFee2, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, claimTx, alice, privateNet)

	// OracleRegisterTx
	fmt.Println("OracleRegisterTx")
	register, err := transactions.NewOracleRegisterTx(alice.Address, "hello", "helloback", config.Client.Oracles.QueryFee, config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, config.Client.Oracles.ABIVersion, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, register, alice, privateNet)

	// OracleExtendTx
	fmt.Println("OracleExtendTx")
	sentTxs.oracleID = register.ID()
	extend, err := transactions.NewOracleExtendTx(alice.Address, sentTxs.oracleID, config.OracleTTLTypeDelta, config.Client.Oracles.QueryTTLValue, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, extend, alice, privateNet)

	// OracleQueryTx
	fmt.Println("OracleQueryTx")
	query, err := transactions.NewOracleQueryTx(alice.Address, sentTxs.oracleID, "How was your day?", config.Client.Oracles.QueryFee, config.OracleTTLTypeDelta, config.Client.Oracles.QueryTTLValue, config.OracleTTLTypeDelta, config.Client.Oracles.ResponseTTLValue, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, query, alice, privateNet)

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
	respond, err := transactions.NewOracleRespondTx(alice.Address, sentTxs.oracleID, *oqID, "My day was fine thank you", config.OracleTTLTypeDelta, config.Client.Oracles.ResponseTTLValue, ttlnoncer)
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, respond, alice, privateNet)

	// ContractCreateTx
	fmt.Println("ContractCreateTx")
	ctCreateBytecode := string(golden.Get(t, "identity_bytecode.txt"))
	ctCreateInitCalldata := string(golden.Get(t, "identity_initcalldata.txt"))
	ctCreate, err := transactions.NewContractCreateTx(alice.Address, ctCreateBytecode, config.Client.Contracts.VMVersion, config.Client.Contracts.ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, ctCreateInitCalldata, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}

	sentTxs.contractID, err = ctCreate.ContractID()
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, ctCreate, alice, privateNet)

	// ContractCallTx
	fmt.Println("ContractCallTx")
	ctCallCalldata := string(golden.Get(t, "identity_main42.txt"))
	ctCall, err := transactions.NewContractCallTx(alice.Address, sentTxs.contractID, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, ctCallCalldata, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _ = signBroadcastWaitKeepTrackOfTx(t, ctCall, alice, privateNet)

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
				name: "OracleRespondTx",
				args: args{
					txHash: sentTxs.OracleRespondTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "ContractCreateTx",
				args: args{
					txHash: sentTxs.ContractCreateTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
			{
				name: "ContractCallTx",
				args: args{
					txHash: sentTxs.ContractCallTx.txHash,
				},
				wantTx:  &models.GenericSignedTx{},
				wantErr: false,
			},
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
	t.Run("GetContractByID", func(t *testing.T) {
		_, err := privateNet.GetContractByID(sentTxs.contractID)
		// t.Logf("%+v\n", gotContractByID)
		if err != nil {
			t.Errorf("Client.GetContractByID() error = %v", err)
			return
		}
	})
}
