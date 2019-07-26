package aeternity

import (
	rlp "github.com/randomshinichi/rlpae"
)

type ContractFunction struct {
	FuncHash []byte
	FuncName string
	ArgType  []byte
	OutType  []byte
}

type contract struct {
	Tag             byte
	RLPVersion      byte
	SourceCodeHash  []byte
	TypeInfo        []ContractFunction
	Bytecode        []byte
	CompilerVersion string
}

type Contract struct {
	ct contract
}

func NewContractFromString(cb string) (Contract, error) {
	rawBytes, err := Decode(cb)
	if err != nil {
		return Contract{}, err
	}
	c := contract{}
	err = rlp.DecodeBytes(rawBytes, &c)
	if err != nil {
		return Contract{}, err
	}

	contract := Contract{ct: c}
	return contract, nil
}
