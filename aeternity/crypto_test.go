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
	type args struct {
		message []byte
	}
	tests := []struct {
		name          string
		fields        KeyPair
		args          args
		wantSignature []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyPair{
				SigningKey: tt.fields.SigningKey,
				Address:    tt.fields.Address,
			}
			if gotSignature := k.Sign(tt.args.message); !reflect.DeepEqual(gotSignature, tt.wantSignature) {
				t.Errorf("KeyPair.Sign() = %v, want %v", gotSignature, tt.wantSignature)
			}
		})
	}
}
