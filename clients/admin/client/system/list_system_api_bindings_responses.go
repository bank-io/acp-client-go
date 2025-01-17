// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ListSystemAPIBindingsReader is a Reader for the ListSystemAPIBindings structure.
type ListSystemAPIBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemAPIBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSystemAPIBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSystemAPIBindingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSystemAPIBindingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListSystemAPIBindingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSystemAPIBindingsOK creates a ListSystemAPIBindingsOK with default headers values
func NewListSystemAPIBindingsOK() *ListSystemAPIBindingsOK {
	return &ListSystemAPIBindingsOK{}
}

/* ListSystemAPIBindingsOK describes a response with status code 200, with default header values.

System API bindings
*/
type ListSystemAPIBindingsOK struct {
	Payload *models.SystemAPIBindings
}

func (o *ListSystemAPIBindingsOK) Error() string {
	return fmt.Sprintf("[GET /system/apis/{wid}/bindings][%d] listSystemApiBindingsOK  %+v", 200, o.Payload)
}
func (o *ListSystemAPIBindingsOK) GetPayload() *models.SystemAPIBindings {
	return o.Payload
}

func (o *ListSystemAPIBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemAPIBindings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemAPIBindingsUnauthorized creates a ListSystemAPIBindingsUnauthorized with default headers values
func NewListSystemAPIBindingsUnauthorized() *ListSystemAPIBindingsUnauthorized {
	return &ListSystemAPIBindingsUnauthorized{}
}

/* ListSystemAPIBindingsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListSystemAPIBindingsUnauthorized struct {
	Payload *models.Error
}

func (o *ListSystemAPIBindingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/apis/{wid}/bindings][%d] listSystemApiBindingsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListSystemAPIBindingsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSystemAPIBindingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemAPIBindingsForbidden creates a ListSystemAPIBindingsForbidden with default headers values
func NewListSystemAPIBindingsForbidden() *ListSystemAPIBindingsForbidden {
	return &ListSystemAPIBindingsForbidden{}
}

/* ListSystemAPIBindingsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListSystemAPIBindingsForbidden struct {
	Payload *models.Error
}

func (o *ListSystemAPIBindingsForbidden) Error() string {
	return fmt.Sprintf("[GET /system/apis/{wid}/bindings][%d] listSystemApiBindingsForbidden  %+v", 403, o.Payload)
}
func (o *ListSystemAPIBindingsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSystemAPIBindingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemAPIBindingsTooManyRequests creates a ListSystemAPIBindingsTooManyRequests with default headers values
func NewListSystemAPIBindingsTooManyRequests() *ListSystemAPIBindingsTooManyRequests {
	return &ListSystemAPIBindingsTooManyRequests{}
}

/* ListSystemAPIBindingsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListSystemAPIBindingsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListSystemAPIBindingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /system/apis/{wid}/bindings][%d] listSystemApiBindingsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListSystemAPIBindingsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSystemAPIBindingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
