package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aeternity/aepp-sdk-go/v8/naet"
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

func compileFunc(conn naet.CompileContracter, args []string) (err error) {
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

func encodeCalldataFunc(conn naet.EncodeCalldataer, args []string) (err error) {
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

type decodeCalldataer interface {
	naet.DecodeCalldataBytecoder
	naet.DecodeCalldataSourcer
}

var decodeCalldataBytecodeCmd = &cobra.Command{
	Use:   "decodeCalldataBytecode BYTECODE CALLDATA [..ARGS]",
	Short: "Decode contract function calls. Needs the path to contract source file/compiled bytecode",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return decodeCalldataBytecodeFunc(compiler, args)
	},
	SilenceUsage: true,
}

func decodeCalldataBytecodeFunc(conn decodeCalldataer, args []string) (err error) {
	if !IsBytecode(args[0]) {
		return fmt.Errorf("%s is not bytecode", args[0])
	}
	if !IsBytecode(args[1]) {
		return fmt.Errorf("%s is not bytecode", args[0])
	}

	r, err := conn.DecodeCalldataBytecode(args[0], args[1])
	if err != nil {
		return
	}

	fmt.Println(*r.Function, r.Arguments)
	return nil
}

var decodeCalldataSourceCmd = &cobra.Command{
	Use:   "decodeCalldataSource SOURCE_FILE FUNCTION_NAME CALLDATA [..ARGS]",
	Short: "Decode contract function calls. Needs the path to contract source file/compiled bytecode",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return decodeCalldataSourceFunc(compiler, args)
	},
	SilenceUsage: true,
}

func decodeCalldataSourceFunc(conn decodeCalldataer, args []string) (err error) {
	source, err := readSource(args[0])
	if err != nil {
		return err
	}
	if !IsBytecode(args[2]) {
		return fmt.Errorf("%s is not bytecode", args[0])
	}

	r, err := conn.DecodeCalldataSource(source, args[1], args[2])

	fmt.Println(*r.Function, r.Arguments)
	return
}

var generateAciCmd = &cobra.Command{
	Use:   "generateaci FILENAME",
	Short: "Generate ACI out of source code",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		compiler := newCompiler()
		return generateAciFunc(compiler, args)
	},
	SilenceUsage: true,
}

func generateAciFunc(conn naet.GenerateACIer, args []string) (err error) {
	source, err := readSource(args[0])
	if err != nil {
		return
	}

	aci, err := conn.GenerateACI(source)
	if err != nil {
		return
	}
	PrintObject("ACI", aci)
	return nil
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
	contractCmd.AddCommand(decodeCalldataBytecodeCmd)
	contractCmd.AddCommand(decodeCalldataSourceCmd)
	contractCmd.AddCommand(generateAciCmd)
}
