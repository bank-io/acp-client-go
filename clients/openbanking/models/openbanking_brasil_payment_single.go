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

// OpenbankingBrasilPaymentSingle OpenbankingBrasilPaymentSingle Single
//
// Define a poltica de agendamento nico.
//
// swagger:model OpenbankingBrasilPaymentSingle
type OpenbankingBrasilPaymentSingle struct {

	// Define a data alvo da liquidao do pagamento. O fuso horrio de Brasilia deve ser utilizado para criao e racionalizao sobre os dados deste campo.
	// OBS:Esse campo dever sempre ser no mnimo D+1 corrido, ou seja, a data imediatamente posterior em relao a data do consentimento considerando o fuso horrio de Braslia e dever ser no mximo um ano corrido a partir da data do consentimento considerando o fuso horrio de Braslia.
	// Example: 2021-01-01
	// Required: true
	// Format: date
	Date strfmt.Date `json:"date"`
}

// Validate validates this openbanking brasil payment single
func (m *OpenbankingBrasilPaymentSingle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilPaymentSingle) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.Date(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openbanking brasil payment single based on context it is used
func (m *OpenbankingBrasilPaymentSingle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentSingle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentSingle) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilPaymentSingle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
