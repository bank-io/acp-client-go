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

// CreatePixPaymentData CreatePixPaymentData
//
// Objeto contendo dados do pagamento e do recebedor (creditor).
//
// swagger:model CreatePixPaymentData
type CreatePixPaymentData struct {

	// CNPJ do Iniciador de Pagamento devidamente habilitado para a prestao de Servio de Iniciao no Pix.
	// Example: 50685362000135
	// Required: true
	// Max Length: 14
	// Pattern: ^\d{14}$
	CnpjInitiator string `json:"cnpjInitiator"`

	// creditor account
	// Required: true
	CreditorAccount *CreditorAccount `json:"creditorAccount"`

	// Traz o cdigo da cidade segundo o IBGE (Instituto Brasileiro de Geografia e Estatstica).
	// Para o preenchimento deste campo, o Iniciador de Pagamentos deve seguir a orientao do arranjo da forma de pagamento.
	// O preenchimento do campo tanto em pix/payments quanto /consents deve ser igual. Caso haja divergncia dos valores, a instituio deve retornar HTTP 422 com o cdigo de erro PAGAMENTO_DIVERGENTE_DO_CONSENTIMENTO.
	// Este campo faz referncia ao campo CodMun do arranjo Pix.
	// Example: 5300108
	// Max Length: 7
	// Min Length: 7
	// Pattern: ^\d{7}$
	IbgeTownCode string `json:"ibgeTownCode,omitempty"`

	// local instrument
	// Required: true
	LocalInstrument *EnumLocalInstrument `json:"localInstrument"`

	// payment
	// Required: true
	Payment *PaymentPix `json:"payment"`

	// Chave cadastrada no DICT pertencente ao recebedor. Os tipos de chaves podem ser: telefone, e-mail, cpf/cnpj ou chave aleatria.
	// No caso de telefone celular deve ser informado no padro E.1641.
	// Para e-mail deve ter o formato xxxxxxxx@xxxxxxx.xxx(.xx) e no mximo 77 caracteres.
	// No caso de CPF dever ser informado com 11 nmeros, sem pontos ou traos.
	// Para o caso de CNPJ dever ser informado com 14 nmeros, sem pontos ou traos.
	// No caso de chave aleatria deve ser informado o UUID gerado pelo DICT, conforme formato especificado na RFC41223.
	// Se informado, a detentora da conta deve validar o proxy no DICT quando localInstrument for igual a DICT, QRDN ou QRES e validar o campo creditorAccount.
	// Esta validao  opcional caso o localInstrument for igual a INIC.
	// [Restrio] Se localInstrument for igual a MANU, o campo proxy no deve ser preenchido. Se localInstrument for igual INIC, DICT, QRDN ou QRES, o campo proxy deve ser sempre preenchido com a chave Pix.
	// Example: 12345678901
	// Max Length: 77
	// Pattern: [\w\W\s]*
	Proxy string `json:"proxy,omitempty"`

	// Sequncia de caracteres que corresponde ao QR Code disponibilizado para o pagador.
	//
	//  a sequncia de caracteres que seria lida pelo leitor de QR Code, e deve propiciar o retorno dos dados do pagador aps consulta na DICT.
	//
	// Essa funcionalidade  possvel tanto para QR Code esttico quanto para QR Code dinmico.
	// No arranjo do Pix esta  a mesma sequncia gerada e/ou lida pela funcionalidade Pix Copia e Cola.
	// Este campo dever ser no formato UTF-8.
	// [Restrio] Preenchimento obrigatrio para pagamentos por QR Code, observado o tamanho mximo de 512 bytes.
	// Example: 00020104141234567890123426660014BR.GOV.BCB.PIX014466756C616E6F32303139406578616D706C652E636F6D27300012\nBR.COM.OUTRO011001234567895204000053039865406123.455802BR5915NOMEDORECEBEDOR6008BRASILIA61087007490062\n530515RP12345678-201950300017BR.GOV.BCB.BRCODE01051.0.080450014BR.GOV.BCB.PIX0123PADRAO.URL.PIX/0123AB\nCD81390012BR.COM.OUTRO01190123.ABCD.3456.WXYZ6304EB76\n
	// Max Length: 512
	// Pattern: [\w\W\s]*
	QrCode string `json:"qrCode,omitempty"`

	// Deve ser preenchido sempre que o usurio pagador inserir alguma informao adicional em um pagamento, a ser enviada ao recebedor.
	// Example: Pagamento da nota XPTO035-002.
	// Max Length: 140
	// Pattern: [\w\W\s]*
	RemittanceInformation string `json:"remittanceInformation,omitempty"`

	// Trata-se de um identificador de transao que deve ser retransmitido intacto pelo PSP do pagador ao gerar a ordem de pagamento. Essa informao permitir ao recebedor identificar e correlacionar a transferncia, quando recebida, com a apresentao das instrues ao pagador.
	// Os caracteres permitidos no contexto do Pix para o campo txid (EMV 62-05) so:
	// - Letras minsculas, de a a z
	// - Letras maisculas, de A a z
	// - Dgitos decimais, de 0 a 9
	//
	// [Restrio] Se localInstrument for igual a INIC, o campo transactionIdentification deve ser preenchido obrigatoriamente.
	// Se localInstrument for igual a MANU ou DICT, o campo transactionIdentification no deve ser preenchido.
	// A detentora de conta deve validar se a condicionalidade do campo foi atendida pela iniciadora de pagamento.
	// Example: E00038166201907261559y6j6
	// Max Length: 25
	// Pattern: ^[a-zA-Z0-9][a-zA-Z0-9]{0,24}$
	TransactionIdentification string `json:"transactionIdentification,omitempty"`
}

// Validate validates this create pix payment data
func (m *CreatePixPaymentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCnpjInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbgeTownCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQrCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemittanceInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePixPaymentData) validateCnpjInitiator(formats strfmt.Registry) error {

	if err := validate.RequiredString("cnpjInitiator", "body", m.CnpjInitiator); err != nil {
		return err
	}

	if err := validate.MaxLength("cnpjInitiator", "body", m.CnpjInitiator, 14); err != nil {
		return err
	}

	if err := validate.Pattern("cnpjInitiator", "body", m.CnpjInitiator, `^\d{14}$`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePixPaymentData) validateCreditorAccount(formats strfmt.Registry) error {

	if err := validate.Required("creditorAccount", "body", m.CreditorAccount); err != nil {
		return err
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePixPaymentData) validateIbgeTownCode(formats strfmt.Registry) error {
	if swag.IsZero(m.IbgeTownCode) { // not required
		return nil
	}

	if err := validate.MinLength("ibgeTownCode", "body", m.IbgeTownCode, 7); err != nil {
		return err
	}

	if err := validate.MaxLength("ibgeTownCode", "body", m.IbgeTownCode, 7); err != nil {
		return err
	}

	if err := validate.Pattern("ibgeTownCode", "body", m.IbgeTownCode, `^\d{7}$`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePixPaymentData) validateLocalInstrument(formats strfmt.Registry) error {

	if err := validate.Required("localInstrument", "body", m.LocalInstrument); err != nil {
		return err
	}

	if err := validate.Required("localInstrument", "body", m.LocalInstrument); err != nil {
		return err
	}

	if m.LocalInstrument != nil {
		if err := m.LocalInstrument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localInstrument")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePixPaymentData) validatePayment(formats strfmt.Registry) error {

	if err := validate.Required("payment", "body", m.Payment); err != nil {
		return err
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePixPaymentData) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if err := validate.MaxLength("proxy", "body", m.Proxy, 77); err != nil {
		return err
	}

	if err := validate.Pattern("proxy", "body", m.Proxy, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePixPaymentData) validateQrCode(formats strfmt.Registry) error {
	if swag.IsZero(m.QrCode) { // not required
		return nil
	}

	if err := validate.MaxLength("qrCode", "body", m.QrCode, 512); err != nil {
		return err
	}

	if err := validate.Pattern("qrCode", "body", m.QrCode, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePixPaymentData) validateRemittanceInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.RemittanceInformation) { // not required
		return nil
	}

	if err := validate.MaxLength("remittanceInformation", "body", m.RemittanceInformation, 140); err != nil {
		return err
	}

	if err := validate.Pattern("remittanceInformation", "body", m.RemittanceInformation, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePixPaymentData) validateTransactionIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionIdentification) { // not required
		return nil
	}

	if err := validate.MaxLength("transactionIdentification", "body", m.TransactionIdentification, 25); err != nil {
		return err
	}

	if err := validate.Pattern("transactionIdentification", "body", m.TransactionIdentification, `^[a-zA-Z0-9][a-zA-Z0-9]{0,24}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create pix payment data based on the context it is used
func (m *CreatePixPaymentData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalInstrument(ctx, formats); err != nil {
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

func (m *CreatePixPaymentData) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePixPaymentData) contextValidateLocalInstrument(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalInstrument != nil {
		if err := m.LocalInstrument.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localInstrument")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePixPaymentData) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePixPaymentData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePixPaymentData) UnmarshalBinary(b []byte) error {
	var res CreatePixPaymentData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
