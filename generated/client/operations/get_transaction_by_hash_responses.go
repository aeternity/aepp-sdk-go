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

// GetTransactionByHashReader is a Reader for the GetTransactionByHash structure.
type GetTransactionByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionByHashOK creates a GetTransactionByHashOK with default headers values
func NewGetTransactionByHashOK() *GetTransactionByHashOK {
	return &GetTransactionByHashOK{}
}

/*GetTransactionByHashOK handles this case with default header values.

Successful operation
*/
type GetTransactionByHashOK struct {
	Payload *models.SignedTxJSON
}

func (o *GetTransactionByHashOK) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}][%d] getTransactionByHashOK  %+v", 200, o.Payload)
}

func (o *GetTransactionByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignedTxJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionByHashBadRequest creates a GetTransactionByHashBadRequest with default headers values
func NewGetTransactionByHashBadRequest() *GetTransactionByHashBadRequest {
	return &GetTransactionByHashBadRequest{}
}

/*GetTransactionByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetTransactionByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetTransactionByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}][%d] getTransactionByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionByHashNotFound creates a GetTransactionByHashNotFound with default headers values
func NewGetTransactionByHashNotFound() *GetTransactionByHashNotFound {
	return &GetTransactionByHashNotFound{}
}

/*GetTransactionByHashNotFound handles this case with default header values.

Transaction not found
*/
type GetTransactionByHashNotFound struct {
	Payload *models.Error
}

func (o *GetTransactionByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}][%d] getTransactionByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}