package aeternity

import (
	"bytes"
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
			tx := NewSpendTx(tt.fields.senderID, tt.fields.recipientID,
				tt.fields.amount,
				tt.fields.fee,
				tt.fields.payload,
				tt.fields.ttl,
				tt.fields.nonce,
			)

			txJSON, err := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
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
			tx := NewNamePreclaimTx(
				tt.fields.AccountID,
				tt.fields.CommitmentID,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)

			txJSON, err := tx.JSON()
			fmt.Println(txJSON)
			gotTx, err := BaseEncodeTx(&tx)
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
		NameSalt  utils.BigInt
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
				Name:      "fdsa.test",
				NameSalt:  *utils.RequireBigIntFromString("9795159241593061970"),
				Fee:       *utils.NewBigIntFromUint64(10),
				TTL:       uint64(10),
				Nonce:     uint64(1),
			},
			wantTx:  "tx_+DogAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBiWZkc2EudGVzdIiH72Vu6YoCUgoKx4dL6Q==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewNameClaimTx(
				tt.fields.AccountID,
				tt.fields.Name,
				tt.fields.NameSalt,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)

			txJSON, err := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
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
		{
			name: "update 1 pointer",
			fields: fields{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers:  []string{"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"},
				NameTTL:   uint64(0),
				ClientTTL: uint64(6),
				Fee:       *utils.NewBigIntFromUint64(1),
				TTL:       5,
				Nonce:     5,
			},
			wantTx:  "tx_+H4iAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKADy8Y5hY2NvdW50X3B1YmtleaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMGAQXLBNnv",
			wantErr: false,
		},
		{
			name: "update 3 pointers",
			fields: fields{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers:  []string{"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi", "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v", "ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt"},
				NameTTL:   uint64(0),
				ClientTTL: uint64(6),
				Fee:       *utils.NewBigIntFromUint64(1),
				TTL:       5,
				Nonce:     5,
			},
			wantTx:  "tx_+OMiAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKAD4lvGOYWNjb3VudF9wdWJrZXmhAQkzfmKK/9rguLQf6vv/O43g1vpP+B727TdTmYbwitiB8Y5hY2NvdW50X3B1YmtleaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKPxjmFjY291bnRfcHVia2V5oQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wYBBYpSjmc=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewNameUpdateTx(
				tt.fields.AccountID,
				tt.fields.NameID,
				tt.fields.Pointers,
				tt.fields.NameTTL,
				tt.fields.ClientTTL,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.Nonce,
			)

			txJSON, err := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
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
				txTTL:          Config.Client.TTL,
			},
			wantTx:  "tx_+FgWAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZACCAfQA5kqYWQ==",
			wantErr: false,
		},
		{
			name: "A 'normal' OracleRegisterTx",
			fields: fields{
				accountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				accountNonce:   uint64(1),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       Config.Client.Oracles.QueryFee,
				oracleTTLType:  0,
				oracleTTLValue: uint64(100),
				abiVersion:     uint64(1),
				vmVersion:      uint64(0),
				txFee:          Config.Client.Fee,
				txTTL:          Config.Client.TTL,
			},
			// from the node's debug endpoint
			wantTx:  "tx_+F4WAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBk3F1ZXJ5IFNwZWNpZmljYXRpb26WcmVzcG9uc2UgU3BlY2lmaWNhdGlvbgAAZIa15iD0gACCAfQB0ylR9Q==",
			wantErr: false,
		},
		{
			name: "Should be valid to post to a private testnet",
			fields: fields{
				accountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				accountNonce:   uint64(17),
				querySpec:      "query Specification",
				responseSpec:   "response Specification",
				queryFee:       Config.Client.Oracles.QueryFee,
				oracleTTLType:  0,
				oracleTTLValue: uint64(100),
				abiVersion:     uint64(1),
				vmVersion:      uint64(0),
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
				tt.fields.vmVersion,
				tt.fields.txFee,
				tt.fields.txTTL,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
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

			gotTx, err := BaseEncodeTx(&tx)

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

func OracleQueryTxRLP(t *testing.T) {
	type fields struct {
		SenderID         string
		AccountNonce     uint64
		OracleID         string
		Query            string
		QueryFee         utils.BigInt
		QueryTTLType     uint64
		QueryTTLValue    uint64
		ResponseTTLType  uint64
		ResponseTTLValue uint64
		TxFee            utils.BigInt
		TxTTL            uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Normal query to an Oracle",
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
				TxFee:            Config.Client.Fee,
				TxTTL:            Config.Client.TTL,
			},
			// from aepp-sdk-js
			wantTx:  "tx_+GgXAaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBoQTOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+841BcmUgeW91IG9rYXk/AACCASwAggEshrXmIPSAAILEzIsypOc=",
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
				TxFee:            tt.fields.TxFee,
				TxTTL:            tt.fields.TxTTL,
			}
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
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

func TestNamePointer_EncodeRLP(t *testing.T) {
	type fields struct {
		ID  string
		Key string
	}
	tests := []struct {
		name    string
		fields  fields
		wantW   []byte
		wantErr bool
	}{
		{
			name: "1 pointer to a normal ak_ account",
			fields: fields{
				Key: "account_pubkey",
				ID:  "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			},
			// the reference value of wantW is taken from a correct serialization of NameUpdateTx.
			// Unfortunately there is no way to get the node to serialize just the NamePointer.
			wantW:   []byte{241, 142, 97, 99, 99, 111, 117, 110, 116, 95, 112, 117, 98, 107, 101, 121, 161, 1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewNamePointer(tt.fields.Key, tt.fields.ID)
			w := &bytes.Buffer{}
			if err := p.EncodeRLP(w); (err != nil) != tt.wantErr {
				t.Errorf("NamePointer.EncodeRLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); !bytes.Equal(gotW, tt.wantW) {
				t.Errorf("NamePointer.EncodeRLP() = %v, want %v", gotW, tt.wantW)
				fmt.Println(DecodeRLPMessage(gotW))
			}
		})
	}
}
