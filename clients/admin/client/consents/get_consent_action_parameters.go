// Code generated by go-swagger; DO NOT EDIT.

package consents

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

// NewGetConsentActionParams creates a new GetConsentActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConsentActionParams() *GetConsentActionParams {
	return &GetConsentActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsentActionParamsWithTimeout creates a new GetConsentActionParams object
// with the ability to set a timeout on a request.
func NewGetConsentActionParamsWithTimeout(timeout time.Duration) *GetConsentActionParams {
	return &GetConsentActionParams{
		timeout: timeout,
	}
}

// NewGetConsentActionParamsWithContext creates a new GetConsentActionParams object
// with the ability to set a context for a request.
func NewGetConsentActionParamsWithContext(ctx context.Context) *GetConsentActionParams {
	return &GetConsentActionParams{
		Context: ctx,
	}
}

// NewGetConsentActionParamsWithHTTPClient creates a new GetConsentActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConsentActionParamsWithHTTPClient(client *http.Client) *GetConsentActionParams {
	return &GetConsentActionParams{
		HTTPClient: client,
	}
}

/* GetConsentActionParams contains all the parameters to send to the API endpoint
   for the get consent action operation.

   Typically these are written to a http.Request.
*/
type GetConsentActionParams struct {

	// Action.
	Action string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get consent action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsentActionParams) WithDefaults() *GetConsentActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get consent action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsentActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get consent action params
func (o *GetConsentActionParams) WithTimeout(timeout time.Duration) *GetConsentActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consent action params
func (o *GetConsentActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consent action params
func (o *GetConsentActionParams) WithContext(ctx context.Context) *GetConsentActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consent action params
func (o *GetConsentActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consent action params
func (o *GetConsentActionParams) WithHTTPClient(client *http.Client) *GetConsentActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consent action params
func (o *GetConsentActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the get consent action params
func (o *GetConsentActionParams) WithAction(action string) *GetConsentActionParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the get consent action params
func (o *GetConsentActionParams) SetAction(action string) {
	o.Action = action
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsentActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", o.Action); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
