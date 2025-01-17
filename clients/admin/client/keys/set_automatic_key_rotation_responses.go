// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// SetAutomaticKeyRotationReader is a Reader for the SetAutomaticKeyRotation structure.
type SetAutomaticKeyRotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAutomaticKeyRotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAutomaticKeyRotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetAutomaticKeyRotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetAutomaticKeyRotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetAutomaticKeyRotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetAutomaticKeyRotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetAutomaticKeyRotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetAutomaticKeyRotationOK creates a SetAutomaticKeyRotationOK with default headers values
func NewSetAutomaticKeyRotationOK() *SetAutomaticKeyRotationOK {
	return &SetAutomaticKeyRotationOK{}
}

/* SetAutomaticKeyRotationOK describes a response with status code 200, with default header values.

Automation key rotation
*/
type SetAutomaticKeyRotationOK struct {
	Payload *models.AutomaticKeyRotation
}

func (o *SetAutomaticKeyRotationOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationOK  %+v", 200, o.Payload)
}
func (o *SetAutomaticKeyRotationOK) GetPayload() *models.AutomaticKeyRotation {
	return o.Payload
}

func (o *SetAutomaticKeyRotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutomaticKeyRotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAutomaticKeyRotationBadRequest creates a SetAutomaticKeyRotationBadRequest with default headers values
func NewSetAutomaticKeyRotationBadRequest() *SetAutomaticKeyRotationBadRequest {
	return &SetAutomaticKeyRotationBadRequest{}
}

/* SetAutomaticKeyRotationBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type SetAutomaticKeyRotationBadRequest struct {
	Payload *models.Error
}

func (o *SetAutomaticKeyRotationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationBadRequest  %+v", 400, o.Payload)
}
func (o *SetAutomaticKeyRotationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAutomaticKeyRotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAutomaticKeyRotationUnauthorized creates a SetAutomaticKeyRotationUnauthorized with default headers values
func NewSetAutomaticKeyRotationUnauthorized() *SetAutomaticKeyRotationUnauthorized {
	return &SetAutomaticKeyRotationUnauthorized{}
}

/* SetAutomaticKeyRotationUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type SetAutomaticKeyRotationUnauthorized struct {
	Payload *models.Error
}

func (o *SetAutomaticKeyRotationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationUnauthorized  %+v", 401, o.Payload)
}
func (o *SetAutomaticKeyRotationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAutomaticKeyRotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAutomaticKeyRotationForbidden creates a SetAutomaticKeyRotationForbidden with default headers values
func NewSetAutomaticKeyRotationForbidden() *SetAutomaticKeyRotationForbidden {
	return &SetAutomaticKeyRotationForbidden{}
}

/* SetAutomaticKeyRotationForbidden describes a response with status code 403, with default header values.

HttpError
*/
type SetAutomaticKeyRotationForbidden struct {
	Payload *models.Error
}

func (o *SetAutomaticKeyRotationForbidden) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationForbidden  %+v", 403, o.Payload)
}
func (o *SetAutomaticKeyRotationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAutomaticKeyRotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAutomaticKeyRotationNotFound creates a SetAutomaticKeyRotationNotFound with default headers values
func NewSetAutomaticKeyRotationNotFound() *SetAutomaticKeyRotationNotFound {
	return &SetAutomaticKeyRotationNotFound{}
}

/* SetAutomaticKeyRotationNotFound describes a response with status code 404, with default header values.

HttpError
*/
type SetAutomaticKeyRotationNotFound struct {
	Payload *models.Error
}

func (o *SetAutomaticKeyRotationNotFound) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationNotFound  %+v", 404, o.Payload)
}
func (o *SetAutomaticKeyRotationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAutomaticKeyRotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAutomaticKeyRotationTooManyRequests creates a SetAutomaticKeyRotationTooManyRequests with default headers values
func NewSetAutomaticKeyRotationTooManyRequests() *SetAutomaticKeyRotationTooManyRequests {
	return &SetAutomaticKeyRotationTooManyRequests{}
}

/* SetAutomaticKeyRotationTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type SetAutomaticKeyRotationTooManyRequests struct {
	Payload *models.Error
}

func (o *SetAutomaticKeyRotationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/keys/automatic-key-rotation][%d] setAutomaticKeyRotationTooManyRequests  %+v", 429, o.Payload)
}
func (o *SetAutomaticKeyRotationTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAutomaticKeyRotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
