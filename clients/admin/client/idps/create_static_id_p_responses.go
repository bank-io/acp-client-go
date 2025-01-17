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

// CreateStaticIDPReader is a Reader for the CreateStaticIDP structure.
type CreateStaticIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStaticIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStaticIDPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStaticIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateStaticIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateStaticIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateStaticIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateStaticIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateStaticIDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStaticIDPCreated creates a CreateStaticIDPCreated with default headers values
func NewCreateStaticIDPCreated() *CreateStaticIDPCreated {
	return &CreateStaticIDPCreated{}
}

/* CreateStaticIDPCreated describes a response with status code 201, with default header values.

StaticIDP
*/
type CreateStaticIDPCreated struct {
	Payload *models.StaticIDP
}

func (o *CreateStaticIDPCreated) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPCreated  %+v", 201, o.Payload)
}
func (o *CreateStaticIDPCreated) GetPayload() *models.StaticIDP {
	return o.Payload
}

func (o *CreateStaticIDPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StaticIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPBadRequest creates a CreateStaticIDPBadRequest with default headers values
func NewCreateStaticIDPBadRequest() *CreateStaticIDPBadRequest {
	return &CreateStaticIDPBadRequest{}
}

/* CreateStaticIDPBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type CreateStaticIDPBadRequest struct {
	Payload *models.Error
}

func (o *CreateStaticIDPBadRequest) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPBadRequest  %+v", 400, o.Payload)
}
func (o *CreateStaticIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPUnauthorized creates a CreateStaticIDPUnauthorized with default headers values
func NewCreateStaticIDPUnauthorized() *CreateStaticIDPUnauthorized {
	return &CreateStaticIDPUnauthorized{}
}

/* CreateStaticIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type CreateStaticIDPUnauthorized struct {
	Payload *models.Error
}

func (o *CreateStaticIDPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateStaticIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPForbidden creates a CreateStaticIDPForbidden with default headers values
func NewCreateStaticIDPForbidden() *CreateStaticIDPForbidden {
	return &CreateStaticIDPForbidden{}
}

/* CreateStaticIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type CreateStaticIDPForbidden struct {
	Payload *models.Error
}

func (o *CreateStaticIDPForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPForbidden  %+v", 403, o.Payload)
}
func (o *CreateStaticIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPNotFound creates a CreateStaticIDPNotFound with default headers values
func NewCreateStaticIDPNotFound() *CreateStaticIDPNotFound {
	return &CreateStaticIDPNotFound{}
}

/* CreateStaticIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type CreateStaticIDPNotFound struct {
	Payload *models.Error
}

func (o *CreateStaticIDPNotFound) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPNotFound  %+v", 404, o.Payload)
}
func (o *CreateStaticIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPUnprocessableEntity creates a CreateStaticIDPUnprocessableEntity with default headers values
func NewCreateStaticIDPUnprocessableEntity() *CreateStaticIDPUnprocessableEntity {
	return &CreateStaticIDPUnprocessableEntity{}
}

/* CreateStaticIDPUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type CreateStaticIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateStaticIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateStaticIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStaticIDPTooManyRequests creates a CreateStaticIDPTooManyRequests with default headers values
func NewCreateStaticIDPTooManyRequests() *CreateStaticIDPTooManyRequests {
	return &CreateStaticIDPTooManyRequests{}
}

/* CreateStaticIDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type CreateStaticIDPTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateStaticIDPTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/idps/static][%d] createStaticIdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateStaticIDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStaticIDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
