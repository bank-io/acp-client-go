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

// NewUpdateIdentityPoolIDPParams creates a new UpdateIdentityPoolIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIdentityPoolIDPParams() *UpdateIdentityPoolIDPParams {
	return &UpdateIdentityPoolIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIdentityPoolIDPParamsWithTimeout creates a new UpdateIdentityPoolIDPParams object
// with the ability to set a timeout on a request.
func NewUpdateIdentityPoolIDPParamsWithTimeout(timeout time.Duration) *UpdateIdentityPoolIDPParams {
	return &UpdateIdentityPoolIDPParams{
		timeout: timeout,
	}
}

// NewUpdateIdentityPoolIDPParamsWithContext creates a new UpdateIdentityPoolIDPParams object
// with the ability to set a context for a request.
func NewUpdateIdentityPoolIDPParamsWithContext(ctx context.Context) *UpdateIdentityPoolIDPParams {
	return &UpdateIdentityPoolIDPParams{
		Context: ctx,
	}
}

// NewUpdateIdentityPoolIDPParamsWithHTTPClient creates a new UpdateIdentityPoolIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIdentityPoolIDPParamsWithHTTPClient(client *http.Client) *UpdateIdentityPoolIDPParams {
	return &UpdateIdentityPoolIDPParams{
		HTTPClient: client,
	}
}

/* UpdateIdentityPoolIDPParams contains all the parameters to send to the API endpoint
   for the update identity pool ID p operation.

   Typically these are written to a http.Request.
*/
type UpdateIdentityPoolIDPParams struct {

	/* IdentityPoolIDP.

	   IdentityPoolIDP
	*/
	IdentityPoolIDP *models.IdentityPoolIDP

	/* Iid.

	   IDP id

	   Default: "default"
	*/
	Iid string

	/* Wid.

	   Authorization server id

	   Default: "default"
	*/
	Wid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update identity pool ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIdentityPoolIDPParams) WithDefaults() *UpdateIdentityPoolIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update identity pool ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIdentityPoolIDPParams) SetDefaults() {
	var (
		iidDefault = string("default")

		widDefault = string("default")
	)

	val := UpdateIdentityPoolIDPParams{
		Iid: iidDefault,
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithTimeout(timeout time.Duration) *UpdateIdentityPoolIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithContext(ctx context.Context) *UpdateIdentityPoolIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithHTTPClient(client *http.Client) *UpdateIdentityPoolIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityPoolIDP adds the identityPoolIDP to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithIdentityPoolIDP(identityPoolIDP *models.IdentityPoolIDP) *UpdateIdentityPoolIDPParams {
	o.SetIdentityPoolIDP(identityPoolIDP)
	return o
}

// SetIdentityPoolIDP adds the identityPoolIdP to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetIdentityPoolIDP(identityPoolIDP *models.IdentityPoolIDP) {
	o.IdentityPoolIDP = identityPoolIDP
}

// WithIid adds the iid to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithIid(iid string) *UpdateIdentityPoolIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) WithWid(wid string) *UpdateIdentityPoolIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the update identity pool ID p params
func (o *UpdateIdentityPoolIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIdentityPoolIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IdentityPoolIDP != nil {
		if err := r.SetBodyParam(o.IdentityPoolIDP); err != nil {
			return err
		}
	}

	// path param iid
	if err := r.SetPathParam("iid", o.Iid); err != nil {
		return err
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