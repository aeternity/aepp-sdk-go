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
	"path/filepath"

	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/shibukawa/configdir"
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

var cfgFile string
var debug, outputFormatJSON bool
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
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file to load (defaults to $HOME/.aeternity/config.yaml")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug")
	RootCmd.PersistentFlags().BoolVar(&outputFormatJSON, "json", false, "print output in json format")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// retrieve the directory (os dependent) where the config file exists
	configDirs := configdir.New("aeternity", "aecli")
	globalCfg := configDirs.QueryFolders(configdir.Global)[0]
	// set configuration paramteres
	viper.SetConfigName(aeternity.ConfigFilename) // name of config file (without extension)
	viper.AddConfigPath(globalCfg.Path)           // adding home directory as first search path
	//viper.AutomaticEnv()                          // read in environment variables that match
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
		aeternity.Config.KeysFolder = filepath.Join(globalCfg.Path, "keys")
	} else {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			if do := utils.AskYes("A configuration file was not found, would you like to generate one?", true); do {
				configFilePath := filepath.Join(globalCfg.Path, aeternity.ConfigFilename+".yml")
				aeternity.GenerateDefaultConfig(configFilePath, RootCmd.Version)
				aeternity.Config.Save()
			} else {
				fmt.Println("Configuration file not found!!")
				os.Exit(1)
			}
		}

	}

	aeternity.Config.P.Tuning.OutputFormatJSON = outputFormatJSON
	aeCli = aeternity.NewCli(aeternity.Config.P.Epoch.URL, debug)

}
