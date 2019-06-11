package aeternity

import (
	"math/big"

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
	NameTTL     uint64  `json:"name_ttl" yaml:"name_ttl" mapstructure:"name_ttl"`
	ClientTTL   uint64  `json:"client_ttl" yaml:"client_ttl" mapstructure:"client_ttl"`
	PreClaimFee big.Int `json:"preclaim_fee" yaml:"preclaim_fee" mapstructure:"preclaim_fee"`
	ClaimFee    big.Int `json:"claim_fee" yaml:"claim_fee" mapstructure:"claim_fee"`
	UpdateFee   big.Int `json:"update_fee" yaml:"update_fee" mapstructure:"update_fee"`
}

// ContractConfig configurations for contracts
type ContractConfig struct {
	Gas        big.Int `json:"gas" yaml:"gas" mapstructure:"gas"`
	GasPrice   big.Int `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	Amount     big.Int `json:"amount" yaml:"amount" mapstructure:"amount"`
	Deposit    big.Int `json:"deposit" yaml:"deposit" mapstructure:"deposit"`
	VMVersion  uint16  `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
	ABIVersion uint16  `json:"abi_version" yaml:"abi_version" mapstructure:"abi_version"`
}

// OracleConfig configurations for contracts
type OracleConfig struct {
	QueryFee         big.Int `json:"query_fee" yaml:"query_fee" mapstructure:"query_fee"`
	QueryTTLType     uint64  `json:"query_ttl_type" yaml:"query_ttl_type" mapstructure:"query_ttl_type"`
	QueryTTLValue    uint64  `json:"query_ttl_value" yaml:"query_ttl_value" mapstructure:"query_ttl_value"`
	ResponseTTLType  uint64  `json:"response_ttl_type" yaml:"response_ttl_type" mapstructure:"response_ttl_type"`
	ResponseTTLValue uint64  `json:"response_ttl_value" yaml:"response_ttl_value" mapstructure:"response_ttl_value"`
	VMVersion        uint64  `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// StateChannelConfig configurations for contracts TODO: not complete
type StateChannelConfig struct {
	LockPeriod     uint64 `json:"lock_period" yaml:"lock_period" mapstructure:"lock_period"`
	ChannelReserve uint64 `json:"channel_reserve" yaml:"channel_reserve" mapstructure:"channel_reserve"`
}

// ClientConfig client parameters configuration
type ClientConfig struct {
	BaseGas            big.Int
	GasPerByte         big.Int
	GasPrice           big.Int
	TTL                uint64             `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	Fee                big.Int            `json:"fee" yaml:"fee" mapstructure:"fee"`
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
			Gas:        *utils.NewIntFromUint64(1e9),
			GasPrice:   *utils.NewIntFromUint64(1e9),
			Amount:     *new(big.Int),
			Deposit:    *new(big.Int),
			VMVersion:  3,
			ABIVersion: 1,
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
