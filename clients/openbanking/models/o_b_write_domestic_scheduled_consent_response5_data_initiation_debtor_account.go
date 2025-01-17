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

// OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
//
// swagger:model OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount
type OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount struct {

	// identification
	// Required: true
	Identification *Identification0 `json:"Identification"`

	// The account name is the name or names of the account owner(s) represented at an account level, as displayed by the ASPSP's online channels.
	// Note, the account name is not the product name or the nickname of the account.
	// Max Length: 350
	// Min Length: 1
	Name string `json:"Name,omitempty"`

	// scheme name
	// Required: true
	SchemeName *OBExternalAccountIdentification4Code `json:"SchemeName"`

	// secondary identification
	SecondaryIdentification SecondaryIdentification `json:"SecondaryIdentification,omitempty"`
}

// Validate validates this o b write domestic scheduled consent response5 data initiation debtor account
func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("Identification", "body", m.Identification); err != nil {
		return err
	}

	if err := validate.Required("Identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("Name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Name", "body", m.Name, 350); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) validateSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if m.SchemeName != nil {
		if err := m.SchemeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) validateSecondaryIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := m.SecondaryIdentification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled consent response5 data initiation debtor account based on the context it is used
func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) contextValidateSchemeName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeName != nil {
		if err := m.SchemeName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) contextValidateSecondaryIdentification(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecondaryIdentification.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduledConsentResponse5DataInitiationDebtorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
