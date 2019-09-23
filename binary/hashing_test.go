package binary

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleDecode() {
	b, err := Decode("ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v")
	if err != nil {
		return
	}
	fmt.Println(b)
	// Output: [31 19 163 176 139 240 1 64 6 98 166 139 105 216 117 247 128 60 236 76 8 100 127 110 213 216 76 120 151 189 80 163]
}

func ExampleEncode() {
	addrB := []byte{31, 19, 163, 176, 139, 240, 1, 64, 6, 98, 166, 139, 105, 216, 117, 247, 128, 60, 236, 76, 8, 100, 127, 110, 213, 216, 76, 120, 151, 189, 80, 163}

	addr := Encode("ak_", addrB)
	fmt.Println(addr)
	// Output: ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v
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
