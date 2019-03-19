package aeternity

import (
	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Ae the aeternity client
type Ae struct {
	*apiclient.Node
	*Wallet
	*Aens
	*Contract
	*Oracle
}

// Wallet high level abstraction for operation on a wallet
type Wallet struct {
	nodeCli *apiclient.Node
	owner   *Account
}

// Aens abstractions for aens operations
type Aens struct {
	nodeCli      *apiclient.Node
	owner        *Account
	name         string
	preClaimSalt []byte
}

// Contract abstractions for contracts
type Contract struct {
	nodeCli *apiclient.Node
	owner   *Account
	source  string
}

// Oracle abstractions for oracles
type Oracle struct {
	nodeCli *apiclient.Node
	owner   *Account
}

// NewCli obtain a new nodeCli instance
func NewCli(nodeURL string, debug bool) *Ae {
	// create the transport
	host, schemas := urlComponents(nodeURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClient := apiclient.New(transport, strfmt.Default)
	aecli := &Ae{
		Node: openAPIClient,
		Wallet: &Wallet{
			nodeCli: openAPIClient,
		},
		Aens: &Aens{
			nodeCli: openAPIClient,
		},
		Contract: &Contract{
			nodeCli: openAPIClient,
		},
		Oracle: &Oracle{
			nodeCli: openAPIClient,
		},
	}
	return aecli
}

// WithAccount associate a Account with the client
func (ae *Ae) WithAccount(account *Account) *Ae {
	ae.Wallet.owner = account
	ae.Aens.owner = account
	ae.Contract.owner = account
	return ae
}
