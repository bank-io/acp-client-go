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

// OpenbankingBrasilPaymentData OpenbankingBrasilPaymentData Data
//
// Objeto contendo as informaes de consentimento para a iniciao de pagamento individual.
//
// swagger:model OpenbankingBrasilPaymentData
type OpenbankingBrasilPaymentData struct {

	// business entity
	BusinessEntity *OpenbankingBrasilPaymentBusinessEntity `json:"businessEntity,omitempty"`

	// creditor
	// Required: true
	Creditor *OpenbankingBrasilPaymentIdentification `json:"creditor"`

	// debtor account
	DebtorAccount *OpenbankingBrasilPaymentDebtorAccount `json:"debtorAccount,omitempty"`

	// logged user
	// Required: true
	LoggedUser *OpenbankingBrasilPaymentLoggedUser `json:"loggedUser"`

	// payment
	// Required: true
	Payment *OpenbankingBrasilPaymentPaymentConsent `json:"payment"`
}

// Validate validates this openbanking brasil payment data
func (m *OpenbankingBrasilPaymentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoggedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilPaymentData) validateBusinessEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessEntity) { // not required
		return nil
	}

	if m.BusinessEntity != nil {
		if err := m.BusinessEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessEntity")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) validateCreditor(formats strfmt.Registry) error {

	if err := validate.Required("creditor", "body", m.Creditor); err != nil {
		return err
	}

	if m.Creditor != nil {
		if err := m.Creditor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditor")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) validateDebtorAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.DebtorAccount) { // not required
		return nil
	}

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("debtorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) validateLoggedUser(formats strfmt.Registry) error {

	if err := validate.Required("loggedUser", "body", m.LoggedUser); err != nil {
		return err
	}

	if m.LoggedUser != nil {
		if err := m.LoggedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loggedUser")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) validatePayment(formats strfmt.Registry) error {

	if err := validate.Required("payment", "body", m.Payment); err != nil {
		return err
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openbanking brasil payment data based on the context it is used
func (m *OpenbankingBrasilPaymentData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusinessEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebtorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoggedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilPaymentData) contextValidateBusinessEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.BusinessEntity != nil {
		if err := m.BusinessEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessEntity")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) contextValidateCreditor(ctx context.Context, formats strfmt.Registry) error {

	if m.Creditor != nil {
		if err := m.Creditor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditor")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) contextValidateDebtorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("debtorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) contextValidateLoggedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.LoggedUser != nil {
		if err := m.LoggedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loggedUser")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilPaymentData) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentData) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilPaymentData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
