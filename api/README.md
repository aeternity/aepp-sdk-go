# Node
The node's `node.json` cannot be used out of the box.
1. `python3 updatedict.py node.json node-temp.json`
   - replaces `"$ref": "#/definitions/UInt64/32/16"` with `"type": "integer", "format": "uint64/32/16"`
   - replaces `"$ref": "#/definitions/EncodedPubkey/Hash/Value/ByteArray"` with `"type": "string"`
   - replaces `#/definitions/TxBlockHeight` with `#/definitions/UInt`
   - replaces `OracleResponseTxJSON` with `OracleRespondTxJSON` (https://github.com/aeternity/aeternity/issues/2799)
   - adds BigInt to `/definitions/UInt`
   - makes all implicit `int64`s explicit `uint64`s
   - remove maximum for uint64 because go-swagger converts it to float
2. generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))
    ```
    rm -rf ../swagguard/node/* && swagger generate client -f node-temp.json -A node  --with-flatten=minimal --target ../swagguard/node  --tags=external --api-package=operations --client-package=client
    ```

3. The node replies with a Generic Transaction but specifies type: "SpendTx" instead of "SpendTxJSON", so the stock generic_tx.go does not pick it up.

TODO: investigate why Python SDK have no problem with this

`python3 generic_tx_json_fix.py ../swagguard/node/models/` to bulk edit all `_tx_json.go` files: their `Type()` should return "*Tx" instead of "*TxJSON"

4. Manually add `generic_tx.go unmarshalGenericTx()`: `case "ChannelCloseMutualTxJSON": add "ChannelCloseMutualTx"` etc for other Tx types
   `generic_tx_json_fix.py` fixes the `*TxJSON` problem partially - you still need to edit generic_tx.go

5. remember to add .String() to Error (https://github.com/go-swagger/go-swagger/issues/872)


# Compiler
```
rm -rf ../swagguard/compiler/* && swagger generate client -f compiler.json -A compiler --with-flatten=minimal --target ../swagguard/compiler
```


# go-swagger
The last time was used:
```
$ swagger version
version: v0.27.0
commit: 43c2774170504d87b104e3e4d68626aac2cd447d
```
