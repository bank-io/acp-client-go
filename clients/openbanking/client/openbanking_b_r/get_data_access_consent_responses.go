// Code generated by go-swagger; DO NOT EDIT.

package openbanking_b_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// GetDataAccessConsentReader is a Reader for the GetDataAccessConsent structure.
type GetDataAccessConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataAccessConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataAccessConsentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDataAccessConsentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDataAccessConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataAccessConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDataAccessConsentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDataAccessConsentNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDataAccessConsentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetDataAccessConsentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDataAccessConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataAccessConsentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataAccessConsentOK creates a GetDataAccessConsentOK with default headers values
func NewGetDataAccessConsentOK() *GetDataAccessConsentOK {
	return &GetDataAccessConsentOK{}
}

/* GetDataAccessConsentOK describes a response with status code 200, with default header values.

Customer data access consent
*/
type GetDataAccessConsentOK struct {
	Payload *models.BrazilCustomerDataAccessConsentResponse
}

func (o *GetDataAccessConsentOK) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentOK  %+v", 200, o.Payload)
}
func (o *GetDataAccessConsentOK) GetPayload() *models.BrazilCustomerDataAccessConsentResponse {
	return o.Payload
}

func (o *GetDataAccessConsentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BrazilCustomerDataAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentBadRequest creates a GetDataAccessConsentBadRequest with default headers values
func NewGetDataAccessConsentBadRequest() *GetDataAccessConsentBadRequest {
	return &GetDataAccessConsentBadRequest{}
}

/* GetDataAccessConsentBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetDataAccessConsentBadRequest struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentBadRequest) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentBadRequest  %+v", 400, o.Payload)
}
func (o *GetDataAccessConsentBadRequest) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentUnauthorized creates a GetDataAccessConsentUnauthorized with default headers values
func NewGetDataAccessConsentUnauthorized() *GetDataAccessConsentUnauthorized {
	return &GetDataAccessConsentUnauthorized{}
}

/* GetDataAccessConsentUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetDataAccessConsentUnauthorized struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDataAccessConsentUnauthorized) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentForbidden creates a GetDataAccessConsentForbidden with default headers values
func NewGetDataAccessConsentForbidden() *GetDataAccessConsentForbidden {
	return &GetDataAccessConsentForbidden{}
}

/* GetDataAccessConsentForbidden describes a response with status code 403, with default header values.

Error
*/
type GetDataAccessConsentForbidden struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentForbidden  %+v", 403, o.Payload)
}
func (o *GetDataAccessConsentForbidden) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentMethodNotAllowed creates a GetDataAccessConsentMethodNotAllowed with default headers values
func NewGetDataAccessConsentMethodNotAllowed() *GetDataAccessConsentMethodNotAllowed {
	return &GetDataAccessConsentMethodNotAllowed{}
}

/* GetDataAccessConsentMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type GetDataAccessConsentMethodNotAllowed struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetDataAccessConsentMethodNotAllowed) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentNotAcceptable creates a GetDataAccessConsentNotAcceptable with default headers values
func NewGetDataAccessConsentNotAcceptable() *GetDataAccessConsentNotAcceptable {
	return &GetDataAccessConsentNotAcceptable{}
}

/* GetDataAccessConsentNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type GetDataAccessConsentNotAcceptable struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentNotAcceptable  %+v", 406, o.Payload)
}
func (o *GetDataAccessConsentNotAcceptable) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentUnsupportedMediaType creates a GetDataAccessConsentUnsupportedMediaType with default headers values
func NewGetDataAccessConsentUnsupportedMediaType() *GetDataAccessConsentUnsupportedMediaType {
	return &GetDataAccessConsentUnsupportedMediaType{}
}

/* GetDataAccessConsentUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type GetDataAccessConsentUnsupportedMediaType struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetDataAccessConsentUnsupportedMediaType) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentUnprocessableEntity creates a GetDataAccessConsentUnprocessableEntity with default headers values
func NewGetDataAccessConsentUnprocessableEntity() *GetDataAccessConsentUnprocessableEntity {
	return &GetDataAccessConsentUnprocessableEntity{}
}

/* GetDataAccessConsentUnprocessableEntity describes a response with status code 422, with default header values.

Error
*/
type GetDataAccessConsentUnprocessableEntity struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetDataAccessConsentUnprocessableEntity) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentTooManyRequests creates a GetDataAccessConsentTooManyRequests with default headers values
func NewGetDataAccessConsentTooManyRequests() *GetDataAccessConsentTooManyRequests {
	return &GetDataAccessConsentTooManyRequests{}
}

/* GetDataAccessConsentTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type GetDataAccessConsentTooManyRequests struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetDataAccessConsentTooManyRequests) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentInternalServerError creates a GetDataAccessConsentInternalServerError with default headers values
func NewGetDataAccessConsentInternalServerError() *GetDataAccessConsentInternalServerError {
	return &GetDataAccessConsentInternalServerError{}
}

/* GetDataAccessConsentInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetDataAccessConsentInternalServerError struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDataAccessConsentInternalServerError) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
