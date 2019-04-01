package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
)

// Prefixing each test with Example makes go-test check the stdout
// For now, just verify that none of the commands segfault.

func setConfigTestParams() {
	aeternity.Config.Node.URL = "http://localhost:3013"
	aeternity.Config.Node.NetworkID = "ae_docker"
}

func TestChainTop(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := topFunc(&emptyCmd, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainBroadcast(t *testing.T) {
	setConfigTestParams()
	emptyCmd := cobra.Command{}
	err := broadcastFunc(&emptyCmd, []string{"tx_+KgLAfhCuEAPX1l3BdFOcLeduH3PPwPV25mETXZE8IBDe6PGuasSEKJeB/cDDm+kW05Cdp38+mpvVSTTPMx7trL/7qxfUr8IuGD4XhYBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wGTcXVlcnkgU3BlY2lmaWNhdGlvbpZyZXNwb25zZSBTcGVjaWZpY2F0aW9uAABkhrXmIPSAAIIB9AHdGxXf"})
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
