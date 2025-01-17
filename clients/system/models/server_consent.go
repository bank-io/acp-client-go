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

// ServerConsent server consent
//
// swagger:model ServerConsent
type ServerConsent struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// custom
	Custom *CustomServerConsent `json:"custom,omitempty"`

	// oidc
	Oidc OIDCServerConsent `json:"oidc,omitempty"`

	// openbanking
	Openbanking *OpenbankingServerConsent `json:"openbanking,omitempty"`

	// server id
	ServerID string `json:"server_id,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this server consent
func (m *ServerConsent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenbanking(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConsent) validateCustom(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	if m.Custom != nil {
		if err := m.Custom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *ServerConsent) validateOpenbanking(formats strfmt.Registry) error {
	if swag.IsZero(m.Openbanking) { // not required
		return nil
	}

	if m.Openbanking != nil {
		if err := m.Openbanking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openbanking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openbanking")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this server consent based on the context it is used
func (m *ServerConsent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenbanking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConsent) contextValidateCustom(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom != nil {
		if err := m.Custom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *ServerConsent) contextValidateOpenbanking(ctx context.Context, formats strfmt.Registry) error {

	if m.Openbanking != nil {
		if err := m.Openbanking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openbanking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openbanking")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerConsent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConsent) UnmarshalBinary(b []byte) error {
	var res ServerConsent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
