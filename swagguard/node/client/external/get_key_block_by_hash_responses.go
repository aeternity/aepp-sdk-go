// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v7/swagguard/node/models"
)

// GetKeyBlockByHashReader is a Reader for the GetKeyBlockByHash structure.
type GetKeyBlockByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeyBlockByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeyBlockByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetKeyBlockByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetKeyBlockByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeyBlockByHashOK creates a GetKeyBlockByHashOK with default headers values
func NewGetKeyBlockByHashOK() *GetKeyBlockByHashOK {
	return &GetKeyBlockByHashOK{}
}

/*GetKeyBlockByHashOK handles this case with default header values.

Successful operation
*/
type GetKeyBlockByHashOK struct {
	Payload *models.KeyBlock
}

func (o *GetKeyBlockByHashOK) Error() string {
	return fmt.Sprintf("[GET /key-blocks/hash/{hash}][%d] getKeyBlockByHashOK  %+v", 200, o.Payload)
}

func (o *GetKeyBlockByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyBlock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeyBlockByHashBadRequest creates a GetKeyBlockByHashBadRequest with default headers values
func NewGetKeyBlockByHashBadRequest() *GetKeyBlockByHashBadRequest {
	return &GetKeyBlockByHashBadRequest{}
}

/*GetKeyBlockByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetKeyBlockByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetKeyBlockByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /key-blocks/hash/{hash}][%d] getKeyBlockByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetKeyBlockByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeyBlockByHashNotFound creates a GetKeyBlockByHashNotFound with default headers values
func NewGetKeyBlockByHashNotFound() *GetKeyBlockByHashNotFound {
	return &GetKeyBlockByHashNotFound{}
}

/*GetKeyBlockByHashNotFound handles this case with default header values.

Block not found
*/
type GetKeyBlockByHashNotFound struct {
	Payload *models.Error
}

func (o *GetKeyBlockByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /key-blocks/hash/{hash}][%d] getKeyBlockByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetKeyBlockByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
