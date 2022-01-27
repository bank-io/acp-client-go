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

// GetDataAccessConsentDeprecatedReader is a Reader for the GetDataAccessConsentDeprecated structure.
type GetDataAccessConsentDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataAccessConsentDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataAccessConsentDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDataAccessConsentDeprecatedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDataAccessConsentDeprecatedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataAccessConsentDeprecatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDataAccessConsentDeprecatedMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDataAccessConsentDeprecatedNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDataAccessConsentDeprecatedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetDataAccessConsentDeprecatedUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDataAccessConsentDeprecatedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataAccessConsentDeprecatedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataAccessConsentDeprecatedOK creates a GetDataAccessConsentDeprecatedOK with default headers values
func NewGetDataAccessConsentDeprecatedOK() *GetDataAccessConsentDeprecatedOK {
	return &GetDataAccessConsentDeprecatedOK{}
}

/* GetDataAccessConsentDeprecatedOK describes a response with status code 200, with default header values.

Customer data access consent
*/
type GetDataAccessConsentDeprecatedOK struct {
	Payload *models.OBBRCustomerDataAccessConsentResponse
}

func (o *GetDataAccessConsentDeprecatedOK) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedOK  %+v", 200, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedOK) GetPayload() *models.OBBRCustomerDataAccessConsentResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRCustomerDataAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedBadRequest creates a GetDataAccessConsentDeprecatedBadRequest with default headers values
func NewGetDataAccessConsentDeprecatedBadRequest() *GetDataAccessConsentDeprecatedBadRequest {
	return &GetDataAccessConsentDeprecatedBadRequest{}
}

/* GetDataAccessConsentDeprecatedBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedBadRequest struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedBadRequest) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedBadRequest  %+v", 400, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedBadRequest) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedUnauthorized creates a GetDataAccessConsentDeprecatedUnauthorized with default headers values
func NewGetDataAccessConsentDeprecatedUnauthorized() *GetDataAccessConsentDeprecatedUnauthorized {
	return &GetDataAccessConsentDeprecatedUnauthorized{}
}

/* GetDataAccessConsentDeprecatedUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedUnauthorized struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedUnauthorized) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedForbidden creates a GetDataAccessConsentDeprecatedForbidden with default headers values
func NewGetDataAccessConsentDeprecatedForbidden() *GetDataAccessConsentDeprecatedForbidden {
	return &GetDataAccessConsentDeprecatedForbidden{}
}

/* GetDataAccessConsentDeprecatedForbidden describes a response with status code 403, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedForbidden struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedForbidden  %+v", 403, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedForbidden) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedMethodNotAllowed creates a GetDataAccessConsentDeprecatedMethodNotAllowed with default headers values
func NewGetDataAccessConsentDeprecatedMethodNotAllowed() *GetDataAccessConsentDeprecatedMethodNotAllowed {
	return &GetDataAccessConsentDeprecatedMethodNotAllowed{}
}

/* GetDataAccessConsentDeprecatedMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedMethodNotAllowed struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedMethodNotAllowed) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedNotAcceptable creates a GetDataAccessConsentDeprecatedNotAcceptable with default headers values
func NewGetDataAccessConsentDeprecatedNotAcceptable() *GetDataAccessConsentDeprecatedNotAcceptable {
	return &GetDataAccessConsentDeprecatedNotAcceptable{}
}

/* GetDataAccessConsentDeprecatedNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedNotAcceptable struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedNotAcceptable  %+v", 406, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedNotAcceptable) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedUnsupportedMediaType creates a GetDataAccessConsentDeprecatedUnsupportedMediaType with default headers values
func NewGetDataAccessConsentDeprecatedUnsupportedMediaType() *GetDataAccessConsentDeprecatedUnsupportedMediaType {
	return &GetDataAccessConsentDeprecatedUnsupportedMediaType{}
}

/* GetDataAccessConsentDeprecatedUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedUnsupportedMediaType struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedUnsupportedMediaType) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedUnprocessableEntity creates a GetDataAccessConsentDeprecatedUnprocessableEntity with default headers values
func NewGetDataAccessConsentDeprecatedUnprocessableEntity() *GetDataAccessConsentDeprecatedUnprocessableEntity {
	return &GetDataAccessConsentDeprecatedUnprocessableEntity{}
}

/* GetDataAccessConsentDeprecatedUnprocessableEntity describes a response with status code 422, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedUnprocessableEntity struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedUnprocessableEntity) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedTooManyRequests creates a GetDataAccessConsentDeprecatedTooManyRequests with default headers values
func NewGetDataAccessConsentDeprecatedTooManyRequests() *GetDataAccessConsentDeprecatedTooManyRequests {
	return &GetDataAccessConsentDeprecatedTooManyRequests{}
}

/* GetDataAccessConsentDeprecatedTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedTooManyRequests struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedTooManyRequests) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataAccessConsentDeprecatedInternalServerError creates a GetDataAccessConsentDeprecatedInternalServerError with default headers values
func NewGetDataAccessConsentDeprecatedInternalServerError() *GetDataAccessConsentDeprecatedInternalServerError {
	return &GetDataAccessConsentDeprecatedInternalServerError{}
}

/* GetDataAccessConsentDeprecatedInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetDataAccessConsentDeprecatedInternalServerError struct {
	Payload *models.OBBRErrorResponse
}

func (o *GetDataAccessConsentDeprecatedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /open-banking-brasil/open-banking/consents/v1/consents/{consentID}][%d] getDataAccessConsentDeprecatedInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDataAccessConsentDeprecatedInternalServerError) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *GetDataAccessConsentDeprecatedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}