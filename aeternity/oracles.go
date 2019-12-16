package aeternity

import (
	"math/big"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type oracleQuery string
type oracleResponse string
type oracleResponder interface {
	SendOracleResponse(ctx ContextInterface, resp oracleResponse)
}

type listener func(node oracleInfoer, oracleID string, queryChan chan *models.OracleQuery, errChan chan error, listenInterval time.Duration) (err error)
type handler func(queryStr oracleQuery, respStr oracleResponder)

type oracleInfoer interface {
	naet.GetOracleByPubkeyer
	naet.GetOracleQueriesByPubkeyer
}
type Oracle struct {
	ID       string
	node     oracleInfoer
	ctx      ContextInterface
	listener listener
}

func DefaultOracleListener(node oracleInfoer, oracleID string, queryChan chan *models.OracleQuery, errChan chan error, listenInterval time.Duration) error {
	// Node always returns all queries, but keeping track of until where we read
	// last iteration ensures we only report newly arriving queries. This means
	// the first time this loop runs, it will always return all the queries to
	// an oracle.
	var readUntilPosition int
	for {
		oQueries, err := node.GetOracleQueriesByPubkey(oracleID)
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

func NewOracle(node oracleInfoer, ctx ContextInterface, ID string) *Oracle {
	return &Oracle{
		ID:       ID,
		ctx:      ctx,
		node:     node,
		listener: DefaultOracleListener,
	}
}

// CreateOracle registers a new oracle with the given queryspec and responsespec
func (o *Oracle) Register(queryspec, responsespec string, queryFee *big.Int, queryTTLType uint64, oracleTTL uint64) (oracleID string, err error) {
	registerTx, err := transactions.NewOracleRegisterTx(o.ctx.SenderAccount(), queryspec, responsespec, queryFee, queryTTLType, oracleTTL, config.Client.Oracles.ABIVersion, o.ctx.TTLNoncer())
	if err != nil {
		return
	}

	o.ctx.SignBroadcastWait(registerTx, config.Client.WaitBlocks)
	return registerTx.ID(), nil
}

func (o *Oracle) Listen() error {
	_, err := o.node.GetOracleByPubkey(o.ID)
	if err != nil { // How can we find out through the error that the oracle does not exist/404?
		return err
	}
	return nil
}
