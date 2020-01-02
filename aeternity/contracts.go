package aeternity

import (
	"fmt"

	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

// Contract is a higher level interface to smart contract functionalities.
type Contract struct {
	ctx ContextInterface
}

// NewContract creates a new Contract higher level interface object
func NewContract(ctx ContextInterface) *Contract {
	return &Contract{ctx: ctx}
}

// Deploy lets one deploy a contract with minimum fuss.
func (c *Contract) Deploy(source, function string, args []string, backend string) (ctID string, createTxReceipt *TxReceipt, err error) {
	bytecode, err := c.ctx.Compiler().CompileContract(source, backend)
	if err != nil {
		return
	}
	calldata, err := c.ctx.Compiler().EncodeCalldata(source, function, args, backend)
	if err != nil {
		return
	}

	_, version, err := c.ctx.NodeInfo()
	if err != nil {
		return
	}
	VMVersion, ABIVersion, err := findVMABIVersion(version, backend)
	if err != nil {
		return
	}

	createTx, err := transactions.NewContractCreateTx(c.ctx.SenderAccount(), bytecode, VMVersion, ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, calldata, c.ctx.TTLNoncer())
	if err != nil {
		return
	}

	createTxStr, _ := transactions.SerializeTx(createTx)
	fmt.Printf("%+v\n", createTx)
	fmt.Println(createTxStr)
	createTxReceipt, err = c.ctx.SignBroadcastWait(createTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	ctID, err = createTx.ContractID()
	return
}

// Call calls a smart contract's function, automatically calling the
// compiler to transform the arguments into bytecode.
func (c *Contract) Call(ctID, source, function string, args []string, backend string) (txReceipt *TxReceipt, err error) {
	callData, err := c.ctx.Compiler().EncodeCalldata(source, function, args, backend)
	if err != nil {
		return
	}

	callTx, err := transactions.NewContractCallTx(c.ctx.SenderAccount(), ctID, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, c.ctx.TTLNoncer())
	if err != nil {
		return
	}

	return c.ctx.SignBroadcastWait(callTx, config.Client.WaitBlocks)
}
