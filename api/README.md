# Node
The node's `swagger.json` cannot be used out of the box.
1. replace `"$ref": "#/definitions/UInt64"` with `"type": "integer": "format": "uint64"`
2. replace `"$ref": "#/definitions/EncodedPubkey/Hash/Value/ByteArray"` with `"type": "string"`
3. replace `"$ref": "#/definitions/TxBlockHeight"` with `"$ref": "#/definitions/UInt"`
4. Inconsistency between `OracleResponse/OracleRespond`
5. updatedict.py adds BigInt to `/definitions/UInt`, makes all implicit `int64`s explicit `uint64`s
6. generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

    ```
    rm -rf swagguard/node/* && swagger generate client -f api/swagger.json -A node  --with-flatten=minimal --target swagguard/node  --tags=external --api-package=operations --client-package=client
    ```

7. The node replies with a Generic Transaction but specifies type: "SpendTx" instead of "SpendTxJSON", so the stock generic_tx.go does not pick it up.
TODO: investigate why Python and JS SDKs have no problem with this
`python api/generic_tx_json_fix.py swagguard/node/models/` to bulk edit all `_tx_json.go` files: their `Type()` should return "*Tx" instead of "*TxJSON"
Manually add `generic_tx.go unmarshalGenericTx()`: `case "ChannelCloseMutualTxJSON": add "ChannelCloseMutualTx"` etc for other Tx types

    generic_tx_json_fix.py fixes the `*TxJSON` problem partially - you still need to edit generic_tx.go

8. remember to add .String() to Error


# Compiler
```
rm -rf swagguard/compiler/* && swagger generate client -f api/compiler.json -A compiler --with-flatten=minimal --target swagguard/compiler
```