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

// BindServerReader is a Reader for the BindServer structure.
type BindServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBindServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBindServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewBindServerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBindServerOK creates a BindServerOK with default headers values
func NewBindServerOK() *BindServerOK {
	return &BindServerOK{}
}

/* BindServerOK describes a response with status code 200, with default header values.

Server to server binding
*/
type BindServerOK struct {
	Payload *models.ServerToServer
}

func (o *BindServerOK) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/bind/{rid}][%d] bindServerOK  %+v", 200, o.Payload)
}
func (o *BindServerOK) GetPayload() *models.ServerToServer {
	return o.Payload
}

func (o *BindServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerToServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindServerUnauthorized creates a BindServerUnauthorized with default headers values
func NewBindServerUnauthorized() *BindServerUnauthorized {
	return &BindServerUnauthorized{}
}

/* BindServerUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type BindServerUnauthorized struct {
	Payload *models.Error
}

func (o *BindServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/bind/{rid}][%d] bindServerUnauthorized  %+v", 401, o.Payload)
}
func (o *BindServerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindServerForbidden creates a BindServerForbidden with default headers values
func NewBindServerForbidden() *BindServerForbidden {
	return &BindServerForbidden{}
}

/* BindServerForbidden describes a response with status code 403, with default header values.

HttpError
*/
type BindServerForbidden struct {
	Payload *models.Error
}

func (o *BindServerForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/bind/{rid}][%d] bindServerForbidden  %+v", 403, o.Payload)
}
func (o *BindServerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindServerNotFound creates a BindServerNotFound with default headers values
func NewBindServerNotFound() *BindServerNotFound {
	return &BindServerNotFound{}
}

/* BindServerNotFound describes a response with status code 404, with default header values.

HttpError
*/
type BindServerNotFound struct {
	Payload *models.Error
}

func (o *BindServerNotFound) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/bind/{rid}][%d] bindServerNotFound  %+v", 404, o.Payload)
}
func (o *BindServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindServerTooManyRequests creates a BindServerTooManyRequests with default headers values
func NewBindServerTooManyRequests() *BindServerTooManyRequests {
	return &BindServerTooManyRequests{}
}

/* BindServerTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type BindServerTooManyRequests struct {
	Payload *models.Error
}

func (o *BindServerTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/bind/{rid}][%d] bindServerTooManyRequests  %+v", 429, o.Payload)
}
func (o *BindServerTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindServerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
