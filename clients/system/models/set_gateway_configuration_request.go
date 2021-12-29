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

// SetGatewayConfigurationRequest set gateway configuration request
//
// swagger:model SetGatewayConfigurationRequest
type SetGatewayConfigurationRequest struct {

	// List of api groups
	APIGroups []*APIGroup `json:"api_groups"`
}

// Validate validates this set gateway configuration request
func (m *SetGatewayConfigurationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetGatewayConfigurationRequest) validateAPIGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.APIGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.APIGroups); i++ {
		if swag.IsZero(m.APIGroups[i]) { // not required
			continue
		}

		if m.APIGroups[i] != nil {
			if err := m.APIGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("api_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("api_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this set gateway configuration request based on the context it is used
func (m *SetGatewayConfigurationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetGatewayConfigurationRequest) contextValidateAPIGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIGroups); i++ {

		if m.APIGroups[i] != nil {
			if err := m.APIGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("api_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("api_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetGatewayConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetGatewayConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res SetGatewayConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}