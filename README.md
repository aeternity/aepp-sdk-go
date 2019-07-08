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
aeNode := aeternity.NewNode("http://localhost:3013", false).WithAccount(acc)
```

Most parameters are set by modifying the variables in `config.go` in this manner:
`aeternity.Config.Client.Fee = *utils.RequireBigIntFromString("100000000000000")`

When using the `Context` struct helper functions in `helpers.go`, chores like getting the TTL, Account Nonce, encoding of the AENS claim etc are done automatically.
```
import "github.com/aeternity/aepp-sdk-go/aeternity"

...

// create a Context for the address you're going to sign the transaction
// with, and an aeternity node to talk to/query the address's nonce.
ctx := aeternity.NewContext(node, alice.Address)

// create the SpendTransaction
tx, err := ctx.SpendTx(alice.Address, bob.Address, *amount, *fee, msg)
if err != nil {
    t.Error(err)
}

// optional: minimize the fee to save money!
est, _ := tx.FeeEstimate()
fmt.Println("Estimated vs Actual Fee:", est, tx.Fee)
tx.Fee = *est

// transform the tx into a tx_base64encodedstring
txB64, err := aeternity.BaseEncodeTx(tx)
if err != nil {
    t.Error(err)
}

signedTxStr, hash, _, err := aeternity.SignEncodeTxStr(acc, txB64, aeternity.Config.Node.NetworkID)
if err != nil {
    t.Fatal(err)
}

err = aeternity.BroadcastTransaction(node, signedTxStr)
if err != nil {
    t.Fatal(err)
}
```