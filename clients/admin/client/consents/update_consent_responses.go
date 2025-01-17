// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// UpdateConsentReader is a Reader for the UpdateConsent structure.
type UpdateConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateConsentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConsentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateConsentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateConsentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConsentCreated creates a UpdateConsentCreated with default headers values
func NewUpdateConsentCreated() *UpdateConsentCreated {
	return &UpdateConsentCreated{}
}

/* UpdateConsentCreated describes a response with status code 201, with default header values.

Consent
*/
type UpdateConsentCreated struct {
	Payload *models.Consent
}

func (o *UpdateConsentCreated) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentCreated  %+v", 201, o.Payload)
}
func (o *UpdateConsentCreated) GetPayload() *models.Consent {
	return o.Payload
}

func (o *UpdateConsentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentUnauthorized creates a UpdateConsentUnauthorized with default headers values
func NewUpdateConsentUnauthorized() *UpdateConsentUnauthorized {
	return &UpdateConsentUnauthorized{}
}

/* UpdateConsentUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateConsentUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateConsentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateConsentUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentForbidden creates a UpdateConsentForbidden with default headers values
func NewUpdateConsentForbidden() *UpdateConsentForbidden {
	return &UpdateConsentForbidden{}
}

/* UpdateConsentForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateConsentForbidden struct {
	Payload *models.Error
}

func (o *UpdateConsentForbidden) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentForbidden  %+v", 403, o.Payload)
}
func (o *UpdateConsentForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentNotFound creates a UpdateConsentNotFound with default headers values
func NewUpdateConsentNotFound() *UpdateConsentNotFound {
	return &UpdateConsentNotFound{}
}

/* UpdateConsentNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateConsentNotFound struct {
	Payload *models.Error
}

func (o *UpdateConsentNotFound) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentNotFound  %+v", 404, o.Payload)
}
func (o *UpdateConsentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentConflict creates a UpdateConsentConflict with default headers values
func NewUpdateConsentConflict() *UpdateConsentConflict {
	return &UpdateConsentConflict{}
}

/* UpdateConsentConflict describes a response with status code 409, with default header values.

HttpError
*/
type UpdateConsentConflict struct {
	Payload *models.Error
}

func (o *UpdateConsentConflict) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentConflict  %+v", 409, o.Payload)
}
func (o *UpdateConsentConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentUnprocessableEntity creates a UpdateConsentUnprocessableEntity with default headers values
func NewUpdateConsentUnprocessableEntity() *UpdateConsentUnprocessableEntity {
	return &UpdateConsentUnprocessableEntity{}
}

/* UpdateConsentUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateConsentUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateConsentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateConsentUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentTooManyRequests creates a UpdateConsentTooManyRequests with default headers values
func NewUpdateConsentTooManyRequests() *UpdateConsentTooManyRequests {
	return &UpdateConsentTooManyRequests{}
}

/* UpdateConsentTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateConsentTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /consents/{consent}][%d] updateConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateConsentTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
