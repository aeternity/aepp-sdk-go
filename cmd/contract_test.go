package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

const contractSimpleStorage = "contract SimpleStorage =\n  record state = { data : int }\n  function init(value : int) : state = { data = value }\n  function get() : int = state.data\n  stateful function set(value : int) = put(state{data = value})"
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
