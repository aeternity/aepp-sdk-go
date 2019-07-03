package aeternity

import (
	"strings"

	apiclient "github.com/aeternity/aepp-sdk-go/swagguard/node/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Node is the HTTP connection to the aeternity node
type Node struct {
	*apiclient.Node
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

// NewNode obtain a new swagger HTTP client to a aeternity node
func NewNode(nodeURL string, debug bool) *Node {
	// create the transport
	host, schemas := urlComponents(nodeURL)
	transport := httptransport.New(host, "/v2", schemas)
	transport.SetDebug(debug)
	// create the API client, with the transport
	openAPIClient := apiclient.New(transport, strfmt.Default)
	aecli := &Node{
		Node: openAPIClient,
	}
	return aecli
}
