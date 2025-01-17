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

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// NewAcceptInternationalPaymentConsentSystemParams creates a new AcceptInternationalPaymentConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptInternationalPaymentConsentSystemParams() *AcceptInternationalPaymentConsentSystemParams {
	return &AcceptInternationalPaymentConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptInternationalPaymentConsentSystemParamsWithTimeout creates a new AcceptInternationalPaymentConsentSystemParams object
// with the ability to set a timeout on a request.
func NewAcceptInternationalPaymentConsentSystemParamsWithTimeout(timeout time.Duration) *AcceptInternationalPaymentConsentSystemParams {
	return &AcceptInternationalPaymentConsentSystemParams{
		timeout: timeout,
	}
}

// NewAcceptInternationalPaymentConsentSystemParamsWithContext creates a new AcceptInternationalPaymentConsentSystemParams object
// with the ability to set a context for a request.
func NewAcceptInternationalPaymentConsentSystemParamsWithContext(ctx context.Context) *AcceptInternationalPaymentConsentSystemParams {
	return &AcceptInternationalPaymentConsentSystemParams{
		Context: ctx,
	}
}

// NewAcceptInternationalPaymentConsentSystemParamsWithHTTPClient creates a new AcceptInternationalPaymentConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptInternationalPaymentConsentSystemParamsWithHTTPClient(client *http.Client) *AcceptInternationalPaymentConsentSystemParams {
	return &AcceptInternationalPaymentConsentSystemParams{
		HTTPClient: client,
	}
}

/* AcceptInternationalPaymentConsentSystemParams contains all the parameters to send to the API endpoint
   for the accept international payment consent system operation.

   Typically these are written to a http.Request.
*/
type AcceptInternationalPaymentConsentSystemParams struct {

	// AcceptConsent.
	AcceptConsent *models.AcceptConsentRequest

	// Login.
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept international payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalPaymentConsentSystemParams) WithDefaults() *AcceptInternationalPaymentConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept international payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalPaymentConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) WithTimeout(timeout time.Duration) *AcceptInternationalPaymentConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) WithContext(ctx context.Context) *AcceptInternationalPaymentConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) WithHTTPClient(client *http.Client) *AcceptInternationalPaymentConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptConsent adds the acceptConsent to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) WithAcceptConsent(acceptConsent *models.AcceptConsentRequest) *AcceptInternationalPaymentConsentSystemParams {
	o.SetAcceptConsent(acceptConsent)
	return o
}

// SetAcceptConsent adds the acceptConsent to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) SetAcceptConsent(acceptConsent *models.AcceptConsentRequest) {
	o.AcceptConsent = acceptConsent
}

// WithLogin adds the login to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) WithLogin(login string) *AcceptInternationalPaymentConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the accept international payment consent system params
func (o *AcceptInternationalPaymentConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptInternationalPaymentConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AcceptConsent != nil {
		if err := r.SetBodyParam(o.AcceptConsent); err != nil {
			return err
		}
	}

	// path param login
	if err := r.SetPathParam("login", o.Login); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
