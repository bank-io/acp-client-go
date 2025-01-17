// Code generated by go-swagger; DO NOT EDIT.

package c_d_r

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

// NewCdrConsentIntrospectParams creates a new CdrConsentIntrospectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCdrConsentIntrospectParams() *CdrConsentIntrospectParams {
	return &CdrConsentIntrospectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCdrConsentIntrospectParamsWithTimeout creates a new CdrConsentIntrospectParams object
// with the ability to set a timeout on a request.
func NewCdrConsentIntrospectParamsWithTimeout(timeout time.Duration) *CdrConsentIntrospectParams {
	return &CdrConsentIntrospectParams{
		timeout: timeout,
	}
}

// NewCdrConsentIntrospectParamsWithContext creates a new CdrConsentIntrospectParams object
// with the ability to set a context for a request.
func NewCdrConsentIntrospectParamsWithContext(ctx context.Context) *CdrConsentIntrospectParams {
	return &CdrConsentIntrospectParams{
		Context: ctx,
	}
}

// NewCdrConsentIntrospectParamsWithHTTPClient creates a new CdrConsentIntrospectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCdrConsentIntrospectParamsWithHTTPClient(client *http.Client) *CdrConsentIntrospectParams {
	return &CdrConsentIntrospectParams{
		HTTPClient: client,
	}
}

/* CdrConsentIntrospectParams contains all the parameters to send to the API endpoint
   for the cdr consent introspect operation.

   Typically these are written to a http.Request.
*/
type CdrConsentIntrospectParams struct {

	// Token.
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cdr consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CdrConsentIntrospectParams) WithDefaults() *CdrConsentIntrospectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cdr consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CdrConsentIntrospectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) WithTimeout(timeout time.Duration) *CdrConsentIntrospectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) WithContext(ctx context.Context) *CdrConsentIntrospectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) WithHTTPClient(client *http.Client) *CdrConsentIntrospectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) WithToken(token *string) *CdrConsentIntrospectParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the cdr consent introspect params
func (o *CdrConsentIntrospectParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CdrConsentIntrospectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Token != nil {

		// form param token
		var frToken string
		if o.Token != nil {
			frToken = *o.Token
		}
		fToken := frToken
		if fToken != "" {
			if err := r.SetFormParam("token", fToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
