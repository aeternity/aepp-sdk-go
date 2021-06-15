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
	"github.com/go-openapi/strfmt"

	"github.com/aeternity/aepp-sdk-go/v9/swagguard/compiler/models"
)

// NewDecodeCallResultParams creates a new DecodeCallResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDecodeCallResultParams() *DecodeCallResultParams {
	return &DecodeCallResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDecodeCallResultParamsWithTimeout creates a new DecodeCallResultParams object
// with the ability to set a timeout on a request.
func NewDecodeCallResultParamsWithTimeout(timeout time.Duration) *DecodeCallResultParams {
	return &DecodeCallResultParams{
		timeout: timeout,
	}
}

// NewDecodeCallResultParamsWithContext creates a new DecodeCallResultParams object
// with the ability to set a context for a request.
func NewDecodeCallResultParamsWithContext(ctx context.Context) *DecodeCallResultParams {
	return &DecodeCallResultParams{
		Context: ctx,
	}
}

// NewDecodeCallResultParamsWithHTTPClient creates a new DecodeCallResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewDecodeCallResultParamsWithHTTPClient(client *http.Client) *DecodeCallResultParams {
	return &DecodeCallResultParams{
		HTTPClient: client,
	}
}

/* DecodeCallResultParams contains all the parameters to send to the API endpoint
   for the decode call result operation.

   Typically these are written to a http.Request.
*/
type DecodeCallResultParams struct {

	/* Body.

	   Binary data in Sophia ABI format
	*/
	Body *models.SophiaCallResultInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the decode call result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecodeCallResultParams) WithDefaults() *DecodeCallResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the decode call result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecodeCallResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the decode call result params
func (o *DecodeCallResultParams) WithTimeout(timeout time.Duration) *DecodeCallResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decode call result params
func (o *DecodeCallResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decode call result params
func (o *DecodeCallResultParams) WithContext(ctx context.Context) *DecodeCallResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decode call result params
func (o *DecodeCallResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decode call result params
func (o *DecodeCallResultParams) WithHTTPClient(client *http.Client) *DecodeCallResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decode call result params
func (o *DecodeCallResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the decode call result params
func (o *DecodeCallResultParams) WithBody(body *models.SophiaCallResultInput) *DecodeCallResultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the decode call result params
func (o *DecodeCallResultParams) SetBody(body *models.SophiaCallResultInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DecodeCallResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
