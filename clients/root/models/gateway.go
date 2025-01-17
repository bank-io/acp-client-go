// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Gateway gateway
//
// swagger:model Gateway
type Gateway struct {

	// authorization server id
	// Example: default
	AuthorizationServerID string `json:"authorization_server_id,omitempty"`

	// id of a client created for this gateway for authentication
	ClientID string `json:"client_id,omitempty"`

	// if true services are created automatically for each new discovered api group
	CreateAndBindServicesAutomatically bool `json:"create_and_bind_services_automatically,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// unique gateway id
	// Example: 1
	ID string `json:"id,omitempty"`

	// last time a client fetched configuration
	// Format: date-time
	LastActive strfmt.DateTime `json:"last_active,omitempty"`

	// gateway name
	// Example: Cloudentity Pyron
	Name string `json:"name,omitempty"`

	// tenant id
	// Example: default
	TenantID string `json:"tenant_id,omitempty"`

	// token exchange
	TokenExchange *GatewayTokenExchangeSettings `json:"token_exchange,omitempty"`

	// Token exchange client id
	TokenExchangeClientID string `json:"token_exchange_client_id,omitempty"`

	// gateway type, one of: pyron, aws
	// Example: pyron
	Type string `json:"type,omitempty"`
}

// Validate validates this gateway
func (m *Gateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenExchange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Gateway) validateLastActive(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActive) { // not required
		return nil
	}

	if err := validate.FormatOf("last_active", "body", "date-time", m.LastActive.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Gateway) validateTokenExchange(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenExchange) { // not required
		return nil
	}

	if m.TokenExchange != nil {
		if err := m.TokenExchange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token_exchange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token_exchange")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gateway based on the context it is used
func (m *Gateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTokenExchange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Gateway) contextValidateTokenExchange(ctx context.Context, formats strfmt.Registry) error {

	if m.TokenExchange != nil {
		if err := m.TokenExchange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token_exchange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token_exchange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Gateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Gateway) UnmarshalBinary(b []byte) error {
	var res Gateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
