// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// UpdateAdminTenantReader is a Reader for the UpdateAdminTenant structure.
type UpdateAdminTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdminTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAdminTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAdminTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAdminTenantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAdminTenantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAdminTenantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateAdminTenantUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateAdminTenantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAdminTenantOK creates a UpdateAdminTenantOK with default headers values
func NewUpdateAdminTenantOK() *UpdateAdminTenantOK {
	return &UpdateAdminTenantOK{}
}

/* UpdateAdminTenantOK describes a response with status code 200, with default header values.

Tenant
*/
type UpdateAdminTenantOK struct {
	Payload *models.Tenant
}

func (o *UpdateAdminTenantOK) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantOK  %+v", 200, o.Payload)
}
func (o *UpdateAdminTenantOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *UpdateAdminTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantBadRequest creates a UpdateAdminTenantBadRequest with default headers values
func NewUpdateAdminTenantBadRequest() *UpdateAdminTenantBadRequest {
	return &UpdateAdminTenantBadRequest{}
}

/* UpdateAdminTenantBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateAdminTenantBadRequest struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateAdminTenantBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantUnauthorized creates a UpdateAdminTenantUnauthorized with default headers values
func NewUpdateAdminTenantUnauthorized() *UpdateAdminTenantUnauthorized {
	return &UpdateAdminTenantUnauthorized{}
}

/* UpdateAdminTenantUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateAdminTenantUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateAdminTenantUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantForbidden creates a UpdateAdminTenantForbidden with default headers values
func NewUpdateAdminTenantForbidden() *UpdateAdminTenantForbidden {
	return &UpdateAdminTenantForbidden{}
}

/* UpdateAdminTenantForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateAdminTenantForbidden struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantForbidden) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantForbidden  %+v", 403, o.Payload)
}
func (o *UpdateAdminTenantForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantNotFound creates a UpdateAdminTenantNotFound with default headers values
func NewUpdateAdminTenantNotFound() *UpdateAdminTenantNotFound {
	return &UpdateAdminTenantNotFound{}
}

/* UpdateAdminTenantNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateAdminTenantNotFound struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantNotFound) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantNotFound  %+v", 404, o.Payload)
}
func (o *UpdateAdminTenantNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantUnprocessableEntity creates a UpdateAdminTenantUnprocessableEntity with default headers values
func NewUpdateAdminTenantUnprocessableEntity() *UpdateAdminTenantUnprocessableEntity {
	return &UpdateAdminTenantUnprocessableEntity{}
}

/* UpdateAdminTenantUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateAdminTenantUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateAdminTenantUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminTenantTooManyRequests creates a UpdateAdminTenantTooManyRequests with default headers values
func NewUpdateAdminTenantTooManyRequests() *UpdateAdminTenantTooManyRequests {
	return &UpdateAdminTenantTooManyRequests{}
}

/* UpdateAdminTenantTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateAdminTenantTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateAdminTenantTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /tenant][%d] updateAdminTenantTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateAdminTenantTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAdminTenantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
