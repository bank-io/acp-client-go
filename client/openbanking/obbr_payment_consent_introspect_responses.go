// Code generated by go-swagger; DO NOT EDIT.

package openbanking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// ObbrPaymentConsentIntrospectReader is a Reader for the ObbrPaymentConsentIntrospect structure.
type ObbrPaymentConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObbrPaymentConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObbrPaymentConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObbrPaymentConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObbrPaymentConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObbrPaymentConsentIntrospectOK creates a ObbrPaymentConsentIntrospectOK with default headers values
func NewObbrPaymentConsentIntrospectOK() *ObbrPaymentConsentIntrospectOK {
	return &ObbrPaymentConsentIntrospectOK{}
}

/* ObbrPaymentConsentIntrospectOK describes a response with status code 200, with default header values.

IntrospectOBBRPaymentConsentResponse
*/
type ObbrPaymentConsentIntrospectOK struct {
	Payload *models.IntrospectOBBRPaymentConsentResponse
}

func (o *ObbrPaymentConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/payments/v1/consents/introspect][%d] obbrPaymentConsentIntrospectOK  %+v", 200, o.Payload)
}
func (o *ObbrPaymentConsentIntrospectOK) GetPayload() *models.IntrospectOBBRPaymentConsentResponse {
	return o.Payload
}

func (o *ObbrPaymentConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntrospectOBBRPaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObbrPaymentConsentIntrospectUnauthorized creates a ObbrPaymentConsentIntrospectUnauthorized with default headers values
func NewObbrPaymentConsentIntrospectUnauthorized() *ObbrPaymentConsentIntrospectUnauthorized {
	return &ObbrPaymentConsentIntrospectUnauthorized{}
}

/* ObbrPaymentConsentIntrospectUnauthorized describes a response with status code 401, with default header values.

genericError
*/
type ObbrPaymentConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *ObbrPaymentConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/payments/v1/consents/introspect][%d] obbrPaymentConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}
func (o *ObbrPaymentConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ObbrPaymentConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObbrPaymentConsentIntrospectNotFound creates a ObbrPaymentConsentIntrospectNotFound with default headers values
func NewObbrPaymentConsentIntrospectNotFound() *ObbrPaymentConsentIntrospectNotFound {
	return &ObbrPaymentConsentIntrospectNotFound{}
}

/* ObbrPaymentConsentIntrospectNotFound describes a response with status code 404, with default header values.

genericError
*/
type ObbrPaymentConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *ObbrPaymentConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/payments/v1/consents/introspect][%d] obbrPaymentConsentIntrospectNotFound  %+v", 404, o.Payload)
}
func (o *ObbrPaymentConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ObbrPaymentConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}