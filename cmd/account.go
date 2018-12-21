// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/spf13/cobra"
)

var (
	waitForTx       bool
	payload         string
	printPrivateKey bool
	accountFileName string
	password        string
	fee             string
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account PRIVATE_KEY_PATH",
	Short: "Interact with a account",
	Long:  ``,
}

// addressCmd represents the address subcommand
var addressCmd = &cobra.Command{
	Use:   "address ACCOUNT_KEYSTORE",
	Short: "Print the aeternity account address",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ask for th keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		aeternity.Pp("Account address", account.Address)
		if printPrivateKey {
			aeternity.Pp("Account private key", account.SigningKeyToHexString())
		}
	},
}

// createCmd represents the generate subcommand
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		account, _ := aeternity.NewAccount()
		// ask for password
		p, err := utils.AskPassword("Enter a password for your keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			return
		}
		// check if a name was given
		f, err := aeternity.StoreAccountToKeyStoreFile(account, p, accountFileName)
		if err != nil {
			fmt.Println("Error saving the keystore file: ", err)
			return
		}
		aeternity.Pp(
			"Wallet path", f,
			"Account address", account.Address,
		)
	},
}

// balanceCmd represents the generate subcommand
var balanceCmd = &cobra.Command{
	Use:   "balance ACCOUNT_KEYSTORE",
	Short: "Get the balance of an account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ask for th keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		a, err := aeCli.APIGetAccount(account.Address)
		if err != nil {
			fmt.Println("Error retrieving the account: ", err)
			os.Exit(1)
		}
		aeternity.PrintObject("account", a)
	},
}

// listCmd represents the generate subcommand
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the account in the default keys path",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(aeternity.ListWallets())
	},
}

var signCmd = &cobra.Command{
	Use:   "sign ACCOUNT_KEYSTORE UNSIGNED_TRANSACTION",
	Short: "Sign the input (e.g. a transaction)",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// ask for the keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}

		txUnsignedBase64 := args[1]
		txSignedBase64, txHash, signature, err := aeternity.SignEncodeTxStr(account, txUnsignedBase64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		aeternity.Pp(
			"Signing account address", account.Address,
			"Signature", signature,
			"Unsigned", txUnsignedBase64,
			"Signed", txSignedBase64,
			"Hash", txHash,
		)

	},
}

// saveCmd represents the generate subcommand
var saveCmd = &cobra.Command{
	Use:   "save ACCOUNT_HEX_STRING",
	Short: "Save an account from a hex string to a keystore file",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		account, err := aeternity.AccountFromHexString(args[0])
		if err != nil {
			fmt.Println("Error parsing the private key hex string:", err)
			return
		}
		p, err := utils.AskPassword("Enter a password for your keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			return
		}

		f, err := aeternity.StoreAccountToKeyStoreFile(account, p, accountFileName)
		if err != nil {
			fmt.Println("Error saving the keystore file: ", err)
			return
		}
		aeternity.Pp("Keystore path ", f)
	},
}

// spendCmd represents the spend subcommand
var spendCmd = &cobra.Command{
	Use:   "spend ACCOUNT_KEYSTORE RECIPIENT_ADDRESS AMOUNT",
	Short: "Print the aeternity account spend",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			keystorePath string
			recipient    string
			amount       int64
		)

		// load variables
		for _, a := range args {
			if strings.HasPrefix(a, string(aeternity.PrefixAccountPubkey)) {
				recipient = a
				continue
			}
			if m, _ := regexp.MatchString(`^\d+$`, a); m {
				amount, _ = strconv.ParseInt(a, 10, 64)
			}
			if p, err := aeternity.GetWalletPath(a); err == nil {
				keystorePath = p
			}
		}

		// validate variables
		if len(recipient) == 0 {
			fmt.Println("Error, missing or invalid recipient address")
			os.Exit(1)
		}
		if len(keystorePath) == 0 {
			fmt.Println("Error, missing or invalid keystore path")
			os.Exit(1)
		}
		if amount <= 0 {
			fmt.Println("Error, missing or invalid amount")
			os.Exit(1)
		}
		// ask for the keystore password if not already set by CLI flags
		if len(password) == 0 {
			var err error
			password, err = utils.AskPassword("Enter the password to unlock the keystore: ")
			if err != nil {
				fmt.Println("Error reading the password: ", err)
				os.Exit(1)
			}
		}

		// load the account
		account, err := aeternity.LoadAccountFromKeyStoreFile(keystorePath, password)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		// run the transaction
		_, txHash, _, ttl, _, err := aeCli.WithAccount(account).Spend(recipient, amount, payload)

		// TODO: print also the ttl
		aeternity.Pp(
			"Sender Address", account.Address,
			"Recipient Address", recipient,
			"Amount", amount,
			"TransactionHash", txHash,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if waitForTx {
			_, _, _, tx, err := aeCli.WaitForTransactionUntillHeight(ttl, txHash)
			aeternity.PrintObject("Transaction", tx)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

var txCmd = &cobra.Command{
	Use:   "tx SUBCOMMAND [ARGS]...",
	Short: "Handle transactions creation",
	Long:  ``,
}

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
		var feeInt int64
		if len(fee) == 0 {
			feeInt = aeternity.Config.P.Client.Fee
		} else {
			feeInt, _ = strconv.ParseInt(fee, 10, 64)
		}

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
		if feeInt <= 0 {
			fmt.Println(feeInt)
			fmt.Println("Error, missing or invalid fee")
			os.Exit(1)
		}

		base64Tx, ttl, nonce, err := aeternity.SpendTransaction(sender, recipient, amount, feeInt, ``)
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
			"Fee", feeInt,
			"Nonce", nonce,
			"Payload", "not implemented",
			"Encoded", base64Tx,
		)
	},
}

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
			fmt.Println("Error, missing or invalid recipient address")
			os.Exit(1)
		}
		valid, err := aeternity.VerifySignedTx(sender, txSignedBase64)
		if err != nil {
			fmt.Printf("Error while verifying signature: %s\n", err)
		}
		fmt.Printf("The signature is %t\n", valid)
	},
}

var txBroadcastCmd = &cobra.Command{
	Use:   "broadcast SIGNED_TRANSACTION",
	Short: "Broadcast a transaction to the network",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Load variables from arguments
		txSignedBase64 := args[0]

		if len(txSignedBase64) == 0 || txSignedBase64[0:3] != "tx_" {
			fmt.Println("Error, missing or invalid recipient address")
			os.Exit(1)
		}

		err := aeternity.BroadcastTransaction(txSignedBase64)
		if err != nil {
			fmt.Println("Error while broadcasting transaction: ", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(accountCmd)
	RootCmd.AddCommand(txCmd)
	accountCmd.AddCommand(addressCmd)
	accountCmd.AddCommand(spendCmd)
	accountCmd.AddCommand(createCmd)
	accountCmd.AddCommand(saveCmd)
	accountCmd.AddCommand(balanceCmd)
	accountCmd.AddCommand(listCmd)
	accountCmd.AddCommand(signCmd)
	txCmd.AddCommand(txSpendCmd)
	txCmd.AddCommand(txVerifyCmd)
	txCmd.AddCommand(txBroadcastCmd)

	// create flags
	createCmd.Flags().StringVar(&accountFileName, "name", "", "Override the default name of a wallaet")
	// save flags
	saveCmd.Flags().StringVar(&accountFileName, "name", "", "Override the default name of a wallaet")
	// address flags
	addressCmd.Flags().BoolVar(&printPrivateKey, "private-key", false, "Print the private key as hex string")
	// spend command flags
	spendCmd.Flags().BoolVarP(&waitForTx, "wait", "w", false, "Wait for the transaction to be mined before exiting")
	spendCmd.Flags().StringVarP(&payload, "message", "m", "", "Payload to add to the spend transaction")
	spendCmd.Flags().StringVar(&password, "password", "", "Read account password from stdin [WARN: this method is not secure]")
	// tx spend command
	txSpendCmd.Flags().StringVar(&fee, "fee", "", "Set the transaction fee (default=1)")
}
