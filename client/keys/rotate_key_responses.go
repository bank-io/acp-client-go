// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// RotateKeyReader is a Reader for the RotateKey structure.
type RotateKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRotateKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRotateKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRotateKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRotateKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRotateKeyOK creates a RotateKeyOK with default headers values
func NewRotateKeyOK() *RotateKeyOK {
	return &RotateKeyOK{}
}

/* RotateKeyOK describes a response with status code 200, with default header values.

ServerJWK
*/
type RotateKeyOK struct {
	Payload *models.ServerJWK
}

func (o *RotateKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/keys/rotate][%d] rotateKeyOK  %+v", 200, o.Payload)
}
func (o *RotateKeyOK) GetPayload() *models.ServerJWK {
	return o.Payload
}

func (o *RotateKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerJWK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateKeyBadRequest creates a RotateKeyBadRequest with default headers values
func NewRotateKeyBadRequest() *RotateKeyBadRequest {
	return &RotateKeyBadRequest{}
}

/* RotateKeyBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type RotateKeyBadRequest struct {
	Payload *models.Error
}

func (o *RotateKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/keys/rotate][%d] rotateKeyBadRequest  %+v", 400, o.Payload)
}
func (o *RotateKeyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateKeyUnauthorized creates a RotateKeyUnauthorized with default headers values
func NewRotateKeyUnauthorized() *RotateKeyUnauthorized {
	return &RotateKeyUnauthorized{}
}

/* RotateKeyUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RotateKeyUnauthorized struct {
	Payload *models.Error
}

func (o *RotateKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/keys/rotate][%d] rotateKeyUnauthorized  %+v", 401, o.Payload)
}
func (o *RotateKeyUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateKeyForbidden creates a RotateKeyForbidden with default headers values
func NewRotateKeyForbidden() *RotateKeyForbidden {
	return &RotateKeyForbidden{}
}

/* RotateKeyForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RotateKeyForbidden struct {
	Payload *models.Error
}

func (o *RotateKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/keys/rotate][%d] rotateKeyForbidden  %+v", 403, o.Payload)
}
func (o *RotateKeyForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateKeyNotFound creates a RotateKeyNotFound with default headers values
func NewRotateKeyNotFound() *RotateKeyNotFound {
	return &RotateKeyNotFound{}
}

/* RotateKeyNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RotateKeyNotFound struct {
	Payload *models.Error
}

func (o *RotateKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/keys/rotate][%d] rotateKeyNotFound  %+v", 404, o.Payload)
}
func (o *RotateKeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}