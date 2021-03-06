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

// VersionReader is a Reader for the Version structure.
type VersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVersionOK creates a VersionOK with default headers values
func NewVersionOK() *VersionOK {
	return &VersionOK{}
}

/* VersionOK describes a response with status code 200, with default header values.

Sophia compiler version
*/
type VersionOK struct {
	Payload *models.CompilerVersion
}

func (o *VersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] versionOK  %+v", 200, o.Payload)
}
func (o *VersionOK) GetPayload() *models.CompilerVersion {
	return o.Payload
}

func (o *VersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompilerVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionInternalServerError creates a VersionInternalServerError with default headers values
func NewVersionInternalServerError() *VersionInternalServerError {
	return &VersionInternalServerError{}
}

/* VersionInternalServerError describes a response with status code 500, with default header values.

Error
*/
type VersionInternalServerError struct {
	Payload *models.Error
}

func (o *VersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /version][%d] versionInternalServerError  %+v", 500, o.Payload)
}
func (o *VersionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *VersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
