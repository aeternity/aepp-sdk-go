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
  "github.com/aeternity/aepp-sdk-go/aeternity"
  "github.com/skratchdot/open-golang/open"
  "github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
  Use:   "config",
  Short: "Print the configuration of the client",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    aeternity.PrintObject("configuration", aeternity.Config.P)
    // aeternity.Pp(
    // 	"Epoch URL", aeternity.Config.P.Epoch.URL,
    // 	"Epoch Internal URL", aeternity.Config.P.Epoch.InternalURL,
    // 	"Epoch Websocket URL", aeternity.Config.P.Epoch.WebsocketURL,
    // )
  },
}

var editCmd = &cobra.Command{
  Use:   "edit",
  Short: "Open the config file for editing",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    aeternity.PrintObject("Configuration path", aeternity.Config.ConfigPath)
    open.Run(aeternity.Config.ConfigPath)
  },
}

var profileCmd = &cobra.Command{
  Use:   "profile",
  Short: "Print the current profile",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    aeternity.Pp(
      "Active profile", aeternity.Config.P.Name,
    )
  },
}

var profileListCmd = &cobra.Command{
  Use:   "list",
  Short: "List the available profiles",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    for _, p := range aeternity.Config.Profiles {
      prefix := ""
      if p.Name == aeternity.Config.P.Name {
        prefix = "  *  "
      }
      aeternity.Pp(
        prefix, p.Name,
      )
    }
  },
}

var profileUseCmd = &cobra.Command{
  Use:   "use PROFILE_NAME",
  Short: "Activate the profile PROFILE_NAME",
  Long:  ``,
  Args:  cobra.RangeArgs(1, 1),
  Run: func(cmd *cobra.Command, args []string) {
    err := aeternity.Config.ActivateProfile(args[0])
    if err != nil {
      fmt.Println(err)
    }
    aeternity.Config.Save()
    fmt.Println("Profile", args[0], "activated")
  },
}

var profileCreateCmd = &cobra.Command{
  Use:   "create PROFILE_NAME",
  Short: "Create the profile PROFILE_NAME",
  Args:  cobra.RangeArgs(1, 1),
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    aeternity.Config.NewProfile(args[0])
    aeternity.Config.ActivateProfile(args[0])
    aeternity.Config.Save()
    fmt.Println("Profile", args[0], "created")
  },
}

func init() {
  RootCmd.AddCommand(configCmd)
  configCmd.AddCommand(editCmd)
  configCmd.AddCommand(profileCmd)
  profileCmd.AddCommand(profileListCmd)
  profileCmd.AddCommand(profileCreateCmd)
  profileCmd.AddCommand(profileUseCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // configCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
