package aeternity_test

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func getRLPSerialized(tx1 string, tx2 string) ([]interface{}, []interface{}) {
	tx1Bytes, _ := aeternity.Decode(tx1)
	tx1RLP := aeternity.DecodeRLPMessage(tx1Bytes)
	tx2Bytes, _ := aeternity.Decode(tx2)
	tx2RLP := aeternity.DecodeRLPMessage(tx2Bytes)
	return tx1RLP, tx2RLP
}
func TestSpendTx_RLP(t *testing.T) {
	type fields struct {
		senderID    string
		recipientID string
		amount      utils.BigInt
		fee         utils.BigInt
		payload     string
		ttl         uint64
		nonce       uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Spend 10, Fee 10, Hello World",
			fields: fields{
				senderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				recipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				amount:      *utils.NewBigIntFromUint64(10),
				fee:         *utils.NewBigIntFromUint64(10),
				payload:     "Hello World",
				ttl:         uint64(10),
				nonce:       uint64(1),
			},
			wantTx:  "tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCgoKAYtIZWxsbyBXb3JsZPSZjdM=",
			wantErr: false,
		},
		{
			name: "Spend 0, Fee 10, Hello World (check correct RLP serialization of 0)",
			fields: fields{
				senderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				recipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				amount:      *utils.NewBigIntFromUint64(0),
				fee:         *utils.NewBigIntFromUint64(10),
				payload:     "Hello World",
				ttl:         uint64(10),
				nonce:       uint64(1),
			},
			wantTx:  "tx_+FYMAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAAoKAYtIZWxsbyBXb3JsZICI5/w=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewSpendTx(tt.fields.senderID, tt.fields.recipientID,
				tt.fields.amount,
				tt.fields.fee,
				tt.fields.payload,
				tt.fields.ttl,
				tt.fields.nonce,
			)
			gotTx, err := aeternity.BaseEncodeTx(tx)
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

// func TestNamePreclaimTx_RLP(t *testing.T) {
// 	type fields struct {
// 		accountID    string
// 		commitmentID string
// 		fee          uint64
// 		ttl          uint64
// 		nonce        uint64
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantTx  string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tx := aeternity.NewNamePreclaimTx(
// 				tt.fields.accountID,
// 				tt.fields.commitmentID,
// 				tt.fields.fee,
// 				tt.fields.ttl,
// 				tt.fields.nonce,
// 			)
// 			gotTx, err := aeternity.BaseEncodeTx(tx)()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NamePreclaimTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotTx != tt.wantTx {
// 				t.Errorf("NamePreclaimTx.RLP() = %v, want %v", gotTx, tt.wantTx)
// 			}
// 		})
// 	}
// }

// func TestNameClaimTx_RLP(t *testing.T) {
// 	type fields struct {
// 		accountID string
// 		name      string
// 		nameSalt  uint64
// 		fee       uint64
// 		ttl       uint64
// 		nonce     uint64
// 	}
// 	tests := []struct {
// 		name          string
// 		fields        fields
// 		wantRlpRawMsg []byte
// 		wantErr       bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t := &NameClaimTx{
// 				accountID: tt.fields.accountID,
// 				name:      tt.fields.name,
// 				nameSalt:  tt.fields.nameSalt,
// 				fee:       tt.fields.fee,
// 				ttl:       tt.fields.ttl,
// 				nonce:     tt.fields.nonce,
// 			}
// 			gotRlpRawMsg, err := t.RLP()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NameClaimTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotRlpRawMsg, tt.wantRlpRawMsg) {
// 				t.Errorf("NameClaimTx.RLP() = %v, want %v", gotRlpRawMsg, tt.wantRlpRawMsg)
// 			}
// 		})
// 	}
// }

// func TestNameUpdateTx_RLP(t *testing.T) {
// 	type fields struct {
// 		accountID string
// 		nameID    string
// 		pointers  []string
// 		nameTTL   uint64
// 		clientTTL uint64
// 		fee       uint64
// 		ttl       uint64
// 		nonce     uint64
// 	}
// 	tests := []struct {
// 		name          string
// 		fields        fields
// 		wantRlpRawMsg []byte
// 		wantErr       bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t := &NameUpdateTx{
// 				accountID: tt.fields.accountID,
// 				nameID:    tt.fields.nameID,
// 				pointers:  tt.fields.pointers,
// 				nameTTL:   tt.fields.nameTTL,
// 				clientTTL: tt.fields.clientTTL,
// 				fee:       tt.fields.fee,
// 				ttl:       tt.fields.ttl,
// 				nonce:     tt.fields.nonce,
// 			}
// 			gotRlpRawMsg, err := t.RLP()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NameUpdateTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotRlpRawMsg, tt.wantRlpRawMsg) {
// 				t.Errorf("NameUpdateTx.RLP() = %v, want %v", gotRlpRawMsg, tt.wantRlpRawMsg)
// 			}
// 		})
// 	}
// }

func TestOracleRegisterTx_RLP(t *testing.T) {
	type fields struct {
		accountID      string
		accountNonce   uint64
		querySpec      string
		responseSpec   string
		queryFee       utils.BigInt
		oracleTTLType  uint64
		oracleTTLValue uint64
		abiVersion     uint64
		txFee          utils.BigInt
		txTTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "A 0 in a BigInt field shouldn't cause a RLP serialization mismatch",
			fields: fields{
				accountID:      "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				accountNonce:   uint64(0),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       *utils.NewBigIntFromUint64(0),
				oracleTTLType:  uint64(0),
				oracleTTLValue: uint64(100),
				abiVersion:     uint64(0),
				txFee:          *utils.NewBigIntFromUint64(0),
				txTTL:          aeternity.Config.Client.TTL,
			},
			wantTx:  "tx_+F4WAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCAfQAZpU79A==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewOracleRegisterTx(
				tt.fields.accountID,
				tt.fields.accountNonce,
				tt.fields.querySpec,
				tt.fields.responseSpec,
				tt.fields.queryFee,
				tt.fields.oracleTTLType,
				tt.fields.oracleTTLValue,
				tt.fields.abiVersion,
				tt.fields.txFee,
				tt.fields.txTTL,
			)
			gotTx, err := aeternity.BaseEncodeTx(tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("OracleRegisterTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTx != tt.wantTx {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("OracleRegisterTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

// func TestOracleExtendTx_RLP(t *testing.T) {
// 	type fields struct {
// 		oracleID     string
// 		accountNonce uint64
// 		ttlType      uint64
// 		ttlValue     uint64
// 		fee          utils.BigInt
// 		ttl          uint64
// 	}
// 	tests := []struct {
// 		name          string
// 		fields        fields
// 		wantRlpRawMsg []byte
// 		wantErr       bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t := &OracleExtendTx{
// 				oracleID:     tt.fields.oracleID,
// 				accountNonce: tt.fields.accountNonce,
// 				ttlType:      tt.fields.ttlType,
// 				ttlValue:     tt.fields.ttlValue,
// 				fee:          tt.fields.fee,
// 				ttl:          tt.fields.ttl,
// 			}
// 			gotRlpRawMsg, err := t.RLP()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("OracleExtendTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotRlpRawMsg, tt.wantRlpRawMsg) {
// 				t.Errorf("OracleExtendTx.RLP() = %v, want %v", gotRlpRawMsg, tt.wantRlpRawMsg)
// 			}
// 		})
// 	}
// }
