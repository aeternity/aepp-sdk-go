package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
)

// AENS is a higher level interface to AENS functionalities.
type AENS struct {
	ctx ContextInterface
}

// NewAENS creates a new AENS higher level interface object
func NewAENS(ctx ContextInterface) *AENS {
	return &AENS{ctx: ctx}
}

// RegisterName allows one to easily register a name on AENS. It does the
// preclaim, transaction sending, confirmation and claim for you.
func (aens *AENS) RegisterName(name string, nameFee *big.Int) (txReceipt *TxReceipt, err error) {
	preclaimTx, nameSalt, err := transactions.NewNamePreclaimTx(aens.ctx.SenderAccount(), name, aens.ctx.TTLNoncer())
	if err != nil {
		return
	}
	txReceipt, err = aens.ctx.SignBroadcastWait(preclaimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}

	claimTx, err := transactions.NewNameClaimTx(aens.ctx.SenderAccount(), name, nameSalt, nameFee, aens.ctx.TTLNoncer())
	if err != nil {
		return
	}

	txReceipt, err = aens.ctx.SignBroadcastWait(claimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}
