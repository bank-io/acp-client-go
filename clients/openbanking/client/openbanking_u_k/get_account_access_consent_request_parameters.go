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

// NewGetAccountAccessConsentRequestParams creates a new GetAccountAccessConsentRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountAccessConsentRequestParams() *GetAccountAccessConsentRequestParams {
	return &GetAccountAccessConsentRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountAccessConsentRequestParamsWithTimeout creates a new GetAccountAccessConsentRequestParams object
// with the ability to set a timeout on a request.
func NewGetAccountAccessConsentRequestParamsWithTimeout(timeout time.Duration) *GetAccountAccessConsentRequestParams {
	return &GetAccountAccessConsentRequestParams{
		timeout: timeout,
	}
}

// NewGetAccountAccessConsentRequestParamsWithContext creates a new GetAccountAccessConsentRequestParams object
// with the ability to set a context for a request.
func NewGetAccountAccessConsentRequestParamsWithContext(ctx context.Context) *GetAccountAccessConsentRequestParams {
	return &GetAccountAccessConsentRequestParams{
		Context: ctx,
	}
}

// NewGetAccountAccessConsentRequestParamsWithHTTPClient creates a new GetAccountAccessConsentRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountAccessConsentRequestParamsWithHTTPClient(client *http.Client) *GetAccountAccessConsentRequestParams {
	return &GetAccountAccessConsentRequestParams{
		HTTPClient: client,
	}
}

/* GetAccountAccessConsentRequestParams contains all the parameters to send to the API endpoint
   for the get account access consent request operation.

   Typically these are written to a http.Request.
*/
type GetAccountAccessConsentRequestParams struct {

	/* ConsentID.

	   Consent id

	   Format: consentID
	*/
	ConsentID string

	/* XCustomerUserAgent.

	     The header indicates the user-agent that the PSU is using.

	The TPP may populate this field with the user-agent indicated by the PSU.
	If the PSU is using a TPP mobile app, the TPP must ensure that the user-agent string
	is different from browser based user-agent strings.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.

	The value is supplied as a HTTP-date as in section 7.1.1.1 of [RFC7231]
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	     An RFC4122 UID used as a correlation Id.

	If provided, the ASPSP must "play back" this value
	in the x-fapi-interaction-id response header.
	*/
	XFapiInteractionID *string

	/* XJwsSignature.

	     Header containing a detached JWS signature of the body of the payload.

	Refer to resource specific documentation on when this header must be specified.
	*/
	XJwsSignature *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account access consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAccessConsentRequestParams) WithDefaults() *GetAccountAccessConsentRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account access consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAccessConsentRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithTimeout(timeout time.Duration) *GetAccountAccessConsentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithContext(ctx context.Context) *GetAccountAccessConsentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithHTTPClient(client *http.Client) *GetAccountAccessConsentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsentID adds the consentID to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithConsentID(consentID string) *GetAccountAccessConsentRequestParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountAccessConsentRequestParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountAccessConsentRequestParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountAccessConsentRequestParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountAccessConsentRequestParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXJwsSignature adds the xJwsSignature to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) WithXJwsSignature(xJwsSignature *string) *GetAccountAccessConsentRequestParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the get account access consent request params
func (o *GetAccountAccessConsentRequestParams) SetXJwsSignature(xJwsSignature *string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountAccessConsentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param consentID
	if err := r.SetPathParam("consentID", o.ConsentID); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}
	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}
	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}
	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}
	}

	if o.XJwsSignature != nil {

		// header param x-jws-signature
		if err := r.SetHeaderParam("x-jws-signature", *o.XJwsSignature); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
