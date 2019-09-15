package aeternity

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/utils"
	rlp "github.com/randomshinichi/rlpae"
)

func TestNamePointer(t *testing.T) {
	tests := []struct {
		name        string
		namepointer NamePointer
		rlpBytes    []byte
		wantErr     bool
	}{
		{
			name: "1 pointer to a normal ak_ account",
			namepointer: NamePointer{
				Key: "account_pubkey",
				ID:  "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			},
			// the reference value of rlpBytes is taken from a correct serialization of NameUpdateTx.
			// Unfortunately there is no way to get the node to serialize just the NamePointer.
			rlpBytes: []byte{241, 142, 97, 99, 99, 111, 117, 110, 116, 95, 112, 117, 98, 107, 101, 121, 161, 1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s EncodeRLP", tt.name), func(t *testing.T) {
			gotRLP, err := rlp.EncodeToBytes(&tt.namepointer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NamePointer.EncodeRLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(gotRLP, tt.rlpBytes) {
				t.Errorf("NamePointer.EncodeRLP() = %v, want %v", gotRLP, tt.rlpBytes)
				fmt.Println(DecodeRLPMessage(gotRLP))
			}
		})
		t.Run(fmt.Sprintf("%s DecodeRLP", tt.name), func(t *testing.T) {
			np := NamePointer{}
			err := rlp.DecodeBytes(tt.rlpBytes, &np)
			if err != nil {
				t.Errorf("NamePointer.DecodeRLP() error = %s", err)
			}
			if !(reflect.DeepEqual(tt.namepointer, np)) {
				t.Errorf("Deserialized NamePointer %+v does not deep equal %+v", np, tt.namepointer)
			}
		})
	}
}

func TestAENSTx(t *testing.T) {
	cases := []struct {
		name     string
		tx       Transaction
		wantJSON string
		wantRLP  string
		wantErr  bool
	}{
		{
			name: "NamePreclaimTx fdsa.test",
			tx: &NamePreclaimTx{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				CommitmentID: "cm_2jrPGyFKCEFFrsVvQsUzfnSURV5igr2WxvMR679S5DnuFEjet4", // name: fdsa.test, salt: 12345
				Fee:          utils.NewIntFromUint64(10),
				TTL:          uint64(10),
				AccountNonce: uint64(1),
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","commitment_id":"cm_2jrPGyFKCEFFrsVvQsUzfnSURV5igr2WxvMR679S5DnuFEjet4","fee":10,"nonce":1,"ttl":10}`,
			wantRLP:  "tx_+EkhAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoQPk/tyQN11szXxmy4KFOFRzfzopJGCmg7cv5B9SwaJs0goKoCk0Qg==",
			// [[33] [1] [1 206 167 173 228 112 201 249 157 157 78 64 8 128 168 111 29 73 187 68 75 98 241 26 158 187 100 187 207 235 115 254 243] [1] [3 228 254 220 144 55 93 108 205 124 102 203 130 133 56 84 115 127 58 41 36 96 166 131 183 47 228 31 82 193 162 108 210] [10] [10]]
			wantErr: false,
		},
		{
			name: "NameClaimTx claim fdsa.test",
			tx: &NameClaimTx{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				Name:         "fdsa.test",
				NameSalt:     utils.RequireIntFromString("9795159241593061970"),
				Fee:          utils.NewIntFromUint64(10),
				TTL:          uint64(10),
				AccountNonce: uint64(1),
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","fee":10,"name":"nm_9XeniQagC6u2QHpP8f","name_salt":9795159241593061970,"nonce":1,"ttl":10}`,
			wantRLP:  "tx_+DogAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBiWZkc2EudGVzdIiH72Vu6YoCUgoKx4dL6Q==",
			wantErr:  false,
		},
		{
			name: "NameUpdateTx update 1 pointer",
			tx: &NameUpdateTx{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers: []*NamePointer{
					&NamePointer{Key: "account_pubkey", ID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"},
				},
				NameTTL:      uint64(0),
				ClientTTL:    uint64(6),
				Fee:          utils.NewIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","client_ttl":6,"fee":1,"name_id":"nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb","name_ttl":0,"nonce":5,"pointers":[{"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","key":"account_pubkey"}],"ttl":5}`,
			wantRLP:  "tx_+H4iAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKADy8Y5hY2NvdW50X3B1YmtleaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMGAQXLBNnv",
			wantErr:  false,
		},
		{
			name: "NameUpdateTx update 3 pointers",
			tx: &NameUpdateTx{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers: []*NamePointer{
					&NamePointer{Key: "account_pubkey", ID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"},
					&NamePointer{Key: "account_pubkey", ID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"},
					&NamePointer{Key: "account_pubkey", ID: "ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt"},
				},
				NameTTL:      uint64(0),
				ClientTTL:    uint64(6),
				Fee:          utils.NewIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","client_ttl":6,"fee":1,"name_id":"nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb","name_ttl":0,"nonce":5,"pointers":[{"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","key":"account_pubkey"},{"id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","key":"account_pubkey"},{"id":"ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt","key":"account_pubkey"}],"ttl":5}`,
			wantRLP:  "tx_+OMiAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKAD4lvGOYWNjb3VudF9wdWJrZXmhAQkzfmKK/9rguLQf6vv/O43g1vpP+B727TdTmYbwitiB8Y5hY2NvdW50X3B1YmtleaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKPxjmFjY291bnRfcHVia2V5oQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wYBBYpSjmc=",
			wantErr:  false,
		},
		{
			name: "NameUpdateTx update 4 pointers",
			tx: &NameUpdateTx{
				AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Pointers: []*NamePointer{
					&NamePointer{Key: "account_pubkey", ID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"},
					&NamePointer{Key: "account_pubkey", ID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"},
					&NamePointer{Key: "account_pubkey", ID: "ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt"},
					&NamePointer{Key: "account_pubkey", ID: "ak_rHQAmJsLKC2u7Tr1htTGYxy2ga71AESM611tjGGfyUJmLbDYP"},
				},
				NameTTL:      uint64(0),
				ClientTTL:    uint64(6),
				Fee:          utils.NewIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","client_ttl":6,"fee":1,"name_id":"nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb","name_ttl":0,"nonce":5,"pointers":[{"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","key":"account_pubkey"},{"id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","key":"account_pubkey"},{"id":"ak_542o93BKHiANzqNaFj6UurrJuDuxU61zCGr9LJCwtTUg34kWt","key":"account_pubkey"},{"id":"ak_rHQAmJsLKC2u7Tr1htTGYxy2ga71AESM611tjGGfyUJmLbDYP","key":"account_pubkey"}],"ttl":5}`,
			wantRLP:  "tx_+QEVIgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zBaECXotHBnQ1m9xH62NJrWQAxdC6EOMi+s9dCP9xEydH6SgA+MjxjmFjY291bnRfcHVia2V5oQFv5wr11P3EEyoB8Vv8AoWK140cojJEja4CeC3rE+gY5/GOYWNjb3VudF9wdWJrZXmhAQkzfmKK/9rguLQf6vv/O43g1vpP+B727TdTmYbwitiB8Y5hY2NvdW50X3B1YmtleaEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKPxjmFjY291bnRfcHVia2V5oQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wYBBaPTgbo=",
			wantErr:  false,
		},
		{
			name: "NameTransferTx",
			tx: &NameTransferTx{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				RecipientID:  "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
				Fee:          utils.NewIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","fee":1,"name_id":"nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb","nonce":5,"recipient_id":"ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v","ttl":5}`,
			wantRLP:  "tx_+GskAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKKEBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMBBeUht+4=",
			wantErr:  false,
		},
		{
			name: "NameRevoke one name",
			tx: &NameRevokeTx{
				AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
				Fee:          utils.NewIntFromUint64(1),
				TTL:          5,
				AccountNonce: 5,
			},
			wantJSON: `{"account_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","fee":1,"name_id":"nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb","nonce":5,"ttl":5}`,
			wantRLP:  "tx_+EkjAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKAEFCjOeGw==",
			wantErr:  false,
		},
	}
	for _, tt := range cases {
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
