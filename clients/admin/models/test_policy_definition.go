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

// TestPolicyDefinition Definition of a policy that you wish to test
//
// swagger:model TestPolicyDefinition
type TestPolicyDefinition struct {

	// The definition of an Open Policy Agent (OPA) policy provided using the REGO language.
	Definition string `json:"definition,omitempty"`

	// Language of a policy
	//
	// ACP supports creating Cloudentity policies (using a visual editor or defined using JSON or
	// YAML) and Open Policy Agent (OPA) policies (defined using REGO).
	//
	// OPA is a policy engine that unifies policy enforcement and provides a high-level declarative
	// language (REGO) that lets you specify policies as code. REGO supports referencing nested
	// documents and ensures that your queries are correct and unambiguous.
	// Example: cloudentity
	Language string `json:"language,omitempty"`

	// An array of validators for a Cloudentity policy
	Validators []*ValidatorConfig `json:"validators"`
}

// Validate validates this test policy definition
func (m *TestPolicyDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestPolicyDefinition) validateValidators(formats strfmt.Registry) error {
	if swag.IsZero(m.Validators) { // not required
		return nil
	}

	for i := 0; i < len(m.Validators); i++ {
		if swag.IsZero(m.Validators[i]) { // not required
			continue
		}

		if m.Validators[i] != nil {
			if err := m.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test policy definition based on the context it is used
func (m *TestPolicyDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestPolicyDefinition) contextValidateValidators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Validators); i++ {

		if m.Validators[i] != nil {
			if err := m.Validators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestPolicyDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestPolicyDefinition) UnmarshalBinary(b []byte) error {
	var res TestPolicyDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}