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

// GetCIBAAuthenticationServiceReader is a Reader for the GetCIBAAuthenticationService structure.
type GetCIBAAuthenticationServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCIBAAuthenticationServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCIBAAuthenticationServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCIBAAuthenticationServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCIBAAuthenticationServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCIBAAuthenticationServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCIBAAuthenticationServiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCIBAAuthenticationServiceOK creates a GetCIBAAuthenticationServiceOK with default headers values
func NewGetCIBAAuthenticationServiceOK() *GetCIBAAuthenticationServiceOK {
	return &GetCIBAAuthenticationServiceOK{}
}

/* GetCIBAAuthenticationServiceOK describes a response with status code 200, with default header values.

CIBA authentication service
*/
type GetCIBAAuthenticationServiceOK struct {
	Payload *models.CIBAAuthenticationService
}

func (o *GetCIBAAuthenticationServiceOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/ciba-authentication-service][%d] getCIBAAuthenticationServiceOK  %+v", 200, o.Payload)
}
func (o *GetCIBAAuthenticationServiceOK) GetPayload() *models.CIBAAuthenticationService {
	return o.Payload
}

func (o *GetCIBAAuthenticationServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CIBAAuthenticationService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIBAAuthenticationServiceUnauthorized creates a GetCIBAAuthenticationServiceUnauthorized with default headers values
func NewGetCIBAAuthenticationServiceUnauthorized() *GetCIBAAuthenticationServiceUnauthorized {
	return &GetCIBAAuthenticationServiceUnauthorized{}
}

/* GetCIBAAuthenticationServiceUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetCIBAAuthenticationServiceUnauthorized struct {
	Payload *models.Error
}

func (o *GetCIBAAuthenticationServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/ciba-authentication-service][%d] getCIBAAuthenticationServiceUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCIBAAuthenticationServiceUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCIBAAuthenticationServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIBAAuthenticationServiceForbidden creates a GetCIBAAuthenticationServiceForbidden with default headers values
func NewGetCIBAAuthenticationServiceForbidden() *GetCIBAAuthenticationServiceForbidden {
	return &GetCIBAAuthenticationServiceForbidden{}
}

/* GetCIBAAuthenticationServiceForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetCIBAAuthenticationServiceForbidden struct {
	Payload *models.Error
}

func (o *GetCIBAAuthenticationServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/ciba-authentication-service][%d] getCIBAAuthenticationServiceForbidden  %+v", 403, o.Payload)
}
func (o *GetCIBAAuthenticationServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCIBAAuthenticationServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIBAAuthenticationServiceNotFound creates a GetCIBAAuthenticationServiceNotFound with default headers values
func NewGetCIBAAuthenticationServiceNotFound() *GetCIBAAuthenticationServiceNotFound {
	return &GetCIBAAuthenticationServiceNotFound{}
}

/* GetCIBAAuthenticationServiceNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetCIBAAuthenticationServiceNotFound struct {
	Payload *models.Error
}

func (o *GetCIBAAuthenticationServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/ciba-authentication-service][%d] getCIBAAuthenticationServiceNotFound  %+v", 404, o.Payload)
}
func (o *GetCIBAAuthenticationServiceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCIBAAuthenticationServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIBAAuthenticationServiceTooManyRequests creates a GetCIBAAuthenticationServiceTooManyRequests with default headers values
func NewGetCIBAAuthenticationServiceTooManyRequests() *GetCIBAAuthenticationServiceTooManyRequests {
	return &GetCIBAAuthenticationServiceTooManyRequests{}
}

/* GetCIBAAuthenticationServiceTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetCIBAAuthenticationServiceTooManyRequests struct {
	Payload *models.Error
}

func (o *GetCIBAAuthenticationServiceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/ciba-authentication-service][%d] getCIBAAuthenticationServiceTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCIBAAuthenticationServiceTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCIBAAuthenticationServiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
