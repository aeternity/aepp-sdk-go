package aeternity

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/binary"
	"github.com/aeternity/aepp-sdk-go/v5/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/v5/utils"
)

type mockClient struct {
	i uint64
}

func (m *mockClient) GetHeight() (uint64, error) {
	m.i++
	return m.i, nil
}

// GetTransactionByHash pretends that the transaction was not mined until block 9, and this is only visible when the mockClient is at height 10.
func (m *mockClient) GetTransactionByHash(hash string) (tx *models.GenericSignedTx, err error) {
	unminedHeight, _ := utils.NewIntFromString("-1")
	minedHeight, _ := utils.NewIntFromString("9")

	bh := "bh_someblockhash"
	tx = &models.GenericSignedTx{
		BlockHash:   &bh,
		BlockHeight: utils.BigInt{},
		Hash:        &hash,
		Signatures:  nil,
	}

	if m.i == 10 {
		tx.BlockHeight.Set(minedHeight)
	} else {
		tx.BlockHeight.Set(unminedHeight)
	}
	return tx, nil
}
func TestWaitForTransactionForXBlocks(t *testing.T) {
	m := new(mockClient)
	blockHeight, blockHash, err := WaitForTransactionForXBlocks(m, "th_transactionhash", 10)
	if err != nil {
		t.Fatal(err)
	}
	if blockHeight != 9 {
		t.Fatalf("Expected mock blockHeight 9, got %v", blockHeight)
	}
	if blockHash != "bh_someblockhash" {
		t.Fatalf("Expected mock blockHash bh_someblockhash, got %s", blockHash)
	}
}
func Test_Namehash(t *testing.T) {
	// ('welghmolql.aet') == 'nm_2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ'
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{"welghmolql.aet"}, "nm_2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ"},
		{"ok", args{"welghmolql"}, "nm_2nLRBu1FyukEvJuMANjFzx8mubMFeyG2mJ2QpQoYKymYe1d2sr"},
		{"ok", args{"fdsa.test"}, "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb"},
		{"ok", args{""}, "nm_2q1DrgEuxRNCWRp5nTs6FyA7moSEzrPVUSTEpkpFsM4hRL4Dkb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binary.Encode(binary.PrefixName, Namehash(tt.args.name))
			if got != tt.want {
				t.Errorf("Namehash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeCommitmentID(t *testing.T) {
	type args struct {
		name string
		salt []byte
	}
	tests := []struct {
		name    string
		args    args
		wantCh  string
		wantErr bool
	}{
		{
			name: "fdsa.test, 0",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			wantCh:  "cm_2jJov6dn121oKkHo6TuWaAAL4ZEMonnCjpo8jatkCixrLG8Uc4",
			wantErr: false,
		},
		{
			name: "fdsa.test, 255",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255},
			},
			wantCh:  "cm_sa8UUjorPzCTLfYp6YftR4jwF4kPaZVsoP5bKVAqRw9zm43EE",
			wantErr: false,
		},
		{
			// erlang Eshell: rp(<<9795159241593061970:256>>).
			name: "fdsa.test, 9795159241593061970 (do not use Golang to convert salt integers)",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 135, 239, 101, 110, 233, 138, 2, 82},
			},
			wantCh:  "cm_QhtcYow8krP3xQSTsAhFihfBstTjQMiApaPCgZuciDHZmMNtZ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fmt.Println(saltBytes)
			gotCh, err := computeCommitmentID(tt.args.name, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeCommitmentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCh != tt.wantCh {
				t.Errorf("computeCommitmentID() = %v, want %v", gotCh, tt.wantCh)
			}
		})
	}
}
