# aepp-sdk-go

golang sdk for aeternity blockchain
[![Go Report Card](https://goreportcard.com/badge/github.com/aeternity/aepp-sdk-go)](https://goreportcard.com/report/github.com/aeternity/aepp-sdk-go) [![GoDoc](https://godoc.org/github.com/aeternity/aepp-sdk-go?status.svg)](https://godoc.org/github.com/aeternity/aepp-sdk-go)


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