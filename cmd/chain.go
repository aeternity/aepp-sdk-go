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
	"errors"
	"fmt"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return topFunc(aeNode, args)
	},
}

func topFunc(conn aeternity.GetTopBlocker, args []string) error {
	v, err := conn.GetTopBlock()
	if err != nil {
		return err
	}
	PrintObject("block", v)
	return nil
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status and status of the node running the chain",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return statusFunc(aeNode, args)
	},
}

func statusFunc(conn aeternity.GetStatuser, args []string) (err error) {
	v, err := conn.GetStatus()
	if err != nil {
		return err
	}
	PrintObject("node", v)
	return nil
}

var limit, startFromHeight uint64
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Query the blocks of the chain one after the other",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return playFunc(aeNode, args)
	},
}

type playFuncInterface interface {
	aeternity.GetHeighter
	getGenerationMicroBlockTransactioner
}

func playFunc(conn playFuncInterface, args []string) (err error) {
	blockHeight, err := conn.GetHeight()
	if err != nil {
		return err
	}

	// deal with the height parameter
	if startFromHeight > blockHeight {
		err := fmt.Errorf("Height (%d) is greater that the top block (%d)", startFromHeight, blockHeight)
		return err
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
		PrintGenerationByHeight(conn, blockHeight)
		fmt.Println("")
	}

	return nil
}

var broadcastCmd = &cobra.Command{
	Use:   "broadcast SIGNED_TRANSACTION",
	Short: "Broadcast a transaction to the network",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return broadcastFunc(aeNode, args)
	},
}

func broadcastFunc(conn aeternity.PostTransactioner, args []string) (err error) {
	// Load variables from arguments
	txSignedBase64 := args[0]

	if len(txSignedBase64) == 0 || txSignedBase64[0:3] != "tx_" {
		err := errors.New("Error, missing or invalid recipient address")
		return err
	}

	// Transform tx_ string back to an RLP bytearray to calculate its hash
	txRLP, err := aeternity.Decode(txSignedBase64)
	if err != nil {
		return err
	}
	rlpTxHashRaw, err := aeternity.Blake2bHash(txRLP)
	if err != nil {
		return err
	}
	signedEncodedTxHash := aeternity.Encode(aeternity.PrefixTransactionHash, rlpTxHashRaw)

	err = conn.PostTransaction(txSignedBase64, signedEncodedTxHash)
	if err != nil {
		errFinal := fmt.Errorf("Error while broadcasting transaction: %v", err)
		return errFinal
	}

	return nil
}

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: "Get the absolute TTL for a Transaction",
	Long:  `Get the absolute TTL (node's height + recommended TTL offset) for a Transaction`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return ttlFunc(aeNode, args)
	},
}

func ttlFunc(conn aeternity.GetHeighter, args []string) (err error) {
	height, err := conn.GetHeight()
	if err != nil {
		errFinal := fmt.Errorf("Error getting height from the node: %v", err)
		return errFinal
	}
	ttl = height + aeternity.Config.Client.TTL
	fmt.Println(ttl)
	return nil
}

var networkIDCmd = &cobra.Command{
	Use:   "networkid",
	Short: "Get the node's network_id",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		aeNode := newAeNode()
		return networkIDFunc(aeNode, args)
	},
}

func networkIDFunc(conn aeternity.GetStatuser, args []string) (err error) {
	resp, err := conn.GetStatus()
	if err != nil {
		errFinal := fmt.Errorf("Error getting status information from the node: %v", err)
		return errFinal
	}
	fmt.Println(*resp.NetworkID)
	return nil
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
