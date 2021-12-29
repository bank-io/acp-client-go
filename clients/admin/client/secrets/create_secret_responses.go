// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// CreateSecretReader is a Reader for the CreateSecret structure.
type CreateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSecretCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSecretConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSecretUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSecretTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSecretCreated creates a CreateSecretCreated with default headers values
func NewCreateSecretCreated() *CreateSecretCreated {
	return &CreateSecretCreated{}
}

/* CreateSecretCreated describes a response with status code 201, with default header values.

Secret
*/
type CreateSecretCreated struct {
	Payload *models.Secret
}

func (o *CreateSecretCreated) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretCreated  %+v", 201, o.Payload)
}
func (o *CreateSecretCreated) GetPayload() *models.Secret {
	return o.Payload
}

func (o *CreateSecretCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretBadRequest creates a CreateSecretBadRequest with default headers values
func NewCreateSecretBadRequest() *CreateSecretBadRequest {
	return &CreateSecretBadRequest{}
}

/* CreateSecretBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type CreateSecretBadRequest struct {
	Payload *models.Error
}

func (o *CreateSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSecretBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretUnauthorized creates a CreateSecretUnauthorized with default headers values
func NewCreateSecretUnauthorized() *CreateSecretUnauthorized {
	return &CreateSecretUnauthorized{}
}

/* CreateSecretUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type CreateSecretUnauthorized struct {
	Payload *models.Error
}

func (o *CreateSecretUnauthorized) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateSecretUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretForbidden creates a CreateSecretForbidden with default headers values
func NewCreateSecretForbidden() *CreateSecretForbidden {
	return &CreateSecretForbidden{}
}

/* CreateSecretForbidden describes a response with status code 403, with default header values.

HttpError
*/
type CreateSecretForbidden struct {
	Payload *models.Error
}

func (o *CreateSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretForbidden  %+v", 403, o.Payload)
}
func (o *CreateSecretForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretNotFound creates a CreateSecretNotFound with default headers values
func NewCreateSecretNotFound() *CreateSecretNotFound {
	return &CreateSecretNotFound{}
}

/* CreateSecretNotFound describes a response with status code 404, with default header values.

HttpError
*/
type CreateSecretNotFound struct {
	Payload *models.Error
}

func (o *CreateSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretNotFound  %+v", 404, o.Payload)
}
func (o *CreateSecretNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretConflict creates a CreateSecretConflict with default headers values
func NewCreateSecretConflict() *CreateSecretConflict {
	return &CreateSecretConflict{}
}

/* CreateSecretConflict describes a response with status code 409, with default header values.

HttpError
*/
type CreateSecretConflict struct {
	Payload *models.Error
}

func (o *CreateSecretConflict) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretConflict  %+v", 409, o.Payload)
}
func (o *CreateSecretConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretUnprocessableEntity creates a CreateSecretUnprocessableEntity with default headers values
func NewCreateSecretUnprocessableEntity() *CreateSecretUnprocessableEntity {
	return &CreateSecretUnprocessableEntity{}
}

/* CreateSecretUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type CreateSecretUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateSecretUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateSecretUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecretTooManyRequests creates a CreateSecretTooManyRequests with default headers values
func NewCreateSecretTooManyRequests() *CreateSecretTooManyRequests {
	return &CreateSecretTooManyRequests{}
}

/* CreateSecretTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type CreateSecretTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateSecretTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/secrets][%d] createSecretTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateSecretTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSecretTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}