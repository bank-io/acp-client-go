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

// OBWriteDomesticScheduled2 o b write domestic scheduled2
//
// swagger:model OBWriteDomesticScheduled2
type OBWriteDomesticScheduled2 struct {

	// data
	// Required: true
	Data OBWriteDomesticScheduled2Data `json:"Data"`

	// risk
	// Required: true
	Risk *OBRisk1 `json:"Risk"`
}

// Validate validates this o b write domestic scheduled2
func (m *OBWriteDomesticScheduled2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
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

func (m *OBWriteDomesticScheduled2) validateData(formats strfmt.Registry) error {

	if err := m.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2) validateRisk(formats strfmt.Registry) error {

	if err := validate.Required("Risk", "body", m.Risk); err != nil {
		return err
	}

	if m.Risk != nil {
		if err := m.Risk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Risk")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 based on the context it is used
func (m *OBWriteDomesticScheduled2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
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

func (m *OBWriteDomesticScheduled2) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2) contextValidateRisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Risk != nil {
		if err := m.Risk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Risk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2Data o b write domestic scheduled2 data
//
// swagger:model OBWriteDomesticScheduled2Data
type OBWriteDomesticScheduled2Data struct {

	// OB: Unique identification as assigned by the ASPSP to uniquely identify the consent resource.
	// Required: true
	// Max Length: 128
	// Min Length: 1
	ConsentID string `json:"ConsentId"`

	// initiation
	// Required: true
	Initiation OBWriteDomesticScheduled2DataInitiation `json:"Initiation"`
}

// Validate validates this o b write domestic scheduled2 data
func (m *OBWriteDomesticScheduled2Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduled2Data) validateConsentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("Data"+"."+"ConsentId", "body", m.ConsentID); err != nil {
		return err
	}

	if err := validate.MinLength("Data"+"."+"ConsentId", "body", m.ConsentID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"ConsentId", "body", m.ConsentID, 128); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2Data) validateInitiation(formats strfmt.Registry) error {

	if err := m.Initiation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 data based on the context it is used
func (m *OBWriteDomesticScheduled2Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInitiation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduled2Data) contextValidateInitiation(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Initiation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2Data) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2DataInitiation The Initiation payload is sent by the initiating party to the ASPSP. It is used to request movement of funds from the debtor account to a creditor for a single scheduled domestic payment.
//
// swagger:model OBWriteDomesticScheduled2DataInitiation
type OBWriteDomesticScheduled2DataInitiation struct {

	// creditor account
	// Required: true
	CreditorAccount OBWriteDomesticScheduled2DataInitiationCreditorAccount `json:"CreditorAccount"`

	// creditor postal address
	CreditorPostalAddress *OBPostalAddress6 `json:"CreditorPostalAddress,omitempty"`

	// debtor account
	DebtorAccount OBWriteDomesticScheduled2DataInitiationDebtorAccount `json:"DebtorAccount,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	// OB: The Faster Payments Scheme can only access 31 characters for the EndToEndIdentification field.
	// Max Length: 35
	// Min Length: 1
	EndToEndIdentification string `json:"EndToEndIdentification,omitempty"`

	// instructed amount
	// Required: true
	InstructedAmount OBWriteDomesticScheduled2DataInitiationInstructedAmount `json:"InstructedAmount"`

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	// Required: true
	// Max Length: 35
	// Min Length: 1
	InstructionIdentification string `json:"InstructionIdentification"`

	// local instrument
	LocalInstrument OBExternalLocalInstrument1Code `json:"LocalInstrument,omitempty"`

	// remittance information
	RemittanceInformation OBWriteDomesticScheduled2DataInitiationRemittanceInformation `json:"RemittanceInformation,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited.All dates in the JSON payloads are represented in ISO 8601 date-time format.
	// All date-time fields in responses must include the timezone. An example is below:
	// 2017-04-05T10:43:07+00:00
	// Required: true
	// Format: date-time
	RequestedExecutionDateTime strfmt.DateTime `json:"RequestedExecutionDateTime"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b write domestic scheduled2 data initiation
func (m *OBWriteDomesticScheduled2DataInitiation) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRequestedExecutionDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateCreditorAccount(formats strfmt.Registry) error {

	if err := m.CreditorAccount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateCreditorPostalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditorPostalAddress) { // not required
		return nil
	}

	if m.CreditorPostalAddress != nil {
		if err := m.CreditorPostalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorPostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateDebtorAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.DebtorAccount) { // not required
		return nil
	}

	if err := m.DebtorAccount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateEndToEndIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.EndToEndIdentification) { // not required
		return nil
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"EndToEndIdentification", "body", m.EndToEndIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"EndToEndIdentification", "body", m.EndToEndIdentification, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateInstructedAmount(formats strfmt.Registry) error {

	if err := m.InstructedAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateInstructionIdentification(formats strfmt.Registry) error {

	if err := validate.RequiredString("Data"+"."+"Initiation"+"."+"InstructionIdentification", "body", m.InstructionIdentification); err != nil {
		return err
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"InstructionIdentification", "body", m.InstructionIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"InstructionIdentification", "body", m.InstructionIdentification, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateLocalInstrument(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalInstrument) { // not required
		return nil
	}

	if err := m.LocalInstrument.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "LocalInstrument")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateRemittanceInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.RemittanceInformation) { // not required
		return nil
	}

	if err := m.RemittanceInformation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "RemittanceInformation")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) validateRequestedExecutionDateTime(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"RequestedExecutionDateTime", "body", strfmt.DateTime(m.RequestedExecutionDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("Data"+"."+"Initiation"+"."+"RequestedExecutionDateTime", "body", "date-time", m.RequestedExecutionDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 data initiation based on the context it is used
func (m *OBWriteDomesticScheduled2DataInitiation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateCreditorPostalAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorPostalAddress != nil {
		if err := m.CreditorPostalAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorPostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateDebtorAccount(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DebtorAccount.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateInstructedAmount(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InstructedAmount.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateLocalInstrument(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocalInstrument.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "LocalInstrument")
		}
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiation) contextValidateRemittanceInformation(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RemittanceInformation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "RemittanceInformation")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiation) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2DataInitiation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2DataInitiationCreditorAccount Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
//
// swagger:model OBWriteDomesticScheduled2DataInitiationCreditorAccount
type OBWriteDomesticScheduled2DataInitiationCreditorAccount struct {

	// identification
	// Required: true
	Identification *Identification0 `json:"Identification"`

	// The account name is the name or names of the account owner(s) represented at an account level.
	// Note, the account name is not the product name or the nickname of the account.
	// OB: ASPSPs may carry out name validation for Confirmation of Payee, but it is not mandatory.
	// Required: true
	// Max Length: 350
	// Min Length: 1
	Name string `json:"Name"`

	// scheme name
	// Required: true
	SchemeName *OBExternalAccountIdentification4Code `json:"SchemeName"`

	// secondary identification
	SecondaryIdentification SecondaryIdentification `json:"SecondaryIdentification,omitempty"`
}

// Validate validates this o b write domestic scheduled2 data initiation creditor account
func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) Validate(formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"Identification", "body", m.Identification); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"Identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"Name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"Name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"Name", "body", m.Name, 350); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) validateSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"CreditorAccount"+"."+"SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if m.SchemeName != nil {
		if err := m.SchemeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) validateSecondaryIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := m.SecondaryIdentification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "SecondaryIdentification")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 data initiation creditor account based on the context it is used
func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) contextValidateSchemeName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeName != nil {
		if err := m.SchemeName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) contextValidateSecondaryIdentification(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecondaryIdentification.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "CreditorAccount" + "." + "SecondaryIdentification")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationCreditorAccount) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2DataInitiationCreditorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2DataInitiationDebtorAccount Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
//
// swagger:model OBWriteDomesticScheduled2DataInitiationDebtorAccount
type OBWriteDomesticScheduled2DataInitiationDebtorAccount struct {

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

// Validate validates this o b write domestic scheduled2 data initiation debtor account
func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) Validate(formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"Identification", "body", m.Identification); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"Identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"Name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"Name", "body", m.Name, 350); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) validateSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"DebtorAccount"+"."+"SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if m.SchemeName != nil {
		if err := m.SchemeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) validateSecondaryIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := m.SecondaryIdentification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "SecondaryIdentification")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 data initiation debtor account based on the context it is used
func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) contextValidateSchemeName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeName != nil {
		if err := m.SchemeName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) contextValidateSecondaryIdentification(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecondaryIdentification.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Data" + "." + "Initiation" + "." + "DebtorAccount" + "." + "SecondaryIdentification")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationDebtorAccount) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2DataInitiationDebtorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2DataInitiationInstructedAmount Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
// Usage: This amount has to be transported unchanged through the transaction chain.
//
// swagger:model OBWriteDomesticScheduled2DataInitiationInstructedAmount
type OBWriteDomesticScheduled2DataInitiationInstructedAmount struct {

	// amount
	// Required: true
	Amount *OBActiveCurrencyAndAmountSimpleType `json:"Amount"`

	// currency
	// Required: true
	Currency *ActiveOrHistoricCurrencyCode `json:"Currency"`
}

// Validate validates this o b write domestic scheduled2 data initiation instructed amount
func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) Validate(formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"InstructedAmount"+"."+"Amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"InstructedAmount"+"."+"Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount" + "." + "Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("Data"+"."+"Initiation"+"."+"InstructedAmount"+"."+"Currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.Required("Data"+"."+"Initiation"+"."+"InstructedAmount"+"."+"Currency", "body", m.Currency); err != nil {
		return err
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount" + "." + "Currency")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b write domestic scheduled2 data initiation instructed amount based on the context it is used
func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {
		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount" + "." + "Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if m.Currency != nil {
		if err := m.Currency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data" + "." + "Initiation" + "." + "InstructedAmount" + "." + "Currency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationInstructedAmount) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2DataInitiationInstructedAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBWriteDomesticScheduled2DataInitiationRemittanceInformation Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
//
// swagger:model OBWriteDomesticScheduled2DataInitiationRemittanceInformation
type OBWriteDomesticScheduled2DataInitiationRemittanceInformation struct {

	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction.
	// Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money.
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	// OB: The Faster Payments Scheme can only accept 18 characters for the ReferenceInformation field - which is where this ISO field will be mapped.
	// Max Length: 35
	// Min Length: 1
	Reference string `json:"Reference,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	// Max Length: 140
	// Min Length: 1
	Unstructured string `json:"Unstructured,omitempty"`
}

// Validate validates this o b write domestic scheduled2 data initiation remittance information
func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnstructured(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"RemittanceInformation"+"."+"Reference", "body", m.Reference, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"RemittanceInformation"+"."+"Reference", "body", m.Reference, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) validateUnstructured(formats strfmt.Registry) error {
	if swag.IsZero(m.Unstructured) { // not required
		return nil
	}

	if err := validate.MinLength("Data"+"."+"Initiation"+"."+"RemittanceInformation"+"."+"Unstructured", "body", m.Unstructured, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Data"+"."+"Initiation"+"."+"RemittanceInformation"+"."+"Unstructured", "body", m.Unstructured, 140); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this o b write domestic scheduled2 data initiation remittance information based on context it is used
func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteDomesticScheduled2DataInitiationRemittanceInformation) UnmarshalBinary(b []byte) error {
	var res OBWriteDomesticScheduled2DataInitiationRemittanceInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
