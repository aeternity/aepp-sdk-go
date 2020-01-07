// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v8/swagguard/compiler/models"
)

// NewDecodeCallResultBytecodeParams creates a new DecodeCallResultBytecodeParams object
// with the default values initialized.
func NewDecodeCallResultBytecodeParams() *DecodeCallResultBytecodeParams {
	var ()
	return &DecodeCallResultBytecodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDecodeCallResultBytecodeParamsWithTimeout creates a new DecodeCallResultBytecodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDecodeCallResultBytecodeParamsWithTimeout(timeout time.Duration) *DecodeCallResultBytecodeParams {
	var ()
	return &DecodeCallResultBytecodeParams{

		timeout: timeout,
	}
}

// NewDecodeCallResultBytecodeParamsWithContext creates a new DecodeCallResultBytecodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDecodeCallResultBytecodeParamsWithContext(ctx context.Context) *DecodeCallResultBytecodeParams {
	var ()
	return &DecodeCallResultBytecodeParams{

		Context: ctx,
	}
}

// NewDecodeCallResultBytecodeParamsWithHTTPClient creates a new DecodeCallResultBytecodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDecodeCallResultBytecodeParamsWithHTTPClient(client *http.Client) *DecodeCallResultBytecodeParams {
	var ()
	return &DecodeCallResultBytecodeParams{
		HTTPClient: client,
	}
}

/*DecodeCallResultBytecodeParams contains all the parameters to send to the API endpoint
for the decode call result bytecode operation typically these are written to a http.Request
*/
type DecodeCallResultBytecodeParams struct {

	/*Body
	  Call result + compiled contract

	*/
	Body *models.BytecodeCallResultInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) WithTimeout(timeout time.Duration) *DecodeCallResultBytecodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) WithContext(ctx context.Context) *DecodeCallResultBytecodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) WithHTTPClient(client *http.Client) *DecodeCallResultBytecodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) WithBody(body *models.BytecodeCallResultInput) *DecodeCallResultBytecodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the decode call result bytecode params
func (o *DecodeCallResultBytecodeParams) SetBody(body *models.BytecodeCallResultInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DecodeCallResultBytecodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
