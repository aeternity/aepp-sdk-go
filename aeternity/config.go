package aeternity

import (
	"fmt"
	"io/ioutil"
	"strings"

	utils "github.com/aeternity/aepp-sdk-go/utils"

	yaml "gopkg.in/yaml.v2"
)

const (
	ConfigFilename = "aecli.config"
)

// EpochConfig configuration for the epoch node
type EpochConfig struct {
	URL          string `yaml:"url"`
	InternalURL  string `yaml:"internal_url"`
	WebsocketURL string `yaml:"websocket_url"`
}

type AetNameConfig struct {
	TTL       int64 `yaml:"ttl"`
	ClientTTL int64 `yaml:"client_ttl"`
}

type AetContractConfig struct {
	Gas       int64 `yaml:"gas"`
	GasPrice  int64 `yaml:"gas_price"`
	Deposit   int64 `yaml:"deposit"`
	VMVersion int64 `yaml:"vm_version"`
}

// ClientConfig client paramters configuration
type ClientConfig struct {
	TxTTL      int64             `yaml:"tx_ttl"`
	Fee        int64             `yaml:"tx_ttl"`
	DefaultKey string            `yaml:"default_key_name"`
	Names      AetNameConfig     `yaml:"names"`
	Contracts  AetContractConfig `yaml:"contracts"`
}

// TuningConfig fine tuning of parameters of the client
type TuningConfig struct {
	ChainPollInteval float32 `yaml:"chain_poll_interval"`
	TxPayload        string  `yaml:"tx_payload"`
	ResponseEncoding string  `yaml:"msg_encoding"`
}

// ConfigSchema define the configuration object
type ConfigSchema struct {
	Epoch  EpochConfig  `yaml:"epoch"`
	Client ClientConfig `yaml:"client"`
	Tuning TuningConfig `yaml:"tuning"`
}

//Defaults generate configuration defaults
func (c *ConfigSchema) Defaults() {
	c.Client = ClientConfig{
		Contracts: AetContractConfig{},
		Names:     AetNameConfig{},
	}
	// for server
	utils.DefaultIfEmptyStr(&c.Epoch.URL, "https://sdk-testnet.aepps.com")
	utils.DefaultIfEmptyStr(&c.Epoch.InternalURL, "http://127.0.0.1:3113")
	utils.DefaultIfEmptyStr(&c.Epoch.WebsocketURL, "ws://127.0.0.1:3013")
	// for client
	utils.DefaultIfEmptyStr(&c.Client.DefaultKey, "wallet.key")
	utils.DefaultIfEmptyInt64(&c.Client.TxTTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Fee, 1)
	utils.DefaultIfEmptyInt64(&c.Client.Names.TTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.Gas, 40000000)
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.GasPrice, 1)
	// for tuning
	utils.DefaultIfEmptyStr(&c.Tuning.TxPayload, "payload")
	utils.DefaultIfEmptyStr(&c.Tuning.ResponseEncoding, "json")
}

//Validate configuration
func (c *ConfigSchema) Validate() {

}

// Config sytem configuration
var Config ConfigSchema

// GenerateDefaultConfig generate a default configuration file an writes it in the outFile
func GenerateDefaultConfig(outFile, version string) {
	Config.Defaults()
	b, _ := yaml.Marshal(Config)
	data := strings.Join([]string{
		"#",
		fmt.Sprintf("# Default configuration for aepp-sdk-go v%s", version),
		"#\n",
		fmt.Sprintf("%s", b),
		"#",
		"# Config end",
		"#",
	}, "\n")
	err := ioutil.WriteFile(outFile, []byte(data), 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("config file written to", outFile)
}
