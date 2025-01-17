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

// OBWriteDomesticConsent4DataInitiation OBWriteDomesticConsent4DataInitiation The Initiation payload is sent by the initiating party to the ASPSP. It is used to request movement of funds from the debtor account to a creditor for a single domestic payment.
//
// swagger:model OBWriteDomesticConsent4DataInitiation
type OBWriteDomesticConsent4DataInitiation struct {

	// creditor account
	// Required: true
	CreditorAccount *OBWriteDomesticConsent4DataInitiationCreditorAccount `json:"CreditorAccount"`

	// creditor postal address
	CreditorPostalAddress *OBPostalAddress6 `json:"CreditorPostalAddress,omitempty"`

	// debtor account
	DebtorAccount *OBWriteDomesticConsent4DataInitiationDebtorAccount `json:"DebtorAccount,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	// OB: The Faster Payments Scheme can only access 31 characters for the EndToEndIdentification field.
	// Required: true
	// Max Length: 35
	// Min Length: 1
	EndToEndIdentification string `json:"EndToEndIdentification"`

	// instructed amount
	// Required: true
	InstructedAmount *OBWriteDomesticConsent4DataInitiationInstructedAmount `json:"InstructedAmount"`

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	// Required: true
	// Max Length: 35
	// Min Length: 1
	InstructionIdentification string `json:"InstructionIdentification"`

	// local instrument
	LocalInstrument OBExternalLocalInstrument1Code `json:"LocalInstrument,omitempty"`

	// remittance information
	RemittanceInformation *OBWriteDomesticConsent4DataInitiationRemittanceInformation `json:"RemittanceInformation,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b write domestic consent4 data initiation
func (m *OBWriteDomesticConsent4DataInitiation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorPostalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndToEndIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstructedAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstructionIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemittanceInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateCreditorAccount(formats strfmt.Registry) error {

	if err := validate.Required("CreditorAccount", "body", m.CreditorAccount); err != nil {
		return err
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateCreditorPostalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditorPostalAddress) { // not required
		return nil
	}

	if m.CreditorPostalAddress != nil {
		if err := m.CreditorPostalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorPostalAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorPostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateDebtorAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.DebtorAccount) { // not required
		return nil
	}

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DebtorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DebtorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateEndToEndIdentification(formats strfmt.Registry) error {

	if err := validate.RequiredString("EndToEndIdentification", "body", m.EndToEndIdentification); err != nil {
		return err
	}

	if err := validate.MinLength("EndToEndIdentification", "body", m.EndToEndIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("EndToEndIdentification", "body", m.EndToEndIdentification, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateInstructedAmount(formats strfmt.Registry) error {

	if err := validate.Required("InstructedAmount", "body", m.InstructedAmount); err != nil {
		return err
	}

	if m.InstructedAmount != nil {
		if err := m.InstructedAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InstructedAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateInstructionIdentification(formats strfmt.Registry) error {

	if err := validate.RequiredString("InstructionIdentification", "body", m.InstructionIdentification); err != nil {
		return err
	}

	if err := validate.MinLength("InstructionIdentification", "body", m.InstructionIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("InstructionIdentification", "body", m.InstructionIdentification, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateLocalInstrument(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalInstrument) { // not required
		return nil
	}

	if err := m.LocalInstrument.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LocalInstrument")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LocalInstrument")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) validateRemittanceInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.RemittanceInformation) { // not required
		return nil
	}

	if m.RemittanceInformation != nil {
		if err := m.RemittanceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RemittanceInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RemittanceInformation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b write domestic consent4 data initiation based on the context it is used
func (m *OBWriteDomesticConsent4DataInitiation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorPostalAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebtorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstructedAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalInstrument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemittanceInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateCreditorPostalAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorPostalAddress != nil {
		if err := m.CreditorPostalAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorPostalAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorPostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateDebtorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DebtorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DebtorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateInstructedAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.InstructedAmount != nil {
		if err := m.InstructedAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstructedAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InstructedAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateLocalInstrument(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocalInstrument.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LocalInstrument")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LocalInstrument")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticConsent4DataInitiation) contextValidateRemittanceInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.RemittanceInformation != nil {
		if err := m.RemittanceInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RemittanceInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RemittanceInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticConsent4DataInitiation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticConsent4DataInitiation) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticConsent4DataInitiation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
