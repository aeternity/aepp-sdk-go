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
	"strings"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet PRIVATE_KEY_PATH",
	Short: "Interact with a wallet",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Printf("%#v", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside walletCmd PreRun with args: %v\n", args)
	},
}

// addressCmd represents the address subcommand
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Print the aeternity wallet address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Work your own magic here
		fmt.Printf("%#v", args)
	},
}

// spendCmd represents the spend subcommand
var spendCmd = &cobra.Command{
	Use:   "spend",
	Short: "Print the aeternity wallet spend",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		pkPath := ""
		amount := int64(-1)
		recipient := ""
		var err error = nil
		// find arguments
		for _, a := range args {
			if strings.HasPrefix(a, aeternity.PrefixAccount) {
				recipient = a
			} else if ok, v := utils.IsPositiveInt64(a); ok {
				amount = v
			} else {
				pkPath = a
			}
		}
		if len(pkPath) == 0 {
			err = fmt.Errorf("Missing key path")
		}
		if len(recipient) == 0 {
			err = fmt.Errorf("Missing recipient")
		}
		if amount <= 0 {
			err = fmt.Errorf("Amount must be a strictly positve int")
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// load the private key
		kp, err := aeternity.Load(pkPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		txHash, signature, err := aeternity.Spend(epochCli, kp, recipient, amount)
		// TODO: print also the ttl
		utils.Pp(
			"Sender Address", kp.Address,
			"Recipient Address", recipient,
			"Amount", amount,
			"Transaction hash", txHash,
			"Signature", signature,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(walletCmd)
	walletCmd.AddCommand(addressCmd)
	walletCmd.AddCommand(spendCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
