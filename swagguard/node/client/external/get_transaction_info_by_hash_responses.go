// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v6/swagguard/node/models"
)

// GetTransactionInfoByHashReader is a Reader for the GetTransactionInfoByHash structure.
type GetTransactionInfoByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionInfoByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionInfoByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionInfoByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionInfoByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionInfoByHashOK creates a GetTransactionInfoByHashOK with default headers values
func NewGetTransactionInfoByHashOK() *GetTransactionInfoByHashOK {
	return &GetTransactionInfoByHashOK{}
}

/*GetTransactionInfoByHashOK handles this case with default header values.

Successful operation
*/
type GetTransactionInfoByHashOK struct {
	Payload *models.TxInfoObject
}

func (o *GetTransactionInfoByHashOK) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}/info][%d] getTransactionInfoByHashOK  %+v", 200, o.Payload)
}

func (o *GetTransactionInfoByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TxInfoObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionInfoByHashBadRequest creates a GetTransactionInfoByHashBadRequest with default headers values
func NewGetTransactionInfoByHashBadRequest() *GetTransactionInfoByHashBadRequest {
	return &GetTransactionInfoByHashBadRequest{}
}

/*GetTransactionInfoByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetTransactionInfoByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetTransactionInfoByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}/info][%d] getTransactionInfoByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionInfoByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionInfoByHashNotFound creates a GetTransactionInfoByHashNotFound with default headers values
func NewGetTransactionInfoByHashNotFound() *GetTransactionInfoByHashNotFound {
	return &GetTransactionInfoByHashNotFound{}
}

/*GetTransactionInfoByHashNotFound handles this case with default header values.

Transaction not found
*/
type GetTransactionInfoByHashNotFound struct {
	Payload *models.Error
}

func (o *GetTransactionInfoByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /transactions/{hash}/info][%d] getTransactionInfoByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionInfoByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
