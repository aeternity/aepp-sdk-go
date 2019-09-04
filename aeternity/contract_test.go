package aeternity

import (
	"reflect"
	"testing"
)

func TestNewContractFromString(t *testing.T) {
	type args struct {
		cb string
	}
	tests := []struct {
		name    string
		args    args
		wantC   Contract
		wantErr bool
	}{
		{
			name: "Normal Contract",
			args: args{cb: "cb_+QOERgOgcPPaLbPhA5ZUYYI1viyMLIiudfF0RqUBtpvYh9wt5RL5Ao/5Aoyg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0SEaW5pdAC4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuMJiAAA5YgAAbJGAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAAsFdQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81uAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUYOSUICRUFBiAAB0Vok0LjAuMC1yYzEAsRFc7A=="},
			wantC: Contract{
				Tag:            70,
				RLPVersion:     3,
				SourceCodeHash: []byte{112, 243, 218, 45, 179, 225, 3, 150, 84, 97, 130, 53, 190, 44, 140, 44, 136, 174, 117, 241, 116, 70, 165, 1, 182, 155, 216, 135, 220, 45, 229, 18},
				TypeInfo: []ContractFunction{
					{
						FuncHash: []byte{226, 35, 29, 108, 223, 201, 57, 22, 222, 76, 179, 169, 133, 123, 246, 92, 244, 15, 194, 86, 244, 161, 73, 139, 63, 126, 124, 152, 12, 25, 147, 68},
						FuncName: "init",
						Payable:  false,
						ArgType:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
						OutType:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 128, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				Bytecode:        []byte{98, 0, 0, 57, 98, 0, 0, 108, 145, 128, 128, 81, 127, 226, 35, 29, 108, 223, 201, 57, 22, 222, 76, 179, 169, 133, 123, 246, 92, 244, 15, 194, 86, 244, 161, 73, 139, 63, 126, 124, 152, 12, 25, 147, 68, 20, 98, 0, 0, 176, 87, 80, 96, 1, 25, 81, 0, 91, 96, 0, 25, 89, 96, 32, 1, 144, 129, 82, 96, 32, 144, 3, 96, 0, 89, 144, 129, 82, 129, 82, 89, 96, 32, 1, 144, 129, 82, 96, 32, 144, 3, 96, 3, 129, 82, 144, 89, 96, 0, 81, 89, 82, 96, 0, 82, 96, 0, 243, 91, 96, 0, 128, 82, 96, 0, 243, 91, 128, 89, 144, 129, 82, 89, 96, 32, 1, 144, 129, 82, 96, 32, 144, 3, 96, 0, 25, 89, 96, 32, 1, 144, 129, 82, 96, 32, 144, 3, 96, 0, 89, 144, 129, 82, 129, 82, 89, 96, 32, 1, 144, 129, 82, 96, 32, 144, 3, 96, 3, 129, 82, 129, 82, 144, 80, 144, 86, 91, 96, 32, 1, 81, 81, 131, 146, 80, 128, 145, 80, 80, 98, 0, 0, 116, 86},
				CompilerVersion: "4.0.0-rc1",
				Payable:         false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := NewContractFromString(tt.args.cb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewContractFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("NewContractFromString() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
