// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPeerPubkeyParams creates a new GetPeerPubkeyParams object
// with the default values initialized.
func NewGetPeerPubkeyParams() *GetPeerPubkeyParams {

	return &GetPeerPubkeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeerPubkeyParamsWithTimeout creates a new GetPeerPubkeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeerPubkeyParamsWithTimeout(timeout time.Duration) *GetPeerPubkeyParams {

	return &GetPeerPubkeyParams{

		timeout: timeout,
	}
}

// NewGetPeerPubkeyParamsWithContext creates a new GetPeerPubkeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPeerPubkeyParamsWithContext(ctx context.Context) *GetPeerPubkeyParams {

	return &GetPeerPubkeyParams{

		Context: ctx,
	}
}

// NewGetPeerPubkeyParamsWithHTTPClient creates a new GetPeerPubkeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPeerPubkeyParamsWithHTTPClient(client *http.Client) *GetPeerPubkeyParams {

	return &GetPeerPubkeyParams{
		HTTPClient: client,
	}
}

/*GetPeerPubkeyParams contains all the parameters to send to the API endpoint
for the get peer pubkey operation typically these are written to a http.Request
*/
type GetPeerPubkeyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get peer pubkey params
func (o *GetPeerPubkeyParams) WithTimeout(timeout time.Duration) *GetPeerPubkeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get peer pubkey params
func (o *GetPeerPubkeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get peer pubkey params
func (o *GetPeerPubkeyParams) WithContext(ctx context.Context) *GetPeerPubkeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get peer pubkey params
func (o *GetPeerPubkeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get peer pubkey params
func (o *GetPeerPubkeyParams) WithHTTPClient(client *http.Client) *GetPeerPubkeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get peer pubkey params
func (o *GetPeerPubkeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeerPubkeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
