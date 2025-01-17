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

// NewGetGithubEmbeddedIDPParams creates a new GetGithubEmbeddedIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGithubEmbeddedIDPParams() *GetGithubEmbeddedIDPParams {
	return &GetGithubEmbeddedIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGithubEmbeddedIDPParamsWithTimeout creates a new GetGithubEmbeddedIDPParams object
// with the ability to set a timeout on a request.
func NewGetGithubEmbeddedIDPParamsWithTimeout(timeout time.Duration) *GetGithubEmbeddedIDPParams {
	return &GetGithubEmbeddedIDPParams{
		timeout: timeout,
	}
}

// NewGetGithubEmbeddedIDPParamsWithContext creates a new GetGithubEmbeddedIDPParams object
// with the ability to set a context for a request.
func NewGetGithubEmbeddedIDPParamsWithContext(ctx context.Context) *GetGithubEmbeddedIDPParams {
	return &GetGithubEmbeddedIDPParams{
		Context: ctx,
	}
}

// NewGetGithubEmbeddedIDPParamsWithHTTPClient creates a new GetGithubEmbeddedIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGithubEmbeddedIDPParamsWithHTTPClient(client *http.Client) *GetGithubEmbeddedIDPParams {
	return &GetGithubEmbeddedIDPParams{
		HTTPClient: client,
	}
}

/* GetGithubEmbeddedIDPParams contains all the parameters to send to the API endpoint
   for the get github embedded ID p operation.

   Typically these are written to a http.Request.
*/
type GetGithubEmbeddedIDPParams struct {

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

// WithDefaults hydrates default values in the get github embedded ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGithubEmbeddedIDPParams) WithDefaults() *GetGithubEmbeddedIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get github embedded ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGithubEmbeddedIDPParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := GetGithubEmbeddedIDPParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) WithTimeout(timeout time.Duration) *GetGithubEmbeddedIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) WithContext(ctx context.Context) *GetGithubEmbeddedIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) WithHTTPClient(client *http.Client) *GetGithubEmbeddedIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIid adds the iid to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) WithIid(iid string) *GetGithubEmbeddedIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) WithWid(wid string) *GetGithubEmbeddedIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the get github embedded ID p params
func (o *GetGithubEmbeddedIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGithubEmbeddedIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
