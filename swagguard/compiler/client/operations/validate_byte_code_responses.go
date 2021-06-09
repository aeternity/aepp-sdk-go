// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/compiler/models"
)

// ValidateByteCodeReader is a Reader for the ValidateByteCode structure.
type ValidateByteCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateByteCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateByteCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateByteCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateByteCodeOK creates a ValidateByteCodeOK with default headers values
func NewValidateByteCodeOK() *ValidateByteCodeOK {
	return &ValidateByteCodeOK{}
}

/* ValidateByteCodeOK describes a response with status code 200, with default header values.

Validation successful
*/
type ValidateByteCodeOK struct {
}

func (o *ValidateByteCodeOK) Error() string {
	return fmt.Sprintf("[POST /validate-byte-code][%d] validateByteCodeOK ", 200)
}

func (o *ValidateByteCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateByteCodeBadRequest creates a ValidateByteCodeBadRequest with default headers values
func NewValidateByteCodeBadRequest() *ValidateByteCodeBadRequest {
	return &ValidateByteCodeBadRequest{}
}

/* ValidateByteCodeBadRequest describes a response with status code 400, with default header values.

Invalid contract
*/
type ValidateByteCodeBadRequest struct {
	Payload models.CompilerErrors
}

func (o *ValidateByteCodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /validate-byte-code][%d] validateByteCodeBadRequest  %+v", 400, o.Payload)
}
func (o *ValidateByteCodeBadRequest) GetPayload() models.CompilerErrors {
	return o.Payload
}

func (o *ValidateByteCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}