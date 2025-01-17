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

// NewCreateIntelliTrustIDPParams creates a new CreateIntelliTrustIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIntelliTrustIDPParams() *CreateIntelliTrustIDPParams {
	return &CreateIntelliTrustIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIntelliTrustIDPParamsWithTimeout creates a new CreateIntelliTrustIDPParams object
// with the ability to set a timeout on a request.
func NewCreateIntelliTrustIDPParamsWithTimeout(timeout time.Duration) *CreateIntelliTrustIDPParams {
	return &CreateIntelliTrustIDPParams{
		timeout: timeout,
	}
}

// NewCreateIntelliTrustIDPParamsWithContext creates a new CreateIntelliTrustIDPParams object
// with the ability to set a context for a request.
func NewCreateIntelliTrustIDPParamsWithContext(ctx context.Context) *CreateIntelliTrustIDPParams {
	return &CreateIntelliTrustIDPParams{
		Context: ctx,
	}
}

// NewCreateIntelliTrustIDPParamsWithHTTPClient creates a new CreateIntelliTrustIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIntelliTrustIDPParamsWithHTTPClient(client *http.Client) *CreateIntelliTrustIDPParams {
	return &CreateIntelliTrustIDPParams{
		HTTPClient: client,
	}
}

/* CreateIntelliTrustIDPParams contains all the parameters to send to the API endpoint
   for the create intelli trust ID p operation.

   Typically these are written to a http.Request.
*/
type CreateIntelliTrustIDPParams struct {

	/* IntelliTrustIDP.

	   IntelliTrustIDP
	*/
	IntelliTrustIDP *models.IntelliTrustIDP

	/* Wid.

	   Authorization server id

	   Default: "default"
	*/
	Wid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create intelli trust ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIntelliTrustIDPParams) WithDefaults() *CreateIntelliTrustIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create intelli trust ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIntelliTrustIDPParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := CreateIntelliTrustIDPParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) WithTimeout(timeout time.Duration) *CreateIntelliTrustIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) WithContext(ctx context.Context) *CreateIntelliTrustIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) WithHTTPClient(client *http.Client) *CreateIntelliTrustIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntelliTrustIDP adds the intelliTrustIDP to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) WithIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) *CreateIntelliTrustIDPParams {
	o.SetIntelliTrustIDP(intelliTrustIDP)
	return o
}

// SetIntelliTrustIDP adds the intelliTrustIdP to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) SetIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) {
	o.IntelliTrustIDP = intelliTrustIDP
}

// WithWid adds the wid to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) WithWid(wid string) *CreateIntelliTrustIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the create intelli trust ID p params
func (o *CreateIntelliTrustIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIntelliTrustIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IntelliTrustIDP != nil {
		if err := r.SetBodyParam(o.IntelliTrustIDP); err != nil {
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
