// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateGatewayRequest update gateway request
//
// swagger:model UpdateGatewayRequest
type UpdateGatewayRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// gateway name
	// Example: Cloudentity Pyron
	Name string `json:"name,omitempty"`

	// token exchange
	TokenExchange *GatewayTokenExchangeSettings `json:"token_exchange,omitempty"`
}

// Validate validates this update gateway request
func (m *UpdateGatewayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokenExchange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateGatewayRequest) validateTokenExchange(formats strfmt.Registry) error {
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

// ContextValidate validate this update gateway request based on the context it is used
func (m *UpdateGatewayRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTokenExchange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateGatewayRequest) contextValidateTokenExchange(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UpdateGatewayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateGatewayRequest) UnmarshalBinary(b []byte) error {
	var res UpdateGatewayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}