# Updating the swagger.json for newer node versions
The node's `swagger.json` cannot be used out of the box.
1. Manually replace all int64s with uint64s in swagger.json except for time (because go's `time.Unix()` accepts `int64`, not `uint64`)
2. `python3 updatedict.py` adds `Fee/Balance/Amount/NameSalt BigInt` in definitions and replaces all fees and amounts with references to these definitions. Will also make sure there are no implicit int64s left ("type": "integer" without a "format" key)
CAVEATS
`/generations/height/{height}` has an integer inside a list, cannot be targeted, have to replace by hand
`OffChain*` `allOf` is a list, this script cannot target amounts and fees in there
3. The node replies with a Generic Transaction but specifies type: "SpendTx" instead of "SpendTxJSON", so the stock generic_tx.go does not pick it up.
TODO: investigate why Python and JS SDKs have no problem with this
`python api/generic_tx_json_fix.py generated/models/` to bulk edit all `_tx_json.go` files: their `Type()` should return "*Tx" instead of "*TxJSON"
Manually add `generic_tx.go unmarshalGenericTx()`: `case "ChannelCloseMutualTxJSON": add "ChannelCloseMutualTx"` etc for other Tx types


download the openapi specification
```
curl  https://sdk-testnet.aepps.com/api -o api/swagger.json    
```

generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

```
rm -rf generated/* && swagger generate client -f api/swagger.json -A node  --with-flatten=minimal --target generated  --tags=external --api-package=operations --client-package=client
```
