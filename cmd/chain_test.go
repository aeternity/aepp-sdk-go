package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
)

// Prefixing each test with Example makes go-test check the stdout
// For now, just verify that none of the commands segfault.

func setConfigTestParams() {
	aeternity.Config.Epoch.URL = "http://localhost:3013"
	aeternity.Config.Epoch.NetworkID = "ae_docker"
}

func TestChainTop(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := topFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainStatus(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := statusFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainTtl(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := ttlFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainNetworkID(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := networkIDFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}
