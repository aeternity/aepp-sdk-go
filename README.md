# aepp-sdk-go

golang sdk for aeternity blockchain
[![Go Report Card](https://goreportcard.com/badge/github.com/aeternity/aepp-sdk-go)](https://goreportcard.com/report/github.com/aeternity/aepp-sdk-go) [![GoDoc](https://godoc.org/github.com/aeternity/aepp-sdk-go?status.svg)](https://godoc.org/github.com/aeternity/aepp-sdk-go)

## Setup
If your project uses Go Modules (go.mod, go.sum files), you must include the major version in the import line like this:
`import github.com/aepp-sdk-go/v5/aeternity`

If your project won't use Go Modules (no go.mod, go.sum files), ensure your $GOPATH/src/github.com/aeternity/aepp-sdk-go is on the correct branch. Then your import should be:
`import github.com/aepp-sdk-go/aeternity`

## Usage
No matter what kind of transaction you're making, it always follows the same rules:
1. Find the account nonce, get the transaction TTL (in blocks)
2. Make the transaction
3. Sign the transaction with a given network ID
4. Broadcast it to a node of your choosing

See `aeternity/context_test.go` or use godoc for code examples.