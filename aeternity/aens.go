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
func RegisterName(n nodeStatusHeightAccounterBroadcaster, b *Broadcaster, name string, nameFee *big.Int) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	ttlnoncer := transactions.NewTTLNoncer(n)

	preclaimTx, nameSalt, err := transactions.NewNamePreclaimTx(b.Account.Address, name, ttlnoncer)
	if err != nil {
		return
	}
	_, _, _, _, _, err = b.SignBroadcastWait(preclaimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}

	claimTx, err := transactions.NewNameClaimTx(b.Account.Address, name, nameSalt, nameFee, ttlnoncer)
	if err != nil {
		return
	}

	signedTxStr, hash, signature, blockHeight, blockHash, err = b.SignBroadcastWait(claimTx, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}
