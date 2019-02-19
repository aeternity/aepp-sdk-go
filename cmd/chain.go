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
	Run:   topFunc,
}

func topFunc(cmd *cobra.Command, args []string) {
	aeCli := NewAeCli()
	v, err := aeCli.APIGetTopBlock()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	aeternity.PrintObject("block", v)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status and status of the node running the chain",
	Long:  ``,
	Run:   statusFunc,
}

func statusFunc(cmd *cobra.Command, args []string) {
	aeCli := NewAeCli()
	v, err := aeCli.APIGetStatus()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	aeternity.PrintObject("epoch node", v)
}

var limit, startFromHeight uint64
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Query the blocks of the chain one after the other",
	Long:  ``,
	Run:   playFunc,
}

func playFunc(cmd *cobra.Command, args []string) {
	aeCli := NewAeCli()
	blockHeight, err := aeCli.APIGetHeight()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// deal with the height parameter
	if startFromHeight > blockHeight {
		fmt.Printf("Height (%d) is greater that the top block (%d)", startFromHeight, blockHeight)
		os.Exit(1)
	}

	if startFromHeight > 0 {
		blockHeight = startFromHeight
	}
	// deal with the limit parameter
	targetHeight := uint64(0)
	if limit > 0 {
		th := blockHeight - limit
		if th > targetHeight {
			targetHeight = th
		}
	}
	// run the play
	for ; blockHeight > targetHeight; blockHeight-- {
		aeCli.PrintGenerationByHeight(blockHeight)
		fmt.Println("")
	}
}

var broadcastCmd = &cobra.Command{
	Use:   "broadcast SIGNED_TRANSACTION",
	Short: "Broadcast a transaction to the network",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   broadcastFunc,
}

func broadcastFunc(cmd *cobra.Command, args []string) {
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
}

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: "Get the absolute TTL for a Transaction",
	Long:  `Get the absolute TTL (node's height + recommended TTL offset) for a Transaction`,
	Args:  cobra.ExactArgs(0),
	Run:   ttlFunc,
}

func ttlFunc(cmd *cobra.Command, args []string) {
	ae := NewAeCli()
	height, err := ae.APIGetHeight()
	if err != nil {
		fmt.Printf("Error getting height from the node: %v", err)
		return
	}
	fmt.Println(height + aeternity.Config.Client.TTL)
}

var networkIDCmd = &cobra.Command{
	Use:   "networkid",
	Short: "Get the node's network_id",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run:   networkIDFunc,
}

func networkIDFunc(cmd *cobra.Command, args []string) {
	ae := NewAeCli()
	resp, err := ae.APIGetStatus()
	if err != nil {
		fmt.Printf("Error getting status information from the node: %v", err)
		return
	}
	fmt.Println(*resp.NetworkID)
}

func init() {
	RootCmd.AddCommand(chainCmd)
	chainCmd.AddCommand(topCmd)
	chainCmd.AddCommand(statusCmd)
	chainCmd.AddCommand(playCmd)
	chainCmd.AddCommand(broadcastCmd)
	chainCmd.AddCommand(ttlCmd)
	chainCmd.AddCommand(networkIDCmd)

	playCmd.Flags().Uint64Var(&limit, "limit", 0, "Print at max 'limit' generations")
	playCmd.Flags().Uint64Var(&startFromHeight, "height", 0, "Start playing the chain at 'height'")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
