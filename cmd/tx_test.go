package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/v5/aeternity"
	"github.com/spf13/cobra"
)

var alice = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var bob = "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"

func Test_txSpendFunc(t *testing.T) {
	type args struct {
		ttlFunc   aeternity.GetTTLFunc
		nonceFunc aeternity.GetNextNonceFunc
		args      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Alice sends 10 to Bob",
			args: args{
				ttlFunc:   func(offset uint64) (ttl uint64, err error) { return 500, nil },
				nonceFunc: func(address string) (nonce uint64, err error) { return 2, nil },
				args:      []string{alice, bob, "10"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := txSpendFunc(tt.args.ttlFunc, tt.args.nonceFunc, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("txSpendFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxVerify(t *testing.T) {
	// unsigned tx_+FMMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoa15iD0gACCAfQBgIHqJ/Y=
	// sign with ae_mainnet
	signedTx := "tx_+J0LAfhCuEBcvwtyCo3FYqmINcP6lHLH/dRDcj5rUiKDqYKhPpiQ+1SBQ66rF3gdVQ1IcANcw/IayK//YgK2dsDF1VtroQEAuFX4UwwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKhrXmIPSAAIIB9AGAx+EjLg=="
	config.Node.NetworkID = "ae_mainnet"
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

func Test_txContractCreateFunc(t *testing.T) {
	type args struct {
		ttlFunc   aeternity.GetTTLFunc
		nonceFunc aeternity.GetNextNonceFunc
		args      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Deploy SimpleStorage with alice (unsigned)",
			args: args{
				ttlFunc:   func(offset uint64) (ttl uint64, err error) { return 500, nil },
				nonceFunc: func(address string) (nonce uint64, err error) { return 2, nil },
				args:      []string{alice, contractSimpleStorageBytecode, contractSimpleStorageInit42},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := txContractCreateFunc(tt.args.ttlFunc, tt.args.nonceFunc, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("txContractCreateFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
