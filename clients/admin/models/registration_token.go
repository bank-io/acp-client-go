// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistrationToken registration token
//
// swagger:model RegistrationToken
type RegistrationToken struct {

	// client uri
	ClientURI string `json:"client_uri,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`
}

// Validate validates this registration token
func (m *RegistrationToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this registration token based on context it is used
func (m *RegistrationToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationToken) UnmarshalBinary(b []byte) error {
	var res RegistrationToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
