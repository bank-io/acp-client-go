// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// OpenbankingBrasilPaymentEnumAuthorisationStatusType OpenbankingBrasilPaymentEnumAuthorisationStatusType EnumAuthorisationStatusType
//
// Retorna o estado do consentimento, o qual no momento de sua criao ser AWAITING_AUTHORISATION.
// Este estado ser alterado depois da autorizao do consentimento na detentora da conta do pagador (Debtor) para AUTHORISED ou REJECTED.
// O consentimento fica no estado CONSUMED aps ocorrer a iniciao do pagamento referente ao consentimento.
// Em caso de consentimento expirado a detentora dever retornar o status REJECTED.
// Estados possveis:
// AWAITING_AUTHORISATION - Aguardando autorizao
// AUTHORISED - Autorizado
// REJECTED - Rejeitado
// CONSUMED - Consumido
//
// swagger:model OpenbankingBrasilPaymentEnumAuthorisationStatusType
type OpenbankingBrasilPaymentEnumAuthorisationStatusType string

// Validate validates this openbanking brasil payment enum authorisation status type
func (m OpenbankingBrasilPaymentEnumAuthorisationStatusType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openbanking brasil payment enum authorisation status type based on context it is used
func (m OpenbankingBrasilPaymentEnumAuthorisationStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}