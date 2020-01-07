package cmd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v8/binary"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
	"github.com/aeternity/aepp-sdk-go/v8/utils"

	"github.com/spf13/cobra"
)

// txCmd implments the tx command. All tx subcommands should work offline,
// without any connection to the node.
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
	RunE: func(cmd *cobra.Command, args []string) error {
		node := newAeNode()
		ttlFunc := transactions.CreateTTLer(node)
		nonceFunc := transactions.CreateNoncer(node)
		return txSpendFunc(ttlFunc, nonceFunc, args)
	},
}

func txSpendFunc(ttlFunc transactions.TTLer, nonceFunc transactions.Noncer, args []string) (err error) {
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
	if !IsAddress(sender) {
		return errors.New("Error, missing or invalid sender address")
	}
	if !IsAddress(recipient) {
		return errors.New("Error, missing or invalid recipient address")
	}
	if amount.Cmp(big.NewInt(0)) == -1 {
		return errors.New("Error, missing or invalid amount")
	}
	if feeBigInt.Cmp(big.NewInt(0)) == -1 {
		return errors.New("Error, missing or invalid fee")
	}

	// If nonce or TTL was specified, no need to query the node
	if nonce > 0 {
		nonceFunc = func(accountID string) (uint64, error) {
			return nonce, nil
		}
	}
	if ttl > 0 {
		ttlFunc = func(offset uint64) (uint64, uint64, error) {
			return ttl, 0, nil
		}
	}

	tx, err := transactions.NewSpendTx(sender, recipient, amount, []byte(spendTxPayload), transactions.CreateTTLNoncer(ttlFunc, nonceFunc))
	if err != nil {
		return err
	}

	if feeBigInt.Cmp(big.NewInt(0)) != 0 {
		tx.SetFee(feeBigInt)
	}

	base64Tx, err := transactions.SerializeTx(tx)
	if err != nil {
		return err
	}

	// Print the result
	Pp(
		"Sender acount", tx.SenderID,
		"Recipient account", tx.RecipientID,
		"Amount", tx.Amount,
		"TTL", tx.TTL,
		"Fee", tx.Fee,
		"Nonce", tx.Nonce,
		"Payload", tx.Payload,
		"Encoded", base64Tx,
	)
	return nil
}

var txContractCreateCmd = &cobra.Command{
	Use:   "deploy OWNER_ID CONTRACT_BYTECODE INIT_CALLDATA",
	Short: "Create a smart contract on the blockchain",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		node := newAeNode()
		ttlFunc := transactions.CreateTTLer(node)
		nonceFunc := transactions.CreateNoncer(node)
		return txContractCreateFunc(ttlFunc, nonceFunc, args)
	},
}

type getHeightAccounter interface {
	naet.GetHeighter
	naet.GetAccounter
}

func txContractCreateFunc(ttlFunc transactions.TTLer, nonceFunc transactions.Noncer, args []string) (err error) {
	var (
		owner    string
		contract string
		calldata string
	)

	// Load variables from arguments and validate them
	owner = args[0]
	if !IsAddress(owner) {
		return errors.New("Error, missing or invalid owner address")
	}
	contract = args[1]
	if !IsBytecode(contract) {
		return errors.New("Error, missing or invalid contract bytecode")
	}
	calldata = args[2]
	if !IsBytecode(calldata) {
		return errors.New("Error, missing or invalid init calldata bytecode")
	}

	// If nonce was not specified as an argument, connect to the node to
	// query it
	if nonce > 0 {
		nonceFunc = func(accountID string) (uint64, error) {
			return nonce, nil
		}
	}
	// If TTL was not specified as an argument, connect to the node to calculate
	// it
	if ttl > 0 {
		ttlFunc = func(offset uint64) (uint64, uint64, error) {
			return ttl, 0, nil
		}
	}
	ttlnoncer := transactions.CreateTTLNoncer(ttlFunc, nonceFunc)

	tx, err := transactions.NewContractCreateTx(owner, contract, config.Client.Contracts.VMVersion, config.Client.Contracts.ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, calldata, ttlnoncer)
	if err != nil {
		return err
	}
	txStr, err := transactions.SerializeTx(tx)
	if err != nil {
		return err
	}

	// Print the result
	Pp(
		"OwnerID", tx.OwnerID,
		"AccountNonce", tx.AccountNonce,
		"Code", tx.Code,
		"VMVersion", tx.VMVersion,
		"ABIVersion", tx.AbiVersion,
		"Deposit", tx.Deposit,
		"Amount", tx.Amount,
		"GasLimit", tx.GasLimit,
		"GasPrice", tx.GasPrice,
		"TTL", tx.TTL,
		"Fee", tx.Fee,
		"CallData", tx.CallData,
		"Encoded", txStr,
	)

	return
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

	if !IsAddress(sender) {
		return errors.New("Error, missing or invalid sender address")
	}
	if !IsTransaction(txSignedBase64) {
		return errors.New("Error, missing or invalid base64 encoded transaction")
	}
	valid, err := transactions.VerifySignedTx(sender, txSignedBase64, config.Node.NetworkID)
	if err != nil {
		err := fmt.Errorf("error while verifying signature: %s", err)
		return err
	}
	if valid {
		fmt.Printf("The signature is valid (network-id: %s)\n", config.Node.NetworkID)
	} else {
		message := fmt.Sprintf("The signature is invalid (expecting network-id: %s)", config.Node.NetworkID)
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
	if !IsTransaction(tx) {
		return errors.New("Error, missing or invalid base64 encoded transaction")
	}
	txRaw, err := binary.Decode(tx)
	if err != nil {
		return err
	}
	res := binary.DecodeRLPMessage(txRaw)
	fmt.Println(res)
	return nil
}

func init() {
	RootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txSpendCmd)
	txCmd.AddCommand(txContractCreateCmd)
	txCmd.AddCommand(txVerifyCmd)
	txCmd.AddCommand(txDumpRawCmd)

	// tx spend command
	txSpendCmd.Flags().StringVar(&fee, "fee", config.Client.Fee.String(), fmt.Sprintf("Set the transaction fee (default=%s)", config.Client.Fee.String()))
	txSpendCmd.Flags().Uint64Var(&ttl, "ttl", 0, fmt.Sprintf("Set the TTL in keyblocks (default=%d)", 0))
	txSpendCmd.Flags().Uint64Var(&nonce, "nonce", 0, fmt.Sprint("Set the sender account nonce, if not the chain will be queried for its value"))
	txSpendCmd.Flags().StringVar(&spendTxPayload, "payload", "", fmt.Sprint("Optional text payload for Spend Transactions, which will be turned into a bytearray"))
}
