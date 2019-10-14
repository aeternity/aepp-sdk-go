package transactions

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/utils"
)

func TestContractTx(t *testing.T) {
	testCases := []struct {
		name     string
		tx       Transaction
		wantJSON string
		wantRLP  string
		wantErr  bool
	}{
		{
			name: "Fixed Value Contract Create Minerva: VMVersion 3, ABIVersion 1",
			tx: &ContractCreateTx{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  3,
				AbiVersion: 1,
				Deposit:    new(big.Int),
				Amount:     new(big.Int),
				GasLimit:   utils.NewIntFromUint64(1e9),
				GasPrice:   utils.NewIntFromUint64(1e9),
				Fee:        utils.RequireIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantJSON: `{"abi_version":1,"amount":0,"call_data":"cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==","code":"cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib","deposit":0,"fee":200000000000000,"gas":1000000000,"gas_price":1000000000,"nonce":1,"owner_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","ttl":500,"vm_version":3}`,
			wantRLP:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwMAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKiXmeQ=",
			wantErr:  false,
		},
		{
			name: "Fixed Value Contract Create Fortuna: VMVersion 4, ABIVersion 1",
			tx: &ContractCreateTx{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  4,
				AbiVersion: 1,
				Deposit:    new(big.Int),
				Amount:     new(big.Int),
				GasLimit:   utils.NewIntFromUint64(1e9),
				GasPrice:   utils.NewIntFromUint64(1e9),
				Fee:        utils.RequireIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantJSON: `{"abi_version":1,"amount":0,"call_data":"cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==","code":"cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib","deposit":0,"fee":200000000000000,"gas":1000000000,"gas_price":1000000000,"nonce":1,"owner_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","ttl":500,"vm_version":4}`,
			wantRLP:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwQAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6baM=",
			wantErr:  false,
		},
		{
			name: "Fixed Value Contract Call Minerva (AbiVersion 3)",
			tx: &ContractCallTx{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(1),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       utils.NewIntFromUint64(10),
				GasLimit:     utils.NewIntFromUint64(10),
				GasPrice:     utils.NewIntFromUint64(10),
				AbiVersion:   3,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          utils.RequireIntFromString("200000000000000"),
				TTL:          500,
			},
			wantJSON: `{"abi_version":3,"amount":10,"call_data":"cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li","caller_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","contract_id":"ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm","fee":200000000000000,"gas":10,"gas_price":10,"nonce":1,"ttl":500}`,
			wantRLP:  "tx_+NcrAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwOGteYg9IAAggH0CgoKuIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKk9Ku98=",
			wantErr:  false,
		},
		{
			name: "Fixed Value Contract Call Fortuna (AbiVersion 4)",
			tx: &ContractCallTx{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       new(big.Int),
				GasLimit:     utils.NewIntFromUint64(1e9),
				GasPrice:     utils.NewIntFromUint64(1e9),
				AbiVersion:   4,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          utils.RequireIntFromString("200000000000000"),
				TTL:          500,
			},
			wantJSON: `{"abi_version":4,"amount":0,"call_data":"cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li","caller_id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","contract_id":"ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm","fee":200000000000000,"gas":1000000000,"gas_price":1000000000,"nonce":2,"ttl":500}`,
			wantRLP:  "tx_+N8rAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMCoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwSGteYg9IAAggH0AIQ7msoAhDuaygC4gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqty+KkQ==",
			wantErr:  false,
		},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%s EncodeRLP", tt.name), func(t *testing.T) {
			gotRLP, err := SerializeTx(tt.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCreateTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
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

func Test_encodedecodeVMABI(t *testing.T) {
	type args struct {
		VMVersion  uint16
		ABIVersion uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// disabled because VMVersion 0 is only for oracles. Code is here to illustrate the node's behaviour.
		// {
		// 	name: "VMversion 0, AbiVersion 0",
		// 	args: args{
		// 		VMVersion:  0,
		// 		ABIVersion: 0,
		// 	},
		// 	want: []byte{0},
		// },
		// {
		// 	name: "VMversion 0, AbiVersion 1",
		// 	args: args{
		// 		VMVersion:  0,
		// 		ABIVersion: 1,
		// 	},
		// 	want: []byte{1},
		// },
		{
			name: "VMversion 1, AbiVersion 0",
			args: args{
				VMVersion:  1,
				ABIVersion: 0,
			},
			want: []byte{1, 0, 0},
		},
		{
			name: "VMversion 3, AbiVersion 1",
			args: args{
				VMVersion:  3,
				ABIVersion: 1,
			},
			want: []byte{3, 0, 1},
		},
		{
			name: "VMversion 5, AbiVersion 4",
			args: args{
				VMVersion:  5,
				ABIVersion: 4,
			},
			want: []byte{5, 0, 4},
		},
		{
			name: "VMversion 16, AbiVersion 16",
			args: args{
				VMVersion:  16,
				ABIVersion: 16,
			},
			want: []byte{16, 0, 16},
		},
		{
			name: "VMversion 255, AbiVersion 255",
			args: args{
				VMVersion:  255,
				ABIVersion: 255,
			},
			want: []byte{255, 0, 255},
		},
		{
			name: "VMversion 256, AbiVersion 255",
			args: args{
				VMVersion:  256,
				ABIVersion: 255,
			},
			want: []byte{1, 0, 0, 255},
		},
		{
			name: "VMversion 256, AbiVersion 256",
			args: args{
				VMVersion:  256,
				ABIVersion: 256,
			},
			want: []byte{1, 0, 1, 0},
		},
		{
			name: "VMversion 65535, AbiVersion 65535",
			args: args{
				VMVersion:  65535,
				ABIVersion: 65535,
			},
			want: []byte{255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("encodeVMABI %s", tt.name), func(t *testing.T) {
			if got := encodeVMABI(tt.args.VMVersion, tt.args.ABIVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeVMABI() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("decodeVMABI %s", tt.name), func(t *testing.T) {
			vm, abi := decodeVMABI(tt.want)
			if (vm != tt.args.VMVersion) || (abi != tt.args.ABIVersion) {
				t.Errorf("decodeVMABI() = got %v %v, want %v %v", vm, abi, tt.args.VMVersion, tt.args.ABIVersion)
			}
		})
	}
}
