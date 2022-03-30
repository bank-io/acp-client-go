// Code generated by go-swagger; DO NOT EDIT.

package f_d_x

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new f d x API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for f d x API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptFDXConsentSystem(params *AcceptFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptFDXConsentSystemOK, error)

	FdxConsentIntrospect(params *FdxConsentIntrospectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FdxConsentIntrospectOK, error)

	GetFDXConsent(params *GetFDXConsentParams, opts ...ClientOption) (*GetFDXConsentOK, error)

	GetFDXConsentSystem(params *GetFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFDXConsentSystemOK, error)

	RejectFDXConsentSystem(params *RejectFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RejectFDXConsentSystemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcceptFDXConsentSystem accepts f d x consent

  This API can be used by a custom openbanking consent page to notify ACP that user accepted access.
*/
func (a *Client) AcceptFDXConsentSystem(params *AcceptFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptFDXConsentSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptFDXConsentSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "acceptFDXConsentSystem",
		Method:             "POST",
		PathPattern:        "/fdx/fdx/{login}/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptFDXConsentSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptFDXConsentSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acceptFDXConsentSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FdxConsentIntrospect introspects fdx consent

  This endpoint takes an OAuth 2.0 token and, in addition to returning
meta information surrounding the token, returns the fdx consent.
*/
func (a *Client) FdxConsentIntrospect(params *FdxConsentIntrospectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FdxConsentIntrospectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFdxConsentIntrospectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fdxConsentIntrospect",
		Method:             "POST",
		PathPattern:        "/fdx/consents/introspect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FdxConsentIntrospectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FdxConsentIntrospectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fdxConsentIntrospect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFDXConsent gets a consent grant

  Get a Consent Grant
*/
func (a *Client) GetFDXConsent(params *GetFDXConsentParams, opts ...ClientOption) (*GetFDXConsentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFDXConsentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFDXConsent",
		Method:             "GET",
		PathPattern:        "/consents/{consentID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFDXConsentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFDXConsentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFDXConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFDXConsentSystem gets f d x consent

  This API can be used by a custom openbanking consent page.
The consent page must first use client credentials flow to create consent.
*/
func (a *Client) GetFDXConsentSystem(params *GetFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFDXConsentSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFDXConsentSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFDXConsentSystem",
		Method:             "GET",
		PathPattern:        "/fdx/fdx/{login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFDXConsentSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFDXConsentSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFDXConsentSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RejectFDXConsentSystem rejects f d x consent

  This API can be used by a custom openbanking consent page to notify ACP that user rejected access.
*/
func (a *Client) RejectFDXConsentSystem(params *RejectFDXConsentSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RejectFDXConsentSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRejectFDXConsentSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rejectFDXConsentSystem",
		Method:             "POST",
		PathPattern:        "/fdx/fdx/{login}/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectFDXConsentSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RejectFDXConsentSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rejectFDXConsentSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
