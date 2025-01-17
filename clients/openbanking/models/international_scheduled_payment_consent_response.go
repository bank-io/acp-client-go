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

// InternationalScheduledPaymentConsentResponse international scheduled payment consent response
//
// swagger:model InternationalScheduledPaymentConsentResponse
type InternationalScheduledPaymentConsentResponse struct {

	// data
	// Required: true
	Data *OBWriteInternationalScheduledConsentResponse6Data `json:"Data"`

	// links
	Links *Links `json:"Links,omitempty"`

	// meta
	Meta *Meta `json:"Meta,omitempty"`

	// risk
	// Required: true
	Risk *OBRisk1 `json:"Risk"`
}

// Validate validates this international scheduled payment consent response
func (m *InternationalScheduledPaymentConsentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Links")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Meta")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) validateRisk(formats strfmt.Registry) error {

	if err := validate.Required("Risk", "body", m.Risk); err != nil {
		return err
	}

	if m.Risk != nil {
		if err := m.Risk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Risk")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this international scheduled payment consent response based on the context it is used
func (m *InternationalScheduledPaymentConsentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Links")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Meta")
			}
			return err
		}
	}

	return nil
}

func (m *InternationalScheduledPaymentConsentResponse) contextValidateRisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Risk != nil {
		if err := m.Risk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Risk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternationalScheduledPaymentConsentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternationalScheduledPaymentConsentResponse) UnmarshalBinary(b []byte) error {
	var res InternationalScheduledPaymentConsentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
