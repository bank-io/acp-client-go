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

// ListServersDeveloperResponse list servers developer response
//
// swagger:model ListServersDeveloperResponse
type ListServersDeveloperResponse struct {

	// list of authorization servers
	Servers []*GetServerDeveloperResponse `json:"servers"`
}

// Validate validates this list servers developer response
func (m *ListServersDeveloperResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListServersDeveloperResponse) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list servers developer response based on the context it is used
func (m *ListServersDeveloperResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListServersDeveloperResponse) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Servers); i++ {

		if m.Servers[i] != nil {
			if err := m.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListServersDeveloperResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListServersDeveloperResponse) UnmarshalBinary(b []byte) error {
	var res ListServersDeveloperResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
