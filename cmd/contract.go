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
	Use:          "compile FILENAME",
	Short:        "Send a source file to a compiler",
	Long:         ``,
	Args:         cobra.ExactArgs(1),
	RunE:         compileFunc,
	SilenceUsage: true,
}

func compileFunc(cmd *cobra.Command, args []string) (err error) {
	compiler := aeternity.NewCompiler(compilerURL, debug)
	s, err := readSource(args[0])
	if err != nil {
		return err
	}

	bytecode, err := compiler.CompileContract(s)
	fmt.Println(bytecode)
	return err
}

var encodeCalldataCmd = &cobra.Command{
	Use:          "encodeCalldata SOURCE FUNCTIONNAME [..ARGS]",
	Short:        "Encode contract function call data. Needs the original contract source",
	Long:         ``,
	Args:         cobra.MinimumNArgs(2),
	RunE:         encodeCalldataFunc,
	SilenceUsage: true,
}

func encodeCalldataFunc(cmd *cobra.Command, args []string) (err error) {
	compiler := aeternity.NewCompiler(compilerURL, debug)

	s, err := readSource(args[0])
	if err != nil {
		return err
	}

	callData, err := compiler.EncodeCalldata(s, args[1], args[2:])
	if err != nil {
		return err
	}
	fmt.Println(callData)
	return
}

func readSource(path string) (s string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(file)
	return string(b), err
}
func init() {
	RootCmd.AddCommand(contractCmd)
	contractCmd.AddCommand(compileCmd)
	contractCmd.AddCommand(encodeCalldataCmd)
	contractCmd.PersistentFlags().StringVarP(&compilerURL, "compiler-url", "c", "http://localhost:3080", "Compiler URL")
}
