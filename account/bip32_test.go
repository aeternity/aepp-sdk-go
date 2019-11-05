package account

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestNewMasterKey(t *testing.T) {
	tests := []struct {
		name          string
		seed          string
		wantMasterKey string
		wantErr       bool
	}{
		{
			name:          "ring defense obey exhaust boat popular surround supreme front lemon monster number",
			seed:          "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			wantMasterKey: "0e5e20500f73bb98d6ca0f01ed1623b598de7947905b1d3f89604af9fad2bf58",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seed, err := hex.DecodeString(tt.seed)
			if err != nil {
				t.Error(err)
			}
			got, err := NewMasterKey(seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMasterKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if hexify(got.Key) != tt.wantMasterKey {
				t.Errorf("NewMasterKey() = %v, want %v", got, tt.wantMasterKey)
			}
		})
	}
}

func TestKey_NewChildKey(t *testing.T) {
	type args struct {
		childIdx uint32
	}
	tests := []struct {
		name       string
		seed       string
		args       args
		wantKeyHex string
		want       *Key
		wantErr    bool
	}{
		{
			name: "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39 m/0'",
			args: args{
				childIdx: FirstHardenedChild,
			},
			seed:       "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			wantKeyHex: "548fe88d6c3a392066f21cabc2883fadfe8ea4e232c1898e07f07f0a95aec201",
			want: &Key{
				Key:         []byte{84, 143, 232, 141, 108, 58, 57, 32, 102, 242, 28, 171, 194, 136, 63, 173, 254, 142, 164, 226, 50, 193, 137, 142, 7, 240, 127, 10, 149, 174, 194, 1},
				ChildNumber: 2147483648,
				ChainCode:   []byte{168, 183, 22, 32, 11, 110, 109, 130, 79, 20, 114, 55, 181, 157, 10, 36, 252, 45, 205, 134, 168, 91, 223, 73, 252, 140, 76, 87, 99, 213, 254, 207},
				Depth:       1,
				IsPrivate:   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seedBytes, err := hex.DecodeString(tt.seed)
			if err != nil {
				t.Fatal(err)
			}
			key, err := NewMasterKey(seedBytes)
			if err != nil {
				t.Fatal(err)
			}
			got, err := key.NewChildKey(tt.args.childIdx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.NewChildKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Key.NewChildKey() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
