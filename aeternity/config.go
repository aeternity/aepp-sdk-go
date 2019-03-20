package aeternity

import (
	"github.com/aeternity/aepp-sdk-go/utils"
)

const (
	// ConfigFilename default configuration file name
	ConfigFilename = "config"
)

// NodeConfig configuration for the node node
type NodeConfig struct {
	URL         string `yaml:"url" json:"url" mapstructure:"url"`
	URLInternal string `yaml:"url_internal" json:"url_internal" mapstructure:"url_internal"`
	URLChannels string `yaml:"url_channels" json:"url_channels" mapstructure:"url_channels"`
	NetworkID   string `yaml:"network_id" json:"network_id" mapstructure:"network_id"`
}

// AensConfig configurations for Aens
type AensConfig struct {
	NameTTL     uint64 `json:"name_ttl" yaml:"name_ttl" mapstructure:"name_ttl"`
	ClientTTL   uint64 `json:"client_ttl" yaml:"client_ttl" mapstructure:"client_ttl"`
	PreClaimFee uint64 `json:"preclaim_fee" yaml:"preclaim_fee" mapstructure:"preclaim_fee"`
	ClaimFee    uint64 `json:"claim_fee" yaml:"claim_fee" mapstructure:"claim_fee"`
	UpdateFee   uint64 `json:"update_fee" yaml:"update_fee" mapstructure:"update_fee"`
}

// ContractConfig configurations for contracts
type ContractConfig struct {
	Gas       uint64 `json:"gas" yaml:"gas" mapstructure:"gas"`
	GasPrice  uint64 `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	Deposit   uint64 `json:"deposit" yaml:"deposit" mapstructure:"deposit"`
	VMVersion uint64 `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// OracleConfig configurations for contracts
type OracleConfig struct {
	QueryFee  utils.BigInt `json:"query_fee" yaml:"query_fee" mapstructure:"query_fee"`
	VMVersion uint64       `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// StateChannelConfig configurations for contracts TODO: not complete
type StateChannelConfig struct {
	LockPeriod     uint64 `json:"lock_period" yaml:"lock_period" mapstructure:"lock_period"`
	ChannelReserve uint64 `json:"channel_reserve" yaml:"channel_reserve" mapstructure:"channel_reserve"`
}

// ClientConfig client paramters configuration
type ClientConfig struct {
	TTL                uint64             `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	Fee                utils.BigInt       `json:"fee" yaml:"fee" mapstructure:"fee"`
	DefaultKey         string             `json:"default_key_name" yaml:"default_key_name" mapstructure:"default_key_name"`
	Names              AensConfig         `json:"names" yaml:"names" mapstructure:"names"`
	Contracts          ContractConfig     `json:"contracts" yaml:"contracts" mapstructure:"contracts"`
	Oracles            OracleConfig       `json:"oracles" yaml:"oracles" mapstructure:"oracles"`
	StateChannels      StateChannelConfig `json:"state_channels" yaml:"state_channels" mapstructure:"state_channels"`
	NativeTransactions bool               `yaml:"native_transactions" json:"native_transactions" mapstructure:"native_transactions"`
	Offline            bool               `yaml:"offline" json:"offline" mapstructure:"offline"`
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
	Name   string       `json:"name" yaml:"name" mapstructure:"name"`
	Node   NodeConfig   `json:"node" yaml:"node" mapstructure:"node"`
	Client ClientConfig `json:"client" yaml:"client" mapstructure:"client"`
	Tuning TuningConfig `json:"tuning" yaml:"tuning" mapstructure:"tuning"`
}

// Config sytem configuration
var Config = ProfileConfig{
	Name: "Default Config",
	Node: NodeConfig{
		URL:         "https://sdk-mainnet.aepps.com",
		URLInternal: "https://sdk-mainnet.aepps.com",
		URLChannels: "https://sdk-mainnet.aepps.com",
		NetworkID:   "ae_mainnet",
	},
	Client: ClientConfig{
		TTL: 500,
		Fee: *utils.RequireBigIntFromString("200000000000000"),
		Names: AensConfig{
			NameTTL:   500,
			ClientTTL: 500,
		},
		Contracts: ContractConfig{ // UNUSED
			Gas:       1e9,
			GasPrice:  1,
			Deposit:   0,
			VMVersion: 0,
		},
		Oracles: OracleConfig{
			QueryFee:  *utils.NewBigIntFromUint64(0),
			VMVersion: 0,
		},
		StateChannels: StateChannelConfig{ // UNUSED
			LockPeriod:     0,
			ChannelReserve: 0,
		},
		NativeTransactions: false, //UNUSED
		Offline:            false, // UNUSED
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
