# aepp-sdk-go

golang sdk for aeternity blockchain


## Development

download the latest openapi spcecifications

```
curl  https://sdk-testnet.aepps.com/api -o epoch-0.18.0.json     
```

generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

```
rm -rf generated/*
swagger generate client -f api/swagger.json -A epoch  --with-flatten=minimal --target generated 
```