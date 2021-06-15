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

// NewDecodeCalldataSourceParams creates a new DecodeCalldataSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDecodeCalldataSourceParams() *DecodeCalldataSourceParams {
	return &DecodeCalldataSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDecodeCalldataSourceParamsWithTimeout creates a new DecodeCalldataSourceParams object
// with the ability to set a timeout on a request.
func NewDecodeCalldataSourceParamsWithTimeout(timeout time.Duration) *DecodeCalldataSourceParams {
	return &DecodeCalldataSourceParams{
		timeout: timeout,
	}
}

// NewDecodeCalldataSourceParamsWithContext creates a new DecodeCalldataSourceParams object
// with the ability to set a context for a request.
func NewDecodeCalldataSourceParamsWithContext(ctx context.Context) *DecodeCalldataSourceParams {
	return &DecodeCalldataSourceParams{
		Context: ctx,
	}
}

// NewDecodeCalldataSourceParamsWithHTTPClient creates a new DecodeCalldataSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDecodeCalldataSourceParamsWithHTTPClient(client *http.Client) *DecodeCalldataSourceParams {
	return &DecodeCalldataSourceParams{
		HTTPClient: client,
	}
}

/* DecodeCalldataSourceParams contains all the parameters to send to the API endpoint
   for the decode calldata source operation.

   Typically these are written to a http.Request.
*/
type DecodeCalldataSourceParams struct {

	/* Body.

	   Calldata + contract (stub) code
	*/
	Body *models.DecodeCalldataSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the decode calldata source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecodeCalldataSourceParams) WithDefaults() *DecodeCalldataSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the decode calldata source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecodeCalldataSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the decode calldata source params
func (o *DecodeCalldataSourceParams) WithTimeout(timeout time.Duration) *DecodeCalldataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decode calldata source params
func (o *DecodeCalldataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decode calldata source params
func (o *DecodeCalldataSourceParams) WithContext(ctx context.Context) *DecodeCalldataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decode calldata source params
func (o *DecodeCalldataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decode calldata source params
func (o *DecodeCalldataSourceParams) WithHTTPClient(client *http.Client) *DecodeCalldataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decode calldata source params
func (o *DecodeCalldataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the decode calldata source params
func (o *DecodeCalldataSourceParams) WithBody(body *models.DecodeCalldataSource) *DecodeCalldataSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the decode calldata source params
func (o *DecodeCalldataSourceParams) SetBody(body *models.DecodeCalldataSource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DecodeCalldataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
