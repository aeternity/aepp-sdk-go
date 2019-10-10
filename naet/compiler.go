package naet

import (
	compiler_client "github.com/aeternity/aepp-sdk-go/v6/swagguard/compiler/client"
	"github.com/aeternity/aepp-sdk-go/v6/swagguard/compiler/client/operations"
	models "github.com/aeternity/aepp-sdk-go/v6/swagguard/compiler/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Compiler wraps around the swagger-generated Compiler. Unlike the
// swagger-generated Compiler, this Compiler struct builds the swagger HTTP
// requests from the Golang native arguments that its methods receive, and uses
// the swagger-generated Compiler's endpoints to send these requests off. It
// then parses the swagger response and makes it as native-Go-code-friendly as
// possible.
type Compiler struct {
	*compiler_client.Compiler
}

// NewCompiler creates a new Compiler instance from a URL
func NewCompiler(compilerURL string, debug bool) *Compiler {
	host, schemas := urlComponents(compilerURL)
	transport := httptransport.New(host, "", schemas)
	transport.SetDebug(debug)
	c := compiler_client.New(transport, strfmt.Default)
	compiler := &Compiler{
		Compiler: c,
	}
	return compiler
}

// APIVersioner guarantees that one can run a APIVersion() method on the
// mocked/real network connection to the aesophia compiler
type APIVersioner interface {
	APIVersion() (version string, err error)
}

// APIVersion connects to the compiler and returns its version string, e.g.
// 3.1.0
func (c *Compiler) APIVersion() (version string, err error) {
	result, err := c.Compiler.Operations.APIVersion(nil)
	if err != nil {
		return "", err
	}
	version = *result.Payload.APIVersion
	return
}

// CompileContracter guarantees that one can run a CompileContract() method on
// the mocked/real network connection to the aesophia compiler
type CompileContracter interface {
	CompileContract(source string, backend string) (bytecode string, err error)
}

// CompileContract abstracts away the swagger specifics of posting to /compile
func (c *Compiler) CompileContract(source string, backend string) (bytecode string, err error) {
	contract := &models.Contract{Code: &source, Options: &models.CompileOpts{
		Backend:    backend,
		FileSystem: nil,
		SrcFile:    "",
	}}
	params := operations.NewCompileContractParams().WithBody(contract)
	result, err := c.Compiler.Operations.CompileContract(params)
	if err != nil {
		return "", err
	}
	bytecode = string(result.Payload.Bytecode)
	return
}

// DecodeCallResulter guarantees that one can run a DecodeCallResult() method on
// the mocked/real network connection to the aesophia compiler
type DecodeCallResulter interface {
	DecodeCallResult(callResult string, callValue string, function string, source string, backend string) (answer interface{}, err error)
}

// DecodeCallResult abstracts away the swagger specifics of posting to
// /decode-call-result
func (c *Compiler) DecodeCallResult(callResult string, callValue string, function string, source string, backend string) (answer interface{}, err error) {
	sophiaCallResultInput := &models.SophiaCallResultInput{
		CallResult: &callResult,
		CallValue:  &callValue,
		Function:   &function,
		Options: &models.CompileOpts{
			Backend:    backend,
			FileSystem: nil,
			SrcFile:    "",
		},
		Source: &source,
	}
	params := operations.NewDecodeCallResultParams().WithBody(sophiaCallResultInput)
	result, err := c.Compiler.Operations.DecodeCallResult(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// DecodeCalldataBytecoder guarantees that one can run a
// DecodeCalldataBytecode() method on the mocked/real network connection to the
// aesophia compiler
type DecodeCalldataBytecoder interface {
	DecodeCalldataBytecode(bytecode string, calldata string, backend string) (decodedCallData *models.DecodedCalldata, err error)
}

// DecodeCalldataBytecode abstracts away the swagger specifics of posting to
// /decode-calldata/bytecode
func (c *Compiler) DecodeCalldataBytecode(bytecode string, calldata string, backend string) (decodedCallData *models.DecodedCalldata, err error) {
	decodeCalldataBytecode := &models.DecodeCalldataBytecode{
		Bytecode: models.EncodedByteArray(bytecode),
		Calldata: models.EncodedByteArray(calldata),
		Backend:  backend,
	}
	params := operations.NewDecodeCalldataBytecodeParams().WithBody(decodeCalldataBytecode)
	result, err := c.Compiler.Operations.DecodeCalldataBytecode(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// DecodeCalldataSourcer guarantees that one can run a DecodeCalldataSource()
// method on the mocked/real network connection to the aesophia compiler
type DecodeCalldataSourcer interface {
	DecodeCalldataSource(source string, function string, callData string, backend string) (decodedCallData *models.DecodedCalldata, err error)
}

// DecodeCalldataSource abstracts away the swagger specifics of posting to
// /decode-calldata/source
func (c *Compiler) DecodeCalldataSource(source string, function string, callData string, backend string) (decodedCallData *models.DecodedCalldata, err error) {
	p := &models.DecodeCalldataSource{
		Calldata: models.EncodedByteArray(callData),
		Function: &function,
		Options: &models.CompileOpts{
			Backend:    backend,
			FileSystem: nil,
			SrcFile:    "",
		},
		Source: &source,
	}
	params := operations.NewDecodeCalldataSourceParams().WithBody(p)
	result, err := c.Compiler.Operations.DecodeCalldataSource(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// DecodeDataer guarantees that one can run a DecodeData() method on the
// mocked/real network connection to the aesophia compiler
type DecodeDataer interface {
	DecodeData(data string, sophiaType string) (decodedData *models.SophiaJSONData, err error)
}

// DecodeData abstracts away the swagger specifics of posting to /decode-data
func (c *Compiler) DecodeData(data string, sophiaType string) (decodedData *models.SophiaJSONData, err error) {
	p := &models.SophiaBinaryData{
		Data:       &data,
		SophiaType: &sophiaType,
	}
	params := operations.NewDecodeDataParams().WithBody(p)
	result, err := c.Compiler.Operations.DecodeData(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// EncodeCalldataer guarantees that one can run a EncodeCalldata() method on the
// mocked/real network connection to the aesophia compiler
type EncodeCalldataer interface {
	EncodeCalldata(source string, function string, args []string, backend string) (callData string, err error)
}

// EncodeCalldata abstracts away the swagger specifics of posting to
// /encode-calldata
func (c *Compiler) EncodeCalldata(source string, function string, args []string, backend string) (callData string, err error) {
	f := &models.FunctionCallInput{
		Arguments: args,
		Function:  &function,
		Options: &models.CompileOpts{
			Backend:    backend,
			FileSystem: nil,
			SrcFile:    "",
		},
		Source: &source,
	}
	params := operations.NewEncodeCalldataParams().WithBody(f)
	result, err := c.Compiler.Operations.EncodeCalldata(params)
	if err != nil {
		return
	}

	s := string(result.Payload.Calldata)
	return s, err
}

// GenerateACIer guarantees that one can run a GenerateACI() method on the
// mocked/real network connection to the aesophia compiler
type GenerateACIer interface {
	GenerateACI(source string, backend string) (aci *models.ACI, err error)
}

// GenerateACI abstracts away the swagger specifics of posting to /aci
func (c *Compiler) GenerateACI(source string, backend string) (aci *models.ACI, err error) {
	contract := &models.Contract{Code: &source, Options: &models.CompileOpts{
		Backend:    backend,
		FileSystem: nil,
		SrcFile:    "",
	}}
	params := operations.NewGenerateACIParams().WithBody(contract)
	result, err := c.Compiler.Operations.GenerateACI(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// SophiaVersioner guarantees that one can run a SophiaVersion() method on the
// mocked/real network connection to the aesophia compiler
type SophiaVersioner interface {
	SophiaVersion() (version string, err error)
}

// SophiaVersion abstracts away the swagger specifics of getting /version
func (c *Compiler) SophiaVersion() (version string, err error) {
	result, err := c.Compiler.Operations.Version(nil)
	if err != nil {
		return
	}

	return string(*result.Payload.Version), err
}
