// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new consents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for consents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConsentsDeleteConsentsConsentID(params *ConsentsDeleteConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsDeleteConsentsConsentIDNoContent, error)

	ConsentsGetConsentsConsentID(params *ConsentsGetConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsGetConsentsConsentIDOK, error)

	ConsentsPostConsents(params *ConsentsPostConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsPostConsentsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConsentsDeleteConsentsConsentID consents delete consents consent Id

  Mtodo para deletar / revogar o consentimento identificado por consentId.
*/
func (a *Client) ConsentsDeleteConsentsConsentID(params *ConsentsDeleteConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsDeleteConsentsConsentIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConsentsDeleteConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "consentsDeleteConsentsConsentId",
		Method:             "DELETE",
		PathPattern:        "/consents/{consentId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConsentsDeleteConsentsConsentIDReader{formats: a.formats},
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
	success, ok := result.(*ConsentsDeleteConsentsConsentIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConsentsDeleteConsentsConsentIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ConsentsGetConsentsConsentID consents get consents consent Id

  Mtodo para obter detalhes do consentimento identificado por consentId.
*/
func (a *Client) ConsentsGetConsentsConsentID(params *ConsentsGetConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsGetConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConsentsGetConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "consentsGetConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/consents/{consentId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConsentsGetConsentsConsentIDReader{formats: a.formats},
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
	success, ok := result.(*ConsentsGetConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConsentsGetConsentsConsentIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ConsentsPostConsents consents post consents

  Mtodo para a criao de um novo consentimento.
*/
func (a *Client) ConsentsPostConsents(params *ConsentsPostConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConsentsPostConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConsentsPostConsentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "consentsPostConsents",
		Method:             "POST",
		PathPattern:        "/consents",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConsentsPostConsentsReader{formats: a.formats},
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
	success, ok := result.(*ConsentsPostConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConsentsPostConsentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
