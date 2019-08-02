package aeternity

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestContractCreateTx_EncodeRLP(t *testing.T) {
	type fields struct {
		OwnerID      string
		AccountNonce uint64
		Code         string
		VMVersion    uint16
		AbiVersion   uint16
		Deposit      big.Int
		Amount       big.Int
		Gas          big.Int
		GasPrice     big.Int
		Fee          big.Int
		TTL          uint64
		CallData     string
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Create Minerva: VMVersion 3, ABIVersion 1",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  3,
				AbiVersion: 1,
				Deposit:    *new(big.Int),
				Amount:     *new(big.Int),
				Gas:        *utils.NewIntFromUint64(1e9),
				GasPrice:   *utils.NewIntFromUint64(1e9),
				Fee:        *utils.RequireIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwMAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKiXmeQ=",
			wantErr: false,
		},
		{
			name: "Fixed Value Contract Create Fortuna: VMVersion 4, ABIVersion 1",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  4,
				AbiVersion: 1,
				Deposit:    *new(big.Int),
				Amount:     *new(big.Int),
				Gas:        *utils.NewIntFromUint64(1e9),
				GasPrice:   *utils.NewIntFromUint64(1e9),
				Fee:        *utils.RequireIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwQAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6baM=",
			wantErr: false,
		},
		{
			name: "Config Defaults Contract Create should be ok to post to a current node",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract Identity =\n  type state = ()\n  function main(z : int) = z"
				Code:       `cb_+QP1RgKgpVq1Ib2r2ug+UktHvfWSQ8P35HJQHM6qikqBu1DwgtT5Avv5ASqgaPJnYzj/UIg5q6R3Se/6i+h+8oTyB/s9mZhwHNU4h8WEbWFpbrjAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QHLoLnJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqhGluaXS4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7kBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uMxiAABkYgAAhJGAgIBRf7nJVvKLMUmp9Zh6pQXz2hsiCcxXOSNABiu2wb2fn5nqFGIAAMBXUIBRf2jyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFFGIAAK9XUGABGVEAW2AAGVlgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tZWWAgAZCBUmAgkANgABlZYCABkIFSYCCQA2ADgVKBUpBWW2AgAVFRWVCAkVBQgJBQkFZbUFCCkVBQYgAAjFaFMi4xLjBJtQib`,
				VMVersion:  4,
				AbiVersion: 1,
				Deposit:    Config.Client.Contracts.Deposit,
				Amount:     Config.Client.Contracts.Amount,
				Gas:        Config.Client.Contracts.Gas,
				GasPrice:   Config.Client.Contracts.GasPrice,
				Fee:        Config.Client.Fee,
				TTL:        Config.Client.TTL,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACC5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAnHQYrA==",
			},
			wantTx:  "tx_+QScKgGhAc6nreRwyfmdnU5ACICobx1Ju0RLYvEanrtku8/rc/7zAbkD+PkD9UYCoKVatSG9q9roPlJLR731kkPD9+RyUBzOqopKgbtQ8ILU+QL7+QEqoGjyZ2M4/1CIOaukd0nv+ovofvKE8gf7PZmYcBzVOIfFhG1haW64wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPkBy6C5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6oRpbml0uGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////////////////////////////////+5AUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////7jMYgAAZGIAAISRgICAUX+5yVbyizFJqfWYeqUF89obIgnMVzkjQAYrtsG9n5+Z6hRiAADAV1CAUX9o8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxRRiAACvV1BgARlRAFtgABlZYCABkIFSYCCQA2ADgVKQWWAAUVlSYABSYADzW2AAgFJgAPNbWVlgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgA4FSgVKQVltgIAFRUVlQgJFQUICQUJBWW1BQgpFQUGIAAIxWhTIuMS4wgwQAAYa15iD0gACCAfQAAIQ7msoAhDuaygC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAguclW8osxSan1mHqlBfPaGyIJzFc5I0AGK7bBvZ+fmeoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6baM=",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewContractCreateTx(
				tt.fields.OwnerID,
				tt.fields.AccountNonce,
				tt.fields.Code,
				tt.fields.VMVersion,
				tt.fields.AbiVersion,
				tt.fields.Deposit,
				tt.fields.Amount,
				tt.fields.Gas,
				tt.fields.GasPrice,
				tt.fields.Fee,
				tt.fields.TTL,
				tt.fields.CallData,
			)
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := SerializeTx(&tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCreateTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("ContractCreateTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestContractCreateTx_FeeEstimate(t *testing.T) {
	type fields struct {
		OwnerID      string
		AccountNonce uint64
		Code         string
		VMVersion    uint16
		AbiVersion   uint16
		Deposit      big.Int
		Amount       big.Int
		Gas          big.Int
		GasPrice     big.Int
		Fee          big.Int
		TTL          uint64
		CallData     string
	}
	testCases := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Create",
			fields: fields{
				OwnerID:      "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: 1,
				// encoded "contract SimpleStorage =\n  record state = { data : int }\n  function init(value : int) : state = { data = value }\n  function get() : int = state.data\n  function set(value : int) = put(state{data = value})"
				Code:       `cb_+QYYRgKgf6Gy7VnRXycsYSiFGAUHhMs+Oeg+RJvmPzCSAnxk8LT5BKX5AUmgOoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugeDc2V0uMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////////////////////////////////////////jJoEnsSQdsAgNxJqQzA+rc5DsuLDKUV7ETxQp+ItyJgJS3g2dldLhgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QKLoOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEhGluaXS4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQFEYgAAj2IAAMKRgICAUX9J7EkHbAIDcSakMwPq3OQ7LiwylFexE8UKfiLciYCUtxRiAAE5V1CAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAA0VdQgFF/OoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugcUYgABG1dQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tgAFFRkFZbYCABUVGQUIOSUICRUFCAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUVlQgJFQUGAAUYFZkIFSkFBgAFJZkFCQVltQUFlQUGIAAMpWhTIuMS4w4SWVhA==`,
				VMVersion:  3,
				AbiVersion: 1,
				Deposit:    *new(big.Int),
				Amount:     *new(big.Int),
				Gas:        *utils.NewIntFromUint64(1e9),
				GasPrice:   *utils.NewIntFromUint64(1e9),
				Fee:        *utils.RequireIntFromString("200000000000000"),
				TTL:        500,
				CallData:   "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBo8mdjOP9QiDmrpHdJ7/qL6H7yhPIH+z2ZmHAc1TiHxQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7dbVl",
			},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCreateTx{
				OwnerID:      tt.fields.OwnerID,
				AccountNonce: tt.fields.AccountNonce,
				Code:         tt.fields.Code,
				VMVersion:    tt.fields.VMVersion,
				AbiVersion:   tt.fields.AbiVersion,
				Deposit:      tt.fields.Deposit,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
				CallData:     tt.fields.CallData,
			}

			zero := new(big.Int)
			got, err := tx.FeeEstimate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCreateTx.FeeEstimate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(zero) != 1 {
				t.Errorf("ContractCreateTx.FeeEstimate() was not larger than 0: %v", got)
			}
		})
	}
}

func TestContractCallTx_EncodeRLP(t *testing.T) {
	type fields struct {
		CallerID     string
		AccountNonce uint64
		ContractID   string
		Amount       big.Int
		Gas          big.Int
		GasPrice     big.Int
		AbiVersion   uint16
		CallData     string
		Fee          big.Int
		TTL          uint64
	}
	testCases := []struct {
		name    string
		fields  fields
		wantTx  string
		wantErr bool
	}{
		{
			name: "Fixed Value Contract Call Minerva (AbiVersion 3)",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(1),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *utils.NewIntFromUint64(10),
				Gas:          *utils.NewIntFromUint64(10),
				GasPrice:     *utils.NewIntFromUint64(10),
				AbiVersion:   3,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.RequireIntFromString("200000000000000"),
				TTL:          500,
			},
			wantTx:  "tx_+NcrAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMBoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwOGteYg9IAAggH0CgoKuIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKk9Ku98=",
			wantErr: false,
		},
		{
			name: "Fixed Value Contract Call Fortuna (AbiVersion 4)",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *new(big.Int),
				Gas:          *utils.NewIntFromUint64(1e9),
				GasPrice:     *utils.NewIntFromUint64(1e9),
				AbiVersion:   4,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.RequireIntFromString("200000000000000"),
				TTL:          500,
			},
			wantTx:  "tx_+N8rAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMCoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwSGteYg9IAAggH0AIQ7msoAhDuaygC4gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqty+KkQ==",
			wantErr: false,
		},
		{
			name: "Config Defaults Contract Call",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       Config.Client.Contracts.Amount,
				Gas:          Config.Client.Contracts.Gas,
				GasPrice:     Config.Client.Contracts.GasPrice,
				AbiVersion:   Config.Client.Contracts.ABIVersion,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          Config.Client.Fee,
				TTL:          Config.Client.TTL,
			},
			wantTx:  "tx_+N8rAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMCoQXv7ERRuvBfaps6b3yVUqmUUEmGvanaJbGAxkh6t034wwGGteYg9IAAggH0AIQ7msoAhDuaygC4gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAg4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqwDDyuQ==",
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCallTx{
				CallerID:     tt.fields.CallerID,
				AccountNonce: tt.fields.AccountNonce,
				ContractID:   tt.fields.ContractID,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				AbiVersion:   tt.fields.AbiVersion,
				CallData:     tt.fields.CallData,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
			}
			txJSON, _ := tx.JSON()
			fmt.Println(txJSON)

			gotTx, err := SerializeTx(tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCallTx.RLP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				gotTxRawBytes, wantTxRawBytes := getRLPSerialized(gotTx, tt.wantTx)
				t.Errorf("ContractCallTx.RLP() = \n%v\n%v, want \n%v\n%v", gotTx, gotTxRawBytes, tt.wantTx, wantTxRawBytes)
			}
		})
	}
}

func TestContractCallTx_FeeEstimate(t *testing.T) {
	type fields struct {
		CallerID     string
		AccountNonce uint64
		ContractID   string
		Amount       big.Int
		Gas          big.Int
		GasPrice     big.Int
		AbiVersion   uint16
		VMVersion    uint16
		CallData     string
		Fee          big.Int
		TTL          uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Fortuna (AbiVersion 4), with fixed params for fee estimation",
			fields: fields{
				CallerID:     "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
				AccountNonce: uint64(2),
				ContractID:   "ct_2pfWWzeRzWSdm68HXZJn61KhxdsBA46wzYgvo1swkdJZij1rKm",
				Amount:       *utils.NewIntFromUint64(0),
				Gas:          *utils.NewIntFromUint64(1e5),
				GasPrice:     *utils.NewIntFromUint64(1e9),
				AbiVersion:   4,
				CallData:     "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li",
				Fee:          *utils.NewIntFromUint64(2e9),
				TTL:          0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ContractCallTx{
				CallerID:     tt.fields.CallerID,
				AccountNonce: tt.fields.AccountNonce,
				ContractID:   tt.fields.ContractID,
				Amount:       tt.fields.Amount,
				Gas:          tt.fields.Gas,
				GasPrice:     tt.fields.GasPrice,
				AbiVersion:   tt.fields.AbiVersion,
				CallData:     tt.fields.CallData,
				Fee:          tt.fields.Fee,
				TTL:          tt.fields.TTL,
			}
			zero := new(big.Int)
			got, err := tx.FeeEstimate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContractCallTx.FeeEstimate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(zero) != 1 {
				t.Errorf("ContractCallTx.FeeEstimate() was not larger than 0: %v", got)
			}
		})
	}
}
func Test_encodeVMABI(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeVMABI(tt.args.VMVersion, tt.args.ABIVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeVMABI() = %v, want %v", got, tt.want)
			}
		})
	}
}