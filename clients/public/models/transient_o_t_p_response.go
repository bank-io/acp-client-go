// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransientOTPResponse transient o t p response
//
// swagger:model TransientOTPResponse
type TransientOTPResponse struct {

	// address
	// Example: 8675409
	// Required: true
	Address string `json:"address"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// expires in
	// Format: duration
	ExpiresIn strfmt.Duration `json:"expires_in,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// mechanism
	// Example: email
	// Required: true
	// Enum: [sms email]
	Mechanism string `json:"mechanism"`

	// tenant id
	// Example: default
	// Required: true
	TenantID string `json:"tenant_id"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this transient o t p response
func (m *TransientOTPResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMechanism(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransientOTPResponse) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *TransientOTPResponse) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransientOTPResponse) validateExpiresIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresIn) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_in", "body", "duration", m.ExpiresIn.String(), formats); err != nil {
		return err
	}

	return nil
}

var transientOTPResponseTypeMechanismPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transientOTPResponseTypeMechanismPropEnum = append(transientOTPResponseTypeMechanismPropEnum, v)
	}
}

const (

	// TransientOTPResponseMechanismSms captures enum value "sms"
	TransientOTPResponseMechanismSms string = "sms"

	// TransientOTPResponseMechanismEmail captures enum value "email"
	TransientOTPResponseMechanismEmail string = "email"
)

// prop value enum
func (m *TransientOTPResponse) validateMechanismEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transientOTPResponseTypeMechanismPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransientOTPResponse) validateMechanism(formats strfmt.Registry) error {

	if err := validate.RequiredString("mechanism", "body", m.Mechanism); err != nil {
		return err
	}

	// value enum
	if err := m.validateMechanismEnum("mechanism", "body", m.Mechanism); err != nil {
		return err
	}

	return nil
}

func (m *TransientOTPResponse) validateTenantID(formats strfmt.Registry) error {

	if err := validate.RequiredString("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *TransientOTPResponse) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transient o t p response based on context it is used
func (m *TransientOTPResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransientOTPResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransientOTPResponse) UnmarshalBinary(b []byte) error {
	var res TransientOTPResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}