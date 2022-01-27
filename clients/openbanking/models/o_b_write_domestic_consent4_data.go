// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBWriteDomesticConsent4Data OBWriteDomesticConsent4Data o b write domestic consent4 data
//
// swagger:model OBWriteDomesticConsent4Data
type OBWriteDomesticConsent4Data struct {

	// authorisation
	Authorisation *OBWriteDomesticConsent4DataAuthorisation `json:"Authorisation,omitempty"`

	// initiation
	// Required: true
	Initiation *OBWriteDomesticConsent4DataInitiation `json:"Initiation"`

	// Specifies to share the refund account details with PISP
	// Enum: [No Yes]
	ReadRefundAccount string `json:"ReadRefundAccount,omitempty"`

	// s c a support data
	SCASupportData *OBWriteDomesticConsent4DataSCASupportData `json:"SCASupportData,omitempty"`
}

// Validate validates this o b write domestic consent4 data
func (m *OBWriteDomesticConsent4Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorisation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadRefundAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSCASupportData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticConsent4Data) validateAuthorisation(formats strfmt.Registry) error {
	if swag.IsZero(m.Authorisation) { // not required
		return nil
	}

	if m.Authorisation != nil {
		if err := m.Authorisation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authorisation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Authorisation")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4Data) validateInitiation(formats strfmt.Registry) error {

	if err := validate.Required("Initiation", "body", m.Initiation); err != nil {
		return err
	}

	if m.Initiation != nil {
		if err := m.Initiation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Initiation")
			}
			return err
		}
	}

	return nil
}

var oBWriteDomesticConsent4DataTypeReadRefundAccountPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["No","Yes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBWriteDomesticConsent4DataTypeReadRefundAccountPropEnum = append(oBWriteDomesticConsent4DataTypeReadRefundAccountPropEnum, v)
	}
}

const (

	// OBWriteDomesticConsent4DataReadRefundAccountNo captures enum value "No"
	OBWriteDomesticConsent4DataReadRefundAccountNo string = "No"

	// OBWriteDomesticConsent4DataReadRefundAccountYes captures enum value "Yes"
	OBWriteDomesticConsent4DataReadRefundAccountYes string = "Yes"
)

// prop value enum
func (m *OBWriteDomesticConsent4Data) validateReadRefundAccountEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oBWriteDomesticConsent4DataTypeReadRefundAccountPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OBWriteDomesticConsent4Data) validateReadRefundAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadRefundAccount) { // not required
		return nil
	}

	// value enum
	if err := m.validateReadRefundAccountEnum("ReadRefundAccount", "body", m.ReadRefundAccount); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticConsent4Data) validateSCASupportData(formats strfmt.Registry) error {
	if swag.IsZero(m.SCASupportData) { // not required
		return nil
	}

	if m.SCASupportData != nil {
		if err := m.SCASupportData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SCASupportData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SCASupportData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b write domestic consent4 data based on the context it is used
func (m *OBWriteDomesticConsent4Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorisation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSCASupportData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticConsent4Data) contextValidateAuthorisation(ctx context.Context, formats strfmt.Registry) error {

	if m.Authorisation != nil {
		if err := m.Authorisation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authorisation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Authorisation")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4Data) contextValidateInitiation(ctx context.Context, formats strfmt.Registry) error {

	if m.Initiation != nil {
		if err := m.Initiation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Initiation")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticConsent4Data) contextValidateSCASupportData(ctx context.Context, formats strfmt.Registry) error {

	if m.SCASupportData != nil {
		if err := m.SCASupportData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SCASupportData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SCASupportData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticConsent4Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticConsent4Data) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticConsent4Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}