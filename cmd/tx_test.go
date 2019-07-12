package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/spf13/cobra"
)

var alice = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var bob = "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

func Test_txSpendFunc(t *testing.T) {
	type args struct {
		conn aeternity.GetAccounter
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Alice sends 10 to Bob",
			args: args{
				conn: &mockGetAccounter{account: `{"balance":1600000000000000077131306000000000000000,"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","kind":"basic","nonce":0}`},
				args: []string{alice, bob, "10"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := txSpendFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("txSpendFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxVerify(t *testing.T) {
	// unsigned tx_+FMMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoa15iD0gACCAfQBgIHqJ/Y=
	// sign with ae_mainnet
	signedTx := "tx_+J0LAfhCuEBcvwtyCo3FYqmINcP6lHLH/dRDcj5rUiKDqYKhPpiQ+1SBQ66rF3gdVQ1IcANcw/IayK//YgK2dsDF1VtroQEAuFX4UwwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKhrXmIPSAAIIB9AGAx+EjLg=="
	aeternity.Config.Node.NetworkID = "ae_mainnet"
	emptyCmd := cobra.Command{}

	err := txVerifyFunc(&emptyCmd, []string{alice, signedTx})
	if err != nil {
		t.Error(err)
	}
}

func TestTxDumpRaw(t *testing.T) {
	tx := "tx_+H4iAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKADy8Y5hY2NvdW50X3B1YmtleaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMGAQXLBNnv"
	emptyCmd := cobra.Command{}

	err := txDumpRawFunc(&emptyCmd, []string{tx})
	if err != nil {
		t.Error(err)
	}
}

type mockgetHeightAccounter struct {
	height  uint64
	account string
}

func (m *mockgetHeightAccounter) GetHeight() (uint64, error) {
	return m.height, nil
}
func (m *mockgetHeightAccounter) GetAccount(accountID string) (acc *models.Account, err error) {
	acc = &models.Account{}
	err = acc.UnmarshalBinary([]byte(m.account))
	return acc, err
}
func Test_txContractCreateFunc(t *testing.T) {
	aeternity.GetTTLNonce = func(c aeternity.GetHeightAccounter, accountID string, offset uint64) (height uint64, nonce uint64, err error) {
		return 2, 1337, nil
	}
	type args struct {
		conn getHeightAccounter
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Deploy SimpleStorage with alice (unsigned)",
			args: args{
				conn: &mockgetHeightAccounter{},
				args: []string{alice, contractSimpleStorageBytecode, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := txContractCreateFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("txContractCreateFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
