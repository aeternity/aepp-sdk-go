// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v5/swagguard/compiler/models"
)

// EncodeCalldataReader is a Reader for the EncodeCalldata structure.
type EncodeCalldataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EncodeCalldataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEncodeCalldataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEncodeCalldataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEncodeCalldataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEncodeCalldataOK creates a EncodeCalldataOK with default headers values
func NewEncodeCalldataOK() *EncodeCalldataOK {
	return &EncodeCalldataOK{}
}

/*EncodeCalldataOK handles this case with default header values.

Binary encoded calldata
*/
type EncodeCalldataOK struct {
	Payload *models.Calldata
}

func (o *EncodeCalldataOK) Error() string {
	return fmt.Sprintf("[POST /encode-calldata][%d] encodeCalldataOK  %+v", 200, o.Payload)
}

func (o *EncodeCalldataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Calldata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEncodeCalldataBadRequest creates a EncodeCalldataBadRequest with default headers values
func NewEncodeCalldataBadRequest() *EncodeCalldataBadRequest {
	return &EncodeCalldataBadRequest{}
}

/*EncodeCalldataBadRequest handles this case with default header values.

Invalid data
*/
type EncodeCalldataBadRequest struct {
	Payload *models.Error
}

func (o *EncodeCalldataBadRequest) Error() string {
	return fmt.Sprintf("[POST /encode-calldata][%d] encodeCalldataBadRequest  %+v", 400, o.Payload)
}

func (o *EncodeCalldataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEncodeCalldataForbidden creates a EncodeCalldataForbidden with default headers values
func NewEncodeCalldataForbidden() *EncodeCalldataForbidden {
	return &EncodeCalldataForbidden{}
}

/*EncodeCalldataForbidden handles this case with default header values.

Invalid contract
*/
type EncodeCalldataForbidden struct {
	Payload models.CompilerErrors
}

func (o *EncodeCalldataForbidden) Error() string {
	return fmt.Sprintf("[POST /encode-calldata][%d] encodeCalldataForbidden  %+v", 403, o.Payload)
}

func (o *EncodeCalldataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
