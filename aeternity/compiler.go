package aeternity

import (
	compiler_client "github.com/aeternity/aepp-sdk-go/swagguard/compiler/client"
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
	cClient := compiler_client.New(transport, strfmt.Default)
	compiler := &Compiler{
		Compiler: cClient,
	}
	return compiler
}

// GetAPIVersion connects to the compiler and returns its version string, e.g.
// 3.1.0
func (c *Compiler) GetAPIVersion() (version string, err error) {
	result, err := c.Compiler.Operations.APIVersion(nil)
	if err != nil {
		return "", err
	}
	version = *result.Payload.APIVersion
	return
}
