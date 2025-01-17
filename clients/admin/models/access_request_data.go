// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccessRequestData access request data
//
// swagger:model AccessRequestData
type AccessRequestData struct {

	// OAuth client application identifier.
	ClientID string `json:"client_id,omitempty"`

	// Human readable name of a client application
	ClientName string `json:"client_name,omitempty"`

	// Stores information if the owner of the client application is a developer.
	CreatedByDeveloper bool `json:"created_by_developer,omitempty"`

	// Stores the information which grant type was selected to perfom a given action.
	// Matches one of allowed OAuth client grant types for a given client.
	GrantType string `json:"grant_type,omitempty"`

	// Stores information if the client application is a public one.
	Public bool `json:"public,omitempty"`

	// Requester IP address obtained from system network socket information.
	RemoteAddr string `json:"remote_addr,omitempty"`

	// ID of the authorization server (workspace) to which an access request is tied.
	ServerID string `json:"server_id,omitempty"`

	// Session id of a given subject. It's uniform across the authentication processes.
	// It can be used as a correlation ID between a different audit events.
	SessionID string `json:"session_id,omitempty"`

	// Identification of the principal that is the subject of authorization.
	// For the authorization grant, the subject typically identifies an authorized accessor for which the access token is being requested.
	// For client authentication, the subject is the client_id of the OAuth client.
	Subject string `json:"subject,omitempty"`

	// Stores information if the client application is a system tenant's application.
	System bool `json:"system,omitempty"`

	// Token endpoint authentication method configured for a client application.
	// Enum: [client_secret_basic client_secret_post client_secret_jwt private_key_jwt self_signed_tls_client_auth tls_client_auth none]
	TokenEndpointAuthnMethod string `json:"token_endpoint_authn_method,omitempty"`

	// Token signature
	TokenSignature string `json:"token_signature,omitempty"`

	// A characteristic string that lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.
	UserAgent string `json:"user_agent,omitempty"`

	// ID of the authorization server (workspace) to which a resource is tied.
	WorkspaceID string `json:"workspace_id,omitempty"`

	// Requester IP address obtained from X-Forwarded-For header.
	XForwardedFor string `json:"x_forwarded_for,omitempty"`

	// Requester IP address obtained from X-Real-IP header.
	XRealIP string `json:"x_real_ip,omitempty"`
}

// Validate validates this access request data
func (m *AccessRequestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokenEndpointAuthnMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accessRequestDataTypeTokenEndpointAuthnMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret_basic","client_secret_post","client_secret_jwt","private_key_jwt","self_signed_tls_client_auth","tls_client_auth","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessRequestDataTypeTokenEndpointAuthnMethodPropEnum = append(accessRequestDataTypeTokenEndpointAuthnMethodPropEnum, v)
	}
}

const (

	// AccessRequestDataTokenEndpointAuthnMethodClientSecretBasic captures enum value "client_secret_basic"
	AccessRequestDataTokenEndpointAuthnMethodClientSecretBasic string = "client_secret_basic"

	// AccessRequestDataTokenEndpointAuthnMethodClientSecretPost captures enum value "client_secret_post"
	AccessRequestDataTokenEndpointAuthnMethodClientSecretPost string = "client_secret_post"

	// AccessRequestDataTokenEndpointAuthnMethodClientSecretJwt captures enum value "client_secret_jwt"
	AccessRequestDataTokenEndpointAuthnMethodClientSecretJwt string = "client_secret_jwt"

	// AccessRequestDataTokenEndpointAuthnMethodPrivateKeyJwt captures enum value "private_key_jwt"
	AccessRequestDataTokenEndpointAuthnMethodPrivateKeyJwt string = "private_key_jwt"

	// AccessRequestDataTokenEndpointAuthnMethodSelfSignedTLSClientAuth captures enum value "self_signed_tls_client_auth"
	AccessRequestDataTokenEndpointAuthnMethodSelfSignedTLSClientAuth string = "self_signed_tls_client_auth"

	// AccessRequestDataTokenEndpointAuthnMethodTLSClientAuth captures enum value "tls_client_auth"
	AccessRequestDataTokenEndpointAuthnMethodTLSClientAuth string = "tls_client_auth"

	// AccessRequestDataTokenEndpointAuthnMethodNone captures enum value "none"
	AccessRequestDataTokenEndpointAuthnMethodNone string = "none"
)

// prop value enum
func (m *AccessRequestData) validateTokenEndpointAuthnMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessRequestDataTypeTokenEndpointAuthnMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessRequestData) validateTokenEndpointAuthnMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthnMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenEndpointAuthnMethodEnum("token_endpoint_authn_method", "body", m.TokenEndpointAuthnMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this access request data based on context it is used
func (m *AccessRequestData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessRequestData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessRequestData) UnmarshalBinary(b []byte) error {
	var res AccessRequestData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
