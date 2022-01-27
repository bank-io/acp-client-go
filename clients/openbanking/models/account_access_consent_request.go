// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAccessConsentRequest account access consent request
//
// swagger:model AccountAccessConsentRequest
type AccountAccessConsentRequest struct {

	// data
	// Required: true
	Data *OBReadConsent1Data `json:"Data"`

	// risk
	// Required: true
	Risk OBRisk2 `json:"Risk"`
}

// Validate validates this account access consent request
func (m *AccountAccessConsentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAccessConsentRequest) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAccessConsentRequest) validateRisk(formats strfmt.Registry) error {

	if m.Risk == nil {
		return errors.Required("Risk", "body", nil)
	}

	return nil
}

// ContextValidate validate this account access consent request based on the context it is used
func (m *AccountAccessConsentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAccessConsentRequest) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAccessConsentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAccessConsentRequest) UnmarshalBinary(b []byte) error {
	var res AccountAccessConsentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}