package aeternity

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/utils"
)

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

func TestNamePreclaimTx_EncodeRLP(t *testing.T) {
	type fields struct {
		AccountID    string
		CommitmentID string
		Fee          big.Int
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
				Fee:          *utils.NewIntFromUint64(10),
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
			gotTx, err := SerializeTx(&tx)
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

func TestNameClaimTx_EncodeRLP(t *testing.T) {
	type fields struct {
		AccountID string
		Name      string
		NameSalt  big.Int
		Fee       big.Int
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
				NameSalt:  *utils.RequireIntFromString("9795159241593061970"),
				Fee:       *utils.NewIntFromUint64(10),
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

			gotTx, err := SerializeTx(&tx)
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

func TestNameUpdateTx_EncodeRLP(t *testing.T) {
	type fields struct {
		AccountID string
		NameID    string
		Pointers  []string
		NameTTL   uint64
		ClientTTL uint64
		Fee       big.Int
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
				Fee:       *utils.NewIntFromUint64(1),
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
				Fee:       *utils.NewIntFromUint64(1),
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
				Fee:       *utils.NewIntFromUint64(1),
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

			gotTx, err := SerializeTx(&tx)

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

func TestNameRevokeTx_EncodeRLP(t *testing.T) {
	type fields struct {
		AccountID    string
		NameID       string
		Fee          big.Int
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
				Fee:          *utils.NewIntFromUint64(1),
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

			gotTx, err := SerializeTx(&tx)
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

func TestNameTransferTx_EncodeRLP(t *testing.T) {
	type fields struct {
		AccountID    string
		NameID       string
		RecipientID  string
		Fee          big.Int
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
				Fee:          *utils.NewIntFromUint64(1),
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

			gotTx, err := SerializeTx(&tx)
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
