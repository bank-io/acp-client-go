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

// SMSAuth s m s auth
//
// swagger:model SMSAuth
type SMSAuth struct {

	// The Twilio Auth Token.
	// Required: true
	AuthToken string `json:"auth_token"`

	// The Twilio Account SID.
	// Required: true
	Sid string `json:"sid"`
}

// Validate validates this s m s auth
func (m *SMSAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SMSAuth) validateAuthToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("auth_token", "body", m.AuthToken); err != nil {
		return err
	}

	return nil
}

func (m *SMSAuth) validateSid(formats strfmt.Registry) error {

	if err := validate.RequiredString("sid", "body", m.Sid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s m s auth based on context it is used
func (m *SMSAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SMSAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SMSAuth) UnmarshalBinary(b []byte) error {
	var res SMSAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
