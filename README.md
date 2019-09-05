# aepp-sdk-go

golang sdk for aeternity blockchain
[![Go Report Card](https://goreportcard.com/badge/github.com/aeternity/aepp-sdk-go)](https://goreportcard.com/report/github.com/aeternity/aepp-sdk-go) [![GoDoc](https://godoc.org/github.com/aeternity/aepp-sdk-go?status.svg)](https://godoc.org/github.com/aeternity/aepp-sdk-go)


## Usage
No matter what kind of transaction you're making, it always follows the same rules:
1. Find the account nonce, get the transaction TTL (in blocks)
2. Make the transaction
3. Sign the transaction with a given network ID
4. Broadcast it to a node of your choosing

```
acc, err := aeternity.AccountFromHexString(senderPrivateKey)
if err != nil {
    fmt.Println(err)
    return
}
node := aeternity.NewNode("http://localhost:3013", false).WithAccount(acc)
```

Most parameters are set by modifying the variables in `config.go` in this manner:
`aeternity.Config.Client.Fee = big.NewInt(100000000000000)`

When using the `Helper` methods in `helpers.go`, chores like getting the TTL, Account Nonce are done automatically.
When using the `Context` methods in `helpers.go`, additional conveniences for AENS like Commitment ID calculation, Namehashing, are done for you.
For a painless experience when building transactions, use the `Context` methods.
```
import "github.com/aeternity/aepp-sdk-go/aeternity"

...

// create a Context for the address you're going to sign the transaction
// with, and an aeternity node to talk to/query the address's nonce.
ctx := aeternity.NewContextFromURL(node, alice.Address, false)

// create the SpendTransaction
tx, err := ctx.SpendTx(alice.Address, bob.Address, *amount, *fee, msg)
if err != nil {
    t.Error(err)
}

// optional: minimize the fee to save money!
est, _ := tx.FeeEstimate()
fmt.Println("Estimated vs Actual Fee:", est, tx.Fee)
tx.Fee = *est

signedTx, hash, signature, err := aeternity.SignHashTx(acc, tx, aeternity.Config.Node.NetworkID)
if err != nil {
    t.Fatal(err)
}

// transform the tx into a tx_base64encodedstring so you can HTTP POST it
signedTxStr, err := aeternity.SerializeTx(signedTx)
if err != nil {
    t.Error(err)
}

err = aeternity.BroadcastTransaction(node, signedTxStr)
if err != nil {
    t.Fatal(err)
}
```