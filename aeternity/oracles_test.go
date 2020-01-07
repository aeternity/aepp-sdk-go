package aeternity

import (
	"fmt"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

type mockOracleQueryNode struct {
	oracleQueries *models.OracleQueries
	Count         int
}

func newMockOracleQueryNode() *mockOracleQueryNode {
	return &mockOracleQueryNode{
		oracleQueries: &models.OracleQueries{
			OracleQueries: []*models.OracleQuery{},
		},
	}
}
func (m *mockOracleQueryNode) GetOracleByPubkey(pubkey string) (oracleQueries *models.RegisteredOracle, err error) {
	return &models.RegisteredOracle{}, err
}

func (m *mockOracleQueryNode) GetOracleQueriesByPubkey(pubkey string) (oracleQueries *models.OracleQueries, err error) {
	return m.oracleQueries, err
}

func (m *mockOracleQueryNode) AddOracleQuery(n int) (err error) {
	var newOracleQuery = func(i int) (q *models.OracleQuery, err error) {
		q = new(models.OracleQuery)
		qJSON := fmt.Sprintf(`{"fee":%d,"id":"oq_FAKEQUERY","oracle_id":"ok_FAKEID","query":"ov_FAKEQUERY=","response":"or_FAKERESPONSE","response_ttl":{"type":"delta","value":100},"sender_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","sender_nonce":%d,"ttl":137}`, i, i)
		err = q.UnmarshalBinary([]byte(qJSON))
		return
	}
	// Create some fake OracleQuerys to add them to
	// mockOracleQueryNode.OracleQueries. Keep track of how many we have
	// created.
	for index := 0; index < n; index++ {
		q, err := newOracleQuery(m.Count)
		if err != nil {
			break
		}
		m.Count++
		m.oracleQueries.OracleQueries = append(m.oracleQueries.OracleQueries, q)
	}
	return
}

func TestDefaultOracleListener(t *testing.T) {
	n := newMockOracleQueryNode()

	oQueries := make(chan *models.OracleQuery, 30)
	errChan := make(chan error)
	go DefaultOracleListener(n, "ok_FAKEID", oQueries, errChan, 20)
	n.AddOracleQuery(3)
	n.AddOracleQuery(5)
	time.Sleep(20 * time.Millisecond)
	n.AddOracleQuery(15)
	time.Sleep(20 * time.Millisecond)
	t.Logf("oQueries channel contains %d queries; we generated %d\n", len(oQueries), n.Count)
	if n.Count != len(oQueries) {
		t.Fatalf("Oracle Queries channel should contain %d queries but actually had %d", n.Count, len(oQueries))
	}
}

func TestDefaultOracleListenerManyPendingQueries(t *testing.T) {
	n := newMockOracleQueryNode()
	readQueries := []*models.OracleQuery{}
	oQueries := make(chan *models.OracleQuery, 5)
	errChan := make(chan error)

	n.AddOracleQuery(100)

	go DefaultOracleListener(n, "ok_FAKEID", oQueries, errChan, 20)

	var q *models.OracleQuery
	for i := 0; i < 100; i++ {
		q = <-oQueries
		readQueries = append(readQueries, q)
	}

	if len(oQueries) != 0 {
		t.Fatalf("oQueries channel should be empty but still contains %d queries", len(oQueries))
	}
	if len(readQueries) != 100 {
		t.Fatalf("We should have 100 queries in readQueries but instead we have %d queries", len(readQueries))
	}
}
