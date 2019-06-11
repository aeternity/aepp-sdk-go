// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetCurrentKeyBlockHashReader is a Reader for the GetCurrentKeyBlockHash structure.
type GetCurrentKeyBlockHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentKeyBlockHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentKeyBlockHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCurrentKeyBlockHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentKeyBlockHashOK creates a GetCurrentKeyBlockHashOK with default headers values
func NewGetCurrentKeyBlockHashOK() *GetCurrentKeyBlockHashOK {
	return &GetCurrentKeyBlockHashOK{}
}

/*GetCurrentKeyBlockHashOK handles this case with default header values.

Successful operation
*/
type GetCurrentKeyBlockHashOK struct {
	Payload *GetCurrentKeyBlockHashOKBody
}

func (o *GetCurrentKeyBlockHashOK) Error() string {
	return fmt.Sprintf("[GET /key-blocks/current/hash][%d] getCurrentKeyBlockHashOK  %+v", 200, o.Payload)
}

func (o *GetCurrentKeyBlockHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCurrentKeyBlockHashOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentKeyBlockHashNotFound creates a GetCurrentKeyBlockHashNotFound with default headers values
func NewGetCurrentKeyBlockHashNotFound() *GetCurrentKeyBlockHashNotFound {
	return &GetCurrentKeyBlockHashNotFound{}
}

/*GetCurrentKeyBlockHashNotFound handles this case with default header values.

Block not found
*/
type GetCurrentKeyBlockHashNotFound struct {
	Payload *models.Error
}

func (o *GetCurrentKeyBlockHashNotFound) Error() string {
	return fmt.Sprintf("[GET /key-blocks/current/hash][%d] getCurrentKeyBlockHashNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrentKeyBlockHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCurrentKeyBlockHashOKBody get current key block hash o k body
swagger:model GetCurrentKeyBlockHashOKBody
*/
type GetCurrentKeyBlockHashOKBody struct {

	// hash
	Hash string `json:"hash,omitempty"`
}

// Validate validates this get current key block hash o k body
func (o *GetCurrentKeyBlockHashOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCurrentKeyBlockHashOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCurrentKeyBlockHashOKBody) UnmarshalBinary(b []byte) error {
	var res GetCurrentKeyBlockHashOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
