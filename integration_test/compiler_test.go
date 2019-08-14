package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/golden"
)

func TestCompiler(t *testing.T) {
	c := aeternity.NewCompiler("http://localhost:3080", false)
	t.Run("APIVersion", func(t *testing.T) {
		_, err := c.APIVersion()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("CompileContract", func(t *testing.T) {
		_, err := c.CompileContract("contract Identity =\n  type state = ()\n  entrypoint main(z : int) = z")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCallResult", func(t *testing.T) {
		// taken from contract_test.go
		_, err := c.DecodeCallResult("ok", "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACr8s/aY", "main", golden.IdentitySource)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataBytecode", func(t *testing.T) {
		_, err := c.DecodeCalldataBytecode(golden.SimpleStorageBytecode, golden.SimpleStorageCalldata)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataSource", func(t *testing.T) {
		_, err := c.DecodeCalldataSource(golden.SimpleStorageSource, golden.SimpleStorageCalldata)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeData", func(t *testing.T) {
		// taken from testnet Contract Call Tx th_toPLrggySMKVecSkEdy7QYF7VEQ4nANAdSiwNXomtwhdp6ZNw
		_, err := c.DecodeData("cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAArMtts", "int")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("EncodeCalldata", func(t *testing.T) {
		_, err := c.EncodeCalldata(golden.SimpleStorageSource, "set", []string{"123"})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GenerateACI", func(t *testing.T) {
		_, err := c.GenerateACI(golden.SimpleStorageSource)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("SophiaVersion", func(t *testing.T) {
		_, err := c.SophiaVersion()
		if err != nil {
			t.Error(err)
		}
	})

}
