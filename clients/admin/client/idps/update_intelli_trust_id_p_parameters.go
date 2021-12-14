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

// NewUpdateIntelliTrustIDPParams creates a new UpdateIntelliTrustIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIntelliTrustIDPParams() *UpdateIntelliTrustIDPParams {
	return &UpdateIntelliTrustIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntelliTrustIDPParamsWithTimeout creates a new UpdateIntelliTrustIDPParams object
// with the ability to set a timeout on a request.
func NewUpdateIntelliTrustIDPParamsWithTimeout(timeout time.Duration) *UpdateIntelliTrustIDPParams {
	return &UpdateIntelliTrustIDPParams{
		timeout: timeout,
	}
}

// NewUpdateIntelliTrustIDPParamsWithContext creates a new UpdateIntelliTrustIDPParams object
// with the ability to set a context for a request.
func NewUpdateIntelliTrustIDPParamsWithContext(ctx context.Context) *UpdateIntelliTrustIDPParams {
	return &UpdateIntelliTrustIDPParams{
		Context: ctx,
	}
}

// NewUpdateIntelliTrustIDPParamsWithHTTPClient creates a new UpdateIntelliTrustIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIntelliTrustIDPParamsWithHTTPClient(client *http.Client) *UpdateIntelliTrustIDPParams {
	return &UpdateIntelliTrustIDPParams{
		HTTPClient: client,
	}
}

/* UpdateIntelliTrustIDPParams contains all the parameters to send to the API endpoint
   for the update intelli trust ID p operation.

   Typically these are written to a http.Request.
*/
type UpdateIntelliTrustIDPParams struct {

	/* IntelliTrustIDP.

	   IntelliTrustIDP
	*/
	IntelliTrustIDP *models.IntelliTrustIDP

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

// WithDefaults hydrates default values in the update intelli trust ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntelliTrustIDPParams) WithDefaults() *UpdateIntelliTrustIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update intelli trust ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntelliTrustIDPParams) SetDefaults() {
	var (
		iidDefault = string("default")

		widDefault = string("default")
	)

	val := UpdateIntelliTrustIDPParams{
		Iid: iidDefault,
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithTimeout(timeout time.Duration) *UpdateIntelliTrustIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithContext(ctx context.Context) *UpdateIntelliTrustIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithHTTPClient(client *http.Client) *UpdateIntelliTrustIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntelliTrustIDP adds the intelliTrustIDP to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) *UpdateIntelliTrustIDPParams {
	o.SetIntelliTrustIDP(intelliTrustIDP)
	return o
}

// SetIntelliTrustIDP adds the intelliTrustIdP to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) {
	o.IntelliTrustIDP = intelliTrustIDP
}

// WithIid adds the iid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithIid(iid string) *UpdateIntelliTrustIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithWid(wid string) *UpdateIntelliTrustIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntelliTrustIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IntelliTrustIDP != nil {
		if err := r.SetBodyParam(o.IntelliTrustIDP); err != nil {
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
