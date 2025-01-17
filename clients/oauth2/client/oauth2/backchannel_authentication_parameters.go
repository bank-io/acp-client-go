// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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
	"github.com/go-openapi/swag"
)

// NewBackchannelAuthenticationParams creates a new BackchannelAuthenticationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackchannelAuthenticationParams() *BackchannelAuthenticationParams {
	return &BackchannelAuthenticationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackchannelAuthenticationParamsWithTimeout creates a new BackchannelAuthenticationParams object
// with the ability to set a timeout on a request.
func NewBackchannelAuthenticationParamsWithTimeout(timeout time.Duration) *BackchannelAuthenticationParams {
	return &BackchannelAuthenticationParams{
		timeout: timeout,
	}
}

// NewBackchannelAuthenticationParamsWithContext creates a new BackchannelAuthenticationParams object
// with the ability to set a context for a request.
func NewBackchannelAuthenticationParamsWithContext(ctx context.Context) *BackchannelAuthenticationParams {
	return &BackchannelAuthenticationParams{
		Context: ctx,
	}
}

// NewBackchannelAuthenticationParamsWithHTTPClient creates a new BackchannelAuthenticationParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackchannelAuthenticationParamsWithHTTPClient(client *http.Client) *BackchannelAuthenticationParams {
	return &BackchannelAuthenticationParams{
		HTTPClient: client,
	}
}

/* BackchannelAuthenticationParams contains all the parameters to send to the API endpoint
   for the backchannel authentication operation.

   Typically these are written to a http.Request.
*/
type BackchannelAuthenticationParams struct {

	/* AcrValues.

	     Requested Authentication Context Class Reference values.
	A space-separated string that specifies the acr values that the OpenID Provider is being requested to use for processing this Authentication Request, with the values appearing in order of preference
	*/
	AcrValues *string

	/* BindingMessage.

	   A human-readable identifier or message intended to be displayed on both the consumption device and the authentication device to interlock them together for the transaction by way of a visual cue for the end-user.
	*/
	BindingMessage *string

	/* ClientAssertion.

	   client assertion
	*/
	ClientAssertion *string

	/* ClientAssertionType.

	   client assertion type
	*/
	ClientAssertionType *string

	// ClientID.
	ClientID *string

	/* ClientNotificationToken.

	   It is a bearer token provided by the Client that will be used by the OpenID Provider to authenticate the callback request to the Client.
	*/
	ClientNotificationToken *string

	// ClientSecret.
	ClientSecret *string

	/* IDTokenHint.

	   An ID Token previously issued to the Client by the OpenID Provider being passed back as a hint to identify the end-user for whom authentication is being requested.
	*/
	IDTokenHint *string

	/* LoginHint.

	   A hint to the OpenID Provider regarding the end-user for whom authentication is being requested.
	*/
	LoginHint *string

	/* LoginHintToken.

	   A token containing information identifying the end-user for whom authentication is being requested
	*/
	LoginHintToken *string

	/* Request.

	   signed authentication request
	*/
	Request *string

	/* RequestedExpiry.

	   A positive integer allowing the client to request the expires_in value for the auth_req_id the server will return.

	   Format: int64
	*/
	RequestedExpiry *int64

	/* Scope.

	   The scope of the access request. Must contain openid scope value.
	*/
	Scope *string

	/* UserCode.

	   A secret code, such as a password or pin, that is known only to the user but verifiable by the OP. The code is used to authorize sending an authentication request to the user's authentication device.
	*/
	UserCode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backchannel authentication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackchannelAuthenticationParams) WithDefaults() *BackchannelAuthenticationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backchannel authentication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackchannelAuthenticationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithTimeout(timeout time.Duration) *BackchannelAuthenticationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithContext(ctx context.Context) *BackchannelAuthenticationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithHTTPClient(client *http.Client) *BackchannelAuthenticationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcrValues adds the acrValues to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithAcrValues(acrValues *string) *BackchannelAuthenticationParams {
	o.SetAcrValues(acrValues)
	return o
}

// SetAcrValues adds the acrValues to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetAcrValues(acrValues *string) {
	o.AcrValues = acrValues
}

// WithBindingMessage adds the bindingMessage to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithBindingMessage(bindingMessage *string) *BackchannelAuthenticationParams {
	o.SetBindingMessage(bindingMessage)
	return o
}

// SetBindingMessage adds the bindingMessage to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetBindingMessage(bindingMessage *string) {
	o.BindingMessage = bindingMessage
}

// WithClientAssertion adds the clientAssertion to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithClientAssertion(clientAssertion *string) *BackchannelAuthenticationParams {
	o.SetClientAssertion(clientAssertion)
	return o
}

// SetClientAssertion adds the clientAssertion to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetClientAssertion(clientAssertion *string) {
	o.ClientAssertion = clientAssertion
}

// WithClientAssertionType adds the clientAssertionType to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithClientAssertionType(clientAssertionType *string) *BackchannelAuthenticationParams {
	o.SetClientAssertionType(clientAssertionType)
	return o
}

// SetClientAssertionType adds the clientAssertionType to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetClientAssertionType(clientAssertionType *string) {
	o.ClientAssertionType = clientAssertionType
}

// WithClientID adds the clientID to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithClientID(clientID *string) *BackchannelAuthenticationParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientNotificationToken adds the clientNotificationToken to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithClientNotificationToken(clientNotificationToken *string) *BackchannelAuthenticationParams {
	o.SetClientNotificationToken(clientNotificationToken)
	return o
}

// SetClientNotificationToken adds the clientNotificationToken to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetClientNotificationToken(clientNotificationToken *string) {
	o.ClientNotificationToken = clientNotificationToken
}

// WithClientSecret adds the clientSecret to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithClientSecret(clientSecret *string) *BackchannelAuthenticationParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithIDTokenHint adds the iDTokenHint to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithIDTokenHint(iDTokenHint *string) *BackchannelAuthenticationParams {
	o.SetIDTokenHint(iDTokenHint)
	return o
}

// SetIDTokenHint adds the idTokenHint to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetIDTokenHint(iDTokenHint *string) {
	o.IDTokenHint = iDTokenHint
}

// WithLoginHint adds the loginHint to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithLoginHint(loginHint *string) *BackchannelAuthenticationParams {
	o.SetLoginHint(loginHint)
	return o
}

// SetLoginHint adds the loginHint to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetLoginHint(loginHint *string) {
	o.LoginHint = loginHint
}

// WithLoginHintToken adds the loginHintToken to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithLoginHintToken(loginHintToken *string) *BackchannelAuthenticationParams {
	o.SetLoginHintToken(loginHintToken)
	return o
}

// SetLoginHintToken adds the loginHintToken to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetLoginHintToken(loginHintToken *string) {
	o.LoginHintToken = loginHintToken
}

// WithRequest adds the request to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithRequest(request *string) *BackchannelAuthenticationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetRequest(request *string) {
	o.Request = request
}

// WithRequestedExpiry adds the requestedExpiry to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithRequestedExpiry(requestedExpiry *int64) *BackchannelAuthenticationParams {
	o.SetRequestedExpiry(requestedExpiry)
	return o
}

// SetRequestedExpiry adds the requestedExpiry to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetRequestedExpiry(requestedExpiry *int64) {
	o.RequestedExpiry = requestedExpiry
}

// WithScope adds the scope to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithScope(scope *string) *BackchannelAuthenticationParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithUserCode adds the userCode to the backchannel authentication params
func (o *BackchannelAuthenticationParams) WithUserCode(userCode *string) *BackchannelAuthenticationParams {
	o.SetUserCode(userCode)
	return o
}

// SetUserCode adds the userCode to the backchannel authentication params
func (o *BackchannelAuthenticationParams) SetUserCode(userCode *string) {
	o.UserCode = userCode
}

// WriteToRequest writes these params to a swagger request
func (o *BackchannelAuthenticationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcrValues != nil {

		// form param acr_values
		var frAcrValues string
		if o.AcrValues != nil {
			frAcrValues = *o.AcrValues
		}
		fAcrValues := frAcrValues
		if fAcrValues != "" {
			if err := r.SetFormParam("acr_values", fAcrValues); err != nil {
				return err
			}
		}
	}

	if o.BindingMessage != nil {

		// form param binding_message
		var frBindingMessage string
		if o.BindingMessage != nil {
			frBindingMessage = *o.BindingMessage
		}
		fBindingMessage := frBindingMessage
		if fBindingMessage != "" {
			if err := r.SetFormParam("binding_message", fBindingMessage); err != nil {
				return err
			}
		}
	}

	if o.ClientAssertion != nil {

		// form param client_assertion
		var frClientAssertion string
		if o.ClientAssertion != nil {
			frClientAssertion = *o.ClientAssertion
		}
		fClientAssertion := frClientAssertion
		if fClientAssertion != "" {
			if err := r.SetFormParam("client_assertion", fClientAssertion); err != nil {
				return err
			}
		}
	}

	if o.ClientAssertionType != nil {

		// form param client_assertion_type
		var frClientAssertionType string
		if o.ClientAssertionType != nil {
			frClientAssertionType = *o.ClientAssertionType
		}
		fClientAssertionType := frClientAssertionType
		if fClientAssertionType != "" {
			if err := r.SetFormParam("client_assertion_type", fClientAssertionType); err != nil {
				return err
			}
		}
	}

	if o.ClientID != nil {

		// form param client_id
		var frClientID string
		if o.ClientID != nil {
			frClientID = *o.ClientID
		}
		fClientID := frClientID
		if fClientID != "" {
			if err := r.SetFormParam("client_id", fClientID); err != nil {
				return err
			}
		}
	}

	if o.ClientNotificationToken != nil {

		// form param client_notification_token
		var frClientNotificationToken string
		if o.ClientNotificationToken != nil {
			frClientNotificationToken = *o.ClientNotificationToken
		}
		fClientNotificationToken := frClientNotificationToken
		if fClientNotificationToken != "" {
			if err := r.SetFormParam("client_notification_token", fClientNotificationToken); err != nil {
				return err
			}
		}
	}

	if o.ClientSecret != nil {

		// form param client_secret
		var frClientSecret string
		if o.ClientSecret != nil {
			frClientSecret = *o.ClientSecret
		}
		fClientSecret := frClientSecret
		if fClientSecret != "" {
			if err := r.SetFormParam("client_secret", fClientSecret); err != nil {
				return err
			}
		}
	}

	if o.IDTokenHint != nil {

		// form param id_token_hint
		var frIDTokenHint string
		if o.IDTokenHint != nil {
			frIDTokenHint = *o.IDTokenHint
		}
		fIDTokenHint := frIDTokenHint
		if fIDTokenHint != "" {
			if err := r.SetFormParam("id_token_hint", fIDTokenHint); err != nil {
				return err
			}
		}
	}

	if o.LoginHint != nil {

		// form param login_hint
		var frLoginHint string
		if o.LoginHint != nil {
			frLoginHint = *o.LoginHint
		}
		fLoginHint := frLoginHint
		if fLoginHint != "" {
			if err := r.SetFormParam("login_hint", fLoginHint); err != nil {
				return err
			}
		}
	}

	if o.LoginHintToken != nil {

		// form param login_hint_token
		var frLoginHintToken string
		if o.LoginHintToken != nil {
			frLoginHintToken = *o.LoginHintToken
		}
		fLoginHintToken := frLoginHintToken
		if fLoginHintToken != "" {
			if err := r.SetFormParam("login_hint_token", fLoginHintToken); err != nil {
				return err
			}
		}
	}

	if o.Request != nil {

		// form param request
		var frRequest string
		if o.Request != nil {
			frRequest = *o.Request
		}
		fRequest := frRequest
		if fRequest != "" {
			if err := r.SetFormParam("request", fRequest); err != nil {
				return err
			}
		}
	}

	if o.RequestedExpiry != nil {

		// form param requested_expiry
		var frRequestedExpiry int64
		if o.RequestedExpiry != nil {
			frRequestedExpiry = *o.RequestedExpiry
		}
		fRequestedExpiry := swag.FormatInt64(frRequestedExpiry)
		if fRequestedExpiry != "" {
			if err := r.SetFormParam("requested_expiry", fRequestedExpiry); err != nil {
				return err
			}
		}
	}

	if o.Scope != nil {

		// form param scope
		var frScope string
		if o.Scope != nil {
			frScope = *o.Scope
		}
		fScope := frScope
		if fScope != "" {
			if err := r.SetFormParam("scope", fScope); err != nil {
				return err
			}
		}
	}

	if o.UserCode != nil {

		// form param user_code
		var frUserCode string
		if o.UserCode != nil {
			frUserCode = *o.UserCode
		}
		fUserCode := frUserCode
		if fUserCode != "" {
			if err := r.SetFormParam("user_code", fUserCode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
