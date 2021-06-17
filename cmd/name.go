package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aeternity/aepp-sdk-go/v9/naet"

	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Lookup a name on AENS",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node := newAeNode()
		err := nameFunc(node, args[0])
		return err
	},
}

func nameFunc(conn naet.GetNameEntryByNamer, name string) (err error) {
	ans, err := conn.GetNameEntryByName(name)
	if err != nil {
		return err
	}

	o, err := json.MarshalIndent(ans, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(o))
	return nil
}
func init() {
	RootCmd.AddCommand(nameCmd)
}
