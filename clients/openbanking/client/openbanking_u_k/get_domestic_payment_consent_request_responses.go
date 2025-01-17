// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// GetDomesticPaymentConsentRequestReader is a Reader for the GetDomesticPaymentConsentRequest structure.
type GetDomesticPaymentConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomesticPaymentConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomesticPaymentConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomesticPaymentConsentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomesticPaymentConsentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDomesticPaymentConsentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDomesticPaymentConsentRequestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDomesticPaymentConsentRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDomesticPaymentConsentRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDomesticPaymentConsentRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomesticPaymentConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomesticPaymentConsentRequestOK creates a GetDomesticPaymentConsentRequestOK with default headers values
func NewGetDomesticPaymentConsentRequestOK() *GetDomesticPaymentConsentRequestOK {
	return &GetDomesticPaymentConsentRequestOK{}
}

/* GetDomesticPaymentConsentRequestOK describes a response with status code 200, with default header values.

Domestic payment consent
*/
type GetDomesticPaymentConsentRequestOK struct {
	Payload *models.DomesticPaymentConsentResponse
}

func (o *GetDomesticPaymentConsentRequestOK) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestOK  %+v", 200, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestOK) GetPayload() *models.DomesticPaymentConsentResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomesticPaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestBadRequest creates a GetDomesticPaymentConsentRequestBadRequest with default headers values
func NewGetDomesticPaymentConsentRequestBadRequest() *GetDomesticPaymentConsentRequestBadRequest {
	return &GetDomesticPaymentConsentRequestBadRequest{}
}

/* GetDomesticPaymentConsentRequestBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestBadRequest  %+v", 400, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestUnauthorized creates a GetDomesticPaymentConsentRequestUnauthorized with default headers values
func NewGetDomesticPaymentConsentRequestUnauthorized() *GetDomesticPaymentConsentRequestUnauthorized {
	return &GetDomesticPaymentConsentRequestUnauthorized{}
}

/* GetDomesticPaymentConsentRequestUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestForbidden creates a GetDomesticPaymentConsentRequestForbidden with default headers values
func NewGetDomesticPaymentConsentRequestForbidden() *GetDomesticPaymentConsentRequestForbidden {
	return &GetDomesticPaymentConsentRequestForbidden{}
}

/* GetDomesticPaymentConsentRequestForbidden describes a response with status code 403, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestForbidden  %+v", 403, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestMethodNotAllowed creates a GetDomesticPaymentConsentRequestMethodNotAllowed with default headers values
func NewGetDomesticPaymentConsentRequestMethodNotAllowed() *GetDomesticPaymentConsentRequestMethodNotAllowed {
	return &GetDomesticPaymentConsentRequestMethodNotAllowed{}
}

/* GetDomesticPaymentConsentRequestMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestNotAcceptable creates a GetDomesticPaymentConsentRequestNotAcceptable with default headers values
func NewGetDomesticPaymentConsentRequestNotAcceptable() *GetDomesticPaymentConsentRequestNotAcceptable {
	return &GetDomesticPaymentConsentRequestNotAcceptable{}
}

/* GetDomesticPaymentConsentRequestNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestNotAcceptable  %+v", 406, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestUnsupportedMediaType creates a GetDomesticPaymentConsentRequestUnsupportedMediaType with default headers values
func NewGetDomesticPaymentConsentRequestUnsupportedMediaType() *GetDomesticPaymentConsentRequestUnsupportedMediaType {
	return &GetDomesticPaymentConsentRequestUnsupportedMediaType{}
}

/* GetDomesticPaymentConsentRequestUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestTooManyRequests creates a GetDomesticPaymentConsentRequestTooManyRequests with default headers values
func NewGetDomesticPaymentConsentRequestTooManyRequests() *GetDomesticPaymentConsentRequestTooManyRequests {
	return &GetDomesticPaymentConsentRequestTooManyRequests{}
}

/* GetDomesticPaymentConsentRequestTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentRequestInternalServerError creates a GetDomesticPaymentConsentRequestInternalServerError with default headers values
func NewGetDomesticPaymentConsentRequestInternalServerError() *GetDomesticPaymentConsentRequestInternalServerError {
	return &GetDomesticPaymentConsentRequestInternalServerError{}
}

/* GetDomesticPaymentConsentRequestInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetDomesticPaymentConsentRequestInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetDomesticPaymentConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/domestic-payment-consents/{consentID}][%d] getDomesticPaymentConsentRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDomesticPaymentConsentRequestInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomesticPaymentConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
