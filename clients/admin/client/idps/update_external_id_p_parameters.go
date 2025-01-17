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

// NewUpdateExternalIDPParams creates a new UpdateExternalIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateExternalIDPParams() *UpdateExternalIDPParams {
	return &UpdateExternalIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateExternalIDPParamsWithTimeout creates a new UpdateExternalIDPParams object
// with the ability to set a timeout on a request.
func NewUpdateExternalIDPParamsWithTimeout(timeout time.Duration) *UpdateExternalIDPParams {
	return &UpdateExternalIDPParams{
		timeout: timeout,
	}
}

// NewUpdateExternalIDPParamsWithContext creates a new UpdateExternalIDPParams object
// with the ability to set a context for a request.
func NewUpdateExternalIDPParamsWithContext(ctx context.Context) *UpdateExternalIDPParams {
	return &UpdateExternalIDPParams{
		Context: ctx,
	}
}

// NewUpdateExternalIDPParamsWithHTTPClient creates a new UpdateExternalIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateExternalIDPParamsWithHTTPClient(client *http.Client) *UpdateExternalIDPParams {
	return &UpdateExternalIDPParams{
		HTTPClient: client,
	}
}

/* UpdateExternalIDPParams contains all the parameters to send to the API endpoint
   for the update external ID p operation.

   Typically these are written to a http.Request.
*/
type UpdateExternalIDPParams struct {

	/* ExternalIDP.

	   ExternalIDP
	*/
	ExternalIDP *models.ExternalIDP

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

// WithDefaults hydrates default values in the update external ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExternalIDPParams) WithDefaults() *UpdateExternalIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update external ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExternalIDPParams) SetDefaults() {
	var (
		iidDefault = string("default")

		widDefault = string("default")
	)

	val := UpdateExternalIDPParams{
		Iid: iidDefault,
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update external ID p params
func (o *UpdateExternalIDPParams) WithTimeout(timeout time.Duration) *UpdateExternalIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update external ID p params
func (o *UpdateExternalIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update external ID p params
func (o *UpdateExternalIDPParams) WithContext(ctx context.Context) *UpdateExternalIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update external ID p params
func (o *UpdateExternalIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update external ID p params
func (o *UpdateExternalIDPParams) WithHTTPClient(client *http.Client) *UpdateExternalIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update external ID p params
func (o *UpdateExternalIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalIDP adds the externalIDP to the update external ID p params
func (o *UpdateExternalIDPParams) WithExternalIDP(externalIDP *models.ExternalIDP) *UpdateExternalIDPParams {
	o.SetExternalIDP(externalIDP)
	return o
}

// SetExternalIDP adds the externalIdP to the update external ID p params
func (o *UpdateExternalIDPParams) SetExternalIDP(externalIDP *models.ExternalIDP) {
	o.ExternalIDP = externalIDP
}

// WithIid adds the iid to the update external ID p params
func (o *UpdateExternalIDPParams) WithIid(iid string) *UpdateExternalIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update external ID p params
func (o *UpdateExternalIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the update external ID p params
func (o *UpdateExternalIDPParams) WithWid(wid string) *UpdateExternalIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the update external ID p params
func (o *UpdateExternalIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateExternalIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ExternalIDP != nil {
		if err := r.SetBodyParam(o.ExternalIDP); err != nil {
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
