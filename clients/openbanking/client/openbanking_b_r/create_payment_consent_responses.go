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

// CreatePaymentConsentReader is a Reader for the CreatePaymentConsent structure.
type CreatePaymentConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePaymentConsentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePaymentConsentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePaymentConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePaymentConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreatePaymentConsentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreatePaymentConsentNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreatePaymentConsentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePaymentConsentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreatePaymentConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePaymentConsentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePaymentConsentCreated creates a CreatePaymentConsentCreated with default headers values
func NewCreatePaymentConsentCreated() *CreatePaymentConsentCreated {
	return &CreatePaymentConsentCreated{}
}

/* CreatePaymentConsentCreated describes a response with status code 201, with default header values.

Customer payment consent
*/
type CreatePaymentConsentCreated struct {
	Payload *models.OBBRCustomerPaymentConsentResponse
}

func (o *CreatePaymentConsentCreated) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentCreated  %+v", 201, o.Payload)
}
func (o *CreatePaymentConsentCreated) GetPayload() *models.OBBRCustomerPaymentConsentResponse {
	return o.Payload
}

func (o *CreatePaymentConsentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRCustomerPaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentBadRequest creates a CreatePaymentConsentBadRequest with default headers values
func NewCreatePaymentConsentBadRequest() *CreatePaymentConsentBadRequest {
	return &CreatePaymentConsentBadRequest{}
}

/* CreatePaymentConsentBadRequest describes a response with status code 400, with default header values.

Error
*/
type CreatePaymentConsentBadRequest struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentBadRequest) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentBadRequest  %+v", 400, o.Payload)
}
func (o *CreatePaymentConsentBadRequest) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentUnauthorized creates a CreatePaymentConsentUnauthorized with default headers values
func NewCreatePaymentConsentUnauthorized() *CreatePaymentConsentUnauthorized {
	return &CreatePaymentConsentUnauthorized{}
}

/* CreatePaymentConsentUnauthorized describes a response with status code 401, with default header values.

Error
*/
type CreatePaymentConsentUnauthorized struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *CreatePaymentConsentUnauthorized) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentForbidden creates a CreatePaymentConsentForbidden with default headers values
func NewCreatePaymentConsentForbidden() *CreatePaymentConsentForbidden {
	return &CreatePaymentConsentForbidden{}
}

/* CreatePaymentConsentForbidden describes a response with status code 403, with default header values.

Error
*/
type CreatePaymentConsentForbidden struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentForbidden  %+v", 403, o.Payload)
}
func (o *CreatePaymentConsentForbidden) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentMethodNotAllowed creates a CreatePaymentConsentMethodNotAllowed with default headers values
func NewCreatePaymentConsentMethodNotAllowed() *CreatePaymentConsentMethodNotAllowed {
	return &CreatePaymentConsentMethodNotAllowed{}
}

/* CreatePaymentConsentMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type CreatePaymentConsentMethodNotAllowed struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *CreatePaymentConsentMethodNotAllowed) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentNotAcceptable creates a CreatePaymentConsentNotAcceptable with default headers values
func NewCreatePaymentConsentNotAcceptable() *CreatePaymentConsentNotAcceptable {
	return &CreatePaymentConsentNotAcceptable{}
}

/* CreatePaymentConsentNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type CreatePaymentConsentNotAcceptable struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentNotAcceptable  %+v", 406, o.Payload)
}
func (o *CreatePaymentConsentNotAcceptable) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentUnsupportedMediaType creates a CreatePaymentConsentUnsupportedMediaType with default headers values
func NewCreatePaymentConsentUnsupportedMediaType() *CreatePaymentConsentUnsupportedMediaType {
	return &CreatePaymentConsentUnsupportedMediaType{}
}

/* CreatePaymentConsentUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type CreatePaymentConsentUnsupportedMediaType struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreatePaymentConsentUnsupportedMediaType) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentUnprocessableEntity creates a CreatePaymentConsentUnprocessableEntity with default headers values
func NewCreatePaymentConsentUnprocessableEntity() *CreatePaymentConsentUnprocessableEntity {
	return &CreatePaymentConsentUnprocessableEntity{}
}

/* CreatePaymentConsentUnprocessableEntity describes a response with status code 422, with default header values.

Error
*/
type CreatePaymentConsentUnprocessableEntity struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreatePaymentConsentUnprocessableEntity) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentTooManyRequests creates a CreatePaymentConsentTooManyRequests with default headers values
func NewCreatePaymentConsentTooManyRequests() *CreatePaymentConsentTooManyRequests {
	return &CreatePaymentConsentTooManyRequests{}
}

/* CreatePaymentConsentTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type CreatePaymentConsentTooManyRequests struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreatePaymentConsentTooManyRequests) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentConsentInternalServerError creates a CreatePaymentConsentInternalServerError with default headers values
func NewCreatePaymentConsentInternalServerError() *CreatePaymentConsentInternalServerError {
	return &CreatePaymentConsentInternalServerError{}
}

/* CreatePaymentConsentInternalServerError describes a response with status code 500, with default header values.

Error
*/
type CreatePaymentConsentInternalServerError struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreatePaymentConsentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open-banking/payments/v1/consents][%d] createPaymentConsentInternalServerError  %+v", 500, o.Payload)
}
func (o *CreatePaymentConsentInternalServerError) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreatePaymentConsentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}