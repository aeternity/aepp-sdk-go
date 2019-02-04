package aeternity

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

// Config sytem configuration
var Config ProfileConfig
