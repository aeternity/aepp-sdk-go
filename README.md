# aepp-sdk-go

golang sdk for aeternity blockchain


## Development

download the latest openapi spcecifications

```
curl  https://sdk-edgenet.aepps.com/api -o api/swagger.json    
```

generate the client (using [go-swagger](https://github.com/go-swagger/go-swagger))

```
rm -rf generated/*
swagger generate client -f api/swagger.json -A epoch  --with-flatten=minimal --target generated  --tags=external --api-package=operations --client-package=client
```