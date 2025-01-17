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

// CreateDomesticScheduledPaymentConsentRequestReader is a Reader for the CreateDomesticScheduledPaymentConsentRequest structure.
type CreateDomesticScheduledPaymentConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDomesticScheduledPaymentConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDomesticScheduledPaymentConsentRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDomesticScheduledPaymentConsentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDomesticScheduledPaymentConsentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDomesticScheduledPaymentConsentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateDomesticScheduledPaymentConsentRequestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateDomesticScheduledPaymentConsentRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDomesticScheduledPaymentConsentRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDomesticScheduledPaymentConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDomesticScheduledPaymentConsentRequestCreated creates a CreateDomesticScheduledPaymentConsentRequestCreated with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestCreated() *CreateDomesticScheduledPaymentConsentRequestCreated {
	return &CreateDomesticScheduledPaymentConsentRequestCreated{}
}

/* CreateDomesticScheduledPaymentConsentRequestCreated describes a response with status code 201, with default header values.

Domestic scheduled payment consent
*/
type CreateDomesticScheduledPaymentConsentRequestCreated struct {
	Payload *models.DomesticScheduledPaymentConsentResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestCreated) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestCreated  %+v", 201, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestCreated) GetPayload() *models.DomesticScheduledPaymentConsentResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomesticScheduledPaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestBadRequest creates a CreateDomesticScheduledPaymentConsentRequestBadRequest with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestBadRequest() *CreateDomesticScheduledPaymentConsentRequestBadRequest {
	return &CreateDomesticScheduledPaymentConsentRequestBadRequest{}
}

/* CreateDomesticScheduledPaymentConsentRequestBadRequest describes a response with status code 400, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestUnauthorized creates a CreateDomesticScheduledPaymentConsentRequestUnauthorized with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestUnauthorized() *CreateDomesticScheduledPaymentConsentRequestUnauthorized {
	return &CreateDomesticScheduledPaymentConsentRequestUnauthorized{}
}

/* CreateDomesticScheduledPaymentConsentRequestUnauthorized describes a response with status code 401, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestForbidden creates a CreateDomesticScheduledPaymentConsentRequestForbidden with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestForbidden() *CreateDomesticScheduledPaymentConsentRequestForbidden {
	return &CreateDomesticScheduledPaymentConsentRequestForbidden{}
}

/* CreateDomesticScheduledPaymentConsentRequestForbidden describes a response with status code 403, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestForbidden  %+v", 403, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestMethodNotAllowed creates a CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestMethodNotAllowed() *CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed {
	return &CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed{}
}

/* CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestNotAcceptable creates a CreateDomesticScheduledPaymentConsentRequestNotAcceptable with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestNotAcceptable() *CreateDomesticScheduledPaymentConsentRequestNotAcceptable {
	return &CreateDomesticScheduledPaymentConsentRequestNotAcceptable{}
}

/* CreateDomesticScheduledPaymentConsentRequestNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestNotAcceptable  %+v", 406, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType creates a CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType() *CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType {
	return &CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType{}
}

/* CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestTooManyRequests creates a CreateDomesticScheduledPaymentConsentRequestTooManyRequests with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestTooManyRequests() *CreateDomesticScheduledPaymentConsentRequestTooManyRequests {
	return &CreateDomesticScheduledPaymentConsentRequestTooManyRequests{}
}

/* CreateDomesticScheduledPaymentConsentRequestTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentRequestInternalServerError creates a CreateDomesticScheduledPaymentConsentRequestInternalServerError with default headers values
func NewCreateDomesticScheduledPaymentConsentRequestInternalServerError() *CreateDomesticScheduledPaymentConsentRequestInternalServerError {
	return &CreateDomesticScheduledPaymentConsentRequestInternalServerError{}
}

/* CreateDomesticScheduledPaymentConsentRequestInternalServerError describes a response with status code 500, with default header values.

Error
*/
type CreateDomesticScheduledPaymentConsentRequestInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateDomesticScheduledPaymentConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentRequestInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
