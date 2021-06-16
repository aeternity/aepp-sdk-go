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

// NewGenerateACIParams creates a new GenerateACIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateACIParams() *GenerateACIParams {
	return &GenerateACIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateACIParamsWithTimeout creates a new GenerateACIParams object
// with the ability to set a timeout on a request.
func NewGenerateACIParamsWithTimeout(timeout time.Duration) *GenerateACIParams {
	return &GenerateACIParams{
		timeout: timeout,
	}
}

// NewGenerateACIParamsWithContext creates a new GenerateACIParams object
// with the ability to set a context for a request.
func NewGenerateACIParamsWithContext(ctx context.Context) *GenerateACIParams {
	return &GenerateACIParams{
		Context: ctx,
	}
}

// NewGenerateACIParamsWithHTTPClient creates a new GenerateACIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateACIParamsWithHTTPClient(client *http.Client) *GenerateACIParams {
	return &GenerateACIParams{
		HTTPClient: client,
	}
}

/* GenerateACIParams contains all the parameters to send to the API endpoint
   for the generate a c i operation.

   Typically these are written to a http.Request.
*/
type GenerateACIParams struct {

	/* Body.

	   contract code
	*/
	Body *models.Contract

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate a c i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateACIParams) WithDefaults() *GenerateACIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate a c i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateACIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate a c i params
func (o *GenerateACIParams) WithTimeout(timeout time.Duration) *GenerateACIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate a c i params
func (o *GenerateACIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate a c i params
func (o *GenerateACIParams) WithContext(ctx context.Context) *GenerateACIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate a c i params
func (o *GenerateACIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate a c i params
func (o *GenerateACIParams) WithHTTPClient(client *http.Client) *GenerateACIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate a c i params
func (o *GenerateACIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate a c i params
func (o *GenerateACIParams) WithBody(body *models.Contract) *GenerateACIParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate a c i params
func (o *GenerateACIParams) SetBody(body *models.Contract) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateACIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
