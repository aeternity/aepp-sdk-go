package aeternity_test

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestSpendTx(t *testing.T) {
	type args struct {
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
		args    args
		wantTx  string
		wantErr bool
	}{
		{
			name: "Spend 10, Fee 10, Hello World",
			args: args{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTx, err := aeternity.SpendTx(tt.args.senderID, tt.args.recipientID, tt.args.amount, tt.args.fee, tt.args.payload, tt.args.ttl, tt.args.nonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpendTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotTxStr := aeternity.Encode(aeternity.PrefixTransaction, gotTx)
			if gotTxStr != tt.wantTx {
				t.Errorf("SpendTx() = %v, want %v", gotTxStr, tt.wantTx)
			}
		})
	}
}

func TestOracleRegisterTx(t *testing.T) {
	type args struct {
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
		args    args
		wantTx  string
		wantErr bool
	}{
		{
			name: "A 0 in a BigInt field shouldn't cause a RLP serialization mismatch",
			args: args{
				accountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				accountNonce:   uint64(0),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       *utils.NewBigIntFromUint64(0),
				oracleTTLType:  uint64(0),
				oracleTTLValue: uint64(100),
				abiVersion:     uint64(0),
				txFee:          aeternity.Config.Client.Fee,
				txTTL:          aeternity.Config.Client.TTL,
			},
			wantTx:  "tx_+F4WAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCAfQAZpU79A==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txRaw, err := aeternity.OracleRegisterTx(tt.args.accountID, tt.args.accountNonce, tt.args.querySpec, tt.args.responseSpec, tt.args.queryFee, tt.args.oracleTTLType, tt.args.oracleTTLValue, tt.args.abiVersion, tt.args.txFee, tt.args.txTTL)
			if (err != nil) != tt.wantErr {
				t.Errorf("OracleRegisterTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tx := aeternity.Encode(aeternity.PrefixTransaction, txRaw)
			if tx != tt.wantTx {
				t.Errorf("OracleRegisterTx() = %v, want %v", tx, tt.wantTx)
			}
		})
	}
}
