// Code generated by go-swagger; DO NOT EDIT.

package servers

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

// NewGetServerForDeveloperParams creates a new GetServerForDeveloperParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerForDeveloperParams() *GetServerForDeveloperParams {
	return &GetServerForDeveloperParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerForDeveloperParamsWithTimeout creates a new GetServerForDeveloperParams object
// with the ability to set a timeout on a request.
func NewGetServerForDeveloperParamsWithTimeout(timeout time.Duration) *GetServerForDeveloperParams {
	return &GetServerForDeveloperParams{
		timeout: timeout,
	}
}

// NewGetServerForDeveloperParamsWithContext creates a new GetServerForDeveloperParams object
// with the ability to set a context for a request.
func NewGetServerForDeveloperParamsWithContext(ctx context.Context) *GetServerForDeveloperParams {
	return &GetServerForDeveloperParams{
		Context: ctx,
	}
}

// NewGetServerForDeveloperParamsWithHTTPClient creates a new GetServerForDeveloperParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerForDeveloperParamsWithHTTPClient(client *http.Client) *GetServerForDeveloperParams {
	return &GetServerForDeveloperParams{
		HTTPClient: client,
	}
}

/* GetServerForDeveloperParams contains all the parameters to send to the API endpoint
   for the get server for developer operation.

   Typically these are written to a http.Request.
*/
type GetServerForDeveloperParams struct {

	/* Rid.

	   Regular server id

	   Default: "default"
	*/
	Rid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server for developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerForDeveloperParams) WithDefaults() *GetServerForDeveloperParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server for developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerForDeveloperParams) SetDefaults() {
	var (
		ridDefault = string("default")
	)

	val := GetServerForDeveloperParams{
		Rid: ridDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get server for developer params
func (o *GetServerForDeveloperParams) WithTimeout(timeout time.Duration) *GetServerForDeveloperParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server for developer params
func (o *GetServerForDeveloperParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server for developer params
func (o *GetServerForDeveloperParams) WithContext(ctx context.Context) *GetServerForDeveloperParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server for developer params
func (o *GetServerForDeveloperParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server for developer params
func (o *GetServerForDeveloperParams) WithHTTPClient(client *http.Client) *GetServerForDeveloperParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server for developer params
func (o *GetServerForDeveloperParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRid adds the rid to the get server for developer params
func (o *GetServerForDeveloperParams) WithRid(rid string) *GetServerForDeveloperParams {
	o.SetRid(rid)
	return o
}

// SetRid adds the rid to the get server for developer params
func (o *GetServerForDeveloperParams) SetRid(rid string) {
	o.Rid = rid
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerForDeveloperParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rid
	if err := r.SetPathParam("rid", o.Rid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
