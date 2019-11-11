package transactions

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/utils"
	rlp "github.com/randomshinichi/rlpae"
)

func TestSpendTx_EncodeRLP(t *testing.T) {
	tests := []struct {
		name    string
		tx      *SpendTx
		wantTx  string
		wantErr bool
	}{
		{
			name: "Spend 10, Fee 10, Hello World",
			tx: &SpendTx{
				SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Amount:      utils.NewIntFromUint64(10),
				Fee:         utils.NewIntFromUint64(10),
				Payload:     []byte("Hello World"),
				TTL:         uint64(10),
				Nonce:       uint64(1),
			},
			wantTx:  "tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCgoKAYtIZWxsbyBXb3JsZPSZjdM=",
			wantErr: false,
		},
		{
			name: "Spend 0, Fee 10, Hello World (check correct RLP serialization of 0)",
			tx: &SpendTx{
				SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Amount:      utils.NewIntFromUint64(0),
				Fee:         utils.NewIntFromUint64(10),
				Payload:     []byte("Hello World"),
				TTL:         uint64(10),
				Nonce:       uint64(1),
			},
			wantTx:  "tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAAoKAYtIZWxsbyBXb3JsZICI5/w=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txJSON, err := tt.tx.JSON()
			fmt.Println(txJSON)

			fmt.Println(rlp.EncodeToBytes(tt.tx))

			gotTx, err := SerializeTx(tt.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpendTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTx != tt.wantTx {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("SpendTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestSpendTx_DecodeRLP(t *testing.T) {
	type args struct {
		rlpBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		wantTx  SpendTx
		wantErr bool
	}{
		{
			name: "Spend 10, Fee 10, Hello World",
			args: args{
				// tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCgoKAYtIZWxsbyBXb3JsZPSZjdM=
				// [12] [1] [1 206 167 173 228 112 201 249 157 157 78 64 8 128 168 111 29 73 187 68 75 98 241 26 158 187 100 187 207 235 115 254 243] [1 31 19 163 176 139 240 1 64 6 98 166 139 105 216 117 247 128 60 236 76 8 100 127 110 213 216 76 120 151 189 80 163] [10] [10] [10] [1] [72 101 108 108 111 32 87 111 114 108 100]]
				rlpBytes: []byte{248, 86, 12, 1, 161, 1, 206, 167, 173, 228, 112, 201, 249, 157, 157, 78, 64, 8, 128, 168, 111, 29, 73, 187, 68, 75, 98, 241, 26, 158, 187, 100, 187, 207, 235, 115, 254, 243, 161, 1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163, 10, 10, 10, 1, 139, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
			},
			wantTx: SpendTx{
				SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Amount:      utils.NewIntFromUint64(10),
				Fee:         utils.NewIntFromUint64(10),
				Payload:     []byte("Hello World"),
				TTL:         uint64(10),
				Nonce:       uint64(1),
			},
			wantErr: false,
		},
		{
			name: "Spend 0, Fee 10, Hello World (check correct RLP deserialization of 0)",
			args: args{
				// tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAAoKAYtIZWxsbyBXb3JsZICI5/w=
				// [[12] [1] [1 206 167 173 228 112 201 249 157 157 78 64 8 128 168 111 29 73 187 68 75 98 241 26 158 187 100 187 207 235 115 254 243] [1 31 19 163 176 139 240 1 64 6 98 166 139 105 216 117 247 128 60 236 76 8 100 127 110 213 216 76 120 151 189 80 163] [0] [10] [10] [1] [72 101 108 108 111 32 87 111 114 108 100]]
				rlpBytes: []byte{248, 86, 12, 1, 161, 1, 206, 167, 173, 228, 112, 201, 249, 157, 157, 78, 64, 8, 128, 168, 111, 29, 73, 187, 68, 75, 98, 241, 26, 158, 187, 100, 187, 207, 235, 115, 254, 243, 161, 1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163, 0, 10, 10, 1, 139, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
			},
			wantTx: SpendTx{
				SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Amount:      utils.NewIntFromUint64(0),
				Fee:         utils.NewIntFromUint64(10),
				Payload:     []byte("Hello World"),
				TTL:         uint64(10),
				Nonce:       uint64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTx := SpendTx{}
			b := &bytes.Buffer{}
			b.Write(tt.args.rlpBytes)

			err := rlp.Decode(b, &gotTx)
			if err != nil {
				t.Error(err)
			}
			if !(reflect.DeepEqual(gotTx, tt.wantTx)) {
				t.Errorf("Deserialization resulted in different structs: got %+v, want %+v", gotTx, tt.wantTx)
			}
		})
	}
}
