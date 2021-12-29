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

// NewListServersForDeveloperParams creates a new ListServersForDeveloperParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServersForDeveloperParams() *ListServersForDeveloperParams {
	return &ListServersForDeveloperParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServersForDeveloperParamsWithTimeout creates a new ListServersForDeveloperParams object
// with the ability to set a timeout on a request.
func NewListServersForDeveloperParamsWithTimeout(timeout time.Duration) *ListServersForDeveloperParams {
	return &ListServersForDeveloperParams{
		timeout: timeout,
	}
}

// NewListServersForDeveloperParamsWithContext creates a new ListServersForDeveloperParams object
// with the ability to set a context for a request.
func NewListServersForDeveloperParamsWithContext(ctx context.Context) *ListServersForDeveloperParams {
	return &ListServersForDeveloperParams{
		Context: ctx,
	}
}

// NewListServersForDeveloperParamsWithHTTPClient creates a new ListServersForDeveloperParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServersForDeveloperParamsWithHTTPClient(client *http.Client) *ListServersForDeveloperParams {
	return &ListServersForDeveloperParams{
		HTTPClient: client,
	}
}

/* ListServersForDeveloperParams contains all the parameters to send to the API endpoint
   for the list servers for developer operation.

   Typically these are written to a http.Request.
*/
type ListServersForDeveloperParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list servers for developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServersForDeveloperParams) WithDefaults() *ListServersForDeveloperParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list servers for developer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServersForDeveloperParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list servers for developer params
func (o *ListServersForDeveloperParams) WithTimeout(timeout time.Duration) *ListServersForDeveloperParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list servers for developer params
func (o *ListServersForDeveloperParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list servers for developer params
func (o *ListServersForDeveloperParams) WithContext(ctx context.Context) *ListServersForDeveloperParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list servers for developer params
func (o *ListServersForDeveloperParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list servers for developer params
func (o *ListServersForDeveloperParams) WithHTTPClient(client *http.Client) *ListServersForDeveloperParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list servers for developer params
func (o *ListServersForDeveloperParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListServersForDeveloperParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}