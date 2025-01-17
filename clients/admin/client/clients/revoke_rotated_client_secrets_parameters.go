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
)

// NewRevokeRotatedClientSecretsParams creates a new RevokeRotatedClientSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeRotatedClientSecretsParams() *RevokeRotatedClientSecretsParams {
	return &RevokeRotatedClientSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeRotatedClientSecretsParamsWithTimeout creates a new RevokeRotatedClientSecretsParams object
// with the ability to set a timeout on a request.
func NewRevokeRotatedClientSecretsParamsWithTimeout(timeout time.Duration) *RevokeRotatedClientSecretsParams {
	return &RevokeRotatedClientSecretsParams{
		timeout: timeout,
	}
}

// NewRevokeRotatedClientSecretsParamsWithContext creates a new RevokeRotatedClientSecretsParams object
// with the ability to set a context for a request.
func NewRevokeRotatedClientSecretsParamsWithContext(ctx context.Context) *RevokeRotatedClientSecretsParams {
	return &RevokeRotatedClientSecretsParams{
		Context: ctx,
	}
}

// NewRevokeRotatedClientSecretsParamsWithHTTPClient creates a new RevokeRotatedClientSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeRotatedClientSecretsParamsWithHTTPClient(client *http.Client) *RevokeRotatedClientSecretsParams {
	return &RevokeRotatedClientSecretsParams{
		HTTPClient: client,
	}
}

/* RevokeRotatedClientSecretsParams contains all the parameters to send to the API endpoint
   for the revoke rotated client secrets operation.

   Typically these are written to a http.Request.
*/
type RevokeRotatedClientSecretsParams struct {

	/* Cid.

	   Client id

	   Default: "default"
	*/
	Cid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke rotated client secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeRotatedClientSecretsParams) WithDefaults() *RevokeRotatedClientSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke rotated client secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeRotatedClientSecretsParams) SetDefaults() {
	var (
		cidDefault = string("default")
	)

	val := RevokeRotatedClientSecretsParams{
		Cid: cidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) WithTimeout(timeout time.Duration) *RevokeRotatedClientSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) WithContext(ctx context.Context) *RevokeRotatedClientSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) WithHTTPClient(client *http.Client) *RevokeRotatedClientSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) WithCid(cid string) *RevokeRotatedClientSecretsParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the revoke rotated client secrets params
func (o *RevokeRotatedClientSecretsParams) SetCid(cid string) {
	o.Cid = cid
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeRotatedClientSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
