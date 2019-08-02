package aeternity

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestOracleRegisterTx_EncodeRLP(t *testing.T) {
	type fields struct {
		accountID      string
		accountNonce   uint64
		querySpec      string
		responseSpec   string
		queryFee       big.Int
		oracleTTLType  uint64
		oracleTTLValue uint64
		abiVersion     uint16
		txFee          big.Int
		txTTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Oracle Register: A 0 in a BigInt field shouldn't cause a RLP serialization mismatch",
			fields: fields{
				accountID:      "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				accountNonce:   uint64(0),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       *utils.NewIntFromUint64(0),
				oracleTTLType:  uint64(0),
				oracleTTLValue: uint64(100),
				abiVersion:     0,
				txFee:          *utils.NewIntFromUint64(0),
				txTTL:          500,
			},
			wantTx:  "tx_+FgWAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZACCAfQA5kqYWQ==",
			wantErr: false,
		},
		{
			name: "Fixed Value Oracle Register",
			fields: fields{
				accountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				accountNonce:   uint64(1),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       Config.Client.Oracles.QueryFee,
				oracleTTLType:  0,
				oracleTTLValue: uint64(100),
				abiVersion:     1,
				txFee:          *utils.RequireIntFromString("200000000000000"),
				txTTL:          500,
			},
			// from the node's debug endpoint
			wantTx:  "tx_+F4WAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCAfQB0ylR9Q==",
			wantErr: false,
		},
		{
			name: "Config Defaults Oracle Register. Should be valid to post to an actual node.",
			fields: fields{
				accountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				accountNonce:   uint64(17),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       Config.Client.Oracles.QueryFee,
				oracleTTLType:  0,
				oracleTTLValue: uint64(100),
				abiVersion:     0,
				txFee:          Config.Client.Fee,
				txTTL:          uint64(50000),
			},
			// from the node's debug endpoint
			wantTx:  "tx_+F4WAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMRk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCw1AAwIXVNw==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewOracleRegisterTx(
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
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := SerializeTx(&tx)
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

func TestOracleExtendTx_EncodeRLP(t *testing.T) {
	type fields struct {
		OracleID     string
		AccountNonce uint64
		TTLType      uint64
		TTLValue     uint64
		Fee          big.Int
		TTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Oracle Extend, Extend by 300 blocks, delta",
			fields: fields{
				OracleID:     "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				TTLType:      0,
				TTLValue:     300,
				Fee:          *utils.NewIntFromUint64(10),
				TTL:          0,
			},
			// from the node's debug endpoint2
			wantTx:  "tx_6xkBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wEAggEsCgDoA8Ab",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewOracleExtendTx(
				tt.fields.OracleID,
				tt.fields.AccountNonce,
				tt.fields.TTLType,
				tt.fields.TTLValue,
				tt.fields.Fee,
				tt.fields.TTL,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := SerializeTx(&tx)

			if (err != nil) != tt.wantErr {
				t.Errorf("OracleExtendTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("OracleExtendTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestOracleQueryTx_EncodeRLP(t *testing.T) {
	type fields struct {
		SenderID         string
		AccountNonce     uint64
		OracleID         string
		Query            string
		QueryFee         big.Int
		QueryTTLType     uint64
		QueryTTLValue    uint64
		ResponseTTLType  uint64
		ResponseTTLValue uint64
		Fee              big.Int
		TTL              uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Values Oracle Query",
			fields: fields{
				SenderID:         "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				AccountNonce:     uint64(1),
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Query:            "Are you okay?",
				QueryFee:         *utils.NewIntFromUint64(0),
				QueryTTLType:     0,
				QueryTTLValue:    300,
				ResponseTTLType:  0,
				ResponseTTLValue: 300,
				Fee:              *utils.RequireIntFromString("200000000000000"),
				TTL:              500,
			},
			// from the node
			wantTx:  "tx_+GgXAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+841BcmUgeW91IG9rYXk/AACCASwAggEshrXmIPSAAIIB9GPfFkA=",
			wantErr: false,
		},
		{
			name: "Config Defaults Oracle Query",
			fields: fields{
				SenderID:         "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				AccountNonce:     uint64(1),
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Query:            "Are you okay?",
				QueryFee:         Config.Client.Oracles.QueryFee,
				QueryTTLType:     Config.Client.Oracles.QueryTTLType,
				QueryTTLValue:    Config.Client.Oracles.QueryTTLValue,
				ResponseTTLType:  Config.Client.Oracles.ResponseTTLType,
				ResponseTTLValue: Config.Client.Oracles.ResponseTTLValue,
				Fee:              Config.Client.Fee,
				TTL:              Config.Client.TTL,
			},
			// from the node
			wantTx:  "tx_+GgXAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+841BcmUgeW91IG9rYXk/AACCASwAggEshrXmIPSAAIIB9GPfFkA=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := OracleQueryTx{
				SenderID:         tt.fields.SenderID,
				AccountNonce:     tt.fields.AccountNonce,
				OracleID:         tt.fields.OracleID,
				Query:            tt.fields.Query,
				QueryFee:         tt.fields.QueryFee,
				QueryTTLType:     tt.fields.QueryTTLType,
				QueryTTLValue:    tt.fields.QueryTTLValue,
				ResponseTTLType:  tt.fields.ResponseTTLType,
				ResponseTTLValue: tt.fields.ResponseTTLValue,
				Fee:              tt.fields.Fee,
				TTL:              tt.fields.TTL,
			}
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := SerializeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("OracleQueryTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("OracleQueryTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestOracleRespondTx_EncodeRLP(t *testing.T) {
	type fields struct {
		OracleID         string
		AccountNonce     uint64
		QueryID          string
		Response         string
		ResponseTTLType  uint64
		ResponseTTLValue uint64
		Fee              big.Int
		TTL              uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Oracle Response",
			fields: fields{
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:     uint64(1),
				QueryID:          "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
				Response:         "Hello back",
				ResponseTTLType:  0,
				ResponseTTLValue: 100,
				Fee:              *utils.RequireIntFromString("200000000000000"),
				TTL:              500,
			},
			wantTx:  "tx_+F0YAaEEzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoLT1h6fjQDFn1a7j+6wVQ886V47xiFwvkbL+x2yR3J9cikhlbGxvIGJhY2sAZIa15iD0gACCAfQC7+L+",
			wantErr: false,
		},
		{
			name: "Config Defaults Oracle Response",
			fields: fields{
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:     uint64(1),
				QueryID:          "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
				Response:         "Hello back",
				ResponseTTLType:  0,
				ResponseTTLValue: 100,
				Fee:              Config.Client.Fee,
				TTL:              Config.Client.TTL,
			},
			wantTx:  "tx_+F0YAaEEzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoLT1h6fjQDFn1a7j+6wVQ886V47xiFwvkbL+x2yR3J9cikhlbGxvIGJhY2sAZIa15iD0gACCAfQC7+L+",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewOracleRespondTx(
				tt.fields.OracleID,
				tt.fields.AccountNonce,
				tt.fields.QueryID,
				tt.fields.Response,
				tt.fields.ResponseTTLType,
				tt.fields.ResponseTTLValue,
				tt.fields.Fee,
				tt.fields.TTL,
			)
			gotTx, err := SerializeTx(&tx)

			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("OracleRespondTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("OracleRespondTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}
