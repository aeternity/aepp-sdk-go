// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetContractReader is a Reader for the GetContract structure.
type GetContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetContractBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetContractNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractOK creates a GetContractOK with default headers values
func NewGetContractOK() *GetContractOK {
	return &GetContractOK{}
}

/*GetContractOK handles this case with default header values.

Successful operation
*/
type GetContractOK struct {
	Payload *models.ContractObject
}

func (o *GetContractOK) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}][%d] getContractOK  %+v", 200, o.Payload)
}

func (o *GetContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContractObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractBadRequest creates a GetContractBadRequest with default headers values
func NewGetContractBadRequest() *GetContractBadRequest {
	return &GetContractBadRequest{}
}

/*GetContractBadRequest handles this case with default header values.

Invalid pubkey
*/
type GetContractBadRequest struct {
	Payload *models.Error
}

func (o *GetContractBadRequest) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}][%d] getContractBadRequest  %+v", 400, o.Payload)
}

func (o *GetContractBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractNotFound creates a GetContractNotFound with default headers values
func NewGetContractNotFound() *GetContractNotFound {
	return &GetContractNotFound{}
}

/*GetContractNotFound handles this case with default header values.

Contract not found
*/
type GetContractNotFound struct {
}

func (o *GetContractNotFound) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}][%d] getContractNotFound ", 404)
}

func (o *GetContractNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}