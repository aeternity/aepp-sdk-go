package aeternity

import (
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type generateTTLNoncerNodeInterface interface {
	naet.GetAccounter
	naet.GetHeighter
}

// CreateOracle registers a new oracle with the given queryspec and responsespec
func (ctx *Context) CreateOracle(queryspec, responsespec string, queryFee *big.Int, queryTTLType uint64, oracleTTL uint64) (oracleID string, err error) {
	registerTx, err := transactions.NewOracleRegisterTx(ctx.Account.Address, queryspec, responsespec, queryFee, queryTTLType, oracleTTL, config.Client.Oracles.ABIVersion, ctx.ttlnoncer)
	if err != nil {
		return
	}

	ctx.SignBroadcastWait(registerTx, config.Client.WaitBlocks)
	return registerTx.ID(), nil
}

// ListenOracleQueries polls the node at a custom interval and returns queries
// and errors in their respective channels. listenInterval should be specified
// in milliseconds.
func ListenOracleQueries(n naet.GetOracleQueriesByPubkeyer, oracleID string, queryChan chan *models.OracleQuery, errChan chan error, listenInterval time.Duration) (err error) {
	// Node always returns all queries, but keeping track of until where we read
	// last iteration ensures we only report newly arriving queries. This means
	// the first time this loop runs, it will always return all the queries to
	// an oracle.
	var readUntilPosition int
	for {
		oQueries, err := n.GetOracleQueriesByPubkey(oracleID)
		if err != nil {
			errChan <- err
		} else {
			for _, q := range oQueries.OracleQueries[readUntilPosition:] {
				queryChan <- q
				readUntilPosition++
			}
		}

		time.Sleep(listenInterval * time.Millisecond)
	}
}
