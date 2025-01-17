// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntelliTrustSettings IntelliTrust™ IDP specific settings
//
// swagger:model IntelliTrustSettings
type IntelliTrustSettings struct {

	// OAuth client application identifier from your Entrust Datacard® IntelliTrust™ Authentication
	// Service general settings
	// Example: client
	ClientID string `json:"client_id,omitempty"`

	// String represented domain of the Entrust Datacard® IntelliTrust™ Authentication Service for your organization
	// Example: cloudentity-dev.us.trustedauth.com
	Domain string `json:"domain,omitempty"`

	// If enabled, the groups a user belongs to are collected
	//
	// If you want to fetch groups from the IntelliTrust™ IDP, you need to add the `groups` claim
	// for your application on the IDP side.
	FetchGroups bool `json:"fetch_groups,omitempty"`

	// If enabled, users' data is collected by calling the `userinfo` IntelliTrust™ endpoint.
	GetUserInfo bool `json:"get_user_info,omitempty"`

	// An array of additional scopes your client requests
	// Example: ["email","profile","openid"]
	Scopes []string `json:"scopes"`
}

// Validate validates this intelli trust settings
func (m *IntelliTrustSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this intelli trust settings based on context it is used
func (m *IntelliTrustSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntelliTrustSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntelliTrustSettings) UnmarshalBinary(b []byte) error {
	var res IntelliTrustSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
