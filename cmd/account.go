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
	fee             int64
)

// accountCmd implements the account command
var accountCmd = &cobra.Command{
	Use:   "account PRIVATE_KEY_PATH",
	Short: "Interact with a account",
	Long:  ``,
}

// addressCmd implements the account address subcommand
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

// createCmd implements the account generate subcommand
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

// balanceCmd implements the account balance subcommand
var balanceCmd = &cobra.Command{
	Use:   "balance ACCOUNT_KEYSTORE",
	Short: "Get the balance of an account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		aeCli := NewAeCli()
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

// signCmd implements the account sign subcommand
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

// saveCmd implements the account save subcommand
var saveCmd = &cobra.Command{
	Use:   "save ACCOUNT_HEX_STRING",
	Short: "Save an account from a hex string to a keystore file",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		accountFileName := args[0]
		account, err := aeternity.AccountFromHexString(args[1])
		if err != nil {
			fmt.Println("Error parsing the private key hex string:", err)
			return
		}

		if len(password) == 0 {
			var err error
			password, err = utils.AskPassword("Enter a password for your keystore: ")
			if err != nil {
				fmt.Println("Error reading the password: ", err)
				os.Exit(1)
			}
		}

		f, err := aeternity.StoreAccountToKeyStoreFile(account, password, accountFileName)
		if err != nil {
			fmt.Println("Error saving the keystore file: ", err)
			return
		}
		aeternity.Pp("Keystore path ", f)
	},
}

func init() {
	RootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(addressCmd)
	accountCmd.AddCommand(createCmd)
	accountCmd.AddCommand(saveCmd)
	accountCmd.AddCommand(balanceCmd)
	accountCmd.AddCommand(signCmd)

	// account sign flags
	signCmd.Flags().StringVar(&password, "password", "", "Read account password from stdin [WARN: this method is not secure]")
	// account create flags
	createCmd.Flags().StringVar(&accountFileName, "name", "", "Override the default name of a wallet")
	// account save flags
	saveCmd.Flags().StringVar(&password, "password", "", "Read account password from stdin [WARN: this method is not secure]")
	// account address flags
	addressCmd.Flags().BoolVar(&printPrivateKey, "private-key", false, "Print the private key as hex string")
}
