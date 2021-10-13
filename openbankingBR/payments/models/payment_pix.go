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

// PaymentPix PaymentPix
//
// Objeto contendo dados do pagameto como moeda e valor.
//
// swagger:model PaymentPix
type PaymentPix struct {

	// Valor da transao com 2 casas decimais.
	// Example: 100000.12
	// Required: true
	// Max Length: 19
	// Min Length: 4
	// Pattern: ^((\d{1,16}\.\d{2}))$
	Amount string `json:"amount"`

	// Cdigo da moeda nacional segundo modelo ISO-4217, ou seja, 'BRL'.
	// Todos os valores monetrios informados esto representados com a moeda vigente do Brasil.
	// Example: BRL
	// Required: true
	// Max Length: 3
	// Pattern: ^([A-Z]{3})$
	Currency string `json:"currency"`
}

// Validate validates this payment pix
func (m *PaymentPix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentPix) validateAmount(formats strfmt.Registry) error {

	if err := validate.RequiredString("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinLength("amount", "body", m.Amount, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("amount", "body", m.Amount, 19); err != nil {
		return err
	}

	if err := validate.Pattern("amount", "body", m.Amount, `^((\d{1,16}\.\d{2}))$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentPix) validateCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.MaxLength("currency", "body", m.Currency, 3); err != nil {
		return err
	}

	if err := validate.Pattern("currency", "body", m.Currency, `^([A-Z]{3})$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payment pix based on context it is used
func (m *PaymentPix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentPix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentPix) UnmarshalBinary(b []byte) error {
	var res PaymentPix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
