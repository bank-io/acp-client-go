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

// ScriptExecutionPoint script execution point
//
// swagger:model ScriptExecutionPoint
type ScriptExecutionPoint struct {

	// Optional script ID
	// Example: 1
	ScriptID string `json:"script_id,omitempty"`

	// The ID of your authorization server (workspace)
	// Example: default
	// Required: true
	ServerID string `json:"server_id"`

	// String representation of the target's ID
	// Example: 1
	// Required: true
	TargetFk string `json:"target_fk"`

	// The ID of your tenant
	// Example: default
	// Required: true
	TenantID string `json:"tenant_id"`

	// String representation of the script execution point type
	// Example: post_authn_ctx
	// Required: true
	Type interface{} `json:"type"`
}

// Validate validates this script execution point
func (m *ScriptExecutionPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetFk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptExecutionPoint) validateServerID(formats strfmt.Registry) error {

	if err := validate.RequiredString("server_id", "body", m.ServerID); err != nil {
		return err
	}

	return nil
}

func (m *ScriptExecutionPoint) validateTargetFk(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_fk", "body", m.TargetFk); err != nil {
		return err
	}

	return nil
}

func (m *ScriptExecutionPoint) validateTenantID(formats strfmt.Registry) error {

	if err := validate.RequiredString("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *ScriptExecutionPoint) validateType(formats strfmt.Registry) error {

	if m.Type == nil {
		return errors.Required("type", "body", nil)
	}

	return nil
}

// ContextValidate validates this script execution point based on context it is used
func (m *ScriptExecutionPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScriptExecutionPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptExecutionPoint) UnmarshalBinary(b []byte) error {
	var res ScriptExecutionPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
