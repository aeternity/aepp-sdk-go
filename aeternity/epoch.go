package aeternity

import (
	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Ae the aeternity client
type Ae struct {
	*apiclient.Epoch
	*Wallet
	*Aens
	*Contract
}

// Wallet high level abstraction for operation on a wallet
type Wallet struct {
	epochCli *apiclient.Epoch
	owner    *Account
}

// Aens abstractions for aens operations
type Aens struct {
	epochCli     *apiclient.Epoch
	owner        *Account
	name         string
	preClaimSalt []byte
}

// Contract abstractions for contracts
type Contract struct {
	epochCli *apiclient.Epoch
	owner    *Account
	source   string
}

// Oracle abstractions for oracles
type Oracle struct {
	epochCli *apiclient.Epoch
	owner    *Account
}

// NewCli obtain a new epochCli instance
func NewCli(epochURL string, debug bool) *Ae {
	// create the transport
	host, schemas := urlComponents(epochURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClinet := apiclient.New(transport, strfmt.Default)
	aecli := &Ae{
		Epoch: openAPIClinet,
		Wallet: &Wallet{
			epochCli: openAPIClinet,
		},
		Aens: &Aens{
			epochCli: openAPIClinet,
		},
		Contract: &Contract{
			epochCli: openAPIClinet,
		},
	}
	return aecli
}

// NewCliW obtain a new epochCli instance
func NewCliW(epochURL string, kp *Account, debug bool) *Ae {
	// create the transport
	host, schemas := urlComponents(epochURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClinet := apiclient.New(transport, strfmt.Default)
	aecli := &Ae{
		Epoch: openAPIClinet,
		Wallet: &Wallet{
			epochCli: openAPIClinet,
			owner:    kp,
		},
		Aens: &Aens{
			epochCli: openAPIClinet,
			owner:    kp,
		},
		Contract: &Contract{
			epochCli: openAPIClinet,
			owner:    kp,
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
