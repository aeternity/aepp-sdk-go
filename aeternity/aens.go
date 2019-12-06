package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v7/account"
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
func RegisterName(n nodeStatusHeightAccounterBroadcaster, acc *account.Account, name string, nameFee *big.Int) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	networkID, err := getNetworkID(n)
	if err != nil {
		return
	}
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(n)

	preclaimTx, nameSalt, err := transactions.NewNamePreclaimTx(acc.Address, name, ttlnoncer)
	if err != nil {
		return
	}
	_, _, _, _, _, err = SignBroadcastWaitTransaction(preclaimTx, acc, n.(*naet.Node), networkID, config.Client.WaitBlocks)
	if err != nil {
		return
	}

	claimTx, err := transactions.NewNameClaimTx(acc.Address, name, nameSalt, nameFee, ttlnoncer)
	if err != nil {
		return
	}

	signedTxStr, hash, signature, blockHeight, blockHash, err = SignBroadcastWaitTransaction(claimTx, acc, n.(*naet.Node), networkID, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}
