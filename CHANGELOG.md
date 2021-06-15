
<a name="v9.0.0"></a>
## [v9.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v8.0.0...v9.0.0) (2021-06-15)

### Bug Fixes

* define items field to fix generation error ([2c42678](https://github.com/aeternity/aepp-sdk-go/commit/2c42678cbfa9f2090d83708137f788f0b3ddd186))
* define items field to fix generation error ([c688250](https://github.com/aeternity/aepp-sdk-go/commit/c68825043940cd5f05701962a3b392a87c9c90dd))
* revert serialisation to string of compiler models ([bfbb64a](https://github.com/aeternity/aepp-sdk-go/commit/bfbb64af006bea341cb243e5adf9bbb664fd6533))
* remove incompatible definitions ([1f4d7b9](https://github.com/aeternity/aepp-sdk-go/commit/1f4d7b9908b8fcb997fae3043ae5744a8a4afa1e))
* remove incompatible definitions ([fd9ff9b](https://github.com/aeternity/aepp-sdk-go/commit/fd9ff9be9aad0e5f5786be3fc9efc06a1be5c42b))
* don't fail on getting nonce for new account ([8ef7545](https://github.com/aeternity/aepp-sdk-go/commit/8ef75452876d294455cb3718ad7c70e39846d6f3))
* don't panic on Invalid tx ([06b545a](https://github.com/aeternity/aepp-sdk-go/commit/06b545aa5466d11f87a83c95af94f7855d5649ad))
* **big-int:** add missed ContextValidate method ([31b0075](https://github.com/aeternity/aepp-sdk-go/commit/31b00759c72f6266df81e52f995bf2586424ad46))
* **node-error:** revert removed String method ([90ae4b4](https://github.com/aeternity/aepp-sdk-go/commit/90ae4b4e29d6ed09a2e75ca7ef3f5778f4d84d63))
* **node-error:** revert removed String method ([2ed984c](https://github.com/aeternity/aepp-sdk-go/commit/2ed984c3dec106b46f7623e8010219a0f5f74660))
* **node-error:** revert removed String method ([ae6f6df](https://github.com/aeternity/aepp-sdk-go/commit/ae6f6dfffa7c08b4534421efacb0d20cc3b98e3c))

### Chore

* regenerate changelog ([d416701](https://github.com/aeternity/aepp-sdk-go/commit/d416701ec67d37fb1cdfead4f9bc41bdccdd9342))
* remove unused bytecode ([25e16bf](https://github.com/aeternity/aepp-sdk-go/commit/25e16bf2765ca01838e190a81d49760c8eb1b90c))
* setup changelog generator ([113e1a4](https://github.com/aeternity/aepp-sdk-go/commit/113e1a431d5ea28995ad53f220fd653c9fc62d47))
* bump version to 9 ([2464486](https://github.com/aeternity/aepp-sdk-go/commit/246448630ce6310ec473d912b8eb3e57c55c547a))
* remove extra dependencies using `go mod tidy` ([03de926](https://github.com/aeternity/aepp-sdk-go/commit/03de926eda007b7771da98f4ee779b75fde9d9ea))
* enable Iris protocol ([26f720c](https://github.com/aeternity/aepp-sdk-go/commit/26f720c2afd67676ae8881bd3bbc16c649ea947a))
* update compiler api to 6.0.0 ([1fa699c](https://github.com/aeternity/aepp-sdk-go/commit/1fa699cda050152f675589e8429e43de1e5bf375))
* manually fix generic_tx.go according to readme ([62fb47d](https://github.com/aeternity/aepp-sdk-go/commit/62fb47d6fc3b65e893cb3417781c656ec4302d2a))
* run generic_tx_json_fix.py ([b09ee44](https://github.com/aeternity/aepp-sdk-go/commit/b09ee441b81ddb4ff015a008cd51423f4c4364d5))
* regenerate node api ([e4e5daf](https://github.com/aeternity/aepp-sdk-go/commit/e4e5daf8d6fb75365b9a1ea2ab1abf88c2596777))
* manually fix generic_tx.go according to readme ([bfc7047](https://github.com/aeternity/aepp-sdk-go/commit/bfc7047f70f02371b44e173a74fa0aad5c98ce5a))
* update node api to 6.0.0 ([e7984a9](https://github.com/aeternity/aepp-sdk-go/commit/e7984a9c7b5adbc42f22e5d599e7c69f7bf9ce46))
* run generic_tx_json_fix.py ([580ce75](https://github.com/aeternity/aepp-sdk-go/commit/580ce75c0596cee8c1c3bc4d194594bc267c4452))
* regenerate node api ([ae79858](https://github.com/aeternity/aepp-sdk-go/commit/ae79858595cd54218a6b43b4d42fbe12b47f9c71))
* update node api to 5.11.0 ([9af9723](https://github.com/aeternity/aepp-sdk-go/commit/9af9723bc698c8a48129e8282ed678b1ee936b14))
* fix linter errors manually ([2257666](https://github.com/aeternity/aepp-sdk-go/commit/2257666bd16a3ceefe45d8ee2aed2546e87e1da9))
* manually fix generic_tx.go according to readme ([01126e1](https://github.com/aeternity/aepp-sdk-go/commit/01126e129a16914a49b4fa11a1a99ae1da9fcebe))
* run generic_tx_json_fix.py ([e4fa90f](https://github.com/aeternity/aepp-sdk-go/commit/e4fa90f4782d5afdea62acaabb82b26e00adabe0))
* regenerate node api using latest go-swagger ([4cdc679](https://github.com/aeternity/aepp-sdk-go/commit/4cdc679300a456bb48b15b6dd74720c21363aac6))
* process node.js ([9cae02a](https://github.com/aeternity/aepp-sdk-go/commit/9cae02a2ef33343947f03459a1058ed870c524bd))
* update api readme ([037ddbd](https://github.com/aeternity/aepp-sdk-go/commit/037ddbddc4824d384ef866453300628f451e6d7e))
* regenerate compiler api ([9624325](https://github.com/aeternity/aepp-sdk-go/commit/9624325568b28f6cbcb6533edc8374660216031d))
* update compiler api to 4.3.2 ([2ca70fb](https://github.com/aeternity/aepp-sdk-go/commit/2ca70fbc9e8181e530aba650ff2d92eb89adcd54))
* reformat compiler api with jq ([3150369](https://github.com/aeternity/aepp-sdk-go/commit/3150369b75ec4d2ba411e3cde2f1c70e8c021f50))
* regenerate compiler api using latest go-swagger ([f6713eb](https://github.com/aeternity/aepp-sdk-go/commit/f6713eb870e8b073cf6ee97e4ce987096806de16))
* switch to a not deprecated linter package ([4e223f8](https://github.com/aeternity/aepp-sdk-go/commit/4e223f8384700c9c38492a15c2646806efc790d0))
* simplify node configuration ([315f5ae](https://github.com/aeternity/aepp-sdk-go/commit/315f5aecd25fc224f138ba11a98bd9549f6352a8))
* run tests at GitHub Actions ([bbce116](https://github.com/aeternity/aepp-sdk-go/commit/bbce116ae190b8fe4665225ebe38b2dcfc1b4197))
* **changelog:** rename sections and add links ([0aa5091](https://github.com/aeternity/aepp-sdk-go/commit/0aa5091145253adad56489964428e1c040961024))

### Ci

* simplify configuration ([672d3ba](https://github.com/aeternity/aepp-sdk-go/commit/672d3ba5d7fb47a326e194d27548f6b582f3124b))

### Code Refactoring

* drop old nodes support and compiler switch ([f8ab7a5](https://github.com/aeternity/aepp-sdk-go/commit/f8ab7a5bd5eabe9033b62f57399a72b310cbefb6))
* fix linter errors manually ([52a8066](https://github.com/aeternity/aepp-sdk-go/commit/52a8066c8304f1222c9eb322e7a52c1132abe11b))
* fix linter errors manually ([f78ad17](https://github.com/aeternity/aepp-sdk-go/commit/f78ad177e3401345d69af43b6013d1b23da88c5c))
* simplify Makefile ([daef543](https://github.com/aeternity/aepp-sdk-go/commit/daef54318c6c5ef0c5ce3b6004055da429a6837b))
* remove Jenkins-related report package ([815a94f](https://github.com/aeternity/aepp-sdk-go/commit/815a94f40ab298e329b25698882e94adf97f076d))
* update go version ([e548512](https://github.com/aeternity/aepp-sdk-go/commit/e548512a2d6344327ab71486ebe23d6522aa3589))
* update dependencies ([5e71e07](https://github.com/aeternity/aepp-sdk-go/commit/5e71e07d94e7d0c73d5b23fe7be14b6f7658fd70))
* depend on organisation fork of rlp for aeternity ([8f661f0](https://github.com/aeternity/aepp-sdk-go/commit/8f661f014c6f18935176c638e1e77a63a70e436d))
* update node URLs ([a2648aa](https://github.com/aeternity/aepp-sdk-go/commit/a2648aafd76889c7f3defb5891297f1fa92ca089))

### Docs

* add go-swagger version ([ab5ab0d](https://github.com/aeternity/aepp-sdk-go/commit/ab5ab0dd077db5fe3036bbdf3f6fbed41b02af53))

### Test

* fix oracle ([0ff2f08](https://github.com/aeternity/aepp-sdk-go/commit/0ff2f08570bac6692038b3f0a510a399ac32f23c))
* generate random name to don't cleanup node each run ([078f642](https://github.com/aeternity/aepp-sdk-go/commit/078f6422f9b7f80b4b1c233366ea7028ae2a11c2))
* fix GA ([951b02e](https://github.com/aeternity/aepp-sdk-go/commit/951b02e65bdba73e2166c43c49859ff1a40205a3))
* update bob address because of lack of private key ([2967e12](https://github.com/aeternity/aepp-sdk-go/commit/2967e12747a664193ee575c587eec378bab378e9))
* don't connect to incompatible public testnet ([d56bb06](https://github.com/aeternity/aepp-sdk-go/commit/d56bb0653c7281e698050cb52c9203c6f07e23f9))


<a name="v8.0.0"></a>
## [v8.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v7.0.1...v8.0.0) (2020-01-07)

### Bug Fixes

* DefaultOracleListener never returns an error ([35d42a8](https://github.com/aeternity/aepp-sdk-go/commit/35d42a8ce85555454bd70f210edd9191e2037574))
* aens_test.go uses the TxReceipt's features ([8cc0af8](https://github.com/aeternity/aepp-sdk-go/commit/8cc0af8807012e5753cca27f94de2707b08019f5))
* node.Infoer functionality moved to Context.NodeInfo() ([97453a4](https://github.com/aeternity/aepp-sdk-go/commit/97453a4cf6babb654e5d995a7af60274e85eab32))
* avoid naet.NodeInterface whenever possible when writing functions that use naet.Node. Each function that uses the node should specify explicitly its methods that it uses from naet.Node, perhaps by defining its own interface locally. ([245f5e3](https://github.com/aeternity/aepp-sdk-go/commit/245f5e30f80807df452266e0989fa9302ee8a804))
* alternate way of getting a bytes version of Uint64 using binary.BigEndian.PutUint64() ([e403a29](https://github.com/aeternity/aepp-sdk-go/commit/e403a296720d899b1b809515419e201ae3766a48))

### Chore

* bump version to 8 ([42123be](https://github.com/aeternity/aepp-sdk-go/commit/42123be29e8a7efa11434e622c7ec676291df9eb))
* fix linter complaints ([c459dc6](https://github.com/aeternity/aepp-sdk-go/commit/c459dc64bab8e13417bae47406281ed8fe459371))
* comments and exported CompileEncoder interface from aeternity ([dd4e371](https://github.com/aeternity/aepp-sdk-go/commit/dd4e371013f3f42b5be2b7d605e71dbd00a65129))
* .env v5.2.0 ([9cc1a20](https://github.com/aeternity/aepp-sdk-go/commit/9cc1a20342bd2256b4a4d482e4544f8247f609d4))
* rename INTEGRATION_TEST_SENDER_PRIVATE_KEY to AETERNITY_ALICE_PRIVATE_KEY ([07d4a51](https://github.com/aeternity/aepp-sdk-go/commit/07d4a51889568ca0dc127aa98cb159f268ab9dd0))

### Code Refactoring

* Context.SignBroadcast as well as SignBroadcastWait ([6260ed9](https://github.com/aeternity/aepp-sdk-go/commit/6260ed9afb75c54604118b777b4900765721084c))
* integrationtest NewContext no longer returns an err ([12c2047](https://github.com/aeternity/aepp-sdk-go/commit/12c2047ee447bdb489a84768d1ba8f2a3ece6de5))
* NewContext doesn't need to return an error after all; TxReceipt.Tx should not be a pointer to an interface ([1dbe992](https://github.com/aeternity/aepp-sdk-go/commit/1dbe992ef2894f52eb07bbeb82e8e3fbb6ede4a5))
* getTransactionByHashHeighter interface renamed to transactionWaiter ([30e6487](https://github.com/aeternity/aepp-sdk-go/commit/30e648758e2cca507da206045f2dd90427ca9334))
* NewContext does not need to return error after all. ErrWaitTransaction type was too complicated for its usefulness ([a6b5cb1](https://github.com/aeternity/aepp-sdk-go/commit/a6b5cb110d701306cc8d68f57198f9fece155371))
* remove unused WaitTransactionForXBlocks() ([035fcb5](https://github.com/aeternity/aepp-sdk-go/commit/035fcb5a985bdd4bb7bbc7385f345307f626a994))
* use aeternity.SignBroadcast(), aeternity.WaitSynchronous() or ctx.SignBroadcastWait() throughout the codebase ([74f46d9](https://github.com/aeternity/aepp-sdk-go/commit/74f46d99d09c56d96ffca5ae19782f1df6df2372))
* aeternity.WaitSynchronous() ([13b8b55](https://github.com/aeternity/aepp-sdk-go/commit/13b8b550a9c79508bf826a86951d89d194b3e1c9))
* it is not yet time to remove broadcastWaitTransactionNodeCapabilities ([e94898a](https://github.com/aeternity/aepp-sdk-go/commit/e94898a9afce51ef8ffc350cd2211ccf76d6bd95))
* introduced TxReceipt.Watch(),  which replaces SignBroadcastWait. but not sure yet how to use it with existing code ([d56cecb](https://github.com/aeternity/aepp-sdk-go/commit/d56cecb241eb92a52c9b3ebbb57cd4059415a360))
* Oraclize can now Listen and Register if Oracle was not previously registered ([f20236e](https://github.com/aeternity/aepp-sdk-go/commit/f20236e01d3dcc7a58b7f835d0c5297c7e0d1bda))
* make ChainPollInterval useful, 1s by default ([04136e9](https://github.com/aeternity/aepp-sdk-go/commit/04136e98987e7d76c38c08c125e50f4c5c815ce0))
* Oracles uses Context, beginning refactor to mimic http.Listen/Parser/Handler ([9367aeb](https://github.com/aeternity/aepp-sdk-go/commit/9367aeb72eb8b240d331f37a3717f78205d032fe))
* naet.Node.Info() returns netowrkID and node version for use with HLL ([6601570](https://github.com/aeternity/aepp-sdk-go/commit/66015706c72b10cd05ed23121b03596e44e9978e))
* AENS uses Context ([2fa3dbb](https://github.com/aeternity/aepp-sdk-go/commit/2fa3dbbc0a8d65d68b51ba0aef6ba0ead4949b47))
* Contracts uses Context ([d7ea73a](https://github.com/aeternity/aepp-sdk-go/commit/d7ea73a1670c7195886a64c999b4b94ecdf4638c))
* Context struct with ContextInterface ([000f1ed](https://github.com/aeternity/aepp-sdk-go/commit/000f1eddccb6712ce0c7cb862434d815ed5f7263))
* SignBroadcastWaitTransaction returns a *TxReceipt ([479607f](https://github.com/aeternity/aepp-sdk-go/commit/479607f30069bb3e4e2c26822ce37344887d6511))
* introduce TxReceipt, Context. 2nd level functions are now methods of Context ([b02675c](https://github.com/aeternity/aepp-sdk-go/commit/b02675c0c80725066226ddc7fc2d7456af3161dc))
* Noncer now returns the current height in addition to the recommended TTL, so NewNameUpdateTx does not need separate ttler, noncer input arguments and does not have to call ttler twice. transactions.GenerateTTLNoncer renamed to NewTTLNoncer ([e8aa318](https://github.com/aeternity/aepp-sdk-go/commit/e8aa3188225fad4489d8e4384f1f8b5bed17bc6b))
* AENS higher level TestRegisterName uses Broadcaster; CreateContract returns ctID; CallContract introduced ([6b8280f](https://github.com/aeternity/aepp-sdk-go/commit/6b8280fa8dfe262ec92267baf19f3cd85e4607c5))
* try to use Broadcaster struct in higher level interface ([9cf81a8](https://github.com/aeternity/aepp-sdk-go/commit/9cf81a809f576dee3511a6f6248ef819f5b3293c))
* introducing Broadcaster{} struct, which is like Context{} but only handles broadcasting ([0f3625c](https://github.com/aeternity/aepp-sdk-go/commit/0f3625ce783af3c600121a96688fef7fce23e389))

### Docs

* DESIGN.md explains the current state of the Go SDK ([f878a75](https://github.com/aeternity/aepp-sdk-go/commit/f878a75ed5ec43f8e08fa6b0a093129bb2f5eca4))
* add comments for higher level interface constructs ([157e24c](https://github.com/aeternity/aepp-sdk-go/commit/157e24ce26880f8df735f50de83753cd4a47428a))
* comments for Context{} ([f0a24e4](https://github.com/aeternity/aepp-sdk-go/commit/f0a24e4ca76ed672b6259bd5c2b13d035fd8ebfa))

### Feature

* beta DefaultCallResultListener. There's probably a better way to do this, but this seems to work. ([7d31909](https://github.com/aeternity/aepp-sdk-go/commit/7d31909bc8e842e03c96e7ec8905e8b4623195db))
* Node now supports GetTransactionInfoByHash, required for contract call return values ([116fa9d](https://github.com/aeternity/aepp-sdk-go/commit/116fa9d09038bb7c259ef18450d3975e7ae4a416))
* NamePointer generalized to allow for arbitrary pointer types, as well as account_pubkey, oracle_pubkey, contract_pubkey, channel types ([6f62f04](https://github.com/aeternity/aepp-sdk-go/commit/6f62f041728004daab663d06d7707f1b3a91f0f5))
* CreateOracle, ListenOracleQueries higher level convenience functions. New naet.Node capability required: GetOracleQueriesByPubkey ([8b3d66d](https://github.com/aeternity/aepp-sdk-go/commit/8b3d66d9ce6c324e1cc46ea847f27cc8ebfd9fef))

### Features

* OracleRegisterTx.ID(), OracleQueryTx.ID() methods ([f55cd2d](https://github.com/aeternity/aepp-sdk-go/commit/f55cd2de47af1b7467bd13de7347b94ec9bcec0d))

### Tests

* TestSignBroadcast, TestfindVMABIVersion ([fd88f51](https://github.com/aeternity/aepp-sdk-go/commit/fd88f51bb70e435050bc35c64def2956447a1fe7))
* TxReceipt.Watch unittest ([21a82d4](https://github.com/aeternity/aepp-sdk-go/commit/21a82d4eee6e0e300b9156829442cddea20dc57d))
* integration test for oraclize ([c515139](https://github.com/aeternity/aepp-sdk-go/commit/c515139a7004bb7ed89c16b5c86ec5f9f7ba58dc))


<a name="v7.0.1"></a>
## [v7.0.1](https://github.com/aeternity/aepp-sdk-go/compare/v7...v7.0.1) (2019-11-25)

### Bug Fixes

* account sign not updated to use SerializeTx() ([9217c01](https://github.com/aeternity/aepp-sdk-go/commit/9217c01fa9c51752a1ae1d15d8269a28fe3a7fc9))
* make ABIVersion==3 a const ([0f45e0a](https://github.com/aeternity/aepp-sdk-go/commit/0f45e0a7edd023c2b17966d91eb3e3a8245cf185))
* docker-compose v1.24 changed behaviour to complain about duplicate mountpoints. Also expose apparently does not automatically publish ports to the host anymore ([8df9761](https://github.com/aeternity/aepp-sdk-go/commit/8df976131d6d49e210d90c0a61165fedff604b94))

### Chore

* rename GetWalletPath() to CheckWalletExists() ([ec08cb9](https://github.com/aeternity/aepp-sdk-go/commit/ec08cb9ddd6c2d75dc28e87e1d84105bb1ea1e31))

### Docs

* HD wallet example, moved package aeternity example from context_test.go to helpers_test.go. README.md rewritten. Assume little to no knowledge of Go (assume blockchain knowledge). Make it easier for beginners to know where to look for which code examples. ([5498789](https://github.com/aeternity/aepp-sdk-go/commit/549878913bb4c38f180f5108dc2fe7e98b2bb104))

### Features

* updated fee calculation for FATE Contract Calls ([d5c1b18](https://github.com/aeternity/aepp-sdk-go/commit/d5c1b18848a8d41f4f5cc7b3318cddc54782621c))

### Maintain

* it works with node v5.0.2 ([0563b58](https://github.com/aeternity/aepp-sdk-go/commit/0563b584bd9f1d16fe75cc1adbb0145b50a250e5))


<a name="v7"></a>
## [v7](https://github.com/aeternity/aepp-sdk-go/compare/v7.0.0...v7) (2019-11-06)


<a name="v7.0.0"></a>
## [v7.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v6.0.0...v7.0.0) (2019-11-06)

### Bug Fixes

* updated config.go QueryTTLValue, ResponseTTLValue changed Config Defaults Oracle Response unittest ([a7cfbc5](https://github.com/aeternity/aepp-sdk-go/commit/a7cfbc531ba87d2291b301386ec2000db157fa4e))
* some New*Tx arguments were left over ([0d4c9b1](https://github.com/aeternity/aepp-sdk-go/commit/0d4c9b148becbd4a62c1bc2ea727a9297c3d6366))
* Default OracleTTLValue introduced; Oracles have a ABIVersion, not a VMVersion ([24b005e](https://github.com/aeternity/aepp-sdk-go/commit/24b005e1b33b3ea3c09ed9ec8384a24ec53adda7))
* namesalt was accidentally used for namefee ([2ee99ca](https://github.com/aeternity/aepp-sdk-go/commit/2ee99ca02199efa818f7a5f9687edadf693c1305))
* go 1.13 changed package import behaviour, causing test flags to be parsed before they were declared https://github.com/golang/go/issues/31859 In case anyone else lands here and is looking for a simple answer: ([1e8cf5b](https://github.com/aeternity/aepp-sdk-go/commit/1e8cf5bd8e959ea17e73ee277f7c147c585d188d))
* SignBroadcastWaitTransaction and friends shoud take interface instead of *naet.Node ([5848d92](https://github.com/aeternity/aepp-sdk-go/commit/5848d9264a7719592a67fbeb4a6684969523aa58))
* package aeternity accidentally imported v1 of aepp-sdk-go ([25fec10](https://github.com/aeternity/aepp-sdk-go/commit/25fec10c40ae86826de9465fad8e0699615d60f0))
* update config values for Lima Default Compiler Backend is FATE (updated VMVersion/ABIVersion too) Default GasLimit is lowered (previously was larger than protocol's microblock gas limit of 6e6) ([59d65fb](https://github.com/aeternity/aepp-sdk-go/commit/59d65fb29a782f76cd0415eac4c4b53d47ed89e8))
* package aeternity was wrongly importing aepp-sdk-go without a major version ([b9fb5aa](https://github.com/aeternity/aepp-sdk-go/commit/b9fb5aa5e6447aea07d4fc36bc4b6003ece8b8eb))

### Chore

* update aepp-sdk-go import paths to v7 ([961a75b](https://github.com/aeternity/aepp-sdk-go/commit/961a75b0f447e156fa2f66afe14651656d9c8736))
* update go.mod, go.sum ([5986f00](https://github.com/aeternity/aepp-sdk-go/commit/5986f0053a2b186a1aa9e816f9db80eb3e00a7d2))
* integration test testdata is now FATE bytecode ([501001e](https://github.com/aeternity/aepp-sdk-go/commit/501001e94d635002e7a0ff1f3dc67690031c6ade))
* use SignBroadcastWaitTransaction in integration tests ([20f2cdb](https://github.com/aeternity/aepp-sdk-go/commit/20f2cdb04a780944cc74fc13f0dd23e363cca625))
* feel comfortable about removing Default Config Tx RLP serialization unittest cases now ([4458c06](https://github.com/aeternity/aepp-sdk-go/commit/4458c061a3a7018d1fa37ab6e79af70e9c6bb8f8))
* update go.mod ([920a6b7](https://github.com/aeternity/aepp-sdk-go/commit/920a6b76b3a27f8f3a47875db1b8b14aef5c18f6))
* update to compiler v4.0.0 ([1e1d3c7](https://github.com/aeternity/aepp-sdk-go/commit/1e1d3c7d90faa1ab7ad7e028a6d4e041b7be6eeb))
* private testnet node should upgrade to Lima consensus ASAP ([6918690](https://github.com/aeternity/aepp-sdk-go/commit/691869078de2951b853c58dfa12972501f451eb1))
* WaitForTransactionForXBlocks could really use a config parameter that has a suggested waiting time in blocks ([ae20c6a](https://github.com/aeternity/aepp-sdk-go/commit/ae20c6ae8a8cdf9ca5b81ca94d9c9fcbeb9ce385))

### Code Refactoring

* restrict HD wallet functionality to only deriving aeternity addresses ([6f37c3b](https://github.com/aeternity/aepp-sdk-go/commit/6f37c3b5778a491bebce6745a92d9d07bc4f6ff7))
* Transaction interface now includes the TransactionFeeCalculable methods ([4575526](https://github.com/aeternity/aepp-sdk-go/commit/45755269610c36af6cdc00a7270f1626796de929))
* Transaction fee calculation reworked to accomodate the different gas calculation equation for Oracle*Txs ([1875246](https://github.com/aeternity/aepp-sdk-go/commit/18752460e1692a4b68d0c37ccd4fe5776eb6cf9d))
* cmd uses TTLer/Noncer ([2e3924f](https://github.com/aeternity/aepp-sdk-go/commit/2e3924f6d6cc954b31d5e05fc082d848117f0319))
* integration tests use TTLer/Noncer method ([87f5777](https://github.com/aeternity/aepp-sdk-go/commit/87f577770345f47fc421bc7c211bf41d4268a826))
* CreateTTLer, CreateNoncer, GenerateTTLNoncer, CreateTTLNoncer ([5f06729](https://github.com/aeternity/aepp-sdk-go/commit/5f067294f1d1afc457de2e154fc7ace5d145ecc2))
* GA transactions follow new closure system ([8dadb77](https://github.com/aeternity/aepp-sdk-go/commit/8dadb776917761b39bf8ab3cd770fbd0020d7b00))
* removal of Context struct ([e0801a9](https://github.com/aeternity/aepp-sdk-go/commit/e0801a90c1c53f55de06e6cd20f2b461c11cfc24))
* Contract*Tx constructors now do the work previously done by Context ([529a6a0](https://github.com/aeternity/aepp-sdk-go/commit/529a6a07cd2f33eee0d81f0f7a0182364336308e))
* Oracle*Tx constructors now do the work of Context ([f549966](https://github.com/aeternity/aepp-sdk-go/commit/f5499663d45669893a4ac8e64fd4b225ee2ad7fb))
* SpendTx, AENS transaction constructors now do a lot of hard work previously done by Context move GetTTL/GetNextNonce to package transactions ([40f5481](https://github.com/aeternity/aepp-sdk-go/commit/40f5481295a5905a764fbf9698de5e561838b01f))

### Docs

* comments for NameID, ErrWaitTransaction ([a4d7386](https://github.com/aeternity/aepp-sdk-go/commit/a4d7386843f945eef652c6da46af1754d30d7472))

### Feature

* AEX-10 HD Wallet and mnemonic support ([df76365](https://github.com/aeternity/aepp-sdk-go/commit/df763655f8f654fc99036833f9d1a2c4ee44c0a0))
* WaitTransactionForXBlocks and its derivatives now return a ErrWaitTransaction struct which helps callers distinguish between network errors and transaction acceptance errors ([6b4725b](https://github.com/aeternity/aepp-sdk-go/commit/6b4725b7cb4e4013a6086ee033dd25ef344f2b08))
* integration tests for higher level ([c9ee7fb](https://github.com/aeternity/aepp-sdk-go/commit/c9ee7fb041f1a9457d7a4d02f68fd3fcc56431e2))
* higher level AENS supports NameFee ([d7bf5fd](https://github.com/aeternity/aepp-sdk-go/commit/d7bf5fd3cc36f155ab3e06eee94ffce64bd1158a))
* update AENS integration tests for Lima HF ([8d9962d](https://github.com/aeternity/aepp-sdk-go/commit/8d9962df17b668fb239da21d237f2be1a453da88))
* Lima HF NameClaimTx update and NameID calculation changes ([be88be8](https://github.com/aeternity/aepp-sdk-go/commit/be88be8a7dbe487431f9a4fe9d42b913a37c9209))
* AENS auction config params ([b51ebb3](https://github.com/aeternity/aepp-sdk-go/commit/b51ebb33e429ee6cea07112a18fd8cea61bf4c96))
* test with node v5.0.0-rc.5 ([bc13ca2](https://github.com/aeternity/aepp-sdk-go/commit/bc13ca20eb1e78c7b8f8dfd647f014d20ca6382a))
* new name hashing scheme after Lima HF ([237cb40](https://github.com/aeternity/aepp-sdk-go/commit/237cb40c44243b394b98ab294151811bbd63bde0))
* RegisterName(), CreateContract() ([f6a1ed2](https://github.com/aeternity/aepp-sdk-go/commit/f6a1ed291b611c7d7316511e273097e592d176e2))

### Features

* improve vanity account generation experience ([51440e8](https://github.com/aeternity/aepp-sdk-go/commit/51440e8d919f04f2861b3c7a702714ed751e0c34))


<a name="v6.0.0"></a>
## [v6.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v5.1.2...v6.0.0) (2019-10-10)

### Bug Fixes

* ExampleGetTransactionType test was failing when run as part of the unittest suite, but working properly when run individually. This is because TransactionTypes map was being modified somewhere. Instead of investigating what changed it (non-obvious), I simply wrapped TransactionTypes map in a function to provide the const map behaviour originally intended. ([a7fdfae](https://github.com/aeternity/aepp-sdk-go/commit/a7fdfae03a900d484396aa286f906472ab920f51))
* compiler_test.go was not using config package ([35d0eae](https://github.com/aeternity/aepp-sdk-go/commit/35d0eaede96072faa8421ac49c9191e240d6aa0a))
* ensure all import paths have aepp-sdk-go/v5 ([23bf23a](https://github.com/aeternity/aepp-sdk-go/commit/23bf23abc67fef8d1cb8fac45dd5fa2ece59c9f4))
* account's testdata/ folder was not moved with it into account/ ([8f9a4f0](https://github.com/aeternity/aepp-sdk-go/commit/8f9a4f0db69f3b4565460e4106a0b106b651212e))

### Chore

* import paths to aepp-sdk-go/v6 ([bd8ba2a](https://github.com/aeternity/aepp-sdk-go/commit/bd8ba2a3375af65b8a7aeaba4b82a34a2501b3d7))
* there is no need to test the output of CompileContract, just that it finished without error ([71a5ba5](https://github.com/aeternity/aepp-sdk-go/commit/71a5ba5c05c3dcedaf1f58872fa9da538aa952ac))
* goimports moved aeternity import in ga_test.go ([373bdd0](https://github.com/aeternity/aepp-sdk-go/commit/373bdd01261a6c667a323fbfe619eacb625e7f6b))
* update golden file for simplestorage_bytecode.txt compiler v4.0.0-rc4 ([143f004](https://github.com/aeternity/aepp-sdk-go/commit/143f004f233ce1b166b33d88a992830c68cbe93a))
* functionality required of api/updatedict.py has changed since last editing node v3's swagger.json file ([e851cfc](https://github.com/aeternity/aepp-sdk-go/commit/e851cfc32d869043a8ac27dc512c46fc43c1a33c))
* aeternity.VerifySignedTx should be in package transactions ([0c8699e](https://github.com/aeternity/aepp-sdk-go/commit/0c8699e077fe15deb8a325e8fedb77fa57fd5da2))
* move prefixes in package binary to their own file ([b905bc7](https://github.com/aeternity/aepp-sdk-go/commit/b905bc73eeed3304e066829fcb65005b4e1af9e6))
* config.Config was stuttering ([f41d335](https://github.com/aeternity/aepp-sdk-go/commit/f41d335a35c1608c5a5f1dfe54322ba4b3540880))
* remove redundant 'Account' in package account's functions ([8f6b689](https://github.com/aeternity/aepp-sdk-go/commit/8f6b689f7f29c4b90e07ed62a0bf19a104518b9f))

### Code Refactoring

* package models split into models, transactions ([06c5532](https://github.com/aeternity/aepp-sdk-go/commit/06c5532581ce91d9e760c0294e86e3b994cfde4e))
* integration tests follow aeternity split ([c4648ae](https://github.com/aeternity/aepp-sdk-go/commit/c4648ae46e8ad8bcae4211e7ce2c4970dd214ce8))
* cmd follows the new aeternity split ([329766b](https://github.com/aeternity/aepp-sdk-go/commit/329766b6474d65a78977ab99ef2c5e154ad44a6d))
* split aeternity package ([3638ebd](https://github.com/aeternity/aepp-sdk-go/commit/3638ebdc781bd4d4c388fee8f86c8294037b115a))
* split identifiers.go into hashing.go/transactions.go, where they are used ([c05d256](https://github.com/aeternity/aepp-sdk-go/commit/c05d256b9b299821118d2473f56835254fccf65f))
* standardize on *big.Int being passed everywhere, not big.Int ([1311831](https://github.com/aeternity/aepp-sdk-go/commit/1311831932bece09211ee11565061581bfab7056))
* move hashing_test.go unittests to their respective locations too ([e613f71](https://github.com/aeternity/aepp-sdk-go/commit/e613f717250fe7195de267d3c6b10cdfba4f2af5))
* break up hashing.go to reduce dependencies. According to Go practice, too much DRY leads to too many dependencies, thus some duplication is allowed. leftPadByteSlice, buildOracleQueryID, buildContractID moved to transactions.go Namehash moved to helpers.go randomBytes duplicated in generateCommitmentID and moved to keystore.go:KeystoreSeal uuidV4 moved to keystore.go generateCommitmentID/computeCommitmentID moved to helpers.go buildRLPMessage, buildIDTag, readIDTag moved to transactions.go ([5188c43](https://github.com/aeternity/aepp-sdk-go/commit/5188c439c25ddd75664e64135cb8929b4b0be43d))

### Config

* we shouldn't need utils.Int outside of swagguard ([4e24ba3](https://github.com/aeternity/aepp-sdk-go/commit/4e24ba37091b68ab59525ad74aff086b33cde689))

### Doc

* README points to code examples instead ([1904f0e](https://github.com/aeternity/aepp-sdk-go/commit/1904f0e9a6762e13feee677f6c564f3e3f03c997))
* made a Context example from the Spend transaction test ([d0d5ad1](https://github.com/aeternity/aepp-sdk-go/commit/d0d5ad14c373ddd130725055bfd0fcc95eef940e))

### Docs

* ExampleSerialize/DeserializeTx, ExampleGetTransactionType, ExampleSignHashTx, ExampleVerifySignedTx ([40e2171](https://github.com/aeternity/aepp-sdk-go/commit/40e2171f853d7f60a627e6058e532ab2948ec1a6))
* Contract example ([54b4da5](https://github.com/aeternity/aepp-sdk-go/commit/54b4da5ed8744569cdfabb857c3c9b9291e7db57))
* doc.go files for each package ([e2c56d4](https://github.com/aeternity/aepp-sdk-go/commit/e2c56d4b1e7841abd2327607946b4e1c4ef16a1e))
* updated api/README.md ([b29a4e7](https://github.com/aeternity/aepp-sdk-go/commit/b29a4e78346cb37e4db04267f487496186dc5ffd))
* received many problems about setup, added documentation to README ([58a63c4](https://github.com/aeternity/aepp-sdk-go/commit/58a63c402131f31a67d0dbff831ecb0d40dc1b8a))

### Feature

* node v5.0.0-rc4 support, swagger.json rewriting utils updated ([0bbef47](https://github.com/aeternity/aepp-sdk-go/commit/0bbef475ca80c1d3bb10fc5b686b0a1b8b9585a4))
* compiler v4.0.0-rc5 swagger json, .env also updated ([f57b042](https://github.com/aeternity/aepp-sdk-go/commit/f57b0425d303e4f4cb0d2c122e03fa6286437545))
* compiler v4.0.0-rc5 support. CompileContractForbidden was renamed to CompileContractBadRequest ([55abe33](https://github.com/aeternity/aepp-sdk-go/commit/55abe33e3c8f4daa2498361fa15d7c5679e46b03))
* no need to cancel transaction fee size out anymore when calculating the fee ([39d115b](https://github.com/aeternity/aepp-sdk-go/commit/39d115bd2ad9edf3234b6b546d4a104b01cae818))
* accurate tx fee calculation with unittests ([ba9403b](https://github.com/aeternity/aepp-sdk-go/commit/ba9403b089b23192ae426c3b9ab77bbd0b678076))
* swagger.json from node v5.0.0-rc.2 ([a8c7955](https://github.com/aeternity/aepp-sdk-go/commit/a8c795525979d4057fb3f50a0183c212ffece811))
* support node v5.0.0-rc.2 ([2df1c10](https://github.com/aeternity/aepp-sdk-go/commit/2df1c10761cee3e702a06d1db4c874d44d11552d))


<a name="v5.1.2"></a>
## [v5.1.2](https://github.com/aeternity/aepp-sdk-go/compare/v5.1.1...v5.1.2) (2019-09-14)

### Bug Fixes

* Compiler's Error models changed slightly - ensuring that they all return the underlying Reason ([caf7069](https://github.com/aeternity/aepp-sdk-go/commit/caf70697462d2bf813c4ad7aeb3ba180cc815df1))

### Feature

* basic support for compiler v4.0.0-rc4 ([152bac5](https://github.com/aeternity/aepp-sdk-go/commit/152bac531c027367cf66c3dfdb5e00637a94a7c2))


<a name="v5.1.1"></a>
## [v5.1.1](https://github.com/aeternity/aepp-sdk-go/compare/v5.1.0...v5.1.1) (2019-09-12)

### Bug Fixes

* aepp-sdk-go was importing v1 of itself, causing havoc for module support. Now v5 should be totally standalone. ([852dc24](https://github.com/aeternity/aepp-sdk-go/commit/852dc24b33ab3d024af46c57b7c574f8a2ccfa07))

### Chore

* reenable Contract tests in api_test.go. Using hardcoded fee/GasLimit values - can only be found through trial and error. TODO: improve user experience here! ([8732f08](https://github.com/aeternity/aepp-sdk-go/commit/8732f0832e08a1a540de0631e3ec9dacd8996e3f))


<a name="v5.1.0"></a>
## [v5.1.0](https://github.com/aeternity/aepp-sdk-go/compare/v5.0.0...v5.1.0) (2019-09-11)

### Bug Fixes

* Versions before Lima had inconsistent naming for OracleResponse/OracleRespondTx. Now it is standardized to OracleRespondTx ([acbe52c](https://github.com/aeternity/aepp-sdk-go/commit/acbe52c1ebb1cacf9cdd1d13ff2529c619fc8293))
* move away from custom signBroadcast() code for integration tests. Integration tests should look like example code from now on ([18547a1](https://github.com/aeternity/aepp-sdk-go/commit/18547a1e2ea9724aa59a18d16349b187ae7fbaa9))
* OracleConfig.VMVersion should be uint16, not uint64 ([3cb08f7](https://github.com/aeternity/aepp-sdk-go/commit/3cb08f74f4c472936e47ce61e5194101045cfe13))
* Context.OracleRegisterTx helper had a named argument abiVersion but it's actually VMVersion ([755d779](https://github.com/aeternity/aepp-sdk-go/commit/755d77906ecbdd7e5fa2ba942f19838a7b43467c))
* README.md had wrong example code. Oops. ([93e2651](https://github.com/aeternity/aepp-sdk-go/commit/93e2651df4cd86b32f4de33c413c68bf8890d1df))
* stringer interface also wants a String() method with a non-pointer receiver ([ac3533c](https://github.com/aeternity/aepp-sdk-go/commit/ac3533c6f762eb1473eef85fdd987934c91eac6f))
* PrintGenerationByHeight() was swallowing errors and not printing out transactions in a generation ([8a9f619](https://github.com/aeternity/aepp-sdk-go/commit/8a9f619c6400cdd2b4624b6d4814a116f6cb1f19))
* utils.BigInt.MarshalJSON() should, surprisingly, NOT have a pointer receiver. This allows json.Marshal() to work on it. ([ebd8af4](https://github.com/aeternity/aepp-sdk-go/commit/ebd8af4a757ebcbed850dcdf95508eaa07852372))

### Chore

* README was not updated somehow, updating again ([1b7c26d](https://github.com/aeternity/aepp-sdk-go/commit/1b7c26d6be52f7e7be582f9bfeeaf93cdec672c0))
* better comments for config, internal functions Examples for Encode/Decode functions ([0b6e64b](https://github.com/aeternity/aepp-sdk-go/commit/0b6e64b91087085e551898511c74f33c1e058fd8))
* Tx.Gas renamed to GasLimit. GasPrice moved from Contracts to Config.Client (since it is a constant, used in all tx fee calculations) Cleaned up unused PreclaimFee, ClaimFee, UpdateFee in AENSConfig ([4260720](https://github.com/aeternity/aepp-sdk-go/commit/426072071aa7f8083bf7ac0d5f64006393b7cdd4))
* cleaned up outdated CLI docs ([922444a](https://github.com/aeternity/aepp-sdk-go/commit/922444a74701c0d85c87a76035ffe883890c949b))
* improved comments in some places ([f69d9fe](https://github.com/aeternity/aepp-sdk-go/commit/f69d9fe1be71678b609199bf7fdead0193fce0ea))
* better comments for helpers.go ([7344548](https://github.com/aeternity/aepp-sdk-go/commit/73445489b6106560839c529e8d9c54dc8ab2fc86))
* ChainPollInterval typo ([beab494](https://github.com/aeternity/aepp-sdk-go/commit/beab494c1418127e007115a51a618109838e2239))
* updated Contract struct unittest for the new Payable field ([314997b](https://github.com/aeternity/aepp-sdk-go/commit/314997b223a5f31238e1a16525a733effc01b308))
* standardize on returning pointers for Transaction structs ([b5e6c1c](https://github.com/aeternity/aepp-sdk-go/commit/b5e6c1ce0f7a3f435c6465f3b2a1b57fe1af5eb2))
* use consts in config for some config parameters ([2d85818](https://github.com/aeternity/aepp-sdk-go/commit/2d858186eb1c7460a692f3d97369a90efcd129ad))
* update .env to node v5 rc1, compiler v4-rc2 ([a4902a4](https://github.com/aeternity/aepp-sdk-go/commit/a4902a43c558309ed3c73d05b9e0c5ff8d4eb726))
* update .env to lima/compiler v4 rc1 ([9bd8065](https://github.com/aeternity/aepp-sdk-go/commit/9bd8065458805ad4369f4e46a7680e6dcad90805))
* use more reference-quality BlindAuth smart contract for golden smart contracts package, use it in GA integration test ([18c973c](https://github.com/aeternity/aepp-sdk-go/commit/18c973c5e01f9735e7549af4d6d2b12565402eed))
* finish integration test coverage for compiler swagger interface ([811756c](https://github.com/aeternity/aepp-sdk-go/commit/811756c84fa4d080712480be6260059d7f999585))
* contracts and their bytecodes and calldata were getting too much to manage so introduced a 'golden' package to store them all. Only useful for code which just needs some contract artifact but won't care what it really is. ([20b7979](https://github.com/aeternity/aepp-sdk-go/commit/20b79796f9c79124fff8ca6eb4b1088c0ecc8736))

### Cleanup

* newly calculated dependencies, tidied up struct tags in config.go ([27c28a2](https://github.com/aeternity/aepp-sdk-go/commit/27c28a275421a2788f6aa18551aa5fe08c8769f8))

### Code Refactoring

* Replaced Helpers with closures for GetTTL, GetNextNonce, GetAnythingByName. standardizing referring to *aeternity.Node as 'node' better comments for helpers.go ([3c1ad87](https://github.com/aeternity/aepp-sdk-go/commit/3c1ad87c2b11a8952d4fbbd5cb0122039ba8a9ad))
* BroadcastTransaction removed - it was only used in the CLI chain broadcast subcommand, and when used for other code, recalculates the hash needlessly. hashing.go hash() renamed to Blake2bHash() ([7bf70e4](https://github.com/aeternity/aepp-sdk-go/commit/7bf70e4fd8c73684927e241505792db2a3a06c07))
* integration tests use gotest.tools/golden so that compiler tests can easily update values. Support Lima AEVM/FATE option. ([8618d52](https://github.com/aeternity/aepp-sdk-go/commit/8618d5245690d7c19e453601356868b6a094634f))
* removing golden package that was shared between package cmd and integration_test, because infrastructure was too complex for the benefit it provided. cmd and integration_test will have their own golden contract variables, in separate files. Tests will show if these are out of date for the compiler. ([490e19f](https://github.com/aeternity/aepp-sdk-go/commit/490e19f4f83f024de9e1c3eeaa1d4f3e39884ed0))
* aeternity.Config.Compiler struct. Introduced configuration consts for several parameters. ([6f08984](https://github.com/aeternity/aepp-sdk-go/commit/6f089841bac7e9787690c438396e56f4fd48aee4))
* SignBroadcast() is useful enough that it should be a helper, as SignBroadcastTransaction() WaitForTransactionUntilHeight() required wrapping code, but WaitForTransactionForXBlocks() is more useful ([eda4f69](https://github.com/aeternity/aepp-sdk-go/commit/eda4f696b1eb18e26fdaf4cbf4ff90031702aea2))

### Feature

* SignBroadcastWaitTransaction() ([733cde4](https://github.com/aeternity/aepp-sdk-go/commit/733cde4ea9222b0d3b09bca644e24e7c0a06a036))
* Contract struct supports Lima RLP serialization ([df1249b](https://github.com/aeternity/aepp-sdk-go/commit/df1249bd9860a32073fce8fdc5e53a641cf5d96c))
* support for compiler v4's FATE/AEVM backend switch ([0c263d3](https://github.com/aeternity/aepp-sdk-go/commit/0c263d3d3bf13c73533bcf8d0045181b3a95909e))
* compiler 4.0.0-rc1 support ([4cb9265](https://github.com/aeternity/aepp-sdk-go/commit/4cb9265b52f63aefa13c00ece373b5adad6f963b))


<a name="v5.0.0"></a>
## [v5.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v4.0.1...v5.0.0) (2019-08-13)

### Bug Fixes

* GAMetaTx.Tx only accepts a *SignedTx, not just any Transaction (interface) ([b5ebe72](https://github.com/aeternity/aepp-sdk-go/commit/b5ebe72b3d3fce6da20fd6449d93dfff2881a007))
* now that rlpae deserializes big.Int(0) from 0x00 more correctly, the Equal method on the Tx struct is unnecessary ([b622992](https://github.com/aeternity/aepp-sdk-go/commit/b622992c5633fe207520a5420f10d9395814badb))
* Tx struct JSON() method did not serialize big.Ints properly, so they were always coming out as 0 ([a19900f](https://github.com/aeternity/aepp-sdk-go/commit/a19900fd78b0f0ba1a36f7d072e6c688abf70cde))
* Config.Client.Contracts.VMVersion updated for node v3 ([de87f33](https://github.com/aeternity/aepp-sdk-go/commit/de87f3342a22707480d39ce4705c3086121b4f8d))
* working unittest for Contract struct, no need for embedded contract struct, comments ([7bad2d2](https://github.com/aeternity/aepp-sdk-go/commit/7bad2d23403eea4e85b34691c951594ac4fc044b))
* BroadcastTransaction shouldn't ignore returned errors anymore ([aab5748](https://github.com/aeternity/aepp-sdk-go/commit/aab574811b9235a1be45046f811551f86940df3a))
* update Sophia contract in integration test ([f893b51](https://github.com/aeternity/aepp-sdk-go/commit/f893b5134e341f44ad3cfb11a03f63e7346e8501))
* Sophia 3.2 needed changes to the 'should compile' test contract ([ecf96d3](https://github.com/aeternity/aepp-sdk-go/commit/ecf96d316de4a242c8b1604ce2f98cc607b1442b))
* SpendTx Payload should actually be a bytearray ([361aea5](https://github.com/aeternity/aepp-sdk-go/commit/361aea5409f431ae67f612b4d734844b38c93a8b))
* BigInt implements the MarshalJSON() interface, fixing JSON serialization of the tx structs ([8d2bcb3](https://github.com/aeternity/aepp-sdk-go/commit/8d2bcb3a87703052fd0c8f4b29d70c9c19172d86))

### Chore

* update go module version to v5 ([e29afb3](https://github.com/aeternity/aepp-sdk-go/commit/e29afb3e995362cf4ad767984a4485eb2de473c0))
* update README.md for refactored Tx signing ([46c899c](https://github.com/aeternity/aepp-sdk-go/commit/46c899c5d875f534c623aca8bc797a97c331a90a))
* fix misspelt words ([443fcb5](https://github.com/aeternity/aepp-sdk-go/commit/443fcb534451f4e9cdeed5afd6bf44aa1e541821))
* cleanup SpendTx code, better name for embedded spendrlp struct ([46be029](https://github.com/aeternity/aepp-sdk-go/commit/46be029512c41bb83f2d611ac2487cad8038db8c))
* comments for SignedTx struct ([0abf617](https://github.com/aeternity/aepp-sdk-go/commit/0abf61750b573a3277b9136b4a476866726bb062))
* rename transaction RLP unittests to EncodeRLP ([823f14b](https://github.com/aeternity/aepp-sdk-go/commit/823f14b85b48892183cea9b3768addd6defd0276))
* update README.md for refactored code ([a9af711](https://github.com/aeternity/aepp-sdk-go/commit/a9af711bff418a843e0e711077c837ca140e61f6))
* another unittest for GAAttachTx ([de102db](https://github.com/aeternity/aepp-sdk-go/commit/de102db471f28542d72777dfe4df8914708f8d41))
* move compiler-url option in CLI to root command to be next to --external-api ([b15da47](https://github.com/aeternity/aepp-sdk-go/commit/b15da47d4c8c7b7faff82d2c2bc95c85a9004f60))
* comment for aeternity.GetHeightAccounter ([1908df4](https://github.com/aeternity/aepp-sdk-go/commit/1908df4a23c2826087c263e7d1fbd2edd2c503c2))
* Unit tests for most datatypes in CLI inspect subcommand. refactor: break CLI inspect subcommand down so it is unit testable. ([578e69b](https://github.com/aeternity/aepp-sdk-go/commit/578e69b4658c57c55869f25dac5224e2af76494e))
* use node v3.3.0, compiler v3.2.0 ([abc1ae7](https://github.com/aeternity/aepp-sdk-go/commit/abc1ae7a09529eac1df6fd82675020c30f217750))
* api.go wrap to 80 chars ([a575059](https://github.com/aeternity/aepp-sdk-go/commit/a575059b3d24bea9c191643bde36ff3eded9ed46))
* comments for Node{} individual interfaces ([a6f7f54](https://github.com/aeternity/aepp-sdk-go/commit/a6f7f54545b7c4c81b0597f564d3b7bb333eb144))
* remove var names from aeternity.NodeInterface ([54ac957](https://github.com/aeternity/aepp-sdk-go/commit/54ac957c929d623211a5a3673462f0edff044551))
* setupNetwork() in integration_test should accept a bool argument for debug ([7d2db5f](https://github.com/aeternity/aepp-sdk-go/commit/7d2db5fb24c63177cb30ae651b912f814f236d16))
* update README.md with how to use the Node/Context ([593494e](https://github.com/aeternity/aepp-sdk-go/commit/593494e7ad0bd1dd0216172e6910516d2b45dada))
* spend_test should use WaitForTransactionUntilHeight() to check for transaction parsing issues ([ce86e49](https://github.com/aeternity/aepp-sdk-go/commit/ce86e49d0e51178bdfa3cb60b3e6843753d2b1d4))
* update tested node version to v3.3.0 ([c66c1ee](https://github.com/aeternity/aepp-sdk-go/commit/c66c1ee328a4799858cfdcf674d336af99c12a02))

### Code Refactoring

* GAAttachTx DecodeRLP support ([e1b8f97](https://github.com/aeternity/aepp-sdk-go/commit/e1b8f973e0877cbc386edde79020e507940bc629))
* All this refactoring has happened to enable CLI account sign to work in this way ([4f39e9c](https://github.com/aeternity/aepp-sdk-go/commit/4f39e9c1245fffd16a906c67032697e3ccec1a44))
* SignedTx deserialization possible + unittest ([56ed4ac](https://github.com/aeternity/aepp-sdk-go/commit/56ed4ac2b28a7c5b814ae52b902c9cc813db0748))
* Contract Txs implement DecodeRLP() ([9b0c0dd](https://github.com/aeternity/aepp-sdk-go/commit/9b0c0dd73412a752a1f5df57ca33e06512fa2242))
* Oracle Txs implement DecodeRLP() ([01fc623](https://github.com/aeternity/aepp-sdk-go/commit/01fc623a3831ad254b9ef9ca6311a532f90a92fb))
* DecodeRLP for NameTransferTx, ReadRLP() methods that simplify error handling when ID Tag reading introduced. Consistency everywhere within tx_aens.go, tx_aens_test.go. TODO: POST JSON to debug endpoints within Tx unittests ([21fbb64](https://github.com/aeternity/aepp-sdk-go/commit/21fbb64f1856c2f4d97cd1947951487b77f40e8c))
* AENS Tx structs except NameUpdateTx have DecodeRLP() ([3359273](https://github.com/aeternity/aepp-sdk-go/commit/33592730f44e1b995e232fb5a2bedc8c1b8a0a17))
* TransactionTypes hashmap between ObjectTag identifiers and Tx structs; DeserializeTx(), GetTransactionType(), Transaction interface ([6d720ca](https://github.com/aeternity/aepp-sdk-go/commit/6d720ca62bca0af6968b60f3a5c0faf230cd451e))
* break out Tx structs into their own files ([1b08e52](https://github.com/aeternity/aepp-sdk-go/commit/1b08e524303ad1dc382b617c21110a0d62eb8cd2))
* SpendTx struct deserialization ([b5c3812](https://github.com/aeternity/aepp-sdk-go/commit/b5c381299ec84b272362afeea9acd056e9c85f5e))
* readIDTag() introduced ([ca758c9](https://github.com/aeternity/aepp-sdk-go/commit/ca758c939861936e7cd2847e23d5f55cf5c529fc))
* remove old transaction signing code in favour of the SignedTx struct ([f84f224](https://github.com/aeternity/aepp-sdk-go/commit/f84f2243eaa681aa78b9daed1cc5776dfd9a83d9))
* change transaction signing to use the SignedTx{} struct ([6e4e2f5](https://github.com/aeternity/aepp-sdk-go/commit/6e4e2f581e90bb2b01fd56c9287b4e5bf6a71bfa))
* initial SignedTx{} struct ([9848314](https://github.com/aeternity/aepp-sdk-go/commit/984831473df6f9011aaacc771871fdad2241de26))
* comments for GA Integration test (not working) ([d80cba6](https://github.com/aeternity/aepp-sdk-go/commit/d80cba66dde0f33010bc3a8f30f96d027f616780))
* simplified GA integration test contract ([f976269](https://github.com/aeternity/aepp-sdk-go/commit/f976269f11bdf987bb438610363d74aa1a084b77))
* Tx structs implement rlp.Encoder interface instead of custom RLP() method ([b0289e0](https://github.com/aeternity/aepp-sdk-go/commit/b0289e0f86c61bb11a13072280b0874ae3925c43))
* rename BaseEncodeTx() to SerializeTx() ([c872a19](https://github.com/aeternity/aepp-sdk-go/commit/c872a1954299190139ba56d0883de3df082e9358))
* aeternity.NewContextFromURL(url, address) is introduced. Because it creates Helpers within itself, it hinders mocking and thus is not suited for use within the SDK. SDK level code should create the aeternity.Context{} struct directly without a constructor. As SDK code constructs the Context{} struct directly, it is now independent from the interface that it presents to the SDK using code (NewContextFromURL) ([58dc332](https://github.com/aeternity/aepp-sdk-go/commit/58dc332c0cba5d6382197f18e2db7f10f9ce11f7))
* api_test no need for a new Context for AENS, Oracle, and Spend type transactions! ([801e206](https://github.com/aeternity/aepp-sdk-go/commit/801e20617a4dae1e6caa758ea10bc9e30a2c11b4))
* integration tests use Helper{} struct instead of passing the node client directly to Context{} ([0f03064](https://github.com/aeternity/aepp-sdk-go/commit/0f03064c839746fde26c248b2ebe7498fd20afc5))
* tx spend/contract subcommand unittests should rely on HelperInterface, not API Interfaces. Also, fixed tx spend issue where it was printing what you wanted, instead of the actual tx's parameters ([a4faff0](https://github.com/aeternity/aepp-sdk-go/commit/a4faff0f46a65e461cfaa11a52be917ac9a9a8b6))
* removing Node from Context. Context helper methods only depend on the Helpers. This makes testing much easier (for the CLI mocks) ([b31e6b8](https://github.com/aeternity/aepp-sdk-go/commit/b31e6b803a05ce459bc0107fc901afaf79ecf317))
* split out helper functions into 'Helpers' struct ([6de6315](https://github.com/aeternity/aepp-sdk-go/commit/6de63159a8e043d4c8a95ed09b4ef71ae32a9d3e))
* move CLI unittest mock structs to one file for manageability ([411d668](https://github.com/aeternity/aepp-sdk-go/commit/411d668852c44611404aec756d62f13d74a0061a))
* helpers GetTTL(), GetNextNonce(), GetTTLNonce() are now vars so you can mock them out with other anonymous functions ([2a3cbd0](https://github.com/aeternity/aepp-sdk-go/commit/2a3cbd0662782c83f7c6ffd9785991100d62cb5c))
* CLI tx contract deploy unittest does not depend on network connection ([f7c38f6](https://github.com/aeternity/aepp-sdk-go/commit/f7c38f66e6a1fe9ebd7ac55505362eddc650af4a))
* CLI contract subcommand unittests now offline/online ([1852973](https://github.com/aeternity/aepp-sdk-go/commit/185297371fde2678e5c59162efba44226d0745c1))
* package cmd move test helper functions into separate test_utils.go ([317367d](https://github.com/aeternity/aepp-sdk-go/commit/317367d457b53eb4375e5697050852f1a6d5ca61))
* compiler.go interfaces for each method ([bd6fcc2](https://github.com/aeternity/aepp-sdk-go/commit/bd6fcc25f8923e2ea210366cfb2398ac4a380924))
* CLI tx spend unittest no longer needs a network connection ([e36bc54](https://github.com/aeternity/aepp-sdk-go/commit/e36bc548c1a7242f3b02a8d7e380bd1dd21919ab))
* CLI account unittests do not need network anymore ([9049f6c](https://github.com/aeternity/aepp-sdk-go/commit/9049f6ce089f7bd7d6a751c602c083f69cea20c8))
* New pattern in chain, chain_test.go: online/offline tests update comment because newAeNode() and newCompiler() are no longer vars ([e8d589b](https://github.com/aeternity/aepp-sdk-go/commit/e8d589bae8772032148aa2422c8f066945485459))
* interfaces on every method of the Node{} struct make it possible to granularly mock it out in cmd ([371eeb4](https://github.com/aeternity/aepp-sdk-go/commit/371eeb4c58cf0dc1bc4132fa4554062a1ebfca74))
* moved terminal.go from aeternity/ to cmd/, getErrorReason() was only used by api.go so moved there instead ([c769bc0](https://github.com/aeternity/aepp-sdk-go/commit/c769bc0e7534af47865df132f2a5cd19af4bc4c3))

### Feature

* GAMetaTx works, integration test checks with a SpendTx GAMetaTx.EncodeRLP(): serialize the wrapped SignedTx into a plain bytearray before including it in the RLP GAMetaTx.GaID: this should be an AccountID ([45007df](https://github.com/aeternity/aepp-sdk-go/commit/45007df57de3dd3609985d2285a14e286742192d))
* SignHashTx() detects if it is a GAMetaTx and if so makes the SignedTx with no signatures instead ([a384743](https://github.com/aeternity/aepp-sdk-go/commit/a384743f1932da66c7246276340c8b51ae91f7f0))
* GAMetaTx DecodeRLP() support (untested). NewGAMetaTx() now wraps the Tx in an empty SignedTx for you ([1c18110](https://github.com/aeternity/aepp-sdk-go/commit/1c181100c1d3c923fda1143fe1bf65d6aecef786))
* NamePreclaimTx DecodeRLP() and unittest ([c4dfeba](https://github.com/aeternity/aepp-sdk-go/commit/c4dfeba6bd031941b2a6860a441a8d5a060d313a))
* GAMetaTx initial struct ([43cef35](https://github.com/aeternity/aepp-sdk-go/commit/43cef3577ceb828aae09ed1663cb5e1e1c755206))
* oracle integration test is now idempotent ([a613619](https://github.com/aeternity/aepp-sdk-go/commit/a6136198ff2b6a55ee0d6dff40703969ade47926))
* integration test funding function to make tests idempotent ([c1f99d8](https://github.com/aeternity/aepp-sdk-go/commit/c1f99d89d855cdf3d3b99472bd783b602edf4f6a))
* initial Generalized Accounts integration test ([bd30589](https://github.com/aeternity/aepp-sdk-go/commit/bd305891cbf049b24640faf0bdce05aec440d5ce))
* GAAttachTx with working RLP serialization ([2cd3552](https://github.com/aeternity/aepp-sdk-go/commit/2cd355291d13369649ca616e13c6deb4f4f05883))
* Contract struct parses cb_ contracts and gives relevant details ([52f8332](https://github.com/aeternity/aepp-sdk-go/commit/52f83322e71c246c825d43bb188bf550bfc2813e))
* preliminary GAAttachTx ([6da7312](https://github.com/aeternity/aepp-sdk-go/commit/6da73124a0d9ff65af8c07c4c201a003c61f4547))
* NewContextFromURL() accepts a debug parameter ([23b62a1](https://github.com/aeternity/aepp-sdk-go/commit/23b62a18da8b1dfe64b5eb1b5a86c8b31393aeec))
* Helper{} functions to lookup Accounts, Oracles, Contracts, Channels with AENS ([b00ea77](https://github.com/aeternity/aepp-sdk-go/commit/b00ea772b2918c0f497bc359fc07a1b667cd870b))
* CLI AENS lookup ([d9d5eaf](https://github.com/aeternity/aepp-sdk-go/commit/d9d5eaf6991696df824380fdeb1cf701398d4c75))
* helper function to look up a Name's account_pubkey Pointer ([680f502](https://github.com/aeternity/aepp-sdk-go/commit/680f5029bd9db5f34f6cb0d06b70038c7c1d4e5a))
* CLI contract generateaci command with unittest. Output format could be improved ([70ba5f4](https://github.com/aeternity/aepp-sdk-go/commit/70ba5f4f1881d19b4bc11ff278bb6cdd5921299e))
* NewSpendTx() ensures only valid UTF-8 is used as the SpendTx payload ([41fb8d4](https://github.com/aeternity/aepp-sdk-go/commit/41fb8d44dd400d1115ede7312f781d8bb4cbcf04))
* Context.SpendTx helper, NewContext() constructor ([0275d3c](https://github.com/aeternity/aepp-sdk-go/commit/0275d3ca63a30bd0d87df0a510534d03f8c809df))


<a name="v4.0.1"></a>
## [v4.0.1](https://github.com/aeternity/aepp-sdk-go/compare/v4.0.0...v4.0.1) (2019-07-05)

### Bug Fixes

* **version:** update the module with semantic import versioning ([1d5621c](https://github.com/aeternity/aepp-sdk-go/commit/1d5621c89faf8a129fcdbef6b865a14de644dca0))


<a name="v4.0.0"></a>
## [v4.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v3.0.0...v4.0.0) (2019-07-03)

### Bug Fixes

* renamed swagguard tests to differentiate between node and compiler swagger code ([5793b1f](https://github.com/aeternity/aepp-sdk-go/commit/5793b1f74b36bbe756a4f98e67e65a7226e71c1c))
* NameRevoke doesn't need an Address ([799bb71](https://github.com/aeternity/aepp-sdk-go/commit/799bb710076f0eacb2471e31e029647045c8b94e))
* CLI chain top was crashing when printing time because node v3.0.1 defines time now as Uint64, not int64 ([23cf1c0](https://github.com/aeternity/aepp-sdk-go/commit/23cf1c0aefbd517041e08e8153b62159eac75047))
* contracts, oracle integration tests did not get their model types updated ([5710128](https://github.com/aeternity/aepp-sdk-go/commit/5710128137a03f36eb74d288bd2d79c97034a695))

### Chore

* proper comments for Compiler endpoints ([5c1a91e](https://github.com/aeternity/aepp-sdk-go/commit/5c1a91e3bac8fef9ee7b09753fc80ccbf2ae5152))
* CLI contract subcommand tests for encodeCalldata, compile ([5955bc4](https://github.com/aeternity/aepp-sdk-go/commit/5955bc467262578719bda27161252f8511599507))
* Golang - unlike %#v, %+v doesn't print integers in hex ([11f8af3](https://github.com/aeternity/aepp-sdk-go/commit/11f8af3fcf50ae7e77f1724bc262a2ccbcd71be9))
* Client{} method unittests against testnet WIP ([95b57d6](https://github.com/aeternity/aepp-sdk-go/commit/95b57d68af0bc1ce73bd248eb948faed5c7fafbc))

### Code Refactoring

* break logic/network part of some chain subcommands out, so that the network can be mocked out ([d4be4d6](https://github.com/aeternity/aepp-sdk-go/commit/d4be4d627b7de023232ddf06bbfa7c50d12622b5))
* many helper structs were redundant. Now they are all reduced down to Context{}, which stores only a Node connection and Address on which to query the nonce. Signing is to be done totally separately, the helpers only create unsigned txs ([62a7590](https://github.com/aeternity/aepp-sdk-go/commit/62a75908a62d5173409e8cb30960cf09ad1b39e9))
* Client{} helper struct renamed to Node. So, Node is a connection to a node. Compiler is a connection to a compiler. ([568853e](https://github.com/aeternity/aepp-sdk-go/commit/568853efb092cc3b6de6aca7c6327ac5e1cde376))
* Contract{} helper struct does not need to know the Account. It just needs a ak_ public key. In fact many helper structs should be refactored this way in the future. ([1a4cb39](https://github.com/aeternity/aepp-sdk-go/commit/1a4cb39c3a47b6b7670a72ad0a1c43dfcff4b0cf))
* move node swagger generated code into swagguard/node ([2e644fd](https://github.com/aeternity/aepp-sdk-go/commit/2e644fdaceb82971dc6bf8d114a3fcb188c533c2))
* Client API tests moved to integrationtest package, tested node version bumped to v3.1.0 ([428343e](https://github.com/aeternity/aepp-sdk-go/commit/428343ea4be8975e8283d44623da79f7656879cd))
* EncodedHash, EncodedPubkey, EncodedValue, EncodedByteArray -> string ([0d365f8](https://github.com/aeternity/aepp-sdk-go/commit/0d365f8141ec8f507e8650c2e3634952eacf40ed))
* UInt16, UInt32, UInt64 -> type: integer, format: uint* ([6d3bd5c](https://github.com/aeternity/aepp-sdk-go/commit/6d3bd5cfa0b9d8dc3de379c444ab5624af3b7f02))
* urlComponents() does not belong in helpers. It is a function used only once in node.go Comments for Contract* structs and constructors ([7e4286e](https://github.com/aeternity/aepp-sdk-go/commit/7e4286e24a00ee001280d6bb5a2d992569a798b2))
* API refactoring to use more interfaces and be unit testable Client{} is now a struct that has methods that correspond directly to the node's endpoints helpers.go helper functions that implement useful business logic like GetTTLNonce, WaitForTransactionUntilHeight, BroadcastTransaction etc. now take an interface instead of being attached to Client{}, making them testable. Unittest for WaitForTransactionUntilHeight() AENS/Contract/Oracle helpers stay as they are. Not much testing needed there since they only use GetTTLNonce. ([25ef16e](https://github.com/aeternity/aepp-sdk-go/commit/25ef16e415a564abd0d152c5736643602bc28067))
* Client{} API* methods renamed to not have an API prefix. See Dave Cheney's advice on naming ([6610ae8](https://github.com/aeternity/aepp-sdk-go/commit/6610ae8a403561177b6560e3d4a9cf63bbc8916e))

### Feature

* CLI contract decodeCalldata subcommand, tx deploy subcommand and integrationtest (fails because alice does not exist on mainnet) refactor: argument validation moved into IsAddress(), IsBytecode(), IsTransaction() ([fe8c046](https://github.com/aeternity/aepp-sdk-go/commit/fe8c04666c76121811f7dcb25b21d6ca06ed2d0b))
* CLI contract encode(Calldata) subcommand ([b6f00e1](https://github.com/aeternity/aepp-sdk-go/commit/b6f00e1bc2ddc66b4f7a90c338fc786d2ef6335e))
* CLI contract compile subcommand with --compiler-url option ([ae0891b](https://github.com/aeternity/aepp-sdk-go/commit/ae0891b501b741bc8ec8f2278e9760c62581fcc7))
* preliminary support for compiler's HTTP endpoints. Contract{} helper struct gets a *Compiler attribute ([0bdf629](https://github.com/aeternity/aepp-sdk-go/commit/0bdf629835d468406eabf5ad835b32479b04ce2b))
* compiler v3.1.0 swagger spec and generated code ([e98eb32](https://github.com/aeternity/aepp-sdk-go/commit/e98eb323ef41bfa1a6018e855c7b314de7c39ebd))

### Rlp

* use own fork of go-ethereum rlp. Remove rlp/ from aepp-sdk-go repo. ([96ddd4a](https://github.com/aeternity/aepp-sdk-go/commit/96ddd4a7f5108a1256a9bd624e3694ade2e1ea9f))


<a name="v3.0.0"></a>
## [v3.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v2.2.0...v3.0.0) (2019-06-06)

### Bug Fixes

* Error.Reason dereferencing and generic_tx deserialization ([9988115](https://github.com/aeternity/aepp-sdk-go/commit/99881157b59af5bf9fb175a8373975d2c071795a))
* tx verify unittest was failing ([6dccb3c](https://github.com/aeternity/aepp-sdk-go/commit/6dccb3c0b914e765bc7bf913145cd582986c3b2f))
* FeeEstimate should be less brittle by not looking for a specific value ([ff7e74f](https://github.com/aeternity/aepp-sdk-go/commit/ff7e74ff2fe05843c111749d46a562d6e886021e))
* oracle integration test was brittle when it queried the node too quickly for OracleRespondTx stage ([90ed1a2](https://github.com/aeternity/aepp-sdk-go/commit/90ed1a24a4ff41fa5c2b7663e512d76aeeb99683))
* forgot to update OracleRegisterTx ABIVersion in a unittest ([1079858](https://github.com/aeternity/aepp-sdk-go/commit/1079858b11563f823a79ee423857f70e5859f9ed))
* oracle integration test could be brittle because it didn't wait after the OracleRegisterTx to query the node by oraclepubkey ([77c4e20](https://github.com/aeternity/aepp-sdk-go/commit/77c4e2099477627b97e635aa34e2ddd4e8d345d5))

### Chore

* big.Int constructor rename NewBigIntFromString -> NewIntFromString RequireBigIntFromString -> RequireIntFromString NewBigIntFromUint64 -> NewIntFromUint64 ([0846805](https://github.com/aeternity/aepp-sdk-go/commit/08468055e751f5bb25b418e9e1d7c60bb66fc54e))
* BigInt's own Cmp() simplifies code in LargerThanComparisons() ([5c356e4](https://github.com/aeternity/aepp-sdk-go/commit/5c356e420d8f096d193b342c74c1a0604a7c0b5a))
* cleanup utils, cmd was the only package using it ([21f6e2f](https://github.com/aeternity/aepp-sdk-go/commit/21f6e2fdfcfd3133c0faefa46d38a761bb96b8ad))
* better comment for UnmarshalJSON ([2d86bb3](https://github.com/aeternity/aepp-sdk-go/commit/2d86bb3d197563523005631bcf1c768d146a39e1))
* better names for 'utils' ([c4186c2](https://github.com/aeternity/aepp-sdk-go/commit/c4186c29ccff960776dcb9b49a928df8999d8cb1))

### Code Refactoring

* BigInt newtype has a working Set(), allowing UnmarshalJSON() to work properly ([6ef4866](https://github.com/aeternity/aepp-sdk-go/commit/6ef4866fe38ac79a808f51af0577777b70ea7617))
* terminal printIf() recognizes big.Int and prints accordingly ([54fce99](https://github.com/aeternity/aepp-sdk-go/commit/54fce99c83fa7356c0f12ad920e7b5e88a2b0643))
* transaction structs now cast to utils.BigInt whenever necessary ([df11145](https://github.com/aeternity/aepp-sdk-go/commit/df11145bcc0475aae03c1fba913eb79357e371d9))
* utils.BigInt is now just a type aliased big.Int ([62768a9](https://github.com/aeternity/aepp-sdk-go/commit/62768a90908ff5f8faab6caafd2e855e6bcd316b))
* integration tests cleaned up with closures and test setup functions. More robust because it waits longer in delayableCode ([b3dc468](https://github.com/aeternity/aepp-sdk-go/commit/b3dc468597ed1a5b01b083c9bd8f8bb2fa37fe5f))
* setupNetwork(), setupAccounts(), usage of t.Log(), go-ish variable naming. Includes Spend Transaction Integration Test ([f8951c1](https://github.com/aeternity/aepp-sdk-go/commit/f8951c190aa32b33b0094977e56c9906981ff65f))
* BigInt LargerThanZero functions return a bool, not an error ([454d0a7](https://github.com/aeternity/aepp-sdk-go/commit/454d0a78248fcfefdb641a3ca7ade077aa466e4b))

### Feature

* update .env for v3.0.1 integration test docker image ([1c0ffdc](https://github.com/aeternity/aepp-sdk-go/commit/1c0ffdc21fd9397a63bdda214bc27cb99d70b7cc))
* tx structs, unittests, helpers, config, api updated to support v3.0.1 types ([e90ecfc](https://github.com/aeternity/aepp-sdk-go/commit/e90ecfc369588fb4b85d55c2a566f14f3fe9c83f))
* go-swagger v3.0.1 generated code ([63844f9](https://github.com/aeternity/aepp-sdk-go/commit/63844f9d11bc4563f63ce908d920855b6505a9c8))
* aeternity node v3.0.1 swagger.json ([2bd3ca0](https://github.com/aeternity/aepp-sdk-go/commit/2bd3ca0fb2d2277aa1d7403899923e862b5ad7c8))


<a name="v2.2.0"></a>
## [v2.2.0](https://github.com/aeternity/aepp-sdk-go/compare/v2.1.0...v2.2.0) (2019-05-27)

### Bug Fixes

* tx verify was using blank env variable as a default instead of the built in default. TX verification works now ([2106cae](https://github.com/aeternity/aepp-sdk-go/commit/2106cae5e0ff76dbf110b449b8be0dff553eb9da))
* cmd: tx spend queries node for Nonce if not specified. also accepts text payload ([b2879fb](https://github.com/aeternity/aepp-sdk-go/commit/b2879fbfedc46f66be8363dc196b45bdc8dab163))
* CLI --json flag now works ([911392d](https://github.com/aeternity/aepp-sdk-go/commit/911392d70ee931cc3b073afad09373df41e49455))
* terminal pretty printer also prints BigInt's value/amount ([3f239a1](https://github.com/aeternity/aepp-sdk-go/commit/3f239a15c6e6fbb3f2625d254e82938da86cb6bc))
* calcFeeContract wasn't paying attention to the base gas multiplier ([934ef00](https://github.com/aeternity/aepp-sdk-go/commit/934ef0003865ecdf6e54116d7347c4fd051133b6))
* swagger.json get_generation_by_height should be uint64 (manual fix) ([a5b5a87](https://github.com/aeternity/aepp-sdk-go/commit/a5b5a87107a60f987b6bd107802369a208cedcb0))
* make no_implicit_int64() filter print out what it did ([efdb963](https://github.com/aeternity/aepp-sdk-go/commit/efdb9631b4fd7958b76049ddd61b9141728913f4))
* generated.models.Error is patched to satisfy the Error interface. This dereferences the error.Reason pointer if the models.Error is passed as a pointer ([5806eb9](https://github.com/aeternity/aepp-sdk-go/commit/5806eb946db865050ff4c2d7eb25a4ce117149b3))
* unittest for GenericTx deserialization ([d293f30](https://github.com/aeternity/aepp-sdk-go/commit/d293f30b720b1ba7cf3605eda9e28a7f02f0c156))
* GenericTx was not deserializing properly into go-swagger models. generated code was not parsing node responses that involved GenericTx. First fixed in 2cccbeaefd49e8a037af5b4b2029d697477e683b but resurfaced after regenerating from new swagger.json make go-swagger generated *Tx models return SpendTx, not SpendTxJSON closes [#54](https://github.com/aeternity/aepp-sdk-go/issues/54) ([4716a8e](https://github.com/aeternity/aepp-sdk-go/commit/4716a8e64ce030e04eeddf1ff7120cda1bc8b18e))
* canonical NameUpdateTx RLP serialization seems to want NamePointers in the reverse order than specified in the JSON. ([f2bc093](https://github.com/aeternity/aepp-sdk-go/commit/f2bc09355f2549d2176a809827594e4c9e25a17d))

### Chore

* break integration tests up, make AENS test use random names ([571c9c2](https://github.com/aeternity/aepp-sdk-go/commit/571c9c22dd67a47a23d1c7f9035b4cd38e75ad4c))
* Transactions tests should test with fixed values and Config Defaults ([3686578](https://github.com/aeternity/aepp-sdk-go/commit/3686578d83b8f5f2633b32fb1ce6b7dc8923d1f9))
* AENS integration test now includes NameTransfer, NameRevoke ([7f1c37b](https://github.com/aeternity/aepp-sdk-go/commit/7f1c37bddaf5c32c0c86d3852181b9dd3d4a7bab))
* NameUpdateTx unittest with 4 pointers, just in case ordering issues pop up ([14fa242](https://github.com/aeternity/aepp-sdk-go/commit/14fa242621d13df530fb6d76dc8d35f692b7e7c2))
* another python tool to quickly edit _tx_json.go files ([2794c74](https://github.com/aeternity/aepp-sdk-go/commit/2794c74e9be57e91b66d50611483b6928c33f668))
* improved tools to rewrite swagger.json ([ed9e6bd](https://github.com/aeternity/aepp-sdk-go/commit/ed9e6bd2f607ccd5d49233952ae7ead2513cf62f))
* standardize field names across transactions ([07aa8a7](https://github.com/aeternity/aepp-sdk-go/commit/07aa8a77e5cc0ebe4cf6c6f5b97d9d3dd64f9807))
* cleanup - comments and casing ([3169ca7](https://github.com/aeternity/aepp-sdk-go/commit/3169ca74c2dae188ef8fc9d76b7a3b9cbc89e02f))
* Oracle Integration Test, OracleRegister, OracleExtend ([df5bef7](https://github.com/aeternity/aepp-sdk-go/commit/df5bef7610682bf85489f97664398a30d5c34e5f))
* small unittest to test that generated models.Error has its Reason automatically dereferenced when printing it ([1f0df4d](https://github.com/aeternity/aepp-sdk-go/commit/1f0df4d2147afca366f072a9f868ec5eff4d7e42))
* cleanup obsolete code ([ef86705](https://github.com/aeternity/aepp-sdk-go/commit/ef86705d78fc50fe62c8fe9d3d0e216b9ceade81))
* AENS integration test generalize to name as a variable, do not use stdin n as prompt to continue. use aeternity helpers instead ([cb5e252](https://github.com/aeternity/aepp-sdk-go/commit/cb5e252780a83bed052b897e833a221566695705))
* AENS workflow integration test ([db9ad5c](https://github.com/aeternity/aepp-sdk-go/commit/db9ad5c216eff1bdbcbc8ab3069c342fb12a8414))
* test genesis account is now richer in private testnet for integration tests ([897e78b](https://github.com/aeternity/aepp-sdk-go/commit/897e78bfe2b4b98f6e843c68fb7676575f2248a2))

### Code Refactoring

* separate AENS/Oracle/Contract helpers from aeternity node HTTP client. Also renamed them. Now, 1 helper = 1 account, and they can share the same http client instance ([2174b08](https://github.com/aeternity/aepp-sdk-go/commit/2174b089756a0956a23bb1cbfa44b9784c614a70))
* renaming to prepare for Client/Helper change ([ac47d04](https://github.com/aeternity/aepp-sdk-go/commit/ac47d04d5cc5ba687dfece6c8a5e8ee298f4d4ab))
* no need for SpendTxnormal vs SpendTxLarge anymore ([0fe777e](https://github.com/aeternity/aepp-sdk-go/commit/0fe777e1ee30ad8a33dfc658133c612772af7221))
* ContractConfig - Amount added, Gas/GasPrice as BigInt ([3ab86f5](https://github.com/aeternity/aepp-sdk-go/commit/3ab86f538cc78db2849bfdc25a569fa97353fc00))
* Contract transactions now use utils.BigInt for Gas ([61df44d](https://github.com/aeternity/aepp-sdk-go/commit/61df44d91a2bb08f1dcf29206eb8fed1101db1d5))
* Gas, GasPrice now in utils.BigInt instead of uint64. updatedict.py changed to do this too. edit.py no longer used, dropping. ([bd5e5a6](https://github.com/aeternity/aepp-sdk-go/commit/bd5e5a654f85d3ebb1c5192b6a09e3e66860578d))
* fee estimation ([b6f1054](https://github.com/aeternity/aepp-sdk-go/commit/b6f10545f58ee65b75bc455bcd9056575a9ed9a2))
* rename TxTTL, TxFee back to Fee, TTL so that all Txs have a standard name. This might be useful for interfaces later ([1aa0418](https://github.com/aeternity/aepp-sdk-go/commit/1aa0418725161440d9a7037cc4225b85054fed5b))
* Config includes BaseGas, GasPerByte, GasPrice for all Transaction types, not just Oracle/Contract ([890267e](https://github.com/aeternity/aepp-sdk-go/commit/890267e588d5a07fe8aebaf080c388b2312c40c0))
* Tx interface does not include JSON() method. This was more trouble than it was worth ([5d59dca](https://github.com/aeternity/aepp-sdk-go/commit/5d59dcac7de1ace14845482cd96ccafb8136be8e))
* NamePointer struct now embeds models.NamePointer. This gives automatic JSON serialization ([4729a9e](https://github.com/aeternity/aepp-sdk-go/commit/4729a9ed93d4e93a971d63aa50a804c3354285b7))
* NamePointer is now a struct of its own ([a2de13a](https://github.com/aeternity/aepp-sdk-go/commit/a2de13a3c17636bb1c8b7f9772094587053c4fce))

### Docs

* TODO consider refactoring WaitForTransactionUntilHeight ([03af1ab](https://github.com/aeternity/aepp-sdk-go/commit/03af1aba6b648d8c9aeb5d6cbcc8c41173484ca9))
* added comment for NamePointer.EncodeRLP() ([0788935](https://github.com/aeternity/aepp-sdk-go/commit/07889351d723e4c48fc3b01497312b5f8a7b25f7))
* add minor comments on exported functions ([df01f3e](https://github.com/aeternity/aepp-sdk-go/commit/df01f3ec24f1d49204177eb36be21ed5228c007a))

### Feature

* Fee estimation for SpendTx, AENS Txs, and Contract Txs ([225da65](https://github.com/aeternity/aepp-sdk-go/commit/225da659eb1441f8e384c91f46027f1b1fe7577d))

### Features

* helpers for Contract transactions ([92a1ed8](https://github.com/aeternity/aepp-sdk-go/commit/92a1ed8da4d4a2198a1c8461b33f6ee23a21480f))
* prototype Contract Integration Test ([07cb5df](https://github.com/aeternity/aepp-sdk-go/commit/07cb5dfd88868104ad86ff37ec7f74b90047adf7))
* Contract Call Tx works ([ac88116](https://github.com/aeternity/aepp-sdk-go/commit/ac88116c47063d78296541ec8e7ff3c7205cd491))
* APIGetContractByID (needed for integration test) ([1bf0497](https://github.com/aeternity/aepp-sdk-go/commit/1bf049775494ccb16bfad51bbc15718a6c864b90))
* Contract ID attached to ContractCreateTx struct, encodeVMABI function and unittests, ContractCreateTx unittests test common VM/ABI combinations ([9bf402a](https://github.com/aeternity/aepp-sdk-go/commit/9bf402a97a973195bceced0d0f74326291921df5))
* calculate contract ID ([a072966](https://github.com/aeternity/aepp-sdk-go/commit/a072966de3c1d8187c55781036d33f9ea87694b8))
* Fee Estimation for SpendTx, Name*Tx ([8cff1a8](https://github.com/aeternity/aepp-sdk-go/commit/8cff1a81ef77d5b7fe34175b12d7306c7dfaa2b9))
* ContractCallTx, not working yet ([bb74df6](https://github.com/aeternity/aepp-sdk-go/commit/bb74df6c163a9f6a96cf01122cceef2e3b18ff1d))
* ContractCreateTx works ([74ac97a](https://github.com/aeternity/aepp-sdk-go/commit/74ac97ad7d73b0674b5158a3a92b9705245e722d))
* add the aesophia compiler image to the docker-compose ([d58baa2](https://github.com/aeternity/aepp-sdk-go/commit/d58baa20699c5ca344c54a120556def7fdc7e652))
* generated code for v2.4.0 swagger.json ([6e79cf2](https://github.com/aeternity/aepp-sdk-go/commit/6e79cf2a4d77c57ec593cad68be51d1c250480f2))
* support aeternity v2.4.0 node ([5b0f04f](https://github.com/aeternity/aepp-sdk-go/commit/5b0f04fc294932965ac23b2c8c5285054e9b421e))
* helper functions for NameTransferTx, NameRevokeTx ([4f5ba90](https://github.com/aeternity/aepp-sdk-go/commit/4f5ba90c95fc483145cc829f33a6935b97a006e2))
* NameRevokeTx, NameTransferTx and their unittests ([1c9eb4e](https://github.com/aeternity/aepp-sdk-go/commit/1c9eb4e46f557842824cd31df14b7f224aa70c07))
* OracleRespondTx integration test ([aa9a742](https://github.com/aeternity/aepp-sdk-go/commit/aa9a7424240699208cf7507449f2a6c355756395))
* OracleRespondTx struct, RLP() and JSON() serialization and unittests ([4fd8171](https://github.com/aeternity/aepp-sdk-go/commit/4fd8171b2a368e8fbd5e43529f74d1cc3933c87c))
* OracleRespondTx prerequisites: APIGetOracleQueriesByPubkey(), buildOracleQueryID(), leftPadByteSlice() and unittests ([db05b75](https://github.com/aeternity/aepp-sdk-go/commit/db05b75cfaac52402cee55979d090bb721272390))
* OracleQueryTx works. Unittest reference value taken from node, not JS SDK. Integration test updated. ([24f5f87](https://github.com/aeternity/aepp-sdk-go/commit/24f5f874852b5a6b6fb531bd7c8b0c102f492ece))
* Oracle helper functions for OracleRegisterTx, OracleExtendTx ([0964c8c](https://github.com/aeternity/aepp-sdk-go/commit/0964c8cc0a1801f83a9befada0c618e785d97965))
* account address will ask for confirmation before printing private key ([5de88b5](https://github.com/aeternity/aepp-sdk-go/commit/5de88b5425309aab6eae8d6cee03efacbb346476))
* vanity account search takes regular expressions that can match anywhere ([e119c1d](https://github.com/aeternity/aepp-sdk-go/commit/e119c1d51baad95ca198abd7617e27f964e68b3a))
* AENS integration test now tests NamePreclaim, NameClaim, NameUpdate for simple ak_ type pointers ([d675cf2](https://github.com/aeternity/aepp-sdk-go/commit/d675cf2b12380e816409e1c132c575a3775283ba))

### Test

* Contracts Integration Test, working ([309363d](https://github.com/aeternity/aepp-sdk-go/commit/309363d5f6bd2f574bfa02e72e77dc7890b67385))
* TestContractCallTx_FeeEstimate makes sure that the fee estimate is working properly ([d750d14](https://github.com/aeternity/aepp-sdk-go/commit/d750d14a47981bdbedf9ed82ca3fa9c05620c3e3))


<a name="v2.1.0"></a>
## [v2.1.0](https://github.com/aeternity/aepp-sdk-go/compare/v2.0.0...v2.1.0) (2019-04-26)

### Bug Fixes

* generated.models.Error is patched to satisfy the Error interface. This dereferences the error.Reason pointer if the models.Error is passed as a pointer ([14ce6ed](https://github.com/aeternity/aepp-sdk-go/commit/14ce6ed347333f1ee0888cdebcc1f031aa74af7f))
* unittest for GenericTx deserialization ([1c17e40](https://github.com/aeternity/aepp-sdk-go/commit/1c17e402c31eb0ff9a84e27f596e8b10375946dd))
* GenericTx was not deserializing properly into go-swagger models. generated code was not parsing node responses that involved GenericTx. First fixed in 2cccbeaefd49e8a037af5b4b2029d697477e683b but resurfaced after regenerating from new swagger.json make go-swagger generated *Tx models return SpendTx, not SpendTxJSON closes [#54](https://github.com/aeternity/aepp-sdk-go/issues/54) ([b2b9cd5](https://github.com/aeternity/aepp-sdk-go/commit/b2b9cd536f70ac8cf52886b75ec819a6a8e28938))
* canonical NameUpdateTx RLP serialization seems to want NamePointers in the reverse order than specified in the JSON. ([cc6a56e](https://github.com/aeternity/aepp-sdk-go/commit/cc6a56edd33e5505a49e92c7664aa11bf868c735))

### Chore

* standardize field names across transactions ([aaca2e5](https://github.com/aeternity/aepp-sdk-go/commit/aaca2e58d461b49f3393f0caf3822aa1666b042c))
* cleanup - comments and casing ([21b68d9](https://github.com/aeternity/aepp-sdk-go/commit/21b68d991fd6c39b0ef46215740c24625a4785d7))
* Oracle Integration Test, OracleRegister, OracleExtend ([46a4192](https://github.com/aeternity/aepp-sdk-go/commit/46a41920731e847e358f9b8b9747ba2a23e8e9a7))
* small unittest to test that generated models.Error has its Reason automatically dereferenced when printing it ([839c5b2](https://github.com/aeternity/aepp-sdk-go/commit/839c5b2ed7505babbb2a4c42b390e8140bacde1d))
* cleanup obsolete code ([053837c](https://github.com/aeternity/aepp-sdk-go/commit/053837cf7f801e41fd8cad477bc6777ecb244c53))
* AENS integration test generalize to name as a variable, do not use stdin n as prompt to continue. use aeternity helpers instead ([ffc4f8a](https://github.com/aeternity/aepp-sdk-go/commit/ffc4f8ab65cced275827afe15d1a3f672ea9d697))
* AENS workflow integration test ([152947c](https://github.com/aeternity/aepp-sdk-go/commit/152947c4ff9cce110391859045f05185b56e3007))
* test genesis account is now richer in private testnet for integration tests ([91dff09](https://github.com/aeternity/aepp-sdk-go/commit/91dff0937b2fde2e963341661aa747ee916f5b5c))

### Code Refactoring

* Tx interface does not include JSON() method. This was more trouble than it was worth ([f79c27b](https://github.com/aeternity/aepp-sdk-go/commit/f79c27b399d08ddb6520ca4db453d9f69cf7dfd8))
* NamePointer struct now embeds models.NamePointer. This gives automatic JSON serialization ([ccff0a6](https://github.com/aeternity/aepp-sdk-go/commit/ccff0a6944b8f6692ee6e948dd970a2a2b2bfe0c))
* NamePointer is now a struct of its own ([9bd57e9](https://github.com/aeternity/aepp-sdk-go/commit/9bd57e9a68a00c436f0186e863c3eb7d4c630172))

### Docs

* TODO consider refactoring WaitForTransactionUntilHeight ([b8a4f0c](https://github.com/aeternity/aepp-sdk-go/commit/b8a4f0c4fd111aea9cd2fc19a93d62133eb74f0f))
* added comment for NamePointer.EncodeRLP() ([72e28be](https://github.com/aeternity/aepp-sdk-go/commit/72e28be32700e012689ba8d310e370b21af72dfc))
* add minor comments on exported functions ([8eb5d88](https://github.com/aeternity/aepp-sdk-go/commit/8eb5d88b9b8d46e00450d26472a151a0f916d2de))

### Features

* OracleRespondTx integration test ([d7c5696](https://github.com/aeternity/aepp-sdk-go/commit/d7c5696900b376f088f094579f7ee644eae7a0d1))
* OracleRespondTx struct, RLP() and JSON() serialization and unittests ([680262a](https://github.com/aeternity/aepp-sdk-go/commit/680262a397f711d0bb9bcd1025de1192fe18be79))
* OracleRespondTx prerequisites: APIGetOracleQueriesByPubkey(), buildOracleQueryID(), leftPadByteSlice() and unittests ([c656fb4](https://github.com/aeternity/aepp-sdk-go/commit/c656fb487c360ba297eac19e7870ab07f082447d))
* OracleQueryTx works. Unittest reference value taken from node, not JS SDK. Integration test updated. ([346f8f3](https://github.com/aeternity/aepp-sdk-go/commit/346f8f323a649b16a31225eda3fdc542dd1d4a5e))
* Oracle helper functions for OracleRegisterTx, OracleExtendTx ([73f908c](https://github.com/aeternity/aepp-sdk-go/commit/73f908c72521f35e8ef7eff5d4b6d11b93ec2075))
* account address will ask for confirmation before printing private key ([cb80311](https://github.com/aeternity/aepp-sdk-go/commit/cb80311c20bd70ae178b8805ec79798f53844ea0))
* vanity account search takes regular expressions that can match anywhere ([9612610](https://github.com/aeternity/aepp-sdk-go/commit/961261084b3a11aeecc253660cc013ec0fcecdf9))
* AENS integration test now tests NamePreclaim, NameClaim, NameUpdate for simple ak_ type pointers ([f31085e](https://github.com/aeternity/aepp-sdk-go/commit/f31085e5a3b5dbb73d063e40d18c071a27a5c45b))


<a name="v2.0.0"></a>
## [v2.0.0](https://github.com/aeternity/aepp-sdk-go/compare/v1.0.2...v2.0.0) (2019-04-09)

### Bug Fixes

* NameClaimTx salt 256bit bytearray was being converted to uint64 and back, which mangles the bytes. Solved by making it a BigInt, which has proper 32 byte big endian representation with Bytes() ([45ff92d](https://github.com/aeternity/aepp-sdk-go/commit/45ff92da01bdedc8e7d82c943c49c17c0de63bba))
* disabled OracleQueryTx, aepp-sdk-js reference RLP serialization is more reliable than aepp-sdk-python atm ([88e7ac4](https://github.com/aeternity/aepp-sdk-go/commit/88e7ac456e128cfbbc8a357315ebe7b5ac7b2d00))
* Aens.NameClaimTx helper should pass the plaintext name down to the actual function ([c976216](https://github.com/aeternity/aepp-sdk-go/commit/c976216ce70335fd202215044fb025b3351fabac))
* NameClaimTx correct serialization in JSON and RLP formats ([0e6799d](https://github.com/aeternity/aepp-sdk-go/commit/0e6799d5d4a6769d1049f0dce45162b6451113f8))
* AENS.UpdateFee specified, made into a BigInt ([de42965](https://github.com/aeternity/aepp-sdk-go/commit/de429656195f2b9581329d0a9808d5bf6c88081a))
* integration test was broken after BigInt.Int was changed from copy-by-value to pointer, because swagger did not know it had to initialize BigInt.Int when creating a BigInt. Fixed. ([b9dd2f3](https://github.com/aeternity/aepp-sdk-go/commit/b9dd2f3f12c3c9e6ce1dce2d8debd54d6c541613))
* OracleQueryTx.RLP() was not correct, renamed members of OracleRespondTx ([b679804](https://github.com/aeternity/aepp-sdk-go/commit/b679804138246754f6abbb68368d5536a9adebef))
* OracleExtendTx.RLP() use t.Fee.Int instead of t.Fee.Bytes() for RLP serialization ([857338f](https://github.com/aeternity/aepp-sdk-go/commit/857338f3088aa6f27c3ec380b11b68e3b21effa7))
* integration test now works with BigInt's pointer receivers ([d373565](https://github.com/aeternity/aepp-sdk-go/commit/d373565eaa2e26cd1b62f1ef59ef3114e6f633b6))
* tx command shouldn't use BaseEncodeTx() as a returned function ([21c5a7c](https://github.com/aeternity/aepp-sdk-go/commit/21c5a7c1732368631963c961c5f8587fc37ee745))
* SpendTx, OracleRegisterTx RLP() methods should always use utils.BigInt.Int while serializing, not the utils.BigInt directly, because otherwise there will be a list within a list ([fcaa28d](https://github.com/aeternity/aepp-sdk-go/commit/fcaa28da23b443c2c82f58592b03dc2580196121))
* rlp Encode() was encoding 0 value of uint64 as 0x00, but big.Int 0 value as 0x80. Changed big.Int 0 value to 0x00 ([e21c6f5](https://github.com/aeternity/aepp-sdk-go/commit/e21c6f5d635e7ca13041d729a9619d8821d030f2))
* Ae.WithAccount() was not copying the Account object to the Oracle struct ([ef9eab6](https://github.com/aeternity/aepp-sdk-go/commit/ef9eab68e47a912495b88c1191238ab0a0940358))
* renamed helper function arguments to accomodate different kinds of TTLs in other transaction types ([5697537](https://github.com/aeternity/aepp-sdk-go/commit/5697537685e72aa3032ea7bff2bceb9e047ef345))
* rearranged function arguments and correct param types in OracleRegisterTx() ([2b00541](https://github.com/aeternity/aepp-sdk-go/commit/2b00541eea5b147ae61a7d5e66918bfadfe31a08))
* uint64/utils.BigInt instead of int64 for certain variables in config.go ([4c76f06](https://github.com/aeternity/aepp-sdk-go/commit/4c76f064735070550e5632ed5a19f3e298ad4852))

### Chore

* separate salt generation from calculating the commitmentId, so we can unittest commitmentId generation against GET /v2/debug/names/commitment-id ([4881ad9](https://github.com/aeternity/aepp-sdk-go/commit/4881ad9211360da8b2256fdb05a75614aa6b7291))
* tx dumpraw unittest ([cef5b99](https://github.com/aeternity/aepp-sdk-go/commit/cef5b9990f4c2956536ca66e556356018dc3691e))
* unittests for OracleExtendTx ([c6ca727](https://github.com/aeternity/aepp-sdk-go/commit/c6ca72798cbff34dbf9aeb0e601fb80d4c4a19e8))
* **release:** 2.0.0 ([7852cf0](https://github.com/aeternity/aepp-sdk-go/commit/7852cf075e5e307c4ee3cdb0f4ec1d7fe8785a55))

### Cleanup

* deleted handwritten unittests, use new generated ones. More unittests that reliably check for RLP serialization. getRLPSerialized() is a convenience function that quickly gives you the RLP deserialized representation ([fb9dc82](https://github.com/aeternity/aepp-sdk-go/commit/fb9dc82087a45c1788b6806374a6b9c9673db645))
* transaction unittest should not be in helpers, whose role needs to be redefined anyway. New OracleRegisterTx unittest ([b5230e6](https://github.com/aeternity/aepp-sdk-go/commit/b5230e6af202b4929bc9cace129bf2d243d3f5a9))
* helpers_test.go should use new BigInt constructor methods ([1869bc7](https://github.com/aeternity/aepp-sdk-go/commit/1869bc7ef0e4c280359281478e0c66dbb0f41e9a))
* fix minor typo in node.go ([7cd0bd5](https://github.com/aeternity/aepp-sdk-go/commit/7cd0bd57e9ca960f6628730d59ca263751014112))

### Code Refactoring

* helper functions now return the Tx structs, which then need to be separately base64 encoded. However, TTL and Nonce are automatically determined, because that was just tedious ([40fa167](https://github.com/aeternity/aepp-sdk-go/commit/40fa1679c42e170da4cc88ab2d29aafe2e63e8c8))
* Namehash() exported because it is useful when you need to debug AENS ([d4ac51d](https://github.com/aeternity/aepp-sdk-go/commit/d4ac51db38e341e9390a18a6f22633158694970a))
* AENS Tx fees should be denominated in BigInt ([db198cf](https://github.com/aeternity/aepp-sdk-go/commit/db198cf3f43bcf6b31ba0ef347e236b57bffd7a2))
* transactions_test does not have to be an aeternity_test package - this creates more problems for now. Leave aeternity_test to code that tests the user convenience stuff above this level ([9b0cc65](https://github.com/aeternity/aepp-sdk-go/commit/9b0cc65c98f0f542e652926ed7c7a69771e16a82))
* utils.BigInt should embed a *big.Int after all, not a big.Int. Also, its methods should have pointer receivers, not value receivers. This lets us use big.Int's methods, since they are defined on the pointer see https://stackoverflow.com/questions/55337900/marshaljson-a-type-with-an-embedded-type-ends-up-as-instead-of-the-value/55338067[#55338067](https://github.com/aeternity/aepp-sdk-go/issues/55338067) ([218578a](https://github.com/aeternity/aepp-sdk-go/commit/218578a809e7e1dcb57373e1e05401150a8fd2ce))
* integration test uses new Tx struct method ([b9dcdc1](https://github.com/aeternity/aepp-sdk-go/commit/b9dcdc16939e7a00cfc04b53c762269dd1637023))
* Tx struct methods now use a pointer receiver Also cleaned up unittests ([fe513d8](https://github.com/aeternity/aepp-sdk-go/commit/fe513d8b6d6a7d1d285086213514a64de9e7ac01))
* Tx structs should have exported fields ([42d94ac](https://github.com/aeternity/aepp-sdk-go/commit/42d94ac15c20c1f2c38f755b959f8649d8a68ad2))
* BaseEncodeTx() doesn't need to return a function - it can just convert within itself ([533ae43](https://github.com/aeternity/aepp-sdk-go/commit/533ae43654f701cc420aee35320c9fedcf41791b))
* many helper functions not needed anymore. Commenting out to focus on other issues ([69bef33](https://github.com/aeternity/aepp-sdk-go/commit/69bef33e898919c51446fe300ccd9de8eaa45eb2))
* cmd/tx.go uses new refactored SpendTx struct ([f48377d](https://github.com/aeternity/aepp-sdk-go/commit/f48377dca2b2f40c3daf3b96a0987a8fe186c75d))
* made Transaction functions Go-ish, with structs and interfaces. More unittests, but some failing due to big.Int 0-value serialization issue ([99ca385](https://github.com/aeternity/aepp-sdk-go/commit/99ca385266b8c92510ed70bffe33fa40eed62e91))
* DecodeRLPMessage() exported. Now we can have deep debugging functions from the cmd package, for example ([31de960](https://github.com/aeternity/aepp-sdk-go/commit/31de9601653b7281bde0f587adfe68780dfda5a0))
* BroadcastTransaction() helper becomes part of Ae struct, so aeClient.BroadcastTransaction() ([1f1652c](https://github.com/aeternity/aepp-sdk-go/commit/1f1652cb8b9b18f699c40d75f672ca996e305fd2))
* Ae.GetTTL/Nonce() is for SDK user convenience - under the hood, they use unexported normal functions getTTL/getNonce() ([7cd8679](https://github.com/aeternity/aepp-sdk-go/commit/7cd8679b50d35ef1ed8df231983d56e2c331991e))

### Docs

* show how to use the SDK ([7895e74](https://github.com/aeternity/aepp-sdk-go/commit/7895e744cacc0027bc745622ac5bdad295b252ec))
* comments for JSON() methods ([d1a5f64](https://github.com/aeternity/aepp-sdk-go/commit/d1a5f644d995fd9a40e4fd0cec3b3efc9150adbf))
* explain the TTL types in AENS ([9265779](https://github.com/aeternity/aepp-sdk-go/commit/9265779a39b7f6c9336f50f0e36bde5bbb782555))

### Feature

* fleshing out NameUpdateTx, especially the JSON and Pointers overhaul. buildPointers is completely different now. Still does not work 100% ([6d460af](https://github.com/aeternity/aepp-sdk-go/commit/6d460af2af5aa15e7f09dd12e3fe0d32534b7353))

### Features

* account vanity generator chore: chain broadcast unittest ([3b5c5bd](https://github.com/aeternity/aepp-sdk-go/commit/3b5c5bd47db127cbd0b54c26d52aeb896e861de5))
* OracleQueryTx struct revision, unittest structure (but no working test) ([41f9821](https://github.com/aeternity/aepp-sdk-go/commit/41f9821f1df415a4de37e293188f0a4a77a5b47e))
* OracleQueryTx, OracleRespondTx structs ([be508b1](https://github.com/aeternity/aepp-sdk-go/commit/be508b114cbaa8a6038c2a31bcb4078a0b6437db))
* AENS transactions in new struct format, given JSON() methods, some types fixed. Tx interface now includes the JSON() method ([edabb99](https://github.com/aeternity/aepp-sdk-go/commit/edabb99178c92c4867b18584285fffeb9f93fe3c))
* Tx structs now feature a JSON() serialization that uses swagger models ([944daa7](https://github.com/aeternity/aepp-sdk-go/commit/944daa70ccd0abcdd52835079c8598db3c389217))
* CLI tx dumpraw command (helps with debugging) ([05516cd](https://github.com/aeternity/aepp-sdk-go/commit/05516cd87b50fef8246c03c7d3f97342e28fae83))
* OracleRegisterTxStr() helper function ([a654936](https://github.com/aeternity/aepp-sdk-go/commit/a654936622d2769a4c22e95485b858fce5f52993))
* OracleRegisterTx, OracleExtendTx. Incomplete unittest ([f41a021](https://github.com/aeternity/aepp-sdk-go/commit/f41a0210080500a65c70f684af221f59c01edda2))

### Style

* rename nodeCli to nodeClient - less confusing ([f1aaf15](https://github.com/aeternity/aepp-sdk-go/commit/f1aaf15eaef8d129226cdc9e9e2e436171b83dca))


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/aeternity/aepp-sdk-go/compare/0.25.0-0.1.0-alpha...v1.0.2) (2019-03-20)

### AENS

* fees and salts are now uint64 ([09c2ea4](https://github.com/aeternity/aepp-sdk-go/commit/09c2ea463617453ddeb8f5ee5781a99f53279c19))

### Aeternity

* Base58/64 Encode/Decode functions now exported ([4a17065](https://github.com/aeternity/aepp-sdk-go/commit/4a17065c9d2009bef9d65572860373646e57f301))

### Aeternity_test

* TestSpendTransactionWithNode integration test ([084e9f3](https://github.com/aeternity/aepp-sdk-go/commit/084e9f38854dc21cbdc51a147bc510072c46063d))

### Api

* added design comment on Ae.API*() methods ([5762cbf](https://github.com/aeternity/aepp-sdk-go/commit/5762cbf33e1c577ebc99acf99b5a35b39ca52be7))

### Bug Fixes

* renamed helper function arguments to accomodate different kinds of TTLs in other transaction types ([2970997](https://github.com/aeternity/aepp-sdk-go/commit/297099784a95b3339af4c9326a790a32aff998e4))
* rearranged function arguments and correct param types in OracleRegisterTx() ([7c47342](https://github.com/aeternity/aepp-sdk-go/commit/7c47342fb7a8320a34eed985db3c056fe4256bd4))
* uint64/utils.BigInt instead of int64 for certain variables in config.go ([1c65c9a](https://github.com/aeternity/aepp-sdk-go/commit/1c65c9a3ef0b2f5faa0e38fcbbf4119a18f5731e))

### Chain

* added ttl and networkid command ([a25c0a9](https://github.com/aeternity/aepp-sdk-go/commit/a25c0a93f117ff7ee136ba4c41779c103aaf60d0))

### Chore

* helpers_test.go should use new BigInt constructor methods ([a0325fc](https://github.com/aeternity/aepp-sdk-go/commit/a0325fc9d64ce8b8b109a3f52cc7ce954e93517d))
* fix minor typo in node.go ([5b80101](https://github.com/aeternity/aepp-sdk-go/commit/5b801014e86261ea0a6a2f4d3ceded5f9d9997a4))

### Cmd

* tx spend better help message for --nonce ([226427e](https://github.com/aeternity/aepp-sdk-go/commit/226427e3f71bc636c93da377e8ee300fffb1b79d))
* tx spend accepts ttl, nonce TODO: ACTUALLY IMPLEMENT IT ([109f432](https://github.com/aeternity/aepp-sdk-go/commit/109f43202ad28c9dd990bbded5f9bc686ce470e0))
* removed intermediate nodeExternalURL variable. Now the single source of truth for Node URL and NetworkID is aeternity.Config.Epoch. Until helper functions accept these via arguments, this allows dependent code to change Node URL and NetworkID as well without worrying that it might be touched elsewhere. ([a340ec1](https://github.com/aeternity/aepp-sdk-go/commit/a340ec1a1fbdcdfacc904e6096d7f393b1d4a09f))
* amount, fee are now always processed as uint64 ([4520db9](https://github.com/aeternity/aepp-sdk-go/commit/4520db95bb3007f79adea2ded4c9995098143aca))
* account balance unittest for when an account was found, and not found. Need to rework terminal.go to unwrap model.Error messages properly ([5ba5fba](https://github.com/aeternity/aepp-sdk-go/commit/5ba5fbaf6b7ea00282996b4cc89f7d5018ecbbb9))
* tx spend unittest reenable ([c939931](https://github.com/aeternity/aepp-sdk-go/commit/c9399310f4803f2fd67b1758ed3a9439b052283a))
* inspect unittest and error object ([d52eb1e](https://github.com/aeternity/aepp-sdk-go/commit/d52eb1e1a24dc996b18fc6cea31cdd14d4363069))
* account unittests use t.Error() instead of t.Fatal, renamed Test*() functions ([84719cd](https://github.com/aeternity/aepp-sdk-go/commit/84719cd7b27126b08483b14dc825e9338c218c37))
* chain use error return objects ([5093af2](https://github.com/aeternity/aepp-sdk-go/commit/5093af229e8310c038e81bbe2ed8b0ca856671df))
* tx unittests (failing) ([c396532](https://github.com/aeternity/aepp-sdk-go/commit/c396532f80f101c5a8cead0d5f9c62d4861e2d99))
* tx return error objects. Now that Config is hardcoded, I can reference it directly for the --fee argument ([bcea426](https://github.com/aeternity/aepp-sdk-go/commit/bcea4262d8f6fab3038153298c811e82060692af))
* tx uses named functions now ([52c6a10](https://github.com/aeternity/aepp-sdk-go/commit/52c6a109e261a991246121a42c9d9a10831e483f))
* account subcommands return errors. This helps in testing, and keeps it simple ([ccdf565](https://github.com/aeternity/aepp-sdk-go/commit/ccdf565c7b8b54b44ea80b6084c2c067683d3f16))
* separated functions from cobra commands, some unittests. --password now belongs to account command, not its subcommands ([6bd18f7](https://github.com/aeternity/aepp-sdk-go/commit/6bd18f76647ec50b57b86d8a7be469fab4066212))
* inspect unittest (MORE TESTS) ([f10e442](https://github.com/aeternity/aepp-sdk-go/commit/f10e442232d66bc2950fd39153f414eb4250deda))
* chain unittests ([721eca4](https://github.com/aeternity/aepp-sdk-go/commit/721eca4db5dbe7af4c2cff8efa98fc9e311a17d5))
* removed aeCli as a global variable, so each subcommand has to instantiate its own aeCli now. This helps go-vet detect problems where I would otherwise only get nil reference panics because aeCli wasn't initialized yet ([321b199](https://github.com/aeternity/aepp-sdk-go/commit/321b199ea109899dbdd26c748e5e1d43de100903))
* tx broadcast -> chain broadcast ([a11a62d](https://github.com/aeternity/aepp-sdk-go/commit/a11a62d4b5f4921fd5d56624e902e1d49eb68306))

### Code Refactoring

* BroadcastTransaction() helper becomes part of Ae struct, so aeClient.BroadcastTransaction() ([adb1b7f](https://github.com/aeternity/aepp-sdk-go/commit/adb1b7faec47df7b7fa6294685ee4ba8d38f1b05))
* Ae.GetTTL/Nonce() is for SDK user convenience - under the hood, they use unexported normal functions getTTL/getNonce() ([76a0ff6](https://github.com/aeternity/aepp-sdk-go/commit/76a0ff64a354be37cd41abb55bab03955fd36348))

### Config

* use sdk-mainnet.aepps.com ([68c4ff7](https://github.com/aeternity/aepp-sdk-go/commit/68c4ff7937f81135c17209635b62650cdae5f038))

### Features

* OracleRegisterTxStr() helper function ([a178c87](https://github.com/aeternity/aepp-sdk-go/commit/a178c8738fd40518af1576618d4c08552fbb4e6f))
* OracleRegisterTx, OracleExtendTx. Incomplete unittest ([7c7dbcb](https://github.com/aeternity/aepp-sdk-go/commit/7c7dbcb49909a9d0d7ee1fb1e03741879f8b7a40))

### Generated

* updated to 2.0.0 ([3e7114b](https://github.com/aeternity/aepp-sdk-go/commit/3e7114b8c5c14608f008f2051617906ee60fc585))

### Gitignore

* ignore docker-compose.override.yml ([5c23a85](https://github.com/aeternity/aepp-sdk-go/commit/5c23a85f682d95342b5df86bc625b740d41ab14b))

### Helpers

* GetNextNonce, GetTTL renamed and exported, SpendTransaction->SpendTxStr, signature changed. All helper functions require ttl, nonce as arguments ([3eb5e14](https://github.com/aeternity/aepp-sdk-go/commit/3eb5e140325840a93b826028fa9d7d2c52f3faf3))
* TestSpendTransaction ([7beb27f](https://github.com/aeternity/aepp-sdk-go/commit/7beb27fc152e6f132b0a6edda2146ec8c2deb3e9))

### README

* single line swagger generation command ([944a56c](https://github.com/aeternity/aepp-sdk-go/commit/944a56c8e42f7046acfc4c59fa3f83b72de5abbc))

### Rlp

* fixing 0 encoding issue ([991f2b4](https://github.com/aeternity/aepp-sdk-go/commit/991f2b4385af85985c2a1de356f4455498ec8ee1))

### Style

* rename nodeCli to nodeClient - less confusing ([528eba0](https://github.com/aeternity/aepp-sdk-go/commit/528eba032794778d3d75c90bc72d9196beb3853e))

### Swagger

* Custom Bigint ([9d3c604](https://github.com/aeternity/aepp-sdk-go/commit/9d3c6044dc5fa014d9ba531ff6b07c2433ba5c95))

### Utils

* BigInt implements UnmarshalJSON and MarshalJSON. Unittests included for peace of mind, although code does not actually do much. ([b6714c8](https://github.com/aeternity/aepp-sdk-go/commit/b6714c81fd32499d0fd284ff949be8042a42a14c))
* BigInt now uses big.Int as embedded struct. Validation unittest. ([5a4a67b](https://github.com/aeternity/aepp-sdk-go/commit/5a4a67b1f518e8a1f85efbd29f23c2ec5096aa49))


<a name="0.25.0-0.1.0-alpha"></a>
## [0.25.0-0.1.0-alpha](https://github.com/aeternity/aepp-sdk-go/compare/0.22.0-0.1.0-alpha...0.25.0-0.1.0-alpha) (2018-11-22)


<a name="0.22.0-0.1.0-alpha"></a>
## 0.22.0-0.1.0-alpha (2018-09-19)

