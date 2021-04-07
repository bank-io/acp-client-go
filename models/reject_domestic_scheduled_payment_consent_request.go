// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RejectDomesticScheduledPaymentConsentRequest reject domestic scheduled payment consent request
//
// swagger:model RejectDomesticScheduledPaymentConsentRequest
type RejectDomesticScheduledPaymentConsentRequest struct {

	// reject error
	Error string `json:"error,omitempty"`

	// error cause
	ErrorCause string `json:"error_cause,omitempty"`

	// reject error description
	ErrorDescription string `json:"error_description,omitempty"`

	// login identifier
	ID string `json:"id,omitempty"`

	// login state
	LoginState string `json:"login_state,omitempty"`

	// reject http status code
	StatusCode int64 `json:"status_code,omitempty"`
}

// Validate validates this reject domestic scheduled payment consent request
func (m *RejectDomesticScheduledPaymentConsentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reject domestic scheduled payment consent request based on context it is used
func (m *RejectDomesticScheduledPaymentConsentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RejectDomesticScheduledPaymentConsentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RejectDomesticScheduledPaymentConsentRequest) UnmarshalBinary(b []byte) error {
	var res RejectDomesticScheduledPaymentConsentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}