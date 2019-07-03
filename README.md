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
aeNode := aeternity.NewCli("http://localhost:3013", false).WithAccount(acc)
```

Most parameters are set by modifying the variables in `config.go` in this manner:
`aeternity.Config.Client.Fee = *utils.RequireBigIntFromString("100000000000000")`

When using the `Ae/Aens/Contract/Oracle` struct helper functions in `helpers.go`, chores like getting the TTL, Account Nonce, encoding of the AENS claim etc are done automatically.
```
preclaimTx, salt, err := aeNode.Aens.NamePreclaimTx("fdsa.test", aeternity.Config.Client.Fee)
if err != nil {
    fmt.Println(err)
    return
}
preclaimTxStr, err := aeternity.BaseEncodeTx(&preclaimTx)

signedTxStr, hash, signature, err := aeternity.SignEncodeTxStr(acc, preclaimTxStr, "ae_docker")
if err != nil {
    fmt.Println(err)
    return
}

err = aeNode.BroadcastTransaction(signedTxStr)
if err != nil {
    panic(err)
}
```