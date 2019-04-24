package generated_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
)

func TestErrorModelDereferencing(t *testing.T) {
	reason := "A Very Good Reason"
	postTransactionBadRequest := external.NewPostTransactionBadRequest()
	err := models.Error{Reason: &reason}
	postTransactionBadRequest.Payload = &err
	printedError := fmt.Sprintf("BadRequest %s", postTransactionBadRequest)
	if !strings.Contains(printedError, reason) {
		t.Errorf("Expected to find %s when printing out the models.Error: got %s instead", reason, printedError)
	}
}
