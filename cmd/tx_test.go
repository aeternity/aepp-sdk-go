package cmd

import (
	"testing"

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
	signedTx := "tx_+JkLAfhCuEDXP1xmmSfneNLH6Py6bFVQO8p9Rmc60fmGABDnGAyeZTSaPx0OX/EKjcOwriZ8UA+x4rAIdTBq8qtZ7iUJiRoDuFH4TwwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKgk4ggoLnAYBsrWq9"
	emptyCmd := cobra.Command{}

	err := txVerifyFunc(&emptyCmd, []string{sender, signedTx})
	if err != nil {
		t.Error(err)
	}
}
