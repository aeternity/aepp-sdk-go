package cmd

import (
	"flag"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/spf13/cobra"
)

// Prefixing each test with Example makes go-test check the stdout
// For now, just verify that none of the commands segfault.

var online bool

func init() {
	flag.BoolVar(&online, "online", false, "Run tests that need a running node on localhost:3013, Network ID ae_docker")
	flag.Parse()
}

func setPrivateNetParams() {
	aeternity.Config.Node.URL = "http://localhost:3013"
	aeternity.Config.Node.NetworkID = "ae_docker"
}

func TestChainTop(t *testing.T) {
	setPrivateNetParams()
	err := topFunc(&cobra.Command{}, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainBroadcast(t *testing.T) {
	setPrivateNetParams()
	err := broadcastFunc(&cobra.Command{}, []string{"tx_+KgLAfhCuEAPX1l3BdFOcLeduH3PPwPV25mETXZE8IBDe6PGuasSEKJeB/cDDm+kW05Cdp38+mpvVSTTPMx7trL/7qxfUr8IuGD4XhYBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wGTcXVlcnkgU3BlY2lmaWNhdGlvbpZyZXNwb25zZSBTcGVjaWZpY2F0aW9uAABkhrXmIPSAAIIB9AHdGxXf"})
	if err != nil {
		t.Error(err)
	}
}

func TestChainStatus(t *testing.T) {
	setPrivateNetParams()
	err := statusFunc(&cobra.Command{}, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainTtl(t *testing.T) {
	setPrivateNetParams()
	err := ttlFunc(&cobra.Command{}, []string{})
	if err != nil {
		t.Error(err)
	}
}

func TestChainNetworkID(t *testing.T) {
	setPrivateNetParams()
	err := networkIDFunc(&cobra.Command{}, []string{})
	if err != nil {
		t.Error(err)
	}
}
