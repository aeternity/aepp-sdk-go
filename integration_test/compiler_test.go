package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/aeternity"
	"gotest.tools/golden"
)

func TestCompiler(t *testing.T) {
	simplestorageSource := "simplestorage.aes"
	simplestorageBytecode := "simplestorage_bytecode.txt"
	simplestorageCalldata := "simplestorage_init42.txt"
	identitySource := "identity.aes"

	c := aeternity.NewCompiler("http://localhost:3080", false)
	t.Run("APIVersion", func(t *testing.T) {
		_, err := c.APIVersion()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("CompileContract", func(t *testing.T) {
		compiled, err := c.CompileContract(string(golden.Get(t, simplestorageSource)), aeternity.Config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, compiled, simplestorageBytecode)
	})
	t.Run("DecodeCallResult", func(t *testing.T) {
		// taken from contract_test.go
		_, err := c.DecodeCallResult("ok", "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACr8s/aY", "main", string(golden.Get(t, identitySource)), aeternity.Config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataBytecode", func(t *testing.T) {
		_, err := c.DecodeCalldataBytecode(string(golden.Get(t, simplestorageBytecode)), string(golden.Get(t, simplestorageCalldata)), aeternity.Config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataSource", func(t *testing.T) {
		_, err := c.DecodeCalldataSource(string(golden.Get(t, simplestorageSource)), "init", string(golden.Get(t, simplestorageCalldata)), aeternity.Config.Compiler.Backend)
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
	t.Run("EncodeCalldata SimpleStorage set(123)", func(t *testing.T) {
		encodedCalldata, err := c.EncodeCalldata(string(golden.Get(t, simplestorageSource)), "set", []string{"123"}, aeternity.Config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, encodedCalldata, "simplestorage_set123.txt")
	})
	t.Run("EncodeCalldata SimpleStorage init(42)", func(t *testing.T) {
		encodedCalldata, err := c.EncodeCalldata(string(golden.Get(t, simplestorageSource)), "init", []string{"42"}, aeternity.Config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, encodedCalldata, "simplestorage_init42.txt")
	})
	t.Run("GenerateACI", func(t *testing.T) {
		_, err := c.GenerateACI(string(golden.Get(t, simplestorageSource)), aeternity.Config.Compiler.Backend)
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
