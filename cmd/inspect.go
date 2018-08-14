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
	"strings"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/client/operations"
	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect an object of the blockchain",
	Long: `Inspect an object of the chain

Valid object to inspect are block hash, transaction hash, accounts`,
	Run:  inspect,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func inspect(cmd *cobra.Command, args []string) {
	for _, object := range args {

		// name
		if strings.HasSuffix(object, ".aet") {
			p := operations.NewGetNameParams().WithName(object)
			if r, err := epochCli.Operations.GetName(p); err == nil {
				utils.PrintObjectT("Name", r.Payload)
			} else {
				switch err.(type) {
				case *operations.GetNameBadRequest:
					utils.PrintError("Bad request:", err.(*operations.GetNameBadRequest).Payload)
				case *operations.GetNameNotFound:
					utils.PrintError("Name not found:", err.(*operations.GetNameNotFound).Payload)
				default:
					utils.Pp("Unknown error:", err)
				}
			}
			continue
		}

		switch object[0:3] {
		case aeternity.PrefixAccount:
			// account balance
			p := operations.NewGetAccountBalanceParams().WithAddress(object)
			if r, err := epochCli.Operations.GetAccountBalance(p); err == nil {
				utils.Pp("Balance", r.Payload.Balance)
			} else {
				switch err.(type) {
				case *operations.GetAccountBalanceBadRequest:
					utils.PrintError("Bad request:", err.(*operations.GetAccountBalanceBadRequest).Payload)
				case *operations.GetAccountBalanceNotFound:
					utils.PrintError("Account not found:", err.(*operations.GetAccountBalanceNotFound).Payload)
				default:
					utils.Pp("Unknown error:", err)
				}
			}

		case aeternity.PrefixBlockHash:
			// block
			p := operations.NewGetBlockByHashParams().WithHash(object)
			if r, err := epochCli.Operations.GetBlockByHash(p); err == nil {
				utils.PrintObject(r.Payload)
			} else {
				switch err.(type) {
				case *operations.GetBlockByHashBadRequest:
					utils.PrintError("Bad request:", err.(*operations.GetBlockByHashBadRequest).Payload)
				case *operations.GetBlockByHashNotFound:
					utils.PrintError("Block not found:", err.(*operations.GetBlockByHashNotFound).Payload)
				default:
					utils.Pp("Unknown error:", err)
				}
			}
		case aeternity.PrefixTxHash:
			// transaction
			p := operations.NewGetTxParams().
				WithTxHash(object).
				WithTxEncoding(&aeternity.Config.Tuning.ResponseEncoding)

			if r, err := epochCli.Operations.GetTx(p); err == nil {
				utils.PrintObject(r.Payload)
			} else {
				switch err.(type) {
				case *operations.GetTxBadRequest:
					utils.PrintError("Bad request:", err.(*operations.GetTxBadRequest).Payload)
				case *operations.GetTxNotFound:
					utils.PrintError("Tx not found:", err.(*operations.GetTxBadRequest).Payload)
				default:
					utils.Pp("Unknown error:", err)
				}
			}

		case aeternity.PrefixBlockTxHash:
			// block transaction
			p := operations.NewGetTransactionFromBlockHashParams().
				WithHash(object)
			if r, err := epochCli.Operations.GetTransactionFromBlockHash(p); err == nil {
				utils.PrintObject(r.Payload)
			} else {
				switch err.(type) {
				case *operations.GetTransactionFromBlockHashNotFound:
					utils.PrintError("Tx not found:", err.(*operations.GetTransactionFromBlockHashNotFound).Payload)
				default:
					utils.Pp("Unknown error:", err)
				}
			}
			//utils.PrintObject(r.Payload, err)
		default:
			fmt.Println("Object", object, "not yet supported")
		}

	}
}
