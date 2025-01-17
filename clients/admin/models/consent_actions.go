// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsentActions consent actions
//
// swagger:model ConsentActions
type ConsentActions struct {

	// consent actions
	ConsentActions []*ConsentAction `json:"consent_actions"`
}

// Validate validates this consent actions
func (m *ConsentActions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsentActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentActions) validateConsentActions(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsentActions) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsentActions); i++ {
		if swag.IsZero(m.ConsentActions[i]) { // not required
			continue
		}

		if m.ConsentActions[i] != nil {
			if err := m.ConsentActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consent_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consent_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this consent actions based on the context it is used
func (m *ConsentActions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsentActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentActions) contextValidateConsentActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsentActions); i++ {

		if m.ConsentActions[i] != nil {
			if err := m.ConsentActions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consent_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consent_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentActions) UnmarshalBinary(b []byte) error {
	var res ConsentActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
