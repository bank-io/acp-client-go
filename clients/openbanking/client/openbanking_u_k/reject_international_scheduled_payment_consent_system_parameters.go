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

// NewRejectInternationalScheduledPaymentConsentSystemParams creates a new RejectInternationalScheduledPaymentConsentSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRejectInternationalScheduledPaymentConsentSystemParams() *RejectInternationalScheduledPaymentConsentSystemParams {
	return &RejectInternationalScheduledPaymentConsentSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRejectInternationalScheduledPaymentConsentSystemParamsWithTimeout creates a new RejectInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a timeout on a request.
func NewRejectInternationalScheduledPaymentConsentSystemParamsWithTimeout(timeout time.Duration) *RejectInternationalScheduledPaymentConsentSystemParams {
	return &RejectInternationalScheduledPaymentConsentSystemParams{
		timeout: timeout,
	}
}

// NewRejectInternationalScheduledPaymentConsentSystemParamsWithContext creates a new RejectInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a context for a request.
func NewRejectInternationalScheduledPaymentConsentSystemParamsWithContext(ctx context.Context) *RejectInternationalScheduledPaymentConsentSystemParams {
	return &RejectInternationalScheduledPaymentConsentSystemParams{
		Context: ctx,
	}
}

// NewRejectInternationalScheduledPaymentConsentSystemParamsWithHTTPClient creates a new RejectInternationalScheduledPaymentConsentSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewRejectInternationalScheduledPaymentConsentSystemParamsWithHTTPClient(client *http.Client) *RejectInternationalScheduledPaymentConsentSystemParams {
	return &RejectInternationalScheduledPaymentConsentSystemParams{
		HTTPClient: client,
	}
}

/* RejectInternationalScheduledPaymentConsentSystemParams contains all the parameters to send to the API endpoint
   for the reject international scheduled payment consent system operation.

   Typically these are written to a http.Request.
*/
type RejectInternationalScheduledPaymentConsentSystemParams struct {

	// RejectConsent.
	RejectConsent *models.RejectConsentRequest

	// Login.
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reject international scheduled payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithDefaults() *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reject international scheduled payment consent system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithTimeout(timeout time.Duration) *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithContext(ctx context.Context) *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithHTTPClient(client *http.Client) *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRejectConsent adds the rejectConsent to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithRejectConsent(rejectConsent *models.RejectConsentRequest) *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetRejectConsent(rejectConsent)
	return o
}

// SetRejectConsent adds the rejectConsent to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetRejectConsent(rejectConsent *models.RejectConsentRequest) {
	o.RejectConsent = rejectConsent
}

// WithLogin adds the login to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WithLogin(login string) *RejectInternationalScheduledPaymentConsentSystemParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the reject international scheduled payment consent system params
func (o *RejectInternationalScheduledPaymentConsentSystemParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *RejectInternationalScheduledPaymentConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RejectConsent != nil {
		if err := r.SetBodyParam(o.RejectConsent); err != nil {
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
