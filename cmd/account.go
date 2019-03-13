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
	fee             string // leave it as a string because viper cannot parse it directly into a BigInt
	ttl             uint64
	nonce           uint64
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
	RunE:  addressFunc,
}

func getPassword() (p string, err error) {
	if len(password) != 0 {
		return password, nil
	}
	p, err = utils.AskPassword("Enter the password to unlock the keystore: ")
	if err != nil {
		return "", err
	}
	return p, nil
}

func addressFunc(cmd *cobra.Command, args []string) error {
	p, err := getPassword()
	if err != nil {
		return err
	}

	// load the account
	account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	aeternity.Pp("Account address", account.Address)
	if printPrivateKey {
		aeternity.Pp("Account private key", account.SigningKeyToHexString())
	}

	return nil
}

// createCmd implements the account generate subcommand
var createCmd = &cobra.Command{
	Use:   "create ACCOUNT_KEYSTORE",
	Short: "Create a new account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  createFunc,
}

func createFunc(cmd *cobra.Command, args []string) (err error) {
	account, _ := aeternity.NewAccount()
	p, err := getPassword()
	if err != nil {
		return err
	}
	accountFileName = args[0]

	// check if a name was given
	f, err := aeternity.StoreAccountToKeyStoreFile(account, p, accountFileName)
	if err != nil {
		return err
	}

	aeternity.Pp(
		"Wallet path", f,
		"Account address", account.Address,
	)

	return nil
}

// balanceCmd implements the account balance subcommand
var balanceCmd = &cobra.Command{
	Use:   "balance ACCOUNT_KEYSTORE",
	Short: "Get the balance of an account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  balanceFunc,
}

func balanceFunc(cmd *cobra.Command, args []string) (err error) {
	aeCli := NewAeCli()
	p, err := getPassword()

	// load the account
	account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	a, err := aeCli.APIGetAccount(account.Address)
	if err != nil {
		return err
	}

	aeternity.PrintObject("account", a)
	return nil
}

// signCmd implements the account sign subcommand
var signCmd = &cobra.Command{
	Use:   "sign ACCOUNT_KEYSTORE UNSIGNED_TRANSACTION",
	Short: "Sign the input (e.g. a transaction)",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  signFunc,
}

func signFunc(cmd *cobra.Command, args []string) (err error) {
	p, err := getPassword()

	// load the account
	account, err := aeternity.LoadAccountFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	txUnsignedBase64 := args[1]
	txSignedBase64, txHash, signature, err := aeternity.SignEncodeTxStr(account, txUnsignedBase64, aeternity.Config.Node.NetworkID)
	if err != nil {
		return err
	}

	aeternity.Pp(
		"Signing account address", account.Address,
		"Signature", signature,
		"Unsigned", txUnsignedBase64,
		"Signed", txSignedBase64,
		"Hash", txHash,
	)
	return nil
}

// saveCmd implements the account save subcommand
var saveCmd = &cobra.Command{
	Use:   "save ACCOUNT_KEYSTORE ACCOUNT_HEX_STRING",
	Short: "Save an account from a hex string to a keystore file",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  saveFunc,
}

func saveFunc(cmd *cobra.Command, args []string) (err error) {
	accountFileName := args[0]
	account, err := aeternity.AccountFromHexString(args[1])
	if err != nil {
		return err
	}

	p, err := getPassword()

	f, err := aeternity.StoreAccountToKeyStoreFile(account, p, accountFileName)
	if err != nil {
		return err
	}

	aeternity.Pp("Keystore path ", f)

	return nil
}

func init() {
	RootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(addressCmd)
	accountCmd.AddCommand(createCmd)
	accountCmd.AddCommand(saveCmd)
	accountCmd.AddCommand(balanceCmd)
	accountCmd.AddCommand(signCmd)
	accountCmd.PersistentFlags().StringVar(&password, "password", "", "Read account password from stdin [WARN: this method is not secure]")
	// account address flags
	addressCmd.Flags().BoolVar(&printPrivateKey, "private-key", false, "Print the private key as hex string")
}
