// Code generated by go-swagger; DO NOT EDIT.

package idps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// CreateCustomIDPReader is a Reader for the CreateCustomIDP structure.
type CreateCustomIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCustomIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCustomIDPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCustomIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCustomIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCustomIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCustomIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCustomIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateCustomIDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCustomIDPCreated creates a CreateCustomIDPCreated with default headers values
func NewCreateCustomIDPCreated() *CreateCustomIDPCreated {
	return &CreateCustomIDPCreated{}
}

/* CreateCustomIDPCreated describes a response with status code 201, with default header values.

CustomIDP
*/
type CreateCustomIDPCreated struct {
	Payload *models.CustomIDP
}

func (o *CreateCustomIDPCreated) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPCreated  %+v", 201, o.Payload)
}
func (o *CreateCustomIDPCreated) GetPayload() *models.CustomIDP {
	return o.Payload
}

func (o *CreateCustomIDPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPBadRequest creates a CreateCustomIDPBadRequest with default headers values
func NewCreateCustomIDPBadRequest() *CreateCustomIDPBadRequest {
	return &CreateCustomIDPBadRequest{}
}

/* CreateCustomIDPBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type CreateCustomIDPBadRequest struct {
	Payload *models.Error
}

func (o *CreateCustomIDPBadRequest) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPBadRequest  %+v", 400, o.Payload)
}
func (o *CreateCustomIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPUnauthorized creates a CreateCustomIDPUnauthorized with default headers values
func NewCreateCustomIDPUnauthorized() *CreateCustomIDPUnauthorized {
	return &CreateCustomIDPUnauthorized{}
}

/* CreateCustomIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type CreateCustomIDPUnauthorized struct {
	Payload *models.Error
}

func (o *CreateCustomIDPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateCustomIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPForbidden creates a CreateCustomIDPForbidden with default headers values
func NewCreateCustomIDPForbidden() *CreateCustomIDPForbidden {
	return &CreateCustomIDPForbidden{}
}

/* CreateCustomIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type CreateCustomIDPForbidden struct {
	Payload *models.Error
}

func (o *CreateCustomIDPForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPForbidden  %+v", 403, o.Payload)
}
func (o *CreateCustomIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPNotFound creates a CreateCustomIDPNotFound with default headers values
func NewCreateCustomIDPNotFound() *CreateCustomIDPNotFound {
	return &CreateCustomIDPNotFound{}
}

/* CreateCustomIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type CreateCustomIDPNotFound struct {
	Payload *models.Error
}

func (o *CreateCustomIDPNotFound) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPNotFound  %+v", 404, o.Payload)
}
func (o *CreateCustomIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPUnprocessableEntity creates a CreateCustomIDPUnprocessableEntity with default headers values
func NewCreateCustomIDPUnprocessableEntity() *CreateCustomIDPUnprocessableEntity {
	return &CreateCustomIDPUnprocessableEntity{}
}

/* CreateCustomIDPUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type CreateCustomIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateCustomIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateCustomIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomIDPTooManyRequests creates a CreateCustomIDPTooManyRequests with default headers values
func NewCreateCustomIDPTooManyRequests() *CreateCustomIDPTooManyRequests {
	return &CreateCustomIDPTooManyRequests{}
}

/* CreateCustomIDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type CreateCustomIDPTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateCustomIDPTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/custom][%d] createCustomIdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateCustomIDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomIDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
