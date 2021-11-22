// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureB2CCredentials Azure IDP B2C specific credentials
//
// swagger:model AzureB2CCredentials
type AzureB2CCredentials struct {

	// Application secret from your Microsoft Azure application settings
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this azure b2 c credentials
func (m *AzureB2CCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure b2 c credentials based on context it is used
func (m *AzureB2CCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureB2CCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureB2CCredentials) UnmarshalBinary(b []byte) error {
	var res AzureB2CCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
