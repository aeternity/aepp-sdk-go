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

	"github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

// NewProtectedDryRunTxsParams creates a new ProtectedDryRunTxsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProtectedDryRunTxsParams() *ProtectedDryRunTxsParams {
	return &ProtectedDryRunTxsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProtectedDryRunTxsParamsWithTimeout creates a new ProtectedDryRunTxsParams object
// with the ability to set a timeout on a request.
func NewProtectedDryRunTxsParamsWithTimeout(timeout time.Duration) *ProtectedDryRunTxsParams {
	return &ProtectedDryRunTxsParams{
		timeout: timeout,
	}
}

// NewProtectedDryRunTxsParamsWithContext creates a new ProtectedDryRunTxsParams object
// with the ability to set a context for a request.
func NewProtectedDryRunTxsParamsWithContext(ctx context.Context) *ProtectedDryRunTxsParams {
	return &ProtectedDryRunTxsParams{
		Context: ctx,
	}
}

// NewProtectedDryRunTxsParamsWithHTTPClient creates a new ProtectedDryRunTxsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProtectedDryRunTxsParamsWithHTTPClient(client *http.Client) *ProtectedDryRunTxsParams {
	return &ProtectedDryRunTxsParams{
		HTTPClient: client,
	}
}

/* ProtectedDryRunTxsParams contains all the parameters to send to the API endpoint
   for the protected dry run txs operation.

   Typically these are written to a http.Request.
*/
type ProtectedDryRunTxsParams struct {

	/* Body.

	   transactions
	*/
	Body *models.DryRunInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the protected dry run txs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProtectedDryRunTxsParams) WithDefaults() *ProtectedDryRunTxsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the protected dry run txs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProtectedDryRunTxsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) WithTimeout(timeout time.Duration) *ProtectedDryRunTxsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) WithContext(ctx context.Context) *ProtectedDryRunTxsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) WithHTTPClient(client *http.Client) *ProtectedDryRunTxsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) WithBody(body *models.DryRunInput) *ProtectedDryRunTxsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the protected dry run txs params
func (o *ProtectedDryRunTxsParams) SetBody(body *models.DryRunInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProtectedDryRunTxsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
