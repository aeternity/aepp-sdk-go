// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

// GetCurrentGenerationReader is a Reader for the GetCurrentGeneration structure.
type GetCurrentGenerationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentGenerationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentGenerationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCurrentGenerationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentGenerationOK creates a GetCurrentGenerationOK with default headers values
func NewGetCurrentGenerationOK() *GetCurrentGenerationOK {
	return &GetCurrentGenerationOK{}
}

/* GetCurrentGenerationOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetCurrentGenerationOK struct {
	Payload *models.Generation
}

func (o *GetCurrentGenerationOK) Error() string {
	return fmt.Sprintf("[GET /generations/current][%d] getCurrentGenerationOK  %+v", 200, o.Payload)
}
func (o *GetCurrentGenerationOK) GetPayload() *models.Generation {
	return o.Payload
}

func (o *GetCurrentGenerationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Generation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentGenerationNotFound creates a GetCurrentGenerationNotFound with default headers values
func NewGetCurrentGenerationNotFound() *GetCurrentGenerationNotFound {
	return &GetCurrentGenerationNotFound{}
}

/* GetCurrentGenerationNotFound describes a response with status code 404, with default header values.

Generation not found
*/
type GetCurrentGenerationNotFound struct {
	Payload *models.Error
}

func (o *GetCurrentGenerationNotFound) Error() string {
	return fmt.Sprintf("[GET /generations/current][%d] getCurrentGenerationNotFound  %+v", 404, o.Payload)
}
func (o *GetCurrentGenerationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentGenerationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
