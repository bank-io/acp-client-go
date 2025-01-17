// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DefaultSMSSettings default s m s settings
//
// swagger:model DefaultSMSSettings
type DefaultSMSSettings struct {

	// message template
	MessageTemplate string `json:"message_template,omitempty"`

	// source
	Source string `json:"source,omitempty"`
}

// Validate validates this default s m s settings
func (m *DefaultSMSSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this default s m s settings based on context it is used
func (m *DefaultSMSSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DefaultSMSSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultSMSSettings) UnmarshalBinary(b []byte) error {
	var res DefaultSMSSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
