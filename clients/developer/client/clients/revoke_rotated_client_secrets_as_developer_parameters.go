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

// NewRevokeRotatedClientSecretsAsDeveloperParams creates a new RevokeRotatedClientSecretsAsDeveloperParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeRotatedClientSecretsAsDeveloperParams() *RevokeRotatedClientSecretsAsDeveloperParams {
	return &RevokeRotatedClientSecretsAsDeveloperParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeRotatedClientSecretsAsDeveloperParamsWithTimeout creates a new RevokeRotatedClientSecretsAsDeveloperParams object
// with the ability to set a timeout on a request.
func NewRevokeRotatedClientSecretsAsDeveloperParamsWithTimeout(timeout time.Duration) *RevokeRotatedClientSecretsAsDeveloperParams {
	return &RevokeRotatedClientSecretsAsDeveloperParams{
		timeout: timeout,
	}
}

// NewRevokeRotatedClientSecretsAsDeveloperParamsWithContext creates a new RevokeRotatedClientSecretsAsDeveloperParams object
// with the ability to set a context for a request.
func NewRevokeRotatedClientSecretsAsDeveloperParamsWithContext(ctx context.Context) *RevokeRotatedClientSecretsAsDeveloperParams {
	return &RevokeRotatedClientSecretsAsDeveloperParams{
		Context: ctx,
	}
}

// NewRevokeRotatedClientSecretsAsDeveloperParamsWithHTTPClient creates a new RevokeRotatedClientSecretsAsDeveloperParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeRotatedClientSecretsAsDeveloperParamsWithHTTPClient(client *http.Client) *RevokeRotatedClientSecretsAsDeveloperParams {
	return &RevokeRotatedClientSecretsAsDeveloperParams{
		HTTPClient: client,
	}
}

/* RevokeRotatedClientSecretsAsDeveloperParams contains all the parameters to send to the API endpoint
   for the revoke rotated client secrets as developer operation.

   Typically these are written to a http.Request.
*/
type RevokeRotatedClientSecretsAsDeveloperParams struct {

	/* Cid.

	   Client id

	   Default: "default"
	*/
	Cid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke rotated client secrets as developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WithDefaults() *RevokeRotatedClientSecretsAsDeveloperParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke rotated client secrets as developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeRotatedClientSecretsAsDeveloperParams) SetDefaults() {
	var (
		cidDefault = string("default")
	)

	val := RevokeRotatedClientSecretsAsDeveloperParams{
		Cid: cidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WithTimeout(timeout time.Duration) *RevokeRotatedClientSecretsAsDeveloperParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WithContext(ctx context.Context) *RevokeRotatedClientSecretsAsDeveloperParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WithHTTPClient(client *http.Client) *RevokeRotatedClientSecretsAsDeveloperParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WithCid(cid string) *RevokeRotatedClientSecretsAsDeveloperParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the revoke rotated client secrets as developer params
func (o *RevokeRotatedClientSecretsAsDeveloperParams) SetCid(cid string) {
	o.Cid = cid
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeRotatedClientSecretsAsDeveloperParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
