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

// UpdateGatewayReader is a Reader for the UpdateGateway structure.
type UpdateGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGatewayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateGatewayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateGatewayUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateGatewayTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGatewayOK creates a UpdateGatewayOK with default headers values
func NewUpdateGatewayOK() *UpdateGatewayOK {
	return &UpdateGatewayOK{}
}

/* UpdateGatewayOK describes a response with status code 200, with default header values.

Gateway
*/
type UpdateGatewayOK struct {
	Payload *models.Gateway
}

func (o *UpdateGatewayOK) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayOK  %+v", 200, o.Payload)
}
func (o *UpdateGatewayOK) GetPayload() *models.Gateway {
	return o.Payload
}

func (o *UpdateGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Gateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayBadRequest creates a UpdateGatewayBadRequest with default headers values
func NewUpdateGatewayBadRequest() *UpdateGatewayBadRequest {
	return &UpdateGatewayBadRequest{}
}

/* UpdateGatewayBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateGatewayBadRequest struct {
	Payload *models.Error
}

func (o *UpdateGatewayBadRequest) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateGatewayBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayUnauthorized creates a UpdateGatewayUnauthorized with default headers values
func NewUpdateGatewayUnauthorized() *UpdateGatewayUnauthorized {
	return &UpdateGatewayUnauthorized{}
}

/* UpdateGatewayUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateGatewayUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateGatewayUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateGatewayUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayForbidden creates a UpdateGatewayForbidden with default headers values
func NewUpdateGatewayForbidden() *UpdateGatewayForbidden {
	return &UpdateGatewayForbidden{}
}

/* UpdateGatewayForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateGatewayForbidden struct {
	Payload *models.Error
}

func (o *UpdateGatewayForbidden) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayForbidden  %+v", 403, o.Payload)
}
func (o *UpdateGatewayForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayNotFound creates a UpdateGatewayNotFound with default headers values
func NewUpdateGatewayNotFound() *UpdateGatewayNotFound {
	return &UpdateGatewayNotFound{}
}

/* UpdateGatewayNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateGatewayNotFound struct {
	Payload *models.Error
}

func (o *UpdateGatewayNotFound) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayNotFound  %+v", 404, o.Payload)
}
func (o *UpdateGatewayNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayUnprocessableEntity creates a UpdateGatewayUnprocessableEntity with default headers values
func NewUpdateGatewayUnprocessableEntity() *UpdateGatewayUnprocessableEntity {
	return &UpdateGatewayUnprocessableEntity{}
}

/* UpdateGatewayUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateGatewayUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateGatewayUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateGatewayUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayTooManyRequests creates a UpdateGatewayTooManyRequests with default headers values
func NewUpdateGatewayTooManyRequests() *UpdateGatewayTooManyRequests {
	return &UpdateGatewayTooManyRequests{}
}

/* UpdateGatewayTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateGatewayTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateGatewayTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /gateways/{gw}][%d] updateGatewayTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateGatewayTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGatewayTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
