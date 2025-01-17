// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/system/models"
)

// AcceptScopeGrantRequestReader is a Reader for the AcceptScopeGrantRequest structure.
type AcceptScopeGrantRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptScopeGrantRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptScopeGrantRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptScopeGrantRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptScopeGrantRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptScopeGrantRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAcceptScopeGrantRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptScopeGrantRequestOK creates a AcceptScopeGrantRequestOK with default headers values
func NewAcceptScopeGrantRequestOK() *AcceptScopeGrantRequestOK {
	return &AcceptScopeGrantRequestOK{}
}

/* AcceptScopeGrantRequestOK describes a response with status code 200, with default header values.

Scope grant accepted
*/
type AcceptScopeGrantRequestOK struct {
	Payload *models.ScopeGrantAccepted
}

func (o *AcceptScopeGrantRequestOK) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/accept][%d] acceptScopeGrantRequestOK  %+v", 200, o.Payload)
}
func (o *AcceptScopeGrantRequestOK) GetPayload() *models.ScopeGrantAccepted {
	return o.Payload
}

func (o *AcceptScopeGrantRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopeGrantAccepted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptScopeGrantRequestUnauthorized creates a AcceptScopeGrantRequestUnauthorized with default headers values
func NewAcceptScopeGrantRequestUnauthorized() *AcceptScopeGrantRequestUnauthorized {
	return &AcceptScopeGrantRequestUnauthorized{}
}

/* AcceptScopeGrantRequestUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type AcceptScopeGrantRequestUnauthorized struct {
	Payload *models.Error
}

func (o *AcceptScopeGrantRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/accept][%d] acceptScopeGrantRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *AcceptScopeGrantRequestUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptScopeGrantRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptScopeGrantRequestForbidden creates a AcceptScopeGrantRequestForbidden with default headers values
func NewAcceptScopeGrantRequestForbidden() *AcceptScopeGrantRequestForbidden {
	return &AcceptScopeGrantRequestForbidden{}
}

/* AcceptScopeGrantRequestForbidden describes a response with status code 403, with default header values.

HttpError
*/
type AcceptScopeGrantRequestForbidden struct {
	Payload *models.Error
}

func (o *AcceptScopeGrantRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/accept][%d] acceptScopeGrantRequestForbidden  %+v", 403, o.Payload)
}
func (o *AcceptScopeGrantRequestForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptScopeGrantRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptScopeGrantRequestNotFound creates a AcceptScopeGrantRequestNotFound with default headers values
func NewAcceptScopeGrantRequestNotFound() *AcceptScopeGrantRequestNotFound {
	return &AcceptScopeGrantRequestNotFound{}
}

/* AcceptScopeGrantRequestNotFound describes a response with status code 404, with default header values.

HttpError
*/
type AcceptScopeGrantRequestNotFound struct {
	Payload *models.Error
}

func (o *AcceptScopeGrantRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/accept][%d] acceptScopeGrantRequestNotFound  %+v", 404, o.Payload)
}
func (o *AcceptScopeGrantRequestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptScopeGrantRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptScopeGrantRequestTooManyRequests creates a AcceptScopeGrantRequestTooManyRequests with default headers values
func NewAcceptScopeGrantRequestTooManyRequests() *AcceptScopeGrantRequestTooManyRequests {
	return &AcceptScopeGrantRequestTooManyRequests{}
}

/* AcceptScopeGrantRequestTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type AcceptScopeGrantRequestTooManyRequests struct {
	Payload *models.Error
}

func (o *AcceptScopeGrantRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/accept][%d] acceptScopeGrantRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *AcceptScopeGrantRequestTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptScopeGrantRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
