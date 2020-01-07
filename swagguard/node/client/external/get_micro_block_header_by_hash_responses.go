// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

// GetMicroBlockHeaderByHashReader is a Reader for the GetMicroBlockHeaderByHash structure.
type GetMicroBlockHeaderByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicroBlockHeaderByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicroBlockHeaderByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMicroBlockHeaderByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMicroBlockHeaderByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMicroBlockHeaderByHashOK creates a GetMicroBlockHeaderByHashOK with default headers values
func NewGetMicroBlockHeaderByHashOK() *GetMicroBlockHeaderByHashOK {
	return &GetMicroBlockHeaderByHashOK{}
}

/*GetMicroBlockHeaderByHashOK handles this case with default header values.

Successful operation
*/
type GetMicroBlockHeaderByHashOK struct {
	Payload *models.MicroBlockHeader
}

func (o *GetMicroBlockHeaderByHashOK) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/header][%d] getMicroBlockHeaderByHashOK  %+v", 200, o.Payload)
}

func (o *GetMicroBlockHeaderByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MicroBlockHeader)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockHeaderByHashBadRequest creates a GetMicroBlockHeaderByHashBadRequest with default headers values
func NewGetMicroBlockHeaderByHashBadRequest() *GetMicroBlockHeaderByHashBadRequest {
	return &GetMicroBlockHeaderByHashBadRequest{}
}

/*GetMicroBlockHeaderByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetMicroBlockHeaderByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetMicroBlockHeaderByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/header][%d] getMicroBlockHeaderByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetMicroBlockHeaderByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockHeaderByHashNotFound creates a GetMicroBlockHeaderByHashNotFound with default headers values
func NewGetMicroBlockHeaderByHashNotFound() *GetMicroBlockHeaderByHashNotFound {
	return &GetMicroBlockHeaderByHashNotFound{}
}

/*GetMicroBlockHeaderByHashNotFound handles this case with default header values.

Block not found
*/
type GetMicroBlockHeaderByHashNotFound struct {
	Payload *models.Error
}

func (o *GetMicroBlockHeaderByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/header][%d] getMicroBlockHeaderByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetMicroBlockHeaderByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
