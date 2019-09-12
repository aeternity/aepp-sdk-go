// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v5/swagguard/node/models"
)

// GetContractStoreReader is a Reader for the GetContractStore structure.
type GetContractStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContractStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetContractStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetContractStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractStoreOK creates a GetContractStoreOK with default headers values
func NewGetContractStoreOK() *GetContractStoreOK {
	return &GetContractStoreOK{}
}

/*GetContractStoreOK handles this case with default header values.

Contract Store
*/
type GetContractStoreOK struct {
	Payload *models.ContractStore
}

func (o *GetContractStoreOK) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/store][%d] getContractStoreOK  %+v", 200, o.Payload)
}

func (o *GetContractStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContractStore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractStoreBadRequest creates a GetContractStoreBadRequest with default headers values
func NewGetContractStoreBadRequest() *GetContractStoreBadRequest {
	return &GetContractStoreBadRequest{}
}

/*GetContractStoreBadRequest handles this case with default header values.

Invalid pubkey
*/
type GetContractStoreBadRequest struct {
	Payload *models.Error
}

func (o *GetContractStoreBadRequest) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/store][%d] getContractStoreBadRequest  %+v", 400, o.Payload)
}

func (o *GetContractStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractStoreNotFound creates a GetContractStoreNotFound with default headers values
func NewGetContractStoreNotFound() *GetContractStoreNotFound {
	return &GetContractStoreNotFound{}
}

/*GetContractStoreNotFound handles this case with default header values.

Contract not found
*/
type GetContractStoreNotFound struct {
}

func (o *GetContractStoreNotFound) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/store][%d] getContractStoreNotFound ", 404)
}

func (o *GetContractStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
