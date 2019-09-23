package models

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/v5/utils"
)

func TestOracleTx(t *testing.T) {
	tests := []struct {
		name     string
		tx       Transaction
		wantJSON string
		wantRLP  string
		wantErr  bool
	}{
		{
			name: "Oracle Register: A 0 in a BigInt field shouldn't cause a RLP serialization mismatch",
			tx: &OracleRegisterTx{
				AccountID:      "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				AccountNonce:   uint64(0),
				QuerySpec:      "query Specification",
				ResponseSpec:   "response Specification",
				QueryFee:       utils.NewIntFromUint64(0),
				OracleTTLType:  uint64(0),
				OracleTTLValue: uint64(100),
				AbiVersion:     0,
				Fee:            utils.NewIntFromUint64(0),
				TTL:            500,
			},
			wantJSON: `{"account_id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","fee":0,"oracle_ttl":{"type":"delta","value":100},"query_fee":0,"query_format":"query Specification","response_format":"response Specification","ttl":500}`,
			wantRLP:  "tx_+FgWAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZACCAfQA5kqYWQ==",
			wantErr:  false,
		},
		{
			name: "Fixed Value Oracle Register",
			tx: &OracleRegisterTx{
				AccountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:   uint64(1),
				QuerySpec:      "query Specification",
				ResponseSpec:   "response Specification",
				QueryFee:       config.Config.Client.Oracles.QueryFee,
				OracleTTLType:  0,
				OracleTTLValue: uint64(100),
				AbiVersion:     1,
				Fee:            utils.RequireIntFromString("200000000000000"),
				TTL:            500,
			},
			// from the node's debug endpoint
			wantJSON: `{"abi_version":1,"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","fee":200000000000000,"nonce":1,"oracle_ttl":{"type":"delta","value":100},"query_fee":0,"query_format":"query Specification","response_format":"response Specification","ttl":500}`,
			wantRLP:  "tx_+F4WAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCAfQB0ylR9Q==",
			wantErr:  false,
		},
		{
			name: "Config Defaults Oracle Register. Should be valid to post to an actual node.",
			tx: &OracleRegisterTx{
				AccountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:   uint64(17),
				QuerySpec:      "query Specification",
				ResponseSpec:   "response Specification",
				QueryFee:       config.Config.Client.Oracles.QueryFee,
				OracleTTLType:  0,
				OracleTTLValue: uint64(100),
				AbiVersion:     0,
				Fee:            config.Config.Client.Fee,
				TTL:            uint64(50000),
			},
			// from the node's debug endpoint
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","fee":200000000000000,"nonce":17,"oracle_ttl":{"type":"delta","value":100},"query_fee":0,"query_format":"query Specification","response_format":"response Specification","ttl":50000}`,
			wantRLP:  "tx_+F4WAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMRk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCw1AAwIXVNw==",
			wantErr:  false,
		},
		{
			name: "Fixed Value Oracle Extend, Extend by 300 blocks, delta",
			tx: &OracleExtendTx{
				OracleID:       "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:   1,
				OracleTTLType:  0,
				OracleTTLValue: 300,
				Fee:            utils.NewIntFromUint64(10),
				TTL:            0,
			},
			// from the node's debug endpoint2
			wantJSON: `{"fee":10,"nonce":1,"oracle_id":"ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","oracle_ttl":{"type":"delta","value":300}}`,
			wantRLP:  "tx_6xkBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wEAggEsCgDoA8Ab",
			wantErr:  false,
		},
		{
			name: "Fixed Values Oracle Query",
			tx: &OracleQueryTx{
				SenderID:         "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				AccountNonce:     uint64(1),
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Query:            "Are you okay?",
				QueryFee:         utils.NewIntFromUint64(0),
				QueryTTLType:     0,
				QueryTTLValue:    300,
				ResponseTTLType:  0,
				ResponseTTLValue: 300,
				Fee:              utils.RequireIntFromString("200000000000000"),
				TTL:              500,
			},
			// from the node
			wantJSON: `{"fee":200000000000000,"nonce":1,"oracle_id":"ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","query":"Are you okay?","query_fee":0,"query_ttl":{"type":"delta","value":300},"response_ttl":{"type":"delta","value":300},"sender_id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","ttl":500}`,
			wantRLP:  "tx_+GgXAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+841BcmUgeW91IG9rYXk/AACCASwAggEshrXmIPSAAIIB9GPfFkA=",
			wantErr:  false,
		},
		{
			name: "Config Defaults Oracle Query",
			tx: &OracleQueryTx{
				SenderID:         "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				AccountNonce:     uint64(1),
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Query:            "Are you okay?",
				QueryFee:         config.Config.Client.Oracles.QueryFee,
				QueryTTLType:     config.Config.Client.Oracles.QueryTTLType,
				QueryTTLValue:    config.Config.Client.Oracles.QueryTTLValue,
				ResponseTTLType:  config.Config.Client.Oracles.ResponseTTLType,
				ResponseTTLValue: config.Config.Client.Oracles.ResponseTTLValue,
				Fee:              config.Config.Client.Fee,
				TTL:              config.Config.Client.TTL,
			},
			// from the node
			wantJSON: `{"fee":200000000000000,"nonce":1,"oracle_id":"ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","query":"Are you okay?","query_fee":0,"query_ttl":{"type":"delta","value":300},"response_ttl":{"type":"delta","value":300},"sender_id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","ttl":500}`,
			wantRLP:  "tx_+GgXAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+841BcmUgeW91IG9rYXk/AACCASwAggEshrXmIPSAAIIB9GPfFkA=",
			wantErr:  false,
		},
		{
			name: "Fixed Value Oracle Response",
			tx: &OracleRespondTx{
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:     uint64(1),
				QueryID:          "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
				Response:         "Hello back",
				ResponseTTLType:  0,
				ResponseTTLValue: 100,
				Fee:              utils.RequireIntFromString("200000000000000"),
				TTL:              500,
			},
			wantJSON: `{"fee":200000000000000,"nonce":1,"oracle_id":"ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","query_id":"oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx","response":"Hello back","response_ttl":{"type":"delta","value":100},"ttl":500}`,
			wantRLP:  "tx_+F0YAaEEzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoLT1h6fjQDFn1a7j+6wVQ886V47xiFwvkbL+x2yR3J9cikhlbGxvIGJhY2sAZIa15iD0gACCAfQC7+L+",
			wantErr:  false,
		},
		{
			name: "Config Defaults Oracle Response",
			tx: &OracleRespondTx{
				OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce:     uint64(1),
				QueryID:          "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
				Response:         "Hello back",
				ResponseTTLType:  0,
				ResponseTTLValue: 100,
				Fee:              config.Config.Client.Fee,
				TTL:              config.Config.Client.TTL,
			},
			wantJSON: `{"fee":200000000000000,"nonce":1,"oracle_id":"ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","query_id":"oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx","response":"Hello back","response_ttl":{"type":"delta","value":100},"ttl":500}`,
			wantRLP:  "tx_+F0YAaEEzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoLT1h6fjQDFn1a7j+6wVQ886V47xiFwvkbL+x2yR3J9cikhlbGxvIGJhY2sAZIa15iD0gACCAfQC7+L+",
			wantErr:  false,
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
