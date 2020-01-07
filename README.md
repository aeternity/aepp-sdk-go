# aepp-sdk-go
[![Go Report Card](https://goreportcard.com/badge/github.com/aeternity/aepp-sdk-go)](https://goreportcard.com/report/github.com/aeternity/aepp-sdk-go) [![GoDoc](https://godoc.org/github.com/aeternity/aepp-sdk-go?status.svg)](https://godoc.org/github.com/aeternity/aepp-sdk-go)

golang sdk for aeternity blockchain

## Setup
`go get github.com/aeternity/aepp-sdk-go`

If your project uses Go Modules (go.mod, go.sum files), you must include the major version in the import line like this:
`import github.com/aepp-sdk-go/v8/aeternity`

If your project won't use Go Modules (no go.mod, go.sum files), ensure your `$GOPATH/src/github.com/aeternity/aepp-sdk-go` is on the correct branch. Then your import should be:
`import github.com/aepp-sdk-go/aeternity`

## Contextual Knowledge
Every transaction submitted to a node needs a nonce (to ensure its uniqueness), a TTL (how long, in blocks, should the transaction stay in the mempool). Signing a transaction includes the `NetworkID` as well, so a transaction meant for `ae_uat` (testnet) won't make it onto `ae_mainnet` (mainnet). The SDK communicates with the node and/or Sophia compiler over a HTTP REST API to find the current account nonce/current height/broadcast the transaction etc.

In short, creating a transaction always follows this pattern:
1. Find the account nonce, get the transaction TTL (in blocks)
2. Make the transaction
3. Sign the transaction with a given network ID
4. Broadcast it to a node of your choosing

## Where to find examples
All examples are in godoc.org, except the integration tests.
General workflow code examples are in `package aeternity`, or check out the integration tests in `package integration_test`
Account, HD wallet management in `package account`
etc.

## Where to ask for help
[aeternity forum](https://forum.aeternity.com/c/aepplications/sdk)
[rocketchat](https://devchat.aeternity.com/)