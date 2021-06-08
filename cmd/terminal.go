package cmd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"time"

	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
)

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

var (
	defaultIndentSize = 50
)

// Right right-pads the string with pad up to len runes
func right(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

// Pp pretty print
func Pp(data ...interface{}) {
	PpI(0, data...)
}

// PpI pretty print indent
func PpI(indentSize int, data ...interface{}) {
	for i := 0; i < len(data); i += 2 {
		rp := defaultIndentSize - indentSize

		fmt.Printf("%v%v %v\n",
			times("  ", indentSize),
			right(fmt.Sprintf("%v", data[i]), rp, "_"),
			data[i+1],
		)
	}
}

// PpT pretty print indent Title
func PpT(indentSize int, title string) {
	fmt.Printf("%v%v\n", times("  ", indentSize), title)
}

func printIf(title string, v interface{}) {
	var p func(title, n string, v reflect.Value, dept int)
	p = func(title, n string, v reflect.Value, dept int) {
		switch v.Kind() {
		// If it is a pointer we need to unwrap and call once again
		case reflect.Ptr:
			if v.IsValid() {
				p(title, n, v.Elem(), dept)
			}
		case reflect.Interface:
			p(title, n, v.Elem(), dept)
		case reflect.Struct:
			if v.Type().Name() == "Int" {
				vc := v.Interface().(big.Int)
				PpI(dept, "Balance", vc.Text(10))
			} else {
				PpT(dept, fmt.Sprintf("<%s>", v.Type().Name()))
				dept++
				for i := 0; i < v.NumField(); i++ {
					p("", v.Type().Field(i).Name, v.Field(i), dept)
				}
				dept--
				PpT(dept, fmt.Sprintf("</%s>", v.Type().Name()))
			}
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				p("", "", v.Index(i), dept)
			}
		default:
			if len(n) > 0 {
				if n == "Time" {
					t := v.Uint() * uint64(time.Millisecond)
					PpI(dept, n, time.Unix(0, int64(t)).Format(time.RFC3339))
				} else {
					PpI(dept, n, v)
				}
			}
		}
	}
	p(title, "", reflect.ValueOf(v), 0)
}

// PrintObject pretty print an object obtained from the api with a title
func PrintObject(title string, i interface{}) {
	if config.Tuning.OutputFormatJSON {
		j, _ := json.MarshalIndent(i, "", "  ")
		fmt.Printf("%s\n", j)
		return
	}

	printIf(title, i)
	print("\n")

}

type getGenerationMicroBlockTransactioner interface {
	naet.GetTransactionByHasher
	naet.GetMicroBlockTransactionsByHasher
	naet.GetGenerationByHeighter
}

// PrintGenerationByHeight utility function to print a generation by it's height
// TODO needs to be tested with error cases
func PrintGenerationByHeight(c getGenerationMicroBlockTransactioner, height uint64) {
	r, err := c.GetGenerationByHeight(height)
	if err == nil {
		PrintObject("generation", r)
		// search for transaction in the microblocks
		for _, mbh := range r.MicroBlocks {
			// get the microblok
			mbhs := fmt.Sprint(mbh)
			r, err := c.GetMicroBlockTransactionsByHash(mbhs)
			if err != nil {
				Pp("Error:", err)
			}
			// go through all the hashes
			for _, btx := range r.Transactions {
				p, err := c.GetTransactionByHash(fmt.Sprint(*btx.Hash))
				if err == nil {
					PrintObject("transaction", p)
				} else {
					fmt.Println("Error in c.GetTransactionByHash", err, btx.Hash)
					continue
				}
			}
		}
	} else {
		fmt.Println("Something went wrong in PrintGenerationByHeight")
	}
}
