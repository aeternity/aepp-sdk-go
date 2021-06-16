// Code generated by go-swagger; DO NOT EDIT.

package external

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
)

// NewGetChainEndsParams creates a new GetChainEndsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChainEndsParams() *GetChainEndsParams {
	return &GetChainEndsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChainEndsParamsWithTimeout creates a new GetChainEndsParams object
// with the ability to set a timeout on a request.
func NewGetChainEndsParamsWithTimeout(timeout time.Duration) *GetChainEndsParams {
	return &GetChainEndsParams{
		timeout: timeout,
	}
}

// NewGetChainEndsParamsWithContext creates a new GetChainEndsParams object
// with the ability to set a context for a request.
func NewGetChainEndsParamsWithContext(ctx context.Context) *GetChainEndsParams {
	return &GetChainEndsParams{
		Context: ctx,
	}
}

// NewGetChainEndsParamsWithHTTPClient creates a new GetChainEndsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChainEndsParamsWithHTTPClient(client *http.Client) *GetChainEndsParams {
	return &GetChainEndsParams{
		HTTPClient: client,
	}
}

/* GetChainEndsParams contains all the parameters to send to the API endpoint
   for the get chain ends operation.

   Typically these are written to a http.Request.
*/
type GetChainEndsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get chain ends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChainEndsParams) WithDefaults() *GetChainEndsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get chain ends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChainEndsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get chain ends params
func (o *GetChainEndsParams) WithTimeout(timeout time.Duration) *GetChainEndsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chain ends params
func (o *GetChainEndsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chain ends params
func (o *GetChainEndsParams) WithContext(ctx context.Context) *GetChainEndsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chain ends params
func (o *GetChainEndsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chain ends params
func (o *GetChainEndsParams) WithHTTPClient(client *http.Client) *GetChainEndsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chain ends params
func (o *GetChainEndsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetChainEndsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}