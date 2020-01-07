package aeternity

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/binary"
	"github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type oracleListener func(node oracleInfoer, oracleID string, queryChan chan *models.OracleQuery, errChan chan error, listenInterval time.Duration)
type oracleHandler func(query string) (response string, err error)

type oracleInfoer interface {
	naet.GetOracleByPubkeyer
	naet.GetOracleQueriesByPubkeyer
}

// Oracle is a higher level interface to oracle functionalities.
type Oracle struct {
	ID                 string
	QuerySpec          string
	ResponseSpec       string
	Handler            oracleHandler
	Listener           oracleListener
	ListenPollInterval time.Duration
	node               oracleInfoer
	ctx                ContextInterface
}

// DefaultOracleListener uses a oracleInfoer to get all OracleQueries for a
// given oracleID, but keeps track of how many it read last time to that it only
// pushes new OracleQueries into the queryChan channel.
func DefaultOracleListener(node oracleInfoer, oracleID string, queryChan chan *models.OracleQuery, errChan chan error, listenInterval time.Duration) {
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
		time.Sleep(listenInterval)
	}
}

// NewOracle creates a new Oracle higher level interface object
func NewOracle(h oracleHandler, node oracleInfoer, ctx ContextInterface, QuerySpec, ResponseSpec string, pollInterval time.Duration) *Oracle {
	return &Oracle{
		ID:                 "",
		QuerySpec:          QuerySpec,
		ResponseSpec:       ResponseSpec,
		Handler:            h,
		ListenPollInterval: pollInterval,
		ctx:                ctx,
		node:               node,
		Listener:           DefaultOracleListener,
	}
}

func (o *Oracle) register(queryspec, responsespec string, queryFee *big.Int, queryTTLType uint64, oracleTTL uint64) (oracleID string, err error) {
	registerTx, err := transactions.NewOracleRegisterTx(o.ctx.SenderAccount(), queryspec, responsespec, queryFee, queryTTLType, oracleTTL, config.Client.Oracles.ABIVersion, o.ctx.TTLNoncer())
	if err != nil {
		return
	}

	o.ctx.SignBroadcastWait(registerTx, config.Client.WaitBlocks)
	return registerTx.ID(), nil
}

// RegisterIfNotExists checks if an oracle is already registered, using
// Context.SenderAccount() to figure out the oracleID. If not, it sends a
// OracleRegisterTx with default TTL values from config.
func (o *Oracle) RegisterIfNotExists() error {
	possibleOracleID := strings.Replace(o.ctx.SenderAccount(), "ak_", "ok_", 1)
	r, err := o.node.GetOracleByPubkey(possibleOracleID)

	if err != nil && strings.Contains(err.Error(), "Oracle not found") {
		fmt.Println("Couldn't find the Oracle, registering it!")
		oID, err := o.register(o.QuerySpec, o.ResponseSpec, config.Client.Oracles.QueryFee, config.Client.Oracles.QueryTTLType, config.Client.Oracles.OracleTTLValue)
		if err != nil {
			return err
		}
		o.ID = oID
		fmt.Println("Registered an Oracle", oID)
	}
	fmt.Println(r)
	return nil
}

func (o *Oracle) respondToQueries(queryChan chan *models.OracleQuery, errChan chan error) {
	for {
		q := <-queryChan
		qBin, err := binary.Decode(*q.Query)
		if err != nil {
			fmt.Println("Error decoding OracleQuery", *q.ID)
		}

		qStr := string(qBin)
		fmt.Println("Received query", *q.ID, qStr)
		resp, err := o.Handler(qStr)
		if err != nil {
			fmt.Println("Error responding to OracleQuery - reason:", err)
		}
		respTx, err := transactions.NewOracleRespondTx(o.ctx.SenderAccount(), o.ID, *q.ID, resp, config.Client.Oracles.ResponseTTLType, config.Client.Oracles.ResponseTTLValue, o.ctx.TTLNoncer())
		fmt.Println("respTx", respTx)
		receipt, err := o.ctx.SignBroadcastWait(respTx, config.Client.WaitBlocks)
		if err != nil {
			fmt.Println("Error sending a response to OracleQuery, reason:", err)
		}
		fmt.Println(receipt)
	}
}

// Listen starts polling for OracleQueries in a goroutine, passes new queries to
// Handler, and sends out a OracleResponseTx that contains Handler's return value.
func (o *Oracle) Listen() error {
	err := o.RegisterIfNotExists()
	if err != nil {
		return err
	}

	queryChan := make(chan *models.OracleQuery)
	errChan := make(chan error)
	go o.Listener(o.node, o.ID, queryChan, errChan, o.ListenPollInterval)

	o.respondToQueries(queryChan, errChan)
	return nil
}
