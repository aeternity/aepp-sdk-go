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

// NewGetContractPoIParams creates a new GetContractPoIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContractPoIParams() *GetContractPoIParams {
	return &GetContractPoIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContractPoIParamsWithTimeout creates a new GetContractPoIParams object
// with the ability to set a timeout on a request.
func NewGetContractPoIParamsWithTimeout(timeout time.Duration) *GetContractPoIParams {
	return &GetContractPoIParams{
		timeout: timeout,
	}
}

// NewGetContractPoIParamsWithContext creates a new GetContractPoIParams object
// with the ability to set a context for a request.
func NewGetContractPoIParamsWithContext(ctx context.Context) *GetContractPoIParams {
	return &GetContractPoIParams{
		Context: ctx,
	}
}

// NewGetContractPoIParamsWithHTTPClient creates a new GetContractPoIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContractPoIParamsWithHTTPClient(client *http.Client) *GetContractPoIParams {
	return &GetContractPoIParams{
		HTTPClient: client,
	}
}

/* GetContractPoIParams contains all the parameters to send to the API endpoint
   for the get contract po i operation.

   Typically these are written to a http.Request.
*/
type GetContractPoIParams struct {

	/* Pubkey.

	   Contract pubkey to get proof for
	*/
	Pubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contract po i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContractPoIParams) WithDefaults() *GetContractPoIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contract po i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContractPoIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contract po i params
func (o *GetContractPoIParams) WithTimeout(timeout time.Duration) *GetContractPoIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contract po i params
func (o *GetContractPoIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contract po i params
func (o *GetContractPoIParams) WithContext(ctx context.Context) *GetContractPoIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contract po i params
func (o *GetContractPoIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contract po i params
func (o *GetContractPoIParams) WithHTTPClient(client *http.Client) *GetContractPoIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contract po i params
func (o *GetContractPoIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPubkey adds the pubkey to the get contract po i params
func (o *GetContractPoIParams) WithPubkey(pubkey string) *GetContractPoIParams {
	o.SetPubkey(pubkey)
	return o
}

// SetPubkey adds the pubkey to the get contract po i params
func (o *GetContractPoIParams) SetPubkey(pubkey string) {
	o.Pubkey = pubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetContractPoIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pubkey
	if err := r.SetPathParam("pubkey", o.Pubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
