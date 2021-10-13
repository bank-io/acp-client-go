// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

	"github.com/cloudentity/acp-client-go/acp/models"
)

// NewSetRateLimitParams creates a new SetRateLimitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetRateLimitParams() *SetRateLimitParams {
	return &SetRateLimitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetRateLimitParamsWithTimeout creates a new SetRateLimitParams object
// with the ability to set a timeout on a request.
func NewSetRateLimitParamsWithTimeout(timeout time.Duration) *SetRateLimitParams {
	return &SetRateLimitParams{
		timeout: timeout,
	}
}

// NewSetRateLimitParamsWithContext creates a new SetRateLimitParams object
// with the ability to set a context for a request.
func NewSetRateLimitParamsWithContext(ctx context.Context) *SetRateLimitParams {
	return &SetRateLimitParams{
		Context: ctx,
	}
}

// NewSetRateLimitParamsWithHTTPClient creates a new SetRateLimitParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetRateLimitParamsWithHTTPClient(client *http.Client) *SetRateLimitParams {
	return &SetRateLimitParams{
		HTTPClient: client,
	}
}

/* SetRateLimitParams contains all the parameters to send to the API endpoint
   for the set rate limit operation.

   Typically these are written to a http.Request.
*/
type SetRateLimitParams struct {

	// CustomRateLimit.
	CustomRateLimit *models.SetRateLimitRequest

	// Module.
	Module string

	/* Tid.

	   Tenant id

	   Default: "default"
	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set rate limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRateLimitParams) WithDefaults() *SetRateLimitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set rate limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRateLimitParams) SetDefaults() {
	var (
		tidDefault = string("default")
	)

	val := SetRateLimitParams{
		Tid: tidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the set rate limit params
func (o *SetRateLimitParams) WithTimeout(timeout time.Duration) *SetRateLimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set rate limit params
func (o *SetRateLimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set rate limit params
func (o *SetRateLimitParams) WithContext(ctx context.Context) *SetRateLimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set rate limit params
func (o *SetRateLimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set rate limit params
func (o *SetRateLimitParams) WithHTTPClient(client *http.Client) *SetRateLimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set rate limit params
func (o *SetRateLimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomRateLimit adds the customRateLimit to the set rate limit params
func (o *SetRateLimitParams) WithCustomRateLimit(customRateLimit *models.SetRateLimitRequest) *SetRateLimitParams {
	o.SetCustomRateLimit(customRateLimit)
	return o
}

// SetCustomRateLimit adds the customRateLimit to the set rate limit params
func (o *SetRateLimitParams) SetCustomRateLimit(customRateLimit *models.SetRateLimitRequest) {
	o.CustomRateLimit = customRateLimit
}

// WithModule adds the module to the set rate limit params
func (o *SetRateLimitParams) WithModule(module string) *SetRateLimitParams {
	o.SetModule(module)
	return o
}

// SetModule adds the module to the set rate limit params
func (o *SetRateLimitParams) SetModule(module string) {
	o.Module = module
}

// WithTid adds the tid to the set rate limit params
func (o *SetRateLimitParams) WithTid(tid string) *SetRateLimitParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the set rate limit params
func (o *SetRateLimitParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *SetRateLimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CustomRateLimit != nil {
		if err := r.SetBodyParam(o.CustomRateLimit); err != nil {
			return err
		}
	}

	// path param module
	if err := r.SetPathParam("module", o.Module); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
