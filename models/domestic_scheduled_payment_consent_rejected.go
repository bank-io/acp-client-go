// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomesticScheduledPaymentConsentRejected domestic scheduled payment consent rejected
//
// swagger:model DomesticScheduledPaymentConsentRejected
type DomesticScheduledPaymentConsentRejected struct {

	// url where user should be redirected
	RedirectTo string `json:"redirect_to,omitempty"`
}

// Validate validates this domestic scheduled payment consent rejected
func (m *DomesticScheduledPaymentConsentRejected) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domestic scheduled payment consent rejected based on context it is used
func (m *DomesticScheduledPaymentConsentRejected) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomesticScheduledPaymentConsentRejected) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomesticScheduledPaymentConsentRejected) UnmarshalBinary(b []byte) error {
	var res DomesticScheduledPaymentConsentRejected
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}