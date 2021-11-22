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

// StaticUser static user
//
// swagger:model StaticUser
type StaticUser struct {

	// additional attributes
	AdditionalAttributes AuthenticationContext `json:"additional_attributes,omitempty"`

	// authentication context
	AuthenticationContext AuthenticationContext `json:"authentication_context,omitempty"`

	// User's preferred email.
	Email string `json:"email,omitempty"`

	// If set to true, indicates that the user's email was verfied.
	EmailVerified bool `json:"email_verified,omitempty"`

	// User password.
	// Example: secret
	Password string `json:"password,omitempty"`

	// User's preferred phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// If set to true, indicates that the user's phone number was verfied.
	PhoneNumberVerified bool `json:"phone_number_verified,omitempty"`

	// User login.
	// Example: peter
	Username string `json:"username,omitempty"`
}

// Validate validates this static user
func (m *StaticUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StaticUser) validateAdditionalAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalAttributes) { // not required
		return nil
	}

	if m.AdditionalAttributes != nil {
		if err := m.AdditionalAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additional_attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("additional_attributes")
			}
			return err
		}
	}

	return nil
}

func (m *StaticUser) validateAuthenticationContext(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationContext) { // not required
		return nil
	}

	if m.AuthenticationContext != nil {
		if err := m.AuthenticationContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication_context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication_context")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this static user based on the context it is used
func (m *StaticUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StaticUser) contextValidateAdditionalAttributes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AdditionalAttributes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("additional_attributes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("additional_attributes")
		}
		return err
	}

	return nil
}

func (m *StaticUser) contextValidateAuthenticationContext(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthenticationContext.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication_context")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentication_context")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StaticUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StaticUser) UnmarshalBinary(b []byte) error {
	var res StaticUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
