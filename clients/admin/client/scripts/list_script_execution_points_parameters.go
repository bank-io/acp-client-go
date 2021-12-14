// Code generated by go-swagger; DO NOT EDIT.

package scripts

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

// NewListScriptExecutionPointsParams creates a new ListScriptExecutionPointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListScriptExecutionPointsParams() *ListScriptExecutionPointsParams {
	return &ListScriptExecutionPointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListScriptExecutionPointsParamsWithTimeout creates a new ListScriptExecutionPointsParams object
// with the ability to set a timeout on a request.
func NewListScriptExecutionPointsParamsWithTimeout(timeout time.Duration) *ListScriptExecutionPointsParams {
	return &ListScriptExecutionPointsParams{
		timeout: timeout,
	}
}

// NewListScriptExecutionPointsParamsWithContext creates a new ListScriptExecutionPointsParams object
// with the ability to set a context for a request.
func NewListScriptExecutionPointsParamsWithContext(ctx context.Context) *ListScriptExecutionPointsParams {
	return &ListScriptExecutionPointsParams{
		Context: ctx,
	}
}

// NewListScriptExecutionPointsParamsWithHTTPClient creates a new ListScriptExecutionPointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListScriptExecutionPointsParamsWithHTTPClient(client *http.Client) *ListScriptExecutionPointsParams {
	return &ListScriptExecutionPointsParams{
		HTTPClient: client,
	}
}

/* ListScriptExecutionPointsParams contains all the parameters to send to the API endpoint
   for the list script execution points operation.

   Typically these are written to a http.Request.
*/
type ListScriptExecutionPointsParams struct {

	/* Wid.

	   Authorization server id

	   Default: "default"
	*/
	Wid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list script execution points params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListScriptExecutionPointsParams) WithDefaults() *ListScriptExecutionPointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list script execution points params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListScriptExecutionPointsParams) SetDefaults() {
	var (
		widDefault = string("default")
	)

	val := ListScriptExecutionPointsParams{
		Wid: widDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list script execution points params
func (o *ListScriptExecutionPointsParams) WithTimeout(timeout time.Duration) *ListScriptExecutionPointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list script execution points params
func (o *ListScriptExecutionPointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list script execution points params
func (o *ListScriptExecutionPointsParams) WithContext(ctx context.Context) *ListScriptExecutionPointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list script execution points params
func (o *ListScriptExecutionPointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list script execution points params
func (o *ListScriptExecutionPointsParams) WithHTTPClient(client *http.Client) *ListScriptExecutionPointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list script execution points params
func (o *ListScriptExecutionPointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWid adds the wid to the list script execution points params
func (o *ListScriptExecutionPointsParams) WithWid(wid string) *ListScriptExecutionPointsParams {
	o.SetWid(wid)
	return o
}

// SetWid adds the wid to the list script execution points params
func (o *ListScriptExecutionPointsParams) SetWid(wid string) {
	o.Wid = wid
}

// WriteToRequest writes these params to a swagger request
func (o *ListScriptExecutionPointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param wid
	if err := r.SetPathParam("wid", o.Wid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
