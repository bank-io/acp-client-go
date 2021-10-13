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

// OBCurrencyExchange5 Set of elements used to provide details on the currency exchange.
//
// swagger:model OBCurrencyExchange5
type OBCurrencyExchange5 struct {

	// Unique identification to unambiguously identify the foreign exchange contract.
	// Max Length: 35
	// Min Length: 1
	ContractIdentification string `json:"ContractIdentification,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	// Required: true
	ExchangeRate float64 `json:"ExchangeRate"`

	// instructed amount
	InstructedAmount OBCurrencyExchange5InstructedAmount `json:"InstructedAmount,omitempty"`

	// Date and time at which an exchange rate is quoted.All dates in the JSON payloads are represented in ISO 8601 date-time format.
	// All date-time fields in responses must include the timezone. An example is below:
	// 2017-04-05T10:43:07+00:00
	// Format: date-time
	QuotationDate strfmt.DateTime `json:"QuotationDate,omitempty"`

	// Currency from which an amount is to be converted in a currency conversion.
	// Required: true
	// Pattern: ^[A-Z]{3,3}$
	SourceCurrency string `json:"SourceCurrency"`

	// Currency into which an amount is to be converted in a currency conversion.
	// Pattern: ^[A-Z]{3,3}$
	TargetCurrency string `json:"TargetCurrency,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	// Pattern: ^[A-Z]{3,3}$
	UnitCurrency string `json:"UnitCurrency,omitempty"`
}

// Validate validates this o b currency exchange5
func (m *OBCurrencyExchange5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstructedAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBCurrencyExchange5) validateContractIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.ContractIdentification) { // not required
		return nil
	}

	if err := validate.MinLength("ContractIdentification", "body", m.ContractIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ContractIdentification", "body", m.ContractIdentification, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateExchangeRate(formats strfmt.Registry) error {

	if err := validate.Required("ExchangeRate", "body", float64(m.ExchangeRate)); err != nil {
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateInstructedAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.InstructedAmount) { // not required
		return nil
	}

	if err := m.InstructedAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InstructedAmount")
		}
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateQuotationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.QuotationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("QuotationDate", "body", "date-time", m.QuotationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateSourceCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("SourceCurrency", "body", m.SourceCurrency); err != nil {
		return err
	}

	if err := validate.Pattern("SourceCurrency", "body", m.SourceCurrency, `^[A-Z]{3,3}$`); err != nil {
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateTargetCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("TargetCurrency", "body", m.TargetCurrency, `^[A-Z]{3,3}$`); err != nil {
		return err
	}

	return nil
}

func (m *OBCurrencyExchange5) validateUnitCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("UnitCurrency", "body", m.UnitCurrency, `^[A-Z]{3,3}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o b currency exchange5 based on the context it is used
func (m *OBCurrencyExchange5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstructedAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBCurrencyExchange5) contextValidateInstructedAmount(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InstructedAmount.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InstructedAmount")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBCurrencyExchange5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBCurrencyExchange5) UnmarshalBinary(b []byte) error {
	var res OBCurrencyExchange5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBCurrencyExchange5InstructedAmount Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
//
// swagger:model OBCurrencyExchange5InstructedAmount
type OBCurrencyExchange5InstructedAmount struct {

	// amount
	// Required: true
	Amount *OBActiveCurrencyAndAmountSimpleType `json:"Amount"`

	// currency
	// Required: true
	Currency *ActiveOrHistoricCurrencyCode1 `json:"Currency"`
}

// Validate validates this o b currency exchange5 instructed amount
func (m *OBCurrencyExchange5InstructedAmount) Validate(formats strfmt.Registry) error {
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

func (m *OBCurrencyExchange5InstructedAmount) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("InstructedAmount"+"."+"Amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Required("InstructedAmount"+"."+"Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount" + "." + "Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBCurrencyExchange5InstructedAmount) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("InstructedAmount"+"."+"Currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.Required("InstructedAmount"+"."+"Currency", "body", m.Currency); err != nil {
		return err
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount" + "." + "Currency")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b currency exchange5 instructed amount based on the context it is used
func (m *OBCurrencyExchange5InstructedAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBCurrencyExchange5InstructedAmount) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {
		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount" + "." + "Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBCurrencyExchange5InstructedAmount) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if m.Currency != nil {
		if err := m.Currency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount" + "." + "Currency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBCurrencyExchange5InstructedAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBCurrencyExchange5InstructedAmount) UnmarshalBinary(b []byte) error {
	var res OBCurrencyExchange5InstructedAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
