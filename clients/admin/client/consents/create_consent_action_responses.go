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

// CreateConsentActionReader is a Reader for the CreateConsentAction structure.
type CreateConsentActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConsentActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConsentActionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateConsentActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConsentActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConsentActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateConsentActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateConsentActionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateConsentActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConsentActionCreated creates a CreateConsentActionCreated with default headers values
func NewCreateConsentActionCreated() *CreateConsentActionCreated {
	return &CreateConsentActionCreated{}
}

/* CreateConsentActionCreated describes a response with status code 201, with default header values.

Consent action with consents
*/
type CreateConsentActionCreated struct {
	Payload *models.ConsentActionWithConsents
}

func (o *CreateConsentActionCreated) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionCreated  %+v", 201, o.Payload)
}
func (o *CreateConsentActionCreated) GetPayload() *models.ConsentActionWithConsents {
	return o.Payload
}

func (o *CreateConsentActionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentActionWithConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionUnauthorized creates a CreateConsentActionUnauthorized with default headers values
func NewCreateConsentActionUnauthorized() *CreateConsentActionUnauthorized {
	return &CreateConsentActionUnauthorized{}
}

/* CreateConsentActionUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type CreateConsentActionUnauthorized struct {
	Payload *models.Error
}

func (o *CreateConsentActionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateConsentActionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionForbidden creates a CreateConsentActionForbidden with default headers values
func NewCreateConsentActionForbidden() *CreateConsentActionForbidden {
	return &CreateConsentActionForbidden{}
}

/* CreateConsentActionForbidden describes a response with status code 403, with default header values.

HttpError
*/
type CreateConsentActionForbidden struct {
	Payload *models.Error
}

func (o *CreateConsentActionForbidden) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionForbidden  %+v", 403, o.Payload)
}
func (o *CreateConsentActionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionNotFound creates a CreateConsentActionNotFound with default headers values
func NewCreateConsentActionNotFound() *CreateConsentActionNotFound {
	return &CreateConsentActionNotFound{}
}

/* CreateConsentActionNotFound describes a response with status code 404, with default header values.

HttpError
*/
type CreateConsentActionNotFound struct {
	Payload *models.Error
}

func (o *CreateConsentActionNotFound) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionNotFound  %+v", 404, o.Payload)
}
func (o *CreateConsentActionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionConflict creates a CreateConsentActionConflict with default headers values
func NewCreateConsentActionConflict() *CreateConsentActionConflict {
	return &CreateConsentActionConflict{}
}

/* CreateConsentActionConflict describes a response with status code 409, with default header values.

HttpError
*/
type CreateConsentActionConflict struct {
	Payload *models.Error
}

func (o *CreateConsentActionConflict) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionConflict  %+v", 409, o.Payload)
}
func (o *CreateConsentActionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionUnprocessableEntity creates a CreateConsentActionUnprocessableEntity with default headers values
func NewCreateConsentActionUnprocessableEntity() *CreateConsentActionUnprocessableEntity {
	return &CreateConsentActionUnprocessableEntity{}
}

/* CreateConsentActionUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type CreateConsentActionUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateConsentActionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateConsentActionUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionTooManyRequests creates a CreateConsentActionTooManyRequests with default headers values
func NewCreateConsentActionTooManyRequests() *CreateConsentActionTooManyRequests {
	return &CreateConsentActionTooManyRequests{}
}

/* CreateConsentActionTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type CreateConsentActionTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateConsentActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /actions][%d] createConsentActionTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateConsentActionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
