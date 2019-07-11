package cmd

import (
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/swagguard/compiler/models"
)

const contractSimpleStorage = "contract SimpleStorage =\n  record state = { data : int }\n  function init(value : int) : state = { data = value }\n  function get() : int = state.data\n  stateful function set(value : int) = put(state{data = value})"
const contractSimpleStorageBytecode = "cb_+QYYRgKg+HOI9x+n5+MOEpnQ/zO+GoibqhQxGO4bgnvASx0vzB75BKX5AUmgOoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugeDc2V0uMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////////////////////////////////////////jJoEnsSQdsAgNxJqQzA+rc5DsuLDKUV7ETxQp+ItyJgJS3g2dldLhgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QKLoOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEhGluaXS4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQFEYgAAj2IAAMKRgICAUX9J7EkHbAIDcSakMwPq3OQ7LiwylFexE8UKfiLciYCUtxRiAAE5V1CAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAA0VdQgFF/OoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugcUYgABG1dQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tgAFFRkFZbYCABUVGQUIOSUICRUFCAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUVlQgJFQUGAAUYFZkIFSkFBgAFJZkFCQVltQUFlQUGIAAMpWhTMuMS4wHchc+w=="
const contractSimpleStorageInitCalldata = "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li"
const contractSimpleStorageErr = "contract SimpleStorage =\n  record state = { data : int }\n  function init(value : int) : state = { data = value }\n  function get() : int = state.data\n  function set(value : int) = put(state{data = value})"

type mockCompileContracter struct{}

func (m *mockCompileContracter) CompileContract(source string) (bytecode string, err error) {
	return "cb_+QYYRgKg+HOI9x+n5+MOEpnQ/zO+GoibqhQxGO4bgnvASx0vzB75BKX5AUmgOoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugeDc2V0uMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////////////////////////////////////////jJoEnsSQdsAgNxJqQzA+rc5DsuLDKUV7ETxQp+ItyJgJS3g2dldLhgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QKLoOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEhGluaXS4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQFEYgAAj2IAAMKRgICAUX9J7EkHbAIDcSakMwPq3OQ7LiwylFexE8UKfiLciYCUtxRiAAE5V1CAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAA0VdQgFF/OoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugcUYgABG1dQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tgAFFRkFZbYCABUVGQUIOSUICRUFCAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUVlQgJFQUGAAUYFZkIFSkFBgAFJZkFCQVltQUFlQUGIAAMpWhTMuMS4wHchc+w==", nil
}
func Test_compileFunc(t *testing.T) {
	type args struct {
		conn   aeternity.CompileContracter
		source string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Simple storage, mocked compiler",
			args: args{
				conn:   &mockCompileContracter{},
				source: contractSimpleStorage,
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Simple storage, online compiler: should compile",
			args: args{
				conn:   newCompiler(),
				source: contractSimpleStorage,
			},
			wantErr: false,
			online:  true,
		},
		{
			name: "Simple storage with syntax error, online compiler: shouldn't compile",
			args: args{
				conn:   newCompiler(),
				source: contractSimpleStorageErr,
			},
			wantErr: true,
			online:  true,
		},
	}
	for _, tt := range tests {
		if !online && tt.online {
			t.Skip("Skipping online test")
		}
		tempdir, path := writeTestContractFile(t, tt.args.source)
		defer os.RemoveAll(tempdir)

		t.Run(tt.name, func(t *testing.T) {
			if err := compileFunc(tt.args.conn, []string{path}); (err != nil) != tt.wantErr {
				t.Errorf("compileFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockEncodeCalldataer struct{}

func (m *mockEncodeCalldataer) EncodeCalldata(source string, function string, args []string) (bytecode string, err error) {
	return "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li", nil
}
func Test_encodeCalldataFunc(t *testing.T) {
	type args struct {
		conn   aeternity.EncodeCalldataer
		args   []string
		source string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "1 function argument",
			args: args{
				conn:   &mockEncodeCalldataer{},
				args:   []string{"init", "42"},
				source: contractSimpleStorage,
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "1 function argument (online)",
			args: args{
				conn:   newCompiler(),
				args:   []string{"init", "42"},
				source: contractSimpleStorage,
			},
			wantErr: false,
			online:  true,
		},
	}
	for _, tt := range tests {
		if !online && tt.online {
			t.Skip("Skipping online test")
		}
		tempdir, path := writeTestContractFile(t, tt.args.source)
		defer os.RemoveAll(tempdir)

		t.Run(tt.name, func(t *testing.T) {
			a := append([]string{path}, tt.args.args...)
			if err := encodeCalldataFunc(tt.args.conn, a); (err != nil) != tt.wantErr {
				t.Errorf("encodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockdecodeCalldataer struct {
	decodedCalldata string
}

func (m *mockdecodeCalldataer) DecodeCalldataSource(source string, callData string) (decodedCallData *models.DecodedCalldata, err error) {
	decodedCallData = &models.DecodedCalldata{}
	decodedCallData.UnmarshalBinary([]byte(m.decodedCalldata))
	return decodedCallData, nil
}
func (m *mockdecodeCalldataer) DecodeCalldataBytecode(bytecode string, calldata string) (decodedCallData *models.DecodedCalldata, err error) {
	decodedCallData = &models.DecodedCalldata{}
	decodedCallData.UnmarshalBinary([]byte(m.decodedCalldata))
	return decodedCallData, nil
}

func Test_decodeCalldataFunc(t *testing.T) {
	type args struct {
		conn decodeCalldataer
		args []string
	}
	// Write source file for Decode with source file test
	tempdir, path := writeTestContractFile(t, contractSimpleStorage)
	defer os.RemoveAll(tempdir)

	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Decode with bytecode",
			args: args{
				conn: &mockdecodeCalldataer{decodedCalldata: `{"arguments":[{"type":"word","value":42}],"function":"init"}`},
				args: []string{contractSimpleStorageBytecode, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Decode with source file",
			args: args{
				conn: &mockdecodeCalldataer{decodedCalldata: `{"arguments":[{"type":"word","value":42}],"function":"init"}`},
				args: []string{path, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Decode with bytecode (online)",
			args: args{
				conn: newCompiler(),
				args: []string{contractSimpleStorageBytecode, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
			online:  true,
		},
		{
			name: "Decode with source file (online)",
			args: args{
				conn: newCompiler(),
				args: []string{path, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
			online:  true,
		},
	}

	for _, tt := range tests {
		if !online && tt.online {
			t.Skip("Skipping online test")
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeCalldataFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("decodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
