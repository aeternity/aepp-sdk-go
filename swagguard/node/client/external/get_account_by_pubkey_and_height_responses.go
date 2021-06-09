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

// GetAccountByPubkeyAndHeightReader is a Reader for the GetAccountByPubkeyAndHeight structure.
type GetAccountByPubkeyAndHeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountByPubkeyAndHeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountByPubkeyAndHeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountByPubkeyAndHeightBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountByPubkeyAndHeightNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountByPubkeyAndHeightOK creates a GetAccountByPubkeyAndHeightOK with default headers values
func NewGetAccountByPubkeyAndHeightOK() *GetAccountByPubkeyAndHeightOK {
	return &GetAccountByPubkeyAndHeightOK{}
}

/* GetAccountByPubkeyAndHeightOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAccountByPubkeyAndHeightOK struct {
	Payload *models.Account
}

func (o *GetAccountByPubkeyAndHeightOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/height/{height}][%d] getAccountByPubkeyAndHeightOK  %+v", 200, o.Payload)
}
func (o *GetAccountByPubkeyAndHeightOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *GetAccountByPubkeyAndHeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByPubkeyAndHeightBadRequest creates a GetAccountByPubkeyAndHeightBadRequest with default headers values
func NewGetAccountByPubkeyAndHeightBadRequest() *GetAccountByPubkeyAndHeightBadRequest {
	return &GetAccountByPubkeyAndHeightBadRequest{}
}

/* GetAccountByPubkeyAndHeightBadRequest describes a response with status code 400, with default header values.

Invalid public key or invalid height
*/
type GetAccountByPubkeyAndHeightBadRequest struct {
	Payload *models.Error
}

func (o *GetAccountByPubkeyAndHeightBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/height/{height}][%d] getAccountByPubkeyAndHeightBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountByPubkeyAndHeightBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountByPubkeyAndHeightBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByPubkeyAndHeightNotFound creates a GetAccountByPubkeyAndHeightNotFound with default headers values
func NewGetAccountByPubkeyAndHeightNotFound() *GetAccountByPubkeyAndHeightNotFound {
	return &GetAccountByPubkeyAndHeightNotFound{}
}

/* GetAccountByPubkeyAndHeightNotFound describes a response with status code 404, with default header values.

Account not found or height not available
*/
type GetAccountByPubkeyAndHeightNotFound struct {
	Payload *models.Error
}

func (o *GetAccountByPubkeyAndHeightNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/height/{height}][%d] getAccountByPubkeyAndHeightNotFound  %+v", 404, o.Payload)
}
func (o *GetAccountByPubkeyAndHeightNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountByPubkeyAndHeightNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
