// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Auth0Credentials Auth0 IDP specific credentials
//
// swagger:model Auth0Credentials
type Auth0Credentials struct {

	// OAuth client application secret
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this auth0 credentials
func (m *Auth0Credentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth0 credentials based on context it is used
func (m *Auth0Credentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Auth0Credentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Auth0Credentials) UnmarshalBinary(b []byte) error {
	var res Auth0Credentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
