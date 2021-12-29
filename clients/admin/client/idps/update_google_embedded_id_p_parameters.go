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

// NewUpdateGoogleEmbeddedIDPParams creates a new UpdateGoogleEmbeddedIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGoogleEmbeddedIDPParams() *UpdateGoogleEmbeddedIDPParams {
	return &UpdateGoogleEmbeddedIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGoogleEmbeddedIDPParamsWithTimeout creates a new UpdateGoogleEmbeddedIDPParams object
// with the ability to set a timeout on a request.
func NewUpdateGoogleEmbeddedIDPParamsWithTimeout(timeout time.Duration) *UpdateGoogleEmbeddedIDPParams {
	return &UpdateGoogleEmbeddedIDPParams{
		timeout: timeout,
	}
}

// NewUpdateGoogleEmbeddedIDPParamsWithContext creates a new UpdateGoogleEmbeddedIDPParams object
// with the ability to set a context for a request.
func NewUpdateGoogleEmbeddedIDPParamsWithContext(ctx context.Context) *UpdateGoogleEmbeddedIDPParams {
	return &UpdateGoogleEmbeddedIDPParams{
		Context: ctx,
	}
}

// NewUpdateGoogleEmbeddedIDPParamsWithHTTPClient creates a new UpdateGoogleEmbeddedIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGoogleEmbeddedIDPParamsWithHTTPClient(client *http.Client) *UpdateGoogleEmbeddedIDPParams {
	return &UpdateGoogleEmbeddedIDPParams{
		HTTPClient: client,
	}
}

/* UpdateGoogleEmbeddedIDPParams contains all the parameters to send to the API endpoint
   for the update google embedded ID p operation.

   Typically these are written to a http.Request.
*/
type UpdateGoogleEmbeddedIDPParams struct {

	/* GoogleEmbeddedIDP.

	   GoogleEmbeddedIDP
	*/
	GoogleEmbeddedIDP *models.GoogleEmbeddedIDP

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

// WithDefaults hydrates default values in the update google embedded ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGoogleEmbeddedIDPParams) WithDefaults() *UpdateGoogleEmbeddedIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update google embedded ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGoogleEmbeddedIDPParams) SetDefaults() {
	var (
		iidDefault = string("default")

		widDefault = string("default")
	)

	val := UpdateGoogleEmbeddedIDPParams{
		Iid: iidDefault,
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithTimeout(timeout time.Duration) *UpdateGoogleEmbeddedIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithContext(ctx context.Context) *UpdateGoogleEmbeddedIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithHTTPClient(client *http.Client) *UpdateGoogleEmbeddedIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGoogleEmbeddedIDP adds the googleEmbeddedIDP to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithGoogleEmbeddedIDP(googleEmbeddedIDP *models.GoogleEmbeddedIDP) *UpdateGoogleEmbeddedIDPParams {
	o.SetGoogleEmbeddedIDP(googleEmbeddedIDP)
	return o
}

// SetGoogleEmbeddedIDP adds the googleEmbeddedIdP to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetGoogleEmbeddedIDP(googleEmbeddedIDP *models.GoogleEmbeddedIDP) {
	o.GoogleEmbeddedIDP = googleEmbeddedIDP
}

// WithIid adds the iid to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithIid(iid string) *UpdateGoogleEmbeddedIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) WithWid(wid string) *UpdateGoogleEmbeddedIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the update google embedded ID p params
func (o *UpdateGoogleEmbeddedIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGoogleEmbeddedIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GoogleEmbeddedIDP != nil {
		if err := r.SetBodyParam(o.GoogleEmbeddedIDP); err != nil {
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