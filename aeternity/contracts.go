package aeternity

import (
	"errors"
	"fmt"

	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

func findVMABIVersion(nodeVersion, compilerBackend string) (VMVersion, ABIVersion uint16, err error) {
	if nodeVersion[0] == '5' && compilerBackend == "fate" {
		return 5, 3, nil
	} else if nodeVersion[0] == '5' && compilerBackend == "aevm" {
		return 6, 1, nil
	} else if nodeVersion[0] == '4' {
		return 4, 1, nil
	} else {
		return 0, 0, errors.New("Other node versions unsupported")
	}
}

// CreateContract lets one deploy a contract with minimum fuss.
func (ctx *Context) CreateContract(source, function string, args []string, backend string) (ctID, signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	bytecode, err := ctx.compiler.CompileContract(source, backend)
	if err != nil {
		return
	}
	calldata, err := ctx.compiler.EncodeCalldata(source, function, args, backend)
	if err != nil {
		return
	}

	status, err := ctx.node.GetStatus()
	if err != nil {
		return
	}
	VMVersion, ABIVersion, err := findVMABIVersion(*status.NodeVersion, backend)
	if err != nil {
		return
	}

	createTx, err := transactions.NewContractCreateTx(ctx.Account.Address, bytecode, VMVersion, ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, calldata, ctx.ttlnoncer)
	if err != nil {
		return
	}

	createTxStr, _ := transactions.SerializeTx(createTx)
	fmt.Printf("%+v\n", createTx)
	fmt.Println(createTxStr)
	signedTxStr, hash, signature, blockHeight, blockHash, err = ctx.SignBroadcastWait(createTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	ctID, err = createTx.ContractID()
	return
}

// CallContract calls a smart contract's function, automatically calling the
// compiler to transform the arguments into bytecode.
func (ctx *Context) CallContract(ctID, source, function string, args []string, backend string) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	callData, err := ctx.compiler.EncodeCalldata(source, function, args, backend)
	if err != nil {
		return
	}

	callTx, err := transactions.NewContractCallTx(ctx.Account.Address, ctID, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, ctx.ttlnoncer)
	if err != nil {
		return
	}

	return ctx.SignBroadcastWait(callTx, config.Client.WaitBlocks)
}
