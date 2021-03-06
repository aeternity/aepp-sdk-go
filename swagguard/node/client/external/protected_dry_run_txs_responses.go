// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aeternity/aepp-sdk-go/v9/swagguard/node/models"
)

// ProtectedDryRunTxsReader is a Reader for the ProtectedDryRunTxs structure.
type ProtectedDryRunTxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProtectedDryRunTxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProtectedDryRunTxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProtectedDryRunTxsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProtectedDryRunTxsOK creates a ProtectedDryRunTxsOK with default headers values
func NewProtectedDryRunTxsOK() *ProtectedDryRunTxsOK {
	return &ProtectedDryRunTxsOK{}
}

/* ProtectedDryRunTxsOK describes a response with status code 200, with default header values.

Dry-run result
*/
type ProtectedDryRunTxsOK struct {
	Payload *models.DryRunResults
}

func (o *ProtectedDryRunTxsOK) Error() string {
	return fmt.Sprintf("[POST /dry-run][%d] protectedDryRunTxsOK  %+v", 200, o.Payload)
}
func (o *ProtectedDryRunTxsOK) GetPayload() *models.DryRunResults {
	return o.Payload
}

func (o *ProtectedDryRunTxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DryRunResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProtectedDryRunTxsForbidden creates a ProtectedDryRunTxsForbidden with default headers values
func NewProtectedDryRunTxsForbidden() *ProtectedDryRunTxsForbidden {
	return &ProtectedDryRunTxsForbidden{}
}

/* ProtectedDryRunTxsForbidden describes a response with status code 403, with default header values.

Invalid input
*/
type ProtectedDryRunTxsForbidden struct {
	Payload *models.Error
}

func (o *ProtectedDryRunTxsForbidden) Error() string {
	return fmt.Sprintf("[POST /dry-run][%d] protectedDryRunTxsForbidden  %+v", 403, o.Payload)
}
func (o *ProtectedDryRunTxsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ProtectedDryRunTxsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
