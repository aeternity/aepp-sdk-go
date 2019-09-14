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

// DecodeCallResultReader is a Reader for the DecodeCallResult structure.
type DecodeCallResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecodeCallResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDecodeCallResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDecodeCallResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDecodeCallResultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDecodeCallResultOK creates a DecodeCallResultOK with default headers values
func NewDecodeCallResultOK() *DecodeCallResultOK {
	return &DecodeCallResultOK{}
}

/*DecodeCallResultOK handles this case with default header values.

Json encoded data
*/
type DecodeCallResultOK struct {
	Payload models.SophiaCallResult
}

func (o *DecodeCallResultOK) Error() string {
	return fmt.Sprintf("[POST /decode-call-result][%d] decodeCallResultOK  %+v", 200, o.Payload)
}

func (o *DecodeCallResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecodeCallResultBadRequest creates a DecodeCallResultBadRequest with default headers values
func NewDecodeCallResultBadRequest() *DecodeCallResultBadRequest {
	return &DecodeCallResultBadRequest{}
}

/*DecodeCallResultBadRequest handles this case with default header values.

Invalid data
*/
type DecodeCallResultBadRequest struct {
	Payload *models.Error
}

func (o *DecodeCallResultBadRequest) Error() string {
	return fmt.Sprintf("[POST /decode-call-result][%d] decodeCallResultBadRequest  %+v", 400, o.Payload)
}

func (o *DecodeCallResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecodeCallResultForbidden creates a DecodeCallResultForbidden with default headers values
func NewDecodeCallResultForbidden() *DecodeCallResultForbidden {
	return &DecodeCallResultForbidden{}
}

/*DecodeCallResultForbidden handles this case with default header values.

Invalid data
*/
type DecodeCallResultForbidden struct {
	Payload models.CompilerErrors
}

func (o *DecodeCallResultForbidden) Error() string {
	return fmt.Sprintf("[POST /decode-call-result][%d] decodeCallResultForbidden  %+v", 403, o.Payload)
}

func (o *DecodeCallResultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
