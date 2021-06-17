// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aeternity/aepp-sdk-go/v9/swagguard/compiler/models"
)

// GetFateAssemblerCodeReader is a Reader for the GetFateAssemblerCode structure.
type GetFateAssemblerCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFateAssemblerCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFateAssemblerCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFateAssemblerCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFateAssemblerCodeOK creates a GetFateAssemblerCodeOK with default headers values
func NewGetFateAssemblerCodeOK() *GetFateAssemblerCodeOK {
	return &GetFateAssemblerCodeOK{}
}

/* GetFateAssemblerCodeOK describes a response with status code 200, with default header values.

The FATE assembler
*/
type GetFateAssemblerCodeOK struct {
	Payload *models.FateAssembler
}

func (o *GetFateAssemblerCodeOK) Error() string {
	return fmt.Sprintf("[POST /fate-assembler][%d] getFateAssemblerCodeOK  %+v", 200, o.Payload)
}
func (o *GetFateAssemblerCodeOK) GetPayload() *models.FateAssembler {
	return o.Payload
}

func (o *GetFateAssemblerCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FateAssembler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFateAssemblerCodeBadRequest creates a GetFateAssemblerCodeBadRequest with default headers values
func NewGetFateAssemblerCodeBadRequest() *GetFateAssemblerCodeBadRequest {
	return &GetFateAssemblerCodeBadRequest{}
}

/* GetFateAssemblerCodeBadRequest describes a response with status code 400, with default header values.

Invalid data
*/
type GetFateAssemblerCodeBadRequest struct {
	Payload *models.Error
}

func (o *GetFateAssemblerCodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /fate-assembler][%d] getFateAssemblerCodeBadRequest  %+v", 400, o.Payload)
}
func (o *GetFateAssemblerCodeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFateAssemblerCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
