package aeternity

import (
	compiler_client "github.com/aeternity/aepp-sdk-go/swagguard/compiler/client"
	"github.com/aeternity/aepp-sdk-go/swagguard/compiler/client/operations"
	models "github.com/aeternity/aepp-sdk-go/swagguard/compiler/models"
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

// CompileContract abstracts away the swagger specifics of posting to /compile
func (c *Compiler) CompileContract(source string) (bytecode string, err error) {
	contract := &models.Contract{Code: &source, Options: &models.CompileOpts{}}
	params := operations.NewCompileContractParams().WithBody(contract)
	result, err := c.Compiler.Operations.CompileContract(params)
	if err != nil {
		return "", err
	}
	bytecode = string(result.Payload.Bytecode)
	return
}

// DecodeCallResult abstracts away the swagger specifics of posting to
// /decode-call-result
func (c *Compiler) DecodeCallResult(callResult string, callValue string, function string, source string) (answer interface{}, err error) {
	sophiaCallResultInput := &models.SophiaCallResultInput{
		CallResult: &callResult,
		CallValue:  &callValue,
		Function:   &function,
		Source:     &source,
	}
	params := operations.NewDecodeCallResultParams().WithBody(sophiaCallResultInput)
	result, err := c.Compiler.Operations.DecodeCallResult(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// DecodeCalldataBytecode abstracts away the swagger specifics of posting to /decode-calldata/bytecode
func (c *Compiler) DecodeCalldataBytecode(bytecode string, calldata string) (decodedCallData *models.DecodedCalldata, err error) {
	decodeCalldataBytecode := &models.DecodeCalldataBytecode{
		Bytecode: models.EncodedByteArray(bytecode),
		Calldata: models.EncodedByteArray(calldata),
	}
	params := operations.NewDecodeCalldataBytecodeParams().WithBody(decodeCalldataBytecode)
	result, err := c.Compiler.Operations.DecodeCalldataBytecode(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// DecodeCalldataSource abstracts away the swagger specifics of posting to /decode-calldata/source
func (c *Compiler) DecodeCalldataSource(source string, callData string) (decodedCallData *models.DecodedCalldata, err error) {
	p := &models.DecodeCalldataSource{
		Source:   source,
		Calldata: models.EncodedByteArray(callData),
	}
	params := operations.NewDecodeCalldataSourceParams().WithBody(p)
	result, err := c.Compiler.Operations.DecodeCalldataSource(params)
	if err != nil {
		return
	}

	return result.Payload, err
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

// EncodeCalldata abstracts away the swagger specifics of posting to /encode-calldata
func (c *Compiler) EncodeCalldata(source string, function string, args []string) (callData string, err error) {
	f := &models.FunctionCallInput{
		Arguments: args,
		Function:  &function,
		Source:    &source,
	}
	params := operations.NewEncodeCalldataParams().WithBody(f)
	result, err := c.Compiler.Operations.EncodeCalldata(params)
	if err != nil {
		return
	}

	s := string(result.Payload.Calldata)
	return s, err
}

// GenerateACI abstracts away the swagger specifics of posting to /aci
func (c *Compiler) GenerateACI(source string) (aci *models.ACI, err error) {
	contract := &models.Contract{Code: &source, Options: &models.CompileOpts{}}
	params := operations.NewGenerateACIParams().WithBody(contract)
	result, err := c.Compiler.Operations.GenerateACI(params)
	if err != nil {
		return
	}

	return result.Payload, err
}

// SophiaVersion abstracts away the swagger specifics of getting /version
func (c *Compiler) SophiaVersion() (version string, err error) {
	result, err := c.Compiler.Operations.Version(nil)
	if err != nil {
		return
	}

	return string(*result.Payload.Version), err
}
