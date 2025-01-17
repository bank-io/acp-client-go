// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// RotateClientSecretReader is a Reader for the RotateClientSecret structure.
type RotateClientSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateClientSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateClientSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRotateClientSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRotateClientSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRotateClientSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRotateClientSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRotateClientSecretTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRotateClientSecretOK creates a RotateClientSecretOK with default headers values
func NewRotateClientSecretOK() *RotateClientSecretOK {
	return &RotateClientSecretOK{}
}

/* RotateClientSecretOK describes a response with status code 200, with default header values.

Rotate client secret response
*/
type RotateClientSecretOK struct {
	Payload *models.RotateClientSecretResponse
}

func (o *RotateClientSecretOK) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretOK  %+v", 200, o.Payload)
}
func (o *RotateClientSecretOK) GetPayload() *models.RotateClientSecretResponse {
	return o.Payload
}

func (o *RotateClientSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RotateClientSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateClientSecretBadRequest creates a RotateClientSecretBadRequest with default headers values
func NewRotateClientSecretBadRequest() *RotateClientSecretBadRequest {
	return &RotateClientSecretBadRequest{}
}

/* RotateClientSecretBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type RotateClientSecretBadRequest struct {
	Payload *models.Error
}

func (o *RotateClientSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretBadRequest  %+v", 400, o.Payload)
}
func (o *RotateClientSecretBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateClientSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateClientSecretUnauthorized creates a RotateClientSecretUnauthorized with default headers values
func NewRotateClientSecretUnauthorized() *RotateClientSecretUnauthorized {
	return &RotateClientSecretUnauthorized{}
}

/* RotateClientSecretUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RotateClientSecretUnauthorized struct {
	Payload *models.Error
}

func (o *RotateClientSecretUnauthorized) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretUnauthorized  %+v", 401, o.Payload)
}
func (o *RotateClientSecretUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateClientSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateClientSecretForbidden creates a RotateClientSecretForbidden with default headers values
func NewRotateClientSecretForbidden() *RotateClientSecretForbidden {
	return &RotateClientSecretForbidden{}
}

/* RotateClientSecretForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RotateClientSecretForbidden struct {
	Payload *models.Error
}

func (o *RotateClientSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretForbidden  %+v", 403, o.Payload)
}
func (o *RotateClientSecretForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateClientSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateClientSecretNotFound creates a RotateClientSecretNotFound with default headers values
func NewRotateClientSecretNotFound() *RotateClientSecretNotFound {
	return &RotateClientSecretNotFound{}
}

/* RotateClientSecretNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RotateClientSecretNotFound struct {
	Payload *models.Error
}

func (o *RotateClientSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretNotFound  %+v", 404, o.Payload)
}
func (o *RotateClientSecretNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateClientSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateClientSecretTooManyRequests creates a RotateClientSecretTooManyRequests with default headers values
func NewRotateClientSecretTooManyRequests() *RotateClientSecretTooManyRequests {
	return &RotateClientSecretTooManyRequests{}
}

/* RotateClientSecretTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type RotateClientSecretTooManyRequests struct {
	Payload *models.Error
}

func (o *RotateClientSecretTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /clients/{cid}/rotateSecret][%d] rotateClientSecretTooManyRequests  %+v", 429, o.Payload)
}
func (o *RotateClientSecretTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateClientSecretTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
