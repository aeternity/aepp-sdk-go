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

// GetGenerationByHeightReader is a Reader for the GetGenerationByHeight structure.
type GetGenerationByHeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGenerationByHeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGenerationByHeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetGenerationByHeightNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGenerationByHeightOK creates a GetGenerationByHeightOK with default headers values
func NewGetGenerationByHeightOK() *GetGenerationByHeightOK {
	return &GetGenerationByHeightOK{}
}

/*GetGenerationByHeightOK handles this case with default header values.

Successful operation
*/
type GetGenerationByHeightOK struct {
	Payload *models.Generation
}

func (o *GetGenerationByHeightOK) Error() string {
	return fmt.Sprintf("[GET /generations/height/{height}][%d] getGenerationByHeightOK  %+v", 200, o.Payload)
}

func (o *GetGenerationByHeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Generation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGenerationByHeightNotFound creates a GetGenerationByHeightNotFound with default headers values
func NewGetGenerationByHeightNotFound() *GetGenerationByHeightNotFound {
	return &GetGenerationByHeightNotFound{}
}

/*GetGenerationByHeightNotFound handles this case with default header values.

Generation not found
*/
type GetGenerationByHeightNotFound struct {
	Payload *models.Error
}

func (o *GetGenerationByHeightNotFound) Error() string {
	return fmt.Sprintf("[GET /generations/height/{height}][%d] getGenerationByHeightNotFound  %+v", 404, o.Payload)
}

func (o *GetGenerationByHeightNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
