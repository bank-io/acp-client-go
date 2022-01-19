// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

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

// NewOpenbankingDomesticStandingOrderConsentIntrospectParams creates a new OpenbankingDomesticStandingOrderConsentIntrospectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenbankingDomesticStandingOrderConsentIntrospectParams() *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	return &OpenbankingDomesticStandingOrderConsentIntrospectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithTimeout creates a new OpenbankingDomesticStandingOrderConsentIntrospectParams object
// with the ability to set a timeout on a request.
func NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithTimeout(timeout time.Duration) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	return &OpenbankingDomesticStandingOrderConsentIntrospectParams{
		timeout: timeout,
	}
}

// NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithContext creates a new OpenbankingDomesticStandingOrderConsentIntrospectParams object
// with the ability to set a context for a request.
func NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithContext(ctx context.Context) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	return &OpenbankingDomesticStandingOrderConsentIntrospectParams{
		Context: ctx,
	}
}

// NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithHTTPClient creates a new OpenbankingDomesticStandingOrderConsentIntrospectParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenbankingDomesticStandingOrderConsentIntrospectParamsWithHTTPClient(client *http.Client) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	return &OpenbankingDomesticStandingOrderConsentIntrospectParams{
		HTTPClient: client,
	}
}

/* OpenbankingDomesticStandingOrderConsentIntrospectParams contains all the parameters to send to the API endpoint
   for the openbanking domestic standing order consent introspect operation.

   Typically these are written to a http.Request.
*/
type OpenbankingDomesticStandingOrderConsentIntrospectParams struct {

	// Token.
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the openbanking domestic standing order consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WithDefaults() *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the openbanking domestic standing order consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WithTimeout(timeout time.Duration) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WithContext(ctx context.Context) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WithHTTPClient(client *http.Client) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WithToken(token *string) *OpenbankingDomesticStandingOrderConsentIntrospectParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the openbanking domestic standing order consent introspect params
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *OpenbankingDomesticStandingOrderConsentIntrospectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
