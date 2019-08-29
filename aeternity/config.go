package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/utils"
)

// Acceptable values for various parameters
const (
	ConfigFilename      = "config"
	NetworkIDMainnet    = "ae_mainnet"
	URLMainnet          = "https://sdk-mainnet.aepps.com"
	NetworkIDTestnet    = "ae_uat"
	URLTestnet          = "https://sdk-testnet.aepps.com"
	CompilerBackendFATE = "fate"
	CompilerBackendAEVM = "aevm"
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

// AensConfig configurations for Aens
type AensConfig struct {
	NameTTL     uint64  `json:"name_ttl" yaml:"name_ttl" mapstructure:"name_ttl"`
	ClientTTL   uint64  `json:"client_ttl" yaml:"client_ttl" mapstructure:"client_ttl"`
	PreClaimFee big.Int `json:"preclaim_fee" yaml:"preclaim_fee" mapstructure:"preclaim_fee"`
	ClaimFee    big.Int `json:"claim_fee" yaml:"claim_fee" mapstructure:"claim_fee"`
	UpdateFee   big.Int `json:"update_fee" yaml:"update_fee" mapstructure:"update_fee"`
}

// ContractConfig configurations for contracts
type ContractConfig struct {
	CompilerURL string  `json:"compiler" yaml:"compiler" mapstructure:"compiler"`
	Gas         big.Int `json:"gas" yaml:"gas" mapstructure:"gas"`
	GasPrice    big.Int `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	Amount      big.Int `json:"amount" yaml:"amount" mapstructure:"amount"`
	Deposit     big.Int `json:"deposit" yaml:"deposit" mapstructure:"deposit"`
	VMVersion   uint16  `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
	ABIVersion  uint16  `json:"abi_version" yaml:"abi_version" mapstructure:"abi_version"`
}

// OracleConfig configurations for contracts
type OracleConfig struct {
	QueryFee         big.Int `json:"query_fee" yaml:"query_fee" mapstructure:"query_fee"`
	QueryTTLType     uint64  `json:"query_ttl_type" yaml:"query_ttl_type" mapstructure:"query_ttl_type"`
	QueryTTLValue    uint64  `json:"query_ttl_value" yaml:"query_ttl_value" mapstructure:"query_ttl_value"`
	ResponseTTLType  uint64  `json:"response_ttl_type" yaml:"response_ttl_type" mapstructure:"response_ttl_type"`
	ResponseTTLValue uint64  `json:"response_ttl_value" yaml:"response_ttl_value" mapstructure:"response_ttl_value"`
	VMVersion        uint16  `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// StateChannelConfig configurations for contracts TODO: not complete
type StateChannelConfig struct {
	LockPeriod     uint64 `json:"lock_period" yaml:"lock_period" mapstructure:"lock_period"`
	ChannelReserve uint64 `json:"channel_reserve" yaml:"channel_reserve" mapstructure:"channel_reserve"`
}

// ClientConfig client parameters configuration
type ClientConfig struct {
	BaseGas       big.Int            `json:"base_gas" yaml:"base_gas" mapstructure:"base_gas"`
	GasPerByte    big.Int            `json:"gas_per_byte" yaml:"gas_per_byte" mapstructure:"gas_per_byte"`
	GasPrice      big.Int            `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	TTL           uint64             `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	Fee           big.Int            `json:"fee" yaml:"fee" mapstructure:"fee"`
	DefaultKey    string             `json:"default_key_name" yaml:"default_key_name" mapstructure:"default_key_name"`
	Names         AensConfig         `json:"names" yaml:"names" mapstructure:"names"`
	Contracts     ContractConfig     `json:"contracts" yaml:"contracts" mapstructure:"contracts"`
	Oracles       OracleConfig       `json:"oracles" yaml:"oracles" mapstructure:"oracles"`
	StateChannels StateChannelConfig `json:"state_channels" yaml:"state_channels" mapstructure:"state_channels"`
	Offline       bool               `json:"offline" yaml:"offline" mapstructure:"offline"`
}

// TuningConfig fine tuning of parameters of the client
type TuningConfig struct {
	ChainPollInteval  int64  `json:"chain_poll_interval" yaml:"chain_poll_interval" mapstructure:"chain_poll_interval"`
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

// Config system configuration
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
		GasPrice:   *utils.NewIntFromUint64(1000000000),
		TTL:        500,
		Fee:        *utils.RequireIntFromString("200000000000000"),
		Names: AensConfig{
			NameTTL:     500, // absolute block height when the name will expire
			ClientTTL:   500, // time in blocks until the name resolver should check again in case the name was updated
			PreClaimFee: *utils.RequireIntFromString("100000000000000"),
			ClaimFee:    *utils.RequireIntFromString("100000000000000"),
			UpdateFee:   *utils.RequireIntFromString("100000000000000"),
		},
		Contracts: ContractConfig{
			CompilerURL: "http://localhost:3080",
			Gas:         *utils.NewIntFromUint64(1e9),
			GasPrice:    *utils.NewIntFromUint64(1e9),
			Amount:      *new(big.Int),
			Deposit:     *new(big.Int),
			VMVersion:   4,
			ABIVersion:  1,
		},
		Oracles: OracleConfig{
			QueryFee:         *utils.NewIntFromUint64(0),
			QueryTTLType:     0,
			QueryTTLValue:    300,
			ResponseTTLType:  0,
			ResponseTTLValue: 300,
			VMVersion:        0,
		},
		StateChannels: StateChannelConfig{ // UNUSED
			LockPeriod:     0,
			ChannelReserve: 0,
		},
		Offline: false, // UNUSED
	},
	Tuning: TuningConfig{
		ChainPollInteval:  100,
		ChainTimeout:      5000,
		CryptoKdfMemlimit: 1024 * 32, // 32MB
		CryptoKdfOpslimit: 3,
		CryptoKdfThreads:  1,
		OutputFormatJSON:  false,
	},
}
