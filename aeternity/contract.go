package aeternity

import (
	rlp "github.com/randomshinichi/rlpae"
)

// ContractFunction struct represents the type information for a single function
// in a Sophia smart contract. FuncHash is the Blake2b hash of the function name
// and function types. All data is provided by the compiler in the cb_ compiled
// bytecode.
type ContractFunction struct {
	FuncHash []byte
	FuncName string
	Payable  bool
	ArgType  []byte
	OutType  []byte
}

// Contract represents the internals of the compiled cb_ bytecode that the
// compiler returns and exposes those internals as fields/methods.
type Contract struct {
	Tag             byte
	RLPVersion      byte
	SourceCodeHash  []byte
	TypeInfo        []ContractFunction
	Bytecode        []byte
	CompilerVersion string
	Payable         bool
}

// NewContractFromString takes a cb_ compiled bytecode string and returns a
// Contract struct
func NewContractFromString(cb string) (c Contract, err error) {
	rawBytes, err := Decode(cb)
	if err != nil {
		return Contract{}, err
	}

	err = rlp.DecodeBytes(rawBytes, &c)
	if err != nil {
		return Contract{}, err
	}

	return c, nil
}
