package aeternity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/utils"
)

func getRLPSerialized(tx1 string, tx2 string) ([]interface{}, []interface{}) {
	tx1Bytes, _ := Decode(tx1)
	tx1RLP := DecodeRLPMessage(tx1Bytes)
	tx2Bytes, _ := Decode(tx2)
	tx2RLP := DecodeRLPMessage(tx2Bytes)
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
				gotIDBytes, _ := Decode(gotID)
				wantIDBytes, _ := Decode(tt.wantID)
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
