package aeternity

import (
	"reflect"
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
		{"genesis", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fced7958216f5e48655c06e189b996dfb5bad32db5c24b7a283eec8e96453acbe493b15a01872f26", "ak$me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTsYwGC"}, false, true},
		{"genesis", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fced7958216f5e48655c06e189b996dfb5bad32db5c24b7a283eec8e96453acbe493b15a01872f26", "ak$me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTs"}, false, false},
		{"genesis", args{"e", "ak$2a1j2Mk9YSmC1gioUq4PWRm3bsv88RVwyv4KaUGoR1eiKi"}, true, false},
		{"genesis", args{"9aaf28231d6c1f4a57bfdd834ee4080c7702106b9e117905fc", "ak$me6L5SSXL4NLWv5EkQ7a16xaA145Br7oV4sz9JphZgsTsYwGC"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp, err := Load(tt.args.pk)
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
	pub := "ak$Gd6iMVsoonGuTF8LeswwDDN2NF5wYHAoTRtzwdEcfS32LWoxm"
	kp, _ := Load(priv)

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
		keyPair       *KeyPair
		args          args
		wantSignature []byte
	}{
		{"ok", kp, args{txBinaryAsArray}, signatureAsArray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSignature := tt.keyPair.Sign(tt.args.message); !reflect.DeepEqual(gotSignature, tt.wantSignature) {
				t.Errorf("KeyPair.Sign() = %v, want %v", gotSignature, tt.wantSignature)
			}
		})
	}
}

func Test_decode(t *testing.T) {
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
		{"invalid", args{"ak$"}, []byte{}, true},
		{"invalid", args{"ak$aasda"}, []byte{}, true},
		{"invalid", args{"000000"}, []byte{}, true},
		{"invalid", args{"0"}, []byte{}, true},
		{"invalid", args{"ak$0"}, []byte{}, true},
		{"valid", args{"ak$Gd6iMVsoonGuTF8LeswwDDN2NF5wYHAoTRtzwdEcfS32LWoxm"}, []byte{35, 120, 248, 146, 183, 204, 130, 194, 210, 115, 158, 153, 78, 201, 149, 58, 163, 100, 97, 241, 235, 90, 74, 73, 165, 176, 222, 23, 179, 210, 58, 232}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := decode(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("decode() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_namehash(t *testing.T) {
	// ('welghmolql.aet') == 'nm$2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ'
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{"welghmolql.aet"}, "nm$2KrC4asc6fdv82uhXDwfiqB1TY2htjhnzwzJJKLxidyMymJRUQ"},
		{"ok", args{"welghmolql"}, "nm$2nLRBu1FyukEvJuMANjFzx8mubMFeyG2mJ2QpQoYKymYe1d2sr"},
		{"ok", args{""}, "nm$2q1DrgEuxRNCWRp5nTs6FyA7moSEzrPVUSTEpkpFsM4hRL4Dkb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeP(PrefixNameHash, namehash(tt.args.name))
			if got != tt.want {
				t.Errorf("namehash() = %v, want %v", got, tt.want)
			}
		})
	}
}
