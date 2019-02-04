package aeternity

import (
	"path/filepath"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

const (
	// ConfigFilename default configuration file name
	ConfigFilename = "config"
)

// EpochConfig configuration for the epoch node
type EpochConfig struct {
	URL         string `yaml:"url" json:"url" mapstructure:"url"`
	URLInternal string `yaml:"url_internal" json:"url_internal" mapstructure:"url_internal"`
	URLChannels string `yaml:"url_channels" json:"url_channels" mapstructure:"url_channels"`
	NetworkID   string `yaml:"network_id" json:"network_id" mapstructure:"network_id"`
}

// AensConfig configurations for Aens
type AensConfig struct {
	NameTTL     uint64 `json:"name_ttl" yaml:"name_ttl" mapstructure:"name_ttl"`
	ClientTTL   uint64 `json:"client_ttl" yaml:"client_ttl" mapstructure:"client_ttl"`
	PreClaimFee int64  `json:"preclaim_fee" yaml:"preclaim_fee" mapstructure:"preclaim_fee"`
	ClaimFee    int64  `json:"claim_fee" yaml:"claim_fee" mapstructure:"claim_fee"`
	UpdateFee   int64  `json:"update_fee" yaml:"update_fee" mapstructure:"update_fee"`
}

// ContractConfig configurations for contracts
type ContractConfig struct {
	Gas       int64 `json:"gas" yaml:"gas" mapstructure:"gas"`
	GasPrice  int64 `json:"gas_price" yaml:"gas_price" mapstructure:"gas_price"`
	Deposit   int64 `json:"deposit" yaml:"deposit" mapstructure:"deposit"`
	VMVersion int64 `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// OracleConfig configurations for contracts
type OracleConfig struct {
	QueryFee  int64 `json:"query_fee" yaml:"query_fee" mapstructure:"query_fee"`
	VMVersion int64 `json:"vm_version" yaml:"vm_version" mapstructure:"vm_version"`
}

// StateChannelConfig configurations for contracts TODO: not complete
type StateChannelConfig struct {
	LockPeriod     int64 `json:"lock_period" yaml:"lock_period" mapstructure:"lock_period"`
	ChannelReserve int64 `json:"channel_reserve" yaml:"channel_reserve" mapstructure:"channel_reserve"`
}

// ClientConfig client paramters configuration
type ClientConfig struct {
	TTL                uint64             `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	Fee                int64              `json:"fee" yaml:"fee" mapstructure:"fee"`
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
	Epoch  EpochConfig  `json:"epoch" yaml:"epoch" mapstructure:"epoch"`
	Client ClientConfig `json:"client" yaml:"client" mapstructure:"client"`
	Tuning TuningConfig `json:"tuning" yaml:"tuning" mapstructure:"tuning"`
}

// ConfigSchema define the configuration object
type ConfigSchema struct {
	P          *ProfileConfig `json:"-" yaml:"-" mapstructure:"-"` // holds the active profile
	ConfigPath string         `json:"-" yaml:"-" mapstructure:"-"` // the path of the configuration file
	KeysFolder string         `json:"-" yaml:"-" mapstructure:"-"`
}

var DefaultConfig = ProfileConfig{
	Name: "Default Config",
	Epoch: EpochConfig{
		URL:         "https://sdk-edgenet.aepps.com",
		URLInternal: "https://sdk-edgenet.aepps.com",
		URLChannels: "https://sdk-edgenet.aepps.com",
		NetworkID:   "ae_mainnet",
	},
	Client: ClientConfig{
		TTL:        500,
		Fee:        20000,
		DefaultKey: "wallet.key", // UNUSED
		Names: AensConfig{
			NameTTL:     500,
			ClientTTL:   500,
			PreClaimFee: 1,
			ClaimFee:    1,
			UpdateFee:   1,
		},
		Contracts: ContractConfig{ // UNUSED
			Gas:       1000,
			GasPrice:  1,
			Deposit:   0,
			VMVersion: 0,
		},
		Oracles: OracleConfig{ // UNUSED
			QueryFee:  0,
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

//Defaults generate configuration defaults
func (c *ProfileConfig) Defaults() *ProfileConfig {
	c.Client = ClientConfig{
		Contracts: ContractConfig{},
		Names:     AensConfig{},
	}
	// for server
	utils.DefaultIfEmptyStr(&c.Epoch.URL, "https://sdk-edgenet.aepps.com")
	utils.DefaultIfEmptyStr(&c.Epoch.URLInternal, "https://sdk-edgenet.aepps.com") // UNUSED
	utils.DefaultIfEmptyStr(&c.Epoch.URLChannels, "https://sdk-edgenet.aepps.com") // UNUSED
	utils.DefaultIfEmptyStr(&c.Epoch.NetworkID, "ae_mainnet")
	// for client
	utils.DefaultIfEmptyStr(&c.Client.DefaultKey, "wallet.key") // UNUSED
	utils.DefaultIfEmptyUint64(&c.Client.TTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Fee, 20000)
	// for aens
	utils.DefaultIfEmptyUint64(&c.Client.Names.NameTTL, 500)
	utils.DefaultIfEmptyUint64(&c.Client.Names.ClientTTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Names.PreClaimFee, 1)
	utils.DefaultIfEmptyInt64(&c.Client.Names.ClaimFee, 1)
	utils.DefaultIfEmptyInt64(&c.Client.Names.UpdateFee, 1)
	// UNUSED for contracts
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.Gas, 1000)
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.GasPrice, 1)
	// for tuning
	utils.DefaultIfEmptyInt64(&c.Tuning.ChainPollInteval, 1000)
	utils.DefaultIfEmptyInt64(&c.Tuning.ChainTimeout, 5000)
	utils.DefaultIfEmptyUint32(&c.Tuning.CryptoKdfMemlimit, 1024*32) // 32mb
	utils.DefaultIfEmptyUint32(&c.Tuning.CryptoKdfOpslimit, 3)
	utils.DefaultIfEmptyUint8(&c.Tuning.CryptoKdfThreads, 1)
	return c
}

//Validate configuration
func (c *ProfileConfig) Validate() {

}

//Validate configuration
func (c *ConfigSchema) Validate() {

}

//Defaults configuration
func (c *ConfigSchema) Defaults() *ConfigSchema {
	p := ProfileConfig{}
	c.P = p.Defaults()
	return c
}

// Config sytem configuration
var Config ConfigSchema

// GenerateDefaultConfig generate a default configuration
func GenerateDefaultConfig(outFile, version string) {
	Config = ConfigSchema{
		KeysFolder: filepath.Join(filepath.Dir(outFile), "accounts"),
		ConfigPath: outFile,
	}
}
