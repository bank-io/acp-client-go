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

// Policies policies
//
// swagger:model Policies
type Policies struct {

	// list of policies
	Policies []*Policy `json:"policies"`
}

// Validate validates this policies
func (m *Policies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policies) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this policies based on the context it is used
func (m *Policies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policies) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {
			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policies) UnmarshalBinary(b []byte) error {
	var res Policies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
