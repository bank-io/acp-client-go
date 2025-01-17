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

// NewAcceptInternationalStandingOrderConsentSystemParams creates a new AcceptInternationalStandingOrderConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptInternationalStandingOrderConsentSystemParams() *AcceptInternationalStandingOrderConsentSystemParams {
	return &AcceptInternationalStandingOrderConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptInternationalStandingOrderConsentSystemParamsWithTimeout creates a new AcceptInternationalStandingOrderConsentSystemParams object
// with the ability to set a timeout on a request.
func NewAcceptInternationalStandingOrderConsentSystemParamsWithTimeout(timeout time.Duration) *AcceptInternationalStandingOrderConsentSystemParams {
	return &AcceptInternationalStandingOrderConsentSystemParams{
		timeout: timeout,
	}
}

// NewAcceptInternationalStandingOrderConsentSystemParamsWithContext creates a new AcceptInternationalStandingOrderConsentSystemParams object
// with the ability to set a context for a request.
func NewAcceptInternationalStandingOrderConsentSystemParamsWithContext(ctx context.Context) *AcceptInternationalStandingOrderConsentSystemParams {
	return &AcceptInternationalStandingOrderConsentSystemParams{
		Context: ctx,
	}
}

// NewAcceptInternationalStandingOrderConsentSystemParamsWithHTTPClient creates a new AcceptInternationalStandingOrderConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptInternationalStandingOrderConsentSystemParamsWithHTTPClient(client *http.Client) *AcceptInternationalStandingOrderConsentSystemParams {
	return &AcceptInternationalStandingOrderConsentSystemParams{
		HTTPClient: client,
	}
}

/* AcceptInternationalStandingOrderConsentSystemParams contains all the parameters to send to the API endpoint
   for the accept international standing order consent system operation.

   Typically these are written to a http.Request.
*/
type AcceptInternationalStandingOrderConsentSystemParams struct {

	// AcceptConsent.
	AcceptConsent *models.AcceptConsentRequest

	// Login.
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept international standing order consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithDefaults() *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept international standing order consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithTimeout(timeout time.Duration) *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithContext(ctx context.Context) *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithHTTPClient(client *http.Client) *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptConsent adds the acceptConsent to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithAcceptConsent(acceptConsent *models.AcceptConsentRequest) *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetAcceptConsent(acceptConsent)
	return o
}

// SetAcceptConsent adds the acceptConsent to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetAcceptConsent(acceptConsent *models.AcceptConsentRequest) {
	o.AcceptConsent = acceptConsent
}

// WithLogin adds the login to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) WithLogin(login string) *AcceptInternationalStandingOrderConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the accept international standing order consent system params
func (o *AcceptInternationalStandingOrderConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptInternationalStandingOrderConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
