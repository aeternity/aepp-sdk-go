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
		{
			name: "update 4 pointers",
			fields: fields{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers:  []string{"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi", "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v", "ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt", "ak_rHQAmJsLKC2u7Tr1htTGYxy2ga71AESM611tjGGfyUJmLbDYP"},
				NameTTL:   uint64(0),
				ClientTTL: uint64(6),
				Fee:       *utils.NewBigIntFromUint64(1),
				TTL:       5,
				Nonce:     5,
			},
			wantTx:  "tx_+QEVIgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zBaECXotHBnQ1m9xH62NJrWQAxdC6EOMi+s9dCP9xEydH6SgA+MjxjmFjY291bnRfcHVia2V5oQFv5wr11P3EEyoB8Vv8AoWK140cojJEja4CeC3rE+gY5/GOYWNjb3VudF9wdWJrZXmhAQkzfmKK/9rguLQf6vv/O43g1vpP+B727TdTmYbwitiB8Y5hY2NvdW50X3B1YmtleaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKPxjmFjY291bnRfcHVia2V5oQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wYBBaPTgbo=",
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

func TestNameRevokeTx_RLP(t *testing.T) {
	type fields struct {
		AccountID    string
		NameID       string
		Fee          utils.BigInt
		TTL          uint64
		AccountNonce uint64
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "normal revoke one name",
			fields: fields{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Fee:          *utils.NewBigIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantTx:  "tx_+EkjAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKAEFCjOeGw==",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewNameRevokeTx(
				tt.fields.AccountID,
				tt.fields.NameID,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.AccountNonce,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameRevokeTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("NameRevokeTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestNameTransferTx_RLP(t *testing.T) {
	type fields struct {
		AccountID    string
		NameID       string
		RecipientID  string
		Fee          utils.BigInt
		TTL          uint64
		AccountNonce uint64
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "normal name transfer transaction",
			fields: fields{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				RecipientID:  "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Fee:          *utils.NewBigIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantTx:  "tx_+GskAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKKEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBBeUht+4=",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewNameTransferTx(
				tt.fields.AccountID,
				tt.fields.NameID,
				tt.fields.RecipientID,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.AccountNonce,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameTransferTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("NameTransferTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
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
			name: "Oracle Register: A 0 in a BigInt field shouldn't cause a RLP serialization mismatch",
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
				abiVersion:     uint64(1),
				vmVersion:      uint64(0),
				txFee:          *utils.RequireBigIntFromString("200000000000000"),
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
				abiVersion:     uint64(0),
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
			name: "Fixed Value Oracle Extend, Extend by 300 blocks, delta",
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

func TestOracleQueryTx_RLP(t *testing.T) {
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
		Fee              utils.BigInt
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
				QueryFee:         *utils.NewBigIntFromUint64(0),
				QueryTTLType:     0,
				QueryTTLValue:    300,
				ResponseTTLType:  0,
				ResponseTTLValue: 300,
				Fee:              *utils.RequireBigIntFromString("200000000000000"),
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

func TestOracleRespondTx_RLP(t *testing.T) {
	type fields struct {
		OracleID         string
		AccountNonce     uint64
		QueryID          string
		Response         string
		ResponseTTLType  uint64
		ResponseTTLValue uint64
		Fee              utils.BigInt
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
				Fee:              *utils.RequireBigIntFromString("200000000000000"),
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
			gotTx, err := BaseEncodeTx(&tx)

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

func TestContractCreateTx_RLP(t *testing.T) {
	type fields struct {
		OwnerID      string
		AccountNonce uint64
		Code         string
		VMVersion    uint64
		AbiVersion   uint64
		Deposit      uint64
		Amount       utils.BigInt
		Gas          utils.BigInt
		GasPrice     utils.BigInt
		Fee          utils.BigInt
		TTL          uint64
		CallData     string
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Create Minerva: VMVersion 3, ABIVersion 1",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  uint64(3),
				AbiVersion: uint64(1),
				Deposit:    0,
				Amount:     *utils.NewBigInt(),
				Gas:        *utils.NewBigIntFromUint64(1e9),
				GasPrice:   *utils.NewBigIntFromUint64(1e9),
				Fee:        *utils.RequireBigIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwMAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKiXmeQ=",
			wantErr: false,
		},
		{
			name: "Fixed Value Contract Create Fortuna: VMVersion 4, ABIVersion 1",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  uint64(4),
				AbiVersion: uint64(1),
				Deposit:    0,
				Amount:     *utils.NewBigInt(),
				Gas:        *utils.NewBigIntFromUint64(1e9),
				GasPrice:   *utils.NewBigIntFromUint64(1e9),
				Fee:        *utils.RequireBigIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwQAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6baM=",
			wantErr: false,
		},
		{
			name: "Config Defaults Contract Create should be ok to post to a current node",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  uint64(4),
				AbiVersion: uint64(1),
				Deposit:    Config.Client.Contracts.Deposit,
				Amount:     Config.Client.Contracts.Amount,
				Gas:        Config.Client.Contracts.Gas,
				GasPrice:   Config.Client.Contracts.GasPrice,
				Fee:        Config.Client.Fee,
				TTL:        Config.Client.TTL,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwQAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6baM=",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewContractCreateTx(
				tt.fields.OwnerID,
				tt.fields.AccountNonce,
				tt.fields.Code,
				tt.fields.VMVersion,
				tt.fields.AbiVersion,
				tt.fields.Deposit,
				tt.fields.Amount,
				tt.fields.Gas,
				tt.fields.GasPrice,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.CallData,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCreateTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("ContractCreateTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestContractCreateTx_FeeEstimate(t *testing.T) {
	type fields struct {
		OwnerID      string
		AccountNonce uint64
		Code         string
		VMVersion    uint64
		AbiVersion   uint64
		Deposit      uint64
		Amount       utils.BigInt
		Gas          utils.BigInt
		GasPrice     utils.BigInt
		Fee          utils.BigInt
		TTL          uint64
		CallData     string
	}
	testCases := []struct {
		name    string
		fields  fields
		want    *utils.BigInt
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Create",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract SimpleStorage =\n  record state = { data : int }\n  function init(value : int) : state = { data = value }\n  function get() : int = state.data\n  function set(value : int) = put(state{data = value})"
				Code:       `cb_+QYYRgKgf6Gy7VnRXycsYSiFGAUHhMs+Oeg+RJvmPzCSAnxk8LT5BKX5AUmgOoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugeDc2V0uMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////////////////////////////////////////jJoEnsSQdsAgNxJqQzA+rc5DsuLDKUV7ETxQp+ItyJgJS3g2dldLhgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QKLoOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEhGluaXS4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQFEYgAAj2IAAMKRgICAUX9J7EkHbAIDcSakMwPq3OQ7LiwylFexE8UKfiLciYCUtxRiAAE5V1CAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAA0VdQgFF/OoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugcUYgABG1dQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tgAFFRkFZbYCABUVGQUIOSUICRUFCAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUVlQgJFQUGAAUYFZkIFSkFBgAFJZkFCQVltQUFlQUGIAAMpWhTIuMS4w4SWVhA==`,
				VMVersion:  3,
				AbiVersion: 1,
				Deposit:    0,
				Amount:     *utils.NewBigInt(),
				Gas:        *utils.NewBigIntFromUint64(1e9),
				GasPrice:   *utils.NewBigIntFromUint64(1e9),
				Fee:        *utils.RequireBigIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBo8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7dbVl",
			},
			want:    utils.NewBigIntFromUint64(120100000000000),
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCreateTx{
				OwnerID:      tt.fields.OwnerID,
				AccountNonce: tt.fields.AccountNonce,
				Code:         tt.fields.Code,
				VMVersion:    tt.fields.VMVersion,
				AbiVersion:   tt.fields.AbiVersion,
				Deposit:      tt.fields.Deposit,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
				CallData:     tt.fields.CallData,
			}
			got, err := tx.FeeEstimate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCreateTx.FeeEstimate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContractCreateTx.FeeEstimate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContractCallTx_RLP(t *testing.T) {
	type fields struct {
		CallerID     string
		AccountNonce uint64
		ContractID   string
		Amount       utils.BigInt
		Gas          utils.BigInt
		GasPrice     utils.BigInt
		AbiVersion   uint64
		VMVersion    uint64
		CallData     string
		Fee          utils.BigInt
		TTL          uint64
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Call Minerva (AbiVersion 3)",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(1),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *utils.NewBigIntFromUint64(10),
				Gas:          *utils.NewBigIntFromUint64(10),
				GasPrice:     *utils.NewBigIntFromUint64(10),
				AbiVersion:   uint64(3),
				VMVersion:    uint64(1),
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.RequireBigIntFromString("200000000000000"),
				TTL:          500,
			},
			wantTx:  "tx_+NcrAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwOGteYg9IAAggH0CgoKuIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKk9Ku98=",
			wantErr: false,
		},
		{
			name: "Fixed Value Contract Call Fortuna (AbiVersion 4)",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *utils.NewBigInt(),
				Gas:          *utils.NewBigIntFromUint64(1e9),
				GasPrice:     *utils.NewBigIntFromUint64(1e9),
				AbiVersion:   uint64(4),
				VMVersion:    uint64(1),
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.RequireBigIntFromString("200000000000000"),
				TTL:          500,
			},
			wantTx:  "tx_+N8rAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMCoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwSGteYg9IAAggH0AIQ7msoAhDuaygC4gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqty+KkQ==",
			wantErr: false,
		},
		{
			name: "Config Defaults Contract Call",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       Config.Client.Contracts.Amount,
				Gas:          Config.Client.Contracts.Gas,
				GasPrice:     Config.Client.Contracts.GasPrice,
				AbiVersion:   Config.Client.Contracts.ABIVersion,
				VMVersion:    Config.Client.Contracts.VMVersion,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          Config.Client.Fee,
				TTL:          Config.Client.TTL,
			},
			wantTx:  "tx_+N8rAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMCoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwGGteYg9IAAggH0AIQ7msoAhDuaygC4gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqwDDyuQ==",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCallTx{
				CallerID:     tt.fields.CallerID,
				AccountNonce: tt.fields.AccountNonce,
				ContractID:   tt.fields.ContractID,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				AbiVersion:   tt.fields.AbiVersion,
				VMVersion:    tt.fields.VMVersion,
				CallData:     tt.fields.CallData,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
			}
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := BaseEncodeTx(tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCallTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("ContractCallTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestContractCallTx_FeeEstimate(t *testing.T) {
	type fields struct {
		CallerID     string
		AccountNonce uint64
		ContractID   string
		Amount       utils.BigInt
		Gas          utils.BigInt
		GasPrice     utils.BigInt
		AbiVersion   uint64
		VMVersion    uint64
		CallData     string
		Fee          utils.BigInt
		TTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *utils.BigInt
		wantErr bool
	}{
		{
			name: "Fortuna (AbiVersion 4), with fixed params for fee estimation",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *utils.NewBigIntFromUint64(0),
				Gas:          *utils.NewBigIntFromUint64(1e5),
				GasPrice:     *utils.NewBigIntFromUint64(1e9),
				AbiVersion:   uint64(4),
				VMVersion:    uint64(1),
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.NewBigIntFromUint64(2e9),
				TTL:          0,
			},
			want:    utils.NewBigIntFromUint64(554440000000000),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCallTx{
				CallerID:     tt.fields.CallerID,
				AccountNonce: tt.fields.AccountNonce,
				ContractID:   tt.fields.ContractID,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				AbiVersion:   tt.fields.AbiVersion,
				VMVersion:    tt.fields.VMVersion,
				CallData:     tt.fields.CallData,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
			}
			got, err := tx.FeeEstimate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCallTx.FeeEstimate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContractCallTx.FeeEstimate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_encodeVMABI(t *testing.T) {
	type args struct {
		VMVersion  uint64
		ABIVersion uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// disabled because VMVersion 0 is only for oracles. Code is here to illustrate the node's behaviour.
		// {
		// 	name: "VMversion 0, AbiVersion 0",
		// 	args: args{
		// 		VMVersion:  0,
		// 		ABIVersion: 0,
		// 	},
		// 	want: []byte{0},
		// },
		// {
		// 	name: "VMversion 0, AbiVersion 1",
		// 	args: args{
		// 		VMVersion:  0,
		// 		ABIVersion: 1,
		// 	},
		// 	want: []byte{1},
		// },
		{
			name: "VMversion 1, AbiVersion 0",
			args: args{
				VMVersion:  1,
				ABIVersion: 0,
			},
			want: []byte{1, 0, 0},
		},
		{
			name: "VMversion 3, AbiVersion 1",
			args: args{
				VMVersion:  3,
				ABIVersion: 1,
			},
			want: []byte{3, 0, 1},
		},
		{
			name: "VMversion 5, AbiVersion 4",
			args: args{
				VMVersion:  5,
				ABIVersion: 4,
			},
			want: []byte{5, 0, 4},
		},
		{
			name: "VMversion 16, AbiVersion 16",
			args: args{
				VMVersion:  16,
				ABIVersion: 16,
			},
			want: []byte{16, 0, 16},
		},
		{
			name: "VMversion 255, AbiVersion 255",
			args: args{
				VMVersion:  255,
				ABIVersion: 255,
			},
			want: []byte{255, 0, 255},
		},
		{
			name: "VMversion 256, AbiVersion 255",
			args: args{
				VMVersion:  256,
				ABIVersion: 255,
			},
			want: []byte{1, 0, 0, 255},
		},
		{
			name: "VMversion 256, AbiVersion 256",
			args: args{
				VMVersion:  256,
				ABIVersion: 256,
			},
			want: []byte{1, 0, 1, 0},
		},
		{
			name: "VMversion 65535, AbiVersion 65535",
			args: args{
				VMVersion:  65535,
				ABIVersion: 65535,
			},
			want: []byte{255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeVMABI(tt.args.VMVersion, tt.args.ABIVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeVMABI() = %v, want %v", got, tt.want)
			}
		})
	}
}
