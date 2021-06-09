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

// GenerateACIReader is a Reader for the GenerateACI structure.
type GenerateACIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateACIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateACIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateACIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateACIOK creates a GenerateACIOK with default headers values
func NewGenerateACIOK() *GenerateACIOK {
	return &GenerateACIOK{}
}

/* GenerateACIOK describes a response with status code 200, with default header values.

ACI for contract
*/
type GenerateACIOK struct {
	Payload *models.ACI
}

func (o *GenerateACIOK) Error() string {
	return fmt.Sprintf("[POST /aci][%d] generateACIOK  %+v", 200, o.Payload)
}
func (o *GenerateACIOK) GetPayload() *models.ACI {
	return o.Payload
}

func (o *GenerateACIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateACIBadRequest creates a GenerateACIBadRequest with default headers values
func NewGenerateACIBadRequest() *GenerateACIBadRequest {
	return &GenerateACIBadRequest{}
}

/* GenerateACIBadRequest describes a response with status code 400, with default header values.

Compiler errors
*/
type GenerateACIBadRequest struct {
	Payload models.CompilerErrors
}

func (o *GenerateACIBadRequest) Error() string {
	return fmt.Sprintf("[POST /aci][%d] generateACIBadRequest  %+v", 400, o.Payload)
}
func (o *GenerateACIBadRequest) GetPayload() models.CompilerErrors {
	return o.Payload
}

func (o *GenerateACIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
