package cmd

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/utils"
)

func Test_printIf_BigIntBalancePrinted(t *testing.T) {
	type args struct {
		title string
		v     *big.Int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test that printIf() recognizes the big.Int special case",
			args: args{
				title: "Title",
				v:     utils.NewIntFromUint64(1377),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printIf(tt.args.title, tt.args.v)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			outs := string(out)
			os.Stdout = rescueStdout
			fmt.Println("DEBUG", outs)

			expected := tt.args.v.Text(10)
			if !strings.Contains(outs, expected) {
				t.Fatalf("The terminal pretty printer printIf did not print out the value of the BigInt, which was %s", expected)
			}
		})
	}
}
