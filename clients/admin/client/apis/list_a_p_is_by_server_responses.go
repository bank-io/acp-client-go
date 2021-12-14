// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ListAPIsByServerReader is a Reader for the ListAPIsByServer structure.
type ListAPIsByServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAPIsByServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAPIsByServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAPIsByServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAPIsByServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAPIsByServerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAPIsByServerOK creates a ListAPIsByServerOK with default headers values
func NewListAPIsByServerOK() *ListAPIsByServerOK {
	return &ListAPIsByServerOK{}
}

/* ListAPIsByServerOK describes a response with status code 200, with default header values.

APIs grouped by service ids
*/
type ListAPIsByServerOK struct {
	Payload *models.ServerAPIs
}

func (o *ListAPIsByServerOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/apis][%d] listAPIsByServerOK  %+v", 200, o.Payload)
}
func (o *ListAPIsByServerOK) GetPayload() *models.ServerAPIs {
	return o.Payload
}

func (o *ListAPIsByServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerAPIs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServerUnauthorized creates a ListAPIsByServerUnauthorized with default headers values
func NewListAPIsByServerUnauthorized() *ListAPIsByServerUnauthorized {
	return &ListAPIsByServerUnauthorized{}
}

/* ListAPIsByServerUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListAPIsByServerUnauthorized struct {
	Payload *models.Error
}

func (o *ListAPIsByServerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/apis][%d] listAPIsByServerUnauthorized  %+v", 401, o.Payload)
}
func (o *ListAPIsByServerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServerForbidden creates a ListAPIsByServerForbidden with default headers values
func NewListAPIsByServerForbidden() *ListAPIsByServerForbidden {
	return &ListAPIsByServerForbidden{}
}

/* ListAPIsByServerForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListAPIsByServerForbidden struct {
	Payload *models.Error
}

func (o *ListAPIsByServerForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/apis][%d] listAPIsByServerForbidden  %+v", 403, o.Payload)
}
func (o *ListAPIsByServerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServerTooManyRequests creates a ListAPIsByServerTooManyRequests with default headers values
func NewListAPIsByServerTooManyRequests() *ListAPIsByServerTooManyRequests {
	return &ListAPIsByServerTooManyRequests{}
}

/* ListAPIsByServerTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListAPIsByServerTooManyRequests struct {
	Payload *models.Error
}

func (o *ListAPIsByServerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/apis][%d] listAPIsByServerTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListAPIsByServerTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
