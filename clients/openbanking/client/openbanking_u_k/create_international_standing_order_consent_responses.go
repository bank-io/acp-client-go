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

// CreateInternationalStandingOrderConsentReader is a Reader for the CreateInternationalStandingOrderConsent structure.
type CreateInternationalStandingOrderConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternationalStandingOrderConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInternationalStandingOrderConsentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInternationalStandingOrderConsentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInternationalStandingOrderConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInternationalStandingOrderConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateInternationalStandingOrderConsentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateInternationalStandingOrderConsentNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateInternationalStandingOrderConsentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInternationalStandingOrderConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternationalStandingOrderConsentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInternationalStandingOrderConsentCreated creates a CreateInternationalStandingOrderConsentCreated with default headers values
func NewCreateInternationalStandingOrderConsentCreated() *CreateInternationalStandingOrderConsentCreated {
	return &CreateInternationalStandingOrderConsentCreated{}
}

/* CreateInternationalStandingOrderConsentCreated describes a response with status code 201, with default header values.

International standing order consent
*/
type CreateInternationalStandingOrderConsentCreated struct {
	Payload *models.InternationalStandingOrderConsentResponse
}

func (o *CreateInternationalStandingOrderConsentCreated) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentCreated  %+v", 201, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentCreated) GetPayload() *models.InternationalStandingOrderConsentResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternationalStandingOrderConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentBadRequest creates a CreateInternationalStandingOrderConsentBadRequest with default headers values
func NewCreateInternationalStandingOrderConsentBadRequest() *CreateInternationalStandingOrderConsentBadRequest {
	return &CreateInternationalStandingOrderConsentBadRequest{}
}

/* CreateInternationalStandingOrderConsentBadRequest describes a response with status code 400, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentBadRequest) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentBadRequest  %+v", 400, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentUnauthorized creates a CreateInternationalStandingOrderConsentUnauthorized with default headers values
func NewCreateInternationalStandingOrderConsentUnauthorized() *CreateInternationalStandingOrderConsentUnauthorized {
	return &CreateInternationalStandingOrderConsentUnauthorized{}
}

/* CreateInternationalStandingOrderConsentUnauthorized describes a response with status code 401, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentForbidden creates a CreateInternationalStandingOrderConsentForbidden with default headers values
func NewCreateInternationalStandingOrderConsentForbidden() *CreateInternationalStandingOrderConsentForbidden {
	return &CreateInternationalStandingOrderConsentForbidden{}
}

/* CreateInternationalStandingOrderConsentForbidden describes a response with status code 403, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentForbidden  %+v", 403, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentMethodNotAllowed creates a CreateInternationalStandingOrderConsentMethodNotAllowed with default headers values
func NewCreateInternationalStandingOrderConsentMethodNotAllowed() *CreateInternationalStandingOrderConsentMethodNotAllowed {
	return &CreateInternationalStandingOrderConsentMethodNotAllowed{}
}

/* CreateInternationalStandingOrderConsentMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentNotAcceptable creates a CreateInternationalStandingOrderConsentNotAcceptable with default headers values
func NewCreateInternationalStandingOrderConsentNotAcceptable() *CreateInternationalStandingOrderConsentNotAcceptable {
	return &CreateInternationalStandingOrderConsentNotAcceptable{}
}

/* CreateInternationalStandingOrderConsentNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentNotAcceptable  %+v", 406, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentUnsupportedMediaType creates a CreateInternationalStandingOrderConsentUnsupportedMediaType with default headers values
func NewCreateInternationalStandingOrderConsentUnsupportedMediaType() *CreateInternationalStandingOrderConsentUnsupportedMediaType {
	return &CreateInternationalStandingOrderConsentUnsupportedMediaType{}
}

/* CreateInternationalStandingOrderConsentUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentTooManyRequests creates a CreateInternationalStandingOrderConsentTooManyRequests with default headers values
func NewCreateInternationalStandingOrderConsentTooManyRequests() *CreateInternationalStandingOrderConsentTooManyRequests {
	return &CreateInternationalStandingOrderConsentTooManyRequests{}
}

/* CreateInternationalStandingOrderConsentTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrderConsentInternalServerError creates a CreateInternationalStandingOrderConsentInternalServerError with default headers values
func NewCreateInternationalStandingOrderConsentInternalServerError() *CreateInternationalStandingOrderConsentInternalServerError {
	return &CreateInternationalStandingOrderConsentInternalServerError{}
}

/* CreateInternationalStandingOrderConsentInternalServerError describes a response with status code 500, with default header values.

Error
*/
type CreateInternationalStandingOrderConsentInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateInternationalStandingOrderConsentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open-banking/v3.1/pisp/international-standing-order-consents][%d] createInternationalStandingOrderConsentInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateInternationalStandingOrderConsentInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInternationalStandingOrderConsentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
