package aeternity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/utils"
)

func getRLPSerialized(tx1 string, tx2 string) ([]interface{}, []interface{}) {
	tx1Bytes, _ := Decode(tx1)
	tx1RLP := DecodeRLPMessage(tx1Bytes)
	tx2Bytes, _ := Decode(tx2)
	tx2RLP := DecodeRLPMessage(tx2Bytes)
	return tx1RLP, tx2RLP
}

func TestSignedTx(t *testing.T) {
	tests := []struct {
		name    string
		tx      Transaction
		wantRLP string
		wantErr bool
	}{
		{
			name: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi signed SpendTx",
			tx: &SignedTx{
				Signatures: [][]byte{
					[]byte{10, 209, 197, 35, 33, 60, 73, 235, 31, 242, 68, 40, 83, 36, 49, 185, 210, 155, 146, 245, 148, 195, 118, 71, 232, 136, 84, 192, 104, 87, 114, 107, 26, 152, 167, 129, 192, 67, 213, 184, 220, 130, 126, 105, 22, 118, 228, 212, 198, 176, 0, 222, 210, 252, 185, 230, 18, 201, 238, 96, 105, 70, 40, 12},
				},
				Tx: &SpendTx{
					SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
					RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
					Amount:      *utils.NewIntFromUint64(10),
					Fee:         *utils.NewIntFromUint64(10),
					Payload:     []byte("Hello World"),
					TTL:         uint64(10),
					Nonce:       uint64(1),
				},
			},
			wantRLP: "tx_+KALAfhCuEAK0cUjITxJ6x/yRChTJDG50puS9ZTDdkfoiFTAaFdyaxqYp4HAQ9W43IJ+aRZ25NTGsADe0vy55hLJ7mBpRigMuFj4VgwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKCgoBi0hlbGxvIFdvcmxk+91GKg==",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s EncodeRLP", tt.name), func(t *testing.T) {
			gotRLP, err := SerializeTx(tt.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if gotRLP != tt.wantRLP {
				gotRLPRawBytes, wantRLPRawBytes := getRLPSerialized(gotRLP, tt.wantRLP)
				t.Errorf("%s = \n%v\n%v, want \n%v\n%v", tt.name, gotRLP, gotRLPRawBytes, tt.wantRLP, wantRLPRawBytes)
			}
		})
		t.Run(fmt.Sprintf("%s DecodeRLP", tt.name), func(t *testing.T) {
			tx, err := DeserializeTxStr(tt.wantRLP)

			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if !(reflect.DeepEqual(tx, tt.tx)) {
				t.Errorf("Deserialized Transaction %+v does not deep equal %+v", tx, tt.tx)
			}
		})
	}
}
