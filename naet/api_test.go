/*
For best coverage, this needs to be run against a local node, which needs to have
the transactions submitted to it first. Which makes it an integration test.

Mainnet/testnet doesn't have all transactions, and even if they did, oracles/names expire,
so you can't test all the endpoints anyway. Best to do this in a controlled environment.

Therefore this is an integration test, to live in integration_test/api_test.go
*/
package naet
