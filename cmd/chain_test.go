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

func TestTop(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	topFunc(&emptyCmd, []string{})
}

func TestStatus(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	statusFunc(&emptyCmd, []string{})
}

func TestTtl(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	ttlFunc(&emptyCmd, []string{})
}

func TestNetworkID(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	networkIDFunc(&emptyCmd, []string{})
}
