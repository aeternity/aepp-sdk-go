# Change Log

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

#  (2019-04-09)


### Bug Fixes

* Ae.WithAccount() was not copying the Account object to the Oracle struct ([ef9eab6](https://github.com/aeternity/aepp-sdk-go/commit/ef9eab6))
* Aens.NameClaimTx helper should pass the plaintext name down to the actual function ([c976216](https://github.com/aeternity/aepp-sdk-go/commit/c976216))
* AENS.UpdateFee specified, made into a BigInt ([de42965](https://github.com/aeternity/aepp-sdk-go/commit/de42965))
* disabled OracleQueryTx, aepp-sdk-js reference RLP serialization is more reliable than aepp-sdk-python atm ([88e7ac4](https://github.com/aeternity/aepp-sdk-go/commit/88e7ac4))
* integration test now works with BigInt's pointer receivers ([d373565](https://github.com/aeternity/aepp-sdk-go/commit/d373565))
* integration test was broken after BigInt.Int was changed from copy-by-value to pointer, because swagger did not know it had to initialize BigInt.Int when creating a BigInt. Fixed. ([b9dd2f3](https://github.com/aeternity/aepp-sdk-go/commit/b9dd2f3))
* NameClaimTx correct serialization in JSON and RLP formats ([0e6799d](https://github.com/aeternity/aepp-sdk-go/commit/0e6799d))
* NameClaimTx salt 256bit bytearray was being converted to uint64 and back, which mangles the bytes. Solved by making it a BigInt, which has proper 32 byte big endian representation with Bytes() ([45ff92d](https://github.com/aeternity/aepp-sdk-go/commit/45ff92d))
* OracleExtendTx.RLP() use t.Fee.Int instead of t.Fee.Bytes() for RLP serialization ([857338f](https://github.com/aeternity/aepp-sdk-go/commit/857338f))
* OracleQueryTx.RLP() was not correct, renamed members of OracleRespondTx ([b679804](https://github.com/aeternity/aepp-sdk-go/commit/b679804))
* rearranged function arguments and correct param types in OracleRegisterTx() ([7c47342](https://github.com/aeternity/aepp-sdk-go/commit/7c47342))
* rearranged function arguments and correct param types in OracleRegisterTx() ([2b00541](https://github.com/aeternity/aepp-sdk-go/commit/2b00541))
* renamed helper function arguments to accomodate different kinds of TTLs in other transaction types ([2970997](https://github.com/aeternity/aepp-sdk-go/commit/2970997))
* renamed helper function arguments to accomodate different kinds of TTLs in other transaction types ([5697537](https://github.com/aeternity/aepp-sdk-go/commit/5697537))
* rlp Encode() was encoding 0 value of uint64 as 0x00, but big.Int 0 value as 0x80. Changed big.Int 0 value to 0x00 ([e21c6f5](https://github.com/aeternity/aepp-sdk-go/commit/e21c6f5))
* SpendTx, OracleRegisterTx RLP() methods should always use utils.BigInt.Int while serializing, not the utils.BigInt directly, because otherwise there will be a list within a list ([fcaa28d](https://github.com/aeternity/aepp-sdk-go/commit/fcaa28d))
* tx command shouldn't use BaseEncodeTx() as a returned function ([21c5a7c](https://github.com/aeternity/aepp-sdk-go/commit/21c5a7c))
* uint64/utils.BigInt instead of int64 for certain variables in config.go ([1c65c9a](https://github.com/aeternity/aepp-sdk-go/commit/1c65c9a))
* uint64/utils.BigInt instead of int64 for certain variables in config.go ([4c76f06](https://github.com/aeternity/aepp-sdk-go/commit/4c76f06))


### Features

* account vanity generator ([3b5c5bd](https://github.com/aeternity/aepp-sdk-go/commit/3b5c5bd))
* AENS transactions in new struct format, given JSON() methods, some types fixed. ([edabb99](https://github.com/aeternity/aepp-sdk-go/commit/edabb99))
* CLI tx dumpraw command (helps with debugging) ([05516cd](https://github.com/aeternity/aepp-sdk-go/commit/05516cd))
* OracleQueryTx struct revision, unittest structure (but no working test) ([41f9821](https://github.com/aeternity/aepp-sdk-go/commit/41f9821))
* OracleQueryTx, OracleRespondTx structs ([be508b1](https://github.com/aeternity/aepp-sdk-go/commit/be508b1))
* OracleRegisterTx, OracleExtendTx. Incomplete unittest ([7c7dbcb](https://github.com/aeternity/aepp-sdk-go/commit/7c7dbcb))
* OracleRegisterTx, OracleExtendTx. Incomplete unittest ([f41a021](https://github.com/aeternity/aepp-sdk-go/commit/f41a021))
* OracleRegisterTxStr() helper function ([a178c87](https://github.com/aeternity/aepp-sdk-go/commit/a178c87))
* OracleRegisterTxStr() helper function ([a654936](https://github.com/aeternity/aepp-sdk-go/commit/a654936))
* Tx structs now feature a JSON() serialization that uses swagger models ([944daa7](https://github.com/aeternity/aepp-sdk-go/commit/944daa7))



# 0.25.0-0.1.0-alpha (2018-11-22)



# 0.22.0-0.1.0-alpha (2018-09-19)



## [1.0.0-alpha]

### Added

- New subcommands `tx verify`, `tx broadcast`, `tx spend`, `account sign`

### Changed

- `account spend`, `account sign` support the `--password` flag
- default tx fee is now 20000, in line with other SDKs

### Removed


### Fixed

- keystore.json reader was not reading kdfparams properly
- rlp from go-ethereum was encoding 0 values differently from Python/Erlang implementations


## [0.25.0-0.1.0-alpha]

### Added

- Add compatibility with epoch v0.25.0

### Removed

- Remove compatibility with epoch < v0.25.0

### Changed 

- Change keystore implementation with xsalsa-poly1305/argon2id


## [0.22.0-0.1.0-alpha]

### Added

- First alpha release of the appe-sdk-go
- Add compatibility with epoch v0.22.0
- Add command line interface supporting inspection of the chain 
- Add command line interface supporting spend transaction 
