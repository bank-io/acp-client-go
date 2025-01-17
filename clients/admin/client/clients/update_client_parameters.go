// Code generated by go-swagger; DO NOT EDIT.

package clients

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

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// NewUpdateClientParams creates a new UpdateClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClientParams() *UpdateClientParams {
	return &UpdateClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClientParamsWithTimeout creates a new UpdateClientParams object
// with the ability to set a timeout on a request.
func NewUpdateClientParamsWithTimeout(timeout time.Duration) *UpdateClientParams {
	return &UpdateClientParams{
		timeout: timeout,
	}
}

// NewUpdateClientParamsWithContext creates a new UpdateClientParams object
// with the ability to set a context for a request.
func NewUpdateClientParamsWithContext(ctx context.Context) *UpdateClientParams {
	return &UpdateClientParams{
		Context: ctx,
	}
}

// NewUpdateClientParamsWithHTTPClient creates a new UpdateClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClientParamsWithHTTPClient(client *http.Client) *UpdateClientParams {
	return &UpdateClientParams{
		HTTPClient: client,
	}
}

/* UpdateClientParams contains all the parameters to send to the API endpoint
   for the update client operation.

   Typically these are written to a http.Request.
*/
type UpdateClientParams struct {

	// Client.
	Client *models.UpdateClientAdminRequest

	/* Cid.

	   Client id

	   Default: "default"
	*/
	Cid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClientParams) WithDefaults() *UpdateClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClientParams) SetDefaults() {
	var (
		cidDefault = string("default")
	)

	val := UpdateClientParams{
		Cid: cidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update client params
func (o *UpdateClientParams) WithTimeout(timeout time.Duration) *UpdateClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update client params
func (o *UpdateClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update client params
func (o *UpdateClientParams) WithContext(ctx context.Context) *UpdateClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update client params
func (o *UpdateClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update client params
func (o *UpdateClientParams) WithHTTPClient(client *http.Client) *UpdateClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update client params
func (o *UpdateClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the update client params
func (o *UpdateClientParams) WithClient(client *models.UpdateClientAdminRequest) *UpdateClientParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the update client params
func (o *UpdateClientParams) SetClient(client *models.UpdateClientAdminRequest) {
	o.Client = client
}

// WithCid adds the cid to the update client params
func (o *UpdateClientParams) WithCid(cid string) *UpdateClientParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the update client params
func (o *UpdateClientParams) SetCid(cid string) {
	o.Cid = cid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Client != nil {
		if err := r.SetBodyParam(o.Client); err != nil {
			return err
		}
	}

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
