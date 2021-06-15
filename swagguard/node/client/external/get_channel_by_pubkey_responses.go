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

// GetChannelByPubkeyReader is a Reader for the GetChannelByPubkey structure.
type GetChannelByPubkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChannelByPubkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChannelByPubkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetChannelByPubkeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetChannelByPubkeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChannelByPubkeyOK creates a GetChannelByPubkeyOK with default headers values
func NewGetChannelByPubkeyOK() *GetChannelByPubkeyOK {
	return &GetChannelByPubkeyOK{}
}

/* GetChannelByPubkeyOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetChannelByPubkeyOK struct {
	Payload *models.Channel
}

func (o *GetChannelByPubkeyOK) Error() string {
	return fmt.Sprintf("[GET /channels/{pubkey}][%d] getChannelByPubkeyOK  %+v", 200, o.Payload)
}
func (o *GetChannelByPubkeyOK) GetPayload() *models.Channel {
	return o.Payload
}

func (o *GetChannelByPubkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Channel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelByPubkeyBadRequest creates a GetChannelByPubkeyBadRequest with default headers values
func NewGetChannelByPubkeyBadRequest() *GetChannelByPubkeyBadRequest {
	return &GetChannelByPubkeyBadRequest{}
}

/* GetChannelByPubkeyBadRequest describes a response with status code 400, with default header values.

Invalid public key
*/
type GetChannelByPubkeyBadRequest struct {
	Payload *models.Error
}

func (o *GetChannelByPubkeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /channels/{pubkey}][%d] getChannelByPubkeyBadRequest  %+v", 400, o.Payload)
}
func (o *GetChannelByPubkeyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChannelByPubkeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelByPubkeyNotFound creates a GetChannelByPubkeyNotFound with default headers values
func NewGetChannelByPubkeyNotFound() *GetChannelByPubkeyNotFound {
	return &GetChannelByPubkeyNotFound{}
}

/* GetChannelByPubkeyNotFound describes a response with status code 404, with default header values.

Channel not found
*/
type GetChannelByPubkeyNotFound struct {
	Payload *models.Error
}

func (o *GetChannelByPubkeyNotFound) Error() string {
	return fmt.Sprintf("[GET /channels/{pubkey}][%d] getChannelByPubkeyNotFound  %+v", 404, o.Payload)
}
func (o *GetChannelByPubkeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChannelByPubkeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
