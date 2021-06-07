package swagguard_test

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

// Unfortunately, one must patch the generated go-swagger code to correctly parse the GenericTx into a specific Transaction type.
// Ideally the test JSON snippet should include every possible transaction type, but it only tests SpendTx for now.
func TestGenericTxsPolymorphicDeserialization(t *testing.T) {
	genericTxsJSON := []byte("{\"transactions\":[{\"block_hash\":\"mh_qoxFYgZoG7NxocZqBk1Dx9DJZKboT9WCKGFWV26QHKhB3oPDp\",\"block_height\":8835,\"hash\":\"th_uRnWPL3iqiLB7MzsVQ3aAgHsFALd62cmcNkw7ea1n9sfXybcr\",\"signatures\":[\"sg_RBq5vRuRchZ26HCPnZk5Xorn3ooQtfZSxf4mEPCsnpV9UD9KuZvqEKHM3vosoVvCvxvF5CyfdDkzCHnYW8bs3Ai7EtMBY\"],\"tx\":{\"amount\":10,\"fee\":20000000000000,\"nonce\":54,\"payload\":\"Hello World\",\"recipient_id\":\"ak_wJ3iKZcqvgdnQ6YVz8pY2xPjtVTNNEL61qF4AYQdksZfXZLks\",\"sender_id\":\"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi\",\"ttl\":9335,\"type\":\"SpendTx\",\"version\":1}}]}")
	genericTxs := models.GenericTxs{}
	err := genericTxs.UnmarshalBinary(genericTxsJSON)
	if err != nil {
		t.Error(err)
	}
}
