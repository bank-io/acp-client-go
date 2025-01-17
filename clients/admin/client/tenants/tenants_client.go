// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAdminTenant(params *GetAdminTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminTenantOK, error)

	UpdateAdminTenant(params *UpdateAdminTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdminTenantOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAdminTenant gets tenant

  Get tenant.
*/
func (a *Client) GetAdminTenant(params *GetAdminTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdminTenant",
		Method:             "GET",
		PathPattern:        "/tenant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdminTenantReader{formats: a.formats},
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
	success, ok := result.(*GetAdminTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAdminTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAdminTenant updates tenant

  Update tenant.
*/
func (a *Client) UpdateAdminTenant(params *UpdateAdminTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdminTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdminTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdminTenant",
		Method:             "PUT",
		PathPattern:        "/tenant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdminTenantReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdminTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdminTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
