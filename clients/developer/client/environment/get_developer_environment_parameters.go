// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// NewGetDeveloperEnvironmentParams creates a new GetDeveloperEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeveloperEnvironmentParams() *GetDeveloperEnvironmentParams {
	return &GetDeveloperEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeveloperEnvironmentParamsWithTimeout creates a new GetDeveloperEnvironmentParams object
// with the ability to set a timeout on a request.
func NewGetDeveloperEnvironmentParamsWithTimeout(timeout time.Duration) *GetDeveloperEnvironmentParams {
	return &GetDeveloperEnvironmentParams{
		timeout: timeout,
	}
}

// NewGetDeveloperEnvironmentParamsWithContext creates a new GetDeveloperEnvironmentParams object
// with the ability to set a context for a request.
func NewGetDeveloperEnvironmentParamsWithContext(ctx context.Context) *GetDeveloperEnvironmentParams {
	return &GetDeveloperEnvironmentParams{
		Context: ctx,
	}
}

// NewGetDeveloperEnvironmentParamsWithHTTPClient creates a new GetDeveloperEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeveloperEnvironmentParamsWithHTTPClient(client *http.Client) *GetDeveloperEnvironmentParams {
	return &GetDeveloperEnvironmentParams{
		HTTPClient: client,
	}
}

/* GetDeveloperEnvironmentParams contains all the parameters to send to the API endpoint
   for the get developer environment operation.

   Typically these are written to a http.Request.
*/
type GetDeveloperEnvironmentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get developer environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeveloperEnvironmentParams) WithDefaults() *GetDeveloperEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get developer environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeveloperEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get developer environment params
func (o *GetDeveloperEnvironmentParams) WithTimeout(timeout time.Duration) *GetDeveloperEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get developer environment params
func (o *GetDeveloperEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get developer environment params
func (o *GetDeveloperEnvironmentParams) WithContext(ctx context.Context) *GetDeveloperEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get developer environment params
func (o *GetDeveloperEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get developer environment params
func (o *GetDeveloperEnvironmentParams) WithHTTPClient(client *http.Client) *GetDeveloperEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get developer environment params
func (o *GetDeveloperEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeveloperEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}