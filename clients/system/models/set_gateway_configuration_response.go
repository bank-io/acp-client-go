// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetGatewayConfigurationResponse set gateway configuration response
//
// swagger:model SetGatewayConfigurationResponse
type SetGatewayConfigurationResponse struct {

	// number of added api groups
	NumberOfAddedAPIGroups int64 `json:"number_of_added_api_groups,omitempty"`
}

// Validate validates this set gateway configuration response
func (m *SetGatewayConfigurationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set gateway configuration response based on context it is used
func (m *SetGatewayConfigurationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetGatewayConfigurationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetGatewayConfigurationResponse) UnmarshalBinary(b []byte) error {
	var res SetGatewayConfigurationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
