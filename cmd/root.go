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

	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "aecli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var cfgFile string
var debug bool
var aeCli *aeternity.Ae

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string) {
	RootCmd.Version = v
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is /etc/distill/settings.yaml)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// set configuration paramteres
	viper.SetConfigName(aeternity.ConfigFilename) // name of config file (without extension)
	viper.AddConfigPath("$HOME/.aeternity")       // adding home directory as first search path
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // read in environment variables that match
	// if there is the config file read it
	if len(cfgFile) > 0 { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&aeternity.Config)
		aeternity.Config.Defaults()
		aeternity.Config.Validate()
		aeternity.Config.ConfigPath = viper.ConfigFileUsed()
	} else {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			if do := utils.AskYes("A configuration file was not found, would you like to generate one?", true); do {
				aeternity.GenerateDefaultConfig("config/"+aeternity.ConfigFilename+".yaml", RootCmd.Version)
				aeternity.Config.Save()
				fmt.Println("Configuration created")
			} else {
				fmt.Println("Configuration file not found!!")
				os.Exit(1)
			}
		}

	}

	aeCli = aeternity.NewCli(aeternity.Config.P.Epoch.URL, debug)
}
