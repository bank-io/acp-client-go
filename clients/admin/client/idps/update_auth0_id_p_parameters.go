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

// NewUpdateAuth0IDPParams creates a new UpdateAuth0IDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAuth0IDPParams() *UpdateAuth0IDPParams {
	return &UpdateAuth0IDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAuth0IDPParamsWithTimeout creates a new UpdateAuth0IDPParams object
// with the ability to set a timeout on a request.
func NewUpdateAuth0IDPParamsWithTimeout(timeout time.Duration) *UpdateAuth0IDPParams {
	return &UpdateAuth0IDPParams{
		timeout: timeout,
	}
}

// NewUpdateAuth0IDPParamsWithContext creates a new UpdateAuth0IDPParams object
// with the ability to set a context for a request.
func NewUpdateAuth0IDPParamsWithContext(ctx context.Context) *UpdateAuth0IDPParams {
	return &UpdateAuth0IDPParams{
		Context: ctx,
	}
}

// NewUpdateAuth0IDPParamsWithHTTPClient creates a new UpdateAuth0IDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAuth0IDPParamsWithHTTPClient(client *http.Client) *UpdateAuth0IDPParams {
	return &UpdateAuth0IDPParams{
		HTTPClient: client,
	}
}

/* UpdateAuth0IDPParams contains all the parameters to send to the API endpoint
   for the update auth0 ID p operation.

   Typically these are written to a http.Request.
*/
type UpdateAuth0IDPParams struct {

	/* Auth0IDP.

	   Auth0IDP
	*/
	Auth0IDP *models.Auth0IDP

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

// WithDefaults hydrates default values in the update auth0 ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAuth0IDPParams) WithDefaults() *UpdateAuth0IDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update auth0 ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAuth0IDPParams) SetDefaults() {
	var (
		iidDefault = string("default")

		widDefault = string("default")
	)

	val := UpdateAuth0IDPParams{
		Iid: iidDefault,
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithTimeout(timeout time.Duration) *UpdateAuth0IDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithContext(ctx context.Context) *UpdateAuth0IDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithHTTPClient(client *http.Client) *UpdateAuth0IDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuth0IDP adds the auth0IDP to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithAuth0IDP(auth0IDP *models.Auth0IDP) *UpdateAuth0IDPParams {
	o.SetAuth0IDP(auth0IDP)
	return o
}

// SetAuth0IDP adds the auth0IdP to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetAuth0IDP(auth0IDP *models.Auth0IDP) {
	o.Auth0IDP = auth0IDP
}

// WithIid adds the iid to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithIid(iid string) *UpdateAuth0IDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) WithWid(wid string) *UpdateAuth0IDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the update auth0 ID p params
func (o *UpdateAuth0IDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAuth0IDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Auth0IDP != nil {
		if err := r.SetBodyParam(o.Auth0IDP); err != nil {
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
