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

// GetInternationalStandingOrderConsentRequestReader is a Reader for the GetInternationalStandingOrderConsentRequest structure.
type GetInternationalStandingOrderConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternationalStandingOrderConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternationalStandingOrderConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInternationalStandingOrderConsentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInternationalStandingOrderConsentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternationalStandingOrderConsentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInternationalStandingOrderConsentRequestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetInternationalStandingOrderConsentRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetInternationalStandingOrderConsentRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInternationalStandingOrderConsentRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternationalStandingOrderConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternationalStandingOrderConsentRequestOK creates a GetInternationalStandingOrderConsentRequestOK with default headers values
func NewGetInternationalStandingOrderConsentRequestOK() *GetInternationalStandingOrderConsentRequestOK {
	return &GetInternationalStandingOrderConsentRequestOK{}
}

/* GetInternationalStandingOrderConsentRequestOK describes a response with status code 200, with default header values.

International standing order consent
*/
type GetInternationalStandingOrderConsentRequestOK struct {
	Payload *models.InternationalStandingOrderConsentResponse
}

func (o *GetInternationalStandingOrderConsentRequestOK) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestOK  %+v", 200, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestOK) GetPayload() *models.InternationalStandingOrderConsentResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternationalStandingOrderConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestBadRequest creates a GetInternationalStandingOrderConsentRequestBadRequest with default headers values
func NewGetInternationalStandingOrderConsentRequestBadRequest() *GetInternationalStandingOrderConsentRequestBadRequest {
	return &GetInternationalStandingOrderConsentRequestBadRequest{}
}

/* GetInternationalStandingOrderConsentRequestBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestBadRequest  %+v", 400, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestUnauthorized creates a GetInternationalStandingOrderConsentRequestUnauthorized with default headers values
func NewGetInternationalStandingOrderConsentRequestUnauthorized() *GetInternationalStandingOrderConsentRequestUnauthorized {
	return &GetInternationalStandingOrderConsentRequestUnauthorized{}
}

/* GetInternationalStandingOrderConsentRequestUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestForbidden creates a GetInternationalStandingOrderConsentRequestForbidden with default headers values
func NewGetInternationalStandingOrderConsentRequestForbidden() *GetInternationalStandingOrderConsentRequestForbidden {
	return &GetInternationalStandingOrderConsentRequestForbidden{}
}

/* GetInternationalStandingOrderConsentRequestForbidden describes a response with status code 403, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestForbidden  %+v", 403, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestMethodNotAllowed creates a GetInternationalStandingOrderConsentRequestMethodNotAllowed with default headers values
func NewGetInternationalStandingOrderConsentRequestMethodNotAllowed() *GetInternationalStandingOrderConsentRequestMethodNotAllowed {
	return &GetInternationalStandingOrderConsentRequestMethodNotAllowed{}
}

/* GetInternationalStandingOrderConsentRequestMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestNotAcceptable creates a GetInternationalStandingOrderConsentRequestNotAcceptable with default headers values
func NewGetInternationalStandingOrderConsentRequestNotAcceptable() *GetInternationalStandingOrderConsentRequestNotAcceptable {
	return &GetInternationalStandingOrderConsentRequestNotAcceptable{}
}

/* GetInternationalStandingOrderConsentRequestNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestNotAcceptable  %+v", 406, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestUnsupportedMediaType creates a GetInternationalStandingOrderConsentRequestUnsupportedMediaType with default headers values
func NewGetInternationalStandingOrderConsentRequestUnsupportedMediaType() *GetInternationalStandingOrderConsentRequestUnsupportedMediaType {
	return &GetInternationalStandingOrderConsentRequestUnsupportedMediaType{}
}

/* GetInternationalStandingOrderConsentRequestUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestTooManyRequests creates a GetInternationalStandingOrderConsentRequestTooManyRequests with default headers values
func NewGetInternationalStandingOrderConsentRequestTooManyRequests() *GetInternationalStandingOrderConsentRequestTooManyRequests {
	return &GetInternationalStandingOrderConsentRequestTooManyRequests{}
}

/* GetInternationalStandingOrderConsentRequestTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentRequestInternalServerError creates a GetInternationalStandingOrderConsentRequestInternalServerError with default headers values
func NewGetInternationalStandingOrderConsentRequestInternalServerError() *GetInternationalStandingOrderConsentRequestInternalServerError {
	return &GetInternationalStandingOrderConsentRequestInternalServerError{}
}

/* GetInternationalStandingOrderConsentRequestInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetInternationalStandingOrderConsentRequestInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetInternationalStandingOrderConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/international-standing-order-consents/{consentID}][%d] getInternationalStandingOrderConsentRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInternationalStandingOrderConsentRequestInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
