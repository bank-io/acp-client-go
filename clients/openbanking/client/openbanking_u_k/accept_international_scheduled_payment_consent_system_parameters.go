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

// NewAcceptInternationalScheduledPaymentConsentSystemParams creates a new AcceptInternationalScheduledPaymentConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptInternationalScheduledPaymentConsentSystemParams() *AcceptInternationalScheduledPaymentConsentSystemParams {
	return &AcceptInternationalScheduledPaymentConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptInternationalScheduledPaymentConsentSystemParamsWithTimeout creates a new AcceptInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a timeout on a request.
func NewAcceptInternationalScheduledPaymentConsentSystemParamsWithTimeout(timeout time.Duration) *AcceptInternationalScheduledPaymentConsentSystemParams {
	return &AcceptInternationalScheduledPaymentConsentSystemParams{
		timeout: timeout,
	}
}

// NewAcceptInternationalScheduledPaymentConsentSystemParamsWithContext creates a new AcceptInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a context for a request.
func NewAcceptInternationalScheduledPaymentConsentSystemParamsWithContext(ctx context.Context) *AcceptInternationalScheduledPaymentConsentSystemParams {
	return &AcceptInternationalScheduledPaymentConsentSystemParams{
		Context: ctx,
	}
}

// NewAcceptInternationalScheduledPaymentConsentSystemParamsWithHTTPClient creates a new AcceptInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptInternationalScheduledPaymentConsentSystemParamsWithHTTPClient(client *http.Client) *AcceptInternationalScheduledPaymentConsentSystemParams {
	return &AcceptInternationalScheduledPaymentConsentSystemParams{
		HTTPClient: client,
	}
}

/* AcceptInternationalScheduledPaymentConsentSystemParams contains all the parameters to send to the API endpoint
   for the accept international scheduled payment consent system operation.

   Typically these are written to a http.Request.
*/
type AcceptInternationalScheduledPaymentConsentSystemParams struct {

	// AcceptConsent.
	AcceptConsent *models.AcceptConsentRequest

	// Login.
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept international scheduled payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithDefaults() *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept international scheduled payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithTimeout(timeout time.Duration) *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithContext(ctx context.Context) *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithHTTPClient(client *http.Client) *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptConsent adds the acceptConsent to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithAcceptConsent(acceptConsent *models.AcceptConsentRequest) *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetAcceptConsent(acceptConsent)
	return o
}

// SetAcceptConsent adds the acceptConsent to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetAcceptConsent(acceptConsent *models.AcceptConsentRequest) {
	o.AcceptConsent = acceptConsent
}

// WithLogin adds the login to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WithLogin(login string) *AcceptInternationalScheduledPaymentConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the accept international scheduled payment consent system params
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptInternationalScheduledPaymentConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
