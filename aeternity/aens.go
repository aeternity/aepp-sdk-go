package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type nodeStatusHeightAccounterBroadcaster interface {
	naet.GetStatuser
	naet.GetAccounter
	broadcastWaitTransactionNodeCapabilities
}

// RegisterName allows one to easily register a name on AENS. It does the
// preclaim, transaction sending, confirmation and claim for you.
func (ctx *Context) RegisterName(name string, nameFee *big.Int) (claimTxReceipt *TxReceipt, err error) {
	preclaimTx, nameSalt, err := transactions.NewNamePreclaimTx(ctx.Account.Address, name, ctx.TTLNoncer)
	if err != nil {
		return
	}
	_, err = ctx.SignBroadcastWait(preclaimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}

	claimTx, err := transactions.NewNameClaimTx(ctx.Account.Address, name, nameSalt, nameFee, ctx.TTLNoncer)
	if err != nil {
		return
	}

	claimTxReceipt, err = ctx.SignBroadcastWait(claimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}
