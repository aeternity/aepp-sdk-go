package cmd

import (
	"fmt"
	"os"
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
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sender    string
			recipient string
			amount    int64 // TODO potential problem with int64 for amount
		)

		// Load variables from arguments
		sender = args[0]
		recipient = args[1]
		amount, _ = strconv.ParseInt(args[2], 10, 64)

		// Validate arguments
		if len(sender) == 0 {
			fmt.Println("Error, missing or invalid sender address")
			os.Exit(1)
		}
		if len(recipient) == 0 {
			fmt.Println("Error, missing or invalid recipient address")
			os.Exit(1)
		}
		if amount <= 0 {
			fmt.Println("Error, missing or invalid amount")
			os.Exit(1)
		}
		if fee <= 0 {
			fmt.Println("Error, missing or invalid fee")
			os.Exit(1)
		}

		base64Tx, ttl, nonce, err := aeternity.SpendTransaction(sender, recipient, amount, fee, ``)
		if err != nil {
			fmt.Printf("Creating a Spend Transaction failed with %s", err)
			os.Exit(1)
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
	},
}

// txVerifyCmd implements the tx verify subocmmand.
// It verfies the signature of a signed transaction
var txVerifyCmd = &cobra.Command{
	Use:   "verify SENDER_ADDRESS SIGNED_TRANSACTION",
	Short: "Verify the signature of a signed base64 transaction",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Load variables from arguments
		sender := args[0]
		txSignedBase64 := args[1]

		if len(sender) == 0 {
			fmt.Println("Error, missing or invalid sender address")
			os.Exit(1)
		}
		if len(txSignedBase64) == 0 || txSignedBase64[0:3] != "tx_" {
			fmt.Println("Error, missing or invalid base64 encoded transaction")
			os.Exit(1)
		}
		valid, err := aeternity.VerifySignedTx(sender, txSignedBase64)
		if err != nil {
			fmt.Printf("Error while verifying signature: %s\n", err)
		}
		fmt.Printf("The signature is %t\n", valid)
	},
}

func init() {
	RootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txSpendCmd)
	txCmd.AddCommand(txVerifyCmd)

	// tx spend command
	// TODO Config is not initialized within cmd. This means default config vars have to be hardcoded into help messages
	txSpendCmd.Flags().Int64Var(&fee, "fee", 20000, fmt.Sprintf("Set the transaction fee (default=%d)", 20000))
}
