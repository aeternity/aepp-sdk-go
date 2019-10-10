package cmd

import (
	"encoding"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/config"
)

func setPrivateNetParams() {
	config.Node.URL = "http://localhost:3013"
	config.Node.NetworkID = "ae_docker"
}

// dumpV serializes and prints out any swagger model in JSON. Useful when
// writing test mocks
func dumpV(v encoding.BinaryMarshaler) error {
	s, err := v.MarshalBinary()
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	return err
}

func writeTestContractFile(t *testing.T, source string) (tempdir string, path string) {
	tempdir, err := ioutil.TempDir("", "aepp-sdk-go")
	if err != nil {
		t.Fatal(err)
	}
	path = filepath.Join(tempdir, "testcontract.aes")
	err = ioutil.WriteFile(path, []byte(source), 0666)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func testTempdir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Fatal(err)
	}
	return dir, func() { os.RemoveAll(dir) }
}
