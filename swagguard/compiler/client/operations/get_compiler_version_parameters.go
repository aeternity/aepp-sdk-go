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

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/compiler/models"
)

// NewGetCompilerVersionParams creates a new GetCompilerVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompilerVersionParams() *GetCompilerVersionParams {
	return &GetCompilerVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompilerVersionParamsWithTimeout creates a new GetCompilerVersionParams object
// with the ability to set a timeout on a request.
func NewGetCompilerVersionParamsWithTimeout(timeout time.Duration) *GetCompilerVersionParams {
	return &GetCompilerVersionParams{
		timeout: timeout,
	}
}

// NewGetCompilerVersionParamsWithContext creates a new GetCompilerVersionParams object
// with the ability to set a context for a request.
func NewGetCompilerVersionParamsWithContext(ctx context.Context) *GetCompilerVersionParams {
	return &GetCompilerVersionParams{
		Context: ctx,
	}
}

// NewGetCompilerVersionParamsWithHTTPClient creates a new GetCompilerVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompilerVersionParamsWithHTTPClient(client *http.Client) *GetCompilerVersionParams {
	return &GetCompilerVersionParams{
		HTTPClient: client,
	}
}

/* GetCompilerVersionParams contains all the parameters to send to the API endpoint
   for the get compiler version operation.

   Typically these are written to a http.Request.
*/
type GetCompilerVersionParams struct {

	/* Body.

	   contract byte array
	*/
	Body *models.ByteCodeInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compiler version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompilerVersionParams) WithDefaults() *GetCompilerVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compiler version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompilerVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compiler version params
func (o *GetCompilerVersionParams) WithTimeout(timeout time.Duration) *GetCompilerVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compiler version params
func (o *GetCompilerVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compiler version params
func (o *GetCompilerVersionParams) WithContext(ctx context.Context) *GetCompilerVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compiler version params
func (o *GetCompilerVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compiler version params
func (o *GetCompilerVersionParams) WithHTTPClient(client *http.Client) *GetCompilerVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compiler version params
func (o *GetCompilerVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get compiler version params
func (o *GetCompilerVersionParams) WithBody(body *models.ByteCodeInput) *GetCompilerVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get compiler version params
func (o *GetCompilerVersionParams) SetBody(body *models.ByteCodeInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompilerVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
