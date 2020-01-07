# Design of aepp-sdk-go
This document should give you all the information and context you need to know to start modifying the code with confidence.

# Structure
Let's start from the bottom upwards. 

`api/` is where the Swagger files for the node and compiler stored. Some changes need to be made - there are scripts for some tasks, but not everything can/should be automated. The details are in `api/README.md`.

`swagguard/` is the generated Go code from the modified swagger file. Unfortunately, this generated code needs to be modified - the scripts are in `api/`. Tests to ensure the changes work as required are under `swagguard/`. HTTP communication with the node/compiler and the JSON models are all implemented here.  Due to an old design decision, code to talk with the node's debug HTTP interface is not generated. It is mostly not needed, except for dry-running contract transactions (which is a good use case). Swagger generated code is not so trustworthy or nice to deal with, therefore it is kept here and abstracted away as much as is convenient with `naet/`.

`utils/` used to hold more, but now only holds the `BigInt` type required for swagger generated code to work. Use Go's `big.Int` whenever possible and cast it to `utils.BigInt` only when you have to.

`naet/` contains `Node`, `Compiler`. Since using the swagger generated code under `swagguard/` can be cumbersome, this layer exists to make calling the HTTP endpoints of the node/compiler as Go-ish as possible. Some of `Node`'s endpoints convert the swagger model types into Go native types, which makes it easier to use. `naet/` is a convenience layer on top of `swagguard/`, and not all available HTTP endpoints are covered so far (but most are not needed). Anything that is not in `swagguard/` should not be in `naet/`.

`transactions/` continues abstracting away the models in the Swagger generated code, but only for transactions. Expressing transactions as Go structs instead of using the Swagger generated models was a very rewarding design decision. 

`models/` some models that should be abstracted from swagger generated code do not fit in `transactions/` and thus live here.

`binary/` is a simple package that holds hashing functions and constants.

`account/` is a simple package that holds account functions and constants.

`config/` holds configuration variables.

`cmd/` is a CLI client, sporting similar functionality to those of the Python/JS SDKs. It is not meant to be a convenient CLI tool for tasks but rather as an offline transaction creation tool. This forced design decisions downwards, such as the use of closures in `transactions/` for `TTLNoncer`. It hasn't been paid much attention lately.

`aeternity/` is a convenience layer that abstracts away transaction creation, signing, sending, waiting for it to be mined and exposes a simple task-like interface for people who want to use Aeternity's features. Finding the right design for this was not easy, and `aeternity.Contract` in particular could have been designed better. However, `aeternity.Oracle` should be production ready.

# Testing
## Unit tests
`go test $(go list ./... |grep -v integration_test)`

## Integration tests
```
# ensure .env has the node/compiler versions you want to test
docker-compose up node compiler

export AETERNITY_ALICE_PRIVATE_KEY=.....
export AETERNITY_BOB_PRIVATE_KEY=.....
go test ./integration_test -count=1 # count is needed because Go caches test results if they were successful.
```