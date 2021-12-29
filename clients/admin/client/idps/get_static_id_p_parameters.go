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
)

// NewGetStaticIDPParams creates a new GetStaticIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStaticIDPParams() *GetStaticIDPParams {
	return &GetStaticIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStaticIDPParamsWithTimeout creates a new GetStaticIDPParams object
// with the ability to set a timeout on a request.
func NewGetStaticIDPParamsWithTimeout(timeout time.Duration) *GetStaticIDPParams {
	return &GetStaticIDPParams{
		timeout: timeout,
	}
}

// NewGetStaticIDPParamsWithContext creates a new GetStaticIDPParams object
// with the ability to set a context for a request.
func NewGetStaticIDPParamsWithContext(ctx context.Context) *GetStaticIDPParams {
	return &GetStaticIDPParams{
		Context: ctx,
	}
}

// NewGetStaticIDPParamsWithHTTPClient creates a new GetStaticIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStaticIDPParamsWithHTTPClient(client *http.Client) *GetStaticIDPParams {
	return &GetStaticIDPParams{
		HTTPClient: client,
	}
}

/* GetStaticIDPParams contains all the parameters to send to the API endpoint
   for the get static ID p operation.

   Typically these are written to a http.Request.
*/
type GetStaticIDPParams struct {

	/* Iid.

	   IDP id
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

// WithDefaults hydrates default values in the get static ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStaticIDPParams) WithDefaults() *GetStaticIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get static ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStaticIDPParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := GetStaticIDPParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get static ID p params
func (o *GetStaticIDPParams) WithTimeout(timeout time.Duration) *GetStaticIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get static ID p params
func (o *GetStaticIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get static ID p params
func (o *GetStaticIDPParams) WithContext(ctx context.Context) *GetStaticIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get static ID p params
func (o *GetStaticIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get static ID p params
func (o *GetStaticIDPParams) WithHTTPClient(client *http.Client) *GetStaticIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get static ID p params
func (o *GetStaticIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIid adds the iid to the get static ID p params
func (o *GetStaticIDPParams) WithIid(iid string) *GetStaticIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get static ID p params
func (o *GetStaticIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the get static ID p params
func (o *GetStaticIDPParams) WithWid(wid string) *GetStaticIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the get static ID p params
func (o *GetStaticIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStaticIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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