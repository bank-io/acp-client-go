// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ListGatewaysReader is a Reader for the ListGateways structure.
type ListGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListGatewaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListGatewaysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListGatewaysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListGatewaysOK creates a ListGatewaysOK with default headers values
func NewListGatewaysOK() *ListGatewaysOK {
	return &ListGatewaysOK{}
}

/* ListGatewaysOK describes a response with status code 200, with default header values.

Gateways
*/
type ListGatewaysOK struct {
	Payload *models.Gateways
}

func (o *ListGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/gateways][%d] listGatewaysOK  %+v", 200, o.Payload)
}
func (o *ListGatewaysOK) GetPayload() *models.Gateways {
	return o.Payload
}

func (o *ListGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Gateways)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewaysUnauthorized creates a ListGatewaysUnauthorized with default headers values
func NewListGatewaysUnauthorized() *ListGatewaysUnauthorized {
	return &ListGatewaysUnauthorized{}
}

/* ListGatewaysUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListGatewaysUnauthorized struct {
	Payload *models.Error
}

func (o *ListGatewaysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/gateways][%d] listGatewaysUnauthorized  %+v", 401, o.Payload)
}
func (o *ListGatewaysUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListGatewaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewaysForbidden creates a ListGatewaysForbidden with default headers values
func NewListGatewaysForbidden() *ListGatewaysForbidden {
	return &ListGatewaysForbidden{}
}

/* ListGatewaysForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListGatewaysForbidden struct {
	Payload *models.Error
}

func (o *ListGatewaysForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/gateways][%d] listGatewaysForbidden  %+v", 403, o.Payload)
}
func (o *ListGatewaysForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListGatewaysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewaysTooManyRequests creates a ListGatewaysTooManyRequests with default headers values
func NewListGatewaysTooManyRequests() *ListGatewaysTooManyRequests {
	return &ListGatewaysTooManyRequests{}
}

/* ListGatewaysTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListGatewaysTooManyRequests struct {
	Payload *models.Error
}

func (o *ListGatewaysTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/gateways][%d] listGatewaysTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListGatewaysTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListGatewaysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}