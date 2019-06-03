package aeternity

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"time"

	"github.com/aeternity/aepp-sdk-go/generated/models"
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

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func left(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

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
					PpI(dept, n, time.Unix(0, v.Int()*int64(time.Millisecond)).Format(time.RFC3339))
				} else {
					PpI(dept, n, v)
				}
			}
		}
	}
	p(title, "", reflect.ValueOf(v), 0)
}

func getErrorReason(v interface{}) (msg string) {
	var p func(v reflect.Value) (msg string)
	p = func(v reflect.Value) (msg string) {
		switch v.Kind() {
		// If it is a pointer we need to unwrap and call once again
		case reflect.Ptr:
			if v.IsValid() {
				msg = p(v.Elem())
			}
		case reflect.Struct:
			if v.Type() == reflect.TypeOf(models.Error{}) {
				msg = fmt.Sprint(reflect.Indirect(v.FieldByName("Reason")))
				break
			}
			for i := 0; i < v.NumField(); i++ {
				msg = p(v.Field(i))
			}
		}
		return
	}
	msg = p(reflect.ValueOf(v))
	if len(msg) == 0 {
		msg = fmt.Sprint(v)
	}
	return
}

// PrintError print error
func PrintError(code string, e *models.Error) {
	Pp(code, e.Reason)
}

// PrintObject pretty print an object obtained from the api with a title
func PrintObject(title string, i interface{}) {
	if Config.Tuning.OutputFormatJSON {
		j, _ := json.MarshalIndent(i, "", "  ")
		fmt.Printf("%s\n", j)
		return
	}

	printIf(title, i)
	print("\n")

}
