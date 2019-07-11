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

func writeTestContractFile(source string) (tempdir string, path string, err error) {
	tempdir, err = ioutil.TempDir("", "aepp-sdk-go")
	if err != nil {
		return "", "", err
	}
	path = filepath.Join(tempdir, "testcontract.aes")
	err = ioutil.WriteFile(path, []byte(source), 0666)
	if err != nil {
		return "", "", err
	}

	return
}

func Test_compileFunc(t *testing.T) {
	emptyCmd := cobra.Command{}

	type args struct {
		cmd    *cobra.Command
		source string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Simple storage should compile",
			args: args{
				cmd:    &emptyCmd,
				source: contractSimpleStorage,
			},
			wantErr: false,
		},
		{
			name: "Simple storage with syntax error",
			args: args{
				cmd:    &emptyCmd,
				source: contractSimpleStorageErr,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tempdir, path, err := writeTestContractFile(tt.args.source)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tempdir)

		t.Run(tt.name, func(t *testing.T) {
			if err := compileFunc(tt.args.cmd, []string{path}); (err != nil) != tt.wantErr {
				t.Errorf("compileFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeCalldataFunc(t *testing.T) {
	emptyCmd := cobra.Command{}

	type args struct {
		cmd    *cobra.Command
		args   []string
		source string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1 function argument",
			args: args{
				cmd:    &emptyCmd,
				args:   []string{"init", "42"},
				source: contractSimpleStorage,
			},
		},
	}
	for _, tt := range tests {
		tempdir, path, err := writeTestContractFile(tt.args.source)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tempdir)

		t.Run(tt.name, func(t *testing.T) {
			a := append([]string{path}, tt.args.args...)
			if err := encodeCalldataFunc(tt.args.cmd, a); (err != nil) != tt.wantErr {
				t.Errorf("encodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeCalldataFunc(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	// Write source file for Decode with source file test
	tempdir, path, err := writeTestContractFile(contractSimpleStorage)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempdir)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Decode with bytecode",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{contractSimpleStorageBytecode, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
		},
		{
			name: "Decode with source file",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{path, contractSimpleStorageInitCalldata},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeCalldataFunc(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("decodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
