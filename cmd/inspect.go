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

	"github.com/aeternity/aepp-sdk-go/v5/binary"
	"github.com/aeternity/aepp-sdk-go/v5/naet"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect an object of the blockchain",
	Long: `Inspect an object of the chain

Valid object to inspect are block hash, transaction hash, accounts`,
	RunE: func(cmd *cobra.Command, args []string) error {
		node := newAeNode()
		return inspectFunc(node, args)
	},
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
	PrintObject(title, v)
}

type nodeGetters interface {
	naet.GetGenerationByHeighter
	naet.GetNameEntryByNamer
	naet.GetAccounter
	naet.GetMicroBlockHeaderByHasher
	naet.GetMicroBlockTransactionsByHasher
	naet.GetKeyBlockByHasher
	naet.GetTransactionByHasher
	naet.GetOracleByPubkeyer
}

func printNameEntry(conn naet.GetNameEntryByNamer, name string) (err error) {
	v, err := conn.GetNameEntryByName(name)
	if err != nil {
		return err
	}
	printResult("aens", v, err)
	return err
}

func printAccount(conn naet.GetAccounter, accountID string) (err error) {
	v, err := conn.GetAccount(accountID)
	if err != nil {
		return err
	}
	printResult("account", v, err)
	return err
}

type getMicroBlockHeaderTransactions interface {
	naet.GetMicroBlockHeaderByHasher
	naet.GetMicroBlockTransactionsByHasher
}

func printMicroBlockAndTransactions(conn getMicroBlockHeaderTransactions, mbHash string) (err error) {
	v, err := conn.GetMicroBlockHeaderByHash(mbHash)
	if err != nil {
		return err
	}
	printResult("block", v, err)
	v1, err := conn.GetMicroBlockTransactionsByHash(mbHash)
	if err != nil {
		return err
	}
	printResult("transaction", v1, err)
	return err
}

func printKeyBlockByHash(conn naet.GetKeyBlockByHasher, kbHash string) (err error) {
	v, err := conn.GetKeyBlockByHash(kbHash)
	if err != nil {
		return err
	}
	printResult("key-block", v, err)
	dumpV(v)
	return err
}

func printTransactionByHash(conn naet.GetTransactionByHasher, txHash string) (err error) {
	v, err := conn.GetTransactionByHash(txHash)
	if err != nil {
		return err
	}
	printResult("transaction", v, err)
	return err
}

func printOracleByPubkey(conn naet.GetOracleByPubkeyer, oracleID string) (err error) {
	v, err := conn.GetOracleByPubkey(oracleID)
	if err != nil {
		return err
	}
	printResult("oracle", v, err)
	return err
}

func inspectFunc(conn nodeGetters, args []string) (err error) {
	for _, object := range args {
		// height
		if matched, _ := regexp.MatchString(`^\d+$`, object); matched {
			height, _ := strconv.ParseUint(object, 10, 64)
			PrintGenerationByHeight(conn, height)
			continue
		}
		// name
		if strings.HasSuffix(object, ".aet") {
			printNameEntry(conn, object)
			continue
		}

		switch binary.GetHashPrefix(object) {
		case binary.PrefixAccountPubkey:
			printAccount(conn, object)
		case binary.PrefixMicroBlockHash:
			printMicroBlockAndTransactions(conn, object)
		case binary.PrefixKeyBlockHash:
			printKeyBlockByHash(conn, object)
		case binary.PrefixTransactionHash:
			printTransactionByHash(conn, object)
		case binary.PrefixOraclePubkey:
			printOracleByPubkey(conn, object)
		default:
			return fmt.Errorf("Object %v not yet supported", object)
		}
	}
	return nil
}
