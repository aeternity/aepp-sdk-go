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
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
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
	RunE:  topFunc,
}

func topFunc(cmd *cobra.Command, args []string) (err error) {
	aeNode := newAeNode()
	v, err := topDo(aeNode)
	if err != nil {
		return err
	}
	PrintObject("block", v)
	return nil
}

func topDo(conn aeternity.GetTopBlocker) (kb *models.KeyBlockOrMicroBlockHeader, err error) {
	kb, err = conn.GetTopBlock()
	return
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status and status of the node running the chain",
	Long:  ``,
	RunE:  statusFunc,
}

func statusFunc(cmd *cobra.Command, args []string) (err error) {
	aeNode := newAeNode()
	v, err := statusDo(aeNode)
	if err != nil {
		return err
	}
	PrintObject("node", v)
	return nil
}

func statusDo(conn aeternity.GetStatuser) (status *models.Status, err error) {
	status, err = conn.GetStatus()
	return
}

var limit, startFromHeight uint64
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Query the blocks of the chain one after the other",
	Long:  ``,
	RunE:  playFunc,
}

func playFunc(cmd *cobra.Command, args []string) (err error) {
	aeNode := newAeNode()
	blockHeight, err := aeNode.GetHeight()
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
		PrintGenerationByHeight(aeNode, blockHeight)
		fmt.Println("")
	}

	return nil
}

var broadcastCmd = &cobra.Command{
	Use:   "broadcast SIGNED_TRANSACTION",
	Short: "Broadcast a transaction to the network",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  broadcastFunc,
}

func broadcastFunc(cmd *cobra.Command, args []string) (err error) {
	// Load variables from arguments
	txSignedBase64 := args[0]

	if len(txSignedBase64) == 0 || txSignedBase64[0:3] != "tx_" {
		err := errors.New("Error, missing or invalid recipient address")
		return err
	}

	aeNode := newAeNode()
	err = aeternity.BroadcastTransaction(aeNode, txSignedBase64)
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
	RunE:  ttlFunc,
}

func ttlFunc(cmd *cobra.Command, args []string) (err error) {
	ae := newAeNode()
	ans, err := ttlDo(ae)
	if err != nil {
		return
	}
	fmt.Println(ans)
	return nil
}

func ttlDo(conn aeternity.GetHeighter) (ttl uint64, err error) {
	height, err := conn.GetHeight()
	if err != nil {
		errFinal := fmt.Errorf("Error getting height from the node: %v", err)
		return 0, errFinal
	}
	ttl = height + aeternity.Config.Client.TTL
	return
}

var networkIDCmd = &cobra.Command{
	Use:   "networkid",
	Short: "Get the node's network_id",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	RunE:  networkIDFunc,
}

func networkIDFunc(cmd *cobra.Command, args []string) (err error) {
	ae := newAeNode()
	nID, err := networkIDDo(ae)
	if err != nil {
		return err
	}
	fmt.Println(nID)
	return nil
}

func networkIDDo(conn aeternity.GetStatuser) (networkID string, err error) {
	resp, err := conn.GetStatus()
	if err != nil {
		errFinal := fmt.Errorf("Error getting status information from the node: %v", err)
		return "", errFinal
	}
	networkID = *resp.NetworkID
	return
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
