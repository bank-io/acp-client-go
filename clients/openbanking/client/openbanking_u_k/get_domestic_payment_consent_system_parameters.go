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

// NewGetDomesticPaymentConsentSystemParams creates a new GetDomesticPaymentConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomesticPaymentConsentSystemParams() *GetDomesticPaymentConsentSystemParams {
	return &GetDomesticPaymentConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomesticPaymentConsentSystemParamsWithTimeout creates a new GetDomesticPaymentConsentSystemParams object
// with the ability to set a timeout on a request.
func NewGetDomesticPaymentConsentSystemParamsWithTimeout(timeout time.Duration) *GetDomesticPaymentConsentSystemParams {
	return &GetDomesticPaymentConsentSystemParams{
		timeout: timeout,
	}
}

// NewGetDomesticPaymentConsentSystemParamsWithContext creates a new GetDomesticPaymentConsentSystemParams object
// with the ability to set a context for a request.
func NewGetDomesticPaymentConsentSystemParamsWithContext(ctx context.Context) *GetDomesticPaymentConsentSystemParams {
	return &GetDomesticPaymentConsentSystemParams{
		Context: ctx,
	}
}

// NewGetDomesticPaymentConsentSystemParamsWithHTTPClient creates a new GetDomesticPaymentConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomesticPaymentConsentSystemParamsWithHTTPClient(client *http.Client) *GetDomesticPaymentConsentSystemParams {
	return &GetDomesticPaymentConsentSystemParams{
		HTTPClient: client,
	}
}

/* GetDomesticPaymentConsentSystemParams contains all the parameters to send to the API endpoint
   for the get domestic payment consent system operation.

   Typically these are written to a http.Request.
*/
type GetDomesticPaymentConsentSystemParams struct {

	// Login.
	Login string

	// LoginState.
	LoginState *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domestic payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticPaymentConsentSystemParams) WithDefaults() *GetDomesticPaymentConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domestic payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticPaymentConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) WithTimeout(timeout time.Duration) *GetDomesticPaymentConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) WithContext(ctx context.Context) *GetDomesticPaymentConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) WithHTTPClient(client *http.Client) *GetDomesticPaymentConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) WithLogin(login string) *GetDomesticPaymentConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WithLoginState adds the loginState to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) WithLoginState(loginState *string) *GetDomesticPaymentConsentSystemParams {
	o.SetLoginState(loginState)
	return o
}

// SetLoginState adds the loginState to the get domestic payment consent system params
func (o *GetDomesticPaymentConsentSystemParams) SetLoginState(loginState *string) {
	o.LoginState = loginState
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomesticPaymentConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
