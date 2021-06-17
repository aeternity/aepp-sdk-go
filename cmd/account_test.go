package cmd

import (
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v9/naet"
	"github.com/spf13/cobra"
)

func TestAccountCreate(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}

	dir, closer := testTempdir(t)
	defer closer()
	os.Chdir(dir)

	err := createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}
}

func TestAccountAddress(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}

	dir, closer := testTempdir(t)
	defer closer()
	os.Chdir(dir)

	err := createFunc(&emptyCmd, []string{"test.json"})
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
	dir, closer := testTempdir(t)
	defer closer()
	os.Chdir(dir)

	err := saveFunc(&emptyCmd, []string{"test.json", privateKey})
	if err != nil {
		t.Error(err)
	}
}
func TestAccountSign(t *testing.T) {
	password = "password"
	emptyCmd := cobra.Command{}
	dir, closer := testTempdir(t)
	defer closer()
	os.Chdir(dir)

	err := createFunc(&emptyCmd, []string{"test.json"})
	if err != nil {
		t.Error(err)
	}

	err = signFunc(&emptyCmd, []string{"test.json", "tx_+E8MAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vOhAR8To7CL8AFABmKmi2nYdfeAPOxMCGR/btXYTHiXvVCjCoJOIIKC5wGAkBVBMQ=="})
	if err != nil {
		t.Error(err)
	}
}

func Test_balanceFunc(t *testing.T) {
	setPrivateNetParams()
	type args struct {
		conn              naet.GetAccounter
		args              []string
		accountPrivateKey string
		password          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Balance exists",
			args: args{
				conn:              &mockGetAccounter{account: `{"balance":1600000000000000077131306000000000000000,"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","kind":"basic","nonce":0}`},
				args:              []string{"test.json"},
				accountPrivateKey: "e6a91d633c77cf5771329d3354b3bcef1bc5e032c43d70b6d35af923ce1eb74dcea7ade470c9f99d9d4e400880a86f1d49bb444b62f11a9ebb64bbcfeb73fef3",
				password:          "password",
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Online: Balance exists",
			args: args{
				conn:              newAeNode(),
				args:              []string{"test.json"},
				accountPrivateKey: "e6a91d633c77cf5771329d3354b3bcef1bc5e032c43d70b6d35af923ce1eb74dcea7ade470c9f99d9d4e400880a86f1d49bb444b62f11a9ebb64bbcfeb73fef3",
				password:          "password",
			},
			wantErr: false,
			online:  true,
		},
		{
			name: "Online: Balance does not exist",
			args: args{
				conn:              newAeNode(),
				args:              []string{"test.json"},
				accountPrivateKey: "f8f7a742bf75497bc5541b3bb487d1c3b822d73fb182229baf123ecafe926fe5d5ab4dd50536b3889f9b467b6076a9cabcd6d523756f15297967a40e6af3142b",
				password:          "password",
			},
			wantErr: true,
			online:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !online && tt.online {
				t.Skip("Skipping online test")
			}

			// Create the account keystore
			dir, closer := testTempdir(t)
			defer closer()
			os.Chdir(dir)
			password = tt.args.password
			err := saveFunc(&cobra.Command{}, []string{"test.json", tt.args.accountPrivateKey})
			if err != nil {
				t.Fatal(err)
			}
			// Query the balance using the account keystore
			if err := balanceFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("balanceFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
