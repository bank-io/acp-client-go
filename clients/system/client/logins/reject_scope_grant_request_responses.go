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

// RejectScopeGrantRequestReader is a Reader for the RejectScopeGrantRequest structure.
type RejectScopeGrantRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectScopeGrantRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectScopeGrantRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRejectScopeGrantRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRejectScopeGrantRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRejectScopeGrantRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRejectScopeGrantRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRejectScopeGrantRequestOK creates a RejectScopeGrantRequestOK with default headers values
func NewRejectScopeGrantRequestOK() *RejectScopeGrantRequestOK {
	return &RejectScopeGrantRequestOK{}
}

/* RejectScopeGrantRequestOK describes a response with status code 200, with default header values.

Scope grant rejected
*/
type RejectScopeGrantRequestOK struct {
	Payload *models.ScopeGrantRejected
}

func (o *RejectScopeGrantRequestOK) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/reject][%d] rejectScopeGrantRequestOK  %+v", 200, o.Payload)
}
func (o *RejectScopeGrantRequestOK) GetPayload() *models.ScopeGrantRejected {
	return o.Payload
}

func (o *RejectScopeGrantRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopeGrantRejected)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectScopeGrantRequestUnauthorized creates a RejectScopeGrantRequestUnauthorized with default headers values
func NewRejectScopeGrantRequestUnauthorized() *RejectScopeGrantRequestUnauthorized {
	return &RejectScopeGrantRequestUnauthorized{}
}

/* RejectScopeGrantRequestUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RejectScopeGrantRequestUnauthorized struct {
	Payload *models.Error
}

func (o *RejectScopeGrantRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/reject][%d] rejectScopeGrantRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *RejectScopeGrantRequestUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectScopeGrantRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectScopeGrantRequestForbidden creates a RejectScopeGrantRequestForbidden with default headers values
func NewRejectScopeGrantRequestForbidden() *RejectScopeGrantRequestForbidden {
	return &RejectScopeGrantRequestForbidden{}
}

/* RejectScopeGrantRequestForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RejectScopeGrantRequestForbidden struct {
	Payload *models.Error
}

func (o *RejectScopeGrantRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/reject][%d] rejectScopeGrantRequestForbidden  %+v", 403, o.Payload)
}
func (o *RejectScopeGrantRequestForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectScopeGrantRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectScopeGrantRequestNotFound creates a RejectScopeGrantRequestNotFound with default headers values
func NewRejectScopeGrantRequestNotFound() *RejectScopeGrantRequestNotFound {
	return &RejectScopeGrantRequestNotFound{}
}

/* RejectScopeGrantRequestNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RejectScopeGrantRequestNotFound struct {
	Payload *models.Error
}

func (o *RejectScopeGrantRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/reject][%d] rejectScopeGrantRequestNotFound  %+v", 404, o.Payload)
}
func (o *RejectScopeGrantRequestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectScopeGrantRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectScopeGrantRequestTooManyRequests creates a RejectScopeGrantRequestTooManyRequests with default headers values
func NewRejectScopeGrantRequestTooManyRequests() *RejectScopeGrantRequestTooManyRequests {
	return &RejectScopeGrantRequestTooManyRequests{}
}

/* RejectScopeGrantRequestTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type RejectScopeGrantRequestTooManyRequests struct {
	Payload *models.Error
}

func (o *RejectScopeGrantRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /scope-grants/{login}/reject][%d] rejectScopeGrantRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *RejectScopeGrantRequestTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectScopeGrantRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
