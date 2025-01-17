// Code generated by go-swagger; DO NOT EDIT.

package logins

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

// NewGetLoginRequestParams creates a new GetLoginRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoginRequestParams() *GetLoginRequestParams {
	return &GetLoginRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginRequestParamsWithTimeout creates a new GetLoginRequestParams object
// with the ability to set a timeout on a request.
func NewGetLoginRequestParamsWithTimeout(timeout time.Duration) *GetLoginRequestParams {
	return &GetLoginRequestParams{
		timeout: timeout,
	}
}

// NewGetLoginRequestParamsWithContext creates a new GetLoginRequestParams object
// with the ability to set a context for a request.
func NewGetLoginRequestParamsWithContext(ctx context.Context) *GetLoginRequestParams {
	return &GetLoginRequestParams{
		Context: ctx,
	}
}

// NewGetLoginRequestParamsWithHTTPClient creates a new GetLoginRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoginRequestParamsWithHTTPClient(client *http.Client) *GetLoginRequestParams {
	return &GetLoginRequestParams{
		HTTPClient: client,
	}
}

/* GetLoginRequestParams contains all the parameters to send to the API endpoint
   for the get login request operation.

   Typically these are written to a http.Request.
*/
type GetLoginRequestParams struct {

	// Login.
	Login string

	// LoginState.
	LoginState *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get login request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginRequestParams) WithDefaults() *GetLoginRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get login request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get login request params
func (o *GetLoginRequestParams) WithTimeout(timeout time.Duration) *GetLoginRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login request params
func (o *GetLoginRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login request params
func (o *GetLoginRequestParams) WithContext(ctx context.Context) *GetLoginRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login request params
func (o *GetLoginRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get login request params
func (o *GetLoginRequestParams) WithHTTPClient(client *http.Client) *GetLoginRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get login request params
func (o *GetLoginRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the get login request params
func (o *GetLoginRequestParams) WithLogin(login string) *GetLoginRequestParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get login request params
func (o *GetLoginRequestParams) SetLogin(login string) {
	o.Login = login
}

// WithLoginState adds the loginState to the get login request params
func (o *GetLoginRequestParams) WithLoginState(loginState *string) *GetLoginRequestParams {
	o.SetLoginState(loginState)
	return o
}

// SetLoginState adds the loginState to the get login request params
func (o *GetLoginRequestParams) SetLoginState(loginState *string) {
	o.LoginState = loginState
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
