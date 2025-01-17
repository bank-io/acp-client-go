// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// UpdateServiceReader is a Reader for the UpdateService structure.
type UpdateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateServiceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateServiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateServiceOK creates a UpdateServiceOK with default headers values
func NewUpdateServiceOK() *UpdateServiceOK {
	return &UpdateServiceOK{}
}

/* UpdateServiceOK describes a response with status code 200, with default header values.

Service
*/
type UpdateServiceOK struct {
	Payload *models.ServiceWithAudience
}

func (o *UpdateServiceOK) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceOK  %+v", 200, o.Payload)
}
func (o *UpdateServiceOK) GetPayload() *models.ServiceWithAudience {
	return o.Payload
}

func (o *UpdateServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceWithAudience)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceBadRequest creates a UpdateServiceBadRequest with default headers values
func NewUpdateServiceBadRequest() *UpdateServiceBadRequest {
	return &UpdateServiceBadRequest{}
}

/* UpdateServiceBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateServiceBadRequest struct {
	Payload *models.Error
}

func (o *UpdateServiceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateServiceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceUnauthorized creates a UpdateServiceUnauthorized with default headers values
func NewUpdateServiceUnauthorized() *UpdateServiceUnauthorized {
	return &UpdateServiceUnauthorized{}
}

/* UpdateServiceUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateServiceUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateServiceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateServiceUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceForbidden creates a UpdateServiceForbidden with default headers values
func NewUpdateServiceForbidden() *UpdateServiceForbidden {
	return &UpdateServiceForbidden{}
}

/* UpdateServiceForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateServiceForbidden struct {
	Payload *models.Error
}

func (o *UpdateServiceForbidden) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceForbidden  %+v", 403, o.Payload)
}
func (o *UpdateServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceNotFound creates a UpdateServiceNotFound with default headers values
func NewUpdateServiceNotFound() *UpdateServiceNotFound {
	return &UpdateServiceNotFound{}
}

/* UpdateServiceNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateServiceNotFound struct {
	Payload *models.Error
}

func (o *UpdateServiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceNotFound  %+v", 404, o.Payload)
}
func (o *UpdateServiceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceUnprocessableEntity creates a UpdateServiceUnprocessableEntity with default headers values
func NewUpdateServiceUnprocessableEntity() *UpdateServiceUnprocessableEntity {
	return &UpdateServiceUnprocessableEntity{}
}

/* UpdateServiceUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateServiceUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateServiceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateServiceUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceTooManyRequests creates a UpdateServiceTooManyRequests with default headers values
func NewUpdateServiceTooManyRequests() *UpdateServiceTooManyRequests {
	return &UpdateServiceTooManyRequests{}
}

/* UpdateServiceTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateServiceTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateServiceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /services/{sid}][%d] updateServiceTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateServiceTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
