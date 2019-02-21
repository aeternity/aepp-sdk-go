package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

// Prefixing each test with Example makes go-test check the stdout
// For now, just verify that none of the commands segfault.

func setCLIConfig() {
	url := os.Getenv("AETERNITY_EXTERNAL_API")
	if len(url) == 0 {
		nodeExternalAPI = "http://localhost:3013"
	} else {
		nodeExternalAPI = url
	}
}

func TestChainTop(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	err := topFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainStatus(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	err := statusFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainTtl(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	err := ttlFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainNetworkID(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	err := networkIDFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}
