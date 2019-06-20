package integrationtest

import (
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestCompiler(t *testing.T) {
	c := aeternity.NewCompiler("http://localhost:3080", false)
	t.Run("GetAPIVersion", func(t *testing.T) {
		r, err := c.GetAPIVersion()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r)
	})

}
