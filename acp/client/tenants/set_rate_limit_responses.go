// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/acp/models"
)

// SetRateLimitReader is a Reader for the SetRateLimit structure.
type SetRateLimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRateLimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetRateLimitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetRateLimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetRateLimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetRateLimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetRateLimitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetRateLimitNoContent creates a SetRateLimitNoContent with default headers values
func NewSetRateLimitNoContent() *SetRateLimitNoContent {
	return &SetRateLimitNoContent{}
}

/* SetRateLimitNoContent describes a response with status code 204, with default header values.

custom rate limit has been saved
*/
type SetRateLimitNoContent struct {
}

func (o *SetRateLimitNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/admin/tenants/{tid}/rate-limits/{module}][%d] setRateLimitNoContent ", 204)
}

func (o *SetRateLimitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRateLimitUnauthorized creates a SetRateLimitUnauthorized with default headers values
func NewSetRateLimitUnauthorized() *SetRateLimitUnauthorized {
	return &SetRateLimitUnauthorized{}
}

/* SetRateLimitUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type SetRateLimitUnauthorized struct {
	Payload *models.Error
}

func (o *SetRateLimitUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/tenants/{tid}/rate-limits/{module}][%d] setRateLimitUnauthorized  %+v", 401, o.Payload)
}
func (o *SetRateLimitUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRateLimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRateLimitForbidden creates a SetRateLimitForbidden with default headers values
func NewSetRateLimitForbidden() *SetRateLimitForbidden {
	return &SetRateLimitForbidden{}
}

/* SetRateLimitForbidden describes a response with status code 403, with default header values.

HttpError
*/
type SetRateLimitForbidden struct {
	Payload *models.Error
}

func (o *SetRateLimitForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/tenants/{tid}/rate-limits/{module}][%d] setRateLimitForbidden  %+v", 403, o.Payload)
}
func (o *SetRateLimitForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRateLimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRateLimitNotFound creates a SetRateLimitNotFound with default headers values
func NewSetRateLimitNotFound() *SetRateLimitNotFound {
	return &SetRateLimitNotFound{}
}

/* SetRateLimitNotFound describes a response with status code 404, with default header values.

HttpError
*/
type SetRateLimitNotFound struct {
	Payload *models.Error
}

func (o *SetRateLimitNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/tenants/{tid}/rate-limits/{module}][%d] setRateLimitNotFound  %+v", 404, o.Payload)
}
func (o *SetRateLimitNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRateLimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRateLimitTooManyRequests creates a SetRateLimitTooManyRequests with default headers values
func NewSetRateLimitTooManyRequests() *SetRateLimitTooManyRequests {
	return &SetRateLimitTooManyRequests{}
}

/* SetRateLimitTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type SetRateLimitTooManyRequests struct {
	Payload *models.Error
}

func (o *SetRateLimitTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/admin/tenants/{tid}/rate-limits/{module}][%d] setRateLimitTooManyRequests  %+v", 429, o.Payload)
}
func (o *SetRateLimitTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRateLimitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
