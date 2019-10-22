package aeternity

import (
	"math/big"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/naet"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
)

// RegisterName allows one to easily register a name on AENS. It does the
// preclaim, transaction sending, confirmation and claim for you.
func RegisterName(n naet.NodeInterface, acc *account.Account, name string, nameFee *big.Int) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	status, err := n.GetStatus()
	if err != nil {
		return
	}
	networkID := *status.NetworkID
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
