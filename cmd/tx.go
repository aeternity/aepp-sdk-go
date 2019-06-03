package cmd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/utils"

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
		amount    *big.Int
		feeBigInt *big.Int
	)

	// Load variables from arguments
	sender = args[0]
	recipient = args[1]
	amount, err = utils.NewIntFromString(args[2])
	feeBigInt, _ = utils.NewIntFromString(fee)

	// Validate arguments
	if len(sender) == 0 {
		return errors.New("Error, missing or invalid sender address")
	}
	if len(recipient) == 0 {
		return errors.New("Error, missing or invalid recipient address")
	}
	if amount.Cmp(big.NewInt(0)) == -1 {
		return errors.New("Error, missing or invalid amount")
	}
	if feeBigInt.Cmp(big.NewInt(0)) == -1 {
		return errors.New("Error, missing or invalid fee")
	}

	// Connect to the node to find out sender nonce only
	if nonce == 0 {
		client := aeternity.NewClient(aeternity.Config.Node.URL, false)
		nonce, err = client.GetNextNonce(sender)
		if err != nil {
			return err
		}
	}

	tx := aeternity.NewSpendTx(sender, recipient, *amount, *feeBigInt, spendTxPayload, ttl, nonce)
	if err != nil {
		return err
	}
	base64Tx, err := aeternity.BaseEncodeTx(&tx)
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
		"Payload", spendTxPayload,
		"Encoded", base64Tx,
	)
	return nil
}

// txVerifyCmd implements the tx verify subcommand.
// It verfies the signature of a signed transaction
var txVerifyCmd = &cobra.Command{
	Use:          "verify SENDER_ADDRESS SIGNED_TRANSACTION",
	Short:        "Verify the signature of a signed base64 transaction",
	Long:         ``,
	Args:         cobra.ExactArgs(2),
	RunE:         txVerifyFunc,
	SilenceUsage: true,
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
	valid, err := aeternity.VerifySignedTx(sender, txSignedBase64, aeternity.Config.Node.NetworkID)
	if err != nil {
		err := fmt.Errorf("error while verifying signature: %s", err)
		return err
	}
	if valid {
		fmt.Printf("The signature is valid (network-id: %s)\n", aeternity.Config.Node.NetworkID)
	} else {
		message := fmt.Sprintf("The signature is invalid (expecting network-id: %s)", aeternity.Config.Node.NetworkID)
		// fmt.Println(message)
		err = errors.New(message)
	}
	return err
}

// txDumpRawCmd implements the tx dumpraw subcommand.
// It decodes a base58/64 input down into its RLP byte level representation.
var txDumpRawCmd = &cobra.Command{
	Use:   "dumpraw TRANSACTION",
	Short: "Show the RLP byte level representation of a base58/64 encoded object",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  txDumpRawFunc,
}

func txDumpRawFunc(cmd *cobra.Command, args []string) (err error) {
	tx := args[0]
	txRaw, err := aeternity.Decode(tx)
	if err != nil {
		return err
	}
	res := aeternity.DecodeRLPMessage(txRaw)
	fmt.Println(res)
	return nil
}

func init() {
	RootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txSpendCmd)
	txCmd.AddCommand(txVerifyCmd)
	txCmd.AddCommand(txDumpRawCmd)

	// tx spend command
	txSpendCmd.Flags().StringVar(&fee, "fee", aeternity.Config.Client.Fee.String(), fmt.Sprintf("Set the transaction fee (default=%s)", aeternity.Config.Client.Fee.String()))
	txSpendCmd.Flags().Uint64Var(&ttl, "ttl", aeternity.Config.Client.TTL, fmt.Sprintf("Set the TTL in keyblocks (default=%d)", aeternity.Config.Client.TTL))
	txSpendCmd.Flags().Uint64Var(&nonce, "nonce", 0, fmt.Sprint("Set the sender account nonce, if not the chain will be queried for its value"))
	txSpendCmd.Flags().StringVar(&spendTxPayload, "payload", "", fmt.Sprint("Optional text payload for Spend Transactions"))
}
