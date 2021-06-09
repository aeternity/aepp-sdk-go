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

// DecodeDataReader is a Reader for the DecodeData structure.
type DecodeDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecodeDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecodeDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDecodeDataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecodeDataOK creates a DecodeDataOK with default headers values
func NewDecodeDataOK() *DecodeDataOK {
	return &DecodeDataOK{}
}

/* DecodeDataOK describes a response with status code 200, with default header values.

Json encoded data
*/
type DecodeDataOK struct {
	Payload *models.SophiaJSONData
}

func (o *DecodeDataOK) Error() string {
	return fmt.Sprintf("[POST /decode-data][%d] decodeDataOK  %+v", 200, o.Payload)
}
func (o *DecodeDataOK) GetPayload() *models.SophiaJSONData {
	return o.Payload
}

func (o *DecodeDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SophiaJSONData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecodeDataBadRequest creates a DecodeDataBadRequest with default headers values
func NewDecodeDataBadRequest() *DecodeDataBadRequest {
	return &DecodeDataBadRequest{}
}

/* DecodeDataBadRequest describes a response with status code 400, with default header values.

Invalid data
*/
type DecodeDataBadRequest struct {
	Payload models.CompilerErrors
}

func (o *DecodeDataBadRequest) Error() string {
	return fmt.Sprintf("[POST /decode-data][%d] decodeDataBadRequest  %+v", 400, o.Payload)
}
func (o *DecodeDataBadRequest) GetPayload() models.CompilerErrors {
	return o.Payload
}

func (o *DecodeDataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
