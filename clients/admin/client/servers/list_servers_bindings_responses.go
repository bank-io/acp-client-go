// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ListServersBindingsReader is a Reader for the ListServersBindings structure.
type ListServersBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServersBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServersBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServersBindingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServersBindingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServersBindingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServersBindingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServersBindingsOK creates a ListServersBindingsOK with default headers values
func NewListServersBindingsOK() *ListServersBindingsOK {
	return &ListServersBindingsOK{}
}

/* ListServersBindingsOK describes a response with status code 200, with default header values.

Server bindings
*/
type ListServersBindingsOK struct {
	Payload *models.ServersBindings
}

func (o *ListServersBindingsOK) Error() string {
	return fmt.Sprintf("[GET /servers-bindings][%d] listServersBindingsOK  %+v", 200, o.Payload)
}
func (o *ListServersBindingsOK) GetPayload() *models.ServersBindings {
	return o.Payload
}

func (o *ListServersBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServersBindings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersBindingsUnauthorized creates a ListServersBindingsUnauthorized with default headers values
func NewListServersBindingsUnauthorized() *ListServersBindingsUnauthorized {
	return &ListServersBindingsUnauthorized{}
}

/* ListServersBindingsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListServersBindingsUnauthorized struct {
	Payload *models.Error
}

func (o *ListServersBindingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers-bindings][%d] listServersBindingsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServersBindingsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersBindingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersBindingsForbidden creates a ListServersBindingsForbidden with default headers values
func NewListServersBindingsForbidden() *ListServersBindingsForbidden {
	return &ListServersBindingsForbidden{}
}

/* ListServersBindingsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListServersBindingsForbidden struct {
	Payload *models.Error
}

func (o *ListServersBindingsForbidden) Error() string {
	return fmt.Sprintf("[GET /servers-bindings][%d] listServersBindingsForbidden  %+v", 403, o.Payload)
}
func (o *ListServersBindingsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersBindingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersBindingsNotFound creates a ListServersBindingsNotFound with default headers values
func NewListServersBindingsNotFound() *ListServersBindingsNotFound {
	return &ListServersBindingsNotFound{}
}

/* ListServersBindingsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type ListServersBindingsNotFound struct {
	Payload *models.Error
}

func (o *ListServersBindingsNotFound) Error() string {
	return fmt.Sprintf("[GET /servers-bindings][%d] listServersBindingsNotFound  %+v", 404, o.Payload)
}
func (o *ListServersBindingsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersBindingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersBindingsTooManyRequests creates a ListServersBindingsTooManyRequests with default headers values
func NewListServersBindingsTooManyRequests() *ListServersBindingsTooManyRequests {
	return &ListServersBindingsTooManyRequests{}
}

/* ListServersBindingsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListServersBindingsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListServersBindingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers-bindings][%d] listServersBindingsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServersBindingsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersBindingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}