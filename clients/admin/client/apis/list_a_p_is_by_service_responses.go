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

// ListAPIsByServiceReader is a Reader for the ListAPIsByService structure.
type ListAPIsByServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAPIsByServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAPIsByServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAPIsByServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAPIsByServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAPIsByServiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAPIsByServiceOK creates a ListAPIsByServiceOK with default headers values
func NewListAPIsByServiceOK() *ListAPIsByServiceOK {
	return &ListAPIsByServiceOK{}
}

/* ListAPIsByServiceOK describes a response with status code 200, with default header values.

APIs
*/
type ListAPIsByServiceOK struct {
	Payload *models.APIs
}

func (o *ListAPIsByServiceOK) Error() string {
	return fmt.Sprintf("[GET /services/{sid}/apis][%d] listAPIsByServiceOK  %+v", 200, o.Payload)
}
func (o *ListAPIsByServiceOK) GetPayload() *models.APIs {
	return o.Payload
}

func (o *ListAPIsByServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServiceUnauthorized creates a ListAPIsByServiceUnauthorized with default headers values
func NewListAPIsByServiceUnauthorized() *ListAPIsByServiceUnauthorized {
	return &ListAPIsByServiceUnauthorized{}
}

/* ListAPIsByServiceUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListAPIsByServiceUnauthorized struct {
	Payload *models.Error
}

func (o *ListAPIsByServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services/{sid}/apis][%d] listAPIsByServiceUnauthorized  %+v", 401, o.Payload)
}
func (o *ListAPIsByServiceUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServiceForbidden creates a ListAPIsByServiceForbidden with default headers values
func NewListAPIsByServiceForbidden() *ListAPIsByServiceForbidden {
	return &ListAPIsByServiceForbidden{}
}

/* ListAPIsByServiceForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListAPIsByServiceForbidden struct {
	Payload *models.Error
}

func (o *ListAPIsByServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /services/{sid}/apis][%d] listAPIsByServiceForbidden  %+v", 403, o.Payload)
}
func (o *ListAPIsByServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIsByServiceTooManyRequests creates a ListAPIsByServiceTooManyRequests with default headers values
func NewListAPIsByServiceTooManyRequests() *ListAPIsByServiceTooManyRequests {
	return &ListAPIsByServiceTooManyRequests{}
}

/* ListAPIsByServiceTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListAPIsByServiceTooManyRequests struct {
	Payload *models.Error
}

func (o *ListAPIsByServiceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /services/{sid}/apis][%d] listAPIsByServiceTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListAPIsByServiceTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAPIsByServiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
