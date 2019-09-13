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

	"github.com/aeternity/aepp-sdk-go/v5/aeternity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "aecli",
	Short: "The command line client for the Aeternity blockchain",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var debug bool
var online bool

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string) {
	RootCmd.Version = v
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// newAeNode is just a helper function that gives you a Node so that you don't
// have to maintain a Node global variable (which needs the config vars to be
// read immediately, with this helper function you can defer the reading of the
// variables until the subcommand's execution)
func newAeNode() aeternity.NodeInterface {
	return aeternity.NewNode(aeternity.Config.Node.URL, debug)
}

// newCompiler is just a helper function that gives you a Compiler so that you don't
// have to maintain a Compiler global variable (which needs the config vars to be
// read immediately, with this helper function you can defer the reading of the
// variables until the subcommand's execution)
func newCompiler() *aeternity.Compiler {
	return aeternity.NewCompiler(compilerURL, debug)
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("AETERNITY")
	viper.SetDefault("external-api", aeternity.Config.Node.URL)
	viper.SetDefault("network-id", aeternity.Config.Node.NetworkID)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&aeternity.Config.Node.URL, "external-api", "u", aeternity.Config.Node.URL, "node external API endpoint")
	RootCmd.PersistentFlags().StringVarP(&aeternity.Config.Node.NetworkID, "network-id", "n", aeternity.Config.Node.NetworkID, "network ID for custom private net")
	RootCmd.PersistentFlags().StringVarP(&compilerURL, "compiler-url", "c", "http://localhost:3080", "Compiler URL")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug")
	RootCmd.PersistentFlags().BoolVar(&aeternity.Config.Tuning.OutputFormatJSON, "json", false, "print output in json format")
}
