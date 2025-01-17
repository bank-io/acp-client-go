// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// OpenbankingInternationalScheduledPaymentConsentIntrospectReader is a Reader for the OpenbankingInternationalScheduledPaymentConsentIntrospect structure.
type OpenbankingInternationalScheduledPaymentConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenbankingInternationalScheduledPaymentConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenbankingInternationalScheduledPaymentConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenbankingInternationalScheduledPaymentConsentIntrospectOK creates a OpenbankingInternationalScheduledPaymentConsentIntrospectOK with default headers values
func NewOpenbankingInternationalScheduledPaymentConsentIntrospectOK() *OpenbankingInternationalScheduledPaymentConsentIntrospectOK {
	return &OpenbankingInternationalScheduledPaymentConsentIntrospectOK{}
}

/* OpenbankingInternationalScheduledPaymentConsentIntrospectOK describes a response with status code 200, with default header values.

Introspect Openbanking International Scheduled Payment Consent Response
*/
type OpenbankingInternationalScheduledPaymentConsentIntrospectOK struct {
	Payload *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-scheduled-payment-consents/introspect][%d] openbankingInternationalScheduledPaymentConsentIntrospectOK  %+v", 200, o.Payload)
}
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOK) GetPayload() *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody {
	return o.Payload
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized creates a OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized with default headers values
func NewOpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized() *OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized {
	return &OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized{}
}

/* OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized describes a response with status code 401, with default header values.

ErrorResponse
*/
type OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-scheduled-payment-consents/introspect][%d] openbankingInternationalScheduledPaymentConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalScheduledPaymentConsentIntrospectNotFound creates a OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound with default headers values
func NewOpenbankingInternationalScheduledPaymentConsentIntrospectNotFound() *OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound {
	return &OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound{}
}

/* OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound describes a response with status code 404, with default header values.

ErrorResponse
*/
type OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-scheduled-payment-consents/introspect][%d] openbankingInternationalScheduledPaymentConsentIntrospectNotFound  %+v", 404, o.Payload)
}
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests creates a OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests with default headers values
func NewOpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests() *OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests {
	return &OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests{}
}

/* OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests describes a response with status code 429, with default header values.

ErrorResponse
*/
type OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-scheduled-payment-consents/introspect][%d] openbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests  %+v", 429, o.Payload)
}
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody openbanking international scheduled payment consent introspect o k body
swagger:model OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody
*/
type OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody struct {
	models.IntrospectResponse

	models.InternationalScheduledPaymentConsent

	// account i ds
	AccountIDs []string `json:"AccountIDs"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) UnmarshalJSON(raw []byte) error {
	// OpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0
	var openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0 models.IntrospectResponse
	if err := swag.ReadJSON(raw, &openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0); err != nil {
		return err
	}
	o.IntrospectResponse = openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0

	// OpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1
	var openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1 models.InternationalScheduledPaymentConsent
	if err := swag.ReadJSON(raw, &openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1); err != nil {
		return err
	}
	o.InternationalScheduledPaymentConsent = openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1

	// OpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2
	var dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2 struct {
		AccountIDs []string `json:"AccountIDs"`
	}
	if err := swag.ReadJSON(raw, &dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2); err != nil {
		return err
	}

	o.AccountIDs = dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2.AccountIDs

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0, err := swag.WriteJSON(o.IntrospectResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO0)

	openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1, err := swag.WriteJSON(o.InternationalScheduledPaymentConsent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, openbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO1)
	var dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2 struct {
		AccountIDs []string `json:"AccountIDs"`
	}

	dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2.AccountIDs = o.AccountIDs

	jsonDataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2, errOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2 := swag.WriteJSON(dataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2)
	if errOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2 != nil {
		return nil, errOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2
	}
	_parts = append(_parts, jsonDataOpenbankingInternationalScheduledPaymentConsentIntrospectOKBodyAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this openbanking international scheduled payment consent introspect o k body
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IntrospectResponse
	if err := o.IntrospectResponse.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.InternationalScheduledPaymentConsent
	if err := o.InternationalScheduledPaymentConsent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this openbanking international scheduled payment consent introspect o k body based on the context it is used
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IntrospectResponse
	if err := o.IntrospectResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.InternationalScheduledPaymentConsent
	if err := o.InternationalScheduledPaymentConsent.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody) UnmarshalBinary(b []byte) error {
	var res OpenbankingInternationalScheduledPaymentConsentIntrospectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
