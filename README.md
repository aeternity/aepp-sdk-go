# aepp-sdk-go

golang sdk for aeternity blockchain
[![Go Report Card](https://goreportcard.com/badge/github.com/aeternity/aepp-sdk-go)](https://goreportcard.com/report/github.com/aeternity/aepp-sdk-go) [![GoDoc](https://godoc.org/github.com/aeternity/aepp-sdk-go?status.svg)](https://godoc.org/github.com/aeternity/aepp-sdk-go)


## Development

download the latest openapi spcecifications

```
curl  https://sdk-edgenet.aepps.com/api -o api/swagger.json    
```

replace every integer (int64) with a uint64 in the swagger.json (except for time)
generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

```
rm -rf generated/* && swagger generate client -f api/swagger.json -A node  --with-flatten=minimal --target generated  --tags=external --api-package=operations --client-package=client
```

## Structure
No matter what kind of transaction you're making, it always follows the same rules:
1. Find the account nonce, get the transaction TTL (in blocks)
2. Make the transaction
3. Sign the transaction with a given network ID
4. Broadcast it to a node of your choosing

The functions in `aeternity/helpers.go` will help you with these 4 steps, wrapping away the raw implementation details. They use the more barebones functions in `aeternity/transactions.go`.

