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

// OpenbankingBrasilPaymentIdentification OpenbankingBrasilPaymentIdentification Identification
//
// Objeto contendo os dados do recebedor (creditor).
//
// swagger:model OpenbankingBrasilPaymentIdentification
type OpenbankingBrasilPaymentIdentification struct {

	// Identificao da pessoa envolvida na transao.
	// Preencher com o CPF ou CNPJ, de acordo com o valor escolhido no campo type.
	// O CPF ser utilizado com 11 nmeros e dever ser informado sem pontos ou traos.
	// O CNPJ ser utilizado com 14 nmeros e dever ser informado sem pontos ou traos.
	// Example: 58764789000137
	// Required: true
	// Max Length: 14
	// Min Length: 11
	// Pattern: ^\d{11}$|^\d{14}$
	CpfCnpj string `json:"cpfCnpj"`

	// Em caso de pessoa natural deve ser informado o nome completo do titular da conta do recebedor.
	// Em caso de pessoa jurdica deve ser informada a razo social ou o nome fantasia da conta do recebedor.
	// Example: Marco Antonio de Brito
	// Required: true
	// Max Length: 140
	// Pattern: [\w\W\s]*
	Name string `json:"name"`

	// person type
	// Required: true
	PersonType *OpenbankingBrasilPaymentEnumPaymentPersonType `json:"personType"`
}

// Validate validates this openbanking brasil payment identification
func (m *OpenbankingBrasilPaymentIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpfCnpj(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilPaymentIdentification) validateCpfCnpj(formats strfmt.Registry) error {

	if err := validate.RequiredString("cpfCnpj", "body", m.CpfCnpj); err != nil {
		return err
	}

	if err := validate.MinLength("cpfCnpj", "body", m.CpfCnpj, 11); err != nil {
		return err
	}

	if err := validate.MaxLength("cpfCnpj", "body", m.CpfCnpj, 14); err != nil {
		return err
	}

	if err := validate.Pattern("cpfCnpj", "body", m.CpfCnpj, `^\d{11}$|^\d{14}$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilPaymentIdentification) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 140); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilPaymentIdentification) validatePersonType(formats strfmt.Registry) error {

	if err := validate.Required("personType", "body", m.PersonType); err != nil {
		return err
	}

	if err := validate.Required("personType", "body", m.PersonType); err != nil {
		return err
	}

	if m.PersonType != nil {
		if err := m.PersonType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("personType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openbanking brasil payment identification based on the context it is used
func (m *OpenbankingBrasilPaymentIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePersonType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilPaymentIdentification) contextValidatePersonType(ctx context.Context, formats strfmt.Registry) error {

	if m.PersonType != nil {
		if err := m.PersonType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("personType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentIdentification) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilPaymentIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
