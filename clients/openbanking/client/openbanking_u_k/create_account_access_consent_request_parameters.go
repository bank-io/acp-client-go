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

// NewCreateAccountAccessConsentRequestParams creates a new CreateAccountAccessConsentRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAccountAccessConsentRequestParams() *CreateAccountAccessConsentRequestParams {
	return &CreateAccountAccessConsentRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountAccessConsentRequestParamsWithTimeout creates a new CreateAccountAccessConsentRequestParams object
// with the ability to set a timeout on a request.
func NewCreateAccountAccessConsentRequestParamsWithTimeout(timeout time.Duration) *CreateAccountAccessConsentRequestParams {
	return &CreateAccountAccessConsentRequestParams{
		timeout: timeout,
	}
}

// NewCreateAccountAccessConsentRequestParamsWithContext creates a new CreateAccountAccessConsentRequestParams object
// with the ability to set a context for a request.
func NewCreateAccountAccessConsentRequestParamsWithContext(ctx context.Context) *CreateAccountAccessConsentRequestParams {
	return &CreateAccountAccessConsentRequestParams{
		Context: ctx,
	}
}

// NewCreateAccountAccessConsentRequestParamsWithHTTPClient creates a new CreateAccountAccessConsentRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAccountAccessConsentRequestParamsWithHTTPClient(client *http.Client) *CreateAccountAccessConsentRequestParams {
	return &CreateAccountAccessConsentRequestParams{
		HTTPClient: client,
	}
}

/* CreateAccountAccessConsentRequestParams contains all the parameters to send to the API endpoint
   for the create account access consent request operation.

   Typically these are written to a http.Request.
*/
type CreateAccountAccessConsentRequestParams struct {

	/* Request.

	   Request
	*/
	Request *models.AccountAccessConsentRequest

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

// WithDefaults hydrates default values in the create account access consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountAccessConsentRequestParams) WithDefaults() *CreateAccountAccessConsentRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create account access consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountAccessConsentRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithTimeout(timeout time.Duration) *CreateAccountAccessConsentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithContext(ctx context.Context) *CreateAccountAccessConsentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithHTTPClient(client *http.Client) *CreateAccountAccessConsentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithRequest(request *models.AccountAccessConsentRequest) *CreateAccountAccessConsentRequestParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetRequest(request *models.AccountAccessConsentRequest) {
	o.Request = request
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *CreateAccountAccessConsentRequestParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithXFapiAuthDate(xFapiAuthDate *string) *CreateAccountAccessConsentRequestParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *CreateAccountAccessConsentRequestParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithXFapiInteractionID(xFapiInteractionID *string) *CreateAccountAccessConsentRequestParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXJwsSignature adds the xJwsSignature to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) WithXJwsSignature(xJwsSignature *string) *CreateAccountAccessConsentRequestParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the create account access consent request params
func (o *CreateAccountAccessConsentRequestParams) SetXJwsSignature(xJwsSignature *string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountAccessConsentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
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
