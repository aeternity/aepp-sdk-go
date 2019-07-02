package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
)

var compilerURL string
var contractCmd = &cobra.Command{
	Use:   "contract subcommand",
	Short: "Compile, call or deploy smart contracts",
	Long:  ``,
}

var compileCmd = &cobra.Command{
	Use:          "compile FILENAME COMPILER_URL",
	Short:        "Send a source file to a compiler",
	Long:         ``,
	Args:         cobra.ExactArgs(1),
	RunE:         compileFunc,
	SilenceUsage: true,
}

func compileFunc(cmd *cobra.Command, args []string) (err error) {
	compiler := aeternity.NewCompiler(compilerURL, debug)
	file, err := os.Open(args[0])
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	bytecode, err := compiler.CompileContract(string(b))
	fmt.Println(bytecode)
	return err
}

func init() {
	RootCmd.AddCommand(contractCmd)
	contractCmd.AddCommand(compileCmd)
	contractCmd.PersistentFlags().StringVarP(&compilerURL, "compiler-url", "c", "http://localhost:3080", "Compiler URL")
}
