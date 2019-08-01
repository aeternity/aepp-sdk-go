package aeternity

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	type args struct {
		pk string
		ak string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		keyMatch bool
	}{
		{"account", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fced7958216f5e48655c06e189b996dfb5bad32db5c24b7a283eec8e96453acbe493b15a01872f26", "ak_me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTsYwGC"}, false, true},
		{"account", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fced7958216f5e48655c06e189b996dfb5bad32db5c24b7a283eec8e96453acbe493b15a01872f26", "ak_me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTs"}, false, false},
		{"account", args{"e", "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv88RVwyv4KaUGoR1eiKi"}, true, false},
		{"account", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fc", "ak_me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTsYwGC"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp, err := AccountFromHexString(tt.args.pk)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.keyMatch && tt.args.ak != kp.Address {
				t.Errorf("Load() expected = %v, got %v", tt.args.ak, kp.Address)
			}
		})
	}
}

func TestKeyPair_Sign(t *testing.T) {

	priv := "4d881dd1917036cc231f9881a0db978c8899dd76a817252418606b02bf6ab9d22378f892b7cc82c2d2739e994ec9953aa36461f1eb5a4a49a5b0de17b3d23ae8"
	pub := "ak_Gd6iMVsoonGuTF8LeswwDDN2NF5wYHAoTRtzwdEcfS32LWoxm"
	kp, _ := AccountFromHexString(priv)

	if kp.Address != pub {
		t.Errorf("Load() expected = %v, got %v", pub, kp.Address)
	}

	txBinaryAsArray := []byte{248, 76, 12, 1, 160, 35, 120, 248, 146, 183, 204, 130, 194, 210, 115, 158, 153, 78, 201, 149, 58, 163, 100, 97, 241, 235, 90, 74, 73, 165, 176, 222, 23, 179, 210, 58, 232, 160, 63, 40, 35, 12, 40, 65, 38, 215, 218, 236, 136, 133, 42, 120, 160, 179, 18, 191, 241, 162, 198, 203, 209, 173, 89, 136, 202, 211, 158, 59, 12, 122, 1, 1, 1, 132, 84, 101, 115, 116}
	signatureAsArray := []byte{95, 146, 31, 37, 95, 194, 36, 76, 58, 49, 167, 156, 127, 131, 142, 248, 25, 121, 139, 109, 59, 243, 203, 205, 16, 172, 115, 143, 254, 236, 33, 4, 43, 46, 16, 190, 46, 46, 140, 166, 76, 39, 249, 54, 38, 27, 93, 159, 58, 148, 67, 198, 81, 206, 106, 237, 91, 131, 27, 14, 143, 178, 130, 2}

	type args struct {
		message []byte
	}
	tests := []struct {
		name          string
		account       *Account
		args          args
		wantSignature []byte
	}{
		{"ok", kp, args{txBinaryAsArray}, signatureAsArray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSignature := tt.account.Sign(tt.args.message); !reflect.DeepEqual(gotSignature, tt.wantSignature) {
				t.Errorf("KeyPair.Sign() = %v, want %v", gotSignature, tt.wantSignature)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{"invalid", args{"c"}, []byte{}, true},
		{"invalid", args{"cd"}, []byte{}, true},
		{"invalid", args{"ak_"}, []byte{}, true},
		{"invalid", args{"ak_aasda"}, []byte{}, true},
		{"invalid", args{"000000"}, []byte{}, true},
		{"invalid", args{"0"}, []byte{}, true},
		{"invalid", args{"ak_0"}, []byte{}, true},
		{"valid", args{"ak_Gd6iMVsoonGuTF8LeswwDDN2NF5wYHAoTRtzwdEcfS32LWoxm"}, []byte{35, 120, 248, 146, 183, 204, 130, 194, 210, 115, 158, 153, 78, 201, 149, 58, 163, 100, 97, 241, 235, 90, 74, 73, 165, 176, 222, 23, 179, 210, 58, 232}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := Decode(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Decode() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_Namehash(t *testing.T) {
	// ('welghmolql.aet') == 'nm_2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ'
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{"welghmolql.aet"}, "nm_2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ"},
		{"ok", args{"welghmolql"}, "nm_2nLRBu1FyukEvJuMANjFzx8mubMFeyG2mJ2QpQoYKymYe1d2sr"},
		{"ok", args{"fdsa.test"}, "nm_ie148R2qZYBfo1Ek3sZpfTLwBhkkqCRKi2Ce8JJ7yyWVRw2Sb"},
		{"ok", args{""}, "nm_2q1DrgEuxRNCWRp5nTs6FyA7moSEzrPVUSTEpkpFsM4hRL4Dkb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(PrefixName, Namehash(tt.args.name))
			if got != tt.want {
				t.Errorf("Namehash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadFromFile(t *testing.T) {
	// 7d7d43238efe877a76371a23886f7c9924d8ba35dc6845d9db50b7a906a44c5311f23e7f2b4f46a4cca4d6a7ff7b5770adacf4460dab5d24dac35fcfc8b776e3
	// ak_8uQyHDwrW9CzWSptWrLDzJho7NwdSA76B4rpuwEuWtFj8nn4R
	fixtureAddres := "ak_Jt6AzQEiXiEMFXum8NtTXcCQtE9P1RfpkeVSZX87pFddzzynW"

	type args struct {
		path     string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantAc  string
		wantErr bool
	}{
		{"ok", args{"testdata/keystore.json", "aeternity"}, fixtureAddres, false},
		{"no", args{"testdata/keystore.json", "wrongpwd"}, fixtureAddres, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadAccountFromKeyStoreFile(tt.args.path, tt.args.password)
			if tt.wantErr {
				if err == nil {
					t.Errorf("LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !tt.wantErr {
				if err != nil {
					t.Errorf("LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					return
				}
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKp, err := NewAccount()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasPrefix(gotKp.Address, "ak_") {
				t.Errorf("Generate() error = %v", gotKp.Address)
				return
			}
		})
	}
}

func Test_computeCommitmentID(t *testing.T) {
	type args struct {
		name string
		salt []byte
	}
	tests := []struct {
		name    string
		args    args
		wantCh  string
		wantErr bool
	}{
		{
			name: "fdsa.test, 0",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			wantCh:  "cm_2jJov6dn121oKkHo6TuWaAAL4ZEMonnCjpo8jatkCixrLG8Uc4",
			wantErr: false,
		},
		{
			name: "fdsa.test, 255",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255},
			},
			wantCh:  "cm_sa8UUjorPzCTLfYp6YftR4jwF4kPaZVsoP5bKVAqRw9zm43EE",
			wantErr: false,
		},
		{
			// erlang Eshell: rp(<<9795159241593061970:256>>).
			name: "fdsa.test, 9795159241593061970 (do not use Golang to convert salt integers)",
			args: args{
				name: "fdsa.test",
				salt: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 135, 239, 101, 110, 233, 138, 2, 82},
			},
			wantCh:  "cm_QhtcYow8krP3xQSTsAhFihfBstTjQMiApaPCgZuciDHZmMNtZ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fmt.Println(saltBytes)
			gotCh, err := computeCommitmentID(tt.args.name, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeCommitmentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCh != tt.wantCh {
				t.Errorf("computeCommitmentID() = %v, want %v", gotCh, tt.wantCh)
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
