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

// NewGetDomesticStandingOrderConsentSystemParams creates a new GetDomesticStandingOrderConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomesticStandingOrderConsentSystemParams() *GetDomesticStandingOrderConsentSystemParams {
	return &GetDomesticStandingOrderConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomesticStandingOrderConsentSystemParamsWithTimeout creates a new GetDomesticStandingOrderConsentSystemParams object
// with the ability to set a timeout on a request.
func NewGetDomesticStandingOrderConsentSystemParamsWithTimeout(timeout time.Duration) *GetDomesticStandingOrderConsentSystemParams {
	return &GetDomesticStandingOrderConsentSystemParams{
		timeout: timeout,
	}
}

// NewGetDomesticStandingOrderConsentSystemParamsWithContext creates a new GetDomesticStandingOrderConsentSystemParams object
// with the ability to set a context for a request.
func NewGetDomesticStandingOrderConsentSystemParamsWithContext(ctx context.Context) *GetDomesticStandingOrderConsentSystemParams {
	return &GetDomesticStandingOrderConsentSystemParams{
		Context: ctx,
	}
}

// NewGetDomesticStandingOrderConsentSystemParamsWithHTTPClient creates a new GetDomesticStandingOrderConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomesticStandingOrderConsentSystemParamsWithHTTPClient(client *http.Client) *GetDomesticStandingOrderConsentSystemParams {
	return &GetDomesticStandingOrderConsentSystemParams{
		HTTPClient: client,
	}
}

/* GetDomesticStandingOrderConsentSystemParams contains all the parameters to send to the API endpoint
   for the get domestic standing order consent system operation.

   Typically these are written to a http.Request.
*/
type GetDomesticStandingOrderConsentSystemParams struct {

	// Login.
	Login string

	// LoginState.
	LoginState *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domestic standing order consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticStandingOrderConsentSystemParams) WithDefaults() *GetDomesticStandingOrderConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domestic standing order consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticStandingOrderConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) WithTimeout(timeout time.Duration) *GetDomesticStandingOrderConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) WithContext(ctx context.Context) *GetDomesticStandingOrderConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) WithHTTPClient(client *http.Client) *GetDomesticStandingOrderConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) WithLogin(login string) *GetDomesticStandingOrderConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WithLoginState adds the loginState to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) WithLoginState(loginState *string) *GetDomesticStandingOrderConsentSystemParams {
	o.SetLoginState(loginState)
	return o
}

// SetLoginState adds the loginState to the get domestic standing order consent system params
func (o *GetDomesticStandingOrderConsentSystemParams) SetLoginState(loginState *string) {
	o.LoginState = loginState
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomesticStandingOrderConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param login
	if err := r.SetPathParam("login", o.Login); err != nil {
		return err
	}

	if o.LoginState != nil {

		// query param login_state
		var qrLoginState string

		if o.LoginState != nil {
			qrLoginState = *o.LoginState
		}
		qLoginState := qrLoginState
		if qLoginState != "" {

			if err := r.SetQueryParam("login_state", qLoginState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
