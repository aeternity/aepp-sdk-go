// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// NewPostChannelSnapshotSoloParams creates a new PostChannelSnapshotSoloParams object
// with the default values initialized.
func NewPostChannelSnapshotSoloParams() *PostChannelSnapshotSoloParams {
	var ()
	return &PostChannelSnapshotSoloParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostChannelSnapshotSoloParamsWithTimeout creates a new PostChannelSnapshotSoloParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostChannelSnapshotSoloParamsWithTimeout(timeout time.Duration) *PostChannelSnapshotSoloParams {
	var ()
	return &PostChannelSnapshotSoloParams{

		timeout: timeout,
	}
}

// NewPostChannelSnapshotSoloParamsWithContext creates a new PostChannelSnapshotSoloParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostChannelSnapshotSoloParamsWithContext(ctx context.Context) *PostChannelSnapshotSoloParams {
	var ()
	return &PostChannelSnapshotSoloParams{

		Context: ctx,
	}
}

// NewPostChannelSnapshotSoloParamsWithHTTPClient creates a new PostChannelSnapshotSoloParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostChannelSnapshotSoloParamsWithHTTPClient(client *http.Client) *PostChannelSnapshotSoloParams {
	var ()
	return &PostChannelSnapshotSoloParams{
		HTTPClient: client,
	}
}

/*PostChannelSnapshotSoloParams contains all the parameters to send to the API endpoint
for the post channel snapshot solo operation typically these are written to a http.Request
*/
type PostChannelSnapshotSoloParams struct {

	/*Body*/
	Body *models.ChannelSnapshotSoloTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) WithTimeout(timeout time.Duration) *PostChannelSnapshotSoloParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) WithContext(ctx context.Context) *PostChannelSnapshotSoloParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) WithHTTPClient(client *http.Client) *PostChannelSnapshotSoloParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) WithBody(body *models.ChannelSnapshotSoloTx) *PostChannelSnapshotSoloParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post channel snapshot solo params
func (o *PostChannelSnapshotSoloParams) SetBody(body *models.ChannelSnapshotSoloTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostChannelSnapshotSoloParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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