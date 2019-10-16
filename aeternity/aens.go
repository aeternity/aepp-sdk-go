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
	ttlNonceGetter := GenerateGetTTLNonce(GenerateGetTTL(n), GenerateGetNextNonce(n))
	ttl, nonce, err := ttlNonceGetter(acc.Address, config.Client.TTL)
	if err != nil {
		return
	}

	cm, nameSalt, err := generateCommitmentID(name)
	preclaimTx := transactions.NewNamePreclaimTx(acc.Address, cm, config.Client.Fee, ttl, nonce)
	transactions.CalculateFee(preclaimTx)
	_, _, _, _, _, err = SignBroadcastWaitTransaction(preclaimTx, acc, n.(*naet.Node), networkID, config.Client.WaitBlocks)
	if err != nil {
		return
	}

	ttl, nonce, err = ttlNonceGetter(acc.Address, config.Client.TTL)
	if err != nil {
		return
	}
	claimTx := transactions.NewNameClaimTx(acc.Address, name, nameSalt, nameFee, config.Client.Fee, ttl, nonce)
	transactions.CalculateFee(claimTx)
	signedTxStr, hash, signature, blockHeight, blockHash, err = SignBroadcastWaitTransaction(claimTx, acc, n.(*naet.Node), networkID, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}
