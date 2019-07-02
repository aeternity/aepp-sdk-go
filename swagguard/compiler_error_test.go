package swagguard_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/swagguard/compiler/client/operations"
	"github.com/aeternity/aepp-sdk-go/swagguard/compiler/models"
)

func TestCompilerErrorModelDereferencing(t *testing.T) {
	reason := "A Very Good Reason"
	compileContractBadRequest := operations.NewCompileContractBadRequest()
	err := models.Error{Reason: &reason}
	compileContractBadRequest.Payload = &err
	printedError := fmt.Sprintf("BadRequest %s", compileContractBadRequest)
	if !strings.Contains(printedError, reason) {
		t.Errorf("Expected to find %s when printing out the models.Error: got %s instead", reason, printedError)
	}
}
