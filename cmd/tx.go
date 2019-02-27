package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aeternity/aepp-sdk-go/aeternity"

	"github.com/spf13/cobra"
)

// txCmd implments the tx command
var txCmd = &cobra.Command{
	Use:   "tx SUBCOMMAND [ARGS]...",
	Short: "Handle transactions creation",
	Long:  ``,
}

// txSpendCmd implements the tx spend subcommand.
// It returns an unsigned spend transaction (to be signed with account sign)
var txSpendCmd = &cobra.Command{
	Use:   "spend SENDER_ADDRESS RECIPIENT_ADDRESS AMOUNT",
	Short: "Create a transaction to another account (unsigned)",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	RunE:  txSpendFunc,
}

func txSpendFunc(cmd *cobra.Command, args []string) (err error) {
	var (
		sender    string
		recipient string
		amount    uint64 // TODO potential problem with uint64 not big.Int
	)

	// Load variables from arguments
	sender = args[0]
	recipient = args[1]
	amount, _ = strconv.ParseUint(args[2], 10, 64)

	// Validate arguments
	if len(sender) == 0 {
		return errors.New("Error, missing or invalid sender address")
	}
	if len(recipient) == 0 {
		return errors.New("Error, missing or invalid recipient address")
	}
	if amount <= 0 {
		return errors.New("Error, missing or invalid amount")
	}
	if fee <= 0 {
		return errors.New("Error, missing or invalid fee")
	}

	base64Tx, ttl, nonce, err := aeternity.SpendTransaction(sender, recipient, amount, fee, ``)
	if err != nil {
		return err
	}

	// Sender, Recipient, Amount, Ttl, Fee, Nonce, Payload, Encoded
	aeternity.Pp(
		"Sender acount", sender,
		"Recipient account", recipient,
		"Amount", amount,
		"TTL", ttl,
		"Fee", fee,
		"Nonce", nonce,
		"Payload", "not implemented",
		"Encoded", base64Tx,
	)
	return nil
}

// txVerifyCmd implements the tx verify subcommand.
// It verfies the signature of a signed transaction
var txVerifyCmd = &cobra.Command{
	Use:   "verify SENDER_ADDRESS SIGNED_TRANSACTION",
	Short: "Verify the signature of a signed base64 transaction",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  txVerifyFunc,
}

func txVerifyFunc(cmd *cobra.Command, args []string) (err error) {
	// Load variables from arguments
	sender := args[0]
	txSignedBase64 := args[1]

	if len(sender) == 0 {
		return errors.New("Error, missing or invalid sender address")
	}
	if len(txSignedBase64) == 0 || txSignedBase64[0:3] != "tx_" {
		return errors.New("Error, missing or invalid base64 encoded transaction")
	}
	valid, err := aeternity.VerifySignedTx(sender, txSignedBase64)
	if err != nil {
		err := fmt.Errorf("error while verifying signature: %s", err)
		return err
	}
	fmt.Printf("The signature is %t\n", valid)
	return nil
}

func init() {
	RootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txSpendCmd)
	txCmd.AddCommand(txVerifyCmd)

	// tx spend command
	txSpendCmd.Flags().Uint64Var(&fee, "fee", aeternity.Config.Client.Fee, fmt.Sprintf("Set the transaction fee (default=%d)", aeternity.Config.Client.Fee))
	txSpendCmd.Flags().Uint64Var(&ttl, "ttl", aeternity.Config.Client.TTL, fmt.Sprintf("Set the TTL in keyblocks (default=%d)", aeternity.Config.Client.TTL))
	txSpendCmd.Flags().Uint64Var(&nonce, "nonce", 0, fmt.Sprint("Set the transaction nonce, if not it will be automatically generated"))
}
