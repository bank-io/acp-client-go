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

// CreateDataAccessConsentDeprecatedReader is a Reader for the CreateDataAccessConsentDeprecated structure.
type CreateDataAccessConsentDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDataAccessConsentDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDataAccessConsentDeprecatedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDataAccessConsentDeprecatedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDataAccessConsentDeprecatedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDataAccessConsentDeprecatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateDataAccessConsentDeprecatedMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateDataAccessConsentDeprecatedNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDataAccessConsentDeprecatedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateDataAccessConsentDeprecatedUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDataAccessConsentDeprecatedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDataAccessConsentDeprecatedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDataAccessConsentDeprecatedCreated creates a CreateDataAccessConsentDeprecatedCreated with default headers values
func NewCreateDataAccessConsentDeprecatedCreated() *CreateDataAccessConsentDeprecatedCreated {
	return &CreateDataAccessConsentDeprecatedCreated{}
}

/* CreateDataAccessConsentDeprecatedCreated describes a response with status code 201, with default header values.

Customer data access consent
*/
type CreateDataAccessConsentDeprecatedCreated struct {
	Payload *models.BrazilCustomerDataAccessConsentResponse
}

func (o *CreateDataAccessConsentDeprecatedCreated) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedCreated  %+v", 201, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedCreated) GetPayload() *models.BrazilCustomerDataAccessConsentResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BrazilCustomerDataAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedBadRequest creates a CreateDataAccessConsentDeprecatedBadRequest with default headers values
func NewCreateDataAccessConsentDeprecatedBadRequest() *CreateDataAccessConsentDeprecatedBadRequest {
	return &CreateDataAccessConsentDeprecatedBadRequest{}
}

/* CreateDataAccessConsentDeprecatedBadRequest describes a response with status code 400, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedBadRequest struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedBadRequest) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedBadRequest) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedUnauthorized creates a CreateDataAccessConsentDeprecatedUnauthorized with default headers values
func NewCreateDataAccessConsentDeprecatedUnauthorized() *CreateDataAccessConsentDeprecatedUnauthorized {
	return &CreateDataAccessConsentDeprecatedUnauthorized{}
}

/* CreateDataAccessConsentDeprecatedUnauthorized describes a response with status code 401, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedUnauthorized struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedUnauthorized) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedForbidden creates a CreateDataAccessConsentDeprecatedForbidden with default headers values
func NewCreateDataAccessConsentDeprecatedForbidden() *CreateDataAccessConsentDeprecatedForbidden {
	return &CreateDataAccessConsentDeprecatedForbidden{}
}

/* CreateDataAccessConsentDeprecatedForbidden describes a response with status code 403, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedForbidden struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedForbidden  %+v", 403, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedForbidden) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedMethodNotAllowed creates a CreateDataAccessConsentDeprecatedMethodNotAllowed with default headers values
func NewCreateDataAccessConsentDeprecatedMethodNotAllowed() *CreateDataAccessConsentDeprecatedMethodNotAllowed {
	return &CreateDataAccessConsentDeprecatedMethodNotAllowed{}
}

/* CreateDataAccessConsentDeprecatedMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedMethodNotAllowed struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedMethodNotAllowed) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedNotAcceptable creates a CreateDataAccessConsentDeprecatedNotAcceptable with default headers values
func NewCreateDataAccessConsentDeprecatedNotAcceptable() *CreateDataAccessConsentDeprecatedNotAcceptable {
	return &CreateDataAccessConsentDeprecatedNotAcceptable{}
}

/* CreateDataAccessConsentDeprecatedNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedNotAcceptable struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedNotAcceptable  %+v", 406, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedNotAcceptable) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedUnsupportedMediaType creates a CreateDataAccessConsentDeprecatedUnsupportedMediaType with default headers values
func NewCreateDataAccessConsentDeprecatedUnsupportedMediaType() *CreateDataAccessConsentDeprecatedUnsupportedMediaType {
	return &CreateDataAccessConsentDeprecatedUnsupportedMediaType{}
}

/* CreateDataAccessConsentDeprecatedUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedUnsupportedMediaType struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedUnsupportedMediaType) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedUnprocessableEntity creates a CreateDataAccessConsentDeprecatedUnprocessableEntity with default headers values
func NewCreateDataAccessConsentDeprecatedUnprocessableEntity() *CreateDataAccessConsentDeprecatedUnprocessableEntity {
	return &CreateDataAccessConsentDeprecatedUnprocessableEntity{}
}

/* CreateDataAccessConsentDeprecatedUnprocessableEntity describes a response with status code 422, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedUnprocessableEntity struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedUnprocessableEntity) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedTooManyRequests creates a CreateDataAccessConsentDeprecatedTooManyRequests with default headers values
func NewCreateDataAccessConsentDeprecatedTooManyRequests() *CreateDataAccessConsentDeprecatedTooManyRequests {
	return &CreateDataAccessConsentDeprecatedTooManyRequests{}
}

/* CreateDataAccessConsentDeprecatedTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedTooManyRequests struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedTooManyRequests) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataAccessConsentDeprecatedInternalServerError creates a CreateDataAccessConsentDeprecatedInternalServerError with default headers values
func NewCreateDataAccessConsentDeprecatedInternalServerError() *CreateDataAccessConsentDeprecatedInternalServerError {
	return &CreateDataAccessConsentDeprecatedInternalServerError{}
}

/* CreateDataAccessConsentDeprecatedInternalServerError describes a response with status code 500, with default header values.

Error
*/
type CreateDataAccessConsentDeprecatedInternalServerError struct {
	Payload *models.OBBRErrorResponse
}

func (o *CreateDataAccessConsentDeprecatedInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/open-banking/consents/v1/consents][%d] createDataAccessConsentDeprecatedInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDataAccessConsentDeprecatedInternalServerError) GetPayload() *models.OBBRErrorResponse {
	return o.Payload
}

func (o *CreateDataAccessConsentDeprecatedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
