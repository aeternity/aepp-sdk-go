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
	Use:   "compile FILENAME",
	Short: "Send a source file to a compiler",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return compileFunc(compiler, args)
	},
	SilenceUsage: true,
}

func compileFunc(conn aeternity.CompileContracter, args []string) (err error) {
	s, err := readSource(args[0])
	if err != nil {
		return err
	}

	bytecode, err := conn.CompileContract(s)
	fmt.Println(bytecode)
	return err
}

var encodeCalldataCmd = &cobra.Command{
	Use:   "encodeCalldata SOURCE FUNCTIONNAME [..ARGS]",
	Short: "Encode contract function calls. Needs the path to contract source file",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return encodeCalldataFunc(compiler, args)
	},
	SilenceUsage: true,
}

func encodeCalldataFunc(conn aeternity.EncodeCalldataer, args []string) (err error) {
	s, err := readSource(args[0])
	if err != nil {
		return err
	}

	callData, err := conn.EncodeCalldata(s, args[1], args[2:])
	if err != nil {
		return err
	}
	fmt.Println(callData)
	return
}

var decodeCalldataCmd = &cobra.Command{
	Use:   "decodeCalldata SOURCE_FILE/BYTECODE CALLDATA [..ARGS]",
	Short: "Decode contract function calls. Needs the path to contract source file/compiled bytecode",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return decodeCalldataFunc(compiler, args)
	},
	SilenceUsage: true,
}

type decodeCalldataer interface {
	aeternity.DecodeCalldataBytecoder
	aeternity.DecodeCalldataSourcer
}

func decodeCalldataFunc(conn decodeCalldataer, args []string) (err error) {
	var decodeWithSource = func(path string, callData string) (function string, arguments []interface{}, err error) {
		source, err := readSource(path)
		if err != nil {
			return
		}
		r, err := conn.DecodeCalldataSource(source, callData)
		if err != nil {
			return
		}
		arguments = r.Arguments
		function = *r.Function
		return
	}
	var decodeWithBytecode = func(bytecode string, callData string) (function string, arguments []interface{}, err error) {
		r, err := conn.DecodeCalldataBytecode(bytecode, callData)
		if err != nil {
			return
		}
		arguments = r.Arguments
		function = *r.Function
		return
	}

	var function string
	var arguments []interface{}
	if !IsBytecode(args[0]) {
		function, arguments, err = decodeWithSource(args[0], args[1])
	} else {
		function, arguments, err = decodeWithBytecode(args[0], args[1])
	}
	if err != nil {
		return
	}

	fmt.Println(function, arguments)
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
	contractCmd.AddCommand(decodeCalldataCmd)
	contractCmd.PersistentFlags().StringVarP(&compilerURL, "compiler-url", "c", "http://localhost:3080", "Compiler URL")
}
