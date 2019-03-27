package aeternity_test

import (
	"fmt"
	"reflect"
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

			txJson, err := tx.JSON()
			fmt.Println(txJson)

			gotTx, err := aeternity.BaseEncodeTx(&tx)
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

func TestNamePreclaimTx_RLP(t *testing.T) {
	type fields struct {
		AccountID    string
		CommitmentID string
		Fee          utils.BigInt
		TTL          uint64
		Nonce        uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Normal procedure, reserve a name on AENS",
			fields: fields{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				CommitmentID: "cm_2jrPGyFKCEFFrsVvQsUzfnSURV5igr2WxvMR679S5DnuFEjet4", // name: fdsa.test, salt: 12345
				Fee:          *utils.NewBigIntFromUint64(10),
				TTL:          uint64(10),
				Nonce:        uint64(1),
			},
			wantTx:  "tx_+EkhAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoQPk/tyQN11szXxmy4KFOFRzfzopJGCmg7cv5B9SwaJs0goKoCk0Qg==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewNamePreclaimTx(
				tt.fields.AccountID,
				tt.fields.CommitmentID,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)

			txJson, err := tx.JSON()
			fmt.Println(txJson)
			gotTx, err := aeternity.BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NamePreclaimTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("NamePreclaimTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestNameClaimTx_RLP(t *testing.T) {
	type fields struct {
		AccountID string
		Name      string
		NameSalt  uint64
		Fee       utils.BigInt
		TTL       uint64
		Nonce     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Normal operation: claim a reserved name",
			fields: fields{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Name:      "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb",
				NameSalt:  12345,
				Fee:       *utils.NewBigIntFromUint64(10),
				TTL:       uint64(10),
				Nonce:     uint64(1),
			},
			wantTx:  "who knows",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewNameClaimTx(
				tt.fields.AccountID,
				tt.fields.Name,
				tt.fields.NameSalt,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)

			txJson, err := tx.JSON()
			fmt.Println(txJson)

			gotTx, err := aeternity.BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameClaimTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("NameClaimTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestNameUpdateTx_RLP(t *testing.T) {
	type fields struct {
		AccountID string
		NameID    string
		Pointers  []string
		NameTTL   uint64
		ClientTTL uint64
		Fee       utils.BigInt
		TTL       uint64
		Nonce     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewNameUpdateTx(
				tt.fields.AccountID,
				tt.fields.NameID,
				tt.fields.Pointers,
				tt.fields.NameTTL,
				tt.fields.ClientTTL,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)
			gotTx, err := aeternity.BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameUpdateTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("NameUpdateTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

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
		vmVersion      uint64
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
				vmVersion:      uint64(0),
				txFee:          *utils.NewBigIntFromUint64(0),
				txTTL:          aeternity.Config.Client.TTL,
			},
			wantTx:  "tx_+FgWAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZACCAfQA5kqYWQ==",
			wantErr: false,
		},
		{
			name: "A 'normal' OracleRegisterTx",
			fields: fields{
				accountID:      "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				accountNonce:   uint64(0),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       *utils.NewBigIntFromUint64(3000),
				oracleTTLType:  uint64(0),
				oracleTTLValue: uint64(100),
				abiVersion:     uint64(1),
				vmVersion:      uint64(0),
				txFee:          aeternity.Config.Client.Fee,
				txTTL:          aeternity.Config.Client.TTL,
			},
			wantTx:  "tx_+GAWAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvboILuABkhrXmIPSAAIIB9AErxxDN",
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
				tt.fields.vmVersion,
				tt.fields.txFee,
				tt.fields.txTTL,
			)
			txJson, _ := tx.JSON()
			fmt.Println(txJson)

			gotTx, err := aeternity.BaseEncodeTx(&tx)
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

func TestOracleExtendTx_RLP(t *testing.T) {
	type fields struct {
		OracleID     string
		AccountNonce uint64
		TTLType      uint64
		TTLValue     uint64
		Fee          utils.BigInt
		TTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Extend by 300 blocks, delta",
			fields: fields{
				OracleID:     "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				TTLType:      0,
				TTLValue:     300,
				Fee:          *utils.NewBigIntFromUint64(10),
			},
			wantTx:  "tx_6xkBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wEAggEsCgDoA8Ab",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := aeternity.NewOracleExtendTx(
				tt.fields.OracleID,
				tt.fields.AccountNonce,
				tt.fields.TTLType,
				tt.fields.TTLValue,
				tt.fields.Fee,
				tt.fields.TTL,
			)
			txJson, _ := tx.JSON()
			fmt.Println(txJson)

			gotTx, err := aeternity.BaseEncodeTx(&tx)

			if (err != nil) != tt.wantErr {
				t.Errorf("OracleExtendTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("OracleExtendTx.RLP() = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
