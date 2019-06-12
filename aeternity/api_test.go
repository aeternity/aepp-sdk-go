package aeternity

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/generated/models"
)

/*
For best coverage, this needs to be run against a local node, which needs to have
the transactions submitted to it first. Which makes it an integration test.

Mainnet/testnet doesn't have all transactions, and even if they did, oracles/names expire,
so you can't test all the endpoints anyway. Best to do this in a controlled environment.
*/
const url = "http://sdk-testnet.aepps.com"

func TestClient_GetStatus(t *testing.T) {
	c := NewClient(url, false)
	gotStatus, err := c.GetStatus()
	t.Logf("%#v\n", gotStatus)
	if *gotStatus.NetworkID != "ae_uat" {
		t.Errorf("Client.GetStatus(): Client testsuite should be run on testnet (ae_uat), not %s", *gotStatus.NetworkID)
	}
	if err != nil {
		t.Errorf("Client.GetStatus() error = %v", err)
		return
	}
}

func TestClient_GetTopBlock(t *testing.T) {
	c := NewClient(url, false)
	gotTopBlock, err := c.GetTopBlock()
	t.Logf("%#v\n", gotTopBlock)
	if err != nil {
		t.Errorf("Client.GetTopBlock() error = %v", err)
		return
	}
}

func TestClient_GetHeight(t *testing.T) {
	c := NewClient(url, false)
	gotHeight, err := c.GetHeight()
	t.Logf("gotHeight: %d", gotHeight)
	if err != nil {
		t.Errorf("Client.GetHeight() error = %d", err)
		return
	}
	if gotHeight < 1 {
		t.Errorf("Client.GetHeight() returned an invalid height: %v", gotHeight)
	}
}

func TestClient_GetCurrentKeyBlock(t *testing.T) {
	c := NewClient(url, false)
	gotCurrentKeyBlock, err := c.GetCurrentKeyBlock()
	t.Logf("%#v\n", gotCurrentKeyBlock)
	if err != nil {
		t.Errorf("Client.GetCurrentKeyBlock() error = %v", err)
		return
	}
}

func TestClient_GetAccount(t *testing.T) {
	c := NewClient(url, false)
	var account = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	gotAccount, err := c.GetAccount(account)
	t.Logf("%#v\n", gotAccount)
	if err != nil {
		t.Errorf("Client.GetAccount() error = %v", err)
		return
	}
}

func TestClient_GetGenerationByHeight(t *testing.T) {
	c := NewClient(url, false)
	gotGenerationByHeight, err := c.GetGenerationByHeight(94062)
	t.Logf("%#v\n", gotGenerationByHeight)
	if err != nil {
		t.Errorf("Client.GetGenerationByHeight() error = %v", err)
		return
	}
}

func TestClient_GetMicroBlockHeaderByHash(t *testing.T) {
	c := NewClient(url, false)
	gotMicroBlockHeaderByHash, err := c.GetMicroBlockHeaderByHash("mh_ksqiLRuTqJqtEH2iL43HU8Hg3A3ELUMdzzqZSfist8yCBADrq")
	t.Logf("%#v\n", gotMicroBlockHeaderByHash)
	if err != nil {
		t.Errorf("Client.GetMicroBlockHeaderByHash() error = %v", err)
		return
	}
}

func TestClient_GetKeyBlockByHash(t *testing.T) {
	c := NewClient(url, false)
	gotKeyBlockByHash, err := c.GetKeyBlockByHash("kh_2ZPK9GGvXKJ8vfwapBLztd2F8DSr9QdphZRHSdJH8MR298Guao")
	t.Logf("%#v\n", gotKeyBlockByHash)
	if err != nil {
		t.Errorf("Client.GetKeyBlockByHash() error = %v", err)
		return
	}
}

// TODO Oracles/Names/Contracts expire!
func TestClient_GetName(t *testing.T) {
	c := NewClient(url, false)
	name := "artcontract.test"
	gotName, err := c.GetNameEntryByName(name)
	t.Logf("%#v\n", gotName)
	if err != nil {
		t.Errorf("Client.GetName() error = %v", err)
		return
	}
}

// TODO this test fails/succeeds depending on the different types of txs in the microblock!
func TestClient_GetMicroBlockTransactionsByHash(t *testing.T) {
	c := NewClient(url, false)
	gotMicroBlockTransactionsByHash, err := c.GetMicroBlockTransactionsByHash("mh_ksqiLRuTqJqtEH2iL43HU8Hg3A3ELUMdzzqZSfist8yCBADrq")
	t.Logf("%#v\n", gotMicroBlockTransactionsByHash)
	if err != nil {
		t.Errorf("Client.GetMicroBlockTransactionsByHash() error = %v", err)
		return
	}
}

// TODO Oracles/Names/Contracts expire
func TestClient_GetOracleByPubkey(t *testing.T) {
	c := NewClient(url, false)
	oracle := "ok_something"
	gotOracleByPubkey, err := c.GetOracleByPubkey(oracle)
	t.Logf("%#v\n", gotOracleByPubkey)
	if err != nil {
		t.Errorf("Client.GetOracleByPubkey() error = %v", err)
		return
	}
}

// TODO Oracles/Names/Contracts expire
func TestClient_GetOracleQueriesByPubkey(t *testing.T) {
	c := NewClient(url, false)
	oracle := "ok_something"
	gotOracleQueriesByPubkey, err := c.GetOracleQueriesByPubkey(oracle)
	t.Logf("%#v\n", gotOracleQueriesByPubkey)
	if err != nil {
		t.Errorf("Client.GetOracleQueriesByPubkey() error = %v", err)
		return
	}
}

// TODO Oracles/Names/Contracts expire
func TestClient_GetContractByID(t *testing.T) {
	c := NewClient(url, false)
	contract := "ct_something"
	gotContractByID, err := c.GetContractByID(contract)
	t.Logf("%#v\n", gotContractByID)
	if err != nil {
		t.Errorf("Client.GetContractByID() error = %v", err)
		return
	}
}

// TODO find an OracleRespondTx
func TestClient_GetTransactionByHash(t *testing.T) {
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
				txHash: "th_2663oMtP945CjyWUhKgb7ak3jjZjTnbG3VuLK9k4s8ABqobK5Z",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "NamePreclaimTx",
			args: args{
				txHash: "th_Es36mD6NTQDbSHprmVqUui63uTM7TKwkcPGWcKSpf176XryzB",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "NameClaimTx",
			args: args{
				txHash: "th_b2jCJFyoN8Y4P51wwKRJZ3UERWFV8UGxuKHxiLJRSP1qv6GNF",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "NameUpdateTx",
			args: args{
				txHash: "th_2RptWR8B5JbziGebxa5vvr9xuBBCqtwndbrcNQjen56SS1LSHR",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "NameTransferTx",
			args: args{
				txHash: "th_XpwwJqW4S5oVLDRbgouPWo3nF1u8oon9KDmM944aKEJgr63az",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "NameRevokeTx",
			args: args{
				txHash: "th_G4s1Befn1JLws54ZTSAxVidEqJ4vqVPaowzeqELf7u4DPfHks",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "OracleRegisterTx",
			args: args{
				txHash: "th_2tBJBwA7Z866e5DKA3vckc7iojEXHvTNCenZWo4iaKagfRP4g9",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "OracleExtendTx",
			args: args{
				txHash: "th_aeD32ywzPKxEbkumExc7tNauVGN1zWNvt4EuiRxUPtBLDb5o3",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "OracleQueryTx",
			args: args{
				txHash: "th_uchsRuh2a3yGNqYP3JHw5VyqCpfqxvFpvPybD941ycDuru4QB",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		// {
		// 	name: "OracleRespondTx",
		// 	args: args{
		// 		txHash: "nobody has ever sent this on mainnet/testnet",
		// 	},
		// 	wantTx:  &models.GenericSignedTx{},
		// 	wantErr: false,
		// },
		{
			name: "ContractCreateTx",
			args: args{
				txHash: "th_21yVrEEyaoGuzZmwBAS97TKDvQ4BmUFMiK6YjYLRWcRbwNGb3T",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
		{
			name: "ContractCallTx",
			args: args{
				txHash: "th_4Tf5fGRqyTshwU5F6SziS9NwgU9gLToHT9PHfaZiYgSwZidjc",
			},
			wantTx:  &models.GenericSignedTx{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(url, false)
			genericTx, err := c.GetTransactionByHash(tt.args.txHash)
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
}
