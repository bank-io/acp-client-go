// Code generated by go-swagger; DO NOT EDIT.

package scopes

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

// NewDeleteScopeParams creates a new DeleteScopeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteScopeParams() *DeleteScopeParams {
	return &DeleteScopeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScopeParamsWithTimeout creates a new DeleteScopeParams object
// with the ability to set a timeout on a request.
func NewDeleteScopeParamsWithTimeout(timeout time.Duration) *DeleteScopeParams {
	return &DeleteScopeParams{
		timeout: timeout,
	}
}

// NewDeleteScopeParamsWithContext creates a new DeleteScopeParams object
// with the ability to set a context for a request.
func NewDeleteScopeParamsWithContext(ctx context.Context) *DeleteScopeParams {
	return &DeleteScopeParams{
		Context: ctx,
	}
}

// NewDeleteScopeParamsWithHTTPClient creates a new DeleteScopeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteScopeParamsWithHTTPClient(client *http.Client) *DeleteScopeParams {
	return &DeleteScopeParams{
		HTTPClient: client,
	}
}

/* DeleteScopeParams contains all the parameters to send to the API endpoint
   for the delete scope operation.

   Typically these are written to a http.Request.
*/
type DeleteScopeParams struct {

	// Scp.
	Scp string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete scope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScopeParams) WithDefaults() *DeleteScopeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete scope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScopeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete scope params
func (o *DeleteScopeParams) WithTimeout(timeout time.Duration) *DeleteScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scope params
func (o *DeleteScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scope params
func (o *DeleteScopeParams) WithContext(ctx context.Context) *DeleteScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scope params
func (o *DeleteScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scope params
func (o *DeleteScopeParams) WithHTTPClient(client *http.Client) *DeleteScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scope params
func (o *DeleteScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScp adds the scp to the delete scope params
func (o *DeleteScopeParams) WithScp(scp string) *DeleteScopeParams {
	o.SetScp(scp)
	return o
}

// SetScp adds the scp to the delete scope params
func (o *DeleteScopeParams) SetScp(scp string) {
	o.Scp = scp
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scp
	if err := r.SetPathParam("scp", o.Scp); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
