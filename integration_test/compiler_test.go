package integrationtest

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func TestCompiler(t *testing.T) {
	c := aeternity.NewCompiler("http://localhost:3080", false)
	t.Run("APIVersion", func(t *testing.T) {
		_, err := c.APIVersion()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("CompileContract", func(t *testing.T) {
		_, err := c.CompileContract("contract Identity =\n  type state = ()\n  entrypoint main(z : int) = z")
		if err != nil {
			t.Error(err)
		}
	})

}
