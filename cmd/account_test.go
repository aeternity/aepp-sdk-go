package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestAccountCreate(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}

	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}
}

func TestAccountAddress(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}

	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}

	err = addressFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}
}

func TestAccountSave(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}
	privateKey := "025528252ec5db7d77cd57e14ae7819b9205c84abe5eef8353f88330467048f458019537ef2e809fefe1f2513cda8c8aacc74fb30f8c1f8b32d99a16b7f539b8"
	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = saveFunc(&emptyCmd, []string{"test.json", privateKey})
	if err != nil {
		t.Error(err)
	}
}

func TestAccountBalance(t *testing.T) {
	setCLIConfig()
	password = "password"
	emptyCmd := cobra.Command{}
	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}
	err = balanceFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}
}

func TestAccountSign(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}
	dir, err := ioutil.TempDir("", "aecli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}

	err = signFunc(&emptyCmd, []string{"test.json", "tx_+E8MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIIKC5wGAkBVBMQ=="})
	if err != nil {
		t.Error(err)
	}
}
