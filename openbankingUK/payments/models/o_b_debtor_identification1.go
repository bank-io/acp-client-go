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

// OBDebtorIdentification1 Set of elements used to identify a person or an organisation.
//
// swagger:model OBDebtorIdentification1
type OBDebtorIdentification1 struct {

	// The account name is the name or names of the account owner(s) represented at an account level, as displayed by the ASPSP's online channels. Note, the account name is not the product name or the nickname of the account.
	// Max Length: 350
	Name string `json:"Name,omitempty"`
}

// Validate validates this o b debtor identification1
func (m *OBDebtorIdentification1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBDebtorIdentification1) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("Name", "body", m.Name, 350); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this o b debtor identification1 based on context it is used
func (m *OBDebtorIdentification1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OBDebtorIdentification1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBDebtorIdentification1) UnmarshalBinary(b []byte) error {
	var res OBDebtorIdentification1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
