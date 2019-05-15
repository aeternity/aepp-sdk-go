package aeternity

import (
	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client the aeternity client
type Client struct {
	*apiclient.Node
	*Wallet
	*Aens
	*Contract
	*Oracle
}

// Wallet high level abstraction for operation on a wallet
type Wallet struct {
	Client *apiclient.Node
	owner  *Account
}

// Aens abstractions for aens operations
type Aens struct {
	Client       *apiclient.Node
	owner        *Account
	name         string
	preClaimSalt []byte
}

// Contract abstractions for contracts
type Contract struct {
	Client *apiclient.Node
	owner  *Account
	source string
}

// Oracle abstractions for oracles
type Oracle struct {
	Client *apiclient.Node
	owner  *Account
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
		Wallet: &Wallet{
			Client: openAPIClient,
		},
		Aens: &Aens{
			Client: openAPIClient,
		},
		Contract: &Contract{
			Client: openAPIClient,
		},
		Oracle: &Oracle{
			Client: openAPIClient,
		},
	}
	return aecli
}

// WithAccount associate a Account with the client
func (ae *Client) WithAccount(account *Account) *Client {
	ae.Wallet.owner = account
	ae.Aens.owner = account
	ae.Contract.owner = account
	ae.Oracle.owner = account
	return ae
}
