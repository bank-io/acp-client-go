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

// ConsentPayload consent payload
//
// swagger:model ConsentPayload
type ConsentPayload struct {

	// consent id
	ID string `json:"id,omitempty"`

	// kind
	Kind AuditConsentKind `json:"kind,omitempty"`

	// Type of a consent, specifies the subject of consent, e.g.: domestic_payment
	Type string `json:"type,omitempty"`
}

// Validate validates this consent payload
func (m *ConsentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentPayload) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kind")
		}
		return err
	}

	return nil
}

// ContextValidate validate this consent payload based on the context it is used
func (m *ConsentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentPayload) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Kind.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kind")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentPayload) UnmarshalBinary(b []byte) error {
	var res ConsentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
