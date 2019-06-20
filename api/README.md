# Node
The node's `swagger.json` cannot be used out of the box.
1. Replace UInt16, UInt32, UInt64 with `"type": "integer", "format": "uint16/32/64"`
CAVEATS
`/generations/height/{height}` has an integer inside a list, cannot be targeted, have to replace by hand
`OffChain*` `allOf` is a list, this script cannot target amounts and fees in there
2. The node replies with a Generic Transaction but specifies type: "SpendTx" instead of "SpendTxJSON", so the stock generic_tx.go does not pick it up.
TODO: investigate why Python and JS SDKs have no problem with this
`python api/generic_tx_json_fix.py generated/models/` to bulk edit all `_tx_json.go` files: their `Type()` should return "*Tx" instead of "*TxJSON"
Manually add `generic_tx.go unmarshalGenericTx()`: `case "ChannelCloseMutualTxJSON": add "ChannelCloseMutualTx"` etc for other Tx types


download the openapi specification
```
curl  https://sdk-testnet.aepps.com/api -o api/swagger.json    
```

generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

```
rm -rf swagguard/node/* && swagger generate client -f api/swagger.json -A node  --with-flatten=minimal --target swagguard/node  --tags=external --api-package=operations --client-package=client
```

# Compiler
```
rm -rf swagguard/compiler/* && swagger generate client -f api/compiler.json -A compiler --with-flatten=minimal --target swagguard/compiler
```