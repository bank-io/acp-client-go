// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/system/models"
)

// GetGatewayConfigurationReader is a Reader for the GetGatewayConfiguration structure.
type GetGatewayConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGatewayConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGatewayConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGatewayConfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayConfigurationOK creates a GetGatewayConfigurationOK with default headers values
func NewGetGatewayConfigurationOK() *GetGatewayConfigurationOK {
	return &GetGatewayConfigurationOK{}
}

/* GetGatewayConfigurationOK describes a response with status code 200, with default header values.

Gateway configuration
*/
type GetGatewayConfigurationOK struct {
	Payload *models.GatewayConfiguration
}

func (o *GetGatewayConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /gateways/configuration][%d] getGatewayConfigurationOK  %+v", 200, o.Payload)
}
func (o *GetGatewayConfigurationOK) GetPayload() *models.GatewayConfiguration {
	return o.Payload
}

func (o *GetGatewayConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationUnauthorized creates a GetGatewayConfigurationUnauthorized with default headers values
func NewGetGatewayConfigurationUnauthorized() *GetGatewayConfigurationUnauthorized {
	return &GetGatewayConfigurationUnauthorized{}
}

/* GetGatewayConfigurationUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetGatewayConfigurationUnauthorized struct {
	Payload *models.Error
}

func (o *GetGatewayConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /gateways/configuration][%d] getGatewayConfigurationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGatewayConfigurationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationForbidden creates a GetGatewayConfigurationForbidden with default headers values
func NewGetGatewayConfigurationForbidden() *GetGatewayConfigurationForbidden {
	return &GetGatewayConfigurationForbidden{}
}

/* GetGatewayConfigurationForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetGatewayConfigurationForbidden struct {
	Payload *models.Error
}

func (o *GetGatewayConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /gateways/configuration][%d] getGatewayConfigurationForbidden  %+v", 403, o.Payload)
}
func (o *GetGatewayConfigurationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationNotFound creates a GetGatewayConfigurationNotFound with default headers values
func NewGetGatewayConfigurationNotFound() *GetGatewayConfigurationNotFound {
	return &GetGatewayConfigurationNotFound{}
}

/* GetGatewayConfigurationNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetGatewayConfigurationNotFound struct {
	Payload *models.Error
}

func (o *GetGatewayConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /gateways/configuration][%d] getGatewayConfigurationNotFound  %+v", 404, o.Payload)
}
func (o *GetGatewayConfigurationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationTooManyRequests creates a GetGatewayConfigurationTooManyRequests with default headers values
func NewGetGatewayConfigurationTooManyRequests() *GetGatewayConfigurationTooManyRequests {
	return &GetGatewayConfigurationTooManyRequests{}
}

/* GetGatewayConfigurationTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetGatewayConfigurationTooManyRequests struct {
	Payload *models.Error
}

func (o *GetGatewayConfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /gateways/configuration][%d] getGatewayConfigurationTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetGatewayConfigurationTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayConfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
