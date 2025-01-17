// Code generated by go-swagger; DO NOT EDIT.

package openbanking_b_r

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

// NewPatchPaymentConsentParams creates a new PatchPaymentConsentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPaymentConsentParams() *PatchPaymentConsentParams {
	return &PatchPaymentConsentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentConsentParamsWithTimeout creates a new PatchPaymentConsentParams object
// with the ability to set a timeout on a request.
func NewPatchPaymentConsentParamsWithTimeout(timeout time.Duration) *PatchPaymentConsentParams {
	return &PatchPaymentConsentParams{
		timeout: timeout,
	}
}

// NewPatchPaymentConsentParamsWithContext creates a new PatchPaymentConsentParams object
// with the ability to set a context for a request.
func NewPatchPaymentConsentParamsWithContext(ctx context.Context) *PatchPaymentConsentParams {
	return &PatchPaymentConsentParams{
		Context: ctx,
	}
}

// NewPatchPaymentConsentParamsWithHTTPClient creates a new PatchPaymentConsentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPaymentConsentParamsWithHTTPClient(client *http.Client) *PatchPaymentConsentParams {
	return &PatchPaymentConsentParams{
		HTTPClient: client,
	}
}

/* PatchPaymentConsentParams contains all the parameters to send to the API endpoint
   for the patch payment consent operation.

   Typically these are written to a http.Request.
*/
type PatchPaymentConsentParams struct {

	/* Request.

	   Request
	*/
	Request *models.BrazilCustomerPatchPaymentConsentRequest

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

// WithDefaults hydrates default values in the patch payment consent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPaymentConsentParams) WithDefaults() *PatchPaymentConsentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch payment consent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPaymentConsentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch payment consent params
func (o *PatchPaymentConsentParams) WithTimeout(timeout time.Duration) *PatchPaymentConsentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payment consent params
func (o *PatchPaymentConsentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payment consent params
func (o *PatchPaymentConsentParams) WithContext(ctx context.Context) *PatchPaymentConsentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payment consent params
func (o *PatchPaymentConsentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payment consent params
func (o *PatchPaymentConsentParams) WithHTTPClient(client *http.Client) *PatchPaymentConsentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payment consent params
func (o *PatchPaymentConsentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the patch payment consent params
func (o *PatchPaymentConsentParams) WithRequest(request *models.BrazilCustomerPatchPaymentConsentRequest) *PatchPaymentConsentParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the patch payment consent params
func (o *PatchPaymentConsentParams) SetRequest(request *models.BrazilCustomerPatchPaymentConsentRequest) {
	o.Request = request
}

// WithConsentID adds the consentID to the patch payment consent params
func (o *PatchPaymentConsentParams) WithConsentID(consentID string) *PatchPaymentConsentParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the patch payment consent params
func (o *PatchPaymentConsentParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the patch payment consent params
func (o *PatchPaymentConsentParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *PatchPaymentConsentParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the patch payment consent params
func (o *PatchPaymentConsentParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the patch payment consent params
func (o *PatchPaymentConsentParams) WithXFapiAuthDate(xFapiAuthDate *string) *PatchPaymentConsentParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the patch payment consent params
func (o *PatchPaymentConsentParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the patch payment consent params
func (o *PatchPaymentConsentParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *PatchPaymentConsentParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the patch payment consent params
func (o *PatchPaymentConsentParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the patch payment consent params
func (o *PatchPaymentConsentParams) WithXFapiInteractionID(xFapiInteractionID *string) *PatchPaymentConsentParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the patch payment consent params
func (o *PatchPaymentConsentParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXJwsSignature adds the xJwsSignature to the patch payment consent params
func (o *PatchPaymentConsentParams) WithXJwsSignature(xJwsSignature *string) *PatchPaymentConsentParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the patch payment consent params
func (o *PatchPaymentConsentParams) SetXJwsSignature(xJwsSignature *string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentConsentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

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
