package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v5/utils"
)

// Acceptable values for various parameters
const (
	// NetworkIDMainnet is the network ID for aeternity mainnet
	NetworkIDMainnet = "ae_mainnet"
	// URLMainnet is the URL to an aeternity Foundation maintained node
	URLMainnet = "https://sdk-mainnet.aepps.com"
	// NetworkIDTestnet is the network ID for aeternity testnet
	NetworkIDTestnet = "ae_uat"
	// URLTestnet is the URL to an aeternity Foundation maintained node
	URLTestnet = "https://sdk-testnet.aepps.com"
	// CompilerBackendFATE indicates that the compiler should use the FATE VM
	// for contract bytecode execution
	CompilerBackendFATE = "fate"
	// CompilerBackendAEVM indicates that the compiler should use the AEVM for
	// contract bytecode execution
	CompilerBackendAEVM = "aevm"
	// OracleTTLTypeDelta indicates that the accompanying TTL value (in blocks)
	// should be interpreted as currentHeight + TTLValue
	OracleTTLTypeDelta = 0
	// OracleTTLTypeAbsolute indicates that the accompanying TTL value (in
	// blocks) should be interpreted as an absolute block height, after which
	// the TTL expires.
	OracleTTLTypeAbsolute = 1
)

// NodeConfig configuration for the node
type NodeConfig struct {
	URL         string `json:"url" yaml:"url" mapstructure:"url"`
	URLInternal string `json:"url_internal" yaml:"url_internal" mapstructure:"url_internal"`
	URLChannels string `json:"url_channels" yaml:"url_channels" mapstructure:"url_channels"`
	NetworkID   string `json:"network_id" yaml:"network_id" mapstructure:"network_id"`
}

// CompilerConfig configuration for the compiler
type CompilerConfig struct {
	URL     string `json:"url" yaml:"url" mapstructure:"url"`
	Backend string `json:"backend" yaml:"backend" mapstructure:"backend"`
}

// AensConfig contains default parameters for AENS
type AensConfig struct {
	// NameTTL is the block height (in this case not an absolute block height,
	// but a delta) after which the name goes into the 'revoked' state.
	NameTTL uint64 `json:"name_ttl" yaml:"name_ttl" mapstructure:"name_ttl"`
	// ClientTTL suggests how long (in seconds) AENS clients should cache an AENS entry.
	ClientTTL uint64 `json:"client_ttl" yaml:"client_ttl" mapstructure:"client_ttl"`
}

// ContractConfig contains default parameters for contracts
type ContractConfig struct {
	CompilerURL string `json:"compiler" yaml:"compiler" mapstructure:"compiler"`
	// GasLimit is a default value for the maximum amount of gas that a contract
	// execution should consume. see
	// https://github.com/aeternity/protocol/blob/master/consensus/consensus.md
	//
	// In order to control the size and the number of transactions in a micro
	// block, each transaction has a gas. The sum of gas of all the transactions
	// cannot exceed the gas limit per micro block, which is 6 000 000.
	// The gas of a transaction is the sum of:
	// * the base gas;
	// * other gas components, such as gas proportional to the byte size of the
	//   transaction or relative TTL, gas needed for contract execution.
	GasLimit big.Int `json:"gas" yaml:"gas" mapstructure:"gas"`
	// Amount is an optional amount to transfer to the contract account.
	Amount big.Int `json:"amount" yaml:"amount" mapstructure:"amount"`
	// Deposit will be "held by the contract" until it is deactivated.
	Deposit big.Int `json:"deposit" yaml:"deposit" mapstructure:"deposit"`
	// VMVersion indicates which virtual machine should be used for bytecode execution.
	VMVersion uint16 `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
	// ABIVersion indicates the binary interface/calling convention used by the contract.
	ABIVersion uint16 `json:"abi_version" yaml:"abi_version" mapstructure:"abi_version"`
}

// OracleConfig contains default parameters for oracles
type OracleConfig struct {
	// QueryFee is locked up until the oracle answers (and gets the fee) or the
	// transaction TTL expires (and the sender is refunded). In other words, it
	// is a bounty.
	QueryFee big.Int `json:"query_fee" yaml:"query_fee" mapstructure:"query_fee"`
	// QueryTTLType indicates whether the TTLValue should be interpreted as an absolute or delta blockheight.
	QueryTTLType uint64 `json:"query_ttl_type" yaml:"query_ttl_type" mapstructure:"query_ttl_type"`
	// QueryTTLValue indicates how long the query is open for response from the oracle.
	QueryTTLValue uint64 `json:"query_ttl_value" yaml:"query_ttl_value" mapstructure:"query_ttl_value"`
	// ResponseTTLType indicates whether the TTLValue should be interpreted as an absolute or delta blockheight.
	ResponseTTLType uint64 `json:"response_ttl_type" yaml:"response_ttl_type" mapstructure:"response_ttl_type"`
	// ResponseTTLValue indicates how long the response is available when given from the oracle.
	ResponseTTLValue uint64 `json:"response_ttl_value" yaml:"response_ttl_value" mapstructure:"response_ttl_value"`
	VMVersion        uint16 `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// StateChannelConfig configurations for contracts TODO: not complete
type StateChannelConfig struct {
	LockPeriod     uint64 `json:"lock_period" yaml:"lock_period" mapstructure:"lock_period"`
	ChannelReserve uint64 `json:"channel_reserve" yaml:"channel_reserve" mapstructure:"channel_reserve"`
}

// ClientConfig client parameters configuration
type ClientConfig struct {
	// BaseGas is one component of transaction fee calculation.
	BaseGas big.Int `json:"base_gas" yaml:"base_gas" mapstructure:"base_gas"`
	// GasPerByte is multiplied by the RLP serialized transaction's length.
	GasPerByte big.Int `json:"gas_per_byte" yaml:"gas_per_byte" mapstructure:"gas_per_byte"`
	// GasPrice is the conversion factor from gas to AE.
	GasPrice big.Int `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	// TTL is the default blockheight offset that will be added to the current
	// height to determine a transaction's TTL.
	TTL uint64 `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	// Fee is a default transaction fee that should be big enough for most transaction types.
	Fee           big.Int            `json:"fee" yaml:"fee" mapstructure:"fee"`
	Names         AensConfig         `json:"names" yaml:"names" mapstructure:"names"`
	Contracts     ContractConfig     `json:"contracts" yaml:"contracts" mapstructure:"contracts"`
	Oracles       OracleConfig       `json:"oracles" yaml:"oracles" mapstructure:"oracles"`
	StateChannels StateChannelConfig `json:"state_channels" yaml:"state_channels" mapstructure:"state_channels"`
}

// TuningConfig fine tuning of parameters of the client
type TuningConfig struct {
	ChainPollInterval int64  `json:"chain_poll_interval" yaml:"chain_poll_interval" mapstructure:"chain_poll_interval"`
	ChainTimeout      int64  `json:"chain_timeout" yaml:"chain_timeout" mapstructure:"chain_timeout"`
	CryptoKdfMemlimit uint32 `json:"crypto_kdf_memlimit" yaml:"crypto_kdf_memlimit" mapstructure:"crypto_kdf_memlimit"`
	CryptoKdfOpslimit uint32 `json:"crypto_kdf_opslimit" yaml:"crypto_kdf_opslimit" mapstructure:"crypto_kdf_opslimit"`
	CryptoKdfThreads  uint8  `json:"crypto_kdf_threads" yaml:"crypto_kdf_threads" mapstructure:"crypto_kdf_threads"`
	OutputFormatJSON  bool   `json:"-" yaml:"-" mapstructure:"-"`
}

// ProfileConfig a configuration profile
type ProfileConfig struct {
	Name     string         `json:"name" yaml:"name" mapstructure:"name"`
	Node     NodeConfig     `json:"node" yaml:"node" mapstructure:"node"`
	Compiler CompilerConfig `json:"compiler" yaml:"compiler" mapstructure:"compiler"`
	Client   ClientConfig   `json:"client" yaml:"client" mapstructure:"client"`
	Tuning   TuningConfig   `json:"tuning" yaml:"tuning" mapstructure:"tuning"`
}

// Config specifies defaults for all configuration parameters
var Config = ProfileConfig{
	Name: "Default Config",
	Node: NodeConfig{
		URL:         "https://sdk-mainnet.aepps.com",
		URLInternal: "https://sdk-mainnet.aepps.com",
		URLChannels: "https://sdk-mainnet.aepps.com",
		NetworkID:   "ae_mainnet",
	},
	Compiler: CompilerConfig{
		URL:     "http://localhost:3080",
		Backend: CompilerBackendAEVM,
	},
	Client: ClientConfig{
		BaseGas:    *utils.NewIntFromUint64(15000),
		GasPerByte: *utils.NewIntFromUint64(20),
		GasPrice:   *utils.NewIntFromUint64(1e9),
		TTL:        500,
		Fee:        *utils.RequireIntFromString("200000000000000"),
		Names: AensConfig{
			NameTTL:   500,
			ClientTTL: 500,
		},
		Contracts: ContractConfig{
			CompilerURL: "http://localhost:3080",
			GasLimit:    *utils.NewIntFromUint64(1e9),
			Amount:      *new(big.Int),
			Deposit:     *new(big.Int),
			VMVersion:   4,
			ABIVersion:  1,
		},
		Oracles: OracleConfig{
			QueryFee:         *utils.NewIntFromUint64(0),
			QueryTTLType:     OracleTTLTypeDelta,
			QueryTTLValue:    300,
			ResponseTTLType:  OracleTTLTypeDelta,
			ResponseTTLValue: 300,
			VMVersion:        0,
		},
		StateChannels: StateChannelConfig{ // UNUSED
			LockPeriod:     0,
			ChannelReserve: 0,
		},
	},
	Tuning: TuningConfig{
		ChainPollInterval: 100,
		ChainTimeout:      5000,
		CryptoKdfMemlimit: 1024 * 32, // 32MB
		CryptoKdfOpslimit: 3,
		CryptoKdfThreads:  1,
		OutputFormatJSON:  false,
	},
}
