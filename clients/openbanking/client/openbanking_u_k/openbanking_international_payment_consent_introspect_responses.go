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

// OpenbankingInternationalPaymentConsentIntrospectReader is a Reader for the OpenbankingInternationalPaymentConsentIntrospect structure.
type OpenbankingInternationalPaymentConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenbankingInternationalPaymentConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenbankingInternationalPaymentConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOpenbankingInternationalPaymentConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenbankingInternationalPaymentConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOpenbankingInternationalPaymentConsentIntrospectTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenbankingInternationalPaymentConsentIntrospectOK creates a OpenbankingInternationalPaymentConsentIntrospectOK with default headers values
func NewOpenbankingInternationalPaymentConsentIntrospectOK() *OpenbankingInternationalPaymentConsentIntrospectOK {
	return &OpenbankingInternationalPaymentConsentIntrospectOK{}
}

/* OpenbankingInternationalPaymentConsentIntrospectOK describes a response with status code 200, with default header values.

Introspect Openbanking International Payment Consent Response
*/
type OpenbankingInternationalPaymentConsentIntrospectOK struct {
	Payload *OpenbankingInternationalPaymentConsentIntrospectOKBody
}

func (o *OpenbankingInternationalPaymentConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-payment-consents/introspect][%d] openbankingInternationalPaymentConsentIntrospectOK  %+v", 200, o.Payload)
}
func (o *OpenbankingInternationalPaymentConsentIntrospectOK) GetPayload() *OpenbankingInternationalPaymentConsentIntrospectOKBody {
	return o.Payload
}

func (o *OpenbankingInternationalPaymentConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(OpenbankingInternationalPaymentConsentIntrospectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalPaymentConsentIntrospectUnauthorized creates a OpenbankingInternationalPaymentConsentIntrospectUnauthorized with default headers values
func NewOpenbankingInternationalPaymentConsentIntrospectUnauthorized() *OpenbankingInternationalPaymentConsentIntrospectUnauthorized {
	return &OpenbankingInternationalPaymentConsentIntrospectUnauthorized{}
}

/* OpenbankingInternationalPaymentConsentIntrospectUnauthorized describes a response with status code 401, with default header values.

ErrorResponse
*/
type OpenbankingInternationalPaymentConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalPaymentConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-payment-consents/introspect][%d] openbankingInternationalPaymentConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenbankingInternationalPaymentConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalPaymentConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalPaymentConsentIntrospectNotFound creates a OpenbankingInternationalPaymentConsentIntrospectNotFound with default headers values
func NewOpenbankingInternationalPaymentConsentIntrospectNotFound() *OpenbankingInternationalPaymentConsentIntrospectNotFound {
	return &OpenbankingInternationalPaymentConsentIntrospectNotFound{}
}

/* OpenbankingInternationalPaymentConsentIntrospectNotFound describes a response with status code 404, with default header values.

ErrorResponse
*/
type OpenbankingInternationalPaymentConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalPaymentConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-payment-consents/introspect][%d] openbankingInternationalPaymentConsentIntrospectNotFound  %+v", 404, o.Payload)
}
func (o *OpenbankingInternationalPaymentConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalPaymentConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingInternationalPaymentConsentIntrospectTooManyRequests creates a OpenbankingInternationalPaymentConsentIntrospectTooManyRequests with default headers values
func NewOpenbankingInternationalPaymentConsentIntrospectTooManyRequests() *OpenbankingInternationalPaymentConsentIntrospectTooManyRequests {
	return &OpenbankingInternationalPaymentConsentIntrospectTooManyRequests{}
}

/* OpenbankingInternationalPaymentConsentIntrospectTooManyRequests describes a response with status code 429, with default header values.

ErrorResponse
*/
type OpenbankingInternationalPaymentConsentIntrospectTooManyRequests struct {
	Payload *models.GenericError
}

func (o *OpenbankingInternationalPaymentConsentIntrospectTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-payment-consents/introspect][%d] openbankingInternationalPaymentConsentIntrospectTooManyRequests  %+v", 429, o.Payload)
}
func (o *OpenbankingInternationalPaymentConsentIntrospectTooManyRequests) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingInternationalPaymentConsentIntrospectTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OpenbankingInternationalPaymentConsentIntrospectOKBody openbanking international payment consent introspect o k body
swagger:model OpenbankingInternationalPaymentConsentIntrospectOKBody
*/
type OpenbankingInternationalPaymentConsentIntrospectOKBody struct {
	models.IntrospectResponse

	models.InternationalPaymentConsent

	// account i ds
	AccountIDs []string `json:"AccountIDs"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *OpenbankingInternationalPaymentConsentIntrospectOKBody) UnmarshalJSON(raw []byte) error {
	// OpenbankingInternationalPaymentConsentIntrospectOKBodyAO0
	var openbankingInternationalPaymentConsentIntrospectOKBodyAO0 models.IntrospectResponse
	if err := swag.ReadJSON(raw, &openbankingInternationalPaymentConsentIntrospectOKBodyAO0); err != nil {
		return err
	}
	o.IntrospectResponse = openbankingInternationalPaymentConsentIntrospectOKBodyAO0

	// OpenbankingInternationalPaymentConsentIntrospectOKBodyAO1
	var openbankingInternationalPaymentConsentIntrospectOKBodyAO1 models.InternationalPaymentConsent
	if err := swag.ReadJSON(raw, &openbankingInternationalPaymentConsentIntrospectOKBodyAO1); err != nil {
		return err
	}
	o.InternationalPaymentConsent = openbankingInternationalPaymentConsentIntrospectOKBodyAO1

	// OpenbankingInternationalPaymentConsentIntrospectOKBodyAO2
	var dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2 struct {
		AccountIDs []string `json:"AccountIDs"`
	}
	if err := swag.ReadJSON(raw, &dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2); err != nil {
		return err
	}

	o.AccountIDs = dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2.AccountIDs

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o OpenbankingInternationalPaymentConsentIntrospectOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	openbankingInternationalPaymentConsentIntrospectOKBodyAO0, err := swag.WriteJSON(o.IntrospectResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, openbankingInternationalPaymentConsentIntrospectOKBodyAO0)

	openbankingInternationalPaymentConsentIntrospectOKBodyAO1, err := swag.WriteJSON(o.InternationalPaymentConsent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, openbankingInternationalPaymentConsentIntrospectOKBodyAO1)
	var dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2 struct {
		AccountIDs []string `json:"AccountIDs"`
	}

	dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2.AccountIDs = o.AccountIDs

	jsonDataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2, errOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2 := swag.WriteJSON(dataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2)
	if errOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2 != nil {
		return nil, errOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2
	}
	_parts = append(_parts, jsonDataOpenbankingInternationalPaymentConsentIntrospectOKBodyAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this openbanking international payment consent introspect o k body
func (o *OpenbankingInternationalPaymentConsentIntrospectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IntrospectResponse
	if err := o.IntrospectResponse.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.InternationalPaymentConsent
	if err := o.InternationalPaymentConsent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this openbanking international payment consent introspect o k body based on the context it is used
func (o *OpenbankingInternationalPaymentConsentIntrospectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IntrospectResponse
	if err := o.IntrospectResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.InternationalPaymentConsent
	if err := o.InternationalPaymentConsent.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *OpenbankingInternationalPaymentConsentIntrospectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OpenbankingInternationalPaymentConsentIntrospectOKBody) UnmarshalBinary(b []byte) error {
	var res OpenbankingInternationalPaymentConsentIntrospectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
