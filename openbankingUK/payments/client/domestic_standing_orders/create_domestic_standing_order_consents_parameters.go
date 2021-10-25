// Code generated by go-swagger; DO NOT EDIT.

package domestic_standing_orders

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

	"github.com/cloudentity/acp-client-go/openbankingUK/payments/models"
)

// NewCreateDomesticStandingOrderConsentsParams creates a new CreateDomesticStandingOrderConsentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDomesticStandingOrderConsentsParams() *CreateDomesticStandingOrderConsentsParams {
	return &CreateDomesticStandingOrderConsentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDomesticStandingOrderConsentsParamsWithTimeout creates a new CreateDomesticStandingOrderConsentsParams object
// with the ability to set a timeout on a request.
func NewCreateDomesticStandingOrderConsentsParamsWithTimeout(timeout time.Duration) *CreateDomesticStandingOrderConsentsParams {
	return &CreateDomesticStandingOrderConsentsParams{
		timeout: timeout,
	}
}

// NewCreateDomesticStandingOrderConsentsParamsWithContext creates a new CreateDomesticStandingOrderConsentsParams object
// with the ability to set a context for a request.
func NewCreateDomesticStandingOrderConsentsParamsWithContext(ctx context.Context) *CreateDomesticStandingOrderConsentsParams {
	return &CreateDomesticStandingOrderConsentsParams{
		Context: ctx,
	}
}

// NewCreateDomesticStandingOrderConsentsParamsWithHTTPClient creates a new CreateDomesticStandingOrderConsentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDomesticStandingOrderConsentsParamsWithHTTPClient(client *http.Client) *CreateDomesticStandingOrderConsentsParams {
	return &CreateDomesticStandingOrderConsentsParams{
		HTTPClient: client,
	}
}

/* CreateDomesticStandingOrderConsentsParams contains all the parameters to send to the API endpoint
   for the create domestic standing order consents operation.

   Typically these are written to a http.Request.
*/
type CreateDomesticStandingOrderConsentsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* OBWriteDomesticStandingOrderConsent5Param.

	   Default
	*/
	OBWriteDomesticStandingOrderConsent5Param *models.OBWriteDomesticStandingOrderConsent5

	/* XCustomerUserAgent.

	   Indicates the user-agent that the PSU is using.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   An RFC4122 UID used as a correlation id.
	*/
	XFapiInteractionID *string

	/* XIdempotencyKey.

	     Every request will be processed only once per x-idempotency-key.  The
	Idempotency Key will be valid for 24 hours.

	*/
	XIdempotencyKey string

	/* XJwsSignature.

	   A detached JWS signature of the body of the payload.
	*/
	XJwsSignature string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create domestic standing order consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticStandingOrderConsentsParams) WithDefaults() *CreateDomesticStandingOrderConsentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create domestic standing order consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticStandingOrderConsentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithTimeout(timeout time.Duration) *CreateDomesticStandingOrderConsentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithContext(ctx context.Context) *CreateDomesticStandingOrderConsentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithHTTPClient(client *http.Client) *CreateDomesticStandingOrderConsentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithAuthorization(authorization string) *CreateDomesticStandingOrderConsentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithOBWriteDomesticStandingOrderConsent5Param adds the oBWriteDomesticStandingOrderConsent5Param to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithOBWriteDomesticStandingOrderConsent5Param(oBWriteDomesticStandingOrderConsent5Param *models.OBWriteDomesticStandingOrderConsent5) *CreateDomesticStandingOrderConsentsParams {
	o.SetOBWriteDomesticStandingOrderConsent5Param(oBWriteDomesticStandingOrderConsent5Param)
	return o
}

// SetOBWriteDomesticStandingOrderConsent5Param adds the oBWriteDomesticStandingOrderConsent5Param to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetOBWriteDomesticStandingOrderConsent5Param(oBWriteDomesticStandingOrderConsent5Param *models.OBWriteDomesticStandingOrderConsent5) {
	o.OBWriteDomesticStandingOrderConsent5Param = oBWriteDomesticStandingOrderConsent5Param
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXIdempotencyKey adds the xIdempotencyKey to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXIdempotencyKey(xIdempotencyKey string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXIdempotencyKey(xIdempotencyKey)
	return o
}

// SetXIdempotencyKey adds the xIdempotencyKey to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXIdempotencyKey(xIdempotencyKey string) {
	o.XIdempotencyKey = xIdempotencyKey
}

// WithXJwsSignature adds the xJwsSignature to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) WithXJwsSignature(xJwsSignature string) *CreateDomesticStandingOrderConsentsParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the create domestic standing order consents params
func (o *CreateDomesticStandingOrderConsentsParams) SetXJwsSignature(xJwsSignature string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDomesticStandingOrderConsentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.OBWriteDomesticStandingOrderConsent5Param != nil {
		if err := r.SetBodyParam(o.OBWriteDomesticStandingOrderConsent5Param); err != nil {
			return err
		}
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

	// header param x-idempotency-key
	if err := r.SetHeaderParam("x-idempotency-key", o.XIdempotencyKey); err != nil {
		return err
	}

	// header param x-jws-signature
	if err := r.SetHeaderParam("x-jws-signature", o.XJwsSignature); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
