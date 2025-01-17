// Code generated by go-swagger; DO NOT EDIT.

package idps

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

// NewCreateIdentityPoolIDPParams creates a new CreateIdentityPoolIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIdentityPoolIDPParams() *CreateIdentityPoolIDPParams {
	return &CreateIdentityPoolIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIdentityPoolIDPParamsWithTimeout creates a new CreateIdentityPoolIDPParams object
// with the ability to set a timeout on a request.
func NewCreateIdentityPoolIDPParamsWithTimeout(timeout time.Duration) *CreateIdentityPoolIDPParams {
	return &CreateIdentityPoolIDPParams{
		timeout: timeout,
	}
}

// NewCreateIdentityPoolIDPParamsWithContext creates a new CreateIdentityPoolIDPParams object
// with the ability to set a context for a request.
func NewCreateIdentityPoolIDPParamsWithContext(ctx context.Context) *CreateIdentityPoolIDPParams {
	return &CreateIdentityPoolIDPParams{
		Context: ctx,
	}
}

// NewCreateIdentityPoolIDPParamsWithHTTPClient creates a new CreateIdentityPoolIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIdentityPoolIDPParamsWithHTTPClient(client *http.Client) *CreateIdentityPoolIDPParams {
	return &CreateIdentityPoolIDPParams{
		HTTPClient: client,
	}
}

/* CreateIdentityPoolIDPParams contains all the parameters to send to the API endpoint
   for the create identity pool ID p operation.

   Typically these are written to a http.Request.
*/
type CreateIdentityPoolIDPParams struct {

	/* IdentityPoolIDP.

	   IdentityPoolIDP
	*/
	IdentityPoolIDP *models.IdentityPoolIDP

	/* Wid.

	   Authorization server id

	   Default: "default"
	*/
	Wid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create identity pool ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentityPoolIDPParams) WithDefaults() *CreateIdentityPoolIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create identity pool ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentityPoolIDPParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := CreateIdentityPoolIDPParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) WithTimeout(timeout time.Duration) *CreateIdentityPoolIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) WithContext(ctx context.Context) *CreateIdentityPoolIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) WithHTTPClient(client *http.Client) *CreateIdentityPoolIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityPoolIDP adds the identityPoolIDP to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) WithIdentityPoolIDP(identityPoolIDP *models.IdentityPoolIDP) *CreateIdentityPoolIDPParams {
	o.SetIdentityPoolIDP(identityPoolIDP)
	return o
}

// SetIdentityPoolIDP adds the identityPoolIdP to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) SetIdentityPoolIDP(identityPoolIDP *models.IdentityPoolIDP) {
	o.IdentityPoolIDP = identityPoolIDP
}

// WithWid adds the wid to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) WithWid(wid string) *CreateIdentityPoolIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the create identity pool ID p params
func (o *CreateIdentityPoolIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIdentityPoolIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IdentityPoolIDP != nil {
		if err := r.SetBodyParam(o.IdentityPoolIDP); err != nil {
			return err
		}
	}

	// path param wid
	if err := r.SetPathParam("wid", o.Wid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
