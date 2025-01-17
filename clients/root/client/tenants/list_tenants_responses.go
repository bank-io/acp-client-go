// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/root/models"
)

// ListTenantsReader is a Reader for the ListTenants structure.
type ListTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListTenantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTenantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTenantsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListTenantsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTenantsOK creates a ListTenantsOK with default headers values
func NewListTenantsOK() *ListTenantsOK {
	return &ListTenantsOK{}
}

/* ListTenantsOK describes a response with status code 200, with default header values.

List of tenants
*/
type ListTenantsOK struct {
	Payload *models.Tenants
}

func (o *ListTenantsOK) Error() string {
	return fmt.Sprintf("[GET /api/system/tenants][%d] listTenantsOK  %+v", 200, o.Payload)
}
func (o *ListTenantsOK) GetPayload() *models.Tenants {
	return o.Payload
}

func (o *ListTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenants)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTenantsUnauthorized creates a ListTenantsUnauthorized with default headers values
func NewListTenantsUnauthorized() *ListTenantsUnauthorized {
	return &ListTenantsUnauthorized{}
}

/* ListTenantsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListTenantsUnauthorized struct {
	Payload *models.Error
}

func (o *ListTenantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/system/tenants][%d] listTenantsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListTenantsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTenantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTenantsForbidden creates a ListTenantsForbidden with default headers values
func NewListTenantsForbidden() *ListTenantsForbidden {
	return &ListTenantsForbidden{}
}

/* ListTenantsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListTenantsForbidden struct {
	Payload *models.Error
}

func (o *ListTenantsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/system/tenants][%d] listTenantsForbidden  %+v", 403, o.Payload)
}
func (o *ListTenantsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTenantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTenantsNotFound creates a ListTenantsNotFound with default headers values
func NewListTenantsNotFound() *ListTenantsNotFound {
	return &ListTenantsNotFound{}
}

/* ListTenantsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type ListTenantsNotFound struct {
	Payload *models.Error
}

func (o *ListTenantsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/system/tenants][%d] listTenantsNotFound  %+v", 404, o.Payload)
}
func (o *ListTenantsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTenantsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTenantsTooManyRequests creates a ListTenantsTooManyRequests with default headers values
func NewListTenantsTooManyRequests() *ListTenantsTooManyRequests {
	return &ListTenantsTooManyRequests{}
}

/* ListTenantsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListTenantsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListTenantsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/system/tenants][%d] listTenantsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListTenantsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTenantsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
