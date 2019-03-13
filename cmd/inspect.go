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

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect an object of the blockchain",
	Long: `Inspect an object of the chain

Valid object to inspect are block hash, transaction hash, accounts`,
	RunE: inspectFunc,
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

func printResult(title string, v interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	aeternity.PrintObject(title, v)
}

func inspectFunc(cmd *cobra.Command, args []string) (err error) {
	aeCli := NewAeCli()
	for _, object := range args {
		// height
		if matched, _ := regexp.MatchString(`^\d+$`, object); matched {
			height, _ := strconv.ParseUint(object, 10, 64)
			aeCli.PrintGenerationByHeight(height)
			continue
		}
		// name
		if strings.HasSuffix(object, ".aet") {
			v, err := aeCli.APIGetNameEntryByName(object)
			if err != nil {
				return err
			}
			printResult("aens", v, err)
			continue
		}

		switch aeternity.GetHashPrefix(object) {
		case aeternity.PrefixAccountPubkey:
			// account balance
			v, err := aeCli.APIGetAccount(object)
			if err != nil {
				return err
			}

			printResult("account", v, err)

		case aeternity.PrefixMicroBlockHash:
			v, err := aeCli.APIGetMicroBlockHeaderByHash(object)
			if err != nil {
				return err
			}
			printResult("block", v, err)
			v1, err := aeCli.APIGetMicroBlockTransactionsByHash(object)
			if err != nil {
				return err
			}
			printResult("transaction", v1, err)

		case aeternity.PrefixKeyBlockHash:
			// block
			v, err := aeCli.APIGetKeyBlockByHash(object)
			if err != nil {
				return err
			}
			printResult("key-block", v, err)

		case aeternity.PrefixTransactionHash:
			// transaction
			v, err := aeCli.APIGetTransactionByHash(object)
			if err != nil {
				return err
			}
			printResult("transaction", v, err)

		case aeternity.PrefixOraclePubkey:
			// oracle
			v, err := aeCli.APIGetOracleByPubkey(object)
			if err != nil {
				return err
			}
			printResult("oracle", v, err)

		default:
			return fmt.Errorf("Object %v not yet supported", object)
		}
	}
	return nil
}
