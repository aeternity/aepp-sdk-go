package swagguard_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/swagguard/compiler/client/operations"
	"github.com/aeternity/aepp-sdk-go/v6/swagguard/compiler/models"
)

func TestCompilerErrorModelDereferencing(t *testing.T) {
	reason := "A Very Good Reason"
	internalServerErr := operations.APIVersionInternalServerError{}
	err := models.Error{Reason: &reason}
	internalServerErr.Payload = &err
	printedError := fmt.Sprintf("BadRequest %s", internalServerErr)
	if !strings.Contains(printedError, reason) {
		t.Errorf("Expected to find %s when printing out the models.Error: got %s instead", reason, printedError)
	}
}

func TestCompilerCompilationErrorsModelDereferencing(t *testing.T) {
	err1 := &models.CompilerError{}
	err1.UnmarshalBinary([]byte(`{"message":"Unbound variable ae_addres at line 4, column 9","pos":{"col":9,"line":4},"type":"type_error"}`))
	err2 := &models.CompilerError{}
	err2.UnmarshalBinary([]byte(`{"message":"Also I don't like your face","pos":{"col":0,"line":0},"type":"wrong_programmer_error"}`))

	compileContractErr := operations.CompileContractBadRequest{
		Payload: []*models.CompilerError{err1, err2},
	}
	printedError := fmt.Sprintf("%s", compileContractErr)
	lookForError1 := "Unbound variable ae_addres"
	lookForError2 := "Also I don't like your face"

	if !(strings.Contains(printedError, lookForError1) && strings.Contains(printedError, lookForError2)) {
		t.Errorf("Expected []*models.CompilerError to include the messages %s and %s; got %s instead", lookForError1, lookForError2, printedError)
	}
}
