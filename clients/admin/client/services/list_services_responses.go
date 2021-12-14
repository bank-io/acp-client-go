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

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/* ListServicesOK describes a response with status code 200, with default header values.

Services
*/
type ListServicesOK struct {
	Payload *models.ServicesResponse
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesOK  %+v", 200, o.Payload)
}
func (o *ListServicesOK) GetPayload() *models.ServicesResponse {
	return o.Payload
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesBadRequest creates a ListServicesBadRequest with default headers values
func NewListServicesBadRequest() *ListServicesBadRequest {
	return &ListServicesBadRequest{}
}

/* ListServicesBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type ListServicesBadRequest struct {
	Payload *models.Error
}

func (o *ListServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesBadRequest  %+v", 400, o.Payload)
}
func (o *ListServicesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesUnauthorized creates a ListServicesUnauthorized with default headers values
func NewListServicesUnauthorized() *ListServicesUnauthorized {
	return &ListServicesUnauthorized{}
}

/* ListServicesUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListServicesUnauthorized struct {
	Payload *models.Error
}

func (o *ListServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServicesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesForbidden creates a ListServicesForbidden with default headers values
func NewListServicesForbidden() *ListServicesForbidden {
	return &ListServicesForbidden{}
}

/* ListServicesForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListServicesForbidden struct {
	Payload *models.Error
}

func (o *ListServicesForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesForbidden  %+v", 403, o.Payload)
}
func (o *ListServicesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesNotFound creates a ListServicesNotFound with default headers values
func NewListServicesNotFound() *ListServicesNotFound {
	return &ListServicesNotFound{}
}

/* ListServicesNotFound describes a response with status code 404, with default header values.

HttpError
*/
type ListServicesNotFound struct {
	Payload *models.Error
}

func (o *ListServicesNotFound) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesNotFound  %+v", 404, o.Payload)
}
func (o *ListServicesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesTooManyRequests creates a ListServicesTooManyRequests with default headers values
func NewListServicesTooManyRequests() *ListServicesTooManyRequests {
	return &ListServicesTooManyRequests{}
}

/* ListServicesTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListServicesTooManyRequests struct {
	Payload *models.Error
}

func (o *ListServicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/services][%d] listServicesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServicesTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
