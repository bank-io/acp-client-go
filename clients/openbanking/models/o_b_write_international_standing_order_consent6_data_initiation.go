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

// OBWriteInternationalStandingOrderConsent6DataInitiation OBWriteInternationalStandingOrderConsent6DataInitiation The Initiation payload is sent by the initiating party to the ASPSP. It is used to request movement of funds from the debtor account to a creditor for an international standing order.
//
// swagger:model OBWriteInternationalStandingOrderConsent6DataInitiation
type OBWriteInternationalStandingOrderConsent6DataInitiation struct {

	// charge bearer
	ChargeBearer OBChargeBearerType1Code `json:"ChargeBearer,omitempty"`

	// creditor
	Creditor *OBWriteInternationalStandingOrderConsent6DataInitiationCreditor `json:"Creditor,omitempty"`

	// creditor account
	// Required: true
	CreditorAccount *OBWriteInternationalStandingOrderConsent6DataInitiationCreditorAccount `json:"CreditorAccount"`

	// creditor agent
	CreditorAgent *OBWriteInternationalStandingOrderConsent6DataInitiationCreditorAgent `json:"CreditorAgent,omitempty"`

	// Specifies the currency of the to be transferred amount, which is different from the currency of the debtor's account.
	// Required: true
	// Pattern: ^[A-Z]{3,3}$
	CurrencyOfTransfer string `json:"CurrencyOfTransfer"`

	// debtor account
	DebtorAccount *OBWriteInternationalStandingOrderConsent6DataInitiationDebtorAccount `json:"DebtorAccount,omitempty"`

	// Country in which Credit Account is domiciled. Code to identify a country, a dependency, or another area of particular geopolitical interest, on the basis of country names obtained from the United Nations (ISO 3166, Alpha-2 code).
	// Pattern: [A-Z]{2,2}
	DestinationCountryCode string `json:"DestinationCountryCode,omitempty"`

	// Specifies the purpose of an international payment, when there is no corresponding 4 character code available in the ISO20022 list of Purpose Codes.
	// Max Length: 140
	// Min Length: 1
	ExtendedPurpose string `json:"ExtendedPurpose,omitempty"`

	// The date on which the final payment for a Standing Order schedule will be made.All dates in the JSON payloads are represented in ISO 8601 date-time format.
	// All date-time fields in responses must include the timezone. An example is below:
	// 2017-04-05T10:43:07+00:00
	// Format: date-time
	// Format: date-time
	FinalPaymentDateTime strfmt.DateTime `json:"FinalPaymentDateTime,omitempty"`

	// The date on which the first payment for a Standing Order schedule will be made.All dates in the JSON payloads are represented in ISO 8601 date-time format.
	// All date-time fields in responses must include the timezone. An example is below:
	// 2017-04-05T10:43:07+00:00
	// Required: true
	// Format: date-time
	FirstPaymentDateTime strfmt.DateTime `json:"FirstPaymentDateTime"`

	// Individual Definitions:
	// EvryDay - Every day
	// EvryWorkgDay - Every working day
	// IntrvlWkDay - An interval specified in weeks (01 to 09), and the day within the week (01 to 07)
	// WkInMnthDay - A monthly interval, specifying the week of the month (01 to 05) and day within the week (01 to 07)
	// IntrvlMnthDay - An interval specified in months (between 01 to 06, 12, 24), specifying the day within the month (-5 to -1, 1 to 31)
	// QtrDay - Quarterly (either ENGLISH, SCOTTISH, or RECEIVED).
	// ENGLISH = Paid on the 25th March, 24th June, 29th September and 25th December.
	// SCOTTISH = Paid on the 2nd February, 15th May, 1st August and 11th November.
	// RECEIVED = Paid on the 20th March, 19th June, 24th September and 20th December.
	// Individual Patterns:
	// EvryDay (ScheduleCode)
	// EvryWorkgDay (ScheduleCode)
	// IntrvlWkDay:IntervalInWeeks:DayInWeek (ScheduleCode + IntervalInWeeks + DayInWeek)
	// WkInMnthDay:WeekInMonth:DayInWeek (ScheduleCode + WeekInMonth + DayInWeek)
	// IntrvlMnthDay:IntervalInMonths:DayInMonth (ScheduleCode + IntervalInMonths + DayInMonth)
	// QtrDay: + either (ENGLISH, SCOTTISH or RECEIVED) ScheduleCode + QuarterDay
	// The regular expression for this element combines five smaller versions for each permitted pattern. To aid legibility - the components are presented individually here:
	// EvryDay
	// EvryWorkgDay
	// IntrvlWkDay:0[1-9]:0[1-7]
	// WkInMnthDay:0[1-5]:0[1-7]
	// IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01])
	// QtrDay:(ENGLISH|SCOTTISH|RECEIVED)
	// Full Regular Expression:
	// ^(EvryDay)$|^(EvryWorkgDay)$|^(IntrvlWkDay:0[1-9]:0[1-7])$|^(WkInMnthDay:0[1-5]:0[1-7])$|^(IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01]))$|^(QtrDay:(ENGLISH|SCOTTISH|RECEIVED))$
	// Required: true
	// Pattern: ^(EvryDay)$|^(EvryWorkgDay)$|^(IntrvlDay:((0[2-9])|([1-2][0-9])|3[0-1]))$|^(IntrvlWkDay:0[1-9]:0[1-7])$|^(WkInMnthDay:0[1-5]:0[1-7])$|^(IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01]))$|^(QtrDay:(ENGLISH|SCOTTISH|RECEIVED))$
	Frequency string `json:"Frequency"`

	// instructed amount
	// Required: true
	InstructedAmount *OBWriteInternationalStandingOrderConsent6DataInitiationInstructedAmount `json:"InstructedAmount"`

	// Number of the payments that will be made in completing this frequency sequence including any executed since the sequence start date.
	// Max Length: 35
	// Min Length: 1
	NumberOfPayments string `json:"NumberOfPayments,omitempty"`

	// Specifies the external purpose code in the format of character string with a maximum length of 4 characters.
	// The list of valid codes is an external code list published separately.
	// External code sets can be downloaded from www.iso20022.org.
	// Max Length: 4
	// Min Length: 1
	Purpose string `json:"Purpose,omitempty"`

	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction.
	// Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money.
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	// Max Length: 35
	// Min Length: 1
	Reference string `json:"Reference,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b write international standing order consent6 data initiation
func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeBearer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyOfTransfer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtendedPurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstructedAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfPayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateChargeBearer(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeBearer) { // not required
		return nil
	}

	if err := m.ChargeBearer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ChargeBearer")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ChargeBearer")
		}
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateCreditor(formats strfmt.Registry) error {
	if swag.IsZero(m.Creditor) { // not required
		return nil
	}

	if m.Creditor != nil {
		if err := m.Creditor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Creditor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Creditor")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateCreditorAccount(formats strfmt.Registry) error {

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

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateCreditorAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditorAgent) { // not required
		return nil
	}

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateCurrencyOfTransfer(formats strfmt.Registry) error {

	if err := validate.RequiredString("CurrencyOfTransfer", "body", m.CurrencyOfTransfer); err != nil {
		return err
	}

	if err := validate.Pattern("CurrencyOfTransfer", "body", m.CurrencyOfTransfer, `^[A-Z]{3,3}$`); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateDebtorAccount(formats strfmt.Registry) error {
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

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateDestinationCountryCode(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationCountryCode) { // not required
		return nil
	}

	if err := validate.Pattern("DestinationCountryCode", "body", m.DestinationCountryCode, `[A-Z]{2,2}`); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateExtendedPurpose(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtendedPurpose) { // not required
		return nil
	}

	if err := validate.MinLength("ExtendedPurpose", "body", m.ExtendedPurpose, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ExtendedPurpose", "body", m.ExtendedPurpose, 140); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateFinalPaymentDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalPaymentDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FinalPaymentDateTime", "body", "date-time", m.FinalPaymentDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateFirstPaymentDateTime(formats strfmt.Registry) error {

	if err := validate.Required("FirstPaymentDateTime", "body", strfmt.DateTime(m.FirstPaymentDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("FirstPaymentDateTime", "body", "date-time", m.FirstPaymentDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateFrequency(formats strfmt.Registry) error {

	if err := validate.RequiredString("Frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Pattern("Frequency", "body", m.Frequency, `^(EvryDay)$|^(EvryWorkgDay)$|^(IntrvlDay:((0[2-9])|([1-2][0-9])|3[0-1]))$|^(IntrvlWkDay:0[1-9]:0[1-7])$|^(WkInMnthDay:0[1-5]:0[1-7])$|^(IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01]))$|^(QtrDay:(ENGLISH|SCOTTISH|RECEIVED))$`); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateInstructedAmount(formats strfmt.Registry) error {

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

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateNumberOfPayments(formats strfmt.Registry) error {
	if swag.IsZero(m.NumberOfPayments) { // not required
		return nil
	}

	if err := validate.MinLength("NumberOfPayments", "body", m.NumberOfPayments, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("NumberOfPayments", "body", m.NumberOfPayments, 35); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validatePurpose(formats strfmt.Registry) error {
	if swag.IsZero(m.Purpose) { // not required
		return nil
	}

	if err := validate.MinLength("Purpose", "body", m.Purpose, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Purpose", "body", m.Purpose, 4); err != nil {
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := validate.MinLength("Reference", "body", m.Reference, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Reference", "body", m.Reference, 35); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o b write international standing order consent6 data initiation based on the context it is used
func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChargeBearer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebtorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstructedAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateChargeBearer(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ChargeBearer.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ChargeBearer")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ChargeBearer")
		}
		return err
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateCreditor(ctx context.Context, formats strfmt.Registry) error {

	if m.Creditor != nil {
		if err := m.Creditor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Creditor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Creditor")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateCreditorAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateDebtorAccount(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) contextValidateInstructedAmount(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBWriteInternationalStandingOrderConsent6DataInitiation) UnmarshalBinary(b []byte) error {
	var res OBWriteInternationalStandingOrderConsent6DataInitiation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
