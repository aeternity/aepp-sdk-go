package aeternity

import (
	"time"

	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
)

type callResultListener interface {
	naet.GetHeighter
	naet.GetTransactionInfoByHasher
}

// Contract is a higher level interface to smart contract functionalities.
type Contract struct {
	ctx ContextInterface
}

// NewContract creates a new Contract higher level interface object
func NewContract(ctx ContextInterface) *Contract {
	return &Contract{ctx: ctx}
}

// DefaultCallResultListener polls the node for the result of a particular
// transaction until /transaction/txhash/info is filled out with the pertinent
// data (only for ContractCallTxs). Then it will push the CallInfo to a channel.
// This function is intended to be run as a goroutine.
func DefaultCallResultListener(node callResultListener, txHash string, callResultChan chan *models.ContractCallObject, errChan chan error, listenInterval time.Duration) {
	for {
		txInfo, err := node.GetTransactionInfoByHash(txHash)
		if err != nil {
			errChan <- err
		}
		if *txInfo.CallInfo.GasUsed != 0 {
			callResultChan <- txInfo.CallInfo
			errChan <- nil
			break
		}
		time.Sleep(listenInterval)
	}
}

// Deploy lets one deploy a contract with minimum fuss.
func (c *Contract) Deploy(source, function string, args []string) (ctID string, createTxReceipt *TxReceipt, err error) {
	bytecode, err := c.ctx.Compiler().CompileContract(source)
	if err != nil {
		return
	}
	calldata, err := c.ctx.Compiler().EncodeCalldata(source, function, args)
	if err != nil {
		return
	}

	_, version, err := c.ctx.NodeInfo()
	if err != nil {
		return
	}
	VMVersion, ABIVersion, err := findVMABIVersion(version)
	if err != nil {
		return
	}

	createTx, err := transactions.NewContractCreateTx(c.ctx.SenderAccount(), bytecode, VMVersion, ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, calldata, c.ctx.TTLNoncer())
	if err != nil {
		return
	}

	createTxReceipt, err = c.ctx.SignBroadcastWait(createTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	ctID, err = createTx.ContractID()
	return
}

// Call calls a smart contract's function, automatically calling the
// compiler to transform the arguments into bytecode.
func (c *Contract) Call(ctID, source, function string, args []string) (txReceipt *TxReceipt, err error) {
	callData, err := c.ctx.Compiler().EncodeCalldata(source, function, args)
	if err != nil {
		return
	}

	callTx, err := transactions.NewContractCallTx(c.ctx.SenderAccount(), ctID, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, c.ctx.TTLNoncer())
	if err != nil {
		return
	}

	return c.ctx.SignBroadcastWait(callTx, config.Client.WaitBlocks)
}
