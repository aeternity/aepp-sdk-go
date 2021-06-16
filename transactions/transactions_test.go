package transactions

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v9/account"
	"github.com/aeternity/aepp-sdk-go/v9/config"

	"github.com/aeternity/aepp-sdk-go/v9/binary"
	"github.com/aeternity/aepp-sdk-go/v9/utils"
)

func getRLPSerialized(tx1 string, tx2 string) ([]interface{}, []interface{}) {
	tx1Bytes, _ := binary.Decode(tx1)
	tx1RLP := binary.DecodeRLPMessage(tx1Bytes)
	tx2Bytes, _ := binary.Decode(tx2)
	tx2RLP := binary.DecodeRLPMessage(tx2Bytes)
	return tx1RLP, tx2RLP
}

func TestSignedTx(t *testing.T) {
	tests := []struct {
		name    string
		tx      Transaction
		wantRLP string
		wantErr bool
	}{
		{
			name: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi signed SpendTx",
			tx: &SignedTx{
				Signatures: [][]byte{
					[]byte{10, 209, 197, 35, 33, 60, 73, 235, 31, 242, 68, 40, 83, 36, 49, 185, 210, 155, 146, 245, 148, 195, 118, 71, 232, 136, 84, 192, 104, 87, 114, 107, 26, 152, 167, 129, 192, 67, 213, 184, 220, 130, 126, 105, 22, 118, 228, 212, 198, 176, 0, 222, 210, 252, 185, 230, 18, 201, 238, 96, 105, 70, 40, 12},
				},
				Tx: &SpendTx{
					SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
					RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
					Amount:      utils.NewIntFromUint64(10),
					Fee:         utils.NewIntFromUint64(10),
					Payload:     []byte("Hello World"),
					TTL:         uint64(10),
					Nonce:       uint64(1),
				},
			},
			wantRLP: "tx_+KALAfhCuEAK0cUjITxJ6x/yRChTJDG50puS9ZTDdkfoiFTAaFdyaxqYp4HAQ9W43IJ+aRZ25NTGsADe0vy55hLJ7mBpRigMuFj4VgwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKCgoBi0hlbGxvIFdvcmxk+91GKg==",
			wantErr: false,
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

func Test_leftPadByteSlice(t *testing.T) {
	type args struct {
		length int
		data   []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Left pad a nonce of 3 to 32 bytes",
			args: args{
				length: 32,
				data:   []byte{3},
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3},
		},
		{
			name: "Left pad a multi-byte value to 32 bytes",
			args: args{
				length: 32,
				data:   []byte{1, 2, 3, 4, 3},
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftPadByteSlice(tt.args.length, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftPadByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildOracleQueryID(t *testing.T) {
	type args struct {
		sender      string
		senderNonce uint64
		recipient   string
	}
	tests := []struct {
		name    string
		args    args
		wantID  string
		wantErr bool
	}{
		{
			name: "a simple oracle query id",
			args: args{
				sender:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				senderNonce: uint64(3),
				recipient:   "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			},
			wantID:  "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
			wantErr: false,
		},
		{
			name: "this test case copied from aepp-middleware",
			args: args{
				sender:      "ak_2ZjpYpJbzq8xbzjgPuEpdq9ahZE7iJRcAYC1weq3xdrNbzRiP4",
				senderNonce: uint64(1),
				recipient:   "ok_2iqfJjbhGgJFRezjX6Q6DrvokkTM5niGEHBEJZ7uAG5fSGJAw1",
			},
			wantID:  "oq_2YvZnoohcSvbQCsPKSMxc98i5HZ1sU5mR6xwJUZC3SvkuSynMj",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := buildOracleQueryID(tt.args.sender, tt.args.senderNonce, tt.args.recipient)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildOracleQueryID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				gotIDBytes, _ := binary.Decode(gotID)
				wantIDBytes, _ := binary.Decode(tt.wantID)
				t.Errorf("buildOracleQueryID() = \n%v\n%v, want \n%v\n%v", gotID, gotIDBytes, tt.wantID, wantIDBytes)
			}
		})
	}
}

func Test_buildContractID(t *testing.T) {
	type args struct {
		sender      string
		senderNonce uint64
	}
	tests := []struct {
		name     string
		args     args
		wantCtID string
		wantErr  bool
	}{
		{
			name: "Genesis address, nonce 1",
			args: args{
				sender:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				senderNonce: uint64(1),
			},
			wantCtID: "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
			wantErr:  false,
		},
		{
			name: "Genesis address, nonce 5",
			args: args{
				sender:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				senderNonce: uint64(5),
			},
			wantCtID: "ct_223vybq7Ljr2VKaVhRyveFoSJMBZ8CyBCpPAFZ1BxgvMXggAA",
			wantErr:  false,
		},
		{
			name: "Genesis address, nonce 256",
			args: args{
				sender:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				senderNonce: uint64(256),
			},
			wantCtID: "ct_FT6XgwatDufGJ2RUaLkMmnebfVHNju5YK7cbjnbtby8LwdcJB",
			wantErr:  false,
		},
		{
			name: "Genesis address, nonce 65536",
			args: args{
				sender:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				senderNonce: uint64(65536),
			},
			wantCtID: "ct_vuq6dPXiAgMuGfVvFveL6j3kEPJC32orJmaG5zL1oHgT3WCLB",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtID, err := buildContractID(tt.args.sender, tt.args.senderNonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildContractID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCtID != tt.wantCtID {
				t.Errorf("buildContractID() = %v, want %v", gotCtID, tt.wantCtID)
			}
		})
	}
}

func Test_buildIDTag(t *testing.T) {
	type args struct {
		IDTag       uint8
		encodedHash string
	}
	tests := []struct {
		name    string
		args    args
		wantV   []uint8
		wantErr bool
	}{
		{
			name: "ID tag for Account",
			args: args{
				IDTag:       IDTagAccount,
				encodedHash: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			},
			wantV:   []uint8{1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163},
			wantErr: false,
		},
		{
			name: "ID tag for Name",
			args: args{
				IDTag:       IDTagName,
				encodedHash: "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb",
			},
			wantV:   []uint8{2, 94, 139, 71, 6, 116, 53, 155, 220, 71, 235, 99, 73, 173, 100, 0, 197, 208, 186, 16, 227, 34, 250, 207, 93, 8, 255, 113, 19, 39, 71, 233, 40},
			wantErr: false,
		},
		{
			name: "ID tag for Commitment",
			args: args{
				IDTag:       IDTagCommitment,
				encodedHash: "cm_2jJov6dn121oKkHo6TuWaAAL4ZEMonnCjpo8jatkCixrLG8Uc4",
			},
			wantV:   []uint8{3, 227, 194, 105, 213, 122, 105, 93, 105, 190, 173, 83, 176, 72, 82, 232, 179, 29, 29, 42, 62, 248, 117, 91, 32, 18, 194, 151, 177, 251, 210, 208, 193},
			wantErr: false,
		},
		{
			name: "ID tag for Oracle",
			args: args{
				IDTag:       IDTagOracle,
				encodedHash: "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			},
			wantV:   []uint8{4, 206, 167, 173, 228, 112, 201, 249, 157, 157, 78, 64, 8, 128, 168, 111, 29, 73, 187, 68, 75, 98, 241, 26, 158, 187, 100, 187, 207, 235, 115, 254, 243},
			wantErr: false,
		},
		{
			name: "ID tag for Contract",
			args: args{
				IDTag:       IDTagContract,
				encodedHash: "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
			},
			wantV:   []uint8{5, 239, 236, 68, 81, 186, 240, 95, 106, 155, 58, 111, 124, 149, 82, 169, 148, 80, 73, 134, 189, 169, 218, 37, 177, 128, 198, 72, 122, 183, 77, 248, 195},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, err := buildIDTag(tt.args.IDTag, tt.args.encodedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildIDTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("buildIDTag() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func Test_readIDTag(t *testing.T) {
	type args struct {
		v []uint8
	}
	tests := []struct {
		name            string
		args            args
		wantIDTag       uint8
		wantEncodedHash string
		wantErr         bool
	}{
		{
			name: "Read ID tag for Account",
			args: args{
				v: []uint8{1, 31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163},
			},
			wantIDTag:       IDTagAccount,
			wantEncodedHash: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			wantErr:         false,
		},
		{
			name: "Read ID tag for Name",
			args: args{
				v: []uint8{2, 94, 139, 71, 6, 116, 53, 155, 220, 71, 235, 99, 73, 173, 100, 0, 197, 208, 186, 16, 227, 34, 250, 207, 93, 8, 255, 113, 19, 39, 71, 233, 40},
			},
			wantIDTag:       IDTagName,
			wantEncodedHash: "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb",
			wantErr:         false,
		},
		{
			name: "Read ID tag for Commitment",
			args: args{
				v: []uint8{3, 227, 194, 105, 213, 122, 105, 93, 105, 190, 173, 83, 176, 72, 82, 232, 179, 29, 29, 42, 62, 248, 117, 91, 32, 18, 194, 151, 177, 251, 210, 208, 193},
			},
			wantIDTag:       IDTagCommitment,
			wantEncodedHash: "cm_2jJov6dn121oKkHo6TuWaAAL4ZEMonnCjpo8jatkCixrLG8Uc4",
			wantErr:         false,
		},
		{
			name: "Read ID tag for Oracle",
			args: args{
				v: []uint8{4, 206, 167, 173, 228, 112, 201, 249, 157, 157, 78, 64, 8, 128, 168, 111, 29, 73, 187, 68, 75, 98, 241, 26, 158, 187, 100, 187, 207, 235, 115, 254, 243},
			},
			wantIDTag:       IDTagOracle,
			wantEncodedHash: "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			wantErr:         false,
		},
		{
			name: "Read ID tag for Contract",
			args: args{
				v: []uint8{5, 239, 236, 68, 81, 186, 240, 95, 106, 155, 58, 111, 124, 149, 82, 169, 148, 80, 73, 134, 189, 169, 218, 37, 177, 128, 198, 72, 122, 183, 77, 248, 195},
			},
			wantIDTag:       IDTagContract,
			wantEncodedHash: "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
			wantErr:         false,
		},
		{
			name: "Unknown ID tag",
			args: args{
				v: []uint8{8, 239, 236, 68, 81, 186, 240, 95, 106, 155, 58, 111, 124, 149, 82, 169, 148, 80, 73, 134, 189, 169, 218, 37, 177, 128, 198, 72, 122, 183, 77, 248, 195},
			},
			wantIDTag:       0,
			wantEncodedHash: "",
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIDTag, gotEncodedHash, err := readIDTag(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("readIDTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIDTag != tt.wantIDTag {
				t.Errorf("readIDTag() gotIDTag = %v, want %v", gotIDTag, tt.wantIDTag)
			}
			if gotEncodedHash != tt.wantEncodedHash {
				t.Errorf("readIDTag() gotEncodedHash = %v, want %v", gotEncodedHash, tt.wantEncodedHash)
			}
		})
	}
}

func Test_CalculateFee(t *testing.T) {
	tests := []Transaction{
		&SpendTx{
			SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			Amount:      big.NewInt(10),
			Fee:         big.NewInt(10),
			Payload:     []byte("Hello World"),
			TTL:         uint64(10),
			Nonce:       uint64(1),
		},
		&NamePreclaimTx{
			AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			CommitmentID: "cm_2jrPGyFKCEFFrsVvQsUzfnSURV5igr2WxvMR679S5DnuFEjet4", // name: fdsa.test, salt: 12345
			Fee:          big.NewInt(10),
			TTL:          uint64(10),
			AccountNonce: uint64(1),
		},
		&NameClaimTx{
			AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			Name:         "fdsa.test",
			NameSalt:     utils.RequireIntFromString("9795159241593061970"),
			Fee:          utils.NewIntFromUint64(10),
			TTL:          uint64(10),
			AccountNonce: uint64(1),
		},
		&NameUpdateTx{
			AccountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			NameID:    "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
			Pointers: []*NamePointer{
				&NamePointer{Key: "account_pubkey", Pointer: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"},
			},
			NameTTL:      uint64(0),
			ClientTTL:    uint64(6),
			Fee:          utils.NewIntFromUint64(1),
			TTL:          5,
			AccountNonce: 5,
		},
		&NameTransferTx{
			AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
			RecipientID:  "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
			Fee:          utils.NewIntFromUint64(1),
			TTL:          5,
			AccountNonce: 5,
		},
		&NameRevokeTx{
			AccountID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			NameID:       "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb", // fdsa.test
			Fee:          utils.NewIntFromUint64(1),
			TTL:          5,
			AccountNonce: 5,
		},
		&OracleRegisterTx{
			AccountID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			AccountNonce:   uint64(1),
			QuerySpec:      "query Specification",
			ResponseSpec:   "response Specification",
			QueryFee:       config.Client.Oracles.QueryFee,
			OracleTTLType:  0,
			OracleTTLValue: uint64(100),
			AbiVersion:     1,
			Fee:            utils.RequireIntFromString("200000000000000"),
			TTL:            500,
		},
		&OracleExtendTx{
			OracleID:       "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			AccountNonce:   1,
			OracleTTLType:  0,
			OracleTTLValue: 300,
			Fee:            utils.NewIntFromUint64(10),
			TTL:            0,
		},
		&OracleQueryTx{
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
		&OracleRespondTx{
			OracleID:         "ok_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			AccountNonce:     uint64(1),
			QueryID:          "oq_2NhMjBdKHJYnQjDbAxanmxoXiSiWDoG9bqDgk2MfK2X6AB9Bwx",
			Response:         "Hello back",
			ResponseTTLType:  0,
			ResponseTTLValue: 100,
			Fee:              config.Client.Fee,
			TTL:              config.Client.TTL,
		},
		&ContractCreateTx{
			OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			AccountNonce: 1,
			// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
			Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
			VMVersion:  4,
			AbiVersion: 1,
			Deposit:    config.Client.Contracts.Deposit,
			Amount:     config.Client.Contracts.Amount,
			GasLimit:   config.Client.Contracts.GasLimit,
			GasPrice:   config.Client.GasPrice,
			Fee:        config.Client.Fee,
			TTL:        config.Client.TTL,
			CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
		},
		&ContractCallTx{
			CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			AccountNonce: uint64(2),
			ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
			Amount:       config.Client.Contracts.Amount,
			GasLimit:     config.Client.Contracts.GasLimit,
			GasPrice:     config.Client.GasPrice,
			AbiVersion:   config.Client.Contracts.ABIVersion,
			CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
			Fee:          config.Client.Fee,
			TTL:          config.Client.TTL,
		},
		&GAAttachTx{
			OwnerID:      "ak_oeoYuVx1wmPxSADDCY6GFVorfJHFYBKia9KonSiWjtbvNQv9Y",
			AccountNonce: 1,
			Code:         "cb_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9246K2XN/CAWn0QFGSjvFGIAAYlXUIBRf2vehD2STuvi98Sjn0btXkSUDTn62IZgndZlXQVNK20SFGIAAg5XUGABGVEAW2AAGVlgIAGQgVJgIJADYABZkIFSgVJZYCABkIFSYCCQA2AAWZCBUoFSWWAgAZCBUmAgkANgA4FSkFlgAFFZUmAAUmAA81tgAIBSYADzW2AA/ZBQkFZbYCABUVGQUFlQgJFQUGAAYABgAGEB9FmQgVJgAGAAWvGAUWAAFGIAASBXgFFgARRiAAFZV1BgARlRAFtQf05vdCBpbiBBdXRoIGNvbnRleHQAAAAAAAAAAAAAAAAAWWAgAZCBUmAgkANgE4FSkFBiAADdVltgIAFRYABRgGAgAVFZYCABkIFSYCCQA2ABYABRUQGBUpBQYABSWVBgAZBQkFCQVltgIAFRgFGQYCABUZFQWVCAgpJQklBQYABgAGAAg1lgIAGQgVJgIJADhYFSWWBAAZCBUmAgkANgABlZYCABkIFSYCCQA2AAWZCBUoFSWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSYCCQA2EBk4FSYABgAFrxkVBQkFZbYCABUVGDklCAkVBQgFlgIAGQgVJgIJADYAGBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYABZkIFSgVJZYCABkIFSYCCQA2ADgVKBUpBQkFaFMy4xLjA3jzdH",
			AuthFunc:     []byte{96, 0, 184, 90, 183, 93, 180, 49, 139, 122, 45, 0, 88, 192, 16, 175, 167, 85, 183, 183, 4, 214, 35, 12, 90, 196, 201, 236, 49, 59, 255, 71},
			VMVersion:    4,
			AbiVersion:   1,
			GasLimit:     big.NewInt(500),
			GasPrice:     big.NewInt(1000000000),
			Fee:          big.NewInt(126720000000000),
			TTL:          0,
			CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBr3oQ9kk7r4vfEo59G7V5ElA05+tiGYJ3WZV0FTSttEgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgae21UoMYpb0U8XZCTsTvUGUNeF/kxvl/87SxMDOBYASarBTN",
		},
	}
	for _, tt := range tests {
		ttType := reflect.TypeOf(tt).String()
		t.Run(ttType, func(t *testing.T) {
			err := CalculateFee(tt)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func ExampleSerializeTx() {
	tx := &SpendTx{
		SenderID:    "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
		RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
		Amount:      &big.Int{},
		Fee:         &big.Int{},
		Payload:     nil,
		TTL:         3627,
		Nonce:       3,
	}
	txStr, err := SerializeTx(tx)
	if err != nil {
		return
	}

	fmt.Println(txStr)
	// Output: tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAACCDisDgLzTETQ=
}

func ExampleDeserializeTx() {
	txRLP, err := binary.Decode("tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAACCDisDgLzTETQ=")
	if err != nil {
		return
	}
	tx, err := DeserializeTx(txRLP)
	fmt.Printf("%T %+v", tx, tx)
	//Output: *transactions.SpendTx &{SenderID:ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi RecipientID:ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v Amount:+0 Fee:+0 Payload:[] TTL:3627 Nonce:3}
}

func ExampleDeserializeTxStr() {
	tx, err := DeserializeTxStr("tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAACCDisDgLzTETQ=")
	if err != nil {
		return
	}
	fmt.Printf("%T, %+v", tx, tx)
	//Output: *transactions.SpendTx, &{SenderID:ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi RecipientID:ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v Amount:+0 Fee:+0 Payload:[] TTL:3627 Nonce:3}
}

func ExampleGetTransactionType() {
	txRLP, err := binary.Decode("tx_+E0MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjAACCDisDgLzTETQ=")
	if err != nil {
		return
	}
	tx, err := GetTransactionType(txRLP)
	fmt.Printf("%T, %+v", tx, tx)
	//Output: *transactions.SpendTx, &{SenderID: RecipientID: Amount:<nil> Fee:<nil> Payload:[] TTL:0 Nonce:0}
}

func ExampleSignHashTx() {
	acc, err := account.New()
	if err != nil {
		return
	}

	tx := &SpendTx{
		SenderID:    acc.Address, // If the SenderID differs from the signing account, the transaction will not validate.
		RecipientID: "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v",
		Amount:      &big.Int{},
		Fee:         &big.Int{},
		Payload:     nil,
		TTL:         3627,
		Nonce:       3,
	}

	stx, txhash, sig, err := SignHashTx(acc, tx, "ae_testnet")
	if err != nil {
		return
	}

	fmt.Println(stx, txhash, sig, err)
}

func ExampleVerifySignedTx() {
	valid, err := VerifySignedTx("ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi", "tx_+JcLAfhCuEB42YQL7o806SO319qTPOiHRPKPPwJpcMbPry9PrAMLVmAZWdoEQNY1Ly5Bo5A2br1MaDrss6zkeR6sotxbf/kCuE/4TQwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMAAIIOKwOA4+Wjcw==", "ae_testnet")
	if err != nil {
		return
	}
	fmt.Println(valid)
	//Output: true
}
