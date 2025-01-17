// Code generated by go-swagger; DO NOT EDIT.

package consents

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

// NewListUserConsentsByActionParams creates a new ListUserConsentsByActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUserConsentsByActionParams() *ListUserConsentsByActionParams {
	return &ListUserConsentsByActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUserConsentsByActionParamsWithTimeout creates a new ListUserConsentsByActionParams object
// with the ability to set a timeout on a request.
func NewListUserConsentsByActionParamsWithTimeout(timeout time.Duration) *ListUserConsentsByActionParams {
	return &ListUserConsentsByActionParams{
		timeout: timeout,
	}
}

// NewListUserConsentsByActionParamsWithContext creates a new ListUserConsentsByActionParams object
// with the ability to set a context for a request.
func NewListUserConsentsByActionParamsWithContext(ctx context.Context) *ListUserConsentsByActionParams {
	return &ListUserConsentsByActionParams{
		Context: ctx,
	}
}

// NewListUserConsentsByActionParamsWithHTTPClient creates a new ListUserConsentsByActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUserConsentsByActionParamsWithHTTPClient(client *http.Client) *ListUserConsentsByActionParams {
	return &ListUserConsentsByActionParams{
		HTTPClient: client,
	}
}

/* ListUserConsentsByActionParams contains all the parameters to send to the API endpoint
   for the list user consents by action operation.

   Typically these are written to a http.Request.
*/
type ListUserConsentsByActionParams struct {

	/* Action.

	   Consent action id

	   Default: "default"
	*/
	Action string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list user consents by action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserConsentsByActionParams) WithDefaults() *ListUserConsentsByActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list user consents by action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserConsentsByActionParams) SetDefaults() {
	var (
		actionDefault = string("default")
	)

	val := ListUserConsentsByActionParams{
		Action: actionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list user consents by action params
func (o *ListUserConsentsByActionParams) WithTimeout(timeout time.Duration) *ListUserConsentsByActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user consents by action params
func (o *ListUserConsentsByActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user consents by action params
func (o *ListUserConsentsByActionParams) WithContext(ctx context.Context) *ListUserConsentsByActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user consents by action params
func (o *ListUserConsentsByActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user consents by action params
func (o *ListUserConsentsByActionParams) WithHTTPClient(client *http.Client) *ListUserConsentsByActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user consents by action params
func (o *ListUserConsentsByActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the list user consents by action params
func (o *ListUserConsentsByActionParams) WithAction(action string) *ListUserConsentsByActionParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the list user consents by action params
func (o *ListUserConsentsByActionParams) SetAction(action string) {
	o.Action = action
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserConsentsByActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", o.Action); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
