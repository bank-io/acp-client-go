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

// NewDeleteIDPParams creates a new DeleteIDPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIDPParams() *DeleteIDPParams {
	return &DeleteIDPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIDPParamsWithTimeout creates a new DeleteIDPParams object
// with the ability to set a timeout on a request.
func NewDeleteIDPParamsWithTimeout(timeout time.Duration) *DeleteIDPParams {
	return &DeleteIDPParams{
		timeout: timeout,
	}
}

// NewDeleteIDPParamsWithContext creates a new DeleteIDPParams object
// with the ability to set a context for a request.
func NewDeleteIDPParamsWithContext(ctx context.Context) *DeleteIDPParams {
	return &DeleteIDPParams{
		Context: ctx,
	}
}

// NewDeleteIDPParamsWithHTTPClient creates a new DeleteIDPParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIDPParamsWithHTTPClient(client *http.Client) *DeleteIDPParams {
	return &DeleteIDPParams{
		HTTPClient: client,
	}
}

/* DeleteIDPParams contains all the parameters to send to the API endpoint
   for the delete ID p operation.

   Typically these are written to a http.Request.
*/
type DeleteIDPParams struct {

	/* Iid.

	   ID of the IDP you wish to delete
	*/
	Iid string

	/* Wid.

	   ID of your authorization server (workspace)

	   Default: "default"
	*/
	Wid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIDPParams) WithDefaults() *DeleteIDPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete ID p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIDPParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := DeleteIDPParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete ID p params
func (o *DeleteIDPParams) WithTimeout(timeout time.Duration) *DeleteIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ID p params
func (o *DeleteIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ID p params
func (o *DeleteIDPParams) WithContext(ctx context.Context) *DeleteIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ID p params
func (o *DeleteIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ID p params
func (o *DeleteIDPParams) WithHTTPClient(client *http.Client) *DeleteIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ID p params
func (o *DeleteIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIid adds the iid to the delete ID p params
func (o *DeleteIDPParams) WithIid(iid string) *DeleteIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the delete ID p params
func (o *DeleteIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithWid adds the wid to the delete ID p params
func (o *DeleteIDPParams) WithWid(wid string) *DeleteIDPParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the delete ID p params
func (o *DeleteIDPParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
