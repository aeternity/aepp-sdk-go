package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
)

func TestTxSpend(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	emptyCmd := cobra.Command{}

	err := txSpendFunc(&emptyCmd, []string{sender, recipient, "10"})
	if err != nil {
		t.Error(err)
	}
}

func TestTxVerify(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	// unsigned tx_+FMMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoa15iD0gACCAfQBgIHqJ/Y=
	// sign with ae_mainnet
	signedTx := "tx_+J0LAfhCuEBcvwtyCo3FYqmINcP6lHLH/dRDcj5rUiKDqYKhPpiQ+1SBQ66rF3gdVQ1IcANcw/IayK//YgK2dsDF1VtroQEAuFX4UwwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKhrXmIPSAAIIB9AGAx+EjLg=="
	aeternity.Config.Node.NetworkID = "ae_mainnet"
	emptyCmd := cobra.Command{}

	err := txVerifyFunc(&emptyCmd, []string{sender, signedTx})
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
