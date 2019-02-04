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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"net/http/httptrace"
	"os"
	"regexp"
	"time"

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
	Run: func(cmd *cobra.Command, args []string) {
		v, err := aeCli.APIGetTopBlock()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		aeternity.PrintObject("block", v)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status and status of the node running the chain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		v, err := aeCli.APIGetStatus()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		aeternity.PrintObject("epoch node", v)
	},
}

var limit, startFromHeight uint64
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Query the blocks of the chain one after the other",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	RootCmd.AddCommand(chainCmd)
	chainCmd.AddCommand(topCmd)
	chainCmd.AddCommand(statusCmd)
	chainCmd.AddCommand(playCmd)
	chainCmd.AddCommand(proxyCmd)

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

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "experimental proxy execution",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		peers := new(Peers)
		fmt.Println("Proxy starting")
		sl := sling.New().Base("https://sdk-mainnet.aepps.com")
		sl = sl.Set("User-Agent", "nodelb")
		request, err := sl.Request()

		trace := &httptrace.ClientTrace{
			DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
				fmt.Printf("DNS Info: %+v\n", dnsInfo)
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				fmt.Printf("Got Conn: %+v\n", connInfo)
			},
		}
		request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

		fmt.Println("Requesting peers")
		_, err = sl.Get("/v2/debug/peers").ReceiveSuccess(peers)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		r := regexp.MustCompile(`\d+\.\d+.\d+.\d+`)
		for _, p := range peers.Peers {
			ip := r.FindString(p)
			fmt.Println("Peer ", p, "found (", ip, ")")
			timeGet(fmt.Sprint("http://", ip, ":3013/v2/blocks/top"))
		}

	},
}

// Peers the struct for peers
type Peers struct {
	Blocked  []string `json:"blocked"`
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
	Peers    []string `json:"peers"`
}

type ChainHeight struct {
	Height uint64 `json:"height"`
}

func timeGet(url string) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	netClient.Get("utl")
	req, _ := http.NewRequest("GET", url, nil)

	var start, connect, dns, tlsHandshake time.Time

	trace := &httptrace.ClientTrace{

		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Done: %v\n", time.Since(dns))
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			fmt.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("Connect time: %v\n", time.Since(connect))
		},

		GotFirstResponseByte: func() {
			fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	if resp, err := http.DefaultTransport.RoundTrip(req); err != nil {
		fmt.Println(err)
		return
	} else {
		defer resp.Body.Close()
		ch := new(ChainHeight)
		json.NewDecoder(resp.Body).Decode(ch)
	}
	fmt.Printf("Total time: %v\n", time.Since(start))
}
