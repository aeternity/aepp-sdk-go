package cmd

import (
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/naet"
)

func Test_compileFunc(t *testing.T) {
	type args struct {
		conn   naet.CompileContracter
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

func Test_encodeCalldataFunc(t *testing.T) {
	type args struct {
		conn   naet.EncodeCalldataer
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

func Test_decodeCalldataBytecodeFunc(t *testing.T) {
	type args struct {
		conn decodeCalldataer
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Decode with bytecode",
			args: args{
				conn: &mockdecodeCalldataer{decodedCalldata: `{"arguments":[{"type":"init","value":42}],"function":"init"}`},
				args: []string{contractSimpleStorageBytecode, contractSimpleStorageInit42},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Decode with bytecode (online)",
			args: args{
				conn: newCompiler(),
				args: []string{contractSimpleStorageBytecode, contractSimpleStorageInit42},
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
			if err := decodeCalldataBytecodeFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("decodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_decodeCalldataSourceFunc(t *testing.T) {
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
			name: "Decode with source file",
			args: args{
				conn: &mockdecodeCalldataer{decodedCalldata: `{"arguments":[{"type":"int","value":42}],"function":"init"}`},
				args: []string{path, "init", contractSimpleStorageInit42},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Decode with source file (online)",
			args: args{
				conn: newCompiler(),
				args: []string{path, "init", contractSimpleStorageInit42},
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
			if err := decodeCalldataSourceFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("decodeCalldataFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_generateAciFunc(t *testing.T) {
	// Write source file for Decode with source file test
	tempdir, path := writeTestContractFile(t, contractSimpleStorage)
	defer os.RemoveAll(tempdir)
	type args struct {
		conn naet.GenerateACIer
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate ACI from SimpleStorage",
			args: args{
				conn: &mockGenerateACIer{
					aci: `{"encoded_aci":{"contract":{"functions":[{"arguments":[{"name":"value","type":"int"}],"name":"init","returns":"SimpleStorage.state","stateful":false}],"name":"SimpleStorage","state":{"record":[{"name":"data","type":"int"}]},"type_defs":[]}},"interface":"contract SimpleStorage =\n  record state = {data : int}\n  entrypoint init : (int) =\u003e SimpleStorage.state\n"}`,
				},
				args: []string{path},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateAciFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("generateAciFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
