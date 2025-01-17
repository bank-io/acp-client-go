// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenResponse token response
//
// swagger:model TokenResponse
type TokenResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// cdr arrangement id
	CdrArrangementID string `json:"cdr_arrangement_id,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// id token
	IDToken string `json:"id_token,omitempty"`

	// issued token type
	IssuedTokenType string `json:"issued_token_type,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this token response
func (m *TokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this token response based on context it is used
func (m *TokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenResponse) UnmarshalBinary(b []byte) error {
	var res TokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
