package aeternity

import (
	"strings"

	apiclient "github.com/aeternity/aepp-sdk-go/swagguard/node/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client is the HTTP connection to the aeternity node
type Client struct {
	*apiclient.Node
}

// Wallet is a account-specific helper that stores state relevant to spending operations
type Wallet struct {
	Client  *Client
	Account *Account
}

// Aens ais a account-specific helper that stores state relevant to AENS operations
type Aens struct {
	Client       *Client
	Account      *Account
	name         string
	preClaimSalt []byte
}

// Contract is a account-specific helper that stores state relevant to smtart contract execution
type Contract struct {
	Client   *Client
	Compiler *Compiler
	Owner    string
}

// Oracle is a account-specific helper that stores state relevant to oracles
type Oracle struct {
	Client  *Client
	Account *Account
}

func urlComponents(url string) (host string, schemas []string) {
	p := strings.Split(url, "://")
	if len(p) == 1 {
		host = p[0]
		schemas = []string{"http"}
		return
	}
	host = p[1]
	schemas = []string{p[0]}
	return
}

// NewClient obtain a new nodeClient instance
func NewClient(nodeURL string, debug bool) *Client {
	// create the transport
	host, schemas := urlComponents(nodeURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClient := apiclient.New(transport, strfmt.Default)
	aecli := &Client{
		Node: openAPIClient,
	}
	return aecli
}
