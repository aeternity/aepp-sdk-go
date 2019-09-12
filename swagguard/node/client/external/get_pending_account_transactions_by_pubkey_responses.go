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

// GetPendingAccountTransactionsByPubkeyReader is a Reader for the GetPendingAccountTransactionsByPubkey structure.
type GetPendingAccountTransactionsByPubkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPendingAccountTransactionsByPubkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPendingAccountTransactionsByPubkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPendingAccountTransactionsByPubkeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPendingAccountTransactionsByPubkeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPendingAccountTransactionsByPubkeyOK creates a GetPendingAccountTransactionsByPubkeyOK with default headers values
func NewGetPendingAccountTransactionsByPubkeyOK() *GetPendingAccountTransactionsByPubkeyOK {
	return &GetPendingAccountTransactionsByPubkeyOK{}
}

/*GetPendingAccountTransactionsByPubkeyOK handles this case with default header values.

Successful operation
*/
type GetPendingAccountTransactionsByPubkeyOK struct {
	Payload *models.GenericTxs
}

func (o *GetPendingAccountTransactionsByPubkeyOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/transactions/pending][%d] getPendingAccountTransactionsByPubkeyOK  %+v", 200, o.Payload)
}

func (o *GetPendingAccountTransactionsByPubkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericTxs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingAccountTransactionsByPubkeyBadRequest creates a GetPendingAccountTransactionsByPubkeyBadRequest with default headers values
func NewGetPendingAccountTransactionsByPubkeyBadRequest() *GetPendingAccountTransactionsByPubkeyBadRequest {
	return &GetPendingAccountTransactionsByPubkeyBadRequest{}
}

/*GetPendingAccountTransactionsByPubkeyBadRequest handles this case with default header values.

Invalid public key
*/
type GetPendingAccountTransactionsByPubkeyBadRequest struct {
	Payload *models.Error
}

func (o *GetPendingAccountTransactionsByPubkeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/transactions/pending][%d] getPendingAccountTransactionsByPubkeyBadRequest  %+v", 400, o.Payload)
}

func (o *GetPendingAccountTransactionsByPubkeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingAccountTransactionsByPubkeyNotFound creates a GetPendingAccountTransactionsByPubkeyNotFound with default headers values
func NewGetPendingAccountTransactionsByPubkeyNotFound() *GetPendingAccountTransactionsByPubkeyNotFound {
	return &GetPendingAccountTransactionsByPubkeyNotFound{}
}

/*GetPendingAccountTransactionsByPubkeyNotFound handles this case with default header values.

Account not found
*/
type GetPendingAccountTransactionsByPubkeyNotFound struct {
	Payload *models.Error
}

func (o *GetPendingAccountTransactionsByPubkeyNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}/transactions/pending][%d] getPendingAccountTransactionsByPubkeyNotFound  %+v", 404, o.Payload)
}

func (o *GetPendingAccountTransactionsByPubkeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
