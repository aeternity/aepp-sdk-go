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

	"github.com/aeternity/aepp-sdk-go/generated/client/operations"
	"github.com/aeternity/aepp-sdk-go/utils"
	"github.com/spf13/cobra"
)

// chainCmd represents the chain command
var chainCmd = &cobra.Command{
	Use:   "chain",
	Short: "Query the state of the chain",
	Long:  ``,
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Query the top block of the chain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if r, err := epochCli.Operations.GetTop(nil); err == nil {
			utils.PrintObject(r.Payload)
		} else {
			utils.Pp("Error:", err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Query the version of the node running the chain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if r, err := epochCli.Operations.GetVersion(nil); err == nil {
			utils.PrintObject(r.Payload)
		} else {
			utils.Pp("Error:", err)
		}
	},
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Query the blocks of the chain one after the other",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var prevHash string
		if r, err := epochCli.Operations.GetTopBlock(nil); err == nil {
			utils.PrintObject(r.Payload)
			prevHash = fmt.Sprint(r.Payload.PrevHash)
		} else {
			utils.Pp("Error:", err)
			os.Exit(1)
		}

		for {
			p := operations.NewGetBlockByHashParams().WithHash(prevHash)
			if r, err := epochCli.Operations.GetBlockByHash(p); err == nil {
				utils.PrintObjectT(" <<>> <<>> <<>> ", r.Payload)
			} else {
				switch err.(type) {
				case *operations.GetBlockByHashBadRequest:
					utils.PrintError("Bad request:", err.(*operations.GetBlockByHashBadRequest).Payload)
				case *operations.GetBlockByHashNotFound:
					utils.PrintError("Block not found:", err.(*operations.GetBlockByHashNotFound).Payload)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(chainCmd)
	chainCmd.AddCommand(topCmd)
	chainCmd.AddCommand(versionCmd)
	chainCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
